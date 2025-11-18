import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     const body = await readBody(event)

     try {
          const response = await fetch(`${base}/admin/sms-patterns/send-by-scope-feature`, {
               method: 'POST',
               headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
               },
               body: JSON.stringify(body)
          })

          const json = await response.json()
          return json
     } catch (error) {
          return {
               success: false,
               message: 'خطا در ارسال پیامک',
               error: error?.message || error
          }
     }
})