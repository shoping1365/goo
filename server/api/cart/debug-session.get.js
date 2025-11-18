import { parseCookies } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const cookies = parseCookies(event)
    
    return {
      cookies,
      session_token: cookies.session_token,
      auth_token: cookies.auth_token,
      cart_id: cookies.cart_id,
      all_cookies: Object.keys(cookies)
    }
  } catch (error) {
    console.error('خطا در بررسی session:', error)
    return {
      error: error.message,
      cookies: {}
    }
  }
}) 