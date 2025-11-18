import type { H3Event } from 'h3';
import { createError, defineEventHandler, getCookie, getRouterParam } from 'h3';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string> }) => Promise<T>

export default defineEventHandler(async (event: H3Event) => {
  const config = useRuntimeConfig()
  const orderId = getRouterParam(event, 'id')

  console.log('Fetching order details for ID:', orderId)
  console.log('Go API Base:', config.public.goApiBase)

  if (!orderId) {
    throw createError({
      statusCode: 400,
      message: 'شناسه سفارش الزامی است'
    })
  }

  try {
    const url = `${config.public.goApiBase}/api/admin/orders/${orderId}`
    console.log('Proxying to URL:', url)

    // دریافت توکن احراز هویت
    const authToken = getCookie(event, 'access_token') || getCookie(event, 'auth-token')

    const response = await $fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        ...(authToken && { 'Authorization': `Bearer ${authToken}` })
      }
    })
    console.log('Go API response:', response)

    return response
  } catch (error: unknown) {
    console.error('Error proxying to Go API:', error)
    const errorMessage = error instanceof Error ? error.message : String(error)
    console.error('Error details:', errorMessage)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در دریافت اطلاعات سفارش'
    })
  }
})
