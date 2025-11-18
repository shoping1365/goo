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
    const backendUrl = `${base}/api/payment-gateways/${id}`
    
    // درخواست به بک‌اند
    const response = await $fetch(backendUrl, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: body,
    })

    return response
  } catch (error) {
    throw createError({
      statusCode: 500,
      message: 'خطا در ویرایش درگاه پرداخت',
    })
  }
})