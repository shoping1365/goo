import { defineEventHandler, getQuery, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

// این هندلر، لیست پرسش‌های محصولات را برای پنل ادمین از بک‌اند Go واکشی می‌کند
// ورودی‌های فیلتر (مانند وضعیت و جستجو) به همان صورت به سرویس Go پاس داده می‌شوند
export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    const params = new URLSearchParams()

    // فیلترهای اختیاری
    if (query.status) params.set('status', String(query.status))
    if (query.search) params.set('search', String(query.search))
    if (query.page) params.set('page', String(query.page))
    if (query.per_page) params.set('per_page', String(query.per_page))

    const urlPath = `/api/questions/admin${params.toString() ? `?${params.toString()}` : ''}`
    const response = await fetchGo(event, urlPath)
    return response
  } catch (error) {
    console.error('Error fetching admin product questions:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت پرسش‌های ادمین'
    })
  }
})


