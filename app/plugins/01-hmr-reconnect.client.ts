// Plugin for handling HMR reconnection issues
declare const defineNuxtPlugin: (plugin: (nuxtApp: any) => void | Promise<void>) => { provide?: Record<string, any> }

export default defineNuxtPlugin(() => {
  if (import.meta.client) {
    // Handle HMR disconnection
    let isReconnecting = false

    // Listen for HMR disconnect events
    window.addEventListener('beforeunload', () => {
      // Clean up any pending operations
    })


  }
})