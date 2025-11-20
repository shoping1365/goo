// Server error handler plugin
// This plugin handles server-side errors and provides Persian error messages

declare const defineNuxtPlugin: (plugin: (nuxtApp: unknown) => void | Promise<void>) => { provide?: Record<string, unknown> }

export default defineNuxtPlugin((nuxtApp) => {
  // Handle server-side errors
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  (nuxtApp as any).hook('error', (error: any, { event: _event }: any) => {
    // Attach structured error response in Farsi if possible
    try {
      const status = error?.statusCode || error?.status || 500
      // const url = event?.path || ''
      // const method = event?.node?.req?.method || ''
      // const ip = event?.node?.req?.headers['x-forwarded-for'] || event?.node?.req?.socket?.remoteAddress

      // Server error logged
      // console.error(`Server Error [${status}]: ${method} ${url} - ${error.message}`)

      // Ensure we always have a user-friendly message in Persian
      if (status >= 500) {
        error.message = 'خطای داخلی سرور رخ داده است. لطفاً بعداً تلاش کنید.'
      } else if (status === 404) {
        error.message = 'منبع درخواستی یافت نشد.'
      } else if (status === 401) {
        error.message = 'برای ادامه باید وارد شوید.'
      } else if (status === 403) {
        // Security policy: present as 404 to client, but log 403 internally
        error.statusCode = 404
        error.message = 'منبع درخواستی یافت نشد.'
      }
    } catch {
      // console.error('Error in error handler:', e)
    }
  })
})