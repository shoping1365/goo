import { defineEventHandler, getRouterParam, createError } from 'h3'
import { getCookieValue } from '../../_utils/cookieHelper'

export default defineEventHandler(async (event) => {
  // دریافت شناسه محصول از مسیر
  const productIdParam = getRouterParam(event, 'productId')
  if (!productIdParam) {
    throw createError({
      statusCode: 400,
      message: 'شناسه محصول ارسال نشده است'
    })
  }

  const productId = Number(productIdParam)
  if (!Number.isInteger(productId) || productId <= 0) {
    throw createError({
      statusCode: 400,
      message: 'شناسه محصول نامعتبر است'
    })
  }

  try {
    // در صفحات عمومی: از اجرای احراز هویت سنگین جلوگیری می‌کنیم
    // فقط اگر کوکی احراز هویت وجود داشته باشد درخواست را به بک‌اند فوروارد می‌کنیم
    const hasAccess = !!(getCookieValue(event, 'access_token') || getCookieValue(event, 'auth-token') || getCookieValue(event, 'refresh_token'))

    if (!hasAccess) {
      // برای کاربران مهمان: لاگ سبک و موفقیت ظاهری
      // استفاده از console.debug برای جلوگیری از نویز لاگ‌ها
      return { ok: true, guest: true }
    }

    // اگر کوکی وجود دارد، بدون verify سمت نود، درخواست را به سرویس Go ارسال کن
    const { fetchGo } = await import('../../_utils/fetchGo')
    try {
      return await fetchGo(event, `/api/recent-views/product/${productId}`, {
        method: 'POST'
      })
    } catch (err: unknown) {
      // اگر توکن نامعتبر/منقضی باشد، به عنوان مهمان ادامه می‌دهیم و خطا پرتاب نمی‌کنیم
      const status = err?.statusCode || err?.response?.status
      if (status === 401 || status === 403) {
        return { ok: true, guest: true }
      }
      throw err
    }
  } catch (error: unknown) {
    console.error('خطا در ثبت بازدید محصول:', error)
    throw createError({
      statusCode: error?.statusCode || 500,
      message: error?.statusMessage || 'خطا در ثبت بازدید محصول'
    })
  }
})