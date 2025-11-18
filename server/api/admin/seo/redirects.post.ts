// @ts-nocheck
export default defineEventHandler(async (event) => {
          const config = useRuntimeConfig()
          const apiBaseUrl = config.public.goApiBase
     let body: any = null

     try {
          console.log('🔍 Starting redirect creation...')
          body = await readBody(event)
          console.log('📝 Request body:', body)

          // اعتبارسنجی داده‌ها
          if (!body?.sourcePath || !body?.targetPath) {
               console.log('❌ Validation failed: missing sourcePath or targetPath')
               throw createError({
                    statusCode: 400,
                    message: 'مسیر مبدا و مقصد الزامی است'
               })
          }

          console.log('🌐 API Base URL:', apiBaseUrl)

          // بررسی احراز هویت
          const authToken = getCookie(event, 'auth-token') || getCookie(event, 'access_token')
          console.log('🔑 Auth token exists:', !!authToken)

          // درخواست به Go backend
          const requestBody = {
               sourcePath: body.sourcePath,
               targetPath: body.targetPath,
               code: body.code || 301,
               groupName: body.groupName || 'دسته‌بندی جدید'
          }
          console.log('📤 Sending request to backend:', requestBody)

          const response = await $fetch(`${apiBaseUrl}/api/admin/seo/redirects`, {
               method: 'POST',
               headers: {
                    'Authorization': `Bearer ${authToken}`,
                    'Content-Type': 'application/json'
               },
               body: requestBody
          })

          console.log('✅ Backend response:', response)
          return response
     } catch (error: any) {
          console.error('❌ Error in redirect creation:', error)
          console.error('Error details:', {
               message: error.message,
               statusCode: error.statusCode,
               statusMessage: error.statusMessage,
               data: error.data
          })

          // بررسی نوع خطا
          if (error.message?.includes('ECONNREFUSED') || error.message?.includes('fetch failed')) {
               console.log('🔄 Connection refused - Go server not running')
               return {
                    success: false,
                    message: 'سرور Go در حال اجرا نیست - لطفاً سرور را اجرا کنید',
                    error: error.message
               }
          }

          throw createError({
               statusCode: error.statusCode || 500,
               message: 'خطا در ایجاد ریدایرکت: ' + (error.message || error.statusMessage || 'خطای نامشخص')
          })
     }
})