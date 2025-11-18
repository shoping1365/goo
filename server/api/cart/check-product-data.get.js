import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (event) => {
  try {
    console.log('Check product data API called')
    
    const db = await getDatabase()
    console.log('Database connected')
    
    // بررسی محصولات با تمام ستون‌ها
    const products = await db.query(`
      SELECT 
        id, 
        name, 
        sku, 
        price, 
        image,
        LENGTH(image) as image_length,
        CASE 
          WHEN image IS NULL THEN 'NULL'
          WHEN image = '' THEN 'EMPTY_STRING'
          ELSE 'HAS_DATA'
        END as image_status
      FROM products 
      ORDER BY id 
      LIMIT 5
    `)
    
    console.log('Products with image analysis:', products)
    
    // بررسی جداول مرتبط با عکس
    const relatedTables = await db.query(`
      SELECT table_name 
      FROM information_schema.tables 
      WHERE table_schema = 'public' 
      AND table_name LIKE '%image%' 
      OR table_name LIKE '%media%'
      OR table_name LIKE '%file%'
    `)
    
    console.log('Related tables:', relatedTables)
    
    return {
      success: true,
      products: products,
      related_tables: relatedTables,
      message: 'Product data analysis completed'
    }
    
  } catch (error) {
    console.error('Check product data API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 