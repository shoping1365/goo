import { defineEventHandler } from 'h3'
import { getCookieValue } from '../_utils/cookieHelper'

/**
 * GET /api/auth/me
 * دریافت اطلاعات کاربر جاری از Go backend
 */
export default defineEventHandler(async (event) => {
  try {
    const { proxy } = await import('~/server/api/_utils/fetchProxy')
    const config = useRuntimeConfig()
    const targetUrl = `${config.goApiBaseUrl}/api/auth/me`
    
    // Debug: بررسی کوکی‌های موجود
    const accessToken = getCookieValue(event, 'access_token')
    const authToken = getCookieValue(event, 'auth-token')
    const allCookies = event.node?.req?.headers?.cookie
    
    // استفاده از proxy برای forward کردن کوکی‌ها به درستی
    const response = await proxy(event, targetUrl)

    return response

  } catch (error: unknown) {
    console.error('❌ Error in auth/me:', error)
    console.error('❌ Error status:', error?.statusCode, error?.status)
    console.error('❌ Error message:', error?.message)

    // اگر 401 unauthorized بود، user authenticated نیست (حالت طبیعی)
    if (error?.statusCode === 401 || error?.status === 401) {
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
