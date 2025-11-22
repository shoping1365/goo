import { defineEventHandler, getRouterParam, createError, getQuery } from 'h3'

export default defineEventHandler(async (event) => {
  // دریافت شناسه کاربر از مسیر
  const userIdParam = getRouterParam(event, 'userId')
  if (!userIdParam) {
    throw createError({
      statusCode: 400,
      message: 'شناسه کاربر ارسال نشده است'
    })
  }

  const userId = Number(userIdParam)
  if (!Number.isInteger(userId) || userId <= 0) {
    throw createError({
      statusCode: 400,
      message: 'شناسه کاربر نامعتبر است'
    })
  }

  // دریافت limit از query params
  const query = getQuery(event)
  const limit = query.limit ? Number(query.limit) : 50

  try {
    // ارسال درخواست به سرویس Go برای دریافت بازدیدهای کاربر (فقط برای ادمین)
    const { fetchGo } = await import('../../../_utils/fetchGo')
    return await fetchGo(event, `/api/admin/recent-views/user/${userId}?limit=${limit}`, {
      method: 'GET'
    })
  } catch (error: unknown) {
    console.error('خطا در دریافت بازدیدهای کاربر:', error)
    throw createError({
      statusCode: error?.statusCode || 500,
      message: error?.statusMessage || 'خطا در دریافت بازدیدهای کاربر'
    })
  }
})

