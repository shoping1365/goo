// Global error handler plugin
declare const defineNuxtPlugin: (plugin: (nuxtApp: any) => void | Promise<void>) => { provide?: Record<string, any> }

export default defineNuxtPlugin((nuxtApp: any) => {
  nuxtApp.vueApp.config.errorHandler = (error: any, instance: any, info: string) => {
    // Log error
    console.error('Global error:', error)
    console.error('Component info:', info)

    // Send to monitoring service
    if (process.client) {
      // TODO: Send to error tracking service (e.g., Sentry)
    }
  }

  // Handle unhandled promise rejections
  if (process.client) {
    window.addEventListener('unhandledrejection', (event) => {
      console.error('Unhandled rejection:', event.reason)
      event.preventDefault()
    })
  }
})
