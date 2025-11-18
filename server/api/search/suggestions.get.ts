export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { getQuery } = await import('h3')
  const { proxy } = await import('../_utils/fetchProxy')

  const query = getQuery(event)
  
  // Validate required query parameter
  if (!query.q || typeof query.q !== 'string' || query.q.trim().length < 1) {
    return {
      success: true,
      data: []
    }
  }

  // forward any query params, e.g., ?q=phone
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as Record<string, string>).toString() : ''
  return proxy(event, `${base}/api/search/suggestions${qs}`)
})