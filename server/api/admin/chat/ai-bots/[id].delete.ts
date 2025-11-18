import type { H3Event } from 'h3'
import { defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../../_utils/fetchGo'

declare const requireAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>

export default defineEventHandler(async (event): Promise<unknown> => {
  await requireAuth(event)
  const id = getRouterParam(event, 'id')
  return await fetchGo(event, `/api/chat/admin/ai-bots/${encodeURIComponent(String(id))}`, {
    method: 'DELETE'
  })
})
