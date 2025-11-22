import { createError, defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const { fetchGo } = await import('./_utils/fetchGo')

  const userId = event.context?.user?.id
  if (!userId) {
    throw createError({
      statusCode: 401,
      message: 'کاربر احراز هویت نشده است'
    })
  }

  try {
    // Fetch recent views from Go backend service
    const response = await fetchGo(event, '/api/recent-views', { method: 'GET' })

    const data = Array.isArray((response as { data?: unknown[] })?.data)
      ? (response as { data: unknown[] }).data
      : Array.isArray(response) ? response : []

    return {
      success: true,
      data,
      count: data.length
    }
  } catch (error: unknown) {
    // Handle errors from Go backend or network issues
    console.error('Error fetching recent views:', error)
    throw createError({
      statusCode: error?.statusCode || 500,
      message: error?.statusMessage || 'Error fetching recent views'
    })
  }
})

