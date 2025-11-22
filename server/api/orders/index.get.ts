import type { H3Event } from 'h3'
import { defineEventHandler } from 'h3'
import { requireAuth } from '~/server/utils/requireAuth'
import { getDatabase } from '../_utils/database'

interface OrderRow {
  id: number
  order_number?: string
  total_amount?: number
  final_amount?: number
  payment_status?: string
  status?: string
  created_at?: string
  updated_at?: string
  product_id?: number
  quantity?: number
  item_price?: number
  product_name?: string
  product_name_from_products?: string
  product_image?: string
}

export default defineEventHandler(async (event: H3Event) => {
  try {
    const user = await requireAuth(event)
    event.context.user = user
    const userId = user.id

    // اتصال به پایگاه‌داده
    const db = await getDatabase()

    // اجرای کوئری سفارش‌ها از پایگاه‌داده
    const ordersResult = await db.query(`
      SELECT 
        o.id,
        o.order_number,
        o.total_amount,
        o.final_amount,
        o.payment_status,
        o.status,
        o.created_at,
        o.updated_at,
        oi.product_id,
        oi.quantity,
        oi.unit_price as item_price,
        oi.product_name,
        p.name as product_name_from_products,
        p.image as product_image
      FROM orders o
      LEFT JOIN order_items oi ON oi.order_id = o.id
      LEFT JOIN products p ON p.id = oi.product_id
      WHERE o.user_id = $1
      ORDER BY o.created_at DESC
    `, [userId])
    const orders = ordersResult as unknown as OrderRow[]


    // ساخت Map برای گروه‌بندی سفارش‌ها
    const ordersMap = new Map()

    orders.forEach((row: OrderRow) => {
      if (!ordersMap.has(row.id)) {
        ordersMap.set(row.id, {
          id: row.id,
          order_number: row.order_number || `#${row.id}`,
          total_amount: row.total_amount || 0,
          final_amount: row.final_amount || row.total_amount || 0,
          payment_status: row.payment_status || 'pending',
          status: row.status || 'pending',
          created_at: row.created_at,
          updated_at: row.updated_at,
          items: []
        })
      }

      if (row.product_id) {
        ordersMap.get(row.id).items.push({
          id: row.product_id,
          product_id: row.product_id,
          product_name: row.product_name || row.product_name_from_products || `محصول ${row.product_id}`,
          quantity: row.quantity || 1,
          price: row.item_price || 0,
          product_image: row.product_image
        })
      }
    })

    // تبدیل Map به آرایه نهایی
    const formattedOrders = Array.from(ordersMap.values())


    return {
      success: true,
      data: formattedOrders
    }
  } catch (error: unknown) {
    const errorWithDetails = error as { message?: string; stack?: string; name?: string }
    console.error('[orders.index] خطا در دریافت سفارش‌ها:', error)
    console.error('[orders.index] جزئیات خطا:', {
      message: errorWithDetails?.message,
      stack: errorWithDetails?.stack,
      name: errorWithDetails?.name
    })
    return {
      success: false,
      data: [],
      message: 'بازیابی سفارش‌ها با خطا روبه‌رو شد'
    }
  }
})

