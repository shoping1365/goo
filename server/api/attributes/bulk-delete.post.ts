import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  const body = await readBody(event)
  return proxy(event, `${base}/api/attributes/bulk-delete`, {
    method: 'POST',
    body: JSON.stringify(body)
  })
})