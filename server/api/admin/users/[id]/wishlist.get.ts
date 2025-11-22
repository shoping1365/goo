import type { H3Event } from 'h3'
import { defineEventHandler, createError } from 'h3'
import { z } from 'zod'
import { getDatabase } from '~/server/api/_utils/database'

declare const requireAdminAuth: (event: H3Event) => Promise<{ id: number | string; role: string } | null>

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

    // دریافت لیست علاقه‌مندی‌ها از دیتابیس
    const db = await getDatabase()
    
    // شمارش کل رکوردها
    const totalCount = await db.query(`
      SELECT COUNT(*) as count 
      FROM user_collection_items uci
      LEFT JOIN user_collections uc ON uc.id = uci.collection_id
      WHERE uc.user_id = $1 AND uc.is_default = true
    `, [userId]) as unknown as { count: string }[]
    
    const total = parseInt(totalCount[0].count)

    // دریافت رکوردها با pagination
    const result = await db.query(`
      SELECT 
        uci.id,
        uci.product_id,
        uci.created_at,
        p.name as product_name,
        p.price as product_price,
        p.sale_price,
        p.sale_start_at,
        p.sale_end_at,
        p.image as product_image,
        p.sku as product_sku,
        CASE
          WHEN p.sale_price IS NOT NULL AND p.sale_price > 0
               AND (p.sale_start_at IS NULL OR p.sale_start_at <= NOW())
               AND (p.sale_end_at IS NULL OR p.sale_end_at >= NOW())
            THEN p.sale_price
          ELSE p.price
        END AS effective_price
      FROM user_collection_items uci
      LEFT JOIN user_collections uc ON uc.id = uci.collection_id
      LEFT JOIN products p ON p.id = uci.product_id
      WHERE uc.user_id = $1 AND uc.is_default = true
      ORDER BY uci.created_at DESC
      LIMIT $2 OFFSET $3
    `, [userId, limitNum, offset]) as unknown as Array<{ id: number; product_id: number; created_at: string; product_name: string | null; product_price: number; sale_price: number | null; sale_start_at: string | null; sale_end_at: string | null; product_image: string | null; product_sku: string | null; effective_price: number }>

    const wishlistItems = result.map((row) => ({
      id: row.id,
      productId: row.product_id,
      productName: row.product_name || `محصول ${row.product_id}`,
      productPrice: row.effective_price || row.product_price,
      productImage: row.product_image || '/default-product.svg',
      productSku: row.product_sku,
      createdAt: row.created_at,
    }))

    return {
      success: true,
      data: {
        wishlistItems,
        pagination: {
          page: pageNum,
          limit: limitNum,
          total,
          totalPages: Math.ceil(total / limitNum)
        }
      }
    }

  } catch (error) {
    console.error('خطا در دریافت لیست علاقه‌مندی‌ها:', error)
    
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت لیست علاقه‌مندی‌ها'
    })
  }
})

