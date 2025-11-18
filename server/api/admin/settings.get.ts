import { proxy } from '../_utils/fetchProxy'
import { defineEventHandler, createError } from 'h3'
import { getCookieValue } from '../_utils/cookieHelper'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const token = getCookieValue(event, 'access_token') || getCookieValue(event, 'auth-token')
  if (!token) {
    throw createError({ statusCode: 401, message: 'برای ادامه باید وارد شوید.' })
  }
  return await proxy(event, `${config.public.goApiBase}/api/admin/settings`, {
    headers: { Authorization: `Bearer ${token}` }
  })
}) 