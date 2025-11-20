import { defineNuxtRouteMiddleware, useCookie } from 'nuxt/app'

export default defineNuxtRouteMiddleware((to, _from) => {
  // API routes را از middleware مستثنا کن
  if (to.path.startsWith('/api/')) {
    return
  }

  // فقط در client-side اجرا شود
  if (import.meta.client) {
    // بررسی وجود session ID برای سبد خرید
    const cartSessionId = useCookie('session_id')

    if (!cartSessionId.value) {
      // ایجاد session ID جدید با استفاده از crypto.getRandomValues برای امنیت بیشتر
      const array = new Uint8Array(16)
      crypto.getRandomValues(array)
      const randomPart = Array.from(array, byte => byte.toString(36)).join('').substring(0, 16)
      const sessionId = `cart_${Date.now()}_${randomPart}`
      cartSessionId.value = sessionId
    }
  }
}) 