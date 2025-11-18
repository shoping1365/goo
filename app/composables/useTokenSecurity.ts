/**
 * Token Security Composable
 * Handle token expiry, refresh, and validation
 */

import { ref, computed } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { useApiClient } from '~/utils/api'

interface TokenPayload {
  sub: string
  email: string
  roles: string[]
  iat: number
  exp: number
}

const tokenRefreshThreshold = 5 * 60 * 1000 // 5 minutes
const refreshTokenTimeout = ref<NodeJS.Timeout | null>(null)

/**
 * Decode JWT token (without verification)
 */
export const decodeToken = (token: string): TokenPayload | null => {
  try {
    const parts = token.split('.')
    if (parts.length !== 3) return null

    const decoded = JSON.parse(atob(parts[1]))
    return decoded as TokenPayload
  } catch {
    return null
  }
}

/**
 * Check if token is expired
 */
export const isTokenExpired = (token: string): boolean => {
  const payload = decodeToken(token)
  if (!payload) return true

  const now = Math.floor(Date.now() / 1000)
  return payload.exp <= now
}

/**
 * Check if token is expiring soon
 */
export const isTokenExpiringSoon = (token: string): boolean => {
  const payload = decodeToken(token)
  if (!payload) return true

  const now = Math.floor(Date.now() / 1000)
  const expiresIn = (payload.exp - now) * 1000 // Convert to milliseconds

  return expiresIn <= tokenRefreshThreshold
}

/**
 * Get time remaining until token expires
 */
export const getTokenTimeRemaining = (token: string): number => {
  const payload = decodeToken(token)
  if (!payload) return 0

  const now = Math.floor(Date.now() / 1000)
  return Math.max(0, payload.exp - now)
}

/**
 * Use Token Security Composable
 */
export const useTokenSecurity = () => {
  const authStore = useAuthStore()

  /**
   * Setup automatic token refresh
   */
  const setupTokenRefresh = () => {
    clearTokenRefresh()

    const token = authStore.accessToken
    if (!token || isTokenExpired(token)) return

    const timeRemaining = getTokenTimeRemaining(token)
    const refreshTime = Math.max(0, (timeRemaining - tokenRefreshThreshold / 1000) * 1000)

    if (refreshTime > 0) {
      refreshTokenTimeout.value = setTimeout(() => {
        refreshAccessToken()
      }, refreshTime)
    }
  }

  /**
   * Refresh access token
   */
  const refreshAccessToken = async () => {
    try {
      const refreshToken = authStore.refreshToken
      if (!refreshToken) {
        authStore.clearAuth()
        return
      }

      const api = useApiClient()
      const response = await api.api.post<any>('/auth/refresh', {
        refresh_token: refreshToken,
      })

      if ((response as any)?.access_token) {
        authStore.setAccessToken((response as any).access_token)

        if ((response as any)?.refresh_token) {
          authStore.refreshToken = (response as any).refresh_token as any
        }

        setupTokenRefresh()
      }
    } catch (error) {
      console.error('Token refresh failed:', error)
      authStore.clearAuth()
    }
  }

  /**
   * Clear token refresh timeout
   */
  const clearTokenRefresh = () => {
    if (refreshTokenTimeout.value) {
      clearTimeout(refreshTokenTimeout.value)
      refreshTokenTimeout.value = null
    }
  }

  /**
   * Validate token before API call
   */
  const validateToken = async (): Promise<boolean> => {
    const token = authStore.accessToken as string

    if (!token || isTokenExpired(token)) {
      authStore.clearAuth()
      return false
    }

    if (isTokenExpiringSoon(token)) {
      await refreshAccessToken()
    }

    return true
  }

  /**
   * Get current user from token
   */
  const getCurrentUser = () => {
    const token = authStore.accessToken as string
    if (!token) return null

    const payload = decodeToken(token)
    return payload
  }

  /**
   * Check if user has specific role
   */
  const hasRole = (roleName: string): boolean => {
    const user = getCurrentUser()
    return user?.roles?.includes(roleName) ?? false
  }

  /**
   * Check if user has any of the roles
   */
  const hasAnyRole = (roleNames: string[]): boolean => {
    const user = getCurrentUser()
    return roleNames.some(role => user?.roles?.includes(role))
  }

  return {
    setupTokenRefresh,
    refreshAccessToken,
    clearTokenRefresh,
    validateToken,
    getCurrentUser,
    hasRole,
    hasAnyRole,
    isTokenExpired,
    isTokenExpiringSoon,
    getTokenTimeRemaining,
  }
}
