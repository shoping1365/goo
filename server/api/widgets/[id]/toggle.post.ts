import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    // دریافت ID از URL
    const id = getRouterParam(event, 'id')

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, `/api/widgets/${id}/toggle`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    return response
  } catch (error: any) {
    console.error('خطا در تغییر وضعیت ویجت:', error)
    
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'خطا در تغییر وضعیت ویجت',
      data: error.data || error.message
    })
  }
})