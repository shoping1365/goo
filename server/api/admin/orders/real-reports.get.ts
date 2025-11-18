import type { H3Event } from 'h3'
import { defineEventHandler } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event: H3Event) => {
  // process.static is a Nuxt 2 property, not available in Nuxt 3/Nitro
  if (process.server && false) {
    return { data: {}, message: 'API disabled during prerendering', success: false }
  }

  const config = useRuntimeConfig()
  const url = `${config.public.goApiBase}/api/admin/orders/analytics`
  
  return await proxy(event, url, { method: 'GET' })
})
