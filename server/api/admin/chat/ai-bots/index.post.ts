import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface User {
  token?: string
  [key: string]: unknown
}

declare const requireAuth: (event: H3Event) => Promise<User | null>
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string>; body?: unknown }) => Promise<T>

export default defineEventHandler(async (event: H3Event) => {
  const user = await requireAuth(event)
  if (!user) throw createError({ statusCode: 401, message: 'Unauthorized' })
  const body = await readBody(event)
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/chat/admin/ai-bots`, {
    method: 'POST',
    headers: { Authorization: `Bearer ${user.token}`, 'Content-Type': 'application/json' },
    body
  })
})




