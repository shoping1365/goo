import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase
          const id = event.context.params?.id
          if (!id) {
               return { status: 'error', message: 'شناسه درگاه الزامی است' }
          }
          const response = await fetch(`${base}/sms-gateways/${id}/balance`, {
               method: 'GET',
               headers: { 'Accept': 'application/json' }
          })
          const json = await response.json()
          return json
     } catch (error) {
          return { status: 'error', message: 'خطا در دریافت موجودی', error: error?.message || error }
     }
}
)