import { defineEventHandler, getRouterParam, createError, getCookie } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  
  const token = getCookie(event, 'access_token')

  const id = getRouterParam(event, 'id')

  return proxy(event, `${base}/api/products/${id}`, {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
})
