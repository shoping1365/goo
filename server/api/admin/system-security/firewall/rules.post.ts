export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/rules`
  const res = await $fetch(url, { method: 'POST', body, headers: getRequestHeaders(event) as any })
  return res
})


