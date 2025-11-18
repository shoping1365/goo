export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const url = `${base}/api/admin/system/db-pool`
  const data = await $fetch(url, { headers: event.node.req.headers as any })
  return data
})