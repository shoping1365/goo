import type { H3Event } from 'h3'
import { defineEventHandler, createError } from 'h3'
import { z } from 'zod'
import { getDatabase } from '~/server/api/_utils/database.js'

declare const requireAdminAuth: (event: H3Event) => Promise<{ id: number | string; role: string } | null>

interface OrderRow {
  id: number
  order_number: string
  user_id: number
  status: string
  payment_status: string
  payment_method: string | null
  total_amount: number
  final_amount: number
  shipping_address: string | null
  billing_address: string | null
  created_at: string
  updated_at: string
  mobile: string | null
  name: string | null
  email: string | null
  last_login_ip: string | null
  profile_data: any | null
}

interface CountRow {
  count: string
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

    // دریافت سفارشات از دیتابیس
    const db = await getDatabase()
    
    // شمارش کل رکوردها
    const totalCount = await db.query(`
      SELECT COUNT(*) as count 
      FROM orders 
      WHERE user_id = $1
    `, [userId]) as unknown as CountRow[]
    
    const total = parseInt(totalCount[0].count)

    // دریافت رکوردها با pagination و اطلاعات کاربر
    const result = await db.query(`
      SELECT 
        o.id,
        o.order_number,
        o.user_id,
        o.status,
        o.payment_status,
        o.payment_method,
        o.total_amount,
        o.final_amount,
        o.shipping_address,
        o.billing_address,
        o.created_at,
        o.updated_at,
        u.mobile,
        u.name,
        u.email,
        u.last_login_ip,
        u.profile_data
      FROM orders o
      LEFT JOIN users u ON o.user_id = u.id
      WHERE o.user_id = $1
      ORDER BY o.created_at DESC
      LIMIT $2 OFFSET $3
    `, [userId, limitNum, offset]) as unknown as OrderRow[]

    console.log('Raw database result:', JSON.stringify(result, null, 2));

    const orders = result.map((row) => {
      // استخراج کد ملی از profile_data
      let nationalCode = '';
      if (row.profile_data && typeof row.profile_data === 'object') {
        nationalCode = row.profile_data.national_code || row.profile_data.nationalCode || '';
      }

      const orderData = {
        id: row.id,
        orderNumber: row.order_number,
        userId: row.user_id,
        status: row.status,
        paymentStatus: row.payment_status,
        paymentMethod: row.payment_method || '',
        totalAmount: row.total_amount,
        finalAmount: row.final_amount,
        shippingAddress: row.shipping_address || '',
        billingAddress: row.billing_address || '',
        createdAt: row.created_at,
        updatedAt: row.updated_at,
        // اطلاعات کاربر
        userMobile: row.mobile || '',
        userName: row.name || '',
        userEmail: row.email || '',
        userLastLoginIP: row.last_login_ip || '',
        userNationalCode: nationalCode,
      };

      console.log('Mapped order data:', JSON.stringify(orderData, null, 2));
      return orderData;
    })

    return {
      success: true,
      data: {
        orders,
        pagination: {
          page: pageNum,
          limit: limitNum,
          total,
          totalPages: Math.ceil(total / limitNum)
        }
      }
    }

  } catch (error) {
    console.error('خطا در دریافت سفارشات کاربر:', error)
    
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت سفارشات کاربر'
    })
  }
})

