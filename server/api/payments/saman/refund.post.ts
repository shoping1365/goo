import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { refNum, amount } = body

    // اعتبارسنجی ورودی‌ها
    if (!refNum) {
      throw createError({
        statusCode: 400,
        message: 'شماره ارجاع الزامی است'
      })
    }

    if (!amount || amount < 1) {
      throw createError({
        statusCode: 400,
        message: 'مبلغ باید عدد مثبت باشد'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, '/api/payments/refund', {
      method: 'POST',
      body: {
        gatewayType: 'saman',
        refID: refNum, // در سامان از refNum استفاده می‌شود
        amount
      }
    }) as { data?: unknown }

    return {
      success: true,
      data: response.data,
      message: 'بازگشت وجه سامان با موفقیت انجام شد'
    }

  } catch (error: unknown) {
    console.error('خطا در بازگشت وجه سامان:', error)
    
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در بازگشت وجه سامان'
    })
  }
}) 