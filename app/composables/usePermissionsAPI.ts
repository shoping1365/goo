import { ref, computed } from 'vue'
import { useApiClient } from '@/utils/api'

export interface APIPermission {
  id: string
  name: string
  description?: string
  category?: string
  created_at?: string
}

export const usePermissionsAPI = () => {
  const { api } = useApiClient()

  const permissions = ref<APIPermission[]>([])
  const loading = ref(false)
  const error = ref('')

  // دریافت تمام دسترسی‌ها
  const fetchPermissions = async () => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.get<APIPermission[]>('/api/admin/permissions')
      permissions.value = response
    } catch (err: any) {
      error.value = err.error || 'خطا در دریافت دسترسی‌ها'
    } finally {
      loading.value = false
    }
  }

  // ایجاد دسترسی جدید
  const createPermission = async (
    permData: Omit<APIPermission, 'id' | 'created_at'>
  ): Promise<APIPermission | null> => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.post<APIPermission>('/api/admin/permissions', permData)
      permissions.value.push(response)
      return response
    } catch (err: any) {
      error.value = err.error
      return null
    } finally {
      loading.value = false
    }
  }

  // ویرایش دسترسی
  const updatePermission = async (
    id: string,
    permData: Partial<APIPermission>
  ): Promise<APIPermission | null> => {
    try {
      const response = await api.put<APIPermission>(`/api/admin/permissions/${id}`, permData)
      const index = permissions.value.findIndex(p => p.id === id)
      if (index !== -1) {
        permissions.value[index] = response
      }
      return response
    } catch (err: any) {
      error.value = err.error
      return null
    }
  }

  // حذف دسترسی
  const deletePermission = async (id: string): Promise<boolean> => {
    try {
      await api.delete(`/api/admin/permissions/${id}`)
      permissions.value = permissions.value.filter(p => p.id !== id)
      return true
    } catch (err: any) {
      error.value = err.error
      return false
    }
  }

  // دریافت دسترسی‌ها براساس دسته‌بندی
  const permissionsByCategory = computed(() => {
    const grouped: Record<string, APIPermission[]> = {}
    permissions.value.forEach(perm => {
      const category = perm.category || 'other'
      if (!grouped[category]) {
        grouped[category] = []
      }
      grouped[category].push(perm)
    })
    return grouped
  })

  return {
    permissions: computed(() => permissions.value),
    permissionsByCategory,
    loading: computed(() => loading.value),
    error: computed(() => error.value),
    fetchPermissions,
    createPermission,
    updatePermission,
    deletePermission,
  }
}
