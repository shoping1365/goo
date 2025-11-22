import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    // دریافت body درخواست
    const body = await readBody(event)

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, '/api/widgets', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در ایجاد ویجت:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string }).message || 'خطا در ایجاد ویجت',
      data: (error as { data?: unknown }).data
    })
  }
}) 