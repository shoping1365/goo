import { proxy } from '../../_utils/fetchProxy'
import { defineEventHandler, getRouterParam, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const id = getRouterParam(event, 'id')
  const body = await readBody(event)
  
  return await proxy(event, `${config.public.goApiBase}/api/roles/${id}/permissions`, {
    method: 'PUT',
    body: body
  })
}) 