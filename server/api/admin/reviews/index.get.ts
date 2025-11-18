import { useRuntimeConfig } from '#imports'
import { defineEventHandler, getCookie, getQuery, getRequestHeader } from 'h3'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase


  const user = event.context.user || { id: 1, role: 'admin' }
  const cookie = getRequestHeader(event, 'cookie') || ''
  const token = getCookie(event, 'access_token') || getCookie(event, 'auth-token')

  const query = getQuery(event)
  const params = new URLSearchParams()
  if (query?.status) params.append('status', String(query.status))
  if (query?.rating) params.append('rating', String(query.rating))
  if (query?.search) params.append('search', String(query.search))
  if (query?.product_id) params.append('product_id', String(query.product_id))
  if (query?.page) params.append('page', String(query.page))
  if (query?.per_page) params.append('per_page', String(query.per_page))

  const url = `${base}/api/admin/reviews?${params.toString()}`
  return await proxy(event, url, {
    method: 'GET',
    headers: {
      cookie,
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      'X-User-ID': String(user?.id || 1),
      'X-User-Role': user?.role || 'admin'
    }
  })
})
