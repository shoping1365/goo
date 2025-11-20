import { createError, defineEventHandler, getHeader, readBody } from 'h3'
import { useRuntimeConfig } from 'nuxt/app'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     const body = await readBody(event)

     try {
          const response = await $fetch(`${base}/api/admin/admin-settings/bulk-edit-columns`, {
               method: 'PUT',
               body: body,
               headers: {
                    cookie: getHeader(event, 'cookie') || '',
                    'Content-Type': 'application/json'
               }
          })

          return response
     } catch {
          throw createError({
               statusCode: 500,
               statusMessage: 'خطا در ذخیره تنظیمات ستون‌ها'
          })
     }
})
