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
   * بررسی کامل پروتکل و جلوگیری از URLهای خطرناک
   */
  const isSafeUrl = (url: string): boolean => {
    if (!url || typeof url !== 'string') {
      return false
    }

    try {
      const urlObj = new URL(url)
      const protocol = urlObj.protocol.toLowerCase()
      
      // فقط http و https مجاز هستند
      if (!['http:', 'https:'].includes(protocol)) {
        return false
      }

      // بررسی hostname برای جلوگیری از localhost و IPهای داخلی
      const hostname = urlObj.hostname.toLowerCase()
      if (hostname === 'localhost' || 
          hostname === '127.0.0.1' || 
          hostname === '::1' ||
          hostname === '0.0.0.0' ||
          hostname.startsWith('192.168.') ||
          hostname.startsWith('10.') ||
          hostname.startsWith('172.16.') ||
          hostname.startsWith('172.17.') ||
          hostname.startsWith('172.18.') ||
          hostname.startsWith('172.19.') ||
          hostname.startsWith('172.20.') ||
          hostname.startsWith('172.21.') ||
          hostname.startsWith('172.22.') ||
          hostname.startsWith('172.23.') ||
          hostname.startsWith('172.24.') ||
          hostname.startsWith('172.25.') ||
          hostname.startsWith('172.26.') ||
          hostname.startsWith('172.27.') ||
          hostname.startsWith('172.28.') ||
          hostname.startsWith('172.29.') ||
          hostname.startsWith('172.30.') ||
          hostname.startsWith('172.31.')) {
        return false
      }

      return true
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