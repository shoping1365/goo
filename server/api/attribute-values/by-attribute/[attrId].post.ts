import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  const { attrId } = event.context.params as { attrId: string }
  const body = await readBody(event)
  return proxy(event, `${base}/api/attribute-values/by-attribute/${attrId}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
})