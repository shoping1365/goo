import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')
    
    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه تراکنش الزامی است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, `/api/payments/transaction/${id}`, {
      method: 'GET'
    })

    return response
  } catch (error: any) {
    console.error('خطا در دریافت اطلاعات تراکنش:', error)
    
    if (error.statusCode) {
      throw createError({
        statusCode: error.statusCode,
        message: error.message || error.statusMessage || 'خطا در دریافت اطلاعات تراکنش'
      })
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی سرور'
    })
  }
}) 