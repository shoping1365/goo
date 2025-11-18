import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (event) => {
  try {
    console.log('Check all tables API called')
    
    const db = await getDatabase()
    console.log('Database connected')
    
    // بررسی تمام جداول
    const allTables = await db.query(`
      SELECT table_name 
      FROM information_schema.tables 
      WHERE table_schema = 'public' 
      ORDER BY table_name
    `)
    
    console.log('All tables:', allTables)
    
    // بررسی جداول مرتبط با محصولات و عکس‌ها
    const productRelatedTables = allTables.filter(table => 
      table.table_name.includes('product') || 
      table.table_name.includes('image') || 
      table.table_name.includes('media') || 
      table.table_name.includes('file') ||
      table.table_name.includes('attachment')
    )
    
    // بررسی ساختار جداول مرتبط
    const tableStructures = {}
    for (const table of productRelatedTables) {
      try {
        const structure = await db.query(`
          SELECT column_name, data_type, is_nullable
          FROM information_schema.columns 
          WHERE table_name = $1 
          ORDER BY ordinal_position
        `, [table.table_name])
        
        tableStructures[table.table_name] = structure
        
        // بررسی تعداد رکوردها
        const count = await db.query(`SELECT COUNT(*) as count FROM ${table.table_name}`)
        tableStructures[table.table_name + '_count'] = count[0]?.count
      } catch (error) {
        tableStructures[table.table_name + '_error'] = error.message
      }
    }
    
    return {
      success: true,
      all_tables: allTables.map(t => t.table_name),
      product_related_tables: productRelatedTables.map(t => t.table_name),
      table_structures: tableStructures,
      message: 'All tables check completed'
    }
    
  } catch (error) {
    console.error('Check all tables API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 