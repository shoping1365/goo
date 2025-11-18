import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MenuItem {
  id: number
  name: string
  url: string
  [key: string]: unknown
}

interface MenuItemResponse {
  data?: MenuItem
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')
    const body = await readBody(event)

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه آیتم منو الزامی است'
      })
    }

    const response = await fetchGo(event, `/api/menu-items/${id}`, {
      method: 'PUT',
      body
    }) as MenuItemResponse

    const item = response?.data

    return {
      success: true,
      data: item,
      message: 'آیتم منو با موفقیت به‌روزرسانی شد'
    }
  } catch (error: unknown) {
    console.error('Error updating menu item:', error)

    const errorWithStatus = error as { statusCode?: number; message?: string; data?: unknown }
    if (errorWithStatus?.statusCode) {
      throw createError({
        statusCode: errorWithStatus.statusCode,
        message: errorWithStatus.message || 'خطا در به‌روزرسانی آیتم منو',
        data: errorWithStatus.data || error
      })
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در به‌روزرسانی آیتم منو',
      data: error
    })
  }
})
