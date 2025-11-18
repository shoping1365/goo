// Plugin for retrying failed fetch requests
interface NuxtApp {
  vueApp?: unknown
  [key: string]: unknown
}

interface FetchOptions {
  method?: string
  body?: unknown
  headers?: Record<string, string>
  [key: string]: unknown
}

interface FetchError {
  message?: string
  name?: string
  stack?: string
}

type FetchFunction = (url: string | Request, options?: FetchOptions) => Promise<unknown>

declare const defineNuxtPlugin: (plugin: (nuxtApp: NuxtApp) => void | Promise<void>) => { provide?: Record<string, unknown> }

export default defineNuxtPlugin(() => {
  if (import.meta.client) {
    // Override $fetch with retry logic
    const originalFetch = (globalThis as { $fetch?: FetchFunction }).$fetch as FetchFunction | undefined

    if (originalFetch) {
      ; (globalThis as { $fetch: FetchFunction }).$fetch = async (url: string | Request, options: FetchOptions = {}) => {
        const maxRetries = 3
        const retryDelay = 1000

        for (let attempt = 1; attempt <= maxRetries; attempt++) {
          try {
            return await originalFetch(url, options)
          } catch (error: unknown) {
            const fetchError = error as FetchError
            // Check if it's a network error or module loading error
            if (
              fetchError.message?.includes('Failed to fetch dynamically imported module') ||
              fetchError.message?.includes('Loading chunk') ||
              fetchError.message?.includes('Loading CSS chunk') ||
              fetchError.name === 'TypeError' ||
              fetchError.name === 'ChunkLoadError'
            ) {
              console.warn(`🔄 Fetch attempt ${attempt}/${maxRetries} failed for ${url}:`, fetchError.message)

              if (attempt < maxRetries) {
                // Wait before retrying
                await new Promise(resolve => setTimeout(resolve, retryDelay * attempt))
                continue
              }
            }

            // If it's not a retryable error or we've exhausted retries, throw
            throw error
          }
        }
      }
    }
  }
})