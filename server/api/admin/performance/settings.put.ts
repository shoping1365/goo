import { defineEventHandler, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const body = await readBody(event)
  return await $fetch(`${base}/api/admin/performance/settings`, {
    method: 'PUT',
    body,
    headers: event.node.req.headers as Record<string, string>,
  })
})