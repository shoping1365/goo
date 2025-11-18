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
  } catch (error: any) {
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || 'خطا در دریافت دسترسی‌های نقش'
    })
  }
}) 