import { defineEventHandler, getRouterParam, readBody, parseCookies, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

interface FooterPayload {
     Title: string
     Description?: string
     PageSelection: string
     SpecificPages?: string
     ExcludedPages?: string
     IsActive: boolean
     Layers?: any[]
}

interface FooterResponse {
     success: boolean
     data?: any
     message?: string
}

export default defineEventHandler(async (event): Promise<FooterResponse> => {
     try {
          const id = getRouterParam(event, 'id')
          const config = useRuntimeConfig()
          const body = await readBody(event) as FooterPayload

          console.log('درخواست ویرایش فوتر:', id, body)

          if (!id) {
               throw createError({
                    statusCode: 400,
                    message: 'شناسه فوتر الزامی است'
               })
          }

          // اعتبارسنجی داده‌ها
          if (!body.Title || !body.Title.trim()) {
               throw createError({
                    statusCode: 400,
                    message: 'عنوان فوتر الزامی است'
               })
          }

          if (!body.PageSelection) {
               throw createError({
                    statusCode: 400,
                    message: 'نوع نمایش صفحه الزامی است'
               })
          }

          if (typeof body.IsActive !== 'boolean') {
               throw createError({
                    statusCode: 400,
                    message: 'وضعیت فعال بودن باید boolean باشد'
               })
          }

          const cookies = parseCookies(event)
          const cookieHeader = Object.entries(cookies)
               .map(([key, value]) => `${key}=${value}`)
               .join('; ')

          const response = await fetch(`${config.public.goApiBase}/api/admin/footer-settings/${id}`, {
               method: 'PUT',
               headers: {
                    'Content-Type': 'application/json',
                    'Cookie': cookieHeader
               },
               body: JSON.stringify(body)
          })

          if (!response.ok) {
               const errorData = await response.json().catch(() => ({}))
               throw createError({
                    statusCode: response.status,
                    message: errorData.message || `خطای HTTP: ${response.status}`,
                    data: errorData
               })
          }

          const responseData = await response.json()

          console.log('پاسخ ویرایش فوتر:', responseData)

          return {
               success: true,
               data: responseData?.data || responseData,
               message: 'فوتر با موفقیت ویرایش شد'
          }

     } catch (error: any) {
          console.error('خطا در ویرایش فوتر:', error)

          if (error.data) {
               throw createError({
                    statusCode: error.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در ویرایش فوتر',
                    data: error.data
               })
          }

          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})
