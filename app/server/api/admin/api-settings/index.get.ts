import { createError, defineEventHandler, getHeader } from 'h3'
import { useRuntimeConfig } from 'nuxt/app'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  try {
    const response = await $fetch(`${base}/api/admin/api-settings`, {
      headers: {
        cookie: getHeader(event, 'cookie') || ''
      }
    })

    return response
  } catch {
    throw createError({
      statusCode: 500,
      statusMessage: 'خطا در دریافت تنظیمات API'
    })
  }
})
