// Global error handler plugin
declare const defineNuxtPlugin: (plugin: (nuxtApp: unknown) => void | Promise<void>) => { provide?: Record<string, unknown> }

interface NuxtApp {
  vueApp: {
    config: {
      errorHandler?: (error: unknown, instance: unknown, info: string) => void
    }
  }
  [key: string]: unknown
}
export default defineNuxtPlugin((nuxtApp: NuxtApp) => {
  nuxtApp.vueApp.config.errorHandler = (_error: unknown, _instance: unknown, _info: string) => {
    // Log error
    // console.error('Global error:', error)
    // console.error('Component info:', info)

    // Send to monitoring service
    if (process.client) {
      // TODO: Send to error tracking service (e.g., Sentry)
    }
  }

  // Handle unhandled promise rejections
  if (process.client) {
    window.addEventListener('unhandledrejection', (event) => {
      // console.error('Unhandled rejection:', event.reason)
      event.preventDefault()
    })
  }
})
