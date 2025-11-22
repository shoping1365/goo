import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const baseURL = config.public.goApiBase

  const id = getRouterParam(event, 'id')

  if (!id) {
    throw createError({
      statusCode: 400,
      message: 'شناسه سوال الزامی است'
    })
  }

  const body = await readBody(event)

  try {
    // نگاشت فیلدها برای سازگاری با بک‌اند Go
    const updatePayload: Record<string, unknown> = {}
    if (typeof body.answer !== 'undefined') updatePayload.admin_reply = body.answer
    if (typeof body.status !== 'undefined') updatePayload.status = body.status
    // دسته‌بندی: هم کلید متنی (category) و هم شناسه (category_id) در صورت ارسال پشتیبانی شود
    if (typeof body.category !== 'undefined') updatePayload.category = body.category
    if (typeof body.category_id !== 'undefined') updatePayload.category_id = body.category_id
    if (typeof body.priority !== 'undefined') updatePayload.priority = body.priority
    if (typeof body.question !== 'undefined') updatePayload.question = body.question
    if (typeof body.product_id !== 'undefined') updatePayload.product_id = body.product_id
    if (typeof body.customer_id !== 'undefined') updatePayload.user_id = body.customer_id
    if (typeof body.is_anonymous !== 'undefined') updatePayload.is_anonymous = body.is_anonymous
    if (typeof body.answeredBy !== 'undefined') updatePayload.answered_by = body.answeredBy
    if (typeof body.answeredAt !== 'undefined') updatePayload.admin_reply_at = body.answeredAt

    const response = await $fetch(`${baseURL}/api/questions/${id}`, {
      method: 'PUT',
      body: updatePayload
    })
    return response
  } catch (error) {
    console.error('Error updating question:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در به‌روزرسانی سوال'
    })
  }
})