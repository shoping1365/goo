import { getHeader } from '../_utils/getHeader'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  const body = await readBody<{ query?: string }>(event)
  const query = (body?.query || '').trim()
  if (!query) {
    throw createError({ statusCode: 400, message: 'query الزامی است' })
  }
  // محدودیت: فقط SELECT مجاز است
  if (!/^select\s+/i.test(query)) {
    throw createError({ statusCode: 400, message: 'فقط SELECT مجاز است' })
  }

  try {
    const start = Date.now()
    // اگر بک‌اند مسیر اختصاصی نداشت، 501 بدهیم
    const data = await $fetch(`${base}/health/query`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', cookie: getHeader(event, 'cookie') || '' },
      body: { query }
    })

    // ذخیره در لاگ‌های درون‌حافظه‌ای اگر مسیر لاگ وجود داشته باشد
    try {
      const elapsed = Date.now() - start
      // @ts-ignore
      const { logs } = await import('../admin/system/custom-query-logs.get')
      if (Array.isArray((logs as unknown[]))) {
        (logs as unknown[]).push({
          id: Date.now(),
          executed_at: new Date().toISOString(),
          response_time_ms: elapsed,
          query,
          result: JSON.stringify((data as any)?.result ?? data)
        })
      }
    } catch {}

    return data
  } catch (_err: unknown) {
    throw createError({ statusCode: 501, message: 'پشتیبانی در بک‌اند موجود نیست' })
  }
})