import { defineNuxtRouteMiddleware } from 'nuxt/app'

export default defineNuxtRouteMiddleware((to, _from) => {
  // API routes را از middleware مستثنا کن
  if (to.path.startsWith('/api/')) {
    return
  }

  // غیرفعال کردن موقت traffic monitoring
  return
}) 