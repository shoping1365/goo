// API Client Configuration
import { addCsrfTokenToHeaders } from '~/utils/csrf';
import { formatApiError, getSafeErrorMessage } from '~/utils/errorHandler';

declare const useRuntimeConfig: () => { public: { goApiBase: string } }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown; credentials?: string; headers?: Record<string, string> }) => Promise<T>

export const useApiClient = () => {
  const config = useRuntimeConfig()
  const apiBase = (config.public.goApiBase as string)

  // Error handling with safe messages  
  const _handleError = (error: unknown): string => {
    return getSafeErrorMessage(error)
  }

  // API calls with auth and security
  const api = {
    get: async <T>(endpoint: string) => {
      try {
        let headers: Record<string, string> = {}

        // Add CSRF token for GET requests to sensitive endpoints
        if (endpoint.startsWith('/api/admin')) {
          headers = addCsrfTokenToHeaders(headers)
        }

        return await $fetch<T>(`${apiBase}${endpoint}`, {
          method: 'GET',
          credentials: 'include',
          headers,
        })
      } catch (error) {
        const formattedError = formatApiError(error)
        throw formattedError
      }
    },

    post: async <T>(endpoint: string, body?: unknown) => {
      try {
        let headers: Record<string, string> = {
          'Content-Type': 'application/json',
        }

        // Add CSRF token
        headers = addCsrfTokenToHeaders(headers)

        return await $fetch<T>(`${apiBase}${endpoint}`, {
          method: 'POST',
          body,
          credentials: 'include',
          headers,
        })
      } catch (error) {
        const formattedError = formatApiError(error)
        throw formattedError
      }
    },

    put: async <T>(endpoint: string, body?: unknown) => {
      try {
        let headers: Record<string, string> = {
          'Content-Type': 'application/json',
        }

        // Add CSRF token
        headers = addCsrfTokenToHeaders(headers)

        return await $fetch<T>(`${apiBase}${endpoint}`, {
          method: 'PUT',
          body,
          credentials: 'include',
          headers,
        })
      } catch (error) {
        const formattedError = formatApiError(error)
        throw formattedError
      }
    },

    delete: async <T>(endpoint: string) => {
      try {
        let headers: Record<string, string> = {}

        // Add CSRF token
        headers = addCsrfTokenToHeaders(headers)

        return await $fetch<T>(`${apiBase}${endpoint}`, {
          method: 'DELETE',
          credentials: 'include',
          headers,
        })
      } catch (error) {
        const formattedError = formatApiError(error)
        throw formattedError
      }
    },
  }

  return { api, apiBase }
}
