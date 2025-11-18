import { defineEventHandler, createError, readRawBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface MobileAppHeaderCreateResponse {
     success: boolean
     data: any
     message?: string
}

export default defineEventHandler(async (event): Promise<MobileAppHeaderCreateResponse> => {
     try {
          let body: any = {}

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
                    } catch (parseErr) {
                         console.error('❌ خطا در parse بدنه هدر موبایل:', parseErr, raw)
                         throw createError({ statusCode: 400, message: 'ساخت هدر موبایل: بدنه نامعتبر است' })
                    }
               }
          } else {
               const rawBody = await readRawBody(event)
               if (typeof rawBody === 'string' && rawBody.length) {
                    try {
                         body = JSON.parse(rawBody)
                    } catch (parseErr) {
                         console.error('❌ خطا در parse بدنه هدر موبایل:', parseErr, rawBody)
                         throw createError({ statusCode: 400, message: 'ساخت هدر موبایل: بدنه نامعتبر است' })
                    }
               } else if (rawBody && typeof rawBody === 'object') {
                    body = rawBody
               }
          }

          console.log('درخواست ایجاد هدر موبایل جدید:', body)

          let responseData: any
          try {
               responseData = await fetchGo(event, '/api/admin/mobile-app-header-settings', {
                    method: 'POST',
                    body
               })
          } catch (fetchErr: any) {
               console.error('❌ fetchGo ایجاد هدر موبایل شکست خورد:', {
                    statusCode: fetchErr?.statusCode,
                    status: fetchErr?.status,
                    message: fetchErr?.message,
                    data: fetchErr?.data || fetchErr?.response?._data
               })
               throw createError({
                    statusCode: fetchErr?.statusCode || fetchErr?.status || fetchErr?.response?.status || 500,
                    message: fetchErr?.data?.message || fetchErr?.data?.error || fetchErr?.message || 'خطا در ایجاد هدر موبایل',
                    data: fetchErr?.data || fetchErr?.response?._data
               })
          }

          console.log('پاسخ Go برای ایجاد هدر موبایل:', responseData)

          console.log('پاسخ ایجاد هدر موبایل:', responseData)

          return {
               success: true,
               data: responseData?.data || null
          }

     } catch (error: any) {
          if (error?.statusCode || error?.status) {
               throw createError({
                    statusCode: error?.statusCode || error?.status,
                    message: error?.message || error?.data?.message || error?.data?.error || 'خطا در ایجاد هدر موبایل',
                    data: error?.data
               })
          }
          console.error('خطا در ایجاد هدر موبایل:', error)

          // اگر خطا از سرور Go آمده باشد
          if (error.data) {
               throw createError({
                    statusCode: error.statusCode || 500,
                    message: error.data.message || error.data.error || 'خطا در ایجاد هدر موبایل',
                    data: error.data
               })
          }

          // اگر خطای شبکه باشد
          throw createError({
               statusCode: 500,
               message: 'خطا در اتصال به سرور'
          })
     }
})