import { createError, defineEventHandler, readBody } from 'h3'
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
    const body = await readBody(event) as MenuBody

    // Validate required fields
    if (!body.name) {
      throw createError({
        statusCode: 400,
        message: 'نام منو الزامی است'
      })
    }

    // Generate slug from name if not provided
    if (!body.slug) {
      body.slug = String(body.name)
        .toLowerCase()
        .replace(/[^a-z0-9\u0600-\u06FF]/g, '-')
        .replace(/-+/g, '-')
        .replace(/^-|-$/g, '')
    }

    const response = await fetchGo(event, '/api/menus', {
      method: 'POST',
      body
    }) as MenuResponse

    const menu = response?.data

    return {
      success: true,
      data: menu,
      message: 'منو با موفقیت ایجاد شد'
    }
  } catch (error: unknown) {
    console.error('Error creating menu:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در ایجاد منو'
    })
  }
})