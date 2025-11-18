// server/api/admin/security/settings.get.ts
// برمی‌گرداند وضعیت فعلی قوانین امنیت (mock: در حافظه)

import { defineEventHandler, createError } from 'h3'

interface SecuritySettings {
  csrf: boolean
  rateLimiter: boolean
  xss: boolean
  frameOptions: boolean
  hsts: boolean
  disabledUntil?: number // timestamp millis
}

// کلید ذخیره در حافظه globalThis
const KEY = '__security_settings__'

function getSettings(): SecuritySettings {
  // @ts-ignore
  if (!globalThis[KEY]) {
    // مقادیر پیش‌فرض (همه فعال)
    // @ts-ignore
    globalThis[KEY] = {
      csrf: false,
      rateLimiter: false,
      xss: false,
      frameOptions: false,
      hsts: false,
    } as SecuritySettings
  }
  // @ts-ignore
  return globalThis[KEY] as SecuritySettings
}

export default defineEventHandler(async (event) => {
  try {
    // اینجا باید از دیتابیس تنظیمات امنیت خوانده شود
    // فعلاً مقادیر پیش‌فرض برمی‌گردانیم

    const securitySettings = {
      csrf: true,
      rateLimiter: true,
      xss: true,
      frameOptions: true,
      hsts: true,
      disabledUntil: null
    }

    const recaptchaSettings = {
      version: 'v2',
      siteKey: '',
      secretKey: '',
      theme: 'light',
      scoreThreshold: 0.5
    }

    const rateLimitSettings = {
      admin: {
        requestsPerMinute: 100,
        requestsPerHour: 1000,
        requestsPerDay: 10000
      },
      public: {
        requestsPerMinute: 30,
        requestsPerHour: 300,
        requestsPerDay: 5000
      }
    }

    return {
      success: true,
      data: {
        security: securitySettings,
        recaptcha: recaptchaSettings,
        rateLimit: rateLimitSettings
      }
    }
  } catch (error) {
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت تنظیمات امنیت'
    })
  }
})