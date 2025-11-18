import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  const body = await readBody(event)
  return await $fetch(`${base}/api/admin/system/auto-refresh-settings`, { method: 'POST', body })
})