export default defineEventHandler(async (_event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await $fetch(`${base}/api/admin/system/reset-database`, { method: 'POST' })
})