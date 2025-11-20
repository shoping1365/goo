import { defineNuxtRouteMiddleware } from 'nuxt/app'
import type { RouteLocationNormalized } from 'vue-router'

// تعریف interface برای Mobile Header
interface MobileHeader {
  id: number
  is_active: boolean
  page_selection: 'all' | 'specific' | 'exclude'
  specific_pages?: string | null
  excluded_pages?: string | null
}

interface MobileHeaderDisplayState {
  shouldShow: boolean
  headerData: MobileHeader | null
  lastFetch: number
  isFetching: boolean
}

interface ApiResponse {
  success: boolean
  data: MobileHeader[]
}

// تعریف useState برای Nuxt 3
declare const useState: <T>(key: string, init?: () => T) => { value: T }

export default defineNuxtRouteMiddleware(async (to: RouteLocationNormalized, from: RouteLocationNormalized) => {
  // اگر مسیر تغییر نکرده، اجرا نکن
  if (to.path === from.path) {
    return
  }

  // اگر در صفحه ادمین هستیم، اجرا نکن
  if (to.path.startsWith('/admin')) {
    return
  }

  // استفاده از state برای جلوگیری از درخواست‌های مکرر
  const mobileHeaderDisplayState = useState<MobileHeaderDisplayState>('mobileHeaderDisplayState', () => ({
    shouldShow: true,
    headerData: null,
    lastFetch: 0,
    isFetching: false
  }))

  // اگر در حال بارگذاری است، صبر کن
  if (mobileHeaderDisplayState.value.isFetching) {
    return
  }

  // اگر کمتر از 60 ثانیه از آخرین بارگذاری گذشته، از کش استفاده کن
  const now = Date.now()
  if (mobileHeaderDisplayState.value.lastFetch > 0 && (now - mobileHeaderDisplayState.value.lastFetch) < 60000) {
    return
  }

  try {
    mobileHeaderDisplayState.value.isFetching = true

    // بارگذاری هدرهای موبایل از سرور
    const response = await $fetch<ApiResponse>('/api/mobile-app-header-settings')

    if (response && response.success && response.data) {
      const mobileHeaders = response.data
      const activeMobileHeader = mobileHeaders.find((header: MobileHeader) => header.is_active)

      if (activeMobileHeader) {
        // استفاده از نام‌های صحیح فیلدهای دیتابیس
        const pageSelection = activeMobileHeader.page_selection
        const specificPages = activeMobileHeader.specific_pages
        const excludedPages = activeMobileHeader.excluded_pages
        const currentPath = to.path

        let shouldShowMobileHeader = true

        // بررسی منطق نمایش هدر موبایل
        if (pageSelection === 'specific' && specificPages) {
          const pages = specificPages.split('\n').map((p: string) => p.trim()).filter((p: string) => p)
          shouldShowMobileHeader = pages.includes(currentPath)
        } else if (pageSelection === 'exclude' && excludedPages) {
          const pages = excludedPages.split('\n').map((p: string) => p.trim()).filter((p: string) => p)
          shouldShowMobileHeader = !pages.includes(currentPath)
        }

        // ذخیره state
        mobileHeaderDisplayState.value = {
          shouldShow: shouldShowMobileHeader,
          headerData: activeMobileHeader,
          lastFetch: now,
          isFetching: false
        }
      } else {
        // هیچ هدر موبایل فعالی یافت نشد
        mobileHeaderDisplayState.value = {
          shouldShow: false,
          headerData: null,
          lastFetch: now,
          isFetching: false
        }
      }
    } else {
      // خطا در دریافت داده
      mobileHeaderDisplayState.value = {
        shouldShow: false,
        headerData: null,
        lastFetch: now,
        isFetching: false
      }
    }
  } catch (_error) {
    // در صورت خطا، هدر را نمایش نده
    mobileHeaderDisplayState.value = {
      shouldShow: false,
      headerData: null,
      lastFetch: now,
      isFetching: false
    }
  }
})
