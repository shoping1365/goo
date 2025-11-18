export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { getQuery, createError } = await import('h3')
  const { proxy } = await import('../_utils/fetchProxy')

  const query = getQuery(event)
  
  // Validate required query parameter
  if (!query.q || typeof query.q !== 'string' || query.q.trim().length < 1) {
    throw createError({
      statusCode: 400,
      message: 'عبارت جستجو الزامی است'
    })
  }

  // forward any query params, e.g., ?q=phone&type=products&limit=10
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as Record<string, string>).toString() : ''
  return proxy(event, `${base}/api/search${qs}`)
})