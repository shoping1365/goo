import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    // دریافت تنظیمات دسته‌بندی social-media از Go backend
    const response = await fetchGo(event, '/api/admin/settings/category/social-media', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })


    // بررسی اینکه response یک object تکی هست یا آرایه
    let responseData = response
    
    // اگر response یک object با key و value هست، به آرایه تبدیل کن
    if (response && !Array.isArray(response) && typeof response === 'object') {
      if (response.key || response.Key) {
        responseData = [response]
      } else if (response.data && Array.isArray(response.data)) {
        responseData = response.data
      }
    }

    // تبدیل داده‌های دریافتی به فرمت مناسب
  const settings: Record<string, unknown> = {}
  const prefixedKeys = new Set<string>()
    
    if (Array.isArray(responseData)) {
      responseData.forEach((item: { key?: string; Key?: string; value?: string; Value?: string }) => {
        const key = item.key || item.Key
        const value = item.value || item.Value
        
        if (key) {
          // حذف پیشوند social-media. از key اگر وجود دارد
          const cleanKey = key.replace(/^social-media\./, '')
          const isPrefixed = key.startsWith('social-media.')

          if (isPrefixed) {
            prefixedKeys.add(cleanKey)
            settings[cleanKey] = value
          } else if (!prefixedKeys.has(cleanKey) && settings[cleanKey] === undefined) {
            settings[cleanKey] = value
          }
        }
      })
    }

    // تبدیل خودکار لینک‌های سفارشی به آرایه قابل استفاده
    const rawCustomLinks = settings.custom_links ?? settings.customLinks
    if (typeof rawCustomLinks === 'string') {
      try {
        settings.customLinks = JSON.parse(rawCustomLinks)
      } catch (_parseError) {
        settings.customLinks = []
      }
    } else if (Array.isArray(rawCustomLinks)) {
      settings.customLinks = rawCustomLinks
    } else {
      settings.customLinks = []
    }

    return {
      success: true,
      data: settings,
      message: 'تنظیمات شبکه‌های اجتماعی با موفقیت دریافت شد'
    }
  } catch (error: unknown) {
    
    return {
      success: false,
      data: {},
      message: error?.data?.message || 'خطا در دریافت تنظیمات شبکه‌های اجتماعی'
    }
  }
})
