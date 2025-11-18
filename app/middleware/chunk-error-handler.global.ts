import { defineNuxtRouteMiddleware } from 'nuxt/app'

// Global middleware to handle chunk loading errors
export default defineNuxtRouteMiddleware((to, from) => {
     // Only run on client side
     if (import.meta.client) {
          // Listen for chunk loading errors
          window.addEventListener('error', (event) => {
               const error = event.error

               if (
                    error?.message?.includes('Failed to fetch dynamically imported module') ||
                    error?.message?.includes('Loading chunk') ||
                    error?.message?.includes('Loading CSS chunk') ||
                    error?.name === 'ChunkLoadError'
               ) {
                    console.warn('🔄 Chunk loading error detected, reloading page...')

                    // Prevent default error handling
                    event.preventDefault()

                    // Reload the page
                    setTimeout(() => {
                         window.location.reload()
                    }, 500)
               }
          })
     }
})

