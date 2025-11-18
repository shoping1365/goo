import { createError, defineEventHandler, readBody } from 'h3'
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
    const body = await readBody(event)

    const response = await fetchGo(event, '/api/menu-items', {
      method: 'POST',
      body
    }) as MenuItemResponse

    const item = response?.data

    return {
      success: true,
      data: item,
      message: 'آیتم منو با موفقیت ایجاد شد'
    }
  } catch (error: unknown) {
    console.error('Error creating menu item:', error)

    const errorWithStatus = error as { statusCode?: number; message?: string; data?: unknown }
    if (errorWithStatus?.statusCode) {
      throw createError({
        statusCode: errorWithStatus.statusCode,
        message: errorWithStatus.message || 'خطا در ایجاد آیتم منو',
        data: errorWithStatus.data || error
      })
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در ایجاد آیتم منو',
      data: error
    })
  }
})
