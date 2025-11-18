export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    
    if (!body.product_id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    const userId = event.context?.user?.id
    
    if (!userId) {
      throw createError({
        statusCode: 401,
        message: 'کاربر احراز هویت نشده است'
      })
    }

    // دریافت اتصال دیتابیس
    const { getDatabase } = await import('../_utils/database')
    const db = await getDatabase()

    // بررسی وجود محصول
    const product = await db.query('SELECT id FROM products WHERE id = $1', [body.product_id])
    if (!product || product.length === 0) {
      throw createError({
        statusCode: 404,
        message: 'محصول یافت نشد'
      })
    }

    // دریافت یا ایجاد wishlist برای کاربر (لیست پیش‌فرض)
    let wishlist = await db.query('SELECT id FROM user_collections WHERE user_id = $1 AND is_default = true', [userId])
    
    if (!wishlist || wishlist.length === 0) {
      // ایجاد wishlist جدید
      const newWishlist = await db.query(
        'INSERT INTO user_collections (user_id, name, is_default) VALUES ($1, $2, true) RETURNING id',
        [userId, 'علاقمندی‌ها']
      )
      wishlist = newWishlist
    }

    const wishlistId = wishlist[0].id

    // بررسی اینکه آیا محصول قبلاً اضافه شده یا نه
    const existingItem = await db.query(
      'SELECT id FROM user_collection_items WHERE collection_id = $1 AND product_id = $2',
      [wishlistId, body.product_id]
    )

    if (existingItem && existingItem.length > 0) {
      return {
        success: true,
        message: 'محصول قبلاً در لیست علاقه‌مندی‌ها موجود است',
        product_id: parseInt(body.product_id)
      }
    }

    // اضافه کردن محصول به wishlist
    await db.query(
      'INSERT INTO user_collection_items (collection_id, product_id) VALUES ($1, $2)',
      [wishlistId, body.product_id]
    )

    console.log(`محصول ${body.product_id} به علاقه‌مندی‌های کاربر ${userId} اضافه شد`)

    return {
      success: true,
      message: 'محصول به علاقه‌مندی‌ها اضافه شد',
      product_id: parseInt(body.product_id)
    }

  } catch (error) {
    console.error('خطا در افزودن به علاقه‌مندی‌ها:', error)
    
    throw createError({
      statusCode: error?.statusCode || 500,
      message: error?.statusMessage || 'خطا در افزودن به علاقه‌مندی‌ها'
    })
  }
})