import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (_event) => {
  try {
    const db = await getDatabase()

    // بررسی ساختار جدول products
    const tableInfo = await db.query(`
      SELECT column_name, data_type, is_nullable, column_default
      FROM information_schema.columns 
      WHERE table_name = 'products' 
      ORDER BY ordinal_position
    `)

    // بررسی یک محصول با تمام ستون‌ها
    const productWithAllColumns = await db.query(`
      SELECT * FROM products WHERE id = 1
    `)

    return {
      success: true,
      table_structure: tableInfo,
      sample_product: productWithAllColumns[0] || null,
      message: 'Product table structure check completed'
    }

  } catch (error) {
    console.error('Check product structure API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 