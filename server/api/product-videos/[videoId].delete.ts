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

    const response: any = await fetchGo(event, `/api/product-videos/${videoId}`, { method: 'DELETE' })

    return {
      success: true,
      data: response
    }
  } catch (error: any) {
    console.error('خطا در حذف ویدیو:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در حذف ویدیو'
    })
  }
}) 