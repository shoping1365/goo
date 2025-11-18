export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { getQuery } = await import('h3')
  const { proxy } = await import('../_utils/fetchProxy')

  const query = getQuery(event)
  // forward any query params, e.g., ?page=2&search=foo
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query).toString() : ''
  
  // استفاده از API عمومی برای محصولات منتشر شده
  return proxy(event, `${base}/api/products/public${qs}`)
})