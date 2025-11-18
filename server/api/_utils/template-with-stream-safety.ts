// Template for API handlers with stream safety
// این template برای ایجاد API handlers امن استفاده می‌شود
import type { H3Event } from 'h3'
import { defineEventHandler, createError, readBody } from 'h3'

declare const goApiFetch: (event: H3Event, endpoint: string, options?: { method?: string; body?: unknown }) => Promise<unknown>
declare const isResponseStreamSafe: (event: H3Event) => boolean
declare const safeSendResponse: (event: H3Event, response: unknown) => boolean
declare const safeHandleError: (event: H3Event, error: unknown) => boolean

export default defineEventHandler(async (event) => {
  try {
    // ✅ 1. بررسی وضعیت response stream
    if (!isResponseStreamSafe(event)) {
      console.warn('Response stream already closed, skipping request processing')
      return
    }

    // ✅ 2. دریافت و اعتبارسنجی داده‌های ورودی
    const body = await readBody(event)
    
    // اعتبارسنجی داده‌ها (مثال)
    if (!body || typeof body !== 'object') {
      throw createError({
        statusCode: 400,
        message: 'Invalid request body'
      })
    }

    // ✅ 3. ارسال درخواست به Go backend
    const response = await goApiFetch(event, '/api/your-endpoint', {
      method: 'POST', // یا GET, PUT, DELETE
      body: body
    })

    // ✅ 4. بررسی مجدد وضعیت response stream قبل از ارسال پاسخ
    if (!safeSendResponse(event, response)) {
      return
    }

    return response
  } catch (error: unknown) {
    console.error('خطا در پردازش درخواست:', error)
    
    // ✅ 5. بررسی وضعیت response stream قبل از ارسال خطا
    if (!safeHandleError(event, error)) {
      return
    }
    
    // ✅ 6. مدیریت خطاهای مختلف
    const errorWithStatus = error as { statusCode?: number; message?: string }
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    // خطای عمومی
    throw createError({
      statusCode: 500,
      message: 'خطا در پردازش درخواست'
    })
  }
})