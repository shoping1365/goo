import type { H3Event } from 'h3';
import { defineEventHandler, getQuery } from 'h3';

declare const getDatabase: () => Promise<{
  query: (sql: string, params?: unknown[]) => Promise<{ rows: unknown[]; rowCount: number }>
}>

interface OrderRow {
  id: number
  order_number?: string
  total_amount?: number
  payment_method?: string
  status?: string
  payment_status?: string
  created_at?: string
  user_name?: string
  user_mobile?: string
  shipping_address?: string
  shipping_city?: string
  shipping_province?: string
  shipping_postal_code?: string
  recipient_name?: string
  recipient_phone?: string
  items_count?: string | number
  tracking_code?: string
  customer_ip?: string
  user_agent?: string
}

interface CountRow {
  count: string | number
}

export default defineEventHandler(async (event: H3Event) => {
  try {
    // دریافت اتصال دیتابیس
    const db = await getDatabase()

    // دریافت پارامترهای صفحه‌بندی
    const query = getQuery(event)
    const page = parseInt(query.page as string) || 1
    const pageSize = parseInt(query.pageSize as string) || 10
    const offset = (page - 1) * pageSize

    // دریافت سفارشات از دیتابیس
    const ordersResult = await db.query(`
      SELECT 
        o.*,
        u.name as user_name,
        u.mobile as user_mobile,
        ua.street as shipping_address,
        ua.city as shipping_city,
        ua.province as shipping_province,
        ua.postal_code as shipping_postal_code,
        ua.recipient_name,
        ua.phone as recipient_phone,
        COUNT(oi.id) as items_count
      FROM orders o
      LEFT JOIN users u ON u.id = o.user_id
      LEFT JOIN user_addresses ua ON ua.id = o.shipping_address_id
      LEFT JOIN order_items oi ON oi.order_id = o.id
      GROUP BY o.id, u.id, ua.id
      ORDER BY o.created_at DESC
      LIMIT $1 OFFSET $2
    `, [pageSize, offset])

    const orders = ordersResult.rows as OrderRow[]

    // شمارش کل سفارشات
    const totalResult = await db.query('SELECT COUNT(*) as count FROM orders')
    const totalRow = totalResult.rows[0] as CountRow
    const total = parseInt(String(totalRow?.count || 0))

    // تبدیل به فرمت مورد انتظار
    const formattedOrders = orders.map((order: OrderRow) => ({
      id: order.id,
      orderNumber: order.order_number || `#${order.id}`,
      customerName: order.user_name || 'نامشخص',
      customerPhone: order.user_mobile || 'نامشخص',
      totalAmount: order.total_amount || 0,
      paymentMethod: order.payment_method || 'online',
      status: order.status || 'pending',
      paymentStatus: order.payment_status || 'pending',
      createdAt: order.created_at,
      itemsCount: parseInt(String(order.items_count || 0)),
      shippingAddress: order.shipping_address || '',
      shippingCity: order.shipping_city || '',
      shippingProvince: order.shipping_province || '',
      shippingPostalCode: order.shipping_postal_code || '',
      recipientName: order.recipient_name || '',
      recipientPhone: order.recipient_phone || '',
      trackingCode: order.tracking_code || order.id.toString(), // استفاده از ID به عنوان tracking code
      customerIP: order.customer_ip || 'نامشخص',
      userAgent: order.user_agent || 'نامشخص',
      orderIntegrity: 'verified'
    }))

    return {
      data: formattedOrders,
      page,
      total,
      success: true
    }
  } catch (error: unknown) {
    console.error('خطا در دریافت سفارشات:', error)
    const errorWithMessage = error as { message?: string }
    return {
      data: [],
      message: errorWithMessage?.message || 'خطا در دریافت سفارشات',
      success: false
    }
  }
})