import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'

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

  // دریافت مدت زمان از body
  const body = await readBody(event)
  const duration = Number(body?.duration)

  if (!duration || duration < 1 || duration > 3600) {
    throw createError({
      statusCode: 400,
      message: 'مدت زمان نامعتبر است (باید بین 1 تا 3600 ثانیه باشد)'
    })
  }

  // بررسی احراز هویت کاربر
  const userId = event.context?.user?.id
  if (!userId) {
    throw createError({
      statusCode: 401,
      message: 'کاربر احراز هویت نشده است'
    })
  }

  try {
    // ارسال درخواست به سرویس Go برای بروزرسانی مدت زمان
    const { fetchGo } = await import('../../../_utils/fetchGo')
    return await fetchGo(event, `/api/recent-views/product/${productId}/duration`, {
      method: 'PATCH',
      body: { duration }
    })
  } catch (error: unknown) {
    console.error('خطا در بروزرسانی مدت زمان بازدید:', error)
    const err = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: err.statusCode || 500,
      message: err.statusMessage || 'خطا در بروزرسانی مدت زمان بازدید'
    })
  }
})

