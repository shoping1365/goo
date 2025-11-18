import { createError, defineEventHandler, getHeader, getMethod } from 'h3'

interface NavigationResponse {
     id: number
     name: string
     platform: string
     [key: string]: unknown
}

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string> }) => Promise<T>

export default defineEventHandler(async (event): Promise<{ success: boolean; data: NavigationResponse[]; message: string }> => {
     try {
          // بررسی متد درخواست
          if (getMethod(event) !== 'GET') {
               throw createError({
                    statusCode: 405,
                    message: 'Method Not Allowed'
               })
          }

          const config = useRuntimeConfig()
          const apiBase = config.public.goApiBase

          // ارسال درخواست به backend Go
          const response = await $fetch<NavigationResponse[]>(`${apiBase}/api/admin/mobile-app-navigation-settings`, {
               method: 'GET',
               headers: {
                    'Authorization': getHeader(event, 'authorization') || ''
               }
          })

          return {
               success: true,
               data: response,
               message: 'لیست ناوبری‌های موبایل دریافت شد'
          }

     } catch (error: unknown) {
          console.error('خطا در دریافت ناوبری‌های موبایل:', error)

          const errorWithStatus = error as { statusCode?: number; message?: string; data?: { message?: string } }
          // اگر خطا از backend باشد
          if (errorWithStatus?.data) {
               throw createError({
                    statusCode: errorWithStatus.statusCode || 500,
                    message: errorWithStatus.data.message || 'خطا در دریافت ناوبری‌های موبایل'
               })
          }

          // اگر خطای شبکه باشد
          const errorMessage = errorWithStatus?.message || ''
          if (errorMessage.includes('fetch') || errorMessage.includes('ECONNREFUSED')) {
               throw createError({
                    statusCode: 503,
                    message: 'خطا در ارتباط با سرور - لطفاً سرور را بررسی کنید'
               })
          }

          // خطای عمومی
          throw createError({
               statusCode: errorWithStatus?.statusCode || 500,
               message: errorWithStatus?.message || 'خطا در دریافت ناوبری‌های موبایل'
          })
     }
})