import { createError, defineEventHandler, getCookie, getRouterParam } from 'h3'
import { requireAuth } from '~/server/utils/requireAuth'

export default defineEventHandler(async (event) => {
  try {
    const config = useRuntimeConfig()
    const apiBase = config.public.goApiBase
    const id = getRouterParam(event, 'id')

    // Check if user is authenticated
    const user = await requireAuth(event)
    if (!user) {
      throw createError({
        statusCode: 401,
        message: 'Unauthorized'
      })
    }

    // Get access token from cookies
    const accessToken = getCookie(event, 'access_token') || getCookie(event, 'auth-token')

    // ارسال درخواست به backend با توکن احراز هویت
    const response = await fetch(`${apiBase}/api/widgets/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`
      }
    }).then(res => res.json())

    return response
  } catch (error: unknown) {
    console.error('خطا در دریافت ویجت:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در دریافت ویجت',
      data: (error as { data?: unknown; message?: string }).data || (error as { data?: unknown; message?: string }).message
    })
  }
})