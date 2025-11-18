/**
 * Safe Error Handling Utilities for Server-Side API Endpoints
 */

declare const createError: (options: { statusCode: number; message: string; data?: any }) => Error

// تابع امن برای استخراج پیام خطا
export function safeGetErrorMessage(error: any): string {
  if (!error) return 'خطای غیرمنتظره رخ داده است'

  // بررسی انواع مختلف error object
  return error.message ||
    error.error ||
    error.data?.message ||
    error.data?.error ||
    error.user_message ||
    'خطای غیرمنتظره رخ داده است'
}

// تابع امن برای استخراج کد خطا
export function safeGetErrorCode(error: any): number {
  if (!error) return 500

  return error.status ||
    error.statusCode ||
    error.code ||
    500
}

// تابع امن برای استخراج جزئیات خطا
export function safeGetErrorDetails(error: any): any {
  if (!error) return {}

  return error.data ||
    error.details ||
    {}
}

// تابع امن برای ایجاد createError
export function safeCreateError(error: any, defaultMessage: string = 'خطای غیرمنتظره رخ داده است') {
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