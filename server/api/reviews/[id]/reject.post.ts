import { defineEventHandler, getRouterParam, createError } from 'h3'
import { useRuntimeConfig } from '#imports'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const baseURL = config.public.goApiBase
  const id = getRouterParam(event, 'id')
  
  try {
    const response = await $fetch(`${baseURL}/api/reviews/${id}/reject`, {
      method: 'POST'
    })
    return response
  } catch (error: unknown) {
    const errorObj = error as { message?: string }
    throw createError({
      statusCode: 500,
      message: errorObj?.message || 'Failed to reject review'
    })
  }
})