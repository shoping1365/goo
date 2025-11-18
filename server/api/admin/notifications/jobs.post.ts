import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface NotificationBody {
  channel?: string
  message?: string
  recipients?: unknown[]
  [key: string]: unknown
}

export default defineEventHandler(async (event: H3Event) => {
  try {
    const body = await readBody(event) as NotificationBody
    const channel = body?.channel
    const message = body?.message
    const recipients = body?.recipients

    if (!channel || !message || !Array.isArray(recipients) || recipients.length === 0) {
      throw createError({ statusCode: 400, message: 'پارامترهای ورودی نامعتبر است' })
    }

    // شبیه سازی ثبت جاب در سیستم اعلان مرکزی
    const jobId = 'JOB-' + Math.random().toString(36).slice(2, 10).toUpperCase()

    await new Promise((r) => setTimeout(r, 200))

    return { jobId, accepted: true }
  } catch (error: unknown) {
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({ statusCode: errorWithStatus?.statusCode || 500, message: errorWithStatus?.statusMessage || 'خطا در ثبت Job اعلان' })
  }
})


