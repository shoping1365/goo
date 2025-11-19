import { computed, ref } from 'vue'
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

export const useWidgetCore = (baseURL: string) => {
  // State
  const widgets = ref<Widget[]>([])
  const widget = ref<Widget | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Helper to get auth headers
  const getAuthHeaders = () => ({
    'Content-Type': 'application/json'
  })

  const clearError = () => {
    error.value = null
  }

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

      const response = await $fetch<WidgetListResponse>(
        `${baseURL}?${params.toString()}`,
        {
          headers: getAuthHeaders(),
          credentials: 'include'
        }
      )

      widgets.value = response.data
      return response.data
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در دریافت ابزارک‌ها'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const fetchWidget = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch<WidgetResponse>(`${baseURL}/${id}`, {
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      widget.value = response.data
      return response.data
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در دریافت ابزارک'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const fetchWidgetByCode = async (code: string) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch<WidgetResponse>(`${baseURL}/code/${code}`, {
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      widget.value = response.data
      return response.data
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در دریافت ابزارک'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const createWidget = async (data: CreateWidgetRequest) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch<WidgetResponse>(baseURL, {
        method: 'POST',
        headers: getAuthHeaders(),
        body: data,
        credentials: 'include'
      })

      widgets.value.push(response.data)
      widget.value = response.data

      return response.data
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در ایجاد ابزارک'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const updateWidget = async (id: number, data: UpdateWidgetRequest) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch<WidgetResponse>(`${baseURL}/${id}`, {
        method: 'PUT',
        headers: getAuthHeaders(),
        body: data,
        credentials: 'include'
      })

      const index = widgets.value.findIndex(w => w.id === id)
      if (index !== -1) {
        widgets.value[index] = response.data
      }

      widget.value = response.data
      return response.data
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در به‌روزرسانی ابزارک'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const deleteWidget = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      await $fetch(`${baseURL}/${id}`, {
        method: 'DELETE',
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      widgets.value = widgets.value.filter(w => w.id !== id)

      if (widget.value?.id === id) {
        widget.value = null
      }

      return true
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در حذف ابزارک'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const toggleWidgetStatus = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch<WidgetResponse>(`${baseURL}/${id}/toggle`, {
        method: 'POST',
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      const index = widgets.value.findIndex(w => w.id === id)
      if (index !== -1) {
        widgets.value[index] = response.data
      }

      if (widget.value?.id === id) {
        widget.value = response.data
      }

      return response.data
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در تغییر وضعیت ابزارک'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const duplicateWidget = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await $fetch<WidgetResponse>(`${baseURL}/${id}/duplicate`, {
        method: 'POST',
        headers: getAuthHeaders(),
        credentials: 'include'
      })

      widgets.value.push(response.data)

      return response.data
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در کپی کردن ابزارک'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

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

      for (const item of orderData.items) {
        const index = widgets.value.findIndex(w => w.id === item.id)
        if (index !== -1) {
          widgets.value[index].order = item.order
        }
      }

      widgets.value.sort((a, b) => a.order - b.order)

      return true
    } catch (err: unknown) {
      const message = err instanceof Error ? err.message : 'خطا در به‌روزرسانی ترتیب'
      error.value = message
      throw err
    } finally {
      loading.value = false
    }
  }

  const sortWidgets = (sortBy: 'order' | 'title' | 'created_at' = 'order', direction: 'asc' | 'desc' = 'asc') => {
    widgets.value.sort((a, b) => {
      let aValue: string | number | Date = a[sortBy]
      let bValue: string | number | Date = b[sortBy]

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

  const isCodeUnique = (code: string, excludeId?: number) => {
    return !widgets.value.some(w => w.code === code && w.id !== excludeId)
  }

  const getRecentWidgets = (limit: number = 5) => {
    return computed(() => {
      return widgets.value
        .slice()
        .sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
        .slice(0, limit)
    })
  }

  const getActiveWidgets = () => {
    return computed(() => {
      return widgets.value.filter(w => w.status === 'active')
    })
  }

  const resetState = () => {
    widgets.value = []
    widget.value = null
    loading.value = false
    error.value = null
  }

  return {
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
    filterWidgets,
    isCodeUnique,
    getRecentWidgets,
    getActiveWidgets,
    resetState
  }
}
