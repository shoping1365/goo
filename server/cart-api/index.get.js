import { defineEventHandler, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // استفاده از همان منطق موجود در get.get.js
    const { getCartItems } = await import('./data')
    const cartItems = getCartItems()

    return {
      success: true,
      items: cartItems,
      total_items: cartItems.length,
      total_price: cartItems.reduce((sum, item) => sum + (item.price * item.quantity), 0)
    }

  } catch (error) {
    console.error('خطا در دریافت سبد خرید:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.statusMessage || 'خطا در دریافت سبد خرید'
    })
  }
})