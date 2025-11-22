import { defineEventHandler, getRouterParam, createError } from 'h3'
import { fetchGo } from '../../../_utils/fetchGo'

// تایید درخواست احراز هویت (ادمین)
// POST /api/admin/verifications/:id/verify
export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  if (!id) {
    throw createError({ statusCode: 400, message: 'شناسه الزامی است' })
  }

  try {
    const response = await fetchGo(event, `/api/admin/verifications/${id}/verify`, {
      method: 'POST',
      body: {}
    })

    return response
  } catch (error: unknown) {
    console.error('Error verifying request:', error)
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string }).message || 'خطا در تایید درخواست'
    })
  }
})

