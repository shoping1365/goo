import { createError, defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه آیتم منو الزامی است'
      })
    }

    await fetchGo(event, `/api/menu-items/${id}`, {
      method: 'DELETE'
    })

    return {
      success: true,
      message: 'آیتم منو با موفقیت حذف شد'
    }
  } catch (error: unknown) {
    console.error('Error deleting menu item:', error)

    const errorWithStatus = error as { statusCode?: number; message?: string; data?: unknown }
    if (errorWithStatus?.statusCode) {
      throw createError({
        statusCode: errorWithStatus.statusCode,
        message: errorWithStatus.message || 'خطا در حذف آیتم منو',
        data: errorWithStatus.data || error
      })
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در حذف آیتم منو',
      data: error
    })
  }
})
