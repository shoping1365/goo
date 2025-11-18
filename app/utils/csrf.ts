/**
 * CSRF Protection Utilities
 * Generate and validate CSRF tokens
 */

/**
 * Generate a random CSRF token
 * Only available on server-side
 */
export const generateCsrfToken = (): string => {
  if (import.meta.server) {
    // Dynamic import only on server
    const crypto = require('crypto')
    return crypto.randomBytes(32).toString('hex')
  }
  // Fallback for client (shouldn't be called)
  return ''
}

/**
 * Validate CSRF token format
 */
export const validateCsrfToken = (token: string): boolean => {
  // Token should be 64 character hex string
  return /^[a-f0-9]{64}$/.test(token)
}

/**
 * Get CSRF token from cookie
 * Frontend utility
 */
export const getCsrfToken = (): string => {
  if (typeof document === 'undefined') return ''

  const cookieString = document.cookie
  const cookies = cookieString.split(';')

  for (const cookie of cookies) {
    const [name, value] = cookie.trim().split('=')
    if (name === 'csrf-token') {
      return decodeURIComponent(value)
    }
  }

  return ''
}

/**
 * Get CSRF token from meta tag
 * Alternative method
 */
export const getCsrfTokenFromMeta = (): string => {
  if (typeof document === 'undefined') return ''

  const meta = document.querySelector('meta[name="csrf-token"]')
  return meta?.getAttribute('content') || ''
}

/**
 * Add CSRF token to fetch headers
 */
export const addCsrfTokenToHeaders = (headers: Record<string, string> = {}): Record<string, string> => {
  const token = getCsrfToken() || getCsrfTokenFromMeta()

  if (token) {
    headers['X-CSRF-Token'] = token
  }

  return headers
}

/**
 * Add CSRF token to form data
 */
export const addCsrfTokenToFormData = (formData: FormData): FormData => {
  const token = getCsrfToken() || getCsrfTokenFromMeta()

  if (token) {
    formData.append('csrf_token', token)
  }

  return formData
}
