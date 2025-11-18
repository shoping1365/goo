import { defineEventHandler, getRouterParam, parseCookies, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

interface FooterResponse {
     success: boolean
     data: any
}

export default defineEventHandler(async (event): Promise<FooterResponse> => {
     try {
          const id = getRouterParam(event, 'id')
          const config = useRuntimeConfig()

          console.log('درخواست دریافت فوتر:', id)

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

          const responseData = await response.json() as any

          console.log('پاسخ دریافت فوتر:', responseData)

          return {
               success: true,
               data: responseData?.data || responseData
          }

     } catch (error: any) {
          console.error('خطا در دریافت فوتر:', error)

          // اگر خطا از سرور Go آمده باشد
          if (error.data) {
               throw createError({
                    statusCode: error.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در دریافت فوتر',
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
