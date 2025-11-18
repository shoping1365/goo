import type { H3Event } from 'h3';

declare const defineEventHandler: (handler: (event: H3Event) => unknown | Promise<unknown>) => unknown
declare const requireAuth: (event: H3Event) => Promise<{ token?: string;[key: string]: unknown } | null>
declare const createError: (options: { statusCode: number; message: string }) => Error
declare const readBody: (event: H3Event) => Promise<unknown>
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string>; body?: unknown }) => Promise<T>

export default defineEventHandler(async (event: H3Event) => {
  const user = await requireAuth(event)
  if (!user) throw createError({ statusCode: 401, message: 'Unauthorized' })
  const body = await readBody(event)
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/chat/admin/operators/invite`, {
    method: 'POST',
    headers: { Authorization: `Bearer ${user.token}`, 'Content-Type': 'application/json' },
    body
  })
})




