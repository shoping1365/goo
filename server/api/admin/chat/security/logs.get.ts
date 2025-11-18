import type { H3Event } from 'h3';
import { createError, defineEventHandler, getQuery } from 'h3';

declare const goApiFetch: (event: H3Event, endpoint: string, options?: { method?: string; body?: unknown }) => Promise<unknown>

export default defineEventHandler(async (event) => {
  try {
    // Get query parameters
    const query = getQuery(event)
    const days = Math.max(1, Math.min(30, parseInt(query.days as string) || 7))
    const riskLevel = query.risk_level as string || 'all'

    // Forward request to Go backend
    const response = await goApiFetch(event, `/api/chat/admin/security/logs?days=${days}&risk_level=${riskLevel}`, {
      method: 'GET'
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در دریافت لاگ‌های امنیتی چت:', error)

    const errorWithStatus = error as { statusCode?: number; message?: string }
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت لاگ‌های امنیتی چت'
    })
  }
})