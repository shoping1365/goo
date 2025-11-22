export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  
  try {
    const response = await $fetch(`${process.env.BACKEND_URL}/api/roles/${id}/permissions`, {
      headers: {
        'Authorization': `Bearer ${getCookie(event, 'auth-token')}`,
        'Content-Type': 'application/json'
      }
    })
    
    return response
  } catch (error: unknown) {
    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string }).message || 'خطا در دریافت دسترسی‌های نقش'
    })
  }
}) 