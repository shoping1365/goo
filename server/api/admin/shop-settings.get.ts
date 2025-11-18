import { proxy } from '../_utils/fetchProxy'
import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  return await proxy(event, `${config.public.goApiBase}/api/admin/shop-settings`)
}) 