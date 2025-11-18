import type { H3Event } from 'h3'
import { exec } from 'child_process'
import { promisify } from 'util'
import { readFile } from 'fs/promises'
import { defineEventHandler, createError } from 'h3'

const execAsync = promisify(exec)

declare const getServerSession: (event: H3Event) => Promise<{ user?: { role?: string } } | null>

export default defineEventHandler(async (event) => {
  try {
    // بررسی احراز هویت ادمین
    const session = await getServerSession(event)
    if (!session?.user?.role || session.user.role !== 'admin') {
      throw createError({
        statusCode: 403,
        message: 'دسترسی غیرمجاز'
      })
    }

    // دریافت لاگ‌های سیستم
    let logs = ''
    
    try {
      // دریافت اطلاعات سیستم
      const { stdout: systemInfo } = await execAsync('systeminfo | findstr /C:"Total Physical Memory" /C:"Available Physical Memory" /C:"Processor"')
      logs += `=== اطلاعات سیستم ===\n${systemInfo}\n\n`
    } catch (error) {
      logs += '=== اطلاعات سیستم ===\nخطا در دریافت اطلاعات سیستم\n\n'
    }

    try {
      // دریافت لاگ‌های Nuxt
      const nuxtLogs = await readFile('./logs/combined.log', 'utf8').catch(() => '')
      if (nuxtLogs) {
        logs += `=== لاگ‌های Nuxt (آخرین 20 خط) ===\n${nuxtLogs.split('\n').slice(-20).join('\n')}\n\n`
      }
    } catch (error) {
      logs += '=== لاگ‌های Nuxt ===\nفایل لاگ یافت نشد\n\n'
    }

    try {
      // دریافت لاگ‌های خطا
      const errorLogs = await readFile('./logs/error.log', 'utf8').catch(() => '')
      if (errorLogs) {
        logs += `=== لاگ‌های خطا (آخرین 10 خط) ===\n${errorLogs.split('\n').slice(-10).join('\n')}\n\n`
      }
    } catch (error) {
      logs += '=== لاگ‌های خطا ===\nفایل لاگ خطا یافت نشد\n\n'
    }

    // اضافه کردن timestamp
    logs += `\n=== آخرین به‌روزرسانی ===\n${new Date().toLocaleString('fa-IR')}`

    return {
      success: true,
      logs: logs || 'هیچ لاگی یافت نشد',
      timestamp: new Date().toISOString()
    }
  } catch (error: unknown) {
    console.error('خطا در دریافت لاگ‌های سیستم:', error)
    const errorObj = error as { message?: string }
    
    return {
      success: false,
      logs: `خطا در دریافت لاگ‌های سیستم: ${errorObj?.message || 'Unknown error'}`,
      error: errorObj?.message || 'Unknown error'
    }
  }
}) 