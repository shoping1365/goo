// Global error handler for unhandled promise rejections and errors
// @ts-ignore
export default defineNuxtPlugin(() => {
     // Only run on client side
     if (import.meta.client) {
          // Handle unhandled promise rejections
          window.addEventListener('unhandledrejection', (event) => {
               const error = event.reason

               // Check for dynamic import failures
               if (
                    error?.message?.includes('Failed to fetch dynamically imported module') ||
                    error?.message?.includes('Loading chunk') ||
                    error?.message?.includes('Loading CSS chunk') ||
                    error?.name === 'ChunkLoadError'
               ) {
                    console.warn('🔄 Dynamic import failed, attempting to reload:', error.message)

                    // Prevent the default error handling
                    event.preventDefault()

                    // Reload the page after a short delay
                    setTimeout(() => {
                         window.location.reload()
                    }, 1000)

                    return
               }

               // Log other errors but don't prevent default handling
               console.error('Unhandled promise rejection:', error)
          })

          // Handle general errors
          window.addEventListener('error', (event) => {
               const error = event.error

               // Check for module loading errors
               if (
                    error?.message?.includes('Failed to fetch dynamically imported module') ||
                    error?.message?.includes('Loading chunk') ||
                    error?.name === 'ChunkLoadError'
               ) {
                    console.warn('🔄 Module loading error detected:', error.message)

                    // Reload the page after a short delay
                    setTimeout(() => {
                         window.location.reload()
                    }, 1000)
               }
          })
     }
})
