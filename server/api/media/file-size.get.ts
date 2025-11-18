import { promises as fs } from 'node:fs'
import path from 'node:path'

export default defineEventHandler(async (event) => {
  try {
    const query = getQuery(event)
    const imageUrl = query.url as string
    
    if (!imageUrl) {
      throw createError({
        statusCode: 400,
        message: 'URL is required'
      })
    }
    
    // تبدیل URL به مسیر فایل
    let filePath = imageUrl
    if (filePath.startsWith('/uploads/')) {
      filePath = path.join('public', filePath)
    } else if (filePath.startsWith('/')) {
      filePath = path.join('public', filePath)
    } else {
      filePath = path.join('public', filePath)
    }
    
    try {
      const stats = await fs.stat(filePath)
      return {
        success: true,
        size: stats.size,
        exists: true
      }
    } catch (error) {
      return {
        success: true,
        size: 0,
        exists: false
      }
    }
  } catch (error) {
    console.error('Error getting file size:', error)
    return {
      success: false,
      error: 'خطا در دریافت سایز فایل',
      size: 0
    }
  }
})
