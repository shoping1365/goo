import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
  try {
    // دریافت ID از URL
    const id = getRouterParam(event, 'id')

    // ارسال درخواست به backend از طریق fetchGo
    const response = await fetchGo(event, `/api/widgets/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    return response
  } catch (error: unknown) {
    console.error('خطا در حذف ویجت:', error)

    throw createError({
      statusCode: (error as { statusCode?: number }).statusCode || 500,
      message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در حذف ویجت',
      data: (error as { data?: unknown; message?: string }).data || (error as { data?: unknown; message?: string }).message
    })
  }
})