export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          const base = config.public.goApiBase

          const { proxy } = await import('~/server/api/_utils/fetchProxy')
          const id = getRouterParam(event, 'id')

          return proxy(event, `${base}/api/admin/widgets/${id}/toggle`, {
               method: 'POST',
               headers: {
                    'Content-Type': 'application/json'
               }
          })
     } catch (error: any) {
          console.error('خطا در تغییر وضعیت ویجت:', error)
          throw createError({
               statusCode: error.statusCode || 500,
               message: error.message || error.statusMessage || 'خطا در تغییر وضعیت ویجت'
          })
     }
})