import { defineEventHandler, getRouterParam, parseCookies, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

interface DeleteResponse {
     success: boolean
     message: string
}

export default defineEventHandler(async (event): Promise<DeleteResponse> => {
     try {
          const id = getRouterParam(event, 'id')
          const config = useRuntimeConfig()

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

          const responseData = await response.json() as any

          console.log('پاسخ حذف فوتر:', responseData)

          return {
               success: true,
               message: 'فوتر با موفقیت حذف شد'
          }

     } catch (error: any) {
          console.error('خطا در حذف فوتر:', error)

          // اگر خطا از سرور Go آمده باشد
          if (error.data) {
               throw createError({
                    statusCode: error.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در حذف فوتر',
                    data: error.data
               })
          }

          // اگر خطای شبکه باشد
          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})
