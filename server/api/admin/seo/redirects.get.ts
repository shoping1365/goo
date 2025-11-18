import { defineEventHandler, getQuery, getCookie, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
          const config = useRuntimeConfig()
          const apiBaseUrl = config.public.goApiBase
          try {
               // دریافت پارامترهای query
               const query = getQuery(event)
               const group = query.group as string || ''
               const q = query.q as string || ''
               const limit = parseInt(query.limit as string) || 100000
               const offset = parseInt(query.offset as string) || 0

               // ایجاد URL با پارامترها
               const params = new URLSearchParams()
               if (group) params.append('group', group)
               if (q) params.append('q', q)
               if (limit) params.append('limit', limit.toString())
               if (offset) params.append('offset', offset.toString())

               const url = `${apiBaseUrl}/api/admin/seo/redirects?${params.toString()}`

               // درخواست به Go backend
               const response = await $fetch(url, {
                    method: 'GET',
                    headers: {
                         'Authorization': `Bearer ${getCookie(event, 'auth-token')}`,
                         'Content-Type': 'application/json'
                    }
               })

               return response
          } catch (error: unknown) {
               const errorObj = error as { message?: string }
               throw createError({
                    statusCode: 500,
                    message: 'خطا در دریافت لیست ریدایرکت‌ها: ' + (errorObj?.message || 'Unknown error')
               })
          }
})