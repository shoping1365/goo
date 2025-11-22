import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    if (!body.product_id || !body.image_url) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول و مسیر عکس الزامی است'
      })
    }

    const db = await getDatabase()

    // اضافه کردن عکس به محصول
    const result = await db.query(`
      INSERT INTO product_images (product_id, image_url, type, sort_order, is_active, created_at)
      VALUES ($1, $2, $3, $4, $5, NOW())
      RETURNING id
    `, [
      body.product_id,
      body.image_url,
      body.type || 'gallery',
      body.sort_order || 0,
      body.is_active !== false
    ])


    return {
      success: true,
      image_id: result[0]?.id,
      message: 'عکس محصول با موفقیت اضافه شد'
    }

  } catch (error) {
    console.error('Error adding product image:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در اضافه کردن عکس محصول'
    })
  }
})