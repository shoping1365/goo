import { defineEventHandler } from 'h3'

export default defineEventHandler(() => {
  return {
    status: 'ok',
    message: 'pong',
    timestamp: new Date().toISOString()
  }
}) 