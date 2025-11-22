import { defineEventHandler, getRouterParam, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const videoId = getRouterParam(event, 'videoId')

    if (!videoId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه ویدیو الزامی است'
      })
    }

    const response = await fetchGo(event, `/api/product-videos/${videoId}`, { method: 'DELETE' })

    return {
      success: true,
      data: response
    }
  } catch (error: unknown) {
    console.error('خطا در حذف ویدیو:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در حذف ویدیو'
    })
  }
}) 