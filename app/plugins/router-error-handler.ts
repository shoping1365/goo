// فیلتر کردن Vue Router warnings برای 404 و static assets
// @ts-ignore
export default defineNuxtPlugin((nuxtApp) => {
  // فقط در سمت کلاینت
  if (process.client) {
    const originalWarn = console.warn
    const originalError = console.error
    
    console.warn = (...args: any[]) => {
      // فیلتر کردن warning های "No match found"
      const message = args[0]
      if (
        typeof message === 'string' && 
        (message.includes('No match found for location') || 
         message.includes('[Vue Router warn]') ||
         message.includes('/_nuxt/') ||
         message.includes('/uploads/') ||
         message.includes('.well-known'))
      ) {
        return
      }
      originalWarn.apply(console, args)
    }
    
    console.error = (...args: any[]) => {
      const message = args[0]
      if (
        typeof message === 'string' && 
        (message.includes('[Vue Router warn]') ||
         message.includes('No match found'))
      ) {
        return
      }
      originalError.apply(console, args)
    }
  }
})
