import { defineEventHandler } from '#imports'
import { getRequestHeader } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const cookie = getRequestHeader(event, 'cookie') || ''
  const id = event.context.params?.id
  const limit = getQuery(event)?.limit || 5

  return await $fetch(`${base}/api/users/${id}/notifications?limit=${limit}`, {
    method: 'GET',
    credentials: 'include',
    headers: { cookie }
  })
})