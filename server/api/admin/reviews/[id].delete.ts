import type { H3Event } from 'h3'
import { defineEventHandler, getRouterParam } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

declare const requireAuth: (event: H3Event) => Promise<{ id: number | string; role: string } | null>

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const user = await requireAuth(event)
  if (!user) {
    throw new Error('Unauthorized')
  }
  
  const id = getRouterParam(event, 'id')
  const url = `${base}/api/admin/reviews/${encodeURIComponent(String(id))}`

  return await proxy(event, url, {
    method: 'DELETE',
    headers: {
      'X-User-ID': String(user.id),
      'X-User-Role': user.role,
    }
  })
})
