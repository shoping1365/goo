import { readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { proxy } = await import('../_utils/fetchProxy')
  const body = await readBody(event)
  return proxy(event, `${base}/api/product-categories`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
})