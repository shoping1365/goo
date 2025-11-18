import { defineEventHandler, createError, setResponseHeader } from 'h3'
import { fetchGo } from './_utils/fetchGo'

interface PublicFooterResponse {
     success: boolean
     data: any[]
}

export default defineEventHandler(async (event): Promise<PublicFooterResponse> => {
     try {
          console.log('درخواست دریافت تنظیمات فوتر (عمومی)')

          const responseData = await fetchGo(event, '/api/footer-settings', {
               method: 'GET'
          })

          console.log('پاسخ تنظیمات فوتر:', responseData)

          const cacheControl = 'no-store, max-age=0, must-revalidate'
          const nodeRes = event.node?.res as any
          if (nodeRes && typeof nodeRes.setHeader === 'function') {
               nodeRes.setHeader('Cache-Control', cacheControl)
          } else {
               try {
                    setResponseHeader(event, 'Cache-Control', cacheControl)
               } catch (headerErr) {
                    console.warn('[footer-settings] Unable to set Cache-Control header', headerErr)
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
          console.error('خطا در دریافت تنظیمات فوتر:', {
               statusCode: error?.statusCode,
               status: error?.status,
               message: error?.message,
               data: error?.data
          })

          throw createError({
               statusCode: error?.statusCode || error?.status || 500,
               message: error?.data?.message || error?.data?.error || error?.message || 'خطا در دریافت تنظیمات فوتر',
               data: error?.data
          })
     }
})
