import { defineEventHandler, getQuery, createError } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const baseURL = config.public.goApiBase
  const query = getQuery(event)
  const { 
    product_id, 
    category, 
    status = 'answered', 
    page = 1, 
    per_page = 10,
    customer_id 
  } = query

  const params = new URLSearchParams()
  if (product_id) params.append('product_id', product_id as string)
  if (category) params.append('category', category as string)
  if (status) params.append('status', status as string)
  if (customer_id) params.append('customer_id', customer_id as string)
  params.append('page', page as string)
  params.append('per_page', per_page as string)
  
  try {
    const response = await $fetch(`${baseURL}/api/questions?${params.toString()}`)
    return response
  } catch (error) {
    console.error('Error fetching questions:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت سوالات'
    })
  }
})