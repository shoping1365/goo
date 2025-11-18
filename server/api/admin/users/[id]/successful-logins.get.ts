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

    // دریافت ورودهای موفق از دیتابیس
    const db = await getDatabase()
    
    // شمارش کل رکوردها
    const totalCount = await db.query(`
      SELECT COUNT(*) as count 
      FROM auth_events 
      WHERE user_id = $1 AND event_type = 'login_attempt' AND status = 'success'
    `, [userId])
    
    const total = parseInt((totalCount as any)[0].count)

    // دریافت رکوردها با pagination
    const result = await db.query(`
      SELECT 
        id,
        user_id,
        mobile,
        ip_address,
        user_agent,
        metadata,
        created_at
      FROM auth_events 
      WHERE user_id = $1 AND event_type = 'login_attempt' AND status = 'success'
      ORDER BY created_at DESC
      LIMIT $2 OFFSET $3
    `, [userId, limitNum, offset])

    const successfulLogins = (result as any).map((row: any) => {
      const metadata = row.metadata || {}
      const deviceInfo = metadata.device_info || {}
      
      return {
        id: row.id,
        userId: row.user_id,
        mobile: row.mobile,
        ipAddress: row.ip_address,
        userAgent: row.user_agent,
        browser: deviceInfo.browser || '',
        browserVersion: deviceInfo.browser_version || '',
        os: deviceInfo.os || '',
        osVersion: deviceInfo.os_version || '',
        deviceType: deviceInfo.device_type || '',
        deviceModel: deviceInfo.device_model || '',
        attemptType: metadata.method || '',
        createdAt: row.created_at,
      }
    })

    return {
      success: true,
      data: {
        successfulLogins,
        pagination: {
          page: pageNum,
          limit: limitNum,
          total,
          totalPages: Math.ceil(total / limitNum)
        }
      }
    }

  } catch (error: any) {
    console.error('خطا در دریافت ورودهای موفق:', error)
    
    if (error.statusCode) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت ورودهای موفق'
    })
  }
})


