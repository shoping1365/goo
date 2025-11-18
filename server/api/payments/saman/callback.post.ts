import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { RefNum, ResNum, state } = body

    // اعتبارسنجی ورودی‌ها
    if (!state) {
      throw createError({
        statusCode: 400,
        message: 'وضعیت تراکنش مشخص نشده'
      })
    }

    // ارسال درخواست به backend برای پردازش callback از طریق fetchGo
    const response = await fetchGo(event, '/api/payments/callback', {
      method: 'POST',
      body: {
        gatewayType: 'saman',
        formData: {
          RefNum,
          ResNum,
          state
        }
      }
    }) as any

    return {
      success: true,
      data: response.data,
      message: 'Callback سامان با موفقیت پردازش شد'
    }

  } catch (error: any) {
    console.error('خطا در پردازش callback سامان:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در پردازش callback سامان'
    })
  }
}) 