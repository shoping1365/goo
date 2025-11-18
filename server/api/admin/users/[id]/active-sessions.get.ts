import type { H3Event } from 'h3'
import { createError, defineEventHandler, getQuery, getRouterParam } from 'h3'
import { z } from 'zod'
import { getDatabase } from '~/server/api/_utils/database'

declare const requireAdminAuth: (event: H3Event) => Promise<{ id: number | string; role: string } | null>

interface SessionRow {
  id: number
  user_id: number
  token: string
  ip_address: string
  user_agent: string
  expires_at: string
  mobile: string
  auth_method: string
  is_verified: boolean
  verified_at: string | null
  is_active: boolean
  last_used_at: string | null
  created_at: string
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

    // دریافت سشن‌های فعال از دیتابیس
    const db = await getDatabase()

    // شمارش کل رکوردها
    const totalCount = await db.query(`
      SELECT COUNT(*) as count 
      FROM sessions 
      WHERE user_id = $1 AND is_active = true AND expires_at > NOW()
    `, [userId]) as unknown as { count: string }[]

    const total = parseInt(totalCount[0].count)

    // دریافت رکوردها با pagination
    const result = await db.query(`
      SELECT 
        id,
        user_id,
        token,
        ip_address,
        user_agent,
        expires_at,
        mobile,
        auth_method,
        is_verified,
        verified_at,
        is_active,
        last_used_at,
        created_at
      FROM sessions 
      WHERE user_id = $1 AND is_active = true AND expires_at > NOW()
      ORDER BY last_used_at DESC NULLS LAST, created_at DESC
      LIMIT $2 OFFSET $3
    `, [userId, limitNum, offset]) as unknown as SessionRow[]

    const activeSessions = result.map((row) => {
      // Parse User-Agent to extract device info
      const userAgent = row.user_agent || ''
      let deviceType: string
      let browser: string

      if (userAgent.includes('Mobile') || userAgent.includes('Android')) {
        deviceType = 'موبایل'
      } else if (userAgent.includes('Tablet') || userAgent.includes('iPad')) {
        deviceType = 'تبلت'
      } else {
        deviceType = 'دسکتاپ'
      }

      if (userAgent.includes('Chrome')) {
        browser = 'Chrome'
      } else if (userAgent.includes('Firefox')) {
        browser = 'Firefox'
      } else if (userAgent.includes('Safari')) {
        browser = 'Safari'
      } else if (userAgent.includes('Edge')) {
        browser = 'Edge';
      }

      return {
        id: row.id,
        userId: row.user_id,
        token: row.token.substring(0, 20) + '...', // نمایش جزئی از توکن
        ipAddress: row.ip_address,
        userAgent: row.user_agent,
        deviceType,
        browser,
        expiresAt: row.expires_at,
        mobile: row.mobile,
        authMethod: row.auth_method,
        isVerified: row.is_verified,
        verifiedAt: row.verified_at,
        isActive: row.is_active,
        lastUsedAt: row.last_used_at,
        createdAt: row.created_at,
      }
    })

    return {
      success: true,
      data: {
        activeSessions,
        pagination: {
          page: pageNum,
          limit: limitNum,
          total,
          totalPages: Math.ceil(total / limitNum)
        }
      }
    }

  } catch (error) {
    console.error('خطا در دریافت سشن‌های فعال:', error)

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت سشن‌های فعال'
    })
  }
})


