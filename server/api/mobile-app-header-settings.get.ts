import { defineEventHandler, createError, setResponseHeader } from 'h3'
import { getHeader } from './_utils/getHeader'
import { fetchGo } from './_utils/fetchGo'

interface PublicMobileHeaderResponse {
     success: boolean
     data: any[]
}

export default defineEventHandler(async (event): Promise<PublicMobileHeaderResponse> => {
     try {
          console.log('درخواست دریافت تنظیمات هدر موبایل (عمومی)')

          // دریافت User-Agent از درخواست فرانت‌اند
          const userAgent = getHeader(event, 'user-agent') || ''

          const responseData = await fetchGo(event, '/api/mobile-app-header-settings', {
               method: 'GET',
               headers: userAgent ? { 'User-Agent': userAgent } : undefined
          })

          console.log('پاسخ تنظیمات هدر موبایل:', responseData)

          // جلوگیری از کش شدن پاسخ برای SSR (سازگار با همه محیط‌ها)
          const cacheControl = 'no-store, max-age=0, must-revalidate'
          const nodeRes = event.node?.res as any
          if (nodeRes && typeof nodeRes.setHeader === 'function') {
               nodeRes.setHeader('Cache-Control', cacheControl)
          } else {
               try {
                    setResponseHeader(event, 'Cache-Control', cacheControl)
               } catch (headerErr) {
                    console.warn('[mobile-app-header-settings] Unable to set Cache-Control header', headerErr)
               }
          }

          if (responseData && typeof responseData === 'object') {
               const payload = responseData as Record<string, any>
               return {
                    success: payload.success !== false,
                    data: payload.data || []
               }
          }

          return {
               success: true,
               data: Array.isArray(responseData) ? responseData : []
          }

     } catch (error: any) {
          console.error('خطا در دریافت تنظیمات هدر موبایل:', {
               statusCode: error?.statusCode,
               status: error?.status,
               message: error?.message,
               data: error?.data
          })

          throw createError({
               statusCode: error?.statusCode || error?.status || 500,
               message: error?.data?.message || error?.data?.error || error?.message || 'خطا در دریافت تنظیمات هدر موبایل',
               data: error?.data
          })
     }
})
