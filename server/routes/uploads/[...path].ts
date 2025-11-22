import { defineEventHandler, sendStream } from 'h3'
import { createReadStream, existsSync } from 'fs'
import { resolve } from 'path'
import { stat } from 'fs/promises'

export default defineEventHandler(async (event) => {
  // دریافت مسیر فایل از URL
  const path = event.context.params?.path || ''
  
  // مسیر کامل به فایل در پوشه public/uploads
  const filePath = resolve(process.cwd(), 'public', 'uploads', path)
  
  // بررسی اینکه فایل وجود دارد یا خیر
  if (!existsSync(filePath)) {
    throw createError({
      statusCode: 404,
      statusMessage: 'File not found'
    })
  }

  // بررسی اینکه مسیر یک فایل است نه یک پوشه
  const fileStat = await stat(filePath)
  if (!fileStat.isFile()) {
    throw createError({
      statusCode: 400,
      statusMessage: 'Invalid file path'
    })
  }

  // تعیین Content-Type بر اساس پسوند فایل
  const ext = path.split('.').pop()?.toLowerCase()
  const mimeTypes: Record<string, string> = {
    'jpg': 'image/jpeg',
    'jpeg': 'image/jpeg',
    'png': 'image/png',
    'gif': 'image/gif',
    'webp': 'image/webp',
    'svg': 'image/svg+xml',
    'pdf': 'application/pdf',
    'mp4': 'video/mp4',
    'webm': 'video/webm'
  }

  const contentType = mimeTypes[ext || ''] || 'application/octet-stream'
  
  // تنظیم هدرهای مناسب
  event.node.res.setHeader('Content-Type', contentType)
  event.node.res.setHeader('Cache-Control', 'public, max-age=31536000, immutable')
  
  // ارسال فایل به عنوان stream
  return sendStream(event, createReadStream(filePath))
})
