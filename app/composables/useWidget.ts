import { useRequestHeaders } from 'nuxt/app'
import { computed, readonly, ref } from 'vue'
import type {
  CreateWidgetRequest,
  UpdateWidgetOrderRequest,
  UpdateWidgetRequest,
  Widget,
  WidgetListResponse,
  WidgetPage,
  WidgetResponse,
  WidgetStatus,
  WidgetType
} from '~/types/widget'

// تعریف useAuth برای Nuxt 3
declare const useAuth: any

/**
 * Composable برای مدیریت ابزارک‌ها
 * شامل تمام عملیات CRUD و مدیریت state
 */
export const useWidget = () => {
  // Global refresh trigger
  const refreshTrigger = ref(0)

  // Function to trigger global refresh
  const triggerGlobalRefresh = () => {
    refreshTrigger.value++
  }

  // State management
  const widgets = ref<Widget[]>([])
  const widget = ref<Widget | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Cache for widgets by page
  const widgetCache = new Map<string, { data: Widget[], timestamp: number }>()

  // API Base URL - استفاده از endpoint ادمین برای مدیریت ویجت‌ها
  const baseURL = '/api/admin/widgets'

  // Helper to get auth headers - استفاده از JWT از HttpOnly cookies
  const getAuthHeaders = () => {
    const headers: Record<string, string> = {
      'Content-Type': 'application/json'
    }
    return headers
  }

  /**
   * پاک کردن خطا
   */
  const clearError = () => {
    error.value = null
  }

  /**
   * دریافت لیست تمام ابزارک‌ها
   */
  const fetchWidgets = async (filters?: {
    page?: WidgetPage
    status?: WidgetStatus
  }) => {
    loading.value = true
    error.value = null

    try {
      const params = new URLSearchParams()
      if (filters?.page) params.append('page', filters.page)
      if (filters?.status) params.append('status', filters.status)

      const response = await $fetch(
        `${baseURL}?${params.toString()}`,
        {
          headers: getAuthHeaders(),
          credentials: 'include'
        }
      )

      const responseData = response as WidgetListResponse

      widgets.value = responseData.data
      return responseData.data
    } catch (err: any) {
      error.value = err.message || 'خطا در دریافت ابزارک‌ها'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * دریافت ابزارک با ID
   */
  const fetchWidget = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch(`${baseURL}/${id}?_=${new Date().getTime()}`, {
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      const responseData = response as WidgetResponse
      widget.value = responseData.data
      return responseData.data
    } catch (err: any) {
      error.value = err.message || 'خطا در دریافت ابزارک'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * دریافت ابزارک با کد
   */
  const fetchWidgetByCode = async (code: string) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch(`${baseURL}/code/${code}`, {
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      const responseData = response as WidgetResponse
      widget.value = responseData.data
      return responseData.data
    } catch (err: any) {
      error.value = err.message || 'خطا در دریافت ابزارک'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * دریافت ابزارک‌های یک صفحه خاص
   */
  const fetchWidgetsByPage = (page: WidgetPage) => {
    loading.value = true
    error.value = null

    // استفاده از endpoint عمومی برای نمایش ویجت‌ها
    const url = `/api/widgets/page/${page}?_=${new Date().getTime()}`

    const headers = process.server
      ? useRequestHeaders(['cookie', 'authorization'])
      : undefined

    return $fetch(url, {
      timeout: 0, // حذف کامل timeout
      credentials: 'include',
      headers
    }).then((response: any) => {
      const responseData = response as WidgetListResponse

      widgets.value = responseData.data
      loading.value = false
      return responseData.data
    }).catch((err: any) => {
      // خطا در بارگذاری ابزارک‌ها

      // بررسی نوع خطا
      if (err.statusCode === 404) {
        error.value = 'صفحه مورد نظر یافت نشد'
      } else if (err.statusCode === 500) {
        error.value = 'خطای سرور - لطفاً بعداً تلاش کنید'
      } else if (err.message?.includes('fetch')) {
        error.value = 'خطا در اتصال به سرور بک اند - با دولوپر وب سایت تماس بگیرید'
      } else {
        error.value = err.message || 'خطا در دریافت ابزارک‌ها'
      }

      // در صورت خطا، ویجت‌ها را خالی کن
      widgets.value = []
      loading.value = false
      throw err
    })
  }

  /**
   * ایجاد ابزارک جدید
   */
  const createWidget = async (data: CreateWidgetRequest) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch(baseURL, {
        method: 'POST',
        headers: getAuthHeaders(),
        body: data,
        credentials: 'include'
      })

      const responseData = response as WidgetResponse

      // افزودن به لیست محلی
      widgets.value.push(responseData.data)
      widget.value = responseData.data

      return responseData.data
    } catch (err: any) {
      error.value = err.message || 'خطا در ایجاد ابزارک'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * به‌روزرسانی ابزارک
   */
  const updateWidget = async (id: number, data: UpdateWidgetRequest) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch(`${baseURL}/${id}`, {
        method: 'PUT',
        headers: getAuthHeaders(),
        body: data,
        credentials: 'include'
      })

      const responseData = response as WidgetResponse

      // به‌روزرسانی لیست محلی
      const index = widgets.value.findIndex(w => w.id === id)
      if (index !== -1) {
        widgets.value[index] = responseData.data
      }

      widget.value = responseData.data
      return responseData.data
    } catch (err: any) {
      error.value = err.message || 'خطا در به‌روزرسانی ابزارک'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * حذف ابزارک
   */
  const deleteWidget = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      await $fetch(`${baseURL}/${id}`, {
        method: 'DELETE',
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      // حذف از لیست محلی
      widgets.value = widgets.value.filter(w => w.id !== id)

      if (widget.value?.id === id) {
        widget.value = null
      }

      return true
    } catch (err: any) {
      error.value = err.message || 'خطا در حذف ابزارک'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * تغییر وضعیت ابزارک
   */
  const toggleWidgetStatus = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch(`${baseURL}/${id}/toggle`, {
        method: 'POST',
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      const responseData = response as WidgetResponse

      // به‌روزرسانی لیست محلی
      const index = widgets.value.findIndex(w => w.id === id)
      if (index !== -1) {
        widgets.value[index] = responseData.data
      }

      if (widget.value?.id === id) {
        widget.value = responseData.data
      }

      return responseData.data
    } catch (err: any) {
      error.value = err.message || 'خطا در تغییر وضعیت ابزارک'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * کپی کردن ابزارک
   */
  const duplicateWidget = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch(`${baseURL}/${id}/duplicate`, {
        method: 'POST',
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      const responseData = response as WidgetResponse

      // افزودن به لیست محلی
      widgets.value.push(responseData.data)

      return responseData.data
    } catch (err: any) {
      error.value = err.message || 'خطا در کپی کردن ابزارک'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * به‌روزرسانی ترتیب ابزارک‌ها
   */
  const updateWidgetOrder = async (orderData: UpdateWidgetOrderRequest) => {
    loading.value = true
    error.value = null

    try {
      await $fetch(`${baseURL}/order`, {
        method: 'POST',
        headers: getAuthHeaders(),
        body: orderData,
        credentials: 'include'
      })

      // به‌روزرسانی ترتیب در لیست محلی
      for (const item of orderData.items) {
        const index = widgets.value.findIndex(w => w.id === item.id)
        if (index !== -1) {
          widgets.value[index].order = item.order
        }
      }

      // مرتب‌سازی لیست بر اساس order
      widgets.value.sort((a, b) => a.order - b.order)

      return true
    } catch (err: any) {
      error.value = err.message || 'خطا در به‌روزرسانی ترتیب'
      throw err
    } finally {
      loading.value = false
    }
  }

  /**
   * مرتب‌سازی محلی ابزارک‌ها
   */
  const sortWidgets = (sortBy: 'order' | 'title' | 'created_at' = 'order', direction: 'asc' | 'desc' = 'asc') => {
    widgets.value.sort((a, b) => {
      let aValue: any = a[sortBy]
      let bValue: any = b[sortBy]

      if (sortBy === 'created_at') {
        aValue = new Date(aValue).getTime()
        bValue = new Date(bValue).getTime()
      }

      if (direction === 'asc') {
        return aValue > bValue ? 1 : -1
      } else {
        return aValue < bValue ? 1 : -1
      }
    })
  }

  /**
   * فیلتر کردن ابزارک‌ها
   */
  const filterWidgets = (filters: {
    type?: WidgetType
    status?: WidgetStatus
    page?: WidgetPage
    search?: string
  }) => {
    return computed(() => {
      let filtered = widgets.value

      if (filters.type) {
        filtered = filtered.filter(w => w.type === filters.type)
      }

      if (filters.status) {
        filtered = filtered.filter(w => w.status === filters.status)
      }

      if (filters.page) {
        filtered = filtered.filter(w => w.page === filters.page)
      }

      if (filters.search) {
        const searchLower = filters.search.toLowerCase()
        filtered = filtered.filter(w =>
          w.title.toLowerCase().includes(searchLower) ||
          w.description?.toLowerCase().includes(searchLower) ||
          w.code.toLowerCase().includes(searchLower)
        )
      }

      return filtered
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
    widgetCache.clear()
  }

  /**
   * پاک کردن کش ویجت‌ها
   */
  const clearCache = () => {
    widgetCache.clear()
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