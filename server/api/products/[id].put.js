import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { id } = event.context.params
  const body = await readBody(event)
  return proxy(event, `${base}/api/products/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
})