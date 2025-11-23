import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const { id } = event.context.params as { id: string }
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const body = await readBody(event)
  return proxy(event, `${base}/api/brands/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
})
