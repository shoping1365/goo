import { createError, defineEventHandler, getRouterParam } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const productId = getRouterParam(event, 'id')

    if (!productId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    // Auth removed - previously: const { requireAuth } = await import('../../_utils/auth')
    // Auth removed - previously: const user = await requireAuth(event)

    // دریافت اتصال دیتابیس
    const { getDatabase } = await import('../../_utils/database')
    const db = await getDatabase()

    // دریافت wishlist کاربر (لیست پیش‌فرض)
    const wishlist = await db.query('SELECT id FROM user_collections WHERE is_default = true', [])

    if (!wishlist || wishlist.length === 0) {
      return {
        success: true,
        message: 'محصول در لیست علاقه‌مندی‌ها یافت نشد',
        product_id: parseInt(productId)
      }
    }

    const wishlistId = wishlist[0].id

    // حذف محصول از wishlist
    await db.query(
      'DELETE FROM user_collection_items WHERE collection_id = $1 AND product_id = $2',
      [wishlistId, productId]
    )

    return {
      success: true,
      message: 'محصول از علاقه‌مندی‌ها حذف شد',
      product_id: parseInt(productId)
    }

  } catch (error) {
    console.error('خطا در حذف از علاقه‌مندی‌ها:', error)
    const err = error as { statusCode?: number; statusMessage?: string }

    throw createError({
      statusCode: err.statusCode || 500,
      message: err.statusMessage || 'خطا در حذف از علاقه‌مندی‌ها'
    })
  }
})
