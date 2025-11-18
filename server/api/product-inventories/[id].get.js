import { defineEventHandler } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { id } = event.context.params
  return proxy(event, `${base}/api/product-inventories/${id}`)
})