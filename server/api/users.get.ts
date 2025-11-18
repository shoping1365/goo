import { defineEventHandler, getRequestHeader } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const cookie = getRequestHeader(event, 'cookie') || ''

  // پراکسی به بک‌اند Go با انتقال کوکی‌ها برای احراز هویت
  return await $fetch(`${base}/api/users`, {
    method: 'GET',
    credentials: 'include',
    headers: { cookie }
  })
})