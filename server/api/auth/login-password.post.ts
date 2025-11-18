import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '~/server/api/_utils/fetchGo'

/**
 * POST /api/auth/login-password
 * ورود با نام کاربری و رمز عبور
 */
export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  return await fetchGo(event, '/api/auth/login-password', {
    method: 'POST',
    body
  })
})
