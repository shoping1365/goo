// توابع امنیتی برای محافظت از اپلیکیشن

/**
 * پاک کردن ورودی کاربر از کاراکترهای خطرناک
 */
export const useSecurity = () => {
  const sanitizeInput = (input: string): string => {
    if (!input) return ''
    
    return input
      .replace(/[<>]/g, '') // حذف کاراکترهای < >
      .replace(/javascript:/gi, '') // حذف javascript:
      .replace(/on\w+=/gi, '') // حذف event handlers
      .replace(/script/gi, '') // حذف کلمه script
      .trim()
  }

  /**
   * بررسی امنیت پسورد
   */
  const validatePassword = (password: string): { isValid: boolean; errors: string[] } => {
    const errors: string[] = []
    
    if (password.length < 8) {
      errors.push('پسورد باید حداقل 8 کاراکتر باشد')
    }
    
    if (!/[A-Z]/.test(password)) {
      errors.push('پسورد باید شامل حرف بزرگ باشد')
    }
    
    if (!/[a-z]/.test(password)) {
      errors.push('پسورد باید شامل حرف کوچک باشد')
    }
    
    if (!/\d/.test(password)) {
      errors.push('پسورد باید شامل عدد باشد')
    }
    
    if (!/[!@#$%^&*(),.?":{}|<>]/.test(password)) {
      errors.push('پسورد باید شامل کاراکتر خاص باشد')
    }
    
    return {
      isValid: errors.length === 0,
      errors
    }
  }

  /**
   * تولید توکن CSRF
   */
  const generateCSRFToken = (): string => {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    let token = ''
    for (let i = 0; i < 32; i++) {
      token += chars.charAt(Math.floor(Math.random() * chars.length))
    }
    return token
  }

  /**
   * بررسی امنیت URL
   */
  const isSafeUrl = (url: string): boolean => {
    try {
      const urlObj = new URL(url)
      return ['http:', 'https:'].includes(urlObj.protocol)
    } catch {
      return false
    }
  }

  /**
   * رمزگذاری داده‌های حساس
   */
  const encryptData = (data: string): string => {
    // در اینجا می‌توانید از کتابخانه‌های رمزگذاری استفاده کنید
    return btoa(data) // Base64 encoding (فقط برای مثال)
  }

  /**
   * رمزگشایی داده‌های حساس
   */
  const decryptData = (encryptedData: string): string => {
    try {
      return atob(encryptedData) // Base64 decoding (فقط برای مثال)
    } catch {
      return ''
    }
  }

  /**
   * بررسی Rate Limiting
   */
  const checkRateLimit = (key: string, limit: number = 100, window: number = 3600000): boolean => {
    const now = Date.now()
    const requests = JSON.parse(localStorage.getItem(`rate_limit_${key}`) || '[]')
    
    // حذف درخواست‌های قدیمی
    const validRequests = requests.filter((timestamp: number) => now - timestamp < window)
    
    if (validRequests.length >= limit) {
      return false
    }
    
    validRequests.push(now)
    localStorage.setItem(`rate_limit_${key}`, JSON.stringify(validRequests))
    
    return true
  }

  return {
    sanitizeInput,
    validatePassword,
    generateCSRFToken,
    isSafeUrl,
    encryptData,
    decryptData,
    checkRateLimit
  }
} 