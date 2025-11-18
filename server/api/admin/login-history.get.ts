import { defineEventHandler, getHeader, getQuery } from 'h3'
import { getCookieValue } from '../_utils/cookieHelper'

interface LoginHistoryResponse {
  data?: unknown[]
  total?: number
  [key: string]: unknown
}

interface LoginAttempt {
  id: string
  userId: string
  username: string
  ipAddress: string
  userAgent: string
  timestamp: string
  status: 'success' | 'failed'
  reason?: string
  [key: string]: unknown
}

declare const $fetch: <T = unknown>(url: string, options?: { headers?: Record<string, string>; query?: Record<string, string | number> }) => Promise<T>

export default defineEventHandler(async (event): Promise<{ attempts: LoginAttempt[]; total: number }> => {
  try {
    const query = getQuery(event)
    const page = Number(query.page || 1)
    const limit = Number(query.limit || 10)

    // نگاشت فیلترها به بک‌اند
    const backendQuery: Record<string, string | number> = { page, limit }
    if (query.search) backendQuery.user = String(query.search)
    if (query.method) backendQuery.type = String(query.method)
    if (query.ip) backendQuery.ip = String(query.ip)
    if (query.dateFrom) backendQuery.date = String(query.dateFrom)
    if (query.success === 'true') backendQuery.status = 'successful'
    if (query.success === 'false') backendQuery.status = 'failed'

    const authHeader = getHeader(event, 'authorization') || `Bearer ${getCookieValue(event, 'auth-token') || ''}`
    const res = await $fetch<LoginHistoryResponse>(`${process.env.BACKEND_URL}/api/admin/system-security/login/history`, {
      headers: { Authorization: authHeader },
      query: backendQuery
    })

    // بک‌اند خروجی را با فیلدهای data/total/page/limit می‌دهد
    const attempts = Array.isArray(res?.data) ? (res.data as LoginAttempt[]) : []
    const total = Number(res?.total || 0)

    return { attempts, total }
  } catch (error: unknown) {
    console.error('login-history proxy error:', error)
    return { attempts: [], total: 0 }
  }
})

