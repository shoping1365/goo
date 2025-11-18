// Stream Protection Middleware
// این میدلور از خطاهای مربوط به stream محافظت می‌کند
import { eventHandler } from 'h3'

export default eventHandler(async (event) => {
  // غیرفعال کردن کامل middleware برای رفع مشکلات encoding
  return
  // بررسی وضعیت response stream در ابتدای درخواست
  const originalEnd = event.node.res.end
  const originalWrite = event.node.res.write
  const originalWriteHead = event.node.res.writeHead

  // ردیابی وضعیت stream
  let streamClosed = false
  const isDev = process.env.NODE_ENV === 'development'

  // Override کردن متدهای response برای ردیابی وضعیت
  event.node.res.end = function (chunk?: any, encoding?: any, cb?: any) {
    streamClosed = true
    return originalEnd.call(this, chunk, encoding, cb)
  }

  event.node.res.write = function (chunk: any, encoding?: any, cb?: any) {
    if (streamClosed) {
      // فقط در محیط development لاگ می‌کنیم
      if (isDev) {
        console.warn('Attempted to write to closed stream, ignoring')
      }
      return false
    }
    return originalWrite.call(this, chunk, encoding, cb)
  }

  event.node.res.writeHead = function (statusCode: number, ...args: any[]) {
    if (streamClosed) {
      // فقط در محیط development لاگ می‌کنیم
      if (isDev) {
        console.warn('Attempted to write headers to closed stream, ignoring')
      }
      return this
    }
    // همه آرگومان‌ها را عیناً پاس بده تا هدرها از بین نروند
    // پشتیبانی از همه امضاهای Node: writeHead(status), writeHead(status, headers), writeHead(status, reason, headers)
    return originalWriteHead.apply(this, [statusCode, ...args])
  }

  // بررسی وضعیت stream در صورت بسته شدن زودهنگام
  event.node.res.on('close', () => {
    streamClosed = true
  })

  event.node.res.on('finish', () => {
    streamClosed = true
  })

  // بررسی وضعیت stream در صورت خطا
  event.node.res.on('error', (error) => {
    if (error.message?.includes('premature close') ||
      error.message?.includes('Cannot pipe to a closed') ||
      error.message?.includes('destroyed stream')) {
      if (isDev) {
        console.warn('Stream error detected:', error.message)
      }
      streamClosed = true
    }
  })

  // بررسی وضعیت stream قبل از ادامه پردازش
  if (event.node.res.destroyed || event.node.res.writableEnded) {
    if (isDev) {
      console.warn('Response stream already closed at middleware level')
    }
    streamClosed = true
  }
})
