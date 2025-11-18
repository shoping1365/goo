import { defineEventHandler, getRouterParam, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const productId = getRouterParam(event, 'productId')

    if (!productId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    // فراخوانی بک‌اند Go (کوکی/هدر Auth به‌صورت خودکار فوروارد می‌شود)
    const response: any = await fetchGo(event, `/api/product-videos/${productId}`, { method: 'GET' })

    return {
      success: true,
      data: Array.isArray(response?.data) ? response.data : (Array.isArray(response) ? response : [])
    }
  } catch (error: any) {
    console.error('خطا در دریافت ویدیوهای محصول:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در دریافت ویدیوهای محصول'
    })
  }
}) 