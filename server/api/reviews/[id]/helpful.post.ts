import { defineEventHandler } from '#imports'
import { getRequestHeader } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  // احراز هویت سمت سرور نیترو (نیاز به لاگین)
  const user = await requireAuth(event)
  const authHeader = user?.token ? { Authorization: `Bearer ${user.token}` } : {}
  const cookie = getRequestHeader(event, 'cookie') || ''

  const id = event.context.params?.id
  const url = `${base}/api/reviews/${encodeURIComponent(String(id))}/helpful`
  const res = await $fetch(url, {
    method: 'POST',
    headers: {
      cookie,
      ...authHeader,
      'X-User-ID': String(user?.id || 1),
      'X-User-Role': user?.role || 'customer'
    }
  })
  return res
})