import { computed } from 'vue'
import { useAuthStore } from '@/stores/auth'

export const usePermission = () => {
  const authStore = useAuthStore()

  // بررسی اینکه آیا کاربر admin است
  const isAdmin = computed(() => {
    const adminRoles = ['admin', 'developer']
    return authStore.user?.role_id === 100 || 
           authStore.user?.username === 'dev' ||
           adminRoles.includes(authStore.user?.role?.toLowerCase() || '')
  })

  // بررسی دسترسی مخصوص
  const hasPermission = (permission: string): boolean => {
    if (!authStore.user) return false
    if (authStore.user.role === 'admin') return true // admins همه چیز رو دارند
    return authStore.user?.permissions?.includes(permission) || false
  }

  // بررسی چند تا دسترسی
  const hasPermissions = (permissions: string[]): boolean => {
    return permissions.every(perm => hasPermission(perm))
  }

  // بررسی حداقل یک دسترسی
  const hasAnyPermission = (permissions: string[]): boolean => {
    return permissions.some(perm => hasPermission(perm))
  }

  // دریافت تمام دسترسی‌های کاربر
  const getAllPermissions = computed(() => {
    return authStore.user?.permissions || []
  })

  // بررسی نقش
  const hasRole = (role: string): boolean => {
    return authStore.user?.role === role
  }

  // بررسی چند نقش
  const hasAnyRole = (roles: string[]): boolean => {
    return roles.includes(authStore.user?.role || '')
  }

  return {
    isAdmin,
    hasPermission,
    hasPermissions,
    hasAnyPermission,
    getAllPermissions,
    hasRole,
    hasAnyRole,
  }
}
