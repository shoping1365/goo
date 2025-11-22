export default defineEventHandler(async (event) => {
  const body = await readBody<{ enabled: boolean }>(event)
  const config = useRuntimeConfig()
  const base = config.public.apiBase || ''
  const url = `${base}/api/admin/system-security/firewall/toggle`
  const res = await $fetch(url, { method: 'POST', body, headers: getRequestHeaders(event) as Record<string, string> })
  return res
})


