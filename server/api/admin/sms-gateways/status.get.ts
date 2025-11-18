import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
     const config = useRuntimeConfig()
     const base = config.public.goApiBase
     try {
          const response = await $fetch(`${base}/admin/sms-gateways/status`, {
               method: 'GET',
               headers: { 'Accept': 'application/json' }
          })
          return response
     } catch (error) {
          console.error('خطا در دریافت وضعیت درگاه‌ها:', error)
          return {
               status: 'error',
               message: 'خطا در دریافت وضعیت درگاه‌های پیامکی',
               error: error?.message || error
          }
     }
})