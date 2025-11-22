export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  // Proxy to Go backend
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/status`
  const res = await $fetch(url, { headers: getRequestHeaders(event) as Record<string, string> })
  return res
})


