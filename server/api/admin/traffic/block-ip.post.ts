export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const payload = {
      ip_address: body?.ip || body?.ip_address,
      reason: body?.reason || '',
      duration: body?.duration || 0
    }

    const authHeader = getHeader(event, 'authorization') || `Bearer ${getCookie(event, 'auth-token') || ''}`

    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/traffic/block-ip`, {
      method: 'POST',
      headers: {
        Authorization: authHeader,
        'Content-Type': 'application/json'
      },
      body: payload
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در مسدود کردن IP:', error)
    throw createError({ statusCode: 500, message: 'خطا در مسدود کردن IP' })
  }
})