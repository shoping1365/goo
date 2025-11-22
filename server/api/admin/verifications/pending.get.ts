import { defineEventHandler, getQuery, createError } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

// دریافت لیست درخواست‌های در انتظار تایید
// GET /api/admin/verifications/pending?page=1&page_size=20
export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    const queryString = new URLSearchParams(query as Record<string, string>).toString()

    const response = await fetchGo(event, `/api/admin/verifications/pending?${queryString}`, {
      method: 'GET'
    })

    return response
  } catch (error: unknown) {
    console.error('Error fetching pending verifications:', error)
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string }).message || 'خطا در دریافت لیست'
    })
  }
})

