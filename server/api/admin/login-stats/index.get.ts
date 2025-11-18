import { createError, defineEventHandler, getHeader } from 'h3'

declare const $fetch: <T = unknown>(url: string, options?: { headers?: Record<string, string> }) => Promise<T>

export default defineEventHandler(async (event) => {
  try {
    // Call backend API
    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/login-stats`, {
      headers: {
        'Authorization': getHeader(event, 'authorization') || ''
      }
    })

    return response
  } catch (error: unknown) {
    console.error('Error fetching login stats:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در دریافت آمار ورود'
    })
  }
}) 