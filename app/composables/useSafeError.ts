import { safeGetErrorCode, safeGetErrorDetails, safeGetErrorMessage, safeIsRetryable, safeLogError, safeShowError } from '~/utils/safeErrorHandler';

export const useSafeError = () => {
  return {
    // استخراج پیام خطا به صورت امن
    getMessage: safeGetErrorMessage,

    // استخراج کد خطا به صورت امن
    getCode: safeGetErrorCode,

    // استخراج جزئیات خطا به صورت امن
    getDetails: safeGetErrorDetails,

    // لاگ کردن خطا به صورت امن
    log: safeLogError,

    // نمایش خطا به کاربر به صورت امن
    show: safeShowError,

    // بررسی قابل retry بودن خطا
    isRetryable: safeIsRetryable,

    // تابع wrapper برای try-catch blocks
    async handle<T>(promise: Promise<T>, context?: string): Promise<{ data: T | null; error: unknown }> {
      try {
        const data = await promise
        return { data, error: null }
      } catch (error) {
        safeLogError(error, context)
        return { data: null, error }
      }
    },

    // تابع wrapper برای API calls
    async apiCall<T>(apiFunction: () => Promise<T>, context?: string): Promise<{ data: T | null; error: unknown }> {
      try {
        const data = await apiFunction()
        return { data, error: null }
      } catch (error) {
        const message = safeGetErrorMessage(error)
        const code = safeGetErrorCode(error)

        // لاگ کردن خطا
        safeLogError(error, context)

        // نمایش خطا به کاربر (اگر toast در دسترس باشد)
        const win = typeof window !== 'undefined' ? (window as unknown as { $toast?: { error: (msg: string, opts?: any) => void } }) : null
        if (win && win.$toast) {
          win.$toast.error(message, {
            duration: 5000
          })
        }

        return { data: null, error: { message, code, original: error } }
      }
    }
  }
} 