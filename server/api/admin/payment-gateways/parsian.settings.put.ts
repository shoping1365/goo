import { defineEventHandler, getCookie, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface SettingsResponse {
  success: boolean
  data?: unknown
  message?: string
}

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    // به‌روزرسانی تنظیمات درگاه پارسیان
    const response = await fetchGo(event, '/api/admin/parsian/settings', {
      method: 'PUT',
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'admin_token') || ''}`,
        'Content-Type': 'application/json'
      },
      body
    }) as SettingsResponse

    return {
      success: true,
      data: response.data,
      message: 'تنظیمات درگاه پارسیان با موفقیت به‌روزرسانی شد'
    }
  } catch (error: unknown) {
    console.error('خطا در به‌روزرسانی تنظیمات درگاه پارسیان:', error)
    const errorObj = error as { data?: { message?: string } }

    return {
      success: false,
      message: errorObj.data?.message || 'خطا در به‌روزرسانی تنظیمات درگاه پارسیان'
    }
  }
})
