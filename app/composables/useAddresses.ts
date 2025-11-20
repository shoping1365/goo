// Composable مدیریت آدرس‌های کاربر
// API endpoints فرضی: 
//   GET  /api/user/addresses        → لیست آدرس‌ها
//   POST /api/user/addresses        → ایجاد آدرس جدید
//   PUT  /api/user/addresses/:id    → ویرایش آدرس
//   DELETE /api/user/addresses/:id  → حذف آدرس
// (در صورتی که مسیرها متفاوت باشد، فقط اینجا اصلاح کنید)

import { ref } from 'vue'

export interface Address {
  id: number
  full_address?: string
  street?: string
  postal_code?: string
  recipient_name?: string
  recipient_mobile?: string
  phone?: string
  is_default?: boolean
  province?: string
  city?: string
  province_id?: number
  city_id?: number
  [key: string]: unknown
}

interface AddressPayload {
  full_address?: string
  street?: string
  postal_code?: string
  recipient_name?: string
  recipient_mobile?: string
  phone?: string
  is_default?: boolean
  province?: string
  city?: string
  province_id?: number
  city_id?: number
}

export const useAddresses = () => {
  const addresses = ref<Address[]>([])
  const loading = ref(false)
  const lastFetchTime = ref<number>(0)
  const cacheTimeout = 5 * 60 * 1000 // 5 دقیقه

  const fetchAddresses = async (force = false) => {
    const now = Date.now()
    
    // اگر داده‌ها اخیراً دریافت شده‌اند و force=false، از کش استفاده کن
    if (!force && addresses.value.length > 0 && (now - lastFetchTime.value) < cacheTimeout) {
      return addresses.value
    }

    try {
      loading.value = true
      const res = await $fetch<{ data: Address[] } | Address[]>('/api/user/addresses', {
        credentials: 'include'
      })
      // ساختار پاسخ بسته به بک‌اند ممکن است data یا لیست مستقیم باشد
      const rawList = Array.isArray((res as { data: Address[] }).data) ? (res as { data: Address[] }).data : (Array.isArray(res) ? res : [])

      // نرمال‌سازی فیلدها برای سازگاری فرانت و بک‌اند
      const normalized = rawList.map((a: Address) => {
        // اگر street از بک‌اند آمد ولی full_address نبود، کپی کن
        if (a.street && !a.full_address) {
          a.full_address = a.street
        }
        // اگر province / city فیلدها snake_case بود، نگه‌دار برای سازگاری
        return a
      })

      addresses.value = normalized
      lastFetchTime.value = now
      return addresses.value
    } catch (e) {
      console.error('خطا در دریافت آدرس‌ها:', e)
      addresses.value = []
      throw e
    } finally {
      loading.value = false
    }
  }

  const addAddress = async (payload: AddressPayload) => {
    // مطابقت نام فیلدها با ساختار مورد انتظار بک‌اند
    const backendPayload: Record<string, unknown> = {
      ...payload,
      street: payload.street ?? payload.full_address ?? '',
      // اضافه کردن فیلدهای استان و شهر
      province: payload.province,
      city: payload.city,
      province_id: payload.province_id,
      city_id: payload.city_id,
      phone: payload.phone,
    }
    // جلوی ارسال full_address اضافه را بگیر
    delete backendPayload.full_address

    const res = await $fetch<Address>('/api/user/addresses', {
      method: 'POST',
      body: backendPayload,
      credentials: 'include'
    })
    // در صورت موفقیت، به لیست اضافه کن
    if (res && res.id) {
      // اطمینان از داشتن فیلد full_address در لیست لوکال
      if (res.street && !res.full_address) res.full_address = res.street
      addresses.value.push(res)
      // اگر آدرس جدید پیش‌فرض است، سایر آدرس‌ها را غیرفعال کن
      if (payload.is_default) {
        addresses.value.forEach(addr => {
          if (addr.id !== res.id) {
            addr.is_default = false
          }
        })
      }
    }
    return res
  }

  // بروزرسانی یک آدرس بر اساس id
  const updateAddress = async (id: number, payload: Partial<AddressPayload>) => {
    const backendPayload: Record<string, unknown> = {
      ...payload,
      street: payload.street ?? payload.full_address ?? '',
      // اضافه کردن فیلدهای استان و شهر
      province: payload.province,
      city: payload.city,
      province_id: payload.province_id,
      city_id: payload.city_id,
      phone: payload.phone,
    }
    delete backendPayload.full_address

    const res = await $fetch<Address>(`/api/user/addresses/${id}`, {
      method: 'PUT',
      body: backendPayload,
      credentials: 'include'
    })
    // همگام‌سازی لیست در حافظه
    const idx = addresses.value.findIndex(a => a.id === id)
    if (idx > -1) {
      addresses.value[idx] = { ...addresses.value[idx], ...res }
      
      // اگر آدرس پیش‌فرض شده، سایر آدرس‌ها را غیرفعال کن
      if (payload.is_default) {
        addresses.value.forEach(addr => {
          if (addr.id !== id) {
            addr.is_default = false
          }
        })
      }
    }
    return res
  }

  // حذف آدرس بر اساس id
  const deleteAddress = async (id: number) => {
    try {
      await $fetch(`/api/user/addresses/${id}`, {
        method: 'DELETE',
        credentials: 'include'
      })
      addresses.value = addresses.value.filter(a => a.id !== id)
    } catch (error) {
      console.error('خطا در حذف آدرس:', error)
      throw error
    }
  }

  // علامت‌گذاری یک آدرس به عنوان پیش‌فرض
  const setDefaultAddress = async (id: number) => {
    try {
      await updateAddress(id, { is_default: true })
      // مطمئن شو فقط یکی پیش‌فرض است
      addresses.value.forEach(a => { 
        a.is_default = (a.id === id) 
      })
    } catch (error) {
      console.error('خطا در تنظیم آدرس پیش‌فرض:', error)
      throw error
    }
  }

  // پاک‌سازی کش
  const clearCache = () => {
    addresses.value = []
    lastFetchTime.value = 0
  }

  return {
    addresses,
    loading,
    fetchAddresses,
    addAddress,
    updateAddress,
    deleteAddress,
    setDefaultAddress,
    clearCache
  }
} 