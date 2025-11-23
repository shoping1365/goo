import { useRuntimeConfig } from '#imports'
import { createError, defineEventHandler, getRouterParam, parseCookies } from 'h3'

interface DeleteResponse {
     success: boolean
     message: string
}

export default defineEventHandler(async (event): Promise<DeleteResponse> => {
     try {
          const id = getRouterParam(event, 'id')
          const config = useRuntimeConfig()

          // eslint-disable-next-line no-console
          console.log('درخواست حذف فوتر:', id)

          if (!id) {
               throw createError({
                    statusCode: 400,
                    message: 'شناسه فوتر الزامی است'
               })
          }

          // دریافت cookies از درخواست
          const cookies = parseCookies(event)
          const cookieHeader = Object.entries(cookies)
               .map(([key, value]) => `${key}=${value}`)
               .join('; ')

          // ارسال درخواست به Go backend با fetch معمولی
          const response = await fetch(`${config.public.goApiBase}/api/admin/footer-settings/${id}`, {
               method: 'DELETE',
               headers: {
                    'Content-Type': 'application/json',
                    'Cookie': cookieHeader
               }
          })

          if (!response.ok) {
               throw new Error(`HTTP error! status: ${response.status}`)
          }

          // eslint-disable-next-line @typescript-eslint/no-unused-vars
          const responseData = await response.json() as unknown

          return {
               success: true,
               message: 'فوتر با موفقیت حذف شد'
          }

     } catch (error: unknown) {
          console.error('خطا در حذف فوتر:', error)

          const err = error as { data?: unknown; statusCode?: number; message?: string; error?: string }
          // اگر خطا از سرور Go آمده باشد
          if (err.data) {
               throw createError({
                    statusCode: err.statusCode || 500,
                    message: (err.data as any).message || (err.data as any).error || 'خطا در حذف فوتر',
                    data: err.data
               })
          }

          // اگر خطای شبکه باشد
          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})
