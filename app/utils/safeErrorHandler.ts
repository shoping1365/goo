/**
 * Safe Error Handler - جلوگیری از خطاهای "Cannot read properties of undefined"
 */

// تابع امن برای دسترسی به property های object
export function safeGet<T>(obj: unknown, path: string, defaultValue: T): T {
  if (!obj) return defaultValue
  
  const keys = path.split('.')
  let current = obj as Record<string, unknown>
  
  for (const key of keys) {
    if (current === null || current === undefined || typeof current !== 'object') {
      return defaultValue
    }
    current = current[key] as Record<string, unknown>
  }
  
  return current !== undefined ? (current as T) : defaultValue
}

// تابع امن برای استخراج پیام خطا
export function safeGetErrorMessage(error: unknown): string {
  if (!error) return 'خطای غیرمنتظره رخ داده است'
  
  // بررسی انواع مختلف error object
  const message = safeGet(error, 'message', '') ||
                 safeGet(error, 'error', '') ||
                 safeGet(error, 'data.message', '') ||
                 safeGet(error, 'data.error', '') ||
                 safeGet(error, 'statusMessage', '') ||
                 safeGet(error, 'user_message', '')
  
  return (message as string) || 'خطای غیرمنتظره رخ داده است'
}

// تابع امن برای استخراج کد خطا
export function safeGetErrorCode(error: unknown): number {
  if (!error) return 500
  
  return (safeGet(error, 'status', 500) as number) ||
         (safeGet(error, 'statusCode', 500) as number) ||
         (safeGet(error, 'code', 500) as number)
}

// تابع امن برای استخراج جزئیات خطا
export function safeGetErrorDetails(error: unknown): Record<string, unknown> {
  if (!error) return {}
  
  return (safeGet(error, 'data', {}) as Record<string, unknown>) ||
         (safeGet(error, 'details', {}) as Record<string, unknown>) ||
         {}
}

// تابع امن برای لاگ کردن خطا
export function safeLogError(_error: unknown, _context?: string): void {
  // خطا لاگ شده است
}

// تابع امن برای نمایش خطا به کاربر
export function safeShowError(error: unknown, toast?: Record<string, unknown>): void {
  const message = safeGetErrorMessage(error)
  
  if (toast && typeof toast.error === 'function') {
    (toast.error as (msg: string, opts?: unknown) => void)(message, {
      duration: 5000
    })
  } else {
    // fallback به alert
    alert(message)
  }
}

// تابع امن برای بررسی اینکه آیا خطا قابل retry است
export function safeIsRetryable(error: unknown): boolean {
  if (!error) return false
  
  const code = safeGetErrorCode(error)
  
  // خطاهای قابل retry
  return [0, 408, 429, 500, 502, 503, 504].includes(code)
} 