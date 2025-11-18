import type { H3Event } from 'h3'
import { defineEventHandler, getRouterParam, readBody } from 'h3'

declare const requireAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>
declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown }) => Promise<T>

export default defineEventHandler(async (event) => {
  await requireAuth(event)

  const config = useRuntimeConfig()
  const id = getRouterParam(event, 'id')
  const body = await readBody(event)

  return await $fetch(`${config.public.goApiBase}/api/admin/fraud/rules/${id}`, {
    method: 'PUT',
    body
  })
})
