import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    
    // اعتبارسنجی پارامترهای ورودی
    if (!body.authority || !body.amount) {
      throw createError({
        statusCode: 400,
        message: 'پارامترهای ورودی ناقص است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, '/api/payments/verify', {
      method: 'POST',
      body: {
        authority: body.authority,
        amount: body.amount
      }
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در تایید پرداخت:', error)
    
    const err = error as { statusCode?: number; message?: string; statusMessage?: string }
    if (err.statusCode) {
      throw createError({
        statusCode: err.statusCode,
        message: err.message || err.statusMessage || 'خطا در تایید پرداخت'
      })
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی سرور'
    })
  }
}) 