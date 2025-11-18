import { fetchGo } from '../../../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const type = getRouterParam(event, 'type')
    
    if (!type) {
      throw createError({
        statusCode: 400,
        message: 'نوع درگاه الزامی است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, `/api/payments/admin/gateway/${type}/capabilities`, {
      method: 'GET'
    })

    return response
  } catch (error: any) {
    console.error('خطا در دریافت قابلیت‌های درگاه:', error)
    
    if (error.statusCode) {
      throw createError({
        statusCode: error.statusCode,
        message: error.statusMessage || 'خطا در دریافت قابلیت‌های درگاه'
      })
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی سرور'
    })
  }
}) 