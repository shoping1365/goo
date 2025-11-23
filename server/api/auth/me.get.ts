import { defineEventHandler } from 'h3'

/**
 * GET /api/auth/me
 * دریافت اطلاعات کاربر جاری از Go backend
 */
export default defineEventHandler(async (event) => {
  try {
    const { proxy } = await import('~/server/api/_utils/fetchProxy')
    const config = useRuntimeConfig()
    const targetUrl = `${config.goApiBaseUrl}/api/auth/me`
    
    // استفاده از proxy برای forward کردن کوکی‌ها به درستی
    const response = await proxy(event, targetUrl)

    return response

  } catch (error: unknown) {
    const err = error as { statusCode?: number; status?: number; message?: string }
    console.error('❌ Error in auth/me:', error)
    console.error('❌ Error status:', err?.statusCode, err?.status)
    console.error('❌ Error message:', err?.message)

    // اگر 401 unauthorized بود، user authenticated نیست (حالت طبیعی)
    if (err?.statusCode === 401 || err?.status === 401) {
      return {
        authenticated: false,
        user: null,
        role: null
      }
    }

    // خطاهای دیگر را throw کن
    throw error
  }
})
