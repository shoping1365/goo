import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface User {
  id: string
  mobile: string
  username?: string
  role_id?: string | number
  role?: string
  permissions?: string[]
  created_at?: string
}

export interface AuthState {
  accessToken: string | null
  refreshToken: string | null
  user: User | null
  isAuthenticated: boolean
}

export const useAuthStore = defineStore('auth', () => {
  // State
  const accessToken = ref<string | null>(null)
  const refreshToken = ref<string | null>(null)
  const user = ref<User | null>(null)

  // Computed
  const isAuthenticated = computed(() => !!accessToken.value && !!user.value)

  // Actions
  const setTokens = (access: string, refresh: string) => {
    accessToken.value = access
    refreshToken.value = refresh
    
    // ذخیره در localStorage
    if (process.client) {
      localStorage.setItem('accessToken', access)
      localStorage.setItem('refreshToken', refresh)
    }
  }

  const setAccessToken = (access: string) => {
    accessToken.value = access
    if (process.client) {
      localStorage.setItem('accessToken', access)
    }
  }

  const setUser = (userData: User) => {
    user.value = userData
    if (process.client) {
      localStorage.setItem('user', JSON.stringify(userData))
    }
  }

  const loadFromStorage = () => {
    if (process.client) {
      const stored = {
        access: localStorage.getItem('accessToken'),
        refresh: localStorage.getItem('refreshToken'),
        userData: localStorage.getItem('user'),
      }

      if (stored.access && stored.refresh) {
        accessToken.value = stored.access
        refreshToken.value = stored.refresh
      }

      if (stored.userData) {
        try {
          user.value = JSON.parse(stored.userData)
        } catch (err) {
          console.error('Failed to parse stored user:', err)
        }
      }
    }
  }

  const clearAuth = () => {
    accessToken.value = null
    refreshToken.value = null
    user.value = null

    if (process.client) {
      localStorage.removeItem('accessToken')
      localStorage.removeItem('refreshToken')
      localStorage.removeItem('user')
    }
  }

  return {
    // State
    accessToken,
    refreshToken,
    user,

    // Computed
    isAuthenticated,

    // Actions
    setTokens,
    setAccessToken,
    setUser,
    loadFromStorage,
    clearAuth,
  }
})
