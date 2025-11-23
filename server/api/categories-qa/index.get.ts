import { createError, defineEventHandler } from 'h3'

export default defineEventHandler(async (_event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  
  try {
    const response = await $fetch(`${base}/api/categories-qa`)
    return response
  } catch (error) {
    console.error('Error fetching QA categories:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت دسته‌بندی‌های سوالات'
    })
  }
})