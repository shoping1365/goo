import fs from 'fs';
import type { H3Event } from 'h3';
import path from 'path';

declare const defineEventHandler: (handler: (event: H3Event) => unknown | Promise<unknown>) => unknown
declare const createError: (options: { statusCode: number; message: string }) => Error
declare const setResponseHeader: (event: H3Event, name: string, value: string) => void

export default defineEventHandler(async (event) => {
  try {
    // بررسی دسترسی ادمین
    const user = event.context.user
    if (!user || user.role !== 'admin') {
      throw createError({
        statusCode: 403,
        message: 'دسترسی غیرمجاز'
      })
    }

    // مسیر فایل‌های build شده Nuxt
    const outputDir = path.join(process.cwd(), '.output', 'public')
    const nuxtDir = path.join(outputDir, '_nuxt')

    let clearedCount = 0

    // پاک کردن کش فایل‌های _nuxt (JS/CSS بیلد شده)
    if (fs.existsSync(nuxtDir)) {
      const files = fs.readdirSync(nuxtDir)
      for (const file of files) {
        const filePath = path.join(nuxtDir, file)
        if (fs.statSync(filePath).isFile()) {
          fs.unlinkSync(filePath)
          clearedCount++
        }
      }
    }

    // ارسال header برای پاک کردن کش مرورگر
    setResponseHeader(event, 'Clear-Site-Data', '"cache", "storage"')

    return {
      success: true,
      message: `کش با موفقیت پاک شد (${clearedCount} فایل)`,
      clearedFiles: clearedCount,
      timestamp: new Date().toISOString()
    }
  } catch (error: unknown) {
    console.error('خطا در پاک کردن کش:', error)
    const errorWithMessage = error as { message?: string }
    throw createError({
      statusCode: 500,
      message: errorWithMessage?.message || 'خطا در پاک کردن کش'
    })
  }
})
