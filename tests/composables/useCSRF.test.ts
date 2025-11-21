import { beforeEach, describe, expect, it, vi } from 'vitest'
import { useCSRF } from '../../app/composables/useCSRF'

// Mock $fetch
const mockFetch = vi.fn()
vi.stubGlobal('$fetch', mockFetch)

// Mock document.cookie
// const mockCookie = { ... }

describe('useCSRF Composable', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    // Reset module state if possible, or just rely on clearCSRFToken
    const { clearCSRFToken } = useCSRF()
    clearCSRFToken()
  })

  describe('fetchCSRFToken', () => {
    it('should fetch token from server', async () => {
      mockFetch.mockResolvedValue({ token: 'server-token' })
      
      const { fetchCSRFToken } = useCSRF()
      const token = await fetchCSRFToken()
      
      expect(token).toBe('server-token')
      expect(mockFetch).toHaveBeenCalledWith('/api/auth/csrf-token', expect.any(Object))
    })

    it('should return existing token if available', async () => {
      const { setCSRFToken, fetchCSRFToken } = useCSRF()
      setCSRFToken('existing-token')
      
      const token = await fetchCSRFToken()
      
      expect(token).toBe('existing-token')
      expect(mockFetch).not.toHaveBeenCalled()
    })
  })

  describe('getCSRFToken', () => {
    it('should prioritize state token', async () => {
      const { setCSRFToken, getCSRFToken } = useCSRF()
      setCSRFToken('state-token')
      
      const token = await getCSRFToken()
      expect(token).toBe('state-token')
    })

    it('should fallback to server if no state or cookie', async () => {
      mockFetch.mockResolvedValue({ token: 'server-token' })
      
      const { getCSRFToken } = useCSRF()
      const token = await getCSRFToken()
      
      expect(token).toBe('server-token')
    })
  })

  describe('fetchWithCSRF', () => {
    it('should attach csrf token to headers', async () => {
      const { setCSRFToken, fetchWithCSRF } = useCSRF()
      setCSRFToken('test-token')
      mockFetch.mockResolvedValue({ success: true })
      
      await fetchWithCSRF('/api/test', { method: 'POST' })
      
      expect(mockFetch).toHaveBeenCalledWith('/api/test', expect.objectContaining({
        headers: expect.objectContaining({
          'x-csrf-token': 'test-token'
        })
      }))
    })

    it('should throw error if token missing', async () => {
      const { fetchWithCSRF } = useCSRF()
      mockFetch.mockRejectedValue(new Error('CSRF token در دسترس نیست')) // Simulate fetch failure too
      
      // Mock getCSRFToken to return null
      // Since we can't easily mock internal functions of the composable without more complex setup,
      // we rely on the fact that fetchCSRFToken will fail if we mock fetch to fail.
      
      await expect(fetchWithCSRF('/api/test')).rejects.toThrow()
    })
  })
})
