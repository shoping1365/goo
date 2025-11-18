import { useAuthStore } from '@/stores/auth'

declare const defineNuxtPlugin: (plugin: (nuxtApp: any) => void | Promise<void>) => { provide?: Record<string, any> }

export default defineNuxtPlugin(() => {
  if (process.client) {
    // بارگذاری auth state از localStorage
    const authStore = useAuthStore()
    authStore.loadFromStorage()

    // Setup automatic token refresh
    const tokenRefreshInterval = setInterval(async () => {
      if (authStore.isAuthenticated) {
        try {
          // فراخوانی refresh token endpoint
          // این نباید خطا دهد اگر token معتبر نیست
        } catch (err) {
          console.error('Token refresh failed:', err)
        }
      }
    }, 14 * 60 * 1000) // هر 14 دقیقه (توکن 15 دقیقه ای است)

    // Cleanup on page unload
    if (process.client) {
      window.addEventListener('beforeunload', () => {
        clearInterval(tokenRefreshInterval)
      })
    }
  }
})
