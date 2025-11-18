import { defineEventHandler, readBody, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // بررسی method و headers
    console.log('Request Method:', event.node.req.method)
    console.log('Content-Type:', event.node.req.headers['content-type'])
    
    const body = await readBody(event)
    
    // دیباگ: نمایش محتوای درخواست
    console.log('Cart Add Request Body:', body)
    console.log('Body type:', typeof body)
    console.log('product_id:', body?.product_id, 'Type:', typeof body?.product_id)
    
    if (!body || !body.product_id) {
      console.error('Missing product_id in body:', body)
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    const { addCartItem, getCartItems } = await import('./data')

    addCartItem({
      product_id: body.product_id,
      quantity: body.quantity || 1,
      name: body.name || 'محصول بدون نام',
      price: body.price || 0,
      image: body.image || '/statics/images/sample1.jpg'
    })

    const cartItems = getCartItems()

    return {
      success: true,
      message: 'محصول با موفقیت به سبد خرید اضافه شد',
      cart_count: cartItems.length
    }

  } catch (error) {
    console.error('خطا در افزودن به سبد خرید:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.statusMessage || 'خطا در افزودن به سبد خرید'
    })
  }
})