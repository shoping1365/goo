import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { amount, refId } = body

    // اعتبارسنجی ورودی‌ها
    if (!amount || amount < 1) {
      throw createError({
        statusCode: 400,
        message: 'مبلغ باید عدد مثبت باشد'
      })
    }

    if (!refId) {
      throw createError({
        statusCode: 400,
        message: 'شماره ارجاع الزامی است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, '/api/payments/verify', {
      method: 'POST',
      body: {
        gatewayType: 'mellat',
        amount,
        authority: refId // در ملت از refId استفاده می‌شود
      }
    }) as { data?: unknown }

    return {
      success: true,
      data: response.data,
      message: 'پرداخت ملت با موفقیت تایید شد'
    }

  } catch (error: unknown) {
    console.error('خطا در تایید پرداخت ملت:', error)
    
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در تایید پرداخت ملت'
    })
  }
}) 
 
 