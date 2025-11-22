import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const productId = getRouterParam(event, 'productId')
    const body = await readBody(event)

    if (!productId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    if (!body.video_url) {
      throw createError({
        statusCode: 400,
        message: 'لینک ویدیو الزامی است'
      })
    }

    // فراخوانی بک‌اند Go با فوروارد هدرهای لازم
    const response = await fetchGo(event, `/api/product-videos/${productId}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        video_url: body.video_url,
        thumbnail_url: body.thumbnail_url || '',
        title: body.title || '',
        description: body.description || '',
        show_in_gallery: body.show_in_gallery !== undefined ? body.show_in_gallery : true,
        autoplay: body.autoplay !== undefined ? body.autoplay : false,
        show_controls: body.show_controls !== undefined ? body.show_controls : true,
        sort_order: body.sort_order || 0,
        lazy_load: body.lazy_load !== undefined ? body.lazy_load : true
      })
    })

    return {
      success: true,
      data: response
    }
  } catch (error: unknown) {
    console.error('خطا در افزودن ویدیو:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در افزودن ویدیو'
    })
  }
}) 