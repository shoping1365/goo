import { H3Event, getRequestHeader } from 'h3'

/**
 * دریافت header از درخواست - سازگار با h3 v2
 * @param event - رویداد H3
 * @param name - نام header
 * @returns مقدار header یا undefined
 */
export function getHeader(event: H3Event, name: string): string | undefined {
  try {
    const value = getRequestHeader(event, name)
    // h3 v2 ممکن است array برگرداند
    return Array.isArray(value) ? value[0] : value
  } catch {
    // fallback برای compatibility
    return undefined
  }
} 