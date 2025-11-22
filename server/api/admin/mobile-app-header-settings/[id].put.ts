import { createError, defineEventHandler, getRouterParam, readBody, readRawBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MobileAppHeaderUpdateResponse {
     success: boolean
     data: unknown
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
                    } catch (_parseErr) {
                         // console.error('❌ خطا در parse بدنه به‌روزرسانی هدر موبایل:', parseErr, raw)
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

          // console.log('درخواست به‌روزرسانی هدر موبایل با ID:', id, 'داده:', body)

          let responseData: unknown
          try {
               responseData = await fetchGo(event, `/api/admin/mobile-app-header-settings/${id}`, {
                    method: 'PUT',
                    body
               })
          } catch (fetchErr: unknown) {
               const err = fetchErr as { statusCode?: number; status?: number; message?: string; data?: { message?: string; error?: string }; response?: { status?: number; _data?: unknown } }
               // console.error('❌ fetchGo به‌روزرسانی هدر موبایل شکست خورد:', {
               //      statusCode: err?.statusCode,
               //      status: err?.status,
               //      message: err?.message,
               //      data: err?.data || err?.response?._data
               // })
               throw createError({
                    statusCode: err?.statusCode || err?.status || err?.response?.status || 500,
                    message: err?.data?.message || err?.data?.error || err?.message || 'خطا در به‌روزرسانی هدر موبایل',
                    data: err?.data || err?.response?._data
               })
          }

          // console.log('پاسخ به‌روزرسانی هدر موبایل:', responseData)

          if (responseData && typeof responseData === 'object') {
               const payload = responseData as Record<string, unknown>
               return {
                    success: payload.success !== false,
                    data: payload.data || null
               }
          }

          return {
               success: true,
               data: responseData || null
          }

     } catch (error: unknown) {
          const err = error as { statusCode?: number; status?: number; message?: string; data?: { message?: string; error?: string } }
          if (err?.statusCode || err?.status) {
               throw createError({
                    statusCode: err?.statusCode || err?.status,
                    message: err?.message || err?.data?.message || err?.data?.error || 'خطا در به‌روزرسانی هدر موبایل',
                    data: err?.data
               })
          }

          // console.error('خطا در به‌روزرسانی هدر موبایل:', error)

          if ((error as { data?: { message?: string; error?: string } })?.data) {
               const e = error as { statusCode?: number; data: { message?: string; error?: string } }
               throw createError({
                    statusCode: e?.statusCode || 500,
                    message: e.data.message || e.data.error || 'خطا در به‌روزرسانی هدر موبایل',
                    data: e.data
               })
          }

          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})