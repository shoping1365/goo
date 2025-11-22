import { defineEventHandler, readBody, getRouterParam, createError } from 'h3'

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
  
  const body = await readBody(event)
  
  try {
    const response = await fetch(`${base}/admin/sms-patterns/${id}`, {
      method: 'PUT',
      headers: { 'Accept': 'application/json', 'Content-Type': 'application/json' },
      body: JSON.stringify(body)
    })
    const json = await response.json()
    return json
  } catch (_error) {
    throw createError({
      statusCode: 500,
      message: 'خطا در ویرایش پترن'
    })
  }
})
