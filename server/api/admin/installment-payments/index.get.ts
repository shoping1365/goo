import type { H3Event } from 'h3'
import { createError, defineEventHandler, getQuery } from 'h3'

interface User {
  id: number | string
  [key: string]: unknown
}

interface InstallmentRow {
  id: number
  customer_name: string
  national_id: string
  mobile: string
  product_id: number
  product_name: string
  amount: number
  installment_count: number
  monthly_payment: number
  status: string
  credit_score: number
  approved_by: number | null
  approved_at: string | null
  created_at: string
  updated_at: string
  product_sku?: string
  approver_name?: string
  [key: string]: unknown
}

interface CountRow {
  total: string
  [key: string]: unknown
}

declare const requireAdminAuth: (event: H3Event) => Promise<User | null>
declare const $db: {
  query: (sql: string, params?: unknown[]) => Promise<{ rows: unknown[]; rowCount: number }>
}

export default defineEventHandler(async (event) => {
  try {
    // بررسی دسترسی ادمین
    const user = await requireAdminAuth(event)
    if (!user) {
      throw createError({
        statusCode: 401,
        message: 'دسترسی غیرمجاز'
      })
    }

    // دریافت پارامترهای query
    const query = getQuery(event)
    const page = parseInt(query.page as string) || 1
    const limit = parseInt(query.limit as string) || 20
    const status = query.status as string
    const search = query.search as string
    const startDate = query.startDate as string
    const endDate = query.endDate as string

    // ساخت query برای دیتابیس
    let whereClause = 'WHERE 1=1'
    const params: unknown[] = []

    if (status) {
      whereClause += ' AND status = $' + (params.length + 1)
      params.push(status)
    }

    if (search) {
      whereClause += ` AND (customer_name ILIKE $${params.length + 1} OR national_id ILIKE $${params.length + 1})`
      params.push(`%${search}%`)
    }

    if (startDate) {
      whereClause += ' AND created_at >= $' + (params.length + 1)
      params.push(startDate)
    }

    if (endDate) {
      whereClause += ' AND created_at <= $' + (params.length + 1)
      params.push(endDate)
    }

    // شمارش کل رکوردها
    const countQuery = `
      SELECT COUNT(*) as total 
      FROM installment_payments 
      ${whereClause}
    `
    const countResult = await $db.query(countQuery, params)
    const total = parseInt((countResult.rows[0] as CountRow).total)

    // محاسبه offset
    const offset = (page - 1) * limit

    // دریافت داده‌ها
    const dataQuery = `
      SELECT 
        ip.id,
        ip.customer_name,
        ip.national_id,
        ip.mobile,
        ip.product_id,
        ip.product_name,
        ip.amount,
        ip.installment_count,
        ip.monthly_payment,
        ip.status,
        ip.credit_score,
        ip.approved_by,
        ip.approved_at,
        ip.created_at,
        ip.updated_at,
        p.name as product_name,
        p.sku as product_sku,
        u.name as approver_name
      FROM installment_payments ip
      LEFT JOIN products p ON ip.product_id = p.id
      LEFT JOIN users u ON ip.approved_by = u.id
      ${whereClause}
      ORDER BY ip.created_at DESC
      LIMIT $${params.length + 1} OFFSET $${params.length + 2}
    `
    params.push(limit, offset)

    const result = await $db.query(dataQuery, params)

    return {
      success: true,
      data: result.rows as InstallmentRow[],
      pagination: {
        page,
        limit,
        total,
        totalPages: Math.ceil(total / limit)
      }
    }

  } catch (error: unknown) {
    console.error('خطا در دریافت خریدهای اقساطی:', error)
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.statusMessage || 'خطا در دریافت خریدهای اقساطی'
    })
  }
}) 
