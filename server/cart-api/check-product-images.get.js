import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (_event) => {
  try {
    const db = await getDatabase()

    // بررسی عکس‌های محصولات
    const products = await db.query(`
      SELECT id, name, sku, price, image 
      FROM products 
      ORDER BY id 
      LIMIT 10
    `)

    return {
      success: true,
      products: products,
      message: 'Product images check completed'
    }

  } catch (error) {
    console.error('Check product images API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 