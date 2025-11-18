import { defineEventHandler, readBody } from 'h3'
import { useRuntimeConfig } from '#imports'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const body = await readBody(event)

  // fetchProxy automatically forwards cookies and creates Authorization header
  return proxy(event, `${base}/api/product-categories`, {
    method: 'POST',
    body: JSON.stringify(body)
  })
})