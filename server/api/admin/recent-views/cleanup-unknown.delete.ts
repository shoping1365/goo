import { defineEventHandler, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // ارسال درخواست به سرویس Go برای پاک کردن بازدیدهای Unknown
    const { fetchGo } = await import('../../_utils/fetchGo')
    return await fetchGo(event, '/api/admin/recent-views/cleanup-unknown', {
      method: 'DELETE'
    })
  } catch (error: any) {
    console.error('خطا در پاک کردن بازدیدهای قدیمی:', error)
    throw createError({
      statusCode: error?.statusCode || 500,
      message: error?.statusMessage || 'خطا در پاک کردن بازدیدهای قدیمی'
    })
  }
})

