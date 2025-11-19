import { defineEventHandler } from 'h3'
import { handleSessionMessages } from './_shared/handleSessionMessages'

export default defineEventHandler((event) =>
  handleSessionMessages(event, {
    method: 'GET',
    errorMessage: 'خطا در دریافت پیام‌ها',
    logLabel: 'Error fetching messages:'
  })
)