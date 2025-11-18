import { defineEventHandler, getRouterParam, createError, getHeader } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const ip = getRouterParam(event, 'ip')
    const response = await $fetch(`${process.env.BACKEND_URL}/api/admin/traffic/unblock-ip/${ip}`, {
      method: 'DELETE',
      headers: { 'Authorization': getHeader(event, 'authorization') || '' }
    })
    return response
  } catch (error: any) {
    throw createError({ statusCode: 500, message: 'خطا در آزاد کردن IP' })
  }
})