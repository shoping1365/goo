import { Pool } from 'pg'
import { useRuntimeConfig } from '#imports'

// تابع برای دریافت اتصال دیتابیس
export async function getDatabase() {
  const config = useRuntimeConfig()
  
  // بررسی وجود databaseUrl
  if (!config.databaseUrl || config.databaseUrl === 'postgresql://user:password@localhost:5432/dbname') {
    throw new Error('DATABASE_URL not configured. Please set DATABASE_URL environment variable.')
  }
  
  const pool = new Pool({
    connectionString: config.databaseUrl
  })

  return {
    query: async (text, params) => {
      try {
        const result = await pool.query(text, params)
        return result.rows
      } catch (error) {
        console.error('Database query error:', error)
        throw error
      }
    },
    end: async () => {
      await pool.end()
    }
  }
} 