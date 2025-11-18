import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  const id = event.context.params!.id as string
  const body = await readBody(event)
  return proxy(event, `${base}/api/attribute-groups/${id}/attributes`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
})