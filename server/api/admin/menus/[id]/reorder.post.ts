import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'
import { fetchGo } from '../../../_utils/fetchGo'

interface ReorderBody {
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')
    const body = await readBody(event) as ReorderBody

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه منو الزامی است'
      })
    }

    if (!body || !Array.isArray(body)) {
      throw createError({
        statusCode: 400,
        message: 'داده‌های ترتیب نامعتبر است'
      })
    }

    await fetchGo(event, `/api/menus/${id}/reorder`, {
      method: 'POST',
      body
    })

    return {
      success: true,
      message: 'ترتیب آیتم‌های منو با موفقیت تغییر یافت'
    }
  } catch (error: unknown) {
    console.error('Error reordering menu items:', error)
    const errorWithStatus = error as { statusCode?: number; message?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در تغییر ترتیب آیتم‌های منو'
    })
  }
})