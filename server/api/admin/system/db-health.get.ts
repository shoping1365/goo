export default defineEventHandler(async (event) => {
  const data = await $fetch('/api/db-health', { headers: { cookie: getHeader(event, 'cookie') || '' } })
  return data
})