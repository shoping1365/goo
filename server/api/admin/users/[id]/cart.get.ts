import type { H3Event } from 'h3'
import { createError, defineEventHandler } from 'h3'
import { z } from 'zod'
import { getDatabase } from '~/server/api/_utils/database'
import { getAllCarts } from '~/server/api/cart/data'

declare const requireAdminAuth: (event: H3Event) => Promise<{ id: number | string; role: string } | null>

interface CartItemRow {
  id: number
  cart_id: number
  product_id: number
  quantity: number
  unit_price: number
  final_price: number
  added_at: string
  product_name: string
  product_image: string | null
  product_sku: string | null
}

interface CartItemResult {
  id: number
  cartId: number | null
  productId: number
  quantity: number
  unitPrice: number
  finalPrice: number
  productName: string
  productImage: string | null
  productSku: string | null
  addedAt: string
}

const querySchema = z.object({
  page: z.string().optional().default('1'),
  limit: z.string().optional().default('20'),
})

export default defineEventHandler(async (event) => {
  try {
    // احراز هویت ادمین
    const adminUser = await requireAdminAuth(event)
    if (!adminUser) {
      throw createError({
        statusCode: 401,
        message: 'احراز هویت ادمین الزامی است'
      })
    }

    // دریافت پارامترهای URL
    const userId = getRouterParam(event, 'id')
    if (!userId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه کاربر الزامی است'
      })
    }

    // اعتبارسنجی query parameters
    const query = await getQuery(event)
    const { page, limit } = querySchema.parse(query)

    const pageNum = parseInt(page)
    const limitNum = parseInt(limit)
    const offset = (pageNum - 1) * limitNum

    // دریافت سبد خرید از دیتابیس
    const db = await getDatabase()

    // شمارش کل رکوردها برای pagination
    await db.query(`
      SELECT COUNT(*) as count 
      FROM cart_items ci
      JOIN carts c ON ci.cart_id = c.id
      WHERE c.user_id = $1 AND ci.deleted_at IS NULL AND ci.is_next_purchase = false
    `, [userId])

    // دریافت رکوردها با pagination
    const result = await db.query(`
      SELECT 
        ci.id,
        ci.cart_id,
        ci.product_id,
        ci.quantity,
        ci.unit_price,
        ci.final_price,
        ci.added_at,
        p.name as product_name,
        p.image as product_image,
        p.sku as product_sku
      FROM cart_items ci
      JOIN carts c ON ci.cart_id = c.id
      JOIN products p ON ci.product_id = p.id
      WHERE c.user_id = $1 AND ci.deleted_at IS NULL AND ci.is_next_purchase = false
      ORDER BY ci.added_at DESC
      LIMIT $2 OFFSET $3
    `, [userId, limitNum, offset]) as unknown as CartItemRow[]

    let cartItems: CartItemResult[] = result.map((row) => ({
      id: row.id,
      cartId: row.cart_id,
      productId: row.product_id,
      quantity: row.quantity,
      unitPrice: row.unit_price,
      finalPrice: row.final_price,
      productName: row.product_name,
      productImage: row.product_image,
      productSku: row.product_sku,
      addedAt: row.added_at,
    }))

    // اگر در database چیزی پیدا نشد، از session storage بخوان
    // توجه: session storage فعلاً user_id ذخیره نمی‌کند، پس همه session ها را نشان می‌دهیم
    if (cartItems.length === 0) {
      try {
        const allCarts = getAllCarts()
        const sessionCartItems = []

        // پیدا کردن همه آیتم‌های سبد خرید جاری در همه session ها
        for (const [, cart] of Object.entries(allCarts)) {
          const cartData = cart as any
          if (cartData && cartData.items && Array.isArray(cartData.items)) {
            for (const item of cartData.items) {
              // چک کنیم که is_next تنظیم نشده یا false است
              if (!item.is_next) {
                // دریافت اطلاعات محصول از database
                const productInfo = await db.query(`
                  SELECT id, name, image, sku, price, sale_price 
                  FROM products WHERE id = $1
                `, [item.product_id]) as any[]

                if (productInfo && productInfo.length > 0) {
                  const product = productInfo[0]
                  sessionCartItems.push({
                    id: item.id,
                    cartId: null,
                    productId: item.product_id,
                    quantity: item.quantity,
                    unitPrice: item.unit_price || item.price || product.price,
                    finalPrice: item.final_price || (item.price || product.price) * item.quantity,
                    productName: item.name || product.name,
                    productImage: item.image || product.image,
                    productSku: item.sku || product.sku,
                    addedAt: item.added_at || new Date().toISOString(),
                  })
                }
              }
            }
          }
        }

        cartItems = sessionCartItems
      } catch (sessionError) {
        console.error('خطا در خواندن session cart:', sessionError)
      }
    }

    return {
      success: true,
      data: {
        cartItems,
        pagination: {
          page: pageNum,
          limit: limitNum,
          total: cartItems.length,
          totalPages: Math.ceil(cartItems.length / limitNum)
        }
      }
    }

  } catch (error) {
    console.error('خطا در دریافت سبد خرید:', error)

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت سبد خرید'
    })
  }
})

