import { createPinia, setActivePinia } from 'pinia'
import { beforeEach, describe, expect, it } from 'vitest'
import { useAuthStore } from '../../app/stores/auth'

describe('Auth Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    // Mock localStorage
    const localStorageMock = (() => {
      let store: Record<string, string> = {}
      return {
        getItem: (key: string) => store[key] || null,
        setItem: (key: string, value: string) => {
          store[key] = value.toString()
        },
        removeItem: (key: string) => {
          delete store[key]
        },
        clear: () => {
          store = {}
        },
      }
    })()
    
    Object.defineProperty(window, 'localStorage', {
      value: localStorageMock,
    })
  })

  it('should have initial state', () => {
    const auth = useAuthStore()
    expect(auth.accessToken).toBeNull()
    expect(auth.refreshToken).toBeNull()
    expect(auth.user).toBeNull()
    expect(auth.isAuthenticated).toBe(false)
  })

  it('should set tokens correctly', () => {
    const auth = useAuthStore()
    const access = 'access-token-123'
    const refresh = 'refresh-token-456'

    auth.setTokens(access, refresh)

    expect(auth.accessToken).toBe(access)
    expect(auth.refreshToken).toBe(refresh)
    expect(localStorage.getItem('accessToken')).toBe(access)
    expect(localStorage.getItem('refreshToken')).toBe(refresh)
  })

  it('should set user correctly', () => {
    const auth = useAuthStore()
    const user = {
      id: '1',
      mobile: '09123456789',
      role: 'user'
    }

    auth.setUser(user)

    expect(auth.user).toEqual(user)
    expect(JSON.parse(localStorage.getItem('user') || '{}')).toEqual(user)
  })

  it('should be authenticated when tokens and user are set', () => {
    const auth = useAuthStore()
    
    auth.setTokens('access', 'refresh')
    auth.setUser({ id: '1', mobile: '09123456789' })

    expect(auth.isAuthenticated).toBe(true)
  })

  it('should clear auth state on logout', () => {
    const auth = useAuthStore()
    
    // Setup initial state
    auth.setTokens('access', 'refresh')
    auth.setUser({ id: '1', mobile: '09123456789' })
    
    // Perform logout
    auth.clearAuth()

    expect(auth.accessToken).toBeNull()
    expect(auth.refreshToken).toBeNull()
    expect(auth.user).toBeNull()
    expect(auth.isAuthenticated).toBe(false)
    
    expect(localStorage.getItem('accessToken')).toBeNull()
    expect(localStorage.getItem('refreshToken')).toBeNull()
    expect(localStorage.getItem('user')).toBeNull()
  })

  it('should load from storage', () => {
    const auth = useAuthStore()
    const storedUser = { id: '1', mobile: '09123456789' }
    
    // Setup localStorage
    localStorage.setItem('accessToken', 'stored-access')
    localStorage.setItem('refreshToken', 'stored-refresh')
    localStorage.setItem('user', JSON.stringify(storedUser))

    auth.loadFromStorage()

    expect(auth.accessToken).toBe('stored-access')
    expect(auth.refreshToken).toBe('stored-refresh')
    expect(auth.user).toEqual(storedUser)
  })
})
