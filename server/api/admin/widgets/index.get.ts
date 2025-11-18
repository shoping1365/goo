export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const { proxy } = await import('~/server/api/_utils/fetchProxy')
  const query = getQuery(event)

  // fetchProxy automatically forwards cookies and creates Authorization header
  return proxy(event, `${base}/api/admin/widgets?${new URLSearchParams(query as any).toString()}`)
})