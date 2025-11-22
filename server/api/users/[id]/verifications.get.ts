import { createError, defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

// دریافت لیست درخواست‌های احراز هویت یک کاربر
// GET /api/users/:id/verifications
export default defineEventHandler(async (event) => {
  const userId = getRouterParam(event, 'id')
  if (!userId) {
    throw createError({ statusCode: 400, message: 'شناسه کاربر الزامی است' })
  }

  try {
    const response = await fetchGo(event, `/api/users/${userId}/verifications`, {
      method: 'GET'
    })

    return response
  } catch (error: unknown) {
    console.error('Error fetching user verifications:', error)
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string }).message || 'خطا در دریافت لیست احراز هویت'
    })
  }
})

