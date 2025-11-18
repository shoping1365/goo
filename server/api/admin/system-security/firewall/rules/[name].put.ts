export default defineEventHandler(async (event) => {
  const name = getRouterParam(event, 'name')
  const body = await readBody(event)
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/rules/${encodeURIComponent(name!)}`
  const res = await $fetch(url, { method: 'PUT', body, headers: getRequestHeaders(event) as any })
  return res
})


