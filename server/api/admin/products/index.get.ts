import { defineEventHandler, getQuery } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase

     // دریافت query parameters
     const query = getQuery(event)

     // ساخت URL با query parameters
     const params = new URLSearchParams()
     if (query.page) params.append('page', String(query.page))
     if (query.limit) params.append('limit', String(query.limit))
     if (query.search) params.append('search', String(query.search))
     if (query.filter_category) params.append('filter_category', String(query.filter_category))
     if (query.filter_status) params.append('filter_status', String(query.filter_status))

     const queryString = params.toString()
     const fullUrl = queryString ? `${base}/api/admin/products?${queryString}` : `${base}/api/admin/products`

     const { proxy } = await import('~/server/api/_utils/fetchProxy')

     // fetchProxy خودش Authorization header را از کوکی می‌سازد
     return proxy(event, fullUrl)
})