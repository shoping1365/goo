export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase

          // دریافت ID از پارامترهای مسیر
          const id = getRouterParam(event, 'id')

          if (!id) {
               throw createError({
                    statusCode: 400,
                    message: 'شناسه درگاه پیامک الزامی است',
               })
          }

          // درخواست به بک‌اند
          const response = await $fetch(`${base}/api/sms-gateways/${id}`, {
               method: 'DELETE',
               headers: {
                    'Content-Type': 'application/json',
               },
          })

          return response
     } catch (error) {
          console.error('خطا در حذف درگاه پیامک:', error)
          throw createError({
               statusCode: 500,
               message: 'خطا در حذف درگاه پیامک',
          })
     }
})