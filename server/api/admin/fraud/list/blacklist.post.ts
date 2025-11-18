import type { H3Event } from 'h3';
import { defineEventHandler, readBody } from 'h3';

declare const goApiFetch: (event: H3Event, endpoint: string, options?: { method?: string; body?: unknown }) => Promise<unknown>

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  return await goApiFetch(event, '/api/admin/fraud/list/blacklist', {
    method: 'POST',
    body
  })
})


