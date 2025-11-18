import { defineEventHandler, createError, getCookie } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // دریافت session ID از cookie
    const sessionId = getCookie(event, 'session_id')

    if (!sessionId) {
      // اگر session ID وجود ندارد، سبد خرید خالی است
      return {
        success: true,
        items: [],
        total_items: 0,
        total_price: 0
      }
    }

    const { getCartItems } = await import('./data')
    const cartItems = getCartItems(sessionId)

    return {
      success: true,
      items: cartItems.map(item => ({
        ...item,
        price: item.price || 0,
        name: item.name || 'محصول بدون نام',
        image: item.image || '/default-product.svg'
      })),
      total_items: cartItems.length,
      total_price: cartItems.reduce((sum, item) => sum + ((item.price || 0) * (item.quantity || 1)), 0)
    }

  } catch (error) {
    console.error('خطا در دریافت سبد خرید:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.statusMessage || 'خطا در دریافت سبد خرید'
    })
  }
}) 