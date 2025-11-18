import { defineEventHandler, readBody, createError, getCookie } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    if (!body || !body.product_id) {
      throw createError({
        statusCode: 400,
        statusMessage: 'شناسه محصول الزامی است'
      })
    }

    // دریافت session ID از cookie
    const sessionId = getCookie(event, 'session_id')

    if (!sessionId) {
      // اگر session ID وجود ندارد، سعی می‌کنیم آن را ایجاد کنیم
      try {
        const createSessionResponse = await $fetch('/api/cart/create-session', {
          method: 'POST',
          headers: event.node.req.headers
        })
        if (createSessionResponse && createSessionResponse.session_id) {
          // session ID جدید را از کوکی دریافت کن
          const newSessionId = getCookie(event, 'session_id')
          if (newSessionId) {
            // ادامه با session ID جدید
            const { addCartItem, getCartItems } = await import('./data')
            addCartItem(newSessionId, {
              product_id: body.product_id,
              quantity: body.quantity || 1,
              name: body.name || 'محصول بدون نام',
              price: body.price || 0,
              image: body.image || '/default-product.svg',
              original_price: body.original_price,
              discount_percentage: body.discount_percentage,
              rating: body.rating,
              rating_count: body.rating_count,
              sku: body.sku
            })
            const cartItems = getCartItems(newSessionId)
            return {
              success: true,
              message: 'محصول با موفقیت به سبد خرید اضافه شد',
              cart_count: cartItems.length
            }
          }
        }
      } catch (createError) {
        // خطا در ایجاد session
      }
      
      throw createError({
        statusCode: 400,
        statusMessage: 'session ID یافت نشد'
      })
    }

    const { addCartItem, getCartItems } = await import('./data')

    addCartItem(sessionId, {
      product_id: body.product_id,
      quantity: body.quantity || 1,
      name: body.name || 'محصول بدون نام',
      price: body.price || 0,
      image: body.image || '/default-product.svg',
      original_price: body.original_price,
      discount_percentage: body.discount_percentage,
      rating: body.rating,
      rating_count: body.rating_count,
      sku: body.sku
    })

    const cartItems = getCartItems(sessionId)

    return {
      success: true,
      message: 'محصول با موفقیت به سبد خرید اضافه شد',
      cart_count: cartItems.length
    }

  } catch (error) {
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.statusMessage || 'خطا در افزودن به سبد خرید'
    })
  }
}) 