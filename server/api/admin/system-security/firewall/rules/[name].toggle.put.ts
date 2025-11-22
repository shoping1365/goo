export default defineEventHandler(async (event) => {
  const name = getRouterParam(event, 'name')
  const query = getQuery(event)
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/rules/${encodeURIComponent(name!)}/toggle?enabled=${encodeURIComponent(String(query.enabled ?? '1'))}`
  const res = await $fetch(url, { method: 'PUT', headers: getRequestHeaders(event) as Record<string, string> })
  return res
})


