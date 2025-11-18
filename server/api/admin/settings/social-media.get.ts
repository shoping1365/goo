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

    console.log('📥 Raw response from backend:', response)

    // بررسی اینکه response یک object تکی هست یا آرایه
    let responseData = response
    
    // اگر response یک object با key و value هست، به آرایه تبدیل کن
    if (response && !Array.isArray(response) && typeof response === 'object') {
      if (response.key || response.Key) {
        console.log('⚠️ Response is single object, wrapping in array')
        responseData = [response]
      } else if (response.data && Array.isArray(response.data)) {
        console.log('📦 Response has data array, using it')
        responseData = response.data
      }
    }

    // تبدیل داده‌های دریافتی به فرمت مناسب
  const settings: any = {}
  const prefixedKeys = new Set<string>()
    
    if (Array.isArray(responseData)) {
      responseData.forEach((item: any) => {
        const key = item.key || item.Key
        const value = item.value || item.Value
        
        console.log(`🔍 Processing: key="${key}", value="${value}"`)
        
        if (key) {
          // حذف پیشوند social-media. از key اگر وجود دارد
          const cleanKey = key.replace(/^social-media\./, '')
          const isPrefixed = key.startsWith('social-media.')

          if (isPrefixed) {
            prefixedKeys.add(cleanKey)
            settings[cleanKey] = value
            console.log(`✅ Mapped (prefixed): "${key}" -> "${cleanKey}" = "${value}"`)
          } else if (!prefixedKeys.has(cleanKey) && settings[cleanKey] === undefined) {
            settings[cleanKey] = value
            console.log(`⚙️ Mapped (legacy): "${key}" -> "${cleanKey}" = "${value}"`)
          } else {
            console.log(`ℹ️ Skipped legacy key for "${cleanKey}" because prefixed value already loaded`)
          }
        }
      })
    } else {
      console.log('❌ Invalid response format - expected array but got:', typeof responseData)
    }

    // تبدیل خودکار لینک‌های سفارشی به آرایه قابل استفاده
    const rawCustomLinks = settings.custom_links ?? settings.customLinks
    if (typeof rawCustomLinks === 'string') {
      try {
        settings.customLinks = JSON.parse(rawCustomLinks)
      } catch (parseError) {
        console.warn('⚠️ Failed to parse custom_links JSON:', parseError)
        settings.customLinks = []
      }
    } else if (Array.isArray(rawCustomLinks)) {
      settings.customLinks = rawCustomLinks
    } else {
      settings.customLinks = []
    }

    console.log('📦 Final settings object:', settings)

    return {
      success: true,
      data: settings,
      message: 'تنظیمات شبکه‌های اجتماعی با موفقیت دریافت شد'
    }
  } catch (error: any) {
    console.error('❌ خطا در دریافت تنظیمات شبکه‌های اجتماعی:', error)
    
    return {
      success: false,
      data: {},
      message: error?.data?.message || 'خطا در دریافت تنظیمات شبکه‌های اجتماعی'
    }
  }
})
