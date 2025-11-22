import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { RefId, ResCode, SaleOrderId, SaleReferenceId } = body

    // اعتبارسنجی ورودی‌ها
    if (!ResCode) {
      throw createError({
        statusCode: 400,
        message: 'کد نتیجه مشخص نشده'
      })
    }

    // ارسال درخواست به backend برای پردازش callback از طریق fetchGo
    const response = await fetchGo(event, '/api/payments/callback', {
      method: 'POST',
      body: {
        gatewayType: 'mellat',
        formData: {
          RefId,
          ResCode,
          SaleOrderId,
          SaleReferenceId
        }
      }
    }) as { data?: unknown }

    return {
      success: true,
      data: response.data,
      message: 'Callback ملت با موفقیت پردازش شد'
    }

  } catch (error: unknown) {
    console.error('خطا در پردازش callback ملت:', error)
    
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { statusMessage?: string }).statusMessage || 'خطا در پردازش callback ملت'
    })
  }
}) 
 
 