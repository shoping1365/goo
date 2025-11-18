import { defineEventHandler, readBody, createError } from 'h3'
import { fetchGo } from '../../../_utils/fetchGo'

// انتقال گروهی پرسش‌ها به یک دسته‌بندی دیگر در بک‌اند Go
export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const questionIds = Array.isArray(body?.question_ids) ? body.question_ids : []
  const toCategory = body?.to_category

  if (!questionIds.length || !toCategory) {
    throw createError({ statusCode: 400, message: 'لیست شناسه‌ها و دسته مقصد الزامی است' })
  }

  try {
    // 1) دریافت لیست دسته‌بندی‌ها برای نگاشت کلید به شناسه (در صورت نیاز)
    const categories: Array<{ id: number; key: string; name: string }> = await fetchGo(event, '/api/categories-qa')
    const target = categories.find(c => c.key === String(toCategory))
    if (!target) {
      throw createError({ statusCode: 404, message: 'دسته‌بندی مقصد یافت نشد' })
    }

    // 2) به‌روزرسانی تک‌تک سوال‌ها با category_id مقصد
    // از مسیر موجود PUT /api/questions/:id استفاده می‌کنیم
    const results: Array<{ id: string | number; ok: boolean }> = []
    for (const qid of questionIds) {
      try {
        await fetchGo(event, `/api/questions/${encodeURIComponent(qid)}`, {
          method: 'PUT',
          body: { category_id: target.id }
        })
        results.push({ id: qid, ok: true })
      } catch (err) {
        results.push({ id: qid, ok: false })
      }
    }

    const transferred = results.filter(r => r.ok).length
    return { transferred_count: transferred, requested: questionIds.length, to_category: toCategory }
  } catch (error) {
    console.error('Error bulk transferring QA questions:', error)
    throw createError({ statusCode: 500, message: 'خطا در انتقال گروهی پرسش‌ها' })
  }
})


