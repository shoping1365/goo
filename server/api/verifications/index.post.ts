import { defineEventHandler, readBody, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

// ایجاد درخواست احراز هویت جدید
// POST /api/verifications
export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)

    const response = await fetchGo(event, '/api/verifications', {
      method: 'POST',
      body: body
    })

    return response
  } catch (error: any) {
    console.error('Error creating verification request:', error)
    throw createError({ statusCode: error.statusCode || 500, message: error.message || 'خطا در ایجاد درخواست احراز هویت' })
  }
})