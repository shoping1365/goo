import { defineEventHandler, setCookie, createError } from 'h3'

import { randomUUID } from 'crypto'

export default defineEventHandler(async (event) => {
  try {
    // ایجاد session ID جدید با استفاده از crypto.randomUUID برای امنیت بیشتر
    const sessionId = `cart_${Date.now()}_${randomUUID()}`
    
    // تنظیم cookie با نام session_id برای هماهنگی با بقیه سیستم
    setCookie(event, 'session_id', sessionId, {
      httpOnly: false,
      secure: false,
      sameSite: 'lax',
      maxAge: 60 * 60 * 24 * 30 // 30 روز
    })
    
    
    return {
      success: true,
      session_id: sessionId,
      message: 'Session ID created successfully'
    }
    
  } catch (error) {
    throw createError({
      statusCode: 500,
      message: 'خطا در ایجاد session ID'
    })
  }
})