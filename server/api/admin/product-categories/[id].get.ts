import { useRuntimeConfig } from '#imports'
import { defineEventHandler, getRouterParam } from 'h3'
import { proxy } from '~/server/api/_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const id = getRouterParam(event, 'id')

  return proxy(event, `${base}/api/product-categories/${id}`)
})