import { defineEventHandler } from 'h3'
import { Client } from 'pg'

// This endpoint is only for Nuxt dev/runtime internal health check
// It lives at /internal/db-health (outside of /api prefix) so it will NOT be proxied
// to the Go backend when devProxy redirects /api/* to localhost:9090.
export default defineEventHandler(async () => {
  try {
    const client = new Client({
      connectionString: process.env.DATABASE_URL ||
        'postgres://postgres:1365@localhost:5432/mydb?sslmode=disable',
      ssl: false
    })

    await client.connect()
    await client.query('SELECT 1')
    await client.end()

    // Database connection was successful
    return {
      status: 'ok',
      message: 'Database connection successful'
    }
  } catch (err: unknown) {
    // DATABASE CONNECTION FAILED
    return {
      status: 'error',
      message: (err as { message?: string }).message || 'Database connection failed'
    }
  }
}) 