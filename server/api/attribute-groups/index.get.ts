import { defineEventHandler, getQuery } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const query = getQuery(event)
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as any).toString() : ''
  return proxy(event, `${base}/api/attribute-groups${qs}`)
})