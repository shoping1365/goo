import { createError, defineEventHandler, setResponseHeader } from 'h3'
import { fetchGo } from './_utils/fetchGo'
import { getHeader } from './_utils/getHeader'

interface PublicMobileHeaderResponse {
     success: boolean
     data: unknown[]
}

export default defineEventHandler(async (event): Promise<PublicMobileHeaderResponse> => {
     try {
          // دریافت User-Agent از درخواست فرانت‌اند
          const userAgent = getHeader(event, 'user-agent') || ''

          const responseData = await fetchGo(event, '/api/mobile-app-header-settings', {
               method: 'GET',
               headers: userAgent ? { 'User-Agent': userAgent } : undefined
          })

          // جلوگیری از کش شدن پاسخ برای SSR (سازگار با همه محیط‌ها)
          const cacheControl = 'no-store, max-age=0, must-revalidate'
          const nodeRes = event.node?.res
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
               const payload = responseData as Record<string, unknown>
               return {
                    success: payload.success !== false,
                    data: (payload.data as unknown[]) || []
               }
          }

          return {
               success: true,
               data: Array.isArray(responseData) ? responseData : []
          }

     } catch (error: unknown) {
          const err = error as { statusCode?: number; status?: number; message?: string; data?: { message?: string; error?: string } }
          console.error('خطا در دریافت تنظیمات هدر موبایل:', {
               statusCode: err?.statusCode,
               status: err?.status,
               message: err?.message,
               data: err?.data
          })

          throw createError({
               statusCode: err?.statusCode || err?.status || 500,
               message: err?.data?.message || err?.data?.error || err?.message || 'خطا در دریافت تنظیمات هدر موبایل',
               data: err?.data
          })
     }
})
