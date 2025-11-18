import type { H3Event } from 'h3'
import { createError, defineEventHandler, getQuery } from 'h3'
import { proxy } from '../_utils/fetchProxy'
import { getGoApiBaseUrl } from '~/server/utils/api-config'

export default defineEventHandler(async (event: H3Event) => {
  try {
    const query = getQuery(event)
    const categoryParam = query.category ?? 'library'
    const resolvedCategory = Array.isArray(categoryParam) ? (categoryParam[0] ?? 'library') : String(categoryParam)

    const baseUrl = getGoApiBaseUrl()
    const targetUrl = new URL('/api/media/list', baseUrl)
    targetUrl.searchParams.set('category', resolvedCategory)

    console.log('[media/list] Target URL:', targetUrl.toString(), { baseUrl, resolvedCategory })

    return await proxy(event, targetUrl.toString(), {
      method: 'GET'
    })
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; status?: number; message?: string; statusMessage?: string; data?: { message?: string; error?: string } }
    const status = errorWithStatus?.statusCode || errorWithStatus?.status
    if (status === 401) {
      // کاربر احراز نشده است؛ به جای پر کردن لاگ، پاسخ کنترل‌شده برمی‌گردانیم
      console.warn('[media/list] Unauthorized access - returning empty dataset')
      return { success: false, data: [], unauthorized: true }
    }

    console.error('[media/list] Error:', {
      statusCode: status,
      status: errorWithStatus?.status,
      message: errorWithStatus?.message,
      data: errorWithStatus?.data
    })
    throw createError({
      statusCode: status || 500,
      message: errorWithStatus?.data?.message || errorWithStatus?.data?.error || errorWithStatus?.statusMessage || errorWithStatus?.message || 'خطا در دریافت لیست مدیا',
      data: errorWithStatus?.data
    })
  }
})