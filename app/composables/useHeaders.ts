import { ref, reactive, readonly } from 'vue'

// تعریف useFetch برای Nuxt 3
declare const useFetch: any

export interface HeaderLayer {
  id?: number
  name: string
  width: number
  height: number
  rowCount: number
  color: string
  opacity: number
  showSeparator: boolean
  separatorType: string
  separatorColor: string
  separatorOpacity: number
  separatorWidth: number
  items: string[]
  created_at?: string
  updated_at?: string
  // برای backward compatibility
  createdAt?: string
  updatedAt?: string
}

export interface Header {
  id?: number
  name: string
  description: string
  pageSelection: string
  specificPages: string
  excludedPages: string
  layers: HeaderLayer[]
  is_active?: boolean
  created_at?: string
  updated_at?: string
  // برای backward compatibility
  createdAt?: string
  updatedAt?: string
}

export const useHeaders = () => {
  const headers = ref<Header[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // بارگذاری لیست هدرها
  const loadHeaders = async () => {
    try {
      loading.value = true
      error.value = null

      const response = await $fetch('/api/admin/header-settings') as any
      if (response?.success) {
        // حذف رکوردهای تکراری بر اساس id
        const map = new Map<number, any>()
        response.data.forEach((h: any) => map.set(h.id, h))
        headers.value = Array.from(map.values())
      } else {
        error.value = response?.message || 'خطا در بارگذاری هدرها'
      }
    } catch (err) {
      console.error('خطا در بارگذاری هدرها:', err)
      error.value = 'خطا در بارگذاری هدرها'
    } finally {
      loading.value = false
    }
  }

  // دریافت هدر خاص
  const getHeader = async (id: number): Promise<Header | null> => {
    try {
      const response = await $fetch(`/api/admin/header-settings/${id}`) as any
      if (response?.success) {
        return response.data
      }
      return null
    } catch (err) {
      console.error('خطا در دریافت هدر:', err)
      return null
    }
  }

  // ایجاد هدر جدید
  const createHeader = async (headerData: Omit<Header, 'id' | 'createdAt' | 'updatedAt'>): Promise<Header | null> => {
    try {
      const response = await $fetch('/api/admin/header-settings', {
        method: 'POST',
        body: headerData
      }) as any

      if (response?.success) {
        // اضافه کردن به لیست محلی بدون ایجاد تکراری
        if (!headers.value.some(h => h.id === response.data.id)) {
          headers.value.push(response.data)
        }
        return response.data
      }
      return null
    } catch (err) {
      console.error('خطا در ایجاد هدر:', err)
      return null
    }
  }

  // به‌روزرسانی هدر
  const updateHeader = async (id: number, headerData: Partial<Header>): Promise<Header | null> => {
    try {
      const response = await $fetch(`/api/admin/header-settings/${id}`, {
        method: 'PUT',
        body: headerData
      }) as any

      if (response?.success) {
        // به‌روزرسانی در لیست محلی
        const index = headers.value.findIndex(h => h.id === id)
        if (index !== -1) {
          headers.value[index] = response.data
        }
        return response.data
      }
      return null
    } catch (err) {
      console.error('خطا در به‌روزرسانی هدر:', err)
      return null
    }
  }

  // فعال/غیرفعال کردن هدر
  const toggleHeaderStatus = async (id: number): Promise<boolean> => {
    try {
      // پیدا کردن هدر فعلی
      const currentHeader = headers.value.find(h => h.id === id)
      if (!currentHeader) {
        console.error('هدر پیدا نشد:', id)
        return false
      }

      // تغییر وضعیت
      const newStatus = !currentHeader.is_active

      const response = await $fetch(`/api/admin/header-settings/${id}`, {
        method: 'PUT',
        body: { is_active: newStatus }
      }) as any

      if (response?.success) {
        // به‌روزرسانی در لیست محلی
        const index = headers.value.findIndex(h => h.id === id)
        if (index !== -1) {
          headers.value[index].is_active = newStatus
        }
        return true
      }
      return false
    } catch (err) {
      console.error('خطا در تغییر وضعیت هدر:', err)
      return false
    }
  }

  // حذف هدر
  const deleteHeader = async (id: number): Promise<boolean> => {
    try {
      const response = await $fetch(`/api/admin/header-settings/${id}`, {
        method: 'DELETE'
      }) as any

      if (response?.success) {
        // حذف از لیست محلی
        headers.value = headers.value.filter(h => h.id !== id)
        return true
      }
      return false
    } catch (err) {
      console.error('خطا در حذف هدر:', err)
      return false
    }
  }

  return {
    headers: readonly(headers),
    loading: readonly(loading),
    error: readonly(error),
    loadHeaders,
    getHeader,
    createHeader,
    updateHeader,
    toggleHeaderStatus,
    deleteHeader
  }
} 