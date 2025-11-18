import { defineEventHandler } from 'h3'
import { getRouterParam, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const id = getRouterParam(event, 'id')
  const payload = await readBody<{ reason?: string }>(event)
  const url = `${base}/api/admin/reviews/${encodeURIComponent(String(id))}/reject`

  return await proxy(event, url, {
    method: 'POST',
    headers: {
  'Content-Type': 'application/json',
      'X-User-ID': String(event.context?.user?.id || 1),
      'X-User-Role': event.context?.user?.role || 'admin',
    },
    body: payload ? JSON.stringify(payload) : undefined
  })
})

