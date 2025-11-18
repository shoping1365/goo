import { defineEventHandler, readBody, getRequestHeader } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const auth = getRequestHeader(event, 'authorization')
  const cookie = getRequestHeader(event, 'cookie')
  return proxy(event, `${base}/api/admin/monitoring/toggle`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ...(auth ? { Authorization: auth } : {}),
      ...(cookie ? { Cookie: cookie } : {}),
    },
    body: JSON.stringify(body),
  })
})