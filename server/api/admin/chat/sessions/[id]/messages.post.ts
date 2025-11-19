import { defineEventHandler } from 'h3'
import { handleSessionMessages } from './_shared/handleSessionMessages'

export default defineEventHandler((event) =>
  handleSessionMessages(event, {
    method: 'POST',
    errorMessage: 'خطا در ارسال پیام',
    logLabel: 'Error sending message:'
  })
)