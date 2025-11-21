import { beforeEach, describe, expect, it, vi } from 'vitest'
import type { RouteLocationNormalized } from 'vue-router'
import authMiddleware from '../../app/middleware/auth'

// Mock dependencies
const mockNavigateTo = vi.fn()
const mockLoadFromStorage = vi.fn()

const mockAuthStore = {
  isAuthenticated: false,
  user: null as Record<string, unknown> | null,
  loadFromStorage: mockLoadFromStorage
}

vi.stubGlobal('navigateTo', mockNavigateTo)
vi.mock('@/stores/auth', () => ({
  useAuthStore: () => mockAuthStore
}))
vi.mock('nuxt/app', () => ({
  defineNuxtRouteMiddleware: (handler: unknown) => handler
}))

describe('Auth Middleware', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    mockAuthStore.isAuthenticated = false
    mockAuthStore.user = null
  })

  it('should redirect to login if route requires auth and user is not authenticated', () => {
    const to = { meta: { requiresAuth: true }, path: '/protected' } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    authMiddleware(to, from)

    expect(mockLoadFromStorage).toHaveBeenCalled()
    expect(mockNavigateTo).toHaveBeenCalledWith('/auth/login')
  })

  it('should allow access if route requires auth and user is authenticated', () => {
    mockAuthStore.isAuthenticated = true
    const to = { meta: { requiresAuth: true }, path: '/protected' } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    authMiddleware(to, from)

    expect(mockNavigateTo).not.toHaveBeenCalled()
  })

  it('should redirect logged in user from login page to account', () => {
    mockAuthStore.isAuthenticated = true
    mockAuthStore.user = { role: 'user' }
    const to = { path: '/auth/login', meta: {} } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    authMiddleware(to, from)

    expect(mockNavigateTo).toHaveBeenCalledWith('/account')
  })

  it('should redirect admin user from login page to admin dashboard', () => {
    mockAuthStore.isAuthenticated = true
    mockAuthStore.user = { role: 'admin' }
    const to = { path: '/auth/login', meta: {} } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    authMiddleware(to, from)

    expect(mockNavigateTo).toHaveBeenCalledWith('/admin')
  })

  it('should redirect to 403 if route requires admin and user is not admin', () => {
    mockAuthStore.isAuthenticated = true
    mockAuthStore.user = { role: 'user' }
    const to = { meta: { requiresAdmin: true }, path: '/admin/users' } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    authMiddleware(to, from)

    expect(mockNavigateTo).toHaveBeenCalledWith('/403')
  })
})
