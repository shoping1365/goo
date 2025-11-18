import { proxy } from '../_utils/fetchProxy'
import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const config = useRuntimeConfig()

  const payload = typeof body === 'string' ? body : JSON.stringify(body ?? {})

  return await proxy(event, `${config.public.goApiBase}/api/admin/shop-settings`, {
    method: 'PUT',
    body: payload,
    headers: {
      'Content-Type': 'application/json'
    }
  })
}) 