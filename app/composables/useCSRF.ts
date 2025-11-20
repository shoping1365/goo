// Composable برای مدیریت CSRF Token
import { readonly, ref } from 'vue'

// State مشترک برای CSRF token
const csrfToken = ref<string | null>(null)

export const useCSRF = () => {
     /**
      * دریافت CSRF Token از سرور
      */
     const fetchCSRFToken = async (): Promise<string | null> => {
          try {
               if (csrfToken.value) {
                    return csrfToken.value
               }

               const response = await $fetch<{ token: string }>('/api/auth/csrf-token', {
                    credentials: 'include'
               })

               csrfToken.value = response.token
               return response.token
          } catch (error) {
               console.error('خطا در دریافت CSRF token:', error)
               return null
          }
     }

     /**
      * دریافت CSRF Token از cookie
      */
     const getCSRFTokenFromCookie = (): string | null => {
          if (import.meta.client) {
               const cookies = document.cookie.split(';')
               const csrfCookie = cookies.find(cookie =>
                    cookie.trim().startsWith('csrf-token=')
               )

               if (csrfCookie) {
                    return csrfCookie.split('=')[1]
               }
          }
          return null
     }

     /**
      * تنظیم CSRF Token در state
      */
     const setCSRFToken = (token: string) => {
          csrfToken.value = token
     }

     /**
      * پاک کردن CSRF Token
      */
     const clearCSRFToken = () => {
          csrfToken.value = null
     }

     /**
      * دریافت CSRF Token (اول از state، سپس از cookie، سپس از سرور)
      */
     const getCSRFToken = async (): Promise<string | null> => {
          // اول از state بررسی کن
          if (csrfToken.value) {
               return csrfToken.value
          }

          // سپس از cookie بررسی کن
          const cookieToken = getCSRFTokenFromCookie()
          if (cookieToken) {
               csrfToken.value = cookieToken
               return cookieToken
          }

          // در نهایت از سرور دریافت کن
          return await fetchCSRFToken()
     }

     /**
      * Wrapper برای درخواست‌های POST/PUT/DELETE که به طور خودکار CSRF token را اضافه می‌کند
      */
     const fetchWithCSRF = async <T>(
          url: string,
          options: Record<string, unknown> = {}
     ): Promise<T> => {
          const token = await getCSRFToken()

          if (!token) {
               throw new Error('CSRF token در دسترس نیست')
          }

          const headers = {
               'x-csrf-token': token,
               ...(options.headers as Record<string, string>)
          }

          return await $fetch<T>(url, {
               ...options,
               headers,
               credentials: 'include'
          })
     }

     return {
          csrfToken: readonly(csrfToken),
          fetchCSRFToken,
          getCSRFTokenFromCookie,
          setCSRFToken,
          clearCSRFToken,
          getCSRFToken,
          fetchWithCSRF
     }
}
