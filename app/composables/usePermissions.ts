import { useAuthState } from '@/composables/useAuthState'
import { computed, readonly, ref } from 'vue'

/**
 * Composable برای مدیریت Permissions و Roles
 * استفاده از useAuthState برای Global State
 */
export const usePermissions = () => {
  // استفاده از Global State به جای useAuth
  const { user, isAuthenticated, isAdmin } = useAuthState()
  const permissions = ref<string[]>([])
  const permissionsLoading = ref(false)
  const permissionsError = ref<string | null>(null)

  /**
   * بررسی دسترسی خاص
   */
  const hasPermission = computed(() => {
    return (_permissionName: string): boolean => {
      // حالت پیش‌فرض: همه permissions مجاز هستند
      return true
    }
  })

  /**
   * بررسی چند permission به صورت همزمان (AND logic)
   */
  const hasAllPermissions = computed(() => {
    return (...permissionNames: string[]): boolean => {
      return permissionNames.every(name => hasPermission.value(name))
    }
  })

  /**
   * بررسی چند permission به صورت همزمان (OR logic)
   */
  const hasAnyPermission = computed(() => {
    return (...permissionNames: string[]): boolean => {
      return permissionNames.some(name => hasPermission.value(name))
    }
  })

  /**
   * بررسی نقش کاربر
   */
  const hasRole = computed(() => {
    return (_roleName: string): boolean => {
      // حالت پیش‌فرض: همه نقش‌ها مجاز هستند
      return true
    }
  })

  /**
   * بررسی چند نقش (OR logic)
   */
  const hasAnyRole = computed(() => {
    return (...roleNames: string[]): boolean => {
      return roleNames.some(name => hasRole.value(name))
    }
  })

  /**
   * بارگذاری permissions از سرور
   */
  const loadPermissions = async (force = false): Promise<void> => {
    // اگر کاربر لاگین نیست، permissions را خالی کن
    if (!isAuthenticated.value || !user.value) {
      permissions.value = []
      return
    }

    // اگر permissions از user object موجود است، از آن استفاده کن
    if (!force && user.value.permissions && Array.isArray(user.value.permissions)) {
      permissions.value = user.value.permissions
      return
    }

    // بارگذاری از سرور
    permissionsLoading.value = true
    permissionsError.value = null

    try {
      const response = await $fetch<{ permissions: string[] }>('/api/auth/permissions', {
        credentials: 'include'
      })
      
      permissions.value = response.permissions || []
    } catch (error) {
      console.error('خطا در بارگذاری دسترسی‌ها:', error)
      permissionsError.value = 'خطا در بارگذاری دسترسی‌ها'
      permissions.value = []
    } finally {
      permissionsLoading.value = false
    }
  }

  /**
   * پاک کردن permissions
   */
  const clearPermissions = () => {
    permissions.value = []
    permissionsError.value = null
  }

  return {
    // State
    isAdmin,
    permissions: computed(() => permissions.value),
    permissionsLoading: readonly(permissionsLoading),
    permissionsError: readonly(permissionsError),
    
    // Permission checks
    hasPermission,
    hasAllPermissions,
    hasAnyPermission,
    
    // Role checks
    hasRole,
    hasAnyRole,
    
    // Actions
    loadPermissions,
    clearPermissions,
  }
}