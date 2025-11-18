import { defineEventHandler, readBody, createError, getCookie } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    if (!body.cart_item_id || !body.quantity) {
      throw createError({
        statusCode: 400,
        statusMessage: 'شناسه آیتم سبد خرید و تعداد الزامی است'
      })
    }

    if (body.quantity <= 0) {
      throw createError({
        statusCode: 400,
        statusMessage: 'تعداد باید بیشتر از صفر باشد'
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

    const { updateCartItem } = await import('./data')

    updateCartItem(sessionId, body.cart_item_id, body.quantity)

    console.log(`تعداد آیتم ${body.cart_item_id} به ${body.quantity} تغییر یافت`)

    return {
      success: true,
      message: 'سبد خرید با موفقیت به‌روزرسانی شد'
    }

  } catch (error) {
    console.error('خطا در به‌روزرسانی سبد خرید:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.statusMessage || 'خطا در به‌روزرسانی سبد خرید'
    })
  }
}) 