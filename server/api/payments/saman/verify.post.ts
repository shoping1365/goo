import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { amount, refNum } = body

    // اعتبارسنجی ورودی‌ها
    if (!amount || amount < 1) {
      throw createError({
        statusCode: 400,
        message: 'مبلغ باید عدد مثبت باشد'
      })
    }

    if (!refNum) {
      throw createError({
        statusCode: 400,
        message: 'شماره ارجاع الزامی است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, '/api/payments/verify', {
      method: 'POST',
      body: {
        gatewayType: 'saman',
        amount,
        authority: refNum // در سامان از refNum استفاده می‌شود
      }
    }) as any

    return {
      success: true,
      data: response.data,
      message: 'پرداخت سامان با موفقیت تایید شد'
    }

  } catch (error: any) {
    console.error('خطا در تایید پرداخت سامان:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در تایید پرداخت سامان'
    })
  }
}) 