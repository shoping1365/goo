import { defineEventHandler, getQuery, createError } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

// دریافت لیست تمام درخواست‌های احراز هویت (برای ادمین)
// GET /api/admin/verifications?status=pending&page=1&page_size=20
export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    const queryString = new URLSearchParams(query as Record<string, string>).toString()

    const response = await fetchGo(event, `/api/admin/verifications?${queryString}`, {
      method: 'GET'
    })

    return response
  } catch (error: any) {
    console.error('Error fetching verifications:', error)
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || 'خطا در دریافت لیست'
    })
  }
})

