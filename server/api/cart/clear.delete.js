import { defineEventHandler, readBody, createError, getCookie } from 'h3'

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

    console.log(`سبد خرید session ${sessionId} پاک شد`)

    return {
      success: true,
      message: 'سبد خرید با موفقیت پاک شد'
    }

  } catch (error) {
    console.error('خطا در پاک کردن سبد خرید:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.statusMessage || 'خطا در پاک کردن سبد خرید'
    })
  }
})
