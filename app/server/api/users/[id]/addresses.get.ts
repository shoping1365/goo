declare const defineEventHandler: (handler: (event: unknown) => unknown | Promise<unknown>) => unknown
declare const getRouterParam: (event: unknown, name: string) => string | undefined
declare const createError: (options: { statusCode: number; statusMessage: string }) => Error
declare const getHeader: (event: unknown, name: string) => string | undefined
declare const getGoApiBaseUrl: () => string
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; headers?: Record<string, string> }) => Promise<T>

export default defineEventHandler(async (event) => {
     const id = getRouterParam(event, 'id')

     if (!id) {
          throw createError({
               statusCode: 400,
               statusMessage: 'User ID is required'
          })
     }

     try {
          const base = getGoApiBaseUrl()
          const response = await $fetch(`${base}/api/users/${id}/addresses`, {
               method: 'GET',
               headers: {
                    'Content-Type': 'application/json',
                    'Cookie': getHeader(event, 'cookie') || ''
               }
          })

          return response
     } catch (error: unknown) {
          const err = error as { statusCode?: number; statusMessage?: string }
          throw createError({
               statusCode: err.statusCode || 500,
               statusMessage: err.statusMessage || 'Failed to fetch addresses'
          })
     }
})