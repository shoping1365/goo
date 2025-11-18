import { ref, computed } from 'vue'
import { useApiClient } from '@/utils/api'

export interface Role {
  id: string
  name: string
  description?: string
  permissions?: string[]
  created_at?: string
  updated_at?: string
}

export const useRoles = () => {
  const { api } = useApiClient()

  const roles = ref<Role[]>([])
  const loading = ref(false)
  const error = ref('')

  // دریافت تمام نقش‌ها
  const fetchRoles = async () => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.get<Role[]>('/api/admin/roles')
      roles.value = response
    } catch (err: any) {
      error.value = err.error || 'خطا در دریافت نقش‌ها'
    } finally {
      loading.value = false
    }
  }

  // دریافت یک نقش
  const getRole = async (id: string): Promise<Role | null> => {
    try {
      return await api.get<Role>(`/api/admin/roles/${id}`)
    } catch (err: any) {
      error.value = err.error
      return null
    }
  }

  // ایجاد نقش جدید
  const createRole = async (roleData: Omit<Role, 'id' | 'created_at' | 'updated_at'>): Promise<Role | null> => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.post<Role>('/api/admin/roles', roleData)
      roles.value.push(response)
      return response
    } catch (err: any) {
      error.value = err.error
      return null
    } finally {
      loading.value = false
    }
  }

  // ویرایش نقش
  const updateRole = async (id: string, roleData: Partial<Role>): Promise<Role | null> => {
    loading.value = true
    error.value = ''

    try {
      const response = await api.put<Role>(`/api/admin/roles/${id}`, roleData)
      const index = roles.value.findIndex(r => r.id === id)
      if (index !== -1) {
        roles.value[index] = response
      }
      return response
    } catch (err: any) {
      error.value = err.error
      return null
    } finally {
      loading.value = false
    }
  }

  // حذف نقش
  const deleteRole = async (id: string): Promise<boolean> => {
    loading.value = true
    error.value = ''

    try {
      await api.delete(`/api/admin/roles/${id}`)
      roles.value = roles.value.filter(r => r.id !== id)
      return true
    } catch (err: any) {
      error.value = err.error
      return false
    } finally {
      loading.value = false
    }
  }

  // دریافت دسترسی‌های نقش
  const getRolePermissions = async (id: string): Promise<string[]> => {
    try {
      const response = await api.get<{
        permissions: string[]
      }>(`/api/admin/roles/${id}/permissions`)
      return response.permissions
    } catch (err: any) {
      error.value = err.error
      return []
    }
  }

  // اختصاص دسترسی به نقش
  const assignPermissionsToRole = async (roleId: string, permissionIds: string[]): Promise<boolean> => {
    try {
      await api.post(`/api/admin/roles/${roleId}/permissions`, {
        permission_ids: permissionIds,
      })
      return true
    } catch (err: any) {
      error.value = err.error
      return false
    }
  }

  return {
    roles: computed(() => roles.value),
    loading: computed(() => loading.value),
    error: computed(() => error.value),
    fetchRoles,
    getRole,
    createRole,
    updateRole,
    deleteRole,
    getRolePermissions,
    assignPermissionsToRole,
  }
}
