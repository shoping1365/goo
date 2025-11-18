// server/api/admin/chat/sessions/[id]/messages.get.ts
import { createError, defineEventHandler, getRouterParam } from 'h3';
import { proxy } from '../../../../_utils/fetchProxy';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event) => {
  try {
    const sessionId = getRouterParam(event, 'id')
    if (!sessionId) {
      throw createError({ statusCode: 400, message: 'شناسه جلسه الزامی است' })
    }
    const config = useRuntimeConfig()
    const url = `${config.public.goApiBase}/api/admin/chat/sessions/${sessionId}/messages`
    const data = await proxy(event, url, { method: 'GET' })
    return data
  } catch (error) {
    console.error('Error fetching messages:', error)
    throw createError({ statusCode: 500, message: 'خطا در دریافت پیام‌ها' })
  }
}) 