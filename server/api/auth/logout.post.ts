import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '~/server/api/_utils/fetchGo'

/**
 * POST /api/auth/logout
 * خروج از سیستم
 */
export default defineEventHandler(async (event) => {
  const body = await readBody(event).catch(() => ({})) // logout معمولاً body ندارد
  return await fetchGo(event, '/api/auth/logout', {
    method: 'POST',
    body
  })
})
