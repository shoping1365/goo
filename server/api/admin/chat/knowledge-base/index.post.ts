import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface User {
  token?: string
  [key: string]: unknown
}

interface KnowledgeBaseResponse {
  [key: string]: unknown
}

declare const requireAuth: (event: H3Event) => Promise<User | null>
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string>; body?: unknown }) => Promise<T>

export default defineEventHandler(async (event: H3Event) => {
  try {
    const user = await requireAuth(event)
    if (!user) throw createError({ statusCode: 401, message: 'Unauthorized' })
    const body = await readBody(event)

    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const response = await $fetch<KnowledgeBaseResponse>(`${base}/api/chat/admin/knowledge-base`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${user.token}`,
        'Content-Type': 'application/json'
      },
      body
    })
    return response
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.statusMessage || 'Internal Error' })
  }
})




