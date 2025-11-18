/**
 * Performance Monitoring Utility
 * Track API performance and system metrics
 */

interface PerformanceMetric {
  endpoint: string
  method: string
  duration: number // milliseconds
  statusCode: number
  timestamp: Date
  memoryUsage: NodeJS.MemoryUsage
  userId?: string
}

interface PerformanceStats {
  endpoint: string
  method: string
  avgDuration: number
  minDuration: number
  maxDuration: number
  totalRequests: number
  errorCount: number
  successRate: number
}

// In-memory metrics store
const metrics: PerformanceMetric[] = []

/**
 * Record performance metric
 */
export const recordMetric = (metric: PerformanceMetric): void => {
  metrics.push(metric)

  // Keep only last 10000 metrics in memory
  if (metrics.length > 10000) {
    metrics.shift()
  }
}

/**
 * Get endpoint statistics
 */
export const getEndpointStats = (endpoint: string, method?: string): PerformanceStats => {
  let filtered = metrics.filter(m => m.endpoint === endpoint)

  if (method) {
    filtered = filtered.filter(m => m.method === method)
  }

  if (filtered.length === 0) {
    return {
      endpoint,
      method: method || 'ALL',
      avgDuration: 0,
      minDuration: 0,
      maxDuration: 0,
      totalRequests: 0,
      errorCount: 0,
      successRate: 0,
    }
  }

  const durations = filtered.map(m => m.duration)
  const errors = filtered.filter(m => m.statusCode >= 400).length

  return {
    endpoint,
    method: method || 'ALL',
    avgDuration: durations.reduce((a, b) => a + b, 0) / durations.length,
    minDuration: Math.min(...durations),
    maxDuration: Math.max(...durations),
    totalRequests: filtered.length,
    errorCount: errors,
    successRate: ((filtered.length - errors) / filtered.length) * 100,
  }
}

/**
 * Get all endpoint statistics
 */
export const getAllEndpointStats = (): PerformanceStats[] => {
  const endpoints = new Set(metrics.map(m => m.endpoint))
  const stats: PerformanceStats[] = []

  endpoints.forEach(endpoint => {
    stats.push(getEndpointStats(endpoint))
  })

  return stats
}

/**
 * Get slow endpoints (average duration > threshold)
 */
export const getSlowEndpoints = (thresholdMs: number = 1000): PerformanceStats[] => {
  return getAllEndpointStats().filter(stat => stat.avgDuration > thresholdMs)
}

/**
 * Get performance summary
 */
export const getPerformanceSummary = (): {
  totalRequests: number
  avgDuration: number
  errorRate: number
  slowEndpoints: number
  memoryUsage: NodeJS.MemoryUsage
} => {
  if (metrics.length === 0) {
    return {
      totalRequests: 0,
      avgDuration: 0,
      errorRate: 0,
      slowEndpoints: 0,
      memoryUsage: process.memoryUsage(),
    }
  }

  const durations = metrics.map(m => m.duration)
  const errors = metrics.filter(m => m.statusCode >= 400).length
  const slowEndpoints = getSlowEndpoints().length

  return {
    totalRequests: metrics.length,
    avgDuration: durations.reduce((a, b) => a + b, 0) / durations.length,
    errorRate: (errors / metrics.length) * 100,
    slowEndpoints,
    memoryUsage: process.memoryUsage(),
  }
}

/**
 * Get metrics for time range
 */
export const getMetricsInRange = (
  startDate: Date,
  endDate: Date,
  endpoint?: string
): PerformanceMetric[] => {
  let filtered = metrics.filter(m => m.timestamp >= startDate && m.timestamp <= endDate)

  if (endpoint) {
    filtered = filtered.filter(m => m.endpoint === endpoint)
  }

  return filtered
}

/**
 * Get top slowest requests
 */
export const getTopSlowestRequests = (limit: number = 10): PerformanceMetric[] => {
  return [...metrics].sort((a, b) => b.duration - a.duration).slice(0, limit)
}

/**
 * Get requests by user
 */
export const getRequestsByUser = (userId: string, limit: number = 50): PerformanceMetric[] => {
  return metrics.filter(m => m.userId === userId).slice(-limit).reverse()
}

/**
 * Get endpoint performance trend
 */
export const getEndpointTrend = (
  endpoint: string,
  timeInterval: 'minute' | 'hour' | 'day' = 'minute'
): Array<{ time: string; avgDuration: number; requestCount: number }> => {
  const filtered = metrics.filter(m => m.endpoint === endpoint)

  if (filtered.length === 0) return []

  const grouped = new Map<string, PerformanceMetric[]>()

  filtered.forEach(metric => {
    let key: string

    if (timeInterval === 'minute') {
      key = metric.timestamp
        .toISOString()
        .substring(0, 16)
    } else if (timeInterval === 'hour') {
      key = metric.timestamp
        .toISOString()
        .substring(0, 13)
    } else {
      key = metric.timestamp
        .toISOString()
        .substring(0, 10)
    }

    if (!grouped.has(key)) {
      grouped.set(key, [])
    }

    grouped.get(key)!.push(metric)
  })

  const trend: Array<{ time: string; avgDuration: number; requestCount: number }> = []

  grouped.forEach((requests, time) => {
    const durations = requests.map(r => r.duration)
    trend.push({
      time,
      avgDuration: durations.reduce((a, b) => a + b, 0) / durations.length,
      requestCount: requests.length,
    })
  })

  return trend
}

/**
 * Clear metrics
 */
export const clearMetrics = (): void => {
  metrics.length = 0
}

/**
 * Get total metrics count
 */
export const getMetricsCount = (): number => {
  return metrics.length
}

/**
 * Analyze performance health
 */
export const analyzePerformanceHealth = (): {
  status: 'healthy' | 'warning' | 'critical'
  issues: string[]
  recommendations: string[]
} => {
  const summary = getPerformanceSummary()
  const slowEndpoints = getSlowEndpoints()
  const issues: string[] = []
  const recommendations: string[] = []

  // Check average duration
  if (summary.avgDuration > 2000) {
    issues.push(`Average response time is high: ${summary.avgDuration.toFixed(2)}ms`)
    recommendations.push('Review slow endpoints and optimize database queries')
  }

  // Check error rate
  if (summary.errorRate > 5) {
    issues.push(`Error rate is high: ${summary.errorRate.toFixed(2)}%`)
    recommendations.push('Review application logs for error patterns')
  }

  // Check memory usage
  const heapUsedPercent = (summary.memoryUsage.heapUsed / summary.memoryUsage.heapTotal) * 100

  if (heapUsedPercent > 85) {
    issues.push(`High memory usage: ${heapUsedPercent.toFixed(2)}%`)
    recommendations.push('Monitor for memory leaks and optimize memory usage')
  }

  // Check slow endpoints
  if (slowEndpoints.length > 0) {
    issues.push(`${slowEndpoints.length} endpoints with high latency`)
    recommendations.push(
      `Review: ${slowEndpoints.map(e => e.endpoint).join(', ')}`
    )
  }

  const status =
    issues.length === 0 ? 'healthy' : issues.length <= 2 ? 'warning' : 'critical'

  return {
    status,
    issues,
    recommendations,
  }
}
