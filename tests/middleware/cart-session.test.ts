import { beforeEach, describe, expect, it, vi } from 'vitest'
import { ref } from 'vue'
import type { RouteLocationNormalized } from 'vue-router'
import cartSessionMiddleware from '../../app/middleware/cart-session.global'

// Mock dependencies
const mockUseCookie = vi.fn()

vi.mock('nuxt/app', () => ({
  defineNuxtRouteMiddleware: (handler: unknown) => handler,
  useCookie: mockUseCookie
}))

describe('Cart Session Middleware', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    // Mock crypto.getRandomValues
    Object.defineProperty(global, 'crypto', {
      value: {
        getRandomValues: (arr: Uint8Array) => arr.fill(1)
      },
      writable: true
    })
  })

  it('should ignore API routes', () => {
    const to = { path: '/api/some-endpoint' } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    cartSessionMiddleware(to, from)

    expect(mockUseCookie).not.toHaveBeenCalled()
  })

  it('should create session id if missing on client side', () => {
    // Simulate client side
    vi.stubGlobal('import.meta', { client: true })
    
    const cookieValue = ref(null)
    mockUseCookie.mockReturnValue(cookieValue)

    const to = { path: '/' } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    cartSessionMiddleware(to, from)

    expect(mockUseCookie).toHaveBeenCalledWith('session_id')
    expect(cookieValue.value).toMatch(/^cart_\d+_/)
  })

  it('should not create session id if already exists', () => {
    vi.stubGlobal('import.meta', { client: true })
    
    const cookieValue = ref('existing-session')
    mockUseCookie.mockReturnValue(cookieValue)

    const to = { path: '/' } as unknown as RouteLocationNormalized
    const from = {} as unknown as RouteLocationNormalized

    cartSessionMiddleware(to, from)

    expect(cookieValue.value).toBe('existing-session')
  })
})
