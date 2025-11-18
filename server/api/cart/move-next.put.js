import { defineEventHandler, readBody, createError, getCookie } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    if (!body.cart_item_id) {
      throw createError({
        statusCode: 400,
        statusMessage: 'شناسه آیتم سبد خرید الزامی است'
      })
    }

    // دریافت session ID از cookie
    const sessionId = getCookie(event, 'session_id')

    if (!sessionId) {
      throw createError({
        statusCode: 400,
        statusMessage: 'session ID یافت نشد'
      })
    }

    // استفاده از session-based cart
    const { moveItemToNext } = await import('./data')
    const success = moveItemToNext(sessionId, body.cart_item_id)

    if (!success) {
      throw createError({
        statusCode: 404,
        statusMessage: 'آیتم سبد خرید یافت نشد'
      })
    }

    return {
      success: true,
      message: 'آیتم به خرید بعدی انتقال یافت'
    }

  } catch (error) {
    console.error('خطا در انتقال به خرید بعدی:', error)

    // اگر خطا از بک‌اند باشد، آن را منتقل می‌کنیم
    if (error.statusCode) {
      throw createError({
        statusCode: error.statusCode,
        statusMessage: error.statusMessage || 'خطا در انتقال به خرید بعدی'
      })
    }

    throw createError({
      statusCode: 500,
      statusMessage: 'خطا در انتقال به خرید بعدی'
    })
  }
})
