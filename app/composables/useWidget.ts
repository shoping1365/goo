import { useRequestHeaders } from 'nuxt/app'
import { computed, readonly, ref } from 'vue'
import type {
    WidgetListResponse,
    WidgetPage
} from '~/types/widget'
import { useWidgetCore } from './_useWidgetCore'

/**
 * Composable برای مدیریت ابزارک‌ها
 * شامل تمام عملیات CRUD و مدیریت state
 */
export const useWidget = () => {
  // Initialize core composable
  const core = useWidgetCore('/api/admin/widgets')
  const {
    widgets,
    widget,
    loading,
    error,
    clearError,
    fetchWidgets,
    fetchWidget,
    fetchWidgetByCode,
    createWidget,
    updateWidget,
    deleteWidget,
    toggleWidgetStatus,
    duplicateWidget,
    updateWidgetOrder,
    sortWidgets,
    filterWidgets
  } = core

  // Global refresh trigger
  const refreshTrigger = ref(0)

  // Function to trigger global refresh
  const triggerGlobalRefresh = () => {
    refreshTrigger.value++
  }

  /**
   * دریافت ابزارک‌های یک صفحه خاص
   * این متد مخصوص useWidget است و در core وجود ندارد
   */
  const fetchWidgetsByPage = (page: WidgetPage) => {
    loading.value = true
    error.value = null

    // استفاده از endpoint عمومی برای نمایش ویجت‌ها
    const url = `/api/widgets/page/${page}?_=${new Date().getTime()}`

    const headers = process.server
      ? useRequestHeaders(['cookie', 'authorization'])
      : undefined

    return $fetch<WidgetListResponse>(url, {
      timeout: 0, // حذف کامل timeout
      credentials: 'include',
      headers
    }).then((response) => {
      widgets.value = response.data
      loading.value = false
      return response.data
    }).catch((err: unknown) => {
      // خطا در بارگذاری ابزارک‌ها
      const fetchError = err as { statusCode?: number; message?: string }

      // بررسی نوع خطا
      if (fetchError.statusCode === 404) {
        error.value = 'صفحه مورد نظر یافت نشد'
      } else if (fetchError.statusCode === 500) {
        error.value = 'خطای سرور - لطفاً بعداً تلاش کنید'
      } else if (fetchError.message?.includes('fetch')) {
        error.value = 'خطا در اتصال به سرور بک اند - با دولوپر وب سایت تماس بگیرید'
      } else {
        error.value = fetchError.message || 'خطا در دریافت ابزارک‌ها'
      }

      // در صورت خطا، ویجت‌ها را خالی کن
      widgets.value = []
      loading.value = false
      throw err
    })
  }

  /**
   * بررسی وجود ابزارک با کد
   */
  const isCodeUnique = (code: string, excludeId?: number) => {
    return !widgets.value.some(w => w.code === code && w.id !== excludeId)
  }

  /**
   * گرفتن آخرین ابزارک‌ها
   */
  const getRecentWidgets = (limit: number = 5) => {
    return computed(() => {
      return widgets.value
        .slice()
        .sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
        .slice(0, limit)
    })
  }

  /**
   * گرفتن ابزارک‌های فعال
   */
  const getActiveWidgets = () => {
    return computed(() => {
      return widgets.value.filter(w => w.status === 'active')
    })
  }

  /**
   * Reset state
   */
  const resetState = () => {
    widgets.value = []
    widget.value = null
    loading.value = false
    error.value = null
  }

  /**
   * پاک کردن کش ویجت‌ها
   */
  const clearCache = () => {
    // No-op since cache is removed
  }

  return {
    // State
    widgets: readonly(widgets),
    widget,
    loading: readonly(loading),
    error: readonly(error),

    // Global refresh
    refreshTrigger: readonly(refreshTrigger),
    triggerGlobalRefresh,

    // Actions
    fetchWidgets,
    fetchWidget,
    fetchWidgetByCode,
    fetchWidgetsByPage,
    createWidget,
    updateWidget,
    deleteWidget,
    toggleWidgetStatus,
    duplicateWidget,
    updateWidgetOrder,

    // Utilities
    sortWidgets,
    filterWidgets,
    isCodeUnique,
    getRecentWidgets,
    getActiveWidgets,
    clearError,
    resetState,
    clearCache
  }
}
