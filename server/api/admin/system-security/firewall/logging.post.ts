export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const base = useRuntimeConfig().public.apiBase || ''
  return await $fetch(`${base}/api/admin/system-security/firewall/logging`, { method: 'POST', body, headers: getRequestHeaders(event) as Record<string, string> })
})


