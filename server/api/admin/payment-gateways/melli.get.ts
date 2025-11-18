import { createError, defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface PaymentGatewayResponse {
  success: boolean
  data?: unknown
  message?: string
}

export default defineEventHandler(async (event) => {
  try {
    // بررسی احراز هویت ادمین
    const user = event.context.user
    if (!user || !user.is_admin) {
      throw createError({
        statusCode: 403,
        message: 'دسترسی غیرمجاز'
      })
    }

    // دریافت تنظیمات درگاه ملی از backend از طریق fetchGo
    const response = await fetchGo(event, '/api/admin/payment-gateways/melli') as PaymentGatewayResponse

    return {
      success: true,
      data: response.data,
      message: 'تنظیمات درگاه ملی با موفقیت دریافت شد'
    }

  } catch (error: unknown) {
    console.error('خطا در دریافت تنظیمات ملی:', error)

    const errorObj = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorObj.statusCode || 500,
      message: errorObj.statusMessage || 'خطا در دریافت تنظیمات ملی'
    })
  }
}) 