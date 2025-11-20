// Global error handler plugin
declare const defineNuxtPlugin: (plugin: (nuxtApp: unknown) => void | Promise<void>) => { provide?: Record<string, unknown> }

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export default defineNuxtPlugin((nuxtApp: any) => {
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
