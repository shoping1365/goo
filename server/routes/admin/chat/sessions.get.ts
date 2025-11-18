import { defineEventHandler, getQuery, getCookie } from 'h3'
import { useRuntimeConfig } from '#imports'

// Compatibility route: handle /admin/chat/sessions and proxy to Go admin chat API
export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  
  const query = getQuery(event)
  const queryParams: Record<string, string> = {}
  for (const [key, value] of Object.entries(query)) {
    if (typeof value === 'string') {
      queryParams[key] = value
    } else if (Array.isArray(value)) {
      queryParams[key] = value.join(',')
    }
  }
  const qs = Object.keys(queryParams).length ? '?' + new URLSearchParams(queryParams).toString() : ''
  const cookie = getCookie(event, 'auth-token') || getCookie(event, 'access_token') || ''
  
  const url = `${base}/api/chat/admin/sessions${qs}`
  return await $fetch(url, {
    method: 'GET',
    credentials: 'include',
    headers: { cookie }
  })
})