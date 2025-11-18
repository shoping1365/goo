import { createError, defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه منو الزامی است'
      })
    }

    await fetchGo(event, `/api/menus/${id}`, {
      method: 'DELETE'
    })

    return {
      success: true,
      message: 'منو با موفقیت حذف شد'
    }
  } catch (error: unknown) {
    console.error('Error deleting menu:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در حذف منو'
    })
  }
}) 