export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const baseURL = config.public.goApiBase
  
  const id = getRouterParam(event, 'id')
  
  if (!id) {
    throw createError({
      statusCode: 400,
      message: 'Review ID is required'
    })
  }

  const body = await readBody(event)
  
  if (!body.reply) {
    throw createError({
      statusCode: 400,
      message: 'Reply text is required'
    })
  }
  
  try {
    const response = await $fetch(`${baseURL}/api/reviews/${id}/reply`, {
      method: 'POST',
      body: {
        reply: body.reply
      }
    })
    return response
  } catch (_error) {
    throw createError({
      statusCode: 500,
      message: 'Failed to add reply to review'
    })
  }
})