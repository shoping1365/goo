import { createError, defineEventHandler, readBody } from 'h3'
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

          let responseData: unknown
          try {
               responseData = await fetchGo(event, '/api/admin/footer-settings', {
                    method: 'POST',
                    body
               })
          } catch (fetchErr: unknown) {
               const err = fetchErr as { statusCode?: number; status?: number; message?: string; data?: any }
               throw createError({
                    statusCode: err?.statusCode || err?.status || 500,
                    message: err?.data?.message || err?.data?.error || err?.message || 'خطا در ایجاد فوتر',
                    data: err?.data
               })
          }

          return {
               success: true,
               data: (responseData as any)?.data || responseData,
               message: 'فوتر با موفقیت ایجاد شد'
          }

     } catch (error: unknown) {
          const err = error as { statusCode?: number; message?: string; data?: any }
          console.error('خطا در ایجاد فوتر:', error)
          console.error('Error details:', {
               message: err.message,
               statusCode: err.statusCode,
               data: err.data
          })

          if (err.data) {
               throw createError({
                    statusCode: err.statusCode || 500,
                    message: err.data.message || err.data.error || 'خطا در ایجاد فوتر',
                    data: err.data
               })
          }

          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})
