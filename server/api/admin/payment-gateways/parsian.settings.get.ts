import { defineEventHandler, getCookie } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface SettingsResponse {
  success: boolean
  data?: unknown
  message?: string
}

export default defineEventHandler(async (event) => {
  try {
    // دریافت تنظیمات درگاه پارسیان
    const response = await fetchGo(event, '/api/admin/parsian/settings', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'admin_token') || ''}`
      }
    }) as SettingsResponse

    return {
      success: true,
      data: response.data,
      message: 'تنظیمات درگاه پارسیان با موفقیت دریافت شد'
    }
  } catch (error: unknown) {
    console.error('خطا در دریافت تنظیمات درگاه پارسیان:', error)
    const errorObj = error as { data?: { message?: string } }

    return {
      success: false,
      message: errorObj.data?.message || 'خطا در دریافت تنظیمات درگاه پارسیان'
    }
  }
})
