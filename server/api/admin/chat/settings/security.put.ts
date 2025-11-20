import type { H3Event } from 'h3';
import { createError, defineEventHandler, readBody } from 'h3';

declare const goApiFetch: (event: H3Event, endpoint: string, options?: { method?: string; body?: unknown }) => Promise<unknown>
declare const isResponseStreamSafe: (event: H3Event) => boolean
declare const safeSendResponse: (event: H3Event, response: unknown) => boolean
declare const safeHandleError: (event: H3Event, error: unknown) => boolean

export default defineEventHandler(async (event) => {
  try {
    // بررسی وضعیت response stream
    if (!isResponseStreamSafe(event)) {
      // console.warn('Response stream already closed, skipping request processing')
      return
    }

    // Get request body
    const body = await readBody(event)

    // Validate security settings
    const bodyAsObject = body as { maxMessageLength?: string | number; blockDuration?: string | number; cooldownDuration?: string | number; strictMode?: boolean; enableAutoBlock?: boolean; enableRateLimit?: boolean }
    const validatedSettings = {
      maxMessageLength: Math.max(100, Math.min(5000, typeof bodyAsObject.maxMessageLength === 'number' ? bodyAsObject.maxMessageLength : parseInt(String(bodyAsObject.maxMessageLength || 1000)))),
      blockDuration: Math.max(1, Math.min(1440, typeof bodyAsObject.blockDuration === 'number' ? bodyAsObject.blockDuration : parseInt(String(bodyAsObject.blockDuration || 10)))),
      cooldownDuration: Math.max(10, Math.min(300, typeof bodyAsObject.cooldownDuration === 'number' ? bodyAsObject.cooldownDuration : parseInt(String(bodyAsObject.cooldownDuration || 30)))),
      strictMode: Boolean(bodyAsObject.strictMode),
      enableAutoBlock: Boolean(bodyAsObject.enableAutoBlock),
      enableRateLimit: Boolean(bodyAsObject.enableRateLimit)
    }

    // Forward request to Go backend
    const response = await goApiFetch(event, '/api/chat/admin/settings/security', {
      method: 'PUT',
      body: validatedSettings
    })

    // بررسی مجدد وضعیت response stream قبل از ارسال پاسخ
    if (!safeSendResponse(event, response)) {
      return
    }

    return response
  } catch (error: unknown) {
    // console.error('خطا در به‌روزرسانی تنظیمات امنیتی چت:', error)

    // بررسی وضعیت response stream قبل از ارسال خطا
    if (!safeHandleError(event, error)) {
      return
    }

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در به‌روزرسانی تنظیمات امنیتی چت'
    })
  }
})