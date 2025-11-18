import { createError, defineEventHandler, getHeader } from 'h3'
import { getCookieValue } from '../_utils/cookieHelper'

interface LoginCounters {
  successful?: number
  failed?: number
  suspicious?: number
  online?: number
  [key: string]: unknown
}

interface LoginStatsResponse {
  successfulLogins: number
  failedLogins: number
  suspiciousActivity: number
  onlineUsers: number
}

declare const $fetch: <T = unknown>(url: string, options?: { headers?: Record<string, string> }) => Promise<T>

export default defineEventHandler(async (event): Promise<LoginStatsResponse> => {
  try {
    const authHeader = getHeader(event, 'authorization') || `Bearer ${getCookieValue(event, 'auth-token') || ''}`
    const counters = await $fetch<LoginCounters>(`${process.env.BACKEND_URL}/api/admin/system-security/login/counters`, {
      headers: { Authorization: authHeader }
    })
    // نگاشت خروجی برای فرانت
    return {
      successfulLogins: counters?.successful || 0,
      failedLogins: counters?.failed || 0,
      suspiciousActivity: counters?.suspicious || 0,
      onlineUsers: counters?.online || 0,
    }
  } catch (error: unknown) {
    console.error('login-stats proxy error:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.message || 'خطا در دریافت آمار ورود' })
  }
})

