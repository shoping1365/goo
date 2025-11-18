import type { H3Event } from 'h3'
import { createError, defineEventHandler, getRouterParam } from 'h3'
import { getCookieValue } from '../../../_utils/cookieHelper'

declare const requireAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { headers?: Record<string, string> }) => Promise<T>

export default defineEventHandler(async (event: H3Event) => {
  try {
    const id = getRouterParam(event, 'id')
    await requireAuth(event)

    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const token = getCookieValue(event, 'access_token')

    return await $fetch(`${base}/api/chat/admin/knowledge-base/${encodeURIComponent(id!)}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.statusMessage || 'Internal Error' })
  }
})
