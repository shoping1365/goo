import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const response = await fetchGo(event, '/api/admin/settings/auth', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    return {
      success: true,
      data: response,
      message: 'تنظیمات احراز هویت با موفقیت دریافت شد'
    }
  } catch (error: any) {
    console.error('خطا در دریافت تنظیمات احراز هویت:', error)
    
    return {
      success: false,
      data: null,
      message: error?.data?.message || 'خطا در دریافت تنظیمات احراز هویت'
    }
  }
}) 