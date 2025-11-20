/**
 * Safe Error Handler Utilities
 * Return safe error messages to users without exposing sensitive info
 */

export interface ApiError {
  statusCode: number
  message: string
  userMessage: string
  details?: Record<string, unknown>
}

/**
 * Error message mappings
 * Map technical errors to user-friendly Persian messages
 */
const errorMessageMap: Record<string, string> = {
  // Authentication
  'Unauthorized': 'احراز هویت ناموفق. لطفاً دوباره وارد شوید.',
  'Invalid credentials': 'نام کاربری یا رمز عبور نادرست است.',
  'Token expired': 'نشست شما منقضی شده است. لطفاً دوباره وارد شوید.',
  'Invalid token': 'توکن نامعتبر است. لطفاً دوباره وارد شوید.',

  // Validation
  'Validation failed': 'داده های وارد شده نادرست است.',
  'Missing required fields': 'لطفاً تمام فیلد های الزامی را پر کنید.',
  'Invalid email': 'ایمیل وارد شده نامعتبر است.',
  'Invalid phone': 'شماره تلفن وارد شده نامعتبر است.',

  // Resource
  'Not found': 'منبع درخواست شده یافت نشد.',
  'Resource not found': 'منبع درخواست شده یافت نشد.',
  'Already exists': 'این منبع قبلاً ثبت شده است.',

  // Permission
  'Forbidden': 'شما دسترسی به این عملیات را ندارید.',
  'Insufficient permissions': 'شما دسترسی کافی برای انجام این عملیات را ندارید.',

  // Server
  'Internal server error': 'خطای داخلی سرور. لطفاً بعداً دوباره تلاش کنید.',
  'Service unavailable': 'سرویس در حال حاضر دردسترس نیست. لطفاً بعداً دوباره تلاش کنید.',
  'Bad gateway': 'خطای اتصال. لطفاً بعداً دوباره تلاش کنید.',

  // Rate limiting
  'Too many requests': 'بیش از حد درخواست ارسال کردید. لطفاً بعداً دوباره تلاش کنید.',

  // Network
  'Network error': 'خطای شبکه. لطفاً اتصال اینترنت خود را بررسی کنید.',
  'Connection timeout': 'مهلت اتصال به پایان رسید. لطفاً دوباره تلاش کنید.',
}

/**
 * Get safe error message
 * Return user-friendly error message without exposing sensitive info
 */
export const getSafeErrorMessage = (error: unknown): string => {
  const err = error as Record<string, unknown>
  // If error has userMessage, use it
  if ((err?.data as { message?: string })?.message) {
    return (err.data as { message: string }).message
  }

  // Try to find matching error message
  const errorMessage = err?.message || (error && typeof error.toString === 'function' ? error.toString() : '') || 'خطای نامشخصی رخ داد'

  for (const [key, value] of Object.entries(errorMessageMap)) {
    if (errorMessage.includes(key)) {
      return value
    }
  }

  // Return generic message for unknown errors
  return 'خطای نامشخصی رخ داد. لطفاً بعداً دوباره تلاش کنید.'
}

/**
 * Format API error response
 */
export const formatApiError = (error: unknown): ApiError => {
  const err = error as Record<string, unknown>
  const statusCode = (err?.status as number) || (err?.statusCode as number) || 500
  const message = (err?.message as string) || 'Unknown error'
  const userMessage = getSafeErrorMessage(error)

  return {
    statusCode,
    message,
    userMessage,
    details: err?.data as Record<string, unknown>,
  }
}

/**
 * Sanitize error object
 * Remove sensitive information from error
 */
export const sanitizeError = (error: unknown): Record<string, unknown> => {
  const err = error as Record<string, unknown>
  const sanitized: Record<string, unknown> = {}

  // Only include safe error properties
  const safeProperties = ['statusCode', 'status', 'message', 'data']

  for (const prop of safeProperties) {
    if (err?.[prop] !== undefined) {
      sanitized[prop] = err[prop]
    }
  }

  return sanitized
}

/**
 * Log error safely
 * Log sensitive information to backend without exposing to user
 */
export const logErrorSafely = (error: unknown, context?: Record<string, unknown>) => {
  const err = error as Record<string, unknown>
  if (process.client) {
    // In browser, send error to backend for logging
    if (typeof fetch !== 'undefined') {
      fetch('/api/admin/logs', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          type: 'error',
          message: typeof err?.message === 'string' ? err.message : 'Unknown error',
          stack: typeof err?.stack === 'string' ? err.stack : undefined,
          context: context || {},
          timestamp: new Date().toISOString(),
        }),
      }).catch(e => console.error('Failed to log error:', e))
    }
  }
}

/**
 * Create API error response
 */
export const createApiError = (
  statusCode: number,
  message: string,
  userMessage?: string
): ApiError => {
  return {
    statusCode,
    message,
    userMessage: userMessage || getSafeErrorMessage({ message }),
  }
}

/**
 * Extract status code from error
 */
export const getErrorStatusCode = (error: unknown): number => {
  const err = error as Record<string, unknown>
  if (typeof err?.status === 'number') return err.status
  if (typeof err?.statusCode === 'number') return err.statusCode
  
  const response = err?.response as Record<string, unknown> | undefined
  if (typeof response?.status === 'number') return response.status
  
  return 500
}

/**
 * Check if error is a validation error
 */
export const isValidationError = (error: unknown): boolean => {
  const err = error as Record<string, unknown>
  const statusCode = getErrorStatusCode(error)
  const data = err?.data as Record<string, unknown> | undefined
  return statusCode === 400 || data?.type === 'validation'
}

/**
 * Check if error is an authentication error
 */
export const isAuthError = (error: unknown): boolean => {
  const statusCode = getErrorStatusCode(error)
  return statusCode === 401 || statusCode === 403
}

/**
 * Check if error is a server error
 */
export const isServerError = (error: unknown): boolean => {
  const statusCode = getErrorStatusCode(error)
  return statusCode >= 500
}

/**
 * Check if error is a network error
 */
export const isNetworkError = (error: unknown): boolean => {
  const err = error as Record<string, unknown>
  const message = typeof err?.message === 'string' ? err.message : ''
  return err?.name === 'NetworkError' || message.includes('Network')
}
