import { defineEventHandler } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const url = `${base}/api/admin/performance/settings`
  const data = await $fetch(url, { headers: event.node.req.headers as Record<string, string> })
  return data
})