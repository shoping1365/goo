import { createError, defineEventHandler, getRouterParam } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface Menu {
  id: number
  name: string
  slug: string
  [key: string]: unknown
}

interface MenuResponse {
  data?: Menu
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه منو الزامی است'
      })
    }

    const response = await fetchGo(event, `/api/menus/${id}`) as MenuResponse
    const menu = response?.data

    return {
      success: true,
      data: menu,
      message: 'منو با موفقیت دریافت شد'
    }
  } catch (error: unknown) {
    console.error('Error fetching menu:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 404,
      message: errorWithStatus?.message || 'منو یافت نشد'
    })
  }
})