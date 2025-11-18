import { defineEventHandler, getRouterParam, createError, readBody } from 'h3'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase

          const { proxy } = await import('~/server/api/_utils/fetchProxy')
          const id = getRouterParam(event, 'id')
          const body = await readBody(event)

          // fetchProxy automatically forwards cookies and creates Authorization header
          return proxy(event, `${base}/api/admin/widgets/${id}`, {
               method: 'PUT',
               body
          })
     } catch (error: any) {
          console.error('خطا در آپدیت ویجت:', error)
          throw createError({
               statusCode: error.statusCode || 500,
               message: error.statusMessage || 'خطا در آپدیت ویجت'
          })
     }
})
