import type { H3Event } from 'h3';
import { createError, defineEventHandler } from 'h3';

declare const goApiFetch: (event: H3Event, endpoint: string, options?: { method?: string; body?: unknown }) => Promise<unknown>

export default defineEventHandler(async (event) => {
  try {
    // Forward request to Go backend
    const response = await goApiFetch(event, '/api/chat/admin/settings/security', {
      method: 'GET'
    })

    return response
  } catch (error: unknown) {
    // console.error('خطا در دریافت تنظیمات امنیتی چت:', error)

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت تنظیمات امنیتی چت'
    })
  }
})