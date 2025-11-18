// این کامپوزبل همیشه آخرین داده هدر را از سرور می‌گیرد و با SSR سازگار است. هیچ کش داخلی ندارد و برای پروژه‌های SSR مناسب است.
import { ref, readonly } from 'vue'
import type { Header, HeaderLayer } from '~/types/header'

interface ApiResponse<T> {
  success: boolean
  data: T
  message?: string
}

export const usePublicHeaders = () => {
  const headers = ref<Header[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // بارگذاری لیست هدرها (عمومی - بدون احراز هویت)
  const loadHeaders = async () => {
    try {
      loading.value = true
      error.value = null

      const response = await $fetch('/api/header-settings') as ApiResponse<Header[]>

      if (response.success) {
        headers.value = response.data
      } else {
        error.value = response.message || 'خطا در بارگذاری هدرها'
      }
    } catch (err) {
      console.error('خطا در بارگذاری هدرها:', err)
      error.value = 'خطا در بارگذاری هدرها'
    } finally {
      loading.value = false
    }
  }

  // دریافت هدر فعال
  const getActiveHeader = (): Header | null => {
    // بررسی اینکه headers.value آرایه معتبری باشد
    if (!headers.value || !Array.isArray(headers.value) || headers.value.length === 0) {
      return null
    }
    return headers.value.find(header => header.is_active) || null
  }

  // بررسی اینکه آیا هدر باید در صفحه فعلی نمایش داده شود
  const getHeaderForPage = (pagePath: string): Header | null => {
    const activeHeader = getActiveHeader()
    if (!activeHeader) return null

    const { pageSelection, specificPages, excludedPages } = activeHeader

    // اگر pageSelection تعریف نشده باشد، همیشه نمایش بده
    if (!pageSelection || pageSelection === 'all') {
      return activeHeader
    }

    // اگر pageSelection برابر 'specific' باشد
    if (pageSelection === 'specific') {
      if (!specificPages || specificPages.trim() === '') {
        // اگر صفحات خاص تعریف نشده باشند، نمایش نده
        return null
      }

      // تبدیل رشته صفحات خاص به آرایه و حذف موارد تکراری
      const pages = specificPages
        .split(/[\n,]/) // تقسیم بر اساس خط جدید یا کاما
        .map(page => page.trim())
        .filter(page => page.length > 0) // حذف خطوط خالی
        .filter((page, index, arr) => arr.indexOf(page) === index) // حذف موارد تکراری



      const shouldShow = pages.some(page => {
        // حذف http://localhost:3000 از ابتدای مسیر اگر وجود داشته باشد
        let cleanPage = page
        if (cleanPage.startsWith('http://localhost:3000')) {
          cleanPage = cleanPage.replace('http://localhost:3000', '')
        }
        if (cleanPage.startsWith('https://localhost:3000')) {
          cleanPage = cleanPage.replace('https://localhost:3000', '')
        }

        // بررسی دقیق مسیر
        if (cleanPage === pagePath) {

          return true
        }

        // بررسی مسیرهای نسبی (مثل /admin/*)
        if (cleanPage.endsWith('*')) {
          const basePath = cleanPage.slice(0, -1)
          if (pagePath.startsWith(basePath)) {

            return true
          }
        }

        return false
      })


      return shouldShow ? activeHeader : null
    }

    // اگر pageSelection برابر 'exclude' باشد
    if (pageSelection === 'exclude') {
      if (!excludedPages || excludedPages.trim() === '') {
        // اگر صفحات مستثنی تعریف نشده باشند، همیشه نمایش بده
        return activeHeader
      }

      // تبدیل رشته صفحات مستثنی به آرایه و حذف موارد تکراری
      const pages = excludedPages
        .split(/[\n,]/) // تقسیم بر اساس خط جدید یا کاما
        .map(page => page.trim())
        .filter(page => page.length > 0) // حذف خطوط خالی
        .filter((page, index, arr) => arr.indexOf(page) === index) // حذف موارد تکراری



      const shouldExclude = pages.some(page => {
        // حذف http://localhost:3000 از ابتدای مسیر اگر وجود داشته باشد
        let cleanPage = page
        if (cleanPage.startsWith('http://localhost:3000')) {
          cleanPage = cleanPage.replace('http://localhost:3000', '')
        }
        if (cleanPage.startsWith('https://localhost:3000')) {
          cleanPage = cleanPage.replace('https://localhost:3000', '')
        }

        // بررسی دقیق مسیر
        if (cleanPage === pagePath) {

          return true
        }

        // بررسی مسیرهای نسبی (مثل /admin/*)
        if (cleanPage.endsWith('*')) {
          const basePath = cleanPage.slice(0, -1)
          if (pagePath.startsWith(basePath)) {

            return true
          }
        }

        return false
      })


      return shouldExclude ? null : activeHeader
    }

    // در حالت پیش‌فرض، هدر را نمایش بده
    return activeHeader
  }

  return {
    headers: readonly(headers),
    loading: readonly(loading),
    error: readonly(error),
    loadHeaders,
    getActiveHeader,
    getHeaderForPage
  }
} 