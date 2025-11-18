import { useRuntimeConfig } from '#imports'
import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/admin/system/reload`, { method: 'POST', headers: event.node.req.headers as Record<string, string> })
})