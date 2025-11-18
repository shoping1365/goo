// server/api/admin/chat/sessions.get.ts
import { createError, defineEventHandler, getQuery } from 'h3';
import { proxy } from '../../_utils/fetchProxy';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event) => {
  try {
    // Auth removed - requireAuth
    if (!event?.context?.user) {
      throw createError({ statusCode: 401, message: 'احراز هویت مورد نیاز است' })
    }
    // Permission check removed - hasPermission

    const query = getQuery(event)
    const config = useRuntimeConfig()
    const qs = new URLSearchParams()
    if (query.status) qs.set('status', String(query.status))
    if (query.limit) qs.set('limit', String(query.limit))
    if (query.offset) qs.set('offset', String(query.offset))

    const url = `${config.public.goApiBase}/api/chat/admin/sessions?${qs.toString()}`
    const data = await proxy(event, url, { method: 'GET' })
    return data
  } catch (error) {
    console.error('Error fetching chat sessions:', error)
    throw createError({ statusCode: 500, message: 'خطا در دریافت لیست جلسات چت' })
  }
})


