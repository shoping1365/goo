import { defineEventHandler, getRouterParam, readBody, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

// به‌روزرسانی اطلاعات پایه کاربر (نام، ایمیل، موبایل یا profile_data)
export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  if (!id) {
    throw createError({ statusCode: 400, message: 'شناسه کاربر الزامی است' })
  }
  const body = await readBody(event)
  try {
    const payload: Record<string, unknown> = {}
    if (typeof body.name !== 'undefined') payload.name = body.name
    if (typeof body.email !== 'undefined') payload.email = body.email
    if (typeof body.mobile !== 'undefined') payload.mobile = body.mobile
    if (typeof body.profile_data !== 'undefined') payload.profile_data = body.profile_data

    const res = await fetchGo(event, `/api/users/${encodeURIComponent(id)}`, {
      method: 'PUT',
      body: payload
    })
    return res
  } catch (_error) {
    console.error('Error updating user:', error)
    throw createError({ statusCode: 500, message: 'خطا در بروزرسانی کاربر' })
  }
})