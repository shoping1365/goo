// Global Stream Error Handler Plugin
// این پلاگین خطاهای مربوط به stream را به صورت سراسری مدیریت می‌کند

// Override کردن console.warn برای فیلتر کردن پیام‌های خاص
const originalConsoleWarn = console.warn
const originalConsoleError = console.error

// فیلتر کردن پیام‌های مربوط به stream
const streamErrorMessages = [
  'Attempted to write to closed stream',
  'Attempted to write headers to closed stream',
  'write after end',
  'Cannot write headers after they are sent',
  'premature close',
  'Cannot pipe to a closed',
  'destroyed stream',
  'ECONNRESET'
]

console.warn = function(...args: any[]) {
  const message = args.join(' ')
  if (streamErrorMessages.some(msg => message.includes(msg))) {
    // خاموش کردن کامل این خطاها
    return
  }
  originalConsoleWarn.apply(console, args)
}

console.error = function(...args: any[]) {
  const message = args.join(' ')
  if (streamErrorMessages.some(msg => message.includes(msg))) {
    // خاموش کردن کامل این خطاها
    return
  }
  originalConsoleError.apply(console, args)
}

export default defineNitroPlugin((nitroApp) => {
  // مدیریت خطاهای unhandled در سطح سرور
  nitroApp.hooks.hook('error', (error) => {
    // بررسی خطاهای مربوط به stream
    if (streamErrorMessages.some(msg => error.message?.includes(msg))) {
      // خاموش کردن کامل - هیچ لاگی نزن
      return
    }
  })
  
  // مدیریت خطاهای مربوط به response
  nitroApp.hooks.hook('render:response', (response, { event }) => {
    // بررسی وضعیت response stream
    if (event.node.res.destroyed || event.node.res.writableEnded) {
      return
    }
  })
  
  // مدیریت خطاهای مربوط به request
  nitroApp.hooks.hook('request', (event) => {
    // بررسی وضعیت request stream
    if (event.node.req.destroyed) {
      return
    }
  })
})