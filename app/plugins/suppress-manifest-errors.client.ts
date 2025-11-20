// حذف خطای manifest 404 در dev mode
// @ts-ignore
export default defineNuxtPlugin(() => {
  if (process.client) {
    const originalError = console.error
    console.error = (...args: unknown[]) => {
      const message = args[0]
      // فیلتر کردن خطاهای manifest
      if (
        typeof message === 'string' && 
        (message.includes('Error fetching app manifest') ||
         message.includes('builds/meta/dev.json') ||
         message.includes('error caught during app initialization'))
      ) {
        return
      }
      originalError.apply(console, args)
    }
  }
})
