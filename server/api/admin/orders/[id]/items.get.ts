import type { H3Event } from 'h3'
import { createError, defineEventHandler, getRouterParam } from 'h3'

interface OrderItemRow {
  id: number
  order_id: number
  product_id: number
  product_name: string
  product_image: string | null
  product_sku: string
  unit_price: number
  quantity: number
  total_price: number
  final_price: number
  created_at: string
  updated_at: string
}

interface User {
  id: number | string
  [key: string]: unknown
}

interface Database {
  query: (sql: string, params?: unknown[]) => Promise<{ rows: unknown[]; rowCount: number }>
}

declare const requireAdminAuth: (event: H3Event) => Promise<User | null>
declare const getDatabase: () => Promise<Database>

export default defineEventHandler(async (event: H3Event) => {
  try {
    console.log('🔍 API آیتم‌های سفارش فراخوانی شد');

    // احراز هویت ادمین
    const adminUser = await requireAdminAuth(event)
    console.log('👤 کاربر ادمین:', adminUser ? 'موجود' : 'ناموجود');

    if (!adminUser) {
      console.log('❌ احراز هویت ناموفق');
      throw createError({
        statusCode: 401,
        message: 'احراز هویت ادمین الزامی است'
      })
    }

    // دریافت پارامترهای URL
    const orderId = getRouterParam(event, 'id')
    console.log('🆔 شناسه سفارش:', orderId);

    if (!orderId) {
      console.log('❌ شناسه سفارش موجود نیست');
      throw createError({
        statusCode: 400,
        message: 'شناسه سفارش الزامی است'
      })
    }

    // دریافت آیتم‌های سفارش از دیتابیس
    const db = await getDatabase()

    const result = await db.query(`
      SELECT 
        id,
        order_id,
        product_id,
        COALESCE(product_name, 'نامشخص') as product_name,
        product_image,
        COALESCE(product_sku, 'نامشخص') as product_sku,
        unit_price,
        quantity,
        COALESCE(total_price, unit_price * quantity) as total_price,
        final_price,
        created_at,
        updated_at
      FROM order_items 
      WHERE order_id = $1
      ORDER BY created_at ASC
    `, [orderId]) as unknown as OrderItemRow[]

    console.log('Order items raw result:', JSON.stringify(result, null, 2));

    const items = result.map((row) => {
      const item = {
        id: row.id,
        orderId: row.order_id,
        productId: row.product_id,
        name: row.product_name,
        image: row.product_image,
        sku: row.product_sku,
        unitPrice: row.unit_price,
        quantity: row.quantity,
        totalPrice: row.total_price,
        finalPrice: row.final_price,
        createdAt: row.created_at,
        updatedAt: row.updated_at,
      };
      console.log('Mapped order item:', item);
      return item;
    })

    return {
      success: true,
      data: {
        items,
        count: items.length
      }
    }

  } catch (error) {
    console.error('خطا در دریافت آیتم‌های سفارش:', error)

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت آیتم‌های سفارش'
    })
  }
})
