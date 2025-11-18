import { defineEventHandler, getQuery, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const config = useRuntimeConfig()
    const apiBase = config.public.goApiBase
    
    // دریافت query parameters
    const query = getQuery(event)
    const page = query.page as string | undefined
    const status = query.status as string | undefined

    // ساخت URL برای درخواست به backend
    let url = `${apiBase}/api/widgets`
    const params = new URLSearchParams()

    if (page) params.append('page', page)
    if (status) params.append('status', status)

    if (params.toString()) {
      url += `?${params.toString()}`
    }

    // ارسال درخواست به backend
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(res => res.json())

    return response
  } catch (error: any) {
    console.error('خطا در دریافت ویجت‌ها:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در دریافت ویجت‌ها',
      data: error.data || error.message
    })
  }
})