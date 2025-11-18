export default defineEventHandler(async (event) => {
  try {
    // دریافت ID از پارامترهای مسیر
    const id = getRouterParam(event, 'id')
    
    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه درگاه پرداخت الزامی است',
      })
    }

    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const backendUrl = `${base}/api/payment-gateways/${id}`
    
    // درخواست به بک‌اند با timeout و retry
    
    const response = await $fetch(backendUrl, {
      timeout: 5000, // 5 ثانیه timeout
      retry: 1, // یک بار retry
      retryDelay: 1000 // 1 ثانیه تاخیر
    })
    
    return response
  } catch (error) {
    console.error('خطا در دریافت درگاه پرداخت:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت درگاه پرداخت',
    })
  }
})