import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    
    // اعتبارسنجی پارامترهای ورودی
    if (!body.gateway_id || !body.amount || !body.order_id || !body.description || !body.callback_url) {
      throw createError({
        statusCode: 400,
        message: 'پارامترهای ورودی ناقص است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo (جلوگیری از IPC socket)
    const response = await fetchGo(event, '/api/payments/create', {
      method: 'POST',
      body: {
        gateway_id: body.gateway_id,
        amount: body.amount,
        order_id: body.order_id,
        description: body.description,
        email: body.email || '',
        mobile: body.mobile || '',
        callback_url: body.callback_url
      }
    })

    return response
  } catch (error: any) {
    console.error('خطا در ایجاد پرداخت:', error)
    
    if (error.statusCode) {
      throw createError({
        statusCode: error.statusCode,
        message: error.message || error.statusMessage || 'خطا در ایجاد پرداخت'
      })
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطای داخلی سرور'
    })
  }
}) 