export default defineEventHandler(async (event) => {
  const query = getQuery(event)
  const { status, rating, search, page = 1, per_page = 10 } = query

  const params = new URLSearchParams()
  if (status) params.append('status', status as string)
  if (rating) params.append('rating', rating as string)
  if (search) params.append('search', search as string)
  params.append('page', page as string)
  params.append('per_page', per_page as string)

  const config = useRuntimeConfig()
  const baseURL = config.public.goApiBase
  
  try {
    const response = await $fetch(`${baseURL}/api/reviews?${params.toString()}`)
    return response
  } catch (error) {
    throw createError({
      statusCode: 500,
      message: 'Failed to fetch reviews'
    })
  }
})