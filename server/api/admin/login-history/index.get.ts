import { createError, defineEventHandler, getHeader, getQuery } from 'h3';

declare const $fetch: <T = unknown>(url: string, options?: { headers?: Record<string, string>; query?: Record<string, string | number> }) => Promise<T>

interface LoginHistoryItem {
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

interface LoginHistoryResponse {
  data: LoginHistoryItem[]
  total: number
  page: number
  limit: number
  [key: string]: unknown
}

export default defineEventHandler(async (event): Promise<LoginHistoryResponse> => {
  try {
    const query = getQuery(event)
    const page = parseInt(query.page as string) || 1
    const limit = parseInt(query.limit as string) || 20
    const search = query.search as string
    const success = query.success as string
    const method = query.method as string
    const dateFrom = query.dateFrom as string
    const ip = query.ip as string

    // Call backend API
    const response = await $fetch<LoginHistoryResponse>(`${process.env.BACKEND_URL}/api/admin/login-history`, {
      query: {
        page,
        limit,
        search,
        success,
        method,
        dateFrom,
        ip
      },
      headers: {
        'Authorization': getHeader(event, 'authorization') || ''
      }
    })

    return response
  } catch (error: unknown) {
    console.error('Error fetching login history:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در دریافت تاریخچه ورود'
    })
  }
}) 