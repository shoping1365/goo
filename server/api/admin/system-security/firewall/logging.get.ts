export default defineEventHandler(async (event) => {
  const base = useRuntimeConfig().public.apiBase || ''
  return await $fetch(`${base}/api/admin/system-security/firewall/logging`, { headers: getRequestHeaders(event) as Record<string, string> })
})


