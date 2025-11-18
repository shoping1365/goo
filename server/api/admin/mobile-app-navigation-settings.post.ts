import type { H3Event } from 'h3'
import { createError, defineEventHandler, getHeader, getMethod, readBody } from 'h3'

interface NavigationBody {
     name?: string
     description?: string
     platform?: string
     background_color?: string
     text_color?: string
     page_selection?: string
     specific_pages?: string
     excluded_pages?: string
     is_active?: boolean
     navigation_items?: unknown[]
     [key: string]: unknown
}

interface NavigationPayload {
     name: string
     description: string
     platform: string
     background_color: string
     text_color: string
     page_selection: string
     specific_pages: string
     excluded_pages: string
     is_active: boolean
     navigation_items: unknown[]
}

interface NavigationResponse {
     id: number
     name: string
     platform: string
     [key: string]: unknown
}

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string>; body?: unknown }) => Promise<T>

export default defineEventHandler(async (event: H3Event): Promise<{ success: boolean; data: NavigationResponse; message: string }> => {
     try {
          // بررسی متد درخواست
          if (getMethod(event) !== 'POST') {
               throw createError({
                    statusCode: 405,
                    message: 'Method Not Allowed'
               })
          }

          // خواندن داده‌های درخواست
          const body = await readBody(event) as NavigationBody

          // اعتبارسنجی داده‌های ورودی
          if (!body.name || !body.platform) {
               throw createError({
                    statusCode: 400,
                    message: 'نام ناوبری و پلتفرم الزامی است'
               })
          }

          // ساخت payload برای ارسال به backend
          const payload: NavigationPayload = {
               name: body.name,
               description: body.description || '',
               platform: body.platform,
               background_color: body.background_color || '#f8fafc',
               text_color: body.text_color || '#000000',
               page_selection: body.page_selection || 'all',
               specific_pages: body.specific_pages || '',
               excluded_pages: body.excluded_pages || '',
               is_active: body.is_active !== false,
               navigation_items: body.navigation_items || []
          }

          const config = useRuntimeConfig()
          const apiBase = config.public.goApiBase

          // ارسال درخواست به backend Go
          const response = await $fetch<NavigationResponse>(`${apiBase}/api/admin/mobile-app-navigation-settings`, {
               method: 'POST',
               headers: {
                    'Content-Type': 'application/json',
                    'Authorization': getHeader(event, 'authorization') || ''
               },
               body: payload
          })

          return {
               success: true,
               data: response,
               message: 'ناوبری موبایل با موفقیت ایجاد شد'
          }

     } catch (error: unknown) {
          console.error('خطا در ایجاد ناوبری موبایل:', error)

          const errorWithStatus = error as { statusCode?: number; message?: string; data?: { message?: string } }
          // اگر خطا از backend باشد
          if (errorWithStatus?.data) {
               throw createError({
                    statusCode: errorWithStatus.statusCode || 500,
                    message: errorWithStatus.data.message || 'خطا در ایجاد ناوبری موبایل'
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
               message: errorWithStatus?.message || 'خطا در ایجاد ناوبری موبایل'
          })
     }
})