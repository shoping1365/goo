// Server error handler plugin
// This plugin handles server-side errors and provides Persian error messages

export default defineNitroPlugin((nitroApp: any) => {
  nitroApp.hooks.hook('error', (error: any, { event }: { event: any }) => {
    // Attach structured error response in Farsi if possible
    try {
      const status = (error as any)?.statusCode || (error as any)?.status || 500
      const url = event?.path || ''
      const method = event?.node?.req?.method || ''
      const ip = event?.node?.req?.headers['x-forwarded-for'] || event?.node?.req?.socket?.remoteAddress



      // Ensure we always have a user-friendly message in Persian
      if (status >= 500) {
        ; (error as any).message = 'خطای داخلی سرور رخ داده است. لطفاً بعداً تلاش کنید.'
      } else if (status === 404) {
        ; (error as any).message = 'منبع درخواستی یافت نشد.'
      } else if (status === 401) {
        ; (error as any).message = 'برای ادامه باید وارد شوید.'
      } else if (status === 403) {
        // Security policy: present as 404 to client, but log 403 internally
        ; (error as any).statusCode = 404
          ; (error as any).message = 'منبع درخواستی یافت نشد.'
      }
    } catch { }
  })
})