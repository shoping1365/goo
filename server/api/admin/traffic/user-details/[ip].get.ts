import { defineEventHandler, getRouterParam, getHeader } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const ip = getRouterParam(event, 'ip')
  
  try {
    const data = await $fetch(`${base}/api/admin/traffic/user-details/${ip}`, {
      credentials: 'include',
      headers: { cookie: getHeader(event, 'cookie') || '' }
    })
    return { data }
  } catch {
    return { data: [] }
  }
})