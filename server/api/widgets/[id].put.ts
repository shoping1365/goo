import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    // دریافت ID از URL
    const id = getRouterParam(event, 'id')
    
    // دریافت body درخواست
    const body = await readBody(event)

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, `/api/widgets/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body
    })

    return response
  } catch (error: any) {
    console.error('خطا در به‌روزرسانی ویجت:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در به‌روزرسانی ویجت',
      data: error.data || error.message
    })
  }
})