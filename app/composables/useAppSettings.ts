import { ref, readonly } from 'vue'

// تنظیمات عمومی اپلیکیشن
const appSettings = ref({
  // تنظیمات URL محصولات
  includeSlugInURL: true, // آیا slug در URL محصول نمایش داده شود یا نه
  
  // سایر تنظیمات عمومی
  defaultLanguage: 'fa',
  currency: 'IRR',
  timezone: 'Asia/Tehran'
})

export const useAppSettings = () => {
  // دریافت تنظیمات از localStorage (اگر موجود باشد)
  const loadSettings = () => {
    if (import.meta.client) {
      try {
        const saved = localStorage.getItem('app-settings')
        if (saved) {
          const parsed = JSON.parse(saved)
          appSettings.value = { ...appSettings.value, ...parsed }
        }
      } catch (error) {
        console.warn('خطا در بارگذاری تنظیمات:', error)
      }
    }
  }

  // ذخیره تنظیمات در localStorage
  const saveSettings = () => {
    if (import.meta.client) {
      try {
        localStorage.setItem('app-settings', JSON.stringify(appSettings.value))
      } catch (error) {
        console.warn('خطا در ذخیره تنظیمات:', error)
      }
    }
  }

  // تغییر تنظیمات
  const updateSetting = (key: string, value: any) => {
    if (key in appSettings.value) {
      appSettings.value[key] = value
      saveSettings()
    }
  }

  // بارگذاری تنظیمات در mount
  if (import.meta.client) {
    loadSettings()
  }

  return {
    appSettings: readonly(appSettings),
    updateSetting,
    saveSettings,
    loadSettings
  }
}





























