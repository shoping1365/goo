// Server error handler plugin
// This plugin handles server-side errors and provides Persian error messages

export default defineNitroPlugin((nitroApp: { hooks: { hook: (name: string, handler: (error: { statusCode?: number; status?: number; message?: string }, context: { event?: { path?: string; node?: { req?: { method?: string; headers?: Record<string, string | string[] | undefined>; socket?: { remoteAddress?: string } } } } }) => void) => void } }) => {
  nitroApp.hooks.hook('error', (error: { statusCode?: number; status?: number; message?: string }, { event }: { event?: { path?: string; node?: { req?: { method?: string; headers?: Record<string, string | string[] | undefined>; socket?: { remoteAddress?: string } } } } }) => {
    // Attach structured error response in Farsi if possible
    try {
      const status = error?.statusCode || error?.status || 500
      const _url = event?.path || ''
      const _method = event?.node?.req?.method || ''
      const _ip = event?.node?.req?.headers?.['x-forwarded-for'] || event?.node?.req?.socket?.remoteAddress



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
    } catch { }
  })
})