import { defineEventHandler, getRouterParam } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const base = getGoApiBaseUrl()
  const id = getRouterParam(event, 'id')
  return proxy(event, `${base}/api/attributes/${id}`, {
    method: 'DELETE'
  })
})