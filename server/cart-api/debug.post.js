export default defineEventHandler(async (_event) => {
  try {
    const body = await readBody(_event)

    // تست اتصال به دیتابیس
    try {
      const { getDatabase } = await import('../_utils/database.js')
      const db = await getDatabase()

      // تست query ساده
      const testResult = await db.query('SELECT 1 as test')

      return {
        success: true,
        message: 'Database connection OK',
        test_result: testResult
      }
    } catch (dbError) {
      console.error('Database error:', dbError)
      return {
        success: false,
        error: 'Database connection failed',
        details: dbError.message
      }
    }

  } catch (error) {
    console.error('Debug API error:', error)
    return {
      success: false,
      error: error.message,
      stack: error.stack
    }
  }
}) 