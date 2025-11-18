import { defineEventHandler, getRouterParam, createError } from 'h3'

// حذف یک دسته‌بندی پرسش و پاسخ بر اساس کلید
export default defineEventHandler(async (event) => {
  const key = getRouterParam(event, 'key')
  if (!key) {
    throw createError({ statusCode: 400, message: 'کلید دسته‌بندی الزامی است' })
  }
  try {
    const config = useRuntimeConfig()
    const baseURL = config.public.goApiBase
    const response = await $fetch(`${baseURL}/api/categories-qa/${encodeURIComponent(key)}`, {
      method: 'DELETE'
    })
    return response
  } catch (error) {
    console.error('Error deleting QA category:', error)
    throw createError({ statusCode: 500, message: 'خطا در حذف دسته‌بندی' })
  }
})