import { defineEventHandler, getRouterParam, createError, getCookie } from 'h3'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase
               // بررسی احراز هویت
               await requireAuth(event)
               const token = getCookie(event, 'access_token')

               const { proxy } = await import('~/server/api/_utils/fetchProxy')
               const page = getRouterParam(event, 'page')

               return proxy(event, `${base}/api/admin/widgets/page/${page}`, {
                    method: 'GET',
                    headers: {
                         'Authorization': `Bearer ${token}`,
                         'Content-Type': 'application/json'
                    }
               })
          } catch (error: unknown) {
               console.error('خطا در دریافت ویجت‌های صفحه:', error)
               throw createError({
                    statusCode: (error as { statusCode?: number }).statusCode || 500,
                    message: (error as { statusMessage?: string }).statusMessage || 'خطا در دریافت ویجت‌های صفحه'
               })
          }
     })
