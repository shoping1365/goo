import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     try {
          const response = await fetch(`${base}/admin/sms-patterns`, {
               method: 'GET',
               headers: { 'Accept': 'application/json' }
          })
          const json = await response.json()
          const data = Array.isArray(json.data) ? json.data : []
          return {
               status: json.status,
               data
          }
     } catch (error: any) {
          return { status: 'error', message: 'خطا در دریافت لیست پترن‌ها', error: error?.message || error }
     }
})