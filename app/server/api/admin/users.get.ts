import { proxy } from '../../../../server/api/_utils/fetchProxy'
import { defineEventHandler, getCookie, createError, getQuery } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const token = getCookie(event, 'access_token') || getCookie(event, 'auth-token')
  
  if (!token) {
    throw createError({ statusCode: 401, message: 'برای ادامه باید وارد شوید.' })
  }

  const query = getQuery(event)
  const page = query.page || 1
  const limit = query.limit || 20
  
  return await proxy(event, `${config.public.goApiBase}/api/admin/users?page=${page}&limit=${limit}`, {
    headers: { Authorization: `Bearer ${token}` }
  })
})
