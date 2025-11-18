import { defineEventHandler } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const url = new URL(event.node.req.url!, 'http://localhost')
  const slug = url.searchParams.get('slug') || ''
  return proxy(event, `${base}/api/products/check-slug?slug=${encodeURIComponent(slug)}`)
})