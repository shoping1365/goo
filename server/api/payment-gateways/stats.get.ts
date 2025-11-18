export default defineEventHandler(async (event) => {
  try {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const response = await $fetch(`${base}/api/payment-gateways/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    return response
  } catch (error) {
    console.error('خطا در دریافت آمار درگاه‌های پرداخت:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت آمار درگاه‌های پرداخت',
    })
  }
})