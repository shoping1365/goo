import { defineNuxtRouteMiddleware, useCookie } from 'nuxt/app'

export default defineNuxtRouteMiddleware((to, from) => {
  // API routes را از middleware مستثنا کن
  if (to.path.startsWith('/api/')) {
    return
  }

  // فقط در client-side اجرا شود
  if (import.meta.client) {
    // بررسی وجود session ID برای سبد خرید
    const cartSessionId = useCookie('session_id')

    if (!cartSessionId.value) {
      // ایجاد session ID جدید
      const sessionId = `cart_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
      cartSessionId.value = sessionId
    }
  }
}) 