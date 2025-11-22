export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/logs`
  const res = await $fetch(url, { method: 'DELETE', headers: getRequestHeaders(event) as Record<string, string> })
  return res
})


