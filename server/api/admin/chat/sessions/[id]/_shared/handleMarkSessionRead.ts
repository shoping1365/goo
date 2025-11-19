import { H3Event, createError, getRouterParam } from 'h3'
import { fetchGo } from '../../../../../_utils/fetchGo'

export async function handleMarkSessionRead(event: H3Event) {
  const sessionId = getRouterParam(event, 'id')
  if (!sessionId) {
    throw createError({ statusCode: 400, message: 'شناسه جلسه الزامی است' })
  }

  try {
    return await fetchGo(event, `/api/chat/admin/sessions/${sessionId}/read`, { method: 'PUT' })
  } catch (error) {
    console.error('Error marking messages as read:', error)
    throw createError({ statusCode: 500, message: 'خطا در علامت‌گذاری پیام‌ها' })
  }
}
