import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '~/server/api/_utils/fetchGo'

/**
 * POST /api/auth/verify-otp
 * تایید کد OTP و ورود
 */
export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  return await fetchGo(event, '/api/auth/verify-otp', {
    method: 'POST',
    body
  })
})
