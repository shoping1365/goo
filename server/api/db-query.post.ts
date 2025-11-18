import { Pool } from 'pg'
import { defineEventHandler, readBody, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const pool = new Pool({
    connectionString: config.databaseUrl
  })

  try {
    const body = await readBody(event)
    const { query } = body

    if (!query) {
      throw createError({
        statusCode: 400,
        message: 'کوئری الزامی است'
      })
    }

    // فقط اجازه SELECT
    if (!query.trim().toLowerCase().startsWith('select')) {
      throw createError({
        statusCode: 404,
        message: 'صفحه مورد نظر پیدا نشد'
      })
    }

    const result = await pool.query(query)
    return { result: result.rows }
  } catch (error: any) {
    throw createError({
      statusCode: 500,
      message: error.message
    })
  } finally {
    await pool.end()
  }
}) 