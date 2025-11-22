/**
 * Error Tracking Utility
 * Track and manage application errors
 */

interface ErrorLog {
  error_type: string
  message: string
  stack_trace: string
  file_name?: string
  line_number?: number
  column_number?: number
  user_id?: string
  url: string
  user_agent: string
  ip_address: string
  context?: Record<string, unknown>
  severity: 'low' | 'medium' | 'high' | 'critical'
  timestamp: Date
  resolved: boolean
}

// In-memory error store (in production, use database)
const errorLogs: ErrorLog[] = []

/**
 * Log an error
 */
export const logError = async (error: ErrorLog): Promise<void> => {
  // Add to in-memory store
  errorLogs.push(error)

  // Keep only last 1000 errors in memory
  if (errorLogs.length > 1000) {
    errorLogs.shift()
  }

  // In production, also save to database
  // await saveErrorToDatabase(error)

  // Log to file in production
  if (process.env.NODE_ENV === 'production') {
    console.error(`[${error.severity}] ${error.error_type}: ${error.message}`)
  }
}

/**
 * Get recent errors
 */
export const getRecentErrors = (limit: number = 50): ErrorLog[] => {
  return errorLogs.slice(-limit).reverse()
}

/**
 * Get errors by severity
 */
export const getErrorsBySeverity = (
  severity: 'low' | 'medium' | 'high' | 'critical',
  limit: number = 50
): ErrorLog[] => {
  return errorLogs.filter(e => e.severity === severity).slice(-limit).reverse()
}

/**
 * Get errors by type
 */
export const getErrorsByType = (type: string, limit: number = 50): ErrorLog[] => {
  return errorLogs.filter(e => e.error_type === type).slice(-limit).reverse()
}

/**
 * Get error statistics
 */
export const getErrorStats = (): {
  total: number
  byType: Record<string, number>
  bySeverity: Record<string, number>
  resolved: number
  unresolved: number
} => {
  const stats = {
    total: errorLogs.length,
    byType: {} as Record<string, number>,
    bySeverity: {
      low: 0,
      medium: 0,
      high: 0,
      critical: 0,
    },
    resolved: 0,
    unresolved: 0,
  }

  errorLogs.forEach(error => {
    // Count by type
    stats.byType[error.error_type] = (stats.byType[error.error_type] || 0) + 1

    // Count by severity
    stats.bySeverity[error.severity]++

    // Count resolved
    if (error.resolved) {
      stats.resolved++
    } else {
      stats.unresolved++
    }
  })

  return stats
}

/**
 * Mark error as resolved
 */
export const markErrorResolved = (errorIndex: number): void => {
  if (errorLogs[errorIndex]) {
    errorLogs[errorIndex].resolved = true
  }
}

/**
 * Clear all errors
 */
export const clearErrors = (): void => {
  errorLogs.length = 0
}

/**
 * Determine error severity
 */
export const determineErrorSeverity = (
  error: Error,
  statusCode?: number
): 'low' | 'medium' | 'high' | 'critical' => {
  // Critical: 500+ status codes
  if (statusCode && statusCode >= 500) return 'critical'

  // High: 400+ status codes (except 404)
  if (statusCode && statusCode >= 400 && statusCode !== 404) return 'high'

  // High: Database errors
  if (
    error.message.includes('Database') ||
    error.message.includes('Prisma') ||
    error.message.includes('Connection')
  ) {
    return 'high'
  }

  // Medium: Application errors
  if (
    error.message.includes('TypeError') ||
    error.message.includes('ReferenceError') ||
    error.message.includes('SyntaxError')
  ) {
    return 'medium'
  }

  // Low: Validation errors
  if (error.message.includes('Validation') || error.message.includes('Invalid')) {
    return 'low'
  }

  return 'medium'
}

/**
 * Format error for logging
 */
export const formatErrorForLogging = (
  error: Error,
  context?: {
    url?: string
    userId?: string
    userAgent?: string
    ipAddress?: string
    statusCode?: number
  }
): ErrorLog => {
  const fileName = error.stack?.split('\n')[1]?.match(/\((.+?):(\d+):(\d+)\)/)?.[1]
  const lineNumber = error.stack?.split('\n')[1]?.match(/:(\d+):/)?.[1]
  const columnNumber = error.stack?.split('\n')[1]?.match(/:(\d+):(\d+)\)/)?.[2]

  return {
    error_type: error.name,
    message: error.message,
    stack_trace: error.stack || '',
    file_name: fileName,
    line_number: lineNumber ? parseInt(lineNumber) : undefined,
    column_number: columnNumber ? parseInt(columnNumber) : undefined,
    user_id: context?.userId,
    url: context?.url || '',
    user_agent: context?.userAgent || 'unknown',
    ip_address: context?.ipAddress || 'unknown',
    severity: determineErrorSeverity(error, context?.statusCode),
    timestamp: new Date(),
    resolved: false,
  }
}

/**
 * Handle and log error
 */
export const handleAndLogError = async (
  error: Error,
  context?: {
    url?: string
    userId?: string
    userAgent?: string
    ipAddress?: string
    statusCode?: number
  }
): Promise<ErrorLog> => {
  const errorLog = formatErrorForLogging(error, context)
  await logError(errorLog)
  return errorLog
}

/**
 * Get errors by user
 */
export const getErrorsByUser = (userId: string, limit: number = 50): ErrorLog[] => {
  return errorLogs.filter(e => e.user_id === userId).slice(-limit).reverse()
}

/**
 * Get errors in time range
 */
export const getErrorsInRange = (
  startDate: Date,
  endDate: Date,
  limit: number = 100
): ErrorLog[] => {
  return errorLogs
    .filter(e => e.timestamp >= startDate && e.timestamp <= endDate)
    .slice(-limit)
    .reverse()
}
