import { defineEventHandler, getRouterParam, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const body = await readBody(event)
  const id = getRouterParam(event, 'id')

  // fetchProxy automatically forwards cookies and creates Authorization header
  return proxy(event, `${base}/api/products/${id}/specs`, {
    method: 'POST',
    body
  })
})