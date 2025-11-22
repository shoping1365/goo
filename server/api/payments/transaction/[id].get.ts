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
  } catch (error: unknown) {
    console.error('خطا در دریافت اطلاعات تراکنش:', error)
    
    const err = error as { statusCode?: number; message?: string; statusMessage?: string }
    if (err.statusCode) {
      throw createError({
        statusCode: err.statusCode,
        message: err.message || err.statusMessage || 'خطا در دریافت اطلاعات تراکنش'
      })
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی سرور'
    })
  }
}) 