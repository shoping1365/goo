import { defineEventHandler, getQuery } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const query = getQuery(event)
  const queryString = Object.keys(query).length
    ? '?' + Object.entries(query).map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(v)}`).join('&')
    : ''
  try {
    const response = await $fetch(`${base}/api/posts/check-slug${queryString}`)
    return response
  } catch (error) {
    console.error('Posts API - Error:', error)
    throw error
  }
})