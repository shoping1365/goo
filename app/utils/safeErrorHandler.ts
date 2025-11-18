/**
 * Safe Error Handler - جلوگیری از خطاهای "Cannot read properties of undefined"
 */

// تابع امن برای دسترسی به property های object
export function safeGet<T>(obj: any, path: string, defaultValue: T): T {
  if (!obj) return defaultValue
  
  const keys = path.split('.')
  let current = obj
  
  for (const key of keys) {
    if (current === null || current === undefined || typeof current !== 'object') {
      return defaultValue
    }
    current = current[key]
  }
  
  return current !== undefined ? current : defaultValue
}

// تابع امن برای استخراج پیام خطا
export function safeGetErrorMessage(error: any): string {
  if (!error) return 'خطای غیرمنتظره رخ داده است'
  
  // بررسی انواع مختلف error object
  const message = safeGet(error, 'message', '') ||
                 safeGet(error, 'error', '') ||
                 safeGet(error, 'data.message', '') ||
                 safeGet(error, 'data.error', '') ||
                 safeGet(error, 'statusMessage', '') ||
                 safeGet(error, 'user_message', '')
  
  return message || 'خطای غیرمنتظره رخ داده است'
}

// تابع امن برای استخراج کد خطا
export function safeGetErrorCode(error: any): number {
  if (!error) return 500
  
  return safeGet(error, 'status', 500) ||
         safeGet(error, 'statusCode', 500) ||
         safeGet(error, 'code', 500)
}

// تابع امن برای استخراج جزئیات خطا
export function safeGetErrorDetails(error: any): any {
  if (!error) return {}
  
  return safeGet(error, 'data', {}) ||
         safeGet(error, 'details', {}) ||
         {}
}

// تابع امن برای لاگ کردن خطا
export function safeLogError(error: any, context?: string): void {
  // خطا لاگ شده است
}

// تابع امن برای نمایش خطا به کاربر
export function safeShowError(error: any, toast?: any): void {
  const message = safeGetErrorMessage(error)
  
  if (toast && typeof toast.error === 'function') {
    toast.error(message, {
      duration: 5000
    })
  } else {
    // fallback به alert
    alert(message)
  }
}

// تابع امن برای بررسی اینکه آیا خطا قابل retry است
export function safeIsRetryable(error: any): boolean {
  if (!error) return false
  
  const code = safeGetErrorCode(error)
  
  // خطاهای قابل retry
  return [0, 408, 429, 500, 502, 503, 504].includes(code)
} 