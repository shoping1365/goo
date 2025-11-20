import { readonly, ref } from 'vue'
import type { Footer } from '~/types/footer'

interface ApiResponse<T> {
     success: boolean
     data: T
     message?: string
}

export const usePublicFooters = () => {
     const footers = ref<Footer[]>([])
     const loading = ref(false)
     const error = ref<string | null>(null)

     // بارگذاری لیست فوترها (عمومی - بدون احراز هویت)
     const loadFooters = async () => {
          try {
               loading.value = true
               error.value = null

               const response = await $fetch('/api/footer-settings') as ApiResponse<Footer[]>

               if (response.success) {
                    footers.value = response.data
               } else {
                    error.value = response.message || 'خطا در بارگذاری فوترها'
               }
          } catch (err) {
               console.error('خطا در بارگذاری فوترها:', err)
               error.value = 'خطا در بارگذاری فوترها'
          } finally {
               loading.value = false
          }
     }

     // دریافت فوتر فعال
     const getActiveFooter = (): Footer | null => {
          // بررسی اینکه footers.value آرایه معتبری باشد
          if (!footers.value || !Array.isArray(footers.value) || footers.value.length === 0) {
               return null
          }
          return footers.value.find(footer => footer.is_active) || null
     }

     // بررسی اینکه آیا فوتر باید در صفحه فعلی نمایش داده شود
     const getFooterForPage = (pagePath: string): Footer | null => {
          const activeFooter = getActiveFooter()
          if (!activeFooter) return null

          const { pageSelection, specificPages, excludedPages } = activeFooter

          // اگر pageSelection تعریف نشده باشد، همیشه نمایش بده
          if (!pageSelection || pageSelection === 'all') {
               return activeFooter
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


               return shouldShow ? activeFooter : null
          }

          // اگر pageSelection برابر 'exclude' باشد
          if (pageSelection === 'exclude') {
               if (!excludedPages || excludedPages.trim() === '') {
                    // اگر صفحات مستثنی تعریف نشده باشند، همیشه نمایش بده
                    return activeFooter
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
               return shouldExclude ? null : activeFooter
          }

          // در حالت پیش‌فرض، فوتر را نمایش بده
          return activeFooter
     }

     return {
          footers: readonly(footers),
          loading: readonly(loading),
          error: readonly(error),
          loadFooters,
          getActiveFooter,
          getFooterForPage
     }
}

