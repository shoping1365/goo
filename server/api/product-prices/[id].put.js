import { defineEventHandler, readBody, getRequestHeader } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { id } = event.context.params
  const body = await readBody(event)
  const auth = getRequestHeader(event, 'authorization')
  const cookie = getRequestHeader(event, 'cookie')
  return proxy(event, `${base}/api/product-prices/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      ...(auth ? { Authorization: auth } : {}),
      ...(cookie ? { Cookie: cookie } : {})
    },
    body: JSON.stringify(body)
  })
})