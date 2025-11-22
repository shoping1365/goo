// @ts-nocheck
export default defineEventHandler(async (event) => {
          const config = useRuntimeConfig()
          const apiBaseUrl = config.public.goApiBase
     let body: unknown = null

     try {
          body = await readBody(event)

          // اعتبارسنجی داده‌ها
          if (!body?.sourcePath || !body?.targetPath) {
               throw createError({
                    statusCode: 400,
                    message: 'مسیر مبدا و مقصد الزامی است'
               })
          }

          // بررسی احراز هویت
          const authToken = getCookie(event, 'auth-token') || getCookie(event, 'access_token')

          // درخواست به Go backend
          const requestBody = {
               sourcePath: body.sourcePath,
               targetPath: body.targetPath,
               code: body.code || 301,
               groupName: body.groupName || 'دسته‌بندی جدید'
          }
          const response = await $fetch(`${apiBaseUrl}/api/admin/seo/redirects`, {
               method: 'POST',
               headers: {
                    'Authorization': `Bearer ${authToken}`,
                    'Content-Type': 'application/json'
               },
               body: requestBody
          })

          return response
     } catch (error: unknown) {
          const err = error as { message?: string; statusCode?: number; statusMessage?: string; data?: unknown }

          // بررسی نوع خطا
          if (err.message?.includes('ECONNREFUSED') || err.message?.includes('fetch failed')) {
               return {
                    success: false,
                    message: 'سرور Go در حال اجرا نیست - لطفاً سرور را اجرا کنید',
                    error: err.message
               }
          }

          throw createError({
               statusCode: err.statusCode || 500,
               message: 'خطا در ایجاد ریدایرکت: ' + (err.message || err.statusMessage || 'خطای نامشخص')
          })
     }
})