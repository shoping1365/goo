import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { amount, callbackURL, description, email, mobile } = body

    // اعتبارسنجی ورودی‌ها
    if (!amount || amount < 1) {
      throw createError({
        statusCode: 400,
        message: 'مبلغ باید عدد مثبت باشد'
      })
    }

    if (!callbackURL) {
      throw createError({
        statusCode: 400,
        message: 'آدرس callback الزامی است'
      })
    }

    if (!description) {
      throw createError({
        statusCode: 400,
        message: 'توضیحات الزامی است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, '/api/payments/create', {
      method: 'POST',
      body: {
        gatewayType: 'saman',
        amount,
        callbackURL,
        description,
        email,
        mobile
      }
    }) as any

    return {
      success: true,
      data: response.data,
      message: 'درخواست پرداخت سامان با موفقیت ایجاد شد'
    }

  } catch (error: any) {
    console.error('خطا در ایجاد پرداخت سامان:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در ایجاد پرداخت سامان'
    })
  }
}) 