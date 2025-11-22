import { defineEventHandler, getRouterParam, createError } from 'h3'

export default defineEventHandler(async (event) => {
  await requireAuth(event)
  
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const id = getRouterParam(event, 'id')
  
  if (!id) {
    throw createError({
      statusCode: 400,
      message: 'شناسه پترن الزامی است'
    })
  }
  
  try {
    const response = await fetch(`${base}/admin/sms-patterns/${id}`, {
      method: 'DELETE',
      headers: { 'Accept': 'application/json' }
    })
    const json = await response.json()
    return json
  } catch (_error) {
    throw createError({
      statusCode: 500,
      message: 'خطا در حذف پترن'
    })
  }
})
