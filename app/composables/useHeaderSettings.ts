import { reactive, ref } from 'vue'

// تعریف تایپ تنظیمات هدر
export interface HeaderSettings {
  // تنظیمات عمومی
  type: 'fixed' | 'sticky' | 'static'
  height: number
  backgroundColor: string
  opacity: number
  showOnMobile: boolean
  showOnTablet: boolean
  showOnDesktop: boolean
  // تنظیمات لوگو
  logoUrl: string
  siteName: string
  logoHeight: number
  showSiteName: boolean
  // تنظیمات منو
  menuId: string
  menuPosition: 'left' | 'center' | 'right'
  menuTextColor: string
  menuBackgroundColor: string
  // تنظیمات جستجو
  showSearch: boolean
  searchType: 'products' | 'all' | 'blog'
  searchPlaceholder: string
  searchPosition: 'left' | 'center' | 'right'
  // تنظیمات دکمه‌های عملیات
  showCart: boolean
  cartDisplay: 'icon' | 'icon-text' | 'full'
  showUserMenu: boolean
  userMenuDisplay: 'icon' | 'avatar' | 'name'
  showLanguageSelector: boolean
  showCurrencySelector: boolean
}

// تنظیمات پیش‌فرض
const defaultSettings: HeaderSettings = {
  type: 'fixed',
  height: 80,
  backgroundColor: '#ffffff',
  opacity: 100,
  showOnMobile: true,
  showOnTablet: true,
  showOnDesktop: true,
  logoUrl: '',
  siteName: 'فروشگاه آنلاین',
  logoHeight: 50,
  showSiteName: true,
  menuId: 'main-menu',
  menuPosition: 'center',
  menuTextColor: '#333333',
  menuBackgroundColor: 'transparent',
  showSearch: true,
  searchType: 'products',
  searchPlaceholder: 'جستجو در محصولات...',
  searchPosition: 'center',
  showCart: true,
  cartDisplay: 'icon',
  showUserMenu: true,
  userMenuDisplay: 'icon',
  showLanguageSelector: false,
  showCurrencySelector: false
}

export const useHeaderSettings = () => {
  const settings = reactive<HeaderSettings>({ ...defaultSettings })
  const loading = ref(false)
  const saving = ref(false)
  const error = ref<string | null>(null)

  // بارگذاری تنظیمات از سرور
  const loadSettings = async () => {
    try {
      loading.value = true
      error.value = null

      const response = await $fetch('/api/admin/header-settings', {
        method: 'GET'
      })

      if (Array.isArray(response)) {
        response.forEach((setting: any) => {
          const key = setting.key || setting.Key
          const value = setting.value ?? setting.Value

          // حذف پیشوند header. از کلید
          const cleanKey = key.replace('header.', '')

          if (cleanKey in settings) {
            const key = cleanKey as keyof HeaderSettings
            const currentValue = settings[key]

            // تبدیل نوع داده
            if (typeof currentValue === 'number') {
              ; (settings as any)[key] = Number(value)
            } else if (typeof currentValue === 'boolean') {
              ; (settings as any)[key] = value === 'true' || value === true
            } else {
              ; (settings as any)[key] = value
            }
          }
        })
      }
    } catch (err) {
      console.error('خطا در بارگذاری تنظیمات هدر:', err)
      error.value = 'خطا در بارگذاری تنظیمات هدر'
      throw err
    } finally {
      loading.value = false
    }
  }

  // ذخیره تنظیمات
  const saveSettings = async () => {
    try {
      saving.value = true
      error.value = null

      // آماده‌سازی تنظیمات برای ارسال
      const settingsData = Object.entries(settings).map(([key, value]) => ({
        key: `header.${key}`,
        value: value.toString(),
        category: 'header',
        type: typeof value === 'number' ? 'number' : typeof value === 'boolean' ? 'boolean' : 'string'
      }))

      // ارسال به سرور
      await $fetch('/api/admin/header-settings', {
        method: 'PUT',
        body: settingsData
      })
    } catch (err) {
      console.error('خطا در ذخیره تنظیمات هدر:', err)
      error.value = 'خطا در ذخیره تنظیمات هدر'
      throw err
    } finally {
      saving.value = false
    }
  }

  // بازنشانی به تنظیمات پیش‌فرض
  const resetToDefaults = () => {
    Object.assign(settings, defaultSettings)
  }

  // آپلود لوگو
  const uploadLogo = async (file: File): Promise<string> => {
    // اعتبارسنجی اندازه فایل
    if (file.size > 2 * 1024 * 1024) {
      throw new Error('حجم فایل نباید بیشتر از 2 مگابایت باشد')
    }

    // اعتبارسنجی نوع فایل
    const allowedTypes = ['image/jpeg', 'image/png', 'image/svg+xml']
    if (!allowedTypes.includes(file.type)) {
      throw new Error('فقط فایل‌های JPG، PNG و SVG مجاز هستند')
    }

    return new Promise((resolve, reject) => {
      const reader = new FileReader()
      reader.onload = (e) => {
        const result = e.target?.result as string
        settings.logoUrl = result
        resolve(result)
      }
      reader.onerror = () => reject(new Error('خطا در خواندن فایل'))
      reader.readAsDataURL(file)
    })
  }

  // حذف لوگو
  const removeLogo = () => {
    settings.logoUrl = ''
  }

  return {
    settings,
    loading,
    saving,
    error,
    loadSettings,
    saveSettings,
    resetToDefaults,
    uploadLogo,
    removeLogo
  }
} 