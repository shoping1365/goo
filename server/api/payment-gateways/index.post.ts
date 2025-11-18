export default defineEventHandler(async (event) => {
  try {
    // دریافت داده‌های ارسالی
    const body = await readBody(event)
    
    if (!body) {
      throw createError({
        statusCode: 400,
        message: 'داده‌های درگاه پرداخت الزامی است',
      })
    }

    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    // درخواست به بک‌اند
    const response = await $fetch(`${base}/api/payment-gateways`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: body,
    })

    return response
  } catch (error) {
    console.error('خطا در ثبت درگاه پرداخت:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در ثبت درگاه پرداخت',
    })
  }
})