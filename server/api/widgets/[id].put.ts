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
  } catch (error: unknown) {
    console.error('خطا در به‌روزرسانی ویجت:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در به‌روزرسانی ویجت',
      data: (error as { data?: unknown; message?: string }).data || (error as { data?: unknown; message?: string }).message
    })
  }
})