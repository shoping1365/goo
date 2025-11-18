import { createError, defineEventHandler, getRouterParam, readBody } from 'h3';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown }) => Promise<T>

export default defineEventHandler(async (event) => {
  try {
    const id = getRouterParam(event, 'id')
    if (!id) {
      throw createError({ statusCode: 400, message: 'شناسه جلسه الزامی است' })
    }
    const body = await readBody(event)
    const config = useRuntimeConfig()
    const url = `${config.public.goApiBase}/api/chat/admin/sessions/${id}`
    const res = await $fetch(url, { method: 'PUT', body })
    return res
  } catch (error) {
    throw createError({ statusCode: 500, message: 'خطا در به‌روزرسانی جلسه' })
  }
}) 