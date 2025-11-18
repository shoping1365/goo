import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MellatSettingsBody {
  merchant_id?: string
  public_key?: string
  private_key?: string
  callback_url?: string
  is_test_mode?: boolean
  is_active?: boolean
  [key: string]: unknown
}

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

    const body = await readBody(event) as MellatSettingsBody
    const { merchant_id, public_key, private_key, callback_url, is_test_mode, is_active } = body

    // اعتبارسنجی ورودی‌ها
    if (!merchant_id) {
      throw createError({
        statusCode: 400,
        message: 'شماره ترمینال الزامی است'
      })
    }

    if (!public_key) {
      throw createError({
        statusCode: 400,
        message: 'نام کاربری الزامی است'
      })
    }

    if (!private_key) {
      throw createError({
        statusCode: 400,
        message: 'رمز عبور الزامی است'
      })
    }

    if (!callback_url) {
      throw createError({
        statusCode: 400,
        message: 'آدرس callback الزامی است'
      })
    }

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, '/api/admin/payment-gateways/mellat', {
      method: 'PUT',
      body: {
        merchant_id,
        public_key,
        private_key,
        callback_url,
        is_test_mode: is_test_mode || false,
        is_active: is_active || false
      }
    }) as MellatSettingsResponse

    return {
      success: true,
      data: response.data,
      message: 'تنظیمات درگاه ملت با موفقیت به‌روزرسانی شد'
    }

  } catch (error: unknown) {
    console.error('خطا در به‌روزرسانی تنظیمات ملت:', error)
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus.statusCode || 500,
      message: errorWithStatus.statusMessage || 'خطا در به‌روزرسانی تنظیمات ملت'
    })
  }
}) 
 
 