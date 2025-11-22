import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (_event) => {
  try {
    const db = await getDatabase()

    // بررسی ساختار جدول product_images
    const tableStructure = await db.query(`
      SELECT column_name, data_type, is_nullable
      FROM information_schema.columns 
      WHERE table_name = 'product_images' 
      ORDER BY ordinal_position
    `)

    // بررسی داده‌های جدول product_images
    const productImages = await db.query(`
      SELECT * FROM product_images 
      ORDER BY id 
      LIMIT 10
    `)

    // بررسی تعداد رکوردها
    const count = await db.query(`SELECT COUNT(*) as count FROM product_images`)

    // بررسی عکس‌های یک محصول خاص (product_id = 4 که در سبد خرید است)
    const product4Images = await db.query(`
      SELECT * FROM product_images 
      WHERE product_id = 4
    `)

    return {
      success: true,
      table_structure: tableStructure,
      sample_data: productImages,
      total_count: count[0]?.count,
      product_4_images: product4Images,
      message: 'Product images table check completed'
    }

  } catch (error) {
    console.error('Check product_images table API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 