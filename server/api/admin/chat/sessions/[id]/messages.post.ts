// server/api/chat/admin/sessions/[id]/messages.post.ts
import type { H3Event } from 'h3'
import { defineEventHandler, createError, getRouterParam } from 'h3'
import { proxy } from '../../../../_utils/fetchProxy';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event) => {
  try {
    const sessionId = getRouterParam(event, 'id')
    if (!sessionId) {
      throw createError({ statusCode: 400, message: 'شناسه جلسه الزامی است' })
    }
    const config = useRuntimeConfig()
    const url = `${config.public.goApiBase}/api/chat/admin/sessions/${sessionId}/messages`
    // Use proxy to forward multipart body
    const res = await proxy(event, url, { method: 'POST' })
    return res
  } catch (error) {
    console.error('Error sending message:', error)
    throw createError({ statusCode: 500, message: 'خطا در ارسال پیام' })
  }
}) 