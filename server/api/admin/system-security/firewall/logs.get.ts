export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/logs`
  const res = await $fetch(url, { headers: getRequestHeaders(event) as any })
  return res
})


