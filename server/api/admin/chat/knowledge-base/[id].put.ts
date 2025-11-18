import type { H3Event } from 'h3'
import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'
import { getCookieValue } from '../../../_utils/cookieHelper'

declare const requireAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string>; body?: unknown }) => Promise<T>

export default defineEventHandler(async (event: H3Event) => {
  try {
    const id = getRouterParam(event, 'id')
    await requireAuth(event)

    const body = await readBody(event)
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const token = getCookieValue(event, 'access_token')

    return await $fetch(`${base}/api/chat/admin/knowledge-base/${encodeURIComponent(id!)}`, {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body
    })
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.statusMessage || 'Internal Error' })
  }
})
