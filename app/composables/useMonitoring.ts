/**
 * Monitoring Composable
 * Access monitoring and logging data
 */

import { computed, ref } from 'vue'
import { useApiClient } from '~/utils/api'
import { getSafeErrorMessage } from '~/utils/errorHandler'

export interface ErrorLog {
  error_type: string
  message: string
  severity: 'low' | 'medium' | 'high' | 'critical'
  timestamp: Date
  resolved: boolean
}

export interface PerformanceStat {
  endpoint: string
  method: string
  avgDuration: number
  totalRequests: number
  errorCount: number
  successRate: number
}

export interface ActivityLog {
  action: string
  resource_type: string
  user_id: string
  timestamp: Date
  status: 'success' | 'failure'
}

export interface ErrorStats {
  total: number
  bySeverity: Record<string, number>
  resolved: number
  [key: string]: unknown
}

export interface PerformanceSummary {
  avgDuration: number
  errorRate: number
  [key: string]: unknown
}

export interface PerformanceHealth {
  status: 'healthy' | 'warning' | 'critical'
  issues: string[]
  [key: string]: unknown
}

interface DashboardResponse {
  errors: {
    recent: ErrorLog[]
    stats: ErrorStats
  }
  performance: {
    summary: PerformanceSummary
    health: PerformanceHealth
    endpoints: PerformanceStat[]
  }
}

export const useMonitoring = () => {
  const { api } = useApiClient()

  // State
  const errors = ref<ErrorLog[]>([])
  const errorStats = ref<ErrorStats | null>(null)
  const performance = ref<PerformanceSummary | null>(null)
  const performanceHealth = ref<PerformanceHealth | null>(null)
  const endpoints = ref<PerformanceStat[]>([])
  const activities = ref<ActivityLog[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  /**
   * Fetch recent errors
   */
  const fetchErrors = async (limit: number = 50) => {
    loading.value = true
    error.value = null

    try {
      const response = await api.get<{ errors: ErrorLog[] }>(`/api/admin/monitoring/errors?limit=${limit}`)
      errors.value = response.errors
    } catch (err) {
      error.value = getSafeErrorMessage(err)
      console.error('Failed to fetch errors:', err)
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch error statistics
   */
  const fetchErrorStats = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await api.get<ErrorStats>('/api/admin/monitoring/errors/stats')
      errorStats.value = response
    } catch (err) {
      error.value = getSafeErrorMessage(err)
      console.error('Failed to fetch error stats:', err)
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch performance data
   */
  const fetchPerformance = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await api.get<PerformanceSummary>('/api/admin/monitoring/performance')
      performance.value = response
    } catch (err) {
      error.value = getSafeErrorMessage(err)
      console.error('Failed to fetch performance:', err)
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch performance health
   */
  const fetchPerformanceHealth = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await api.get<PerformanceHealth>('/api/admin/monitoring/performance/health')
      performanceHealth.value = response
    } catch (err) {
      error.value = getSafeErrorMessage(err)
      console.error('Failed to fetch health:', err)
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch endpoint statistics
   */
  const fetchEndpoints = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await api.get<PerformanceStat[]>('/api/admin/monitoring/performance/endpoints')
      endpoints.value = response
    } catch (err) {
      error.value = getSafeErrorMessage(err)
      console.error('Failed to fetch endpoints:', err)
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch activities
   */
  const fetchActivities = async (limit: number = 50, offset: number = 0) => {
    loading.value = true
    error.value = null

    try {
      const response = await api.get<ActivityLog[]>(
        `/api/admin/monitoring/activities?limit=${limit}&offset=${offset}`
      )
      activities.value = response
    } catch (err) {
      error.value = getSafeErrorMessage(err)
      console.error('Failed to fetch activities:', err)
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch dashboard data
   */
  const fetchDashboard = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await api.get<DashboardResponse>('/api/admin/monitoring/dashboard')
      errors.value = response.errors.recent
      errorStats.value = response.errors.stats
      performance.value = response.performance.summary
      performanceHealth.value = response.performance.health
      endpoints.value = response.performance.endpoints
    } catch (err) {
      error.value = getSafeErrorMessage(err)
      console.error('Failed to fetch dashboard:', err)
    } finally {
      loading.value = false
    }
  }

  /**
   * Get error severity color
   */
  const getErrorSeverityColor = (severity: string) => {
    const colors: Record<string, string> = {
      low: 'bg-blue-500',
      medium: 'bg-yellow-500',
      high: 'bg-orange-500',
      critical: 'bg-red-500',
    }
    return colors[severity] || 'bg-gray-500'
  }

  /**
   * Get health status color
   */
  const getHealthStatusColor = (status: string) => {
    const colors: Record<string, string> = {
      healthy: 'bg-green-500',
      warning: 'bg-yellow-500',
      critical: 'bg-red-500',
    }
    return colors[status] || 'bg-gray-500'
  }

  /**
   * Format duration
   */
  const formatDuration = (ms: number): string => {
    if (ms < 1000) return `${ms.toFixed(0)}ms`
    return `${(ms / 1000).toFixed(2)}s`
  }

  return {
    // State
    errors: computed(() => errors.value),
    errorStats: computed(() => errorStats.value),
    performance: computed(() => performance.value),
    performanceHealth: computed(() => performanceHealth.value),
    endpoints: computed(() => endpoints.value),
    activities: computed(() => activities.value),
    loading: computed(() => loading.value),
    error: computed(() => error.value),

    // Methods
    fetchErrors,
    fetchErrorStats,
    fetchPerformance,
    fetchPerformanceHealth,
    fetchEndpoints,
    fetchActivities,
    fetchDashboard,
    getErrorSeverityColor,
    getHealthStatusColor,
    formatDuration,
  }
}
