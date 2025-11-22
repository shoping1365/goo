import { defineEventHandler, getQuery } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const query = getQuery(event)
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as Record<string, string>).toString() : ''
  return proxy(event, `${base}/api/attributes${qs}`)
});