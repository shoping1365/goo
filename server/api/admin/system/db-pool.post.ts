import { defineEventHandler, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const url = `${base}/api/admin/system/db-pool`
  const body = await readBody(event)
  const data = await $fetch(url, {
    method: 'POST',
    body,
    headers: event.node.req.headers as Record<string, string>
  })
  return data
})