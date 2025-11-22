import { createError, defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

// دریافت اطلاعات یک درخواست احراز هویت
// GET /api/verifications/:id
export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  if (!id) {
    throw createError({ statusCode: 400, message: 'شناسه الزامی است' })
  }

  try {
    const response = await fetchGo(event, `/api/verifications/${id}`, {
      method: 'GET'
    })

    return response
  } catch (error: unknown) {
    console.error('Error fetching verification:', error)
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string }).message || 'خطا در دریافت اطلاعات'
    })
  }
})