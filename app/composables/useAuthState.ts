// State Management مرکزی برای احراز هویت
import { ref, readonly } from 'vue'
import { useRequestHeaders } from 'nuxt/app'

// Global state - shared across all components
const user = ref<any>(null)
const isAuthenticated = ref(false)
const isAdmin = ref(false)
const loading = ref(false)
const error = ref<string | null>(null)
const authCookieNames = ['access_token', 'auth-token', 'refresh_token', 'session_id']

export const useAuthState = () => {
  // محاسبه نقش ادمین
  const calculateIsAdmin = (userData: any): boolean => {
    if (!userData) return false

    const role = (userData.role || '').toLowerCase()
    const username = (userData.username || userData.name || '').toLowerCase()
    const roleId = userData.role_id || 0

    const adminRoles = ['admin', 'developer', 'super_admin', 'manager', 'operator']
    const isAdminRole = adminRoles.includes(role)
    const isDevUser = username === 'dev'
    const isAdminRoleId = roleId > 1 || roleId === 100

    return isAdminRole || isDevUser || isAdminRoleId
  }

  // به‌روزرسانی اطلاعات کاربر
  const setUser = (userData: any) => {
    user.value = userData
    isAuthenticated.value = !!userData
    isAdmin.value = calculateIsAdmin(userData)
  }

  // پاک کردن state احراز هویت
  const clearAuthState = () => {
    user.value = null
    isAuthenticated.value = false
    isAdmin.value = false
    error.value = null
  }

  // بررسی دسترسی
  const hasPermission = (permission: string): boolean => {
    if (!user.value) return false
    if (isAdmin.value) return true

    const permissions = user.value.permissions || []
    return permissions.includes(permission) || permissions.includes('*')
  }

  // بررسی نقش
  const hasRole = (role: string): boolean => {
    if (!user.value) return false
    return user.value.role?.toLowerCase() === role.toLowerCase()
  }

  // دریافت اطلاعات کاربر
  const fetchUser = async (force = false) => {
    if (loading.value && !force) {
      return
    }

    const hasAuthCredentials = () => {
      if (process.server) {
        try {
          const { cookie = '' } = useRequestHeaders(['cookie'])
          return authCookieNames.some(name => cookie?.includes(`${name}=`))
        } catch {
          return false
        }
      }
      if (typeof document !== 'undefined') {
        const cookieStr = document.cookie || ''
        return authCookieNames.some(name => cookieStr.includes(`${name}=`))
      }
      return false
    }

    if (!force && !hasAuthCredentials()) {
      clearAuthState()
      return
    }

    loading.value = true
    error.value = null

    try {
      const res = await $fetch<any>('/api/auth/me')

      if (res?.authenticated && res?.user) {
        if (!res.user.role && res.role) {
          res.user.role = res.role
        }
        if (!res.user.role_id && res.role_id) {
          res.user.role_id = res.role_id
        }

        user.value = res.user
        isAuthenticated.value = true
        isAdmin.value = calculateIsAdmin(res.user)

        console.log('✅ User authenticated:', {
          id: res.user.id,
          username: res.user.username,
          role: res.user.role
        })
      } else {
        console.log('❌ User not authenticated')
        clearAuthState()
      }
    } catch (fetchError) {
      // اگر error بود (مثلاً 404)، فقط user را null کنید
      clearAuthState()
    } finally {
      loading.value = false
    }
  }

  // خروج
  const logout = async () => {
    try {
      await $fetch('/api/auth/logout', {
        method: 'POST',
        credentials: 'include'
      })
    } catch (e) {
      console.warn('Logout request failed:', e)
    }

    clearAuthState()

    if (process.client) {
      document.cookie = 'access_token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT'
      document.cookie = 'auth-token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT'
      document.cookie = 'refresh_token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT'
      document.cookie = 'session_id=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT'
    }

    return true
  }

  return {
    // State (readonly)
    user: readonly(user),
    isAuthenticated: readonly(isAuthenticated),
    isAdmin: readonly(isAdmin),
    loading: readonly(loading),
    error: readonly(error),

    // Actions
    fetchUser,
    setUser,
    clearAuthState,
    hasPermission,
    hasRole,
    logout
  }
}
