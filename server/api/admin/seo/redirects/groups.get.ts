import { defineEventHandler, createError } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

export default defineEventHandler(async (event) => {
  try {
    // درخواست به Go backend
    const response = await goApiFetch(event, '/api/admin/seo/redirects?groups=true', {
      method: 'GET'
    })

    return response
  } catch (error: any) {
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت دسته‌بندی‌ها: ' + error.message
    })
  }
})