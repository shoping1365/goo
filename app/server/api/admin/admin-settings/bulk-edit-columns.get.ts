import { createError, defineEventHandler, getHeader } from 'h3'
import { useRuntimeConfig } from 'nuxt/app'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase

     try {
          const response = await $fetch(`${base}/api/admin/admin-settings/bulk-edit-columns`, {
               method: 'GET',
               headers: {
                    cookie: getHeader(event, 'cookie') || '',
               }
          })

          return response
     } catch (error) {
          throw createError({
               statusCode: 500,
               statusMessage: 'خطا در دریافت تنظیمات ستون‌ها'
          })
     }
})
