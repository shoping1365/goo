import { defineEventHandler, getRequestHeader, getRouterParam } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const cookie = getRequestHeader(event, 'cookie') || ''
  const id = getRouterParam(event, 'id')

  return await $fetch(`${base}/api/users/${id}`, {
    method: 'DELETE',
    credentials: 'include',
    headers: { cookie }
  })
})