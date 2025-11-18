import { defineEventHandler, getRouterParam, readBody, createError } from 'h3'
import { fetchGo } from '../../../_utils/fetchGo'

// رد درخواست احراز هویت (ادمین)
// POST /api/admin/verifications/:id/reject
export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  if (!id) {
    throw createError({ statusCode: 400, message: 'شناسه الزامی است' })
  }

  try {
    const body = await readBody(event)

    if (!body.reason) {
      throw createError({ statusCode: 400, message: 'دلیل رد الزامی است' })
    }

    const response = await fetchGo(event, `/api/admin/verifications/${id}/reject`, {
      method: 'POST',
      body: body
    })

    return response
  } catch (error: any) {
    console.error('Error rejecting request:', error)
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || 'خطا در رد درخواست'
    })
  }
})

