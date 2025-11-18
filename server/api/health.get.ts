import { getHeader } from './_utils/getHeader'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  try {
    const data = await $fetch(`${base}/health`, {
      method: 'GET',
      credentials: 'include',
      headers: {
        cookie: getHeader(event, 'cookie') || ''
      }
    })
    return data
  } catch (err: any) {
    // پروژه از فیلد message استفاده می‌کند
    throw createError({
      statusCode: err?.status || 500,
      message: err?.data?.message || err?.message || 'خطای نامشخص در health'
    })
  }
})