import { getHeader } from '../../_utils/getHeader'

export default defineEventHandler(async (event) => {
  try {
    // دریافت نام صفحه از URL
    const page = getRouterParam(event, 'page')

    // تنظیمات rate limiting را برای ویجت‌ها bypass کنیم
    event.context.security = event.context.security || {}
    event.context.security.rateLimiter = false

    // دریافت API base از runtime config
    const config = useRuntimeConfig()
    const apiBase = config.public.goApiBase

    // دریافت User-Agent از request و ارسال به backend
    const userAgent = getHeader(event, 'user-agent') || ''

    // ارسال درخواست به backend
    const url = `${apiBase}/api/widgets/page/${page}`

    // استفاده از fetch با تایپ مشخص برای جلوگیری از خطای لینتر
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'User-Agent': userAgent
      }
    }).then(res => res.json())

    return response
  } catch (error: unknown) {
    console.error('خطا در دریافت ویجت‌های صفحه:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string }).message || 'خطا در دریافت ویجت‌های صفحه',
      data: (error as { data?: unknown; message?: string }).data || (error as { data?: unknown; message?: string }).message
    })
  }
})