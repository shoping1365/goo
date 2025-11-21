import { createError, defineEventHandler, getHeader, getQuery } from 'h3'
import { useRuntimeConfig } from 'nuxt/app'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     
     // دریافت query parameters از درخواست
     const query = getQuery(event)

     try {
          // ارسال query parameters به Backend
          const url = `${base}/api/admin/products`
          
          const response = await $fetch(url, {
               headers: {
                    cookie: getHeader(event, 'cookie') || '',
               },
               query: query  // ✅ ارسال page, limit, search, filters
          }) as { data: unknown[]; total: number; page: number; limit: number }

          return response
     } catch {
          throw createError({
               statusCode: 500,
               statusMessage: 'خطا در دریافت محصولات'
          })
     }
})
