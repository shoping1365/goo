import { defineEventHandler, getQuery } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const query = getQuery(event)
  // forward any query params, e.g., ?page=2&search=foo
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as Record<string, string>).toString() : ''
  
  // استفاده از API عمومی برای محصولات منتشر شده
  return proxy(event, `${base}/api/products/public${qs}`)
}) 