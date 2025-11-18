import { defineEventHandler, getRequestHeader } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const auth = getRequestHeader(event, 'authorization')
  const cookie = getRequestHeader(event, 'cookie')
  return proxy(event, `${base}/api/admin/monitoring/status`, {
    method: 'GET',
    headers: {
      ...(auth ? { Authorization: auth } : {}),
      ...(cookie ? { Cookie: cookie } : {}),
    },
  })
})