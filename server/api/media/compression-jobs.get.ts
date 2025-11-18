import { defineEventHandler, getQuery, getRequestHeader } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const cookie = getRequestHeader(event, 'cookie') || ''
  const query = getQuery(event)

  // ساخت query string
  const queryString = Object.keys(query).length
    ? '?' + Object.entries(query).map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(v as string)}`).join('&')
    : ''

  // پراکسی به بک‌اند Go با انتقال کوکی‌ها برای احراز هویت
  return await $fetch(`${base}/api/media/compression-jobs${queryString}`, {
    method: 'GET',
    credentials: 'include',
    headers: { cookie }
  })
})