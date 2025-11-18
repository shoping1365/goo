import { defineEventHandler, readBody, createError } from 'h3'

// ایجاد دسته‌بندی جدید پرسش و پاسخ در بک‌اند Go
export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  if (!body?.key || !body?.name) {
    throw createError({ statusCode: 400, message: 'کلید و نام دسته‌بندی الزامی است' })
  }

  try {
    const baseURL = getGoApiBaseUrl()
    const response = await $fetch(`${baseURL}/api/categories-qa`, {
      method: 'POST',
      body: { key: body.key, name: body.name }
    })
    return response
  } catch (error) {
    console.error('Error creating QA category:', error)
    throw createError({ statusCode: 500, message: 'خطا در ایجاد دسته‌بندی' })
  }
})