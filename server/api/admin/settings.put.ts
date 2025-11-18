import { proxy } from '../_utils/fetchProxy'
import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const payload = await readBody(event)
  const config = useRuntimeConfig()
  
  return await proxy(event, `${config.public.goApiBase}/api/admin/settings`, {
    method: 'PUT',
    body: payload
  })
}) 