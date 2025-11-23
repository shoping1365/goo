import { createError, defineEventHandler, getRouterParam, readBody } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const id = getRouterParam(event, 'id')
     if (!id) {
          throw createError({ statusCode: 400, message: 'شناسه فایل لازم است' })
     }

     const body = await readBody(event)
     try {
          const json = await fetchGo(event, `/api/media/${id}`, {
               method: 'PUT',
               body
          }) as { data?: unknown }
          return { success: true, data: json?.data || json || {} }
     } catch (err: unknown) {
          const error = err as { statusCode?: number; message?: string; statusMessage?: string }
          throw createError({ statusCode: error?.statusCode || 400, message: error?.message || error?.statusMessage || 'خطا در بروزرسانی رسانه' })
     }
})


