import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase
          const body = await readBody(event)
          const response = await fetch(`${base}/api/sms-gateways`, {
               method: 'POST',
               headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
               },
               body: JSON.stringify(body)
          })
          const json = await response.json()
          return json
     } catch (error) {
          return {
               status: 'error',
               message: 'خطا در ثبت درگاه پیامکی',
               error: error?.message || error
          }
     }
})