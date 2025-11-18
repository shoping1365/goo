import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '~/server/api/_utils/fetchGo'

/**
 * POST /api/auth/send-otp
 * ارسال کد OTP به شماره موبایل
 */
export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  return await fetchGo(event, '/api/auth/send-otp', {
    method: 'POST',
    body
  })
})
