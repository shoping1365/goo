export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')
    const body = await readBody(event)

    // Call backend API
    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/users/${id}/block`, {
      method: 'POST',
      body,
      headers: {
        'Authorization': getHeader(event, 'authorization') || ''
      }
    })

    return response
  } catch (error) {
    console.error('Error blocking user:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در بلاک کردن کاربر'
    })
  }
}) 