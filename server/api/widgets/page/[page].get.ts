import { getHeader } from '../../_utils/getHeader'

export default defineEventHandler(async (event) => {
  try {
    // دریافت نام صفحه از URL
    const page = getRouterParam(event, 'page')
    console.log('Requested page:', page)

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
    console.log('Backend URL:', url)

    // استفاده از fetch با تایپ مشخص برای جلوگیری از خطای لینتر
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'User-Agent': userAgent
      }
    }).then(res => res.json())

    console.log('Backend response:', response)
    return response
  } catch (error: any) {
    console.error('خطا در دریافت ویجت‌های صفحه:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || 'خطا در دریافت ویجت‌های صفحه',
      data: error.data || error.message
    })
  }
})