import { defineEventHandler, getRouterParam, createError, readBody } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    const videoId = getRouterParam(event, 'videoId')
    const body = await readBody(event)

    if (!videoId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه ویدیو الزامی است'
      })
    }

    const response: any = await fetchGo(event, `/api/product-videos/${videoId}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        title: body.title || '',
        description: body.description || '',
        video_url: body.video_url || '',
        thumbnail_url: body.thumbnail_url || '',
        show_in_gallery: body.show_in_gallery !== undefined ? body.show_in_gallery : true,
        autoplay: body.autoplay !== undefined ? body.autoplay : false,
        show_controls: body.show_controls !== undefined ? body.show_controls : true,
        sort_order: body.sort_order || 0
      })
    })

    return {
      success: true,
      data: response
    }
  } catch (error: any) {
    console.error('خطا در ویرایش ویدیو:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در ویرایش ویدیو'
    })
  }
}) 