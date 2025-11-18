import { createError, defineEventHandler, getHeader, getQuery } from 'h3'
import { useRuntimeConfig } from 'nuxt/app'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     
     // دریافت query parameters از درخواست
     const query = getQuery(event)
     
     console.log('🔍 Nuxt API Route - Query params:', query)

     try {
          // ارسال query parameters به Backend
          const url = `${base}/api/admin/products`
          console.log('📡 Nuxt API Route - Sending to:', url, 'with query:', query)
          
          const response = await $fetch(url, {
               headers: {
                    cookie: getHeader(event, 'cookie') || '',
               },
               query: query  // ✅ ارسال page, limit, search, filters
          }) as any
          
          console.log('📦 Nuxt API Route - Backend response:', {
               dataLength: response?.data?.length,
               total: response?.total,
               page: response?.page,
               limit: response?.limit
          })

          return response
     } catch (error) {
          throw createError({
               statusCode: 500,
               statusMessage: 'خطا در دریافت محصولات'
          })
     }
})
