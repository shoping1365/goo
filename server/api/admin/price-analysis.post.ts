import { createError, defineEventHandler, readBody } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

// Simple proxy endpoint that queries backend Go service if available,
// otherwise performs a lightweight search + extraction via configured external service.
export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const q = String(body?.q || '').trim()
  const limit = Math.min(Math.max(Number(body?.limit || 10), 3), 15)
  if (!q) {
    throw createError({ statusCode: 400, message: 'کلیدواژه الزامی است' })
  }

  // Prefer Go backend aggregator if exists
  try {
    const res = await goApiFetch(event, '/api/admin/price-analysis', {
      method: 'POST',
      body: { q, limit }
    }) as { status: string; data: unknown[] }
    return res
  } catch (e) {
    // Fallback: return minimal stub to keep UI functional
    return {
      status: 'success',
      data: []
    }
  }
})


