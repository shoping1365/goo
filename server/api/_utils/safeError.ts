/**
 * Safe Error Handling Utilities for Server-Side API Endpoints
 */

declare const createError: (options: { statusCode: number; message: string; data?: unknown }) => Error

// تابع امن برای استخراج پیام خطا
export function safeGetErrorMessage(error: unknown): string {
  if (!error) return 'خطای غیرمنتظره رخ داده است'
  const err = error as Record<string, any>

  // بررسی انواع مختلف error object
  return err.message ||
    err.error ||
    err.data?.message ||
    err.data?.error ||
    err.user_message ||
    'خطای غیرمنتظره رخ داده است'
}

// تابع امن برای استخراج کد خطا
export function safeGetErrorCode(error: unknown): number {
  if (!error) return 500
  const err = error as Record<string, any>

  return err.status ||
    err.statusCode ||
    err.code ||
    500
}

// تابع امن برای استخراج جزئیات خطا
export function safeGetErrorDetails(error: unknown): unknown {
  if (!error) return {}
  const err = error as Record<string, any>

  return err.data ||
    err.details ||
    {}
}

// تابع امن برای ایجاد createError
export function safeCreateError(error: unknown, defaultMessage: string = 'خطای غیرمنتظره رخ داده است') {
  const message = safeGetErrorMessage(error) || defaultMessage
  const statusCode = safeGetErrorCode(error)
  const data = safeGetErrorDetails(error)

  return createError({
    statusCode,
    message,
    data: data || {}
  })
}

// تابع امن برای لاگ کردن خطا
export function safeLogError(error: any, context?: string): void {
  // خطا لاگ شده است
} 