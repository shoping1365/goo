import { defineEventHandler } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const url = new URL(event.node.req.url!, 'http://localhost')
  const baseSlug = url.searchParams.get('base_slug') || ''
  return proxy(event, `${base}/api/products/generate-slug?base_slug=${encodeURIComponent(baseSlug)}`)
})