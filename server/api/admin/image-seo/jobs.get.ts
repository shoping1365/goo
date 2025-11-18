import type { H3Event } from 'h3'
import { createError, defineEventHandler, getHeader, getQuery } from 'h3'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

// لیست jobهای Image SEO برای مانیتورینگ در پنل
export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const backendUrl = config.public.goApiBase
     const query = getQuery(event)

     try {
          const queryRecord: Record<string, string> = {}
          for (const [key, value] of Object.entries(query)) {
               if (value !== undefined) {
                    queryRecord[key] = String(value)
               }
          }
          const queryString = new URLSearchParams(queryRecord).toString()
          const url = queryString ? `${backendUrl}/api/admin/image-seo/jobs?${queryString}` : `${backendUrl}/api/admin/image-seo/jobs`

          const response = await fetch(url, {
               headers: {
                    'Content-Type': 'application/json',
                    ...getHeaders(event)
               }
          })

          if (!response.ok) {
               throw createError({
                    statusCode: response.status,
                    message: response.statusText
               })
          }

          return await response.json()
     } catch (error: unknown) {
          console.error('Error fetching image-seo jobs:', error)
          const errorWithStatus = error as { statusCode?: number; message?: string }
          throw createError({
               statusCode: errorWithStatus?.statusCode || 500,
               message: errorWithStatus?.message || 'خطا در دریافت مشاغل image-seo'
          })
     }
})

function getHeaders(event: H3Event): Record<string, string> {
     const headers: Record<string, string> = {}

     // Forward important headers
     const cookie = getHeader(event, 'cookie')
     if (cookie) headers.cookie = cookie

     const authorization = getHeader(event, 'authorization')
     if (authorization) headers.authorization = authorization

     return headers
}


