import { defineEventHandler, getRouterParam, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    
    const { proxy } = await import('~/server/api/_utils/fetchProxy')
    const id = getRouterParam(event, 'id')

    // fetchProxy automatically forwards cookies and creates Authorization header
    return proxy(event, `${base}/api/admin/widgets/${id}`)
  } catch (error: any) {
    console.error('خطا در دریافت ویجت:', error)
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.statusMessage || 'خطا در دریافت ویجت'
    })
  }
})
