import { fetchGo } from '../../../../_utils/fetchGo'

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
    const response = await fetchGo(event, `/api/payments/admin/transactions/${id}/refund`, {
      method: 'POST'
    })

    return response
  } catch (error: any) {
    console.error('خطا در بازگشت وجه:', error)
    
    if (error.statusCode) {
      throw createError({
        statusCode: error.statusCode,
        message: error.statusMessage || 'خطا در بازگشت وجه'
      })
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی سرور'
    })
  }
}) 