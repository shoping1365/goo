import type { Ref } from 'vue'
import { computed, readonly, ref } from 'vue'

interface ApiResponse<T> {
     success: boolean
     data: T
     message?: string
}

interface ApiError {
     data?: { message?: string }
     message?: string
}

interface MobileAppHeader {
     id: number
     name: string
     description?: string
     platform: 'mobile' | 'app' | 'both'
     page_selection: 'all' | 'specific' | 'exclude'
     specific_pages?: string
     excluded_pages?: string
     layers: MobileAppHeaderLayer[]
     is_active: boolean
     created_at: string
     updated_at: string
}

interface MobileAppHeaderLayer {
     id: number
     mobile_app_header_id: number
     name: string
     width: number
     height: number
     row_count: number
     color: string
     opacity: number
     show_separator: boolean
     separator_type: string
     separator_color: string
     separator_opacity: number
     separator_width: number
     items: string // JSON string
     padding_left: number
     padding_right: number
     padding_top: number
     padding_bottom: number
     created_at: string
     updated_at: string
}

export const usePublicMobileHeaders = () => {
     const mobileHeaders: Ref<MobileAppHeader[]> = ref([])
     const loading = ref(false)
     const error = ref<string | null>(null)

     // بارگذاری هدرهای موبایل از سرور
     const loadMobileHeaders = async () => {
          loading.value = true
          error.value = null

          try {
               const response = await $fetch('/api/mobile-app-header-settings') as ApiResponse<MobileAppHeader[]>

               if (response.success) {
                    mobileHeaders.value = response.data || []
               } else {
                    error.value = response.message || 'خطا در بارگذاری هدرهای موبایل'
               }
          } catch (err) {
               const e = err as ApiError
               error.value = e.data?.message || e.message || 'خطا در اتصال به سرور'
               console.error('Error loading mobile headers:', e)
          } finally {
               loading.value = false
          }
     }

     // دریافت هدر فعال موبایل
     const getActiveMobileHeader = computed(() => {
          return mobileHeaders.value.find(header => header.is_active)
     })

     // دریافت هدر موبایل برای صفحه خاص
     const getMobileHeaderForPage = (pagePath: string) => {
          const activeHeader = getActiveMobileHeader.value

          if (!activeHeader) return null

          const pageSelection = activeHeader.page_selection
          const specificPages = activeHeader.specific_pages
          const excludedPages = activeHeader.excluded_pages

          // اگر همه صفحات انتخاب شده
          if (pageSelection === 'all') {
               return activeHeader
          }

          // اگر صفحات خاص انتخاب شده
          if (pageSelection === 'specific' && specificPages) {
               const pages = specificPages.split('\n').map(p => p.trim()).filter(p => p)
               if (pages.includes(pagePath)) {
                    return activeHeader
               }
          }

          // اگر صفحات مستثنی انتخاب شده
          if (pageSelection === 'exclude' && excludedPages) {
               const pages = excludedPages.split('\n').map(p => p.trim()).filter(p => p)
               if (!pages.includes(pagePath)) {
                    return activeHeader
               }
          }

          return null
     }

     // بررسی اینکه آیا هدر موبایل باید نمایش داده شود
     const shouldShowMobileHeader = (pagePath: string) => {
          const header = getMobileHeaderForPage(pagePath)
          return !!header
     }

     return {
          mobileHeaders: readonly(mobileHeaders),
          loading: readonly(loading),
          error: readonly(error),
          loadMobileHeaders,
          getActiveMobileHeader,
          getMobileHeaderForPage,
          shouldShowMobileHeader
     }
}
