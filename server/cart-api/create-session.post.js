import { defineEventHandler, setCookie, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    // ایجاد session ID جدید
    const sessionId = `cart_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
    
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