import { useRuntimeConfig } from '#imports'
import { createError, defineEventHandler, getRouterParam, parseCookies, readBody } from 'h3'

interface FooterPayload {
     Title: string
     Description?: string
     PageSelection: string
     SpecificPages?: string
     ExcludedPages?: string
     IsActive: boolean
     Layers?: unknown[]
}

interface FooterResponse {
     success: boolean
     data?: unknown
     message?: string
}

export default defineEventHandler(async (event): Promise<FooterResponse> => {
     try {
          const id = getRouterParam(event, 'id')
          const config = useRuntimeConfig()
          const body = await readBody(event) as FooterPayload

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

          return {
               success: true,
               data: responseData?.data || responseData,
               message: 'فوتر با موفقیت ویرایش شد'
          }

     } catch (error: unknown) {
          const err = error as { statusCode?: number; data?: unknown; message?: string; error?: string }
          console.error('خطا در ویرایش فوتر:', err)

          if (err.data) {
               throw createError({
                    statusCode: err.statusCode || 500,
                    message: err.data.message || err.data.error || 'خطا در ویرایش فوتر',
                    data: err.data
               })
          }

          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})
