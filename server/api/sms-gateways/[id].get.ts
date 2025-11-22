import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase
          const id = event.context.params?.id

          if (!id) {
               return {
                    success: false,
                    message: 'شناسه درگاه الزامی است',
                    data: null
               }
          }

          const response = await fetch(`${base}/api/sms-gateways/${id}`, {
               method: 'GET',
               headers: {
                    'Accept': 'application/json',
               },
          })

          if (!response.ok) {
               return {
                    success: false,
                    message: 'خطا در دریافت اطلاعات درگاه',
                    data: null
               }
          }

          const json = await response.json()
          return json
     } catch (_error) {
          return {
               success: false,
               message: 'خطا در ارتباط با سرور',
               data: null
          }
     }
})