export default defineEventHandler(async (event) => {
  const body = await readBody<{ mode: 'low'|'medium'|'high'|'strict' }>(event)
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/mode`
  const res = await $fetch(url, { method: 'POST', body, headers: getRequestHeaders(event) as any })
  return res
})


