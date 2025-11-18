import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     const body = await readBody(event)
     try {
          const response = await fetch(`${base}/admin/sms-patterns`, {
               method: 'POST',
               headers: { 'Accept': 'application/json', 'Content-Type': 'application/json' },
               body: JSON.stringify(body)
          })
          const json = await response.json()
          return json
     } catch (error) {
          return { status: 'error', message: 'خطا در ایجاد پترن', error: error?.message || error }
     }
})