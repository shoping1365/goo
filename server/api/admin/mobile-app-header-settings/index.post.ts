import { createError, defineEventHandler, readRawBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MobileAppHeaderCreateResponse {
     success: boolean
     data: unknown
     message?: string
}

export default defineEventHandler(async (event): Promise<MobileAppHeaderCreateResponse> => {
     try {
          let body: unknown = {}

          if (event.node?.req) {
               const nodeReq = event.node.req
               const raw = await new Promise<string>((resolve, reject) => {
                    const chunks: Buffer[] = []
                    nodeReq.on('data', (chunk) => {
                         chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk))
                    })
                    nodeReq.on('end', () => {
                         resolve(Buffer.concat(chunks).toString('utf8'))
                    })
                    nodeReq.on('error', reject)
               })

               if (raw?.length) {
                    try {
                         body = JSON.parse(raw)
                    } catch (_parseErr) {
                         // console.error('❌ خطا در parse بدنه هدر موبایل:', parseErr, raw)
                         throw createError({ statusCode: 400, message: 'ساخت هدر موبایل: بدنه نامعتبر است' })
                    }
               }
          } else {
               const rawBody = await readRawBody(event)
               if (typeof rawBody === 'string' && rawBody.length) {
                    try {
                         body = JSON.parse(rawBody)
                    } catch (_parseErr) {
                         // console.error('❌ خطا در parse بدنه هدر موبایل:', parseErr, rawBody)
                         throw createError({ statusCode: 400, message: 'ساخت هدر موبایل: بدنه نامعتبر است' })
                    }
               } else if (rawBody && typeof rawBody === 'object') {
                    body = rawBody
               }
          }

          // console.log('درخواست ایجاد هدر موبایل جدید:', body)

          let responseData: unknown
          try {
               responseData = await fetchGo(event, '/api/admin/mobile-app-header-settings', {
                    method: 'POST',
                    body
               })
          } catch (fetchErr: unknown) {
               const err = fetchErr as { statusCode?: number; status?: number; message?: string; data?: { message?: string; error?: string }; response?: { status?: number; _data?: unknown } }
               // console.error('❌ fetchGo ایجاد هدر موبایل شکست خورد:', {
               //      statusCode: err?.statusCode,
               //      status: err?.status,
               //      message: err?.message,
               //      data: err?.data || err?.response?._data
               // })
               throw createError({
                    statusCode: err?.statusCode || err?.status || err?.response?.status || 500,
                    message: err?.data?.message || err?.data?.error || err?.message || 'خطا در ایجاد هدر موبایل',
                    data: err?.data || err?.response?._data
               })
          }

          // console.log('پاسخ Go برای ایجاد هدر موبایل:', responseData)

          // console.log('پاسخ ایجاد هدر موبایل:', responseData)

          return {
               success: true,
               data: (responseData as Record<string, unknown>)?.data || null
          }

     } catch (error: unknown) {
          const err = error as { statusCode?: number; status?: number; message?: string; data?: { message?: string; error?: string } }
          if (err?.statusCode || err?.status) {
               throw createError({
                    statusCode: err?.statusCode || err?.status,
                    message: err?.message || err?.data?.message || err?.data?.error || 'خطا در ایجاد هدر موبایل',
                    data: err?.data
               })
          }
          // console.error('خطا در ایجاد هدر موبایل:', error)

          // اگر خطا از سرور Go آمده باشد
          if (err.data) {
               throw createError({
                    statusCode: err.statusCode || 500,
                    message: err.data.message || err.data.error || 'خطا در ایجاد هدر موبایل',
                    data: err.data
               })
          }

          // اگر خطای شبکه باشد
          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})