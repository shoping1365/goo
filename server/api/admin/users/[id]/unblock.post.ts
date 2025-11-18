export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')

    // Call backend API
    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/users/${id}/unblock`, {
      method: 'POST',
      headers: {
        'Authorization': getHeader(event, 'authorization') || ''
      }
    })

    return response
  } catch (error) {
    console.error('Error unblocking user:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در آنبلاک کردن کاربر'
    })
  }
}) 