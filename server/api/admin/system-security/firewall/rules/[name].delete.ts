export default defineEventHandler(async (event) => {
  const name = getRouterParam(event, 'name')
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/rules/${encodeURIComponent(name!)}`
  const res = await $fetch(url, { method: 'DELETE', headers: getRequestHeaders(event) as any })
  return res
})


