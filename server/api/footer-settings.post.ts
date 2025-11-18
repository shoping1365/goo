import { defineEventHandler, readBody, createError } from 'h3'
import { fetchGo } from './_utils/fetchGo'

interface FooterPayload {
     name: string
     description?: string
     page_selection: string
     specific_pages?: string
     excluded_pages?: string
     is_active: boolean
     layers?: any[]
}

interface FooterResponse {
     success: boolean
     data?: any
     message?: string
}

export default defineEventHandler(async (event): Promise<FooterResponse> => {
     try {
          const body = await readBody(event) as FooterPayload

          console.log('درخواست ایجاد فوتر:', body)
          console.log('Body type:', typeof body)
          console.log('Body keys:', Object.keys(body || {}))

          // اعتبارسنجی داده‌ها
          if (!body) {
               throw createError({
                    statusCode: 400,
                    message: 'بدنه درخواست خالی است'
               })
          }

          if (!body.name || !body.name.trim()) {
               throw createError({
                    statusCode: 400,
                    message: 'عنوان فوتر الزامی است'
               })
          }

          if (!body.page_selection) {
               throw createError({
                    statusCode: 400,
                    message: 'نوع نمایش صفحه الزامی است'
               })
          }

          if (typeof body.is_active !== 'boolean') {
               throw createError({
                    statusCode: 400,
                    message: 'وضعیت فعال بودن باید boolean باشد'
               })
          }

          console.log('Validation passed, proceeding with API call...')

          let responseData: any
          try {
               responseData = await fetchGo(event, '/api/admin/footer-settings', {
                    method: 'POST',
                    body
               })
          } catch (fetchErr: any) {
               console.log('Error response:', fetchErr?.data)
               throw createError({
                    statusCode: fetchErr?.statusCode || fetchErr?.status || 500,
                    message: fetchErr?.data?.message || fetchErr?.data?.error || fetchErr?.message || 'خطا در ایجاد فوتر',
                    data: fetchErr?.data
               })
          }

          console.log('پاسخ ایجاد فوتر:', responseData)

          return {
               success: true,
               data: responseData?.data || responseData,
               message: 'فوتر با موفقیت ایجاد شد'
          }

     } catch (error: any) {
          console.error('خطا در ایجاد فوتر:', error)
          console.error('Error details:', {
               message: error.message,
               statusCode: error.statusCode,
               data: error.data
          })

          if (error.data) {
               throw createError({
                    statusCode: error.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در ایجاد فوتر',
                    data: error.data
               })
          }

          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})
