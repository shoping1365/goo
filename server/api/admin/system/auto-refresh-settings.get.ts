export default defineEventHandler(async (_event) => {
  const base = getGoApiBaseUrl()
  return await $fetch(`${base}/api/admin/system/auto-refresh-settings`)
})