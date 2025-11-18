import { defineEventHandler } from 'h3'
import { getRouterParam } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const id = getRouterParam(event, 'id')
  const url = `${base}/api/admin/reviews/${encodeURIComponent(String(id))}/approve`

  return await proxy(event, url, {
    method: 'POST',
    headers: {
      'X-User-ID': String(event.context?.user?.id || 1),
      'X-User-Role': event.context?.user?.role || 'admin',
    },
  })
})

