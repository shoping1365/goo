import { defineEventHandler } from 'h3'
import { proxy } from '../../../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const { id, attrId } = event.context.params as { id: string; attrId: string }
  const base = getGoApiBaseUrl()
  return proxy(event, `${base}/api/attribute-groups/${id}/attributes/${attrId}`, {
    method: 'DELETE'
  })
})