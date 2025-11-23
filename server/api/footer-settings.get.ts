import { createError, defineEventHandler, setResponseHeader } from 'h3'
import { fetchGo } from './_utils/fetchGo'

interface PublicFooterResponse {
     success: boolean
     data: unknown[]
}

export default defineEventHandler(async (event): Promise<PublicFooterResponse> => {
     try {
          const responseData = await fetchGo(event, '/api/footer-settings', {
               method: 'GET'
          })

          const cacheControl = 'no-store, max-age=0, must-revalidate'
          const nodeRes = event.node?.res
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
          const err = error as { statusCode?: number; status?: number; message?: string; data?: unknown }
          console.error('خطا در دریافت تنظیمات فوتر:', {
               statusCode: err?.statusCode,
               status: err?.status,
               message: err?.message,
               data: err?.data
          })

          throw createError({
               statusCode: err?.statusCode || err?.status || 500,
               message: (err?.data as any)?.message || (err?.data as any)?.error || err?.message || 'خطا در دریافت تنظیمات فوتر',
               data: error?.data
          })
     }
})
