export default defineEventHandler(async (event) => {
  try {
    const productId = getRouterParam(event, 'id')
    
    if (!productId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    // شبیه‌سازی بررسی وضعیت علاقه‌مندی
    await new Promise(resolve => setTimeout(resolve, 200))

    // تصادفی تعیین می‌کنیم که آیا در لیست علاقه‌مندی‌ها است یا نه
    const isInWishlist = Math.random() > 0.5

    return {
      isInWishlist,
      product_id: parseInt(productId)
    }

  } catch (error) {
    console.error('خطا در بررسی وضعیت علاقه‌مندی:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.statusMessage || 'خطا در بررسی وضعیت علاقه‌مندی'
    })
  }
})