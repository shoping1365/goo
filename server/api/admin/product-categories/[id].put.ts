import type { H3Event } from 'h3'
import { readBody, defineEventHandler, getRouterParam, getCookie } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

declare const requireAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  // بررسی احراز هویت
  await requireAuth(event)
  const token = getCookie(event, 'access_token')

  const body = await readBody(event)
  const id = getRouterParam(event, 'id')

  return proxy(event, `${base}/api/product-categories/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token || ''}`
    },
    body: JSON.stringify(body)
  })
})

