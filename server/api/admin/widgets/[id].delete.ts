import { defineEventHandler, getRouterParam, createError } from 'h3'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase

          const { proxy } = await import('~/server/api/_utils/fetchProxy')
          const id = getRouterParam(event, 'id')

          // fetchProxy automatically forwards cookies and creates Authorization header
          return proxy(event, `${base}/api/admin/widgets/${id}`, {
               method: 'DELETE'
          })
     } catch (error: unknown) {
          console.error('خطا در حذف ویجت:', error)
          throw createError({
               statusCode: (error as { statusCode?: number }).statusCode || 500,
               message: (error as { message?: string; statusMessage?: string }).message || (error as { message?: string; statusMessage?: string }).statusMessage || 'خطا در حذف ویجت'
          })
     }
})
