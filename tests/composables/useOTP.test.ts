import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { useOTP } from '../../app/composables/useOTP'

// Mock $fetch
const mockFetch = vi.fn()
vi.stubGlobal('$fetch', mockFetch)

describe('useOTP Composable', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    vi.useFakeTimers()
  })

  afterEach(() => {
    vi.useRealTimers()
  })

  it('should initialize with default values', () => {
    const otp = useOTP()
    expect(otp.loading.value).toBe(false)
    expect(otp.error.value).toBe('')
    expect(otp.successMessage.value).toBe('')
    expect(otp.remainingTime.value).toBe(0)
  })

  describe('sendOTP', () => {
    it('should send OTP successfully and start timer', async () => {
      mockFetch.mockResolvedValue({ message: 'OTP sent', expires_in: 120 })
      
      const otp = useOTP()
      const result = await otp.sendOTP('09123456789')

      expect(result).toBe(true)
      expect(mockFetch).toHaveBeenCalledWith('/api/auth/send-otp', {
        method: 'POST',
        body: { mobile: '09123456789' }
      })
      expect(otp.successMessage.value).toBe('OTP sent')
      expect(otp.remainingTime.value).toBe(120)
      expect(otp.loading.value).toBe(false)
      expect(otp.error.value).toBe('')

      // Test timer
      vi.advanceTimersByTime(1000)
      expect(otp.remainingTime.value).toBe(119)
      
      vi.advanceTimersByTime(119000)
      expect(otp.remainingTime.value).toBe(0)
    })

    it('should handle send OTP error', async () => {
      mockFetch.mockRejectedValue({ data: { error: 'Invalid mobile' } })
      
      const otp = useOTP()
      const result = await otp.sendOTP('invalid')

      expect(result).toBe(false)
      expect(otp.error.value).toBe('Invalid mobile')
      expect(otp.loading.value).toBe(false)
      expect(otp.successMessage.value).toBe('')
    })

    it('should handle generic error', async () => {
      mockFetch.mockRejectedValue(new Error('Network error'))
      
      const otp = useOTP()
      const result = await otp.sendOTP('09123456789')

      expect(result).toBe(false)
      expect(otp.error.value).toBe('ارسال OTP ناموفق بود')
    })
  })

  describe('resendOTP', () => {
    it('should prevent resend if timer is active', async () => {
      mockFetch.mockResolvedValue({ message: 'OTP sent', expires_in: 60 })
      
      const otp = useOTP()
      await otp.sendOTP('09123456789')
      
      // Timer is now 60
      const result = await otp.resendOTP('09123456789')
      
      expect(result).toBe(false)
      expect(otp.error.value).toContain('صبر کنید')
      expect(mockFetch).toHaveBeenCalledTimes(1) // Only the first call
    })

    it('should allow resend if timer expired', async () => {
      mockFetch.mockResolvedValue({ message: 'OTP sent', expires_in: 60 })
      
      const otp = useOTP()
      await otp.sendOTP('09123456789')
      
      // Fast forward timer
      vi.advanceTimersByTime(61000)
      expect(otp.remainingTime.value).toBe(0)

      // Try resend
      const result = await otp.resendOTP('09123456789')
      
      expect(result).toBe(true)
      expect(mockFetch).toHaveBeenCalledTimes(2)
    })
  })

  it('should reset error', () => {
    const otp = useOTP()
    otp.error.value = 'Some error'
    
    otp.resetError()
    
    expect(otp.error.value).toBe('')
  })
})
