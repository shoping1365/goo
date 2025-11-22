import { defineEventHandler } from 'h3'

export default defineEventHandler(async (_event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     try {
          const response = await fetch(`${base}/admin/sms-patterns/stats`, {
               method: 'GET',
               headers: { 'Accept': 'application/json' }
          })
          const json = await response.json()
          return json
     } catch (error) {
          return { status: 'error', message: 'خطا در دریافت آمار پترن‌ها', error: error?.message || error }
     }
})