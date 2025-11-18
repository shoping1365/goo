import { defineEventHandler, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const body = await readBody(event)
  const url = `${base}/api/admin/reviews/settings`
  return await proxy(event, url, {
    method: 'PUT',
    headers: {
      'X-User-ID': String(event.context?.user?.id || 1),
      'X-User-Role': event.context?.user?.role || 'admin',
    },
    body
  })
})

