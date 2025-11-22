import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { refId, amount } = body

    // اعتبارسنجی ورودی‌ها
    if (!refId) {
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
        gatewayType: 'mellat',
        refID: refId, // در ملت از refId استفاده می‌شود
        amount
      }
    }) as { data?: unknown }

    return {
      success: true,
      data: response.data,
      message: 'بازگشت وجه ملت با موفقیت انجام شد'
    }

  } catch (error: unknown) {
    console.error('خطا در بازگشت وجه ملت:', error)
    
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در بازگشت وجه ملت'
    })
  }
}) 
 
 