import { createError, defineEventHandler } from 'h3'

export default defineEventHandler(async (_event) => {
  try {
    // در حال حاضر از داده‌های نمونه استفاده می‌کنیم
    // در آینده باید از دیتابیس و session کاربر استفاده شود

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