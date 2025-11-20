import type { User } from '@/stores/auth'
import { useApiClient } from '@/utils/api'
import { computed, ref } from 'vue'

interface ApiError {
  error?: string
  message?: string
}

export const useUsers = () => {
  const { api } = useApiClient()
  
  const users = ref<User[]>([])
  const loading = ref(false)
  const error = ref('')
  const totalCount = ref(0)

  // دریافت تمام کاربران
  const fetchUsers = async (page: number = 1, limit: number = 20) => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.get<{
        data: User[]
        total: number
      }>(`/api/admin/users?page=${page}&limit=${limit}`)

      users.value = response.data
      totalCount.value = response.total
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در دریافت کاربران'
    } finally {
      loading.value = false
    }
  }

  // دریافت یک کاربر
  const getUser = async (id: string): Promise<User | null> => {
    try {
      return await api.get<User>(`/api/admin/users/${id}`)
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در دریافت کاربر'
      return null
    }
  }

  // جستجوی کاربران
  const searchUsers = async (query: string) => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.get<User[]>(
        `/api/admin/users/search?q=${encodeURIComponent(query)}`
      )
      users.value = response
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در جستجو'
    } finally {
      loading.value = false
    }
  }

  // ایجاد کاربر
  const createUser = async (userData: Partial<User>): Promise<User | null> => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.post<User>('/api/admin/users', userData)
      users.value.push(response)
      return response
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در ایجاد کاربر'
      return null
    } finally {
      loading.value = false
    }
  }

  // ویرایش کاربر
  const updateUser = async (id: string, userData: Partial<User>): Promise<User | null> => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.put<User>(`/api/admin/users/${id}`, userData)
      const index = users.value.findIndex(u => u.id === id)
      if (index !== -1) {
        users.value[index] = response
      }
      return response
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در ویرایش کاربر'
      return null
    } finally {
      loading.value = false
    }
  }

  // حذف کاربر
  const deleteUser = async (id: string): Promise<boolean> => {
    loading.value = true
    error.value = ''

    try {
      await api.delete(`/api/admin/users/${id}`)
      users.value = users.value.filter(u => u.id !== id)
      return true
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در حذف کاربر'
      return false
    } finally {
      loading.value = false
    }
  }

  // تغییر نقش کاربر
  const changeUserRole = async (id: string, roleId: string): Promise<boolean> => {
    try {
      const response = await api.put<User>(`/api/admin/users/${id}/role`, {
        role_id: roleId,
      })
      const index = users.value.findIndex(u => u.id === id)
      if (index !== -1) {
        users.value[index] = response
      }
      return true
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در تغییر نقش'
      return false
    }
  }

  // دسترسی‌های کاربر
  const getUserPermissions = async (id: string): Promise<string[]> => {
    try {
      const response = await api.get<{
        permissions: string[]
      }>(`/api/admin/users/${id}/permissions`)
      return response.permissions
    } catch (err) {
      const e = err as ApiError
      error.value = e.error || 'خطا در دریافت دسترسی‌ها'
      return []
    }
  }

  const filteredUsers = computed(() => {
    return users.value
  })

  return {
    users: computed(() => users.value),
    filteredUsers,
    loading: computed(() => loading.value),
    error: computed(() => error.value),
    totalCount: computed(() => totalCount.value),
    fetchUsers,
    getUser,
    searchUsers,
    createUser,
    updateUser,
    deleteUser,
    changeUserRole,
    getUserPermissions,
  }
}
