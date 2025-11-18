import type { H3Event } from 'h3';

declare const defineEventHandler: (handler: (event: H3Event) => unknown | Promise<unknown>) => unknown
declare const getQuery: (event: H3Event) => Record<string, string | string[] | number | undefined>
declare const requireAuth: (event: H3Event) => Promise<{ token?: string;[key: string]: unknown } | null>
declare const createError: (options: { statusCode: number; message: string }) => Error
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { headers?: Record<string, string> }) => Promise<T>

interface KnowledgeBaseResponse {
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  try {
    const q = getQuery(event)
    const user = await requireAuth(event)
    if (!user) throw createError({ statusCode: 401, message: 'Unauthorized' })

    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const url = new URL(`${base}/api/chat/admin/knowledge-base`)
    if (q.q) url.searchParams.set('q', String(q.q))
    if (q.limit) url.searchParams.set('limit', String(q.limit))
    if (q.offset) url.searchParams.set('offset', String(q.offset))

    const response = await $fetch<KnowledgeBaseResponse>(url.toString(), {
      headers: { Authorization: `Bearer ${user.token}` }
    })
    return response
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.statusMessage || 'Internal Error' })
  }
})




