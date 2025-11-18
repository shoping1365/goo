import { defineEventHandler, createError, getRouterParam, readRawBody, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MobileAppHeaderUpdateResponse {
     success: boolean
     data: any
     message?: string
}

export default defineEventHandler(async (event): Promise<MobileAppHeaderUpdateResponse> => {
     try {
          const id = getRouterParam(event, 'id')

          let body = await readBody(event)

          if (!body || typeof body !== 'object') {
               const raw = await readRawBody(event)
               if (raw && typeof raw === 'string') {
                    try {
                         body = JSON.parse(raw)
                    } catch (parseErr) {
                         console.error('❌ خطا در parse بدنه به‌روزرسانی هدر موبایل:', parseErr, raw)
                         throw createError({ statusCode: 400, message: 'به‌روزرسانی هدر موبایل: بدنه نامعتبر است' })
                    }
               }
          }

          if (!id) {
               throw createError({
                    statusCode: 400,
                    message: 'شناسه هدر موبایل الزامی است'
               })
          }

          console.log('درخواست به‌روزرسانی هدر موبایل با ID:', id, 'داده:', body)

          let responseData: any
          try {
               responseData = await fetchGo(event, `/api/admin/mobile-app-header-settings/${id}`, {
                    method: 'PUT',
                    body
               })
          } catch (fetchErr: any) {
               console.error('❌ fetchGo به‌روزرسانی هدر موبایل شکست خورد:', {
                    statusCode: fetchErr?.statusCode,
                    status: fetchErr?.status,
                    message: fetchErr?.message,
                    data: fetchErr?.data || fetchErr?.response?._data
               })
               throw createError({
                    statusCode: fetchErr?.statusCode || fetchErr?.status || fetchErr?.response?.status || 500,
                    message: fetchErr?.data?.message || fetchErr?.data?.error || fetchErr?.message || 'خطا در به‌روزرسانی هدر موبایل',
                    data: fetchErr?.data || fetchErr?.response?._data
               })
          }

          console.log('پاسخ به‌روزرسانی هدر موبایل:', responseData)

          if (responseData && typeof responseData === 'object') {
               const payload = responseData as Record<string, any>
               return {
                    success: payload.success !== false,
                    data: payload.data || null
               }
          }

          return {
               success: true,
               data: responseData || null
          }

     } catch (error: any) {
          if (error?.statusCode || error?.status) {
               throw createError({
                    statusCode: error?.statusCode || error?.status,
                    message: error?.message || error?.data?.message || error?.data?.error || 'خطا در به‌روزرسانی هدر موبایل',
                    data: error?.data
               })
          }

          console.error('خطا در به‌روزرسانی هدر موبایل:', error)

          if (error?.data) {
               throw createError({
                    statusCode: error?.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در به‌روزرسانی هدر موبایل',
                    data: error.data
               })
          }

          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})