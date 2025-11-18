import { defineEventHandler } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const id = event.context.params?.id
  const query = getQuery(event)
  const rating = query?.rating ? `&rating=${encodeURIComponent(String(query.rating))}` : ''

  const url = `${base}/api/products/${encodeURIComponent(String(id))}/reviews?status=approved${rating}`
  const res = await $fetch(url, { method: 'GET' })
  return res
})