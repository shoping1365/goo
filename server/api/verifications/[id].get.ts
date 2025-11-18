import { defineEventHandler, getRouterParam, createError } from 'h3'
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
  } catch (error: any) {
    console.error('Error fetching verification:', error)
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || 'خطا در دریافت اطلاعات'
    })
  }
})