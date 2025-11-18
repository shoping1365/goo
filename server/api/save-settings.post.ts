import { defineEventHandler, readBody, createError } from 'h3'
import { writeFile } from 'fs/promises'
import { join } from 'path'

export default defineEventHandler(async (event) => {
  const body = await readBody<{ databaseUrl?: string }>(event)
  if (!body?.databaseUrl) {
    throw createError({ statusCode: 400, message: 'databaseUrl الزامی است' })
  }
  // صرفاً تایید می‌کنیم؛ ذخیره‌سازی در این نسخه لازم نیست
  return { success: true }
}) 