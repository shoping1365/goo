import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'
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

interface MenuBody {
  name?: string
  slug?: string
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')
    const body = await readBody(event) as MenuBody

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه منو الزامی است'
      })
    }

    if (!body.name) {
      throw createError({
        statusCode: 400,
        message: 'نام منو الزامی است'
      })
    }

    const response = await fetchGo(event, `/api/menus/${id}`, {
      method: 'PUT',
      body
    }) as MenuResponse

    const menu = response?.data

    return {
      success: true,
      data: menu,
      message: 'منو با موفقیت به‌روزرسانی شد'
    }
  } catch (error: unknown) {
    console.error('Error updating menu:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در به‌روزرسانی منو'
    })
  }
}) 