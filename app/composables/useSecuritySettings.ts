// composables/useSecuritySettings.ts
import { ref } from 'vue'

export interface SecuritySettings {
  csrf: boolean
  rateLimiter: boolean
  xss: boolean
  frameOptions: boolean
  hsts: boolean
  disabledUntil?: number
}

export interface RecaptchaSettings {
  version: 'v2' | 'v3'
  siteKey: string
  secretKey: string
  theme?: 'light' | 'dark'
  scoreThreshold?: number
}

// تعریف interface برای response های API
interface ApiResponse<T = any> {
  success: boolean
  data?: T
  message?: string
}

interface SecuritySettingsResponse {
  security?: SecuritySettings
  recaptcha?: RecaptchaSettings
  rateLimit?: any
}

export const useSecuritySettings = () => {
  const settings = ref<SecuritySettings | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Recaptcha settings
  const recaptchaSettings = ref<RecaptchaSettings>({
    version: 'v2',
    siteKey: '',
    secretKey: '',
    theme: 'light',
    scoreThreshold: 0.5
  })

  // Two-factor authentication
  const twoFactorEnabled = ref(false)

  // Active sessions
  const activeSessions = ref<any[]>([])

  // Login history
  const loginHistory = ref<any[]>([])

  const fetchSettings = async () => {
    loading.value = true
    try {
      const res = await $fetch<ApiResponse<SecuritySettingsResponse>>('/api/admin/security/settings')

      if (res.success && res.data) {
        // به‌روزرسانی تنظیمات امنیت
        if (res.data.security) {
          settings.value = res.data.security
        }

        // به‌روزرسانی تنظیمات reCAPTCHA
        if (res.data.recaptcha) {
          recaptchaSettings.value = res.data.recaptcha
        }

        // به‌روزرسانی تنظیمات Rate Limiting
        if (res.data.rateLimit) {
          // این مقدار در صفحه security.vue استفاده می‌شود
          console.log('Rate limit settings loaded:', res.data.rateLimit)
        }
      }
    } catch (err: any) {
      error.value = err?.message || 'خطا در دریافت تنظیمات'
    } finally {
      loading.value = false
    }
  }

  const saveSettings = async (data: Partial<SecuritySettings>) => {
    loading.value = true
    try {
      const res = await $fetch<ApiResponse<SecuritySettings>>('/api/admin/security/settings', {
        method: 'PUT',
        body: data,
      })
      if (res.data) {
        settings.value = res.data
      }
    } catch (err: any) {
      error.value = err?.message || 'خطا در ذخیره تنظیمات'
    } finally {
      loading.value = false
    }
  }

  const saveRecaptchaSettings = async (data: RecaptchaSettings) => {
    loading.value = true
    try {
      const res = await $fetch<ApiResponse>('/api/admin/security/settings', {
        method: 'PUT',
        body: { recaptcha: data }
      })

      if (res.success) {
        recaptchaSettings.value = { ...data }
        console.log('reCAPTCHA settings saved successfully')
      }
    } catch (err: any) {
      error.value = err?.message || 'خطا در ذخیره تنظیمات reCAPTCHA'
    } finally {
      loading.value = false
    }
  }

  const changePassword = async (currentPassword: string, newPassword: string) => {
    loading.value = true
    try {
      // اینجا باید API call برای تغییر رمز عبور اضافه شود
      console.log('Changing password...')
    } catch (err: any) {
      error.value = err?.message || 'خطا در تغییر رمز عبور'
    } finally {
      loading.value = false
    }
  }

  const toggleTwoFactor = async () => {
    loading.value = true
    try {
      twoFactorEnabled.value = !twoFactorEnabled.value
      // اینجا باید API call برای تغییر وضعیت 2FA اضافه شود
    } catch (err: any) {
      error.value = err?.message || 'خطا در تغییر وضعیت احراز هویت دو مرحله‌ای'
    } finally {
      loading.value = false
    }
  }

  const terminateSession = async (sessionId: number) => {
    try {
      // اینجا باید API call برای خاتمه نشست اضافه شود
      activeSessions.value = activeSessions.value.filter(s => s.id !== sessionId)
    } catch (err: any) {
      error.value = err?.message || 'خطا در خاتمه نشست'
    }
  }

  const loadActiveSessions = async () => {
    try {
      // اینجا باید API call برای بارگذاری نشست‌های فعال اضافه شود
      activeSessions.value = []
    } catch (err: any) {
      error.value = err?.message || 'خطا در بارگذاری نشست‌های فعال'
    }
  }

  const loadLoginHistory = async () => {
    try {
      // اینجا باید API call برای بارگذاری تاریخچه ورود اضافه شود
      loginHistory.value = []
    } catch (err: any) {
      error.value = err?.message || 'خطا در بارگذاری تاریخچه ورود'
    }
  }

  const saveRateLimitSettings = async (data: any) => {
    loading.value = true
    try {
      const res = await $fetch<ApiResponse>('/api/admin/security/settings', {
        method: 'PUT',
        body: { rateLimit: data }
      })

      if (res.success) {
        console.log('Rate limit settings saved successfully')
        return { success: true, message: 'تنظیمات محدودیت نرخ درخواست با موفقیت ذخیره شد' }
      }
    } catch (err: any) {
      error.value = err?.message || 'خطا در ذخیره تنظیمات Rate Limiting'
      return { success: false, message: 'خطا در ذخیره تنظیمات' }
    } finally {
      loading.value = false
    }
  }

  return {
    settings,
    loading,
    error,
    fetchSettings,
    saveSettings,
    recaptchaSettings,
    saveRecaptchaSettings,
    changePassword,
    twoFactorEnabled,
    toggleTwoFactor,
    activeSessions,
    terminateSession,
    loginHistory,
    loadActiveSessions,
    loadLoginHistory,
    saveRateLimitSettings
  }
}

// Export default برای سازگاری بیشتر
export default useSecuritySettings