import { useAuth } from '~/composables/useAuth'

export const useApiErrorHandler = () => {
     const { refreshToken, logout } = useAuth()

     // بررسی و مدیریت انقضای توکن
     const handleTokenExpiry = (error: any): boolean => {
          // بررسی خطاهای مربوط به احراز هویت
          if (error?.statusCode === 401 || error?.status === 401) {
               // تلاش برای تازه‌سازی توکن
               refreshToken().then((success) => {
                    if (!success) {
                         // اگر تازه‌سازی ناموفق بود، کاربر را logout کن
                         logout()
                    }
               })
               return true // خطا مدیریت شد
          }
          return false
     }

     // مدیریت خطاهای API
     const handleApiError = (error: any) => {
          console.error('خطای API:', error)

          // بررسی خطای احراز هویت
          if (handleTokenExpiry(error)) {
               return true // خطا مدیریت شد
          }

          // سایر خطاها
          if (error?.statusCode === 500) {
               console.error('خطای سرور داخلی')
               // می‌توانید notification یا toast نمایش دهید
          } else if (error?.statusCode === 404) {
               console.error('منبع یافت نشد')
          } else if (error?.statusCode >= 400 && error?.statusCode < 500) {
               console.error('خطای درخواست:', error.statusMessage)
          }

          return false // خطا مدیریت نشد
     }

     // wrapper برای درخواست‌های API
     const safeApiCall = async <T>(apiCall: () => Promise<T>): Promise<T | null> => {
          try {
               return await apiCall()
          } catch (error: any) {
               if (handleApiError(error)) {
                    return null // خطای احراز هویت مدیریت شد
               }
               throw error // سایر خطاها را دوباره پرتاب کن
          }
     }

     return {
          handleApiError,
          safeApiCall
     }
}

