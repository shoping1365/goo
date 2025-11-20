import { defineNuxtPlugin } from 'nuxt/app'

// پلاگین امن‌سازی console.log برای جلوگیری از infinite loop در consola
export default defineNuxtPlugin(() => {
     if (!import.meta.client) return

     // ذخیره console.log اصلی
     const originalConsoleLog = globalThis.console.log
     const originalConsoleError = globalThis.console.error
     const originalConsoleWarn = globalThis.console.warn

     // تابع امن برای پردازش object ها
     const safeStringify = (obj: unknown, maxDepth = 3): string => {
          try {
               if (obj === null || obj === undefined) {
                    return String(obj)
               }

               if (typeof obj === 'string' || typeof obj === 'number' || typeof obj === 'boolean') {
                    return String(obj)
               }

               if (obj instanceof Error) {
                    return `Error: ${obj.message}`
               }

               if (Array.isArray(obj)) {
                    if (maxDepth <= 0) return '[Array]'
                    return `[${obj.slice(0, 5).map(item => safeStringify(item, maxDepth - 1)).join(', ')}${obj.length > 5 ? '...' : ''}]`
               }

               if (typeof obj === 'object') {
                    if (maxDepth <= 0) return '[Object]'
                    const keys = Object.keys(obj).slice(0, 5)
                    const pairs = keys.map(key => `${key}: ${safeStringify((obj as Record<string, unknown>)[key], maxDepth - 1)}`)
                    return `{${pairs.join(', ')}${Object.keys(obj).length > 5 ? '...' : ''}}`
               }

               return String(obj)
          } catch {
               return '[Circular or complex object]'
          }
     }

     // جایگزینی console.log با نسخه امن
     globalThis.console.log = (...args: unknown[]) => {
          try {
               const safeArgs = args.map(arg => safeStringify(arg))
               originalConsoleLog(...safeArgs)
          } catch {
               // originalConsoleLog('[Console log error]', error)
          }
     }

     // جایگزینی console.error با نسخه امن
     console.error = (...args: unknown[]) => {
          try {
               const safeArgs = args.map(arg => safeStringify(arg))
               originalConsoleError(...safeArgs)
          } catch (error) {
               originalConsoleError('[Console error log error]', error)
          }
     }

     // جایگزینی console.warn با نسخه امن
     console.warn = (...args: unknown[]) => {
          try {
               const safeArgs = args.map(arg => safeStringify(arg))
               originalConsoleWarn(...safeArgs)
          } catch (error) {
               originalConsoleWarn('[Console warn log error]', error)
          }
     }
})
