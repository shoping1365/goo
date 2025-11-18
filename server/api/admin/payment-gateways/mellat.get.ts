import type { H3Event } from 'h3'
import { createError, defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MellatSettingsResponse {
  data?: unknown
  [key: string]: unknown
}

export default defineEventHandler(async (event: H3Event) => {
  try {
    // بررسی احراز هویت ادمین
    const user = event.context.user
    if (!user || !user.is_admin) {
      throw createError({
        statusCode: 403,
        message: 'دسترسی غیرمجاز'
      })
    }

    // دریافت تنظیمات درگاه ملت از backend از طریق fetchGo
    const response = await fetchGo(event, '/api/admin/payment-gateways/mellat') as MellatSettingsResponse

    return {
      success: true,
      data: response.data,
      message: 'تنظیمات درگاه ملت با موفقیت دریافت شد'
    }

  } catch (error: unknown) {
    console.error('خطا در دریافت تنظیمات ملت:', error)
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus.statusCode || 500,
      message: errorWithStatus.statusMessage || 'خطا در دریافت تنظیمات ملت'
    })
  }
}) 
 
 