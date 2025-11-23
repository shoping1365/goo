import { createError, defineEventHandler, getCookie, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    // دریافت session ID از body یا cookie
    const sessionId = body.session_id || getCookie(event, 'session_id')

    if (!sessionId) {
      throw createError({
        statusCode: 400,
        statusMessage: 'session ID یافت نشد'
      })
    }

    const { clearCart } = await import('./data')

    clearCart(sessionId)

    return {
      success: true,
      message: 'سبد خرید با موفقیت پاک شد'
    }

  } catch (error) {
    console.error('خطا در پاک کردن سبد خرید:', error)
    const err = error as { statusCode?: number; statusMessage?: string }

    throw createError({
      statusCode: err.statusCode || 500,
      statusMessage: err.statusMessage || 'خطا در پاک کردن سبد خرید'
    })
  }
})
