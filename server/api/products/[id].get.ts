import { defineEventHandler } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { id } = event.context.params as { id: string }
  
  // اگر پارامتر شبیه sku-XXXX بود، به صورت lookup=sku به بک‌اند پاس می‌دهیم
  const isSkuPrefixed = typeof id === 'string' && id.startsWith('sku-')
  const identifier = isSkuPrefixed ? id.substring(4) : id
  const lookupParam = isSkuPrefixed ? '?lookup=sku' : ''
  return proxy(event, `${base}/api/products/${identifier}${lookupParam}`)
})
