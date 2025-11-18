import { defineEventHandler, getRouterParam, createError } from 'h3'

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
  
  try {
    const response = await $fetch(`${baseURL}/api/questions/${id}`, {
      method: 'DELETE'
    })
    return response
  } catch (error) {
    console.error('Error deleting question:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در حذف سوال'
    })
  }
})