import { defineEventHandler } from 'h3'

export default defineEventHandler(() => {
  return {
    status: 'ok',
    message: 'API is running',
    timestamp: new Date().toISOString()
  }
}) 