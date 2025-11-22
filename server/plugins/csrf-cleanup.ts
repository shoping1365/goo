// server/plugins/csrf-cleanup.ts
// اگر ماژول CSRF در nuxt.config.ts غیرفعال باشد، این پلاگین کوکی قدیمی "csrf"
// را با maxAge 0 پاک می‌کند تا در محیط dev یا staging مرورگر دچار کوکی بی‌مصرف نشود.

import { appendHeader, getCookie, setCookie } from 'h3'

export default defineNitroPlugin((nitro) => {
  nitro.hooks.hook('request', (event) => {
    // در هر درخواست وضعیت CSRF را از runtimeConfig می‌خوانیم
    const cfg = useRuntimeConfig(event) as { security?: { csrf?: boolean } }
    if (cfg?.security?.csrf !== false) return // اگر فعال بود کاری نکن

    const existing = getCookie(event, 'csrf')
    if (!existing) return

    // حذف کوکی "csrf" برای پاک‌سازی مرورگر در dev / staging
    setCookie(event, 'csrf', '', {
      maxAge: 0,
      path: '/',
      httpOnly: true,
      sameSite: 'lax',
    })
    appendHeader(event, 'set-cookie', '')
  })
})