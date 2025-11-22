import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (_event) => {
  try {
    const db = await getDatabase()

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
      } catch (error) {
        results[table] = {
          exists: false,
          error: (error as { message?: string }).message
        }
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