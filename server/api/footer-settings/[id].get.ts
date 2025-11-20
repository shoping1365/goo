import { useRuntimeConfig } from '#imports'
import { createError, defineEventHandler, getRouterParam, parseCookies } from 'h3'

interface FooterResponse {
     success: boolean
     data: unknown
}

export default defineEventHandler(async (event): Promise<FooterResponse> => {
     try {
          const id = getRouterParam(event, 'id')
          const config = useRuntimeConfig()

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
               method: 'GET',
               headers: {
                    'Content-Type': 'application/json',
                    'Cookie': cookieHeader
               }
          })

          if (!response.ok) {
               throw new Error(`HTTP error! status: ${response.status}`)
          }

          const responseData = await response.json() as { data: unknown }

          return {
               success: true,
               data: responseData?.data || responseData
          }

     } catch (error: unknown) {
          const err = error as { statusCode?: number; data?: any; message?: string; error?: string }
          console.error('خطا در دریافت فوتر:', err)

          // اگر خطا از سرور Go آمده باشد
          if (err.data) {
               throw createError({
                    statusCode: err.statusCode || 500,
                    message: err.data.message || err.data.error || 'خطا در دریافت فوتر',
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
