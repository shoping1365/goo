import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     try {
          const response = await fetch(`${base}/api/sms-gateways`, {
               method: 'GET',
               headers: { 'Accept': 'application/json' }
          })
          const json = await response.json()
          return json
     } catch (error: any) {
          return { status: 'error', message: 'خطا در دریافت لیست درگاه‌ها', error: error?.message || error }
     }
})