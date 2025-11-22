// Response Safety Utilities
// ابزارهای امن برای مدیریت response streams
import type { H3Event } from 'h3'

/**
 * بررسی امن بودن response stream برای نوشتن
 */
export function isResponseStreamSafe(event: H3Event): boolean {
  const res = event.node.res
  
  // بررسی وضعیت‌های مختلف stream
  if (res.destroyed || res.writableEnded || res.finished) {
    return false
  }
  
  // بررسی متد کمکی که در middleware تعریف شده
  if (typeof (res as { isStreamClosed?: () => boolean }).isStreamClosed === 'function') {
    return !(res as { isStreamClosed: () => boolean }).isStreamClosed()
  }
  
  return true
}

/**
 * ارسال پاسخ امن با بررسی وضعیت stream
 */
export function safeSendResponse(event: H3Event, _data: unknown): boolean {
  if (!isResponseStreamSafe(event)) {
    console.warn('Cannot send response: stream is closed or destroyed')
    return false
  }
  
  try {
    // بررسی مجدد قبل از ارسال
    if (!isResponseStreamSafe(event)) {
      return false
    }
    
    return true
  } catch (error) {
    console.error('Error sending response:', error)
    return false
  }
}

/**
 * مدیریت خطا با بررسی وضعیت stream
 */
export function safeHandleError(event: H3Event, _error: unknown): boolean {
  if (!isResponseStreamSafe(event)) {
    console.warn('Cannot send error response: stream is closed or destroyed')
    return false
  }
  
  try {
    // بررسی مجدد قبل از ارسال خطا
    if (!isResponseStreamSafe(event)) {
      return false
    }
    
    return true
  } catch (err) {
    console.error('Error handling error response:', err)
    return false
  }
}

/**
 * بررسی و لاگ کردن وضعیت stream
 */
export function logStreamStatus(event: H3Event, context: string = 'Unknown'): void {
  const res = event.node.res
  const status = {
    destroyed: res.destroyed,
    writableEnded: res.writableEnded,
    finished: res.finished,
    writable: res.writable,
    isStreamClosed: typeof (res as { isStreamClosed?: () => boolean }).isStreamClosed === 'function' ? (res as { isStreamClosed: () => boolean }).isStreamClosed() : 'N/A'
  }
  
  console.warn(`Stream status [${context}]:`, status)
}
