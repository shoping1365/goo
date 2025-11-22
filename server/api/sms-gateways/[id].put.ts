import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase
          const id = event.context.params?.id
          const body = await readBody(event)

          if (!id) {
               return {
                    success: false,
                    message: 'شناسه درگاه الزامی است',
                    data: null
               }
          }

          const response = await fetch(`${base}/api/sms-gateways/${id}`, {
               method: 'PUT',
               headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
               },
               body: JSON.stringify(body)
          })

          if (!response.ok) {
               return {
                    success: false,
                    message: 'خطا در ویرایش درگاه',
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