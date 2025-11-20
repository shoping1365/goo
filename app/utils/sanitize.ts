/**
 * XSS Prevention Utilities
 * Sanitize user input to prevent cross-site scripting
 */

/**
 * Sanitize HTML input
 * Remove potentially dangerous HTML tags and attributes
 */
export const sanitizeHtml = (input: string): string => {
  if (!input) return ''

  // Use a more comprehensive approach: remove all potentially dangerous content
  // Note: For production, consider using DOMPurify library for better security
  let sanitized = input

  // Remove script tags (improved regex to handle nested tags and edge cases)
  sanitized = sanitized.replace(/<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gis, '')
  
  // Remove style tags that could contain malicious CSS
  sanitized = sanitized.replace(/<style\b[^<]*(?:(?!<\/style>)<[^<]*)*<\/style>/gis, '')

  // Remove event handlers (improved to catch all variations)
  sanitized = sanitized.replace(/\s*on\w+\s*=\s*["'][^"']*["']/gi, '')
  sanitized = sanitized.replace(/\s*on\w+\s*=\s*[^\s>]*/gi, '')

  // Remove iframe tags completely
  sanitized = sanitized.replace(/<iframe\b[^<]*(?:(?!<\/iframe>)<[^<]*)*<\/iframe>/gis, '')

  // Remove object/embed tags completely
  sanitized = sanitized.replace(/<(object|embed)\b[^<]*(?:(?!<\/\1>)<[^<]*)*<\/\1>/gis, '')

  // Remove javascript: protocol in attributes
  sanitized = sanitized.replace(/javascript:/gi, '')

  // Remove data: URLs that could contain scripts
  sanitized = sanitized.replace(/data:text\/html/gi, '')

  return sanitized
}

/**
 * Sanitize text input
 * Remove all HTML tags
 */
export const sanitizeText = (input: string): string => {
  if (!input) return ''
  return input.replace(/<[^>]*>/g, '')
}

/**
 * Sanitize URL input
 * Allow only safe protocols
 */
export const sanitizeUrl = (input: string): string => {
  if (!input) return ''

  try {
    const url = new URL(input)
    // Only allow http and https
    if (!['http:', 'https:'].includes(url.protocol)) {
      return ''
    }
    return input
  } catch {
    // If not a valid URL, return empty
    return ''
  }
}

/**
 * Escape HTML special characters
 * Use when rendering user content
 */
export const escapeHtml = (input: string): string => {
  if (!input) return ''

  const map: Record<string, string> = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#039;',
  }

  return input.replace(/[&<>"']/g, char => map[char])
}

/**
 * Sanitize JSON input
 * Validate JSON structure
 */
export const sanitizeJson = (input: string): Record<string, unknown> | null => {
  try {
    return JSON.parse(input)
  } catch {
    return null
  }
}

/**
 * Sanitize mobile number
 * Remove invalid characters
 */
export const sanitizeMobileNumber = (input: string): string => {
  if (!input) return ''

  // Remove non-numeric characters except + and -
  let sanitized = input.replace(/[^\d+\-]/g, '')

  // Ensure starts with + or 0
  if (!sanitized.startsWith('+') && !sanitized.startsWith('0')) {
    sanitized = '0' + sanitized
  }

  return sanitized
}

/**
 * Validate email format
 * Basic email validation
 */
export const validateEmail = (email: string): boolean => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

/**
 * Validate password strength
 * Minimum 8 chars, uppercase, lowercase, number
 */
export const validatePasswordStrength = (password: string): {
  valid: boolean
  strength: 'weak' | 'medium' | 'strong'
  errors: string[]
} => {
  const errors: string[] = []

  if (password.length < 8) {
    errors.push('رمز عبور باید حداقل 8 کاراکتر باشد')
  }

  if (!/[A-Z]/.test(password)) {
    errors.push('رمز عبور باید شامل حرف بزرگ انگلیسی باشد')
  }

  if (!/[a-z]/.test(password)) {
    errors.push('رمز عبور باید شامل حرف کوچک انگلیسی باشد')
  }

  if (!/\d/.test(password)) {
    errors.push('رمز عبور باید شامل عدد باشد')
  }

  if (!/[!@#$%^&*]/.test(password)) {
    errors.push('رمز عبور باید شامل نماد خاص باشد')
  }

  let strength: 'weak' | 'medium' | 'strong' = 'weak'
  if (errors.length === 0) {
    strength = 'strong'
  } else if (errors.length <= 2) {
    strength = 'medium'
  }

  return {
    valid: errors.length === 0,
    strength,
    errors,
  }
}
