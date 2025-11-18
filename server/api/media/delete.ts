import { defineEventHandler, getQuery, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const query = getQuery(event)
     const id = query.id

     if (!id) {
          throw createError({ statusCode: 400, message: 'شناسه فایل لازم است' })
     }

     try {
          const json: any = await fetchGo(event, `/api/media/delete?id=${id}`, {
               method: 'DELETE'
          })
          return { success: true, data: json?.data || json || {} }
     } catch (err: any) {
          throw createError({ statusCode: err?.statusCode || 400, message: err?.message || err?.statusMessage || 'خطا در حذف رسانه' })
     }
})
