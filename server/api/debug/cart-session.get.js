import { parseCookies } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const cookies = parseCookies(event)
    
    return {
      success: true,
      cookies: cookies,
      cart_session_id: cookies.cart_session_id || null,
      session_id: cookies.session_id || null,
      mysession: cookies.mysession || null,
      message: 'Session debug info'
    }
    
  } catch (error) {
    console.error('Session debug error:', error)
    return {
      success: false,
      error: error.message,
      message: 'خطا در دریافت اطلاعات session'
    }
  }
}) 