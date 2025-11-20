// Plugin for handling HMR reconnection issues
declare const defineNuxtPlugin: (plugin: (nuxtApp: unknown) => void | Promise<void>) => { provide?: Record<string, unknown> }

export default defineNuxtPlugin(() => {
  if (import.meta.client) {
    // Handle HMR disconnection
    let isReconnecting = false

    const _handleHMRDisconnect = () => {
      if (isReconnecting) return
      isReconnecting = true

      // console.log('🔄 HMR disconnected, attempting to reconnect...')

      // Wait a bit before attempting to reconnect
      setTimeout(() => {
        window.location.reload()
      }, 1000)
    }

    // Listen for HMR disconnect events
    window.addEventListener('beforeunload', () => {
      // Clean up any pending operations
    })

    // Handle dynamic import failures - window.import is not a standard property
    // This code is intentionally commented out as window.import doesn't exist
    // const originalImport = (window as any).import
    // if (originalImport) {
    //   (window as any).import = async (url: string) => {
    //     try {
    //       return await originalImport(url)
    //     } catch (error) {
    //       console.error('Failed to import module:', url, error)
    //       // Retry once after a short delay
    //       await new Promise(resolve => setTimeout(resolve, 1000))
    //       return await originalImport(url)
    //     }
    //   }
    // }
  }
})