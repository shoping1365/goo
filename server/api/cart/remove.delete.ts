import { createError, defineEventHandler, getCookie, readBody } from 'h3'

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

    const { removeCartItem } = await import('./data')

    removeCartItem(sessionId, body.cart_item_id)

    return {
      success: true,
      message: 'محصول با موفقیت از سبد خرید حذف شد'
    }

  } catch (error) {
    console.error('خطا در حذف از سبد خرید:', error)
    const err = error as { statusCode?: number; statusMessage?: string }

    throw createError({
      statusCode: err.statusCode || 500,
      statusMessage: err.statusMessage || 'خطا در حذف از سبد خرید'
    })
  }
}) 