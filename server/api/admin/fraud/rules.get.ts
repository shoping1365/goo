import type { H3Event } from 'h3';
import { defineEventHandler } from 'h3';

declare const goApiFetch: (event: H3Event, endpoint: string, options?: { method?: string; body?: unknown }) => Promise<unknown>

export default defineEventHandler(async (event) => {
  return await goApiFetch(event, '/api/admin/fraud/rules', {
    method: 'GET'
  })
})


