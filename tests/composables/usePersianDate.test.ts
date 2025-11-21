import { describe, expect, it } from 'vitest'
import { gregorianToJalali, usePersianDate } from '../../app/composables/usePersianDate'

describe('usePersianDate Composable', () => {
  describe('gregorianToJalali', () => {
    it('should convert correctly', () => {
      // 2024-03-20 -> 1402/12/29 (Leap year end)
      expect(gregorianToJalali(2024, 3, 20)).toEqual({ jy: 1403, jm: 1, jd: 1 })
      
      // 2023-03-21 -> 1402/01/01
      expect(gregorianToJalali(2023, 3, 21)).toEqual({ jy: 1402, jm: 1, jd: 1 })
      
      // 2025-01-01 -> 1403/10/12
      expect(gregorianToJalali(2025, 1, 1)).toEqual({ jy: 1403, jm: 10, jd: 12 })
    })
  })

  describe('formatPersianDate', () => {
    const { formatPersianDate } = usePersianDate()

    it('should format YYYY/MM/DD correctly', () => {
      const date = '2025-01-01T10:30:00'
      expect(formatPersianDate(date, 'YYYY/MM/DD')).toBe('1403/10/12')
    })

    it('should format HH:MM correctly', () => {
      const date = '2025-01-01T10:30:00'
      expect(formatPersianDate(date, 'HH:MM')).toBe('10:30')
    })

    it('should format full date time correctly', () => {
      const date = '2025-01-01T10:30:00'
      expect(formatPersianDate(date, 'HH:MM - YYYY/MM/DD')).toBe('10:30 - 1403/10/12')
    })

    it('should handle invalid date', () => {
      expect(formatPersianDate('invalid-date')).toBe('نامشخص')
      expect(formatPersianDate('')).toBe('نامشخص')
    })
  })
})
