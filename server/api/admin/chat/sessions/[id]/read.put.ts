import { createError, defineEventHandler, getRouterParam } from 'h3'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string }) => Promise<T>

export default defineEventHandler(async (event) => {
  try {
    const sessionId = getRouterParam(event, 'id')
    if (!sessionId) {
      throw createError({ statusCode: 400, message: 'شناسه جلسه الزامی است' })
    }

    const config = useRuntimeConfig()
    const url = `${config.public.goApiBase}/api/chat/admin/sessions/${sessionId}/read`
    const res = await $fetch(url, { method: 'PUT' })
    return res
  } catch (error) {
    throw createError({ statusCode: 500, message: 'خطا در علامت‌گذاری پیام‌ها' })
  }
})