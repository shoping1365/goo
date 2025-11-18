import { getHeader } from './_utils/getHeader'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  try {
    const data = await $fetch(`${base}/api/admin/monitoring/health`, {
      method: 'GET',
      credentials: 'include',
      headers: {
        cookie: getHeader(event, 'cookie') || ''
      }
    })
    return data
  } catch (err: any) {
    throw createError({
      statusCode: err?.status || 500,
      message: err?.data?.message || err?.message || 'خطا در دریافت وضعیت دیتابیس/سیستم'
    })
  }
})