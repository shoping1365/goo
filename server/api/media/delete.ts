import { createError, defineEventHandler, getQuery } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const query = getQuery(event)
     const id = query.id

     if (!id) {
          throw createError({ statusCode: 400, message: 'شناسه فایل لازم است' })
     }

     try {
          const json = await fetchGo(event, `/api/media/delete?id=${id}`, {
               method: 'DELETE'
          }) as { data?: unknown }
          return { success: true, data: json?.data || json || {} }
     } catch (err: unknown) {
          const error = err as { statusCode?: number; message?: string; statusMessage?: string }
          throw createError({ statusCode: error?.statusCode || 400, message: error?.message || error?.statusMessage || 'خطا در حذف رسانه' })
     }
})
