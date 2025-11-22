import { usePermissions } from '@/composables/usePermissions'
import { useAuthStore, type User } from '@/stores/auth'
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'

interface LoginResponse {
  token: string
  refresh_token: string
  user: User
}

export const useAuth = () => {
  const router = useRouter()
  const authStore = useAuthStore()
  const { hasPermission: hasPermissionComputed, hasRole: hasRoleComputed } = usePermissions()

  const loading = ref(false)
  const error = ref('')

  // تبدیل computed به function برای سهولت استفاده
  const hasPermission = (permission: string) => hasPermissionComputed.value(permission)
  const hasRole = (role: string) => hasRoleComputed.value(role)

  // دریافت اطلاعات کاربر
  const user = computed(() => authStore.user)
  const isAuthenticated = computed(() => authStore.isAuthenticated)
  const token = computed(() => authStore.accessToken)

  // ورود با رمز عبور
  const loginWithPassword = async (email: string, password: string) => {
    loading.value = true
    error.value = ''

    try {
      const response = await $fetch<LoginResponse>('/api/auth/login-password', {
        method: 'POST',
        body: { email, password },
      })

      authStore.setTokens(response.token, response.token) // استفاده از token برای refresh_token هم
      authStore.setUser(response.user)

      await router.push('/admin')
    } catch (err: unknown) {
      interface ErrorResponse {
        data?: { error?: string }
        statusMessage?: string
        message?: string
        [key: string]: unknown
      }
      const errorObj = err as ErrorResponse
      console.error('❌ Login error:', errorObj)
      error.value = (errorObj?.data?.error) || (errorObj?.statusMessage) || 'ورود ناموفق بود'
    } finally {
      loading.value = false
    }
  }

  // ورود با OTP
  const loginWithOTP = async (mobile: string, code: string) => {
    loading.value = true
    error.value = ''

    try {
      const response = await $fetch<LoginResponse>('/api/auth/verify-otp', {
        method: 'POST',
        body: { mobile, code },
      })

      authStore.setTokens(response.token, response.refresh_token)
      // Ensure user has required properties
      const userData: User = {
        id: String(response.user.id || ''),
        mobile: String(response.user.mobile || ''),
        username: response.user.username ? String(response.user.username) : undefined,
        role_id: response.user.role_id,
        role: response.user.role ? String(response.user.role) : undefined,
        permissions: Array.isArray(response.user.permissions) ? response.user.permissions as string[] : undefined,
        created_at: response.user.created_at ? String(response.user.created_at) : undefined
      }
      authStore.setUser(userData)

      await router.push('/dashboard')
    } catch (err: unknown) {
      const errorObj = err as { data?: { error?: string }; [key: string]: unknown }
      error.value = errorObj.data?.error || 'تایید کد ناموفق بود'
    } finally {
      loading.value = false
    }
  }

  // خروج
  const logout = async () => {
    try {
      await $fetch('/api/auth/logout', {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${authStore.accessToken}`,
        },
      })
    } catch (err) {
      console.error('Logout error:', err)
    } finally {
      authStore.clearAuth()
      await router.push('/auth/login')
    }
  }

  // تازه‌سازی token
  const refreshToken = async () => {
    try {
      const response = await $fetch<{ token: string }>('/api/auth/refresh', {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${authStore.refreshToken}`,
        },
      })

      authStore.setAccessToken(response.token)
      return true
    } catch (_err) {
      authStore.clearAuth()
      await router.push('/auth/login')
      return false
    }
  }

  // دریافت اطلاعات کاربر جاری
  const fetchCurrentUser = async () => {
    try {
      const response = await $fetch<User>('/api/auth/me', {
        headers: {
          Authorization: `Bearer ${authStore.accessToken}`,
        },
      })

      authStore.setUser(response)
    } catch (err) {
      console.error('Fetch user error:', err)
    }
  }

  return {
    loading,
    error,
    user,
    isAuthenticated,
    token,
    loginWithPassword,
    loginWithOTP,
    logout,
    refreshToken,
    fetchCurrentUser,
    hasPermission,
    hasRole,
  }
}
