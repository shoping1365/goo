import { defineEventHandler, getRouterParam, getQuery, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const productId = getRouterParam(event, 'id')
    
    if (!productId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه محصول الزامی است'
      })
    }

    const response = await $fetch(`${process.env.BACKEND_URL}/api/questions/product/${productId}`, {
      method: 'GET'
    })

    return response
  } catch (error) {
    console.error('Error fetching product questions:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت پرسش‌های محصول'
    })
  }
}) 