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
    const response = await fetchGo(event, `/api/product-videos/${productId}`, { method: 'GET' }) as { data?: unknown[] }

    return {
      success: true,
      data: Array.isArray(response?.data) ? response.data : (Array.isArray(response) ? response : [])
    }
  } catch (error: unknown) {
    console.error('خطا در دریافت ویدیوهای محصول:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در دریافت ویدیوهای محصول'
    })
  }
}) 