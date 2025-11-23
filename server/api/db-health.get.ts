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
  } catch (err: unknown) {
    const error = err as { status?: number; data?: { message?: string }; message?: string }
    throw createError({
      statusCode: error?.status || 500,
      message: error?.data?.message || error?.message || 'خطا در دریافت وضعیت دیتابیس/سیستم'
    })
  }
})