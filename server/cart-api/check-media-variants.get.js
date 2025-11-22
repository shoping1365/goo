import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (_event) => {
  try {
    const db = await getDatabase()

    // بررسی ساختار جدول media_variants
    const tableStructure = await db.query(`
      SELECT column_name, data_type, is_nullable
      FROM information_schema.columns 
      WHERE table_name = 'media_variants' 
      ORDER BY ordinal_position
    `)

    // بررسی داده‌های جدول media_variants
    const mediaVariants = await db.query(`
      SELECT * FROM media_variants 
      ORDER BY id 
      LIMIT 10
    `)

    // بررسی تعداد رکوردها
    const count = await db.query(`SELECT COUNT(*) as count FROM media_variants`)

    // بررسی variant های 150x150
    const variants150x150 = await db.query(`
      SELECT * FROM media_variants 
      WHERE width = 150 AND height = 150
      LIMIT 5
    `)

    return {
      success: true,
      table_structure: tableStructure,
      sample_data: mediaVariants,
      total_count: count[0]?.count,
      variants_150x150: variants150x150,
      message: 'Media variants check completed'
    }

  } catch (error) {
    console.error('Check media variants API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 