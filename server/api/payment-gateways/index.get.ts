export default defineEventHandler(async (event) => {
  try {
    // دریافت پارامترهای query
    const query = getQuery(event)
    const { type, status, search, is_test_mode, limit, order_by } = query

    // ساخت URL برای درخواست به بک‌اند
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    let url = `${base}/api/payment-gateways`
    
    // اضافه کردن پارامترهای فیلتر
    const params = new URLSearchParams()
    if (type) params.append('type', type as string)
    if (status) params.append('status', status as string)
    if (search) params.append('search', search as string)
    if (is_test_mode) params.append('is_test_mode', is_test_mode as string)
    if (limit) params.append('limit', limit as string)
    if (order_by) params.append('order_by', order_by as string)

    if (params.toString()) {
      url += '/filters?' + params.toString()
    }

    // درخواست به بک‌اند
    const response = await $fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    return response
  } catch (error) {
    console.error('خطا در دریافت درگاه‌های پرداخت:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت درگاه‌های پرداخت',
    })
  }
})