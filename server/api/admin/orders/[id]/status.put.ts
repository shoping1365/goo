import { createError, defineEventHandler, getCookie, getRouterParam, readBody } from 'h3'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: typeof fetch

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const orderId = getRouterParam(event, 'id')
  const body = await readBody(event)

  console.log('Updating order status for ID:', orderId)
  console.log('Update data:', body)
  console.log('Go API Base:', config.public.goApiBase)

  if (!orderId) {
    throw createError({
      statusCode: 400,
      message: 'شناسه سفارش الزامی است'
    })
  }

  if (!body || (!body.status && !body.paymentStatus)) {
    throw createError({
      statusCode: 400,
      message: 'حداقل یکی از وضعیت‌ها باید ارسال شود'
    })
  }

  try {
    const url = `${config.public.goApiBase}/api/admin/orders/${orderId}/status`
    console.log('Proxying to URL:', url)

    // دریافت توکن احراز هویت
    const authToken = getCookie(event, 'access_token') || getCookie(event, 'auth-token')

    const response = await $fetch(url, {
      method: 'PUT',
      body: body,
      headers: {
        'Content-Type': 'application/json',
        ...(authToken && { 'Authorization': `Bearer ${authToken}` })
      }
    })
    console.log('Go API response:', response)

    return response
  } catch (error) {
    console.error('Error proxying to Go API:', error)
    console.error('Error details:', error instanceof Error ? error.message : String(error))
    throw createError({
      statusCode: 500,
      message: 'خطا در به‌روزرسانی وضعیت سفارش'
    })
  }
})
