import { defineEventHandler, getRouterParam, readBody, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const id = getRouterParam(event, 'id')
     if (!id) {
          throw createError({ statusCode: 400, message: 'شناسه فایل لازم است' })
     }

     const body = await readBody(event)
     try {
          const json: any = await fetchGo(event, `/api/media/${id}`, {
               method: 'PUT',
               body
          })
          return { success: true, data: json?.data || json || {} }
     } catch (err: any) {
          throw createError({ statusCode: err?.statusCode || 400, message: err?.message || err?.statusMessage || 'خطا در بروزرسانی رسانه' })
     }
})


