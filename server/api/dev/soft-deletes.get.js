export default defineEventHandler(async (event) => {
  const { getQuery } = await import('h3')
  const { proxy } = await import('../_utils/fetchProxy')
  
  let query = getQuery(event)
  if (!query) query = {}
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query).toString() : ''
  const base = getGoApiBaseUrl()
  
  return proxy(event, `${base}/api/dev/soft-deletes${qs}`)
})