import { defineEventHandler, getRouterParam, readBody, parseCookies, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

interface ProductResponse {
     success: boolean
     data: any
}

export default defineEventHandler(async (event): Promise<ProductResponse> => {
     try {
          const id = getRouterParam(event, 'id')
          const body = await readBody(event)
          const config = useRuntimeConfig()

          console.log('درخواست به‌روزرسانی محصول:', id, JSON.stringify(body, null, 2))

          // دریافت cookies از درخواست
          const cookies = parseCookies(event)
          const cookieHeader = Object.entries(cookies)
               .map(([key, value]) => `${key}=${value}`)
               .join('; ')

          // ارسال درخواست به Go backend با fetch معمولی
          const response = await fetch(`${config.public.goApiBase}/api/products/${id}`, {
               method: 'PUT',
               body: JSON.stringify(body),
               headers: {
                    'Content-Type': 'application/json',
                    'Cookie': cookieHeader
               }
          })

          if (!response.ok) {
               throw new Error(`HTTP error! status: ${response.status}`)
          }

          const responseData = await response.json()

          console.log('پاسخ به‌روزرسانی محصول:', responseData)

          return {
               success: true,
               data: responseData?.data || responseData
          }

     } catch (error: any) {
          console.error('خطا در به‌روزرسانی محصول:', error)

          // اگر خطا از سرور Go آمده باشد
          if (error.data) {
               throw createError({
                    statusCode: error.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در به‌روزرسانی محصول',
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
