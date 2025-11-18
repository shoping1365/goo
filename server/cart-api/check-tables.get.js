import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (event) => {
  try {
    console.log('Check tables API called')
    
    const db = await getDatabase()
    console.log('Database connected')
    
    // بررسی وجود جداول
    const tables = ['carts', 'cart_items', 'products']
    const results = {}
    
    for (const table of tables) {
      try {
        const result = await db.query(`SELECT COUNT(*) as count FROM ${table}`)
        results[table] = {
          exists: true,
          count: result[0]?.count
        }
        console.log(`Table ${table} exists with ${result[0]?.count} records`)
      } catch (error) {
        results[table] = {
          exists: false,
          error: error.message
        }
        console.log(`Table ${table} does not exist:`, error.message)
      }
    }
    
    return {
      success: true,
      tables: results,
      message: 'Table check completed'
    }
    
  } catch (error) {
    console.error('Check tables API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 