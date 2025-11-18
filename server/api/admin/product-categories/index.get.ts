import { defineEventHandler, getQuery } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const query = getQuery(event)
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as Record<string, string>).toString() : ''

  // Use proxy helper from _utils
  const { proxy } = await import('../../_utils/fetchProxy')
  return proxy(event, `${base}/api/product-categories${qs}`)
})

