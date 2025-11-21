// Plugin برای همگام‌سازی client-side
import { useAuthState } from '../composables/useAuthState'

// @ts-ignore
export default defineNuxtPlugin(async (_nuxtApp) => {
  // فقط در client-side اجرا می‌شود
  if (process.server) return

  const { fetchUser, isAuthenticated, clearAuthState } = useAuthState()

  // همگام‌سازی در تب‌های مختلف مرورگر
  if (typeof window !== 'undefined') {
    window.addEventListener('storage', async (e) => {
      // اگر تغییری در auth state رخ داده، اطلاعات را مجدداً بارگذاری کن
      if (e.key === 'auth-sync') {
        // ✅ همیشه fetch کن - کوکی httpOnly است
        try {
          await fetchUser(true) // force refresh
        } catch (error) {
          // اگر خطا رخ داد، state را پاک کن
          console.error('❌ Client Plugin: Error during sync:', error)
          clearAuthState()
        }
      }
    })

    // Heartbeat برای چک کردن وضعیت احراز هویت هر 5 دقیقه
    let heartbeatInterval: NodeJS.Timeout | null = null

    const startHeartbeat = () => {
      if (heartbeatInterval) return

      heartbeatInterval = setInterval(async () => {
        if (isAuthenticated.value) {
          try {
            await fetchUser(true)
          } catch (error) {
            console.error('❌ Client Plugin: Heartbeat failed:', error)
            clearAuthState()
          }
        }
      }, 5 * 60 * 1000) // 5 دقیقه
    }

    const stopHeartbeat = () => {
      if (heartbeatInterval) {
        clearInterval(heartbeatInterval)
        heartbeatInterval = null
      }
    }

    // شروع heartbeat اگر کاربر احراز شده است
    if (isAuthenticated.value) {
      startHeartbeat()
    }

    // متوقف کردن heartbeat هنگام unload
    window.addEventListener('beforeunload', stopHeartbeat)
  }
})
