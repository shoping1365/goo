import { useAuthStore } from '@/stores/auth'
import { defineNuxtRouteMiddleware } from 'nuxt/app'
import type { RouteLocationNormalized } from 'vue-router'

// تعریف navigateTo برای Nuxt 3
declare const navigateTo: (to: string) => Promise<void>

export default defineNuxtRouteMiddleware((to: RouteLocationNormalized, _from: RouteLocationNormalized) => {
  const authStore = useAuthStore()

  // بارگذاری داده‌های ذخیره‌شده
  authStore.loadFromStorage()

  // اگر صفحه محافظت‌شده است
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return navigateTo('/auth/login')
  }

  // اگر کاربر وارد شده و به login می‌رود
  if (to.path === '/auth/login' && authStore.isAuthenticated) {
    // بر اساس role redirect کن
    const isAdmin = authStore.user?.role_id === 100 ||
      authStore.user?.username === 'dev' ||
      ['admin', 'developer',].includes(authStore.user?.role?.toLowerCase() || '')
    return navigateTo(isAdmin ? '/admin' : '/account')
  }

  // بررسی دسترسی admin
  if (to.meta.requiresAdmin && authStore.user?.role !== 'admin') {
    return navigateTo('/403')
  }
})
