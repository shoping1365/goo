/**
 * Monitoring API Routes
 * Get monitoring and logging data
 */

import type { H3Event } from 'h3'
import { createError, defineEventHandler, getQuery, getRouterParam } from 'h3'

declare const getRecentErrors: (limit: number) => unknown[]
declare const getErrorStats: () => unknown
declare const getErrorsByType: (type: string, limit: number) => unknown[]
declare const getEndpointStats: (endpoint: string, method?: string) => unknown
declare const getAllEndpointStats: () => unknown
declare const getPerformanceSummary: () => unknown
declare const analyzePerformanceHealth: () => unknown
declare const getActivities: (options: { limit: number; offset: number; userId?: string; action?: string }) => Promise<unknown[]>
declare const getActivityStats: (startDate: Date, endDate: Date) => Promise<unknown>

/**
 * GET /api/admin/monitoring/errors
 * Get recent errors
 */
export const getErrors = defineEventHandler(async (event: H3Event) => {
  const limit = Number(getQuery(event).limit) || 50

  return {
    errors: getRecentErrors(limit),
  }
})

/**
 * GET /api/admin/monitoring/errors/stats
 * Get error statistics
 */
export const getErrorsStats = defineEventHandler(async () => {
  return getErrorStats()
})

/**
 * GET /api/admin/monitoring/errors/type/:type
 * Get errors by type
 */
export const getErrorsByTypeRoute = defineEventHandler(async (event: H3Event) => {
  const type = getRouterParam(event, 'type')
  const limit = Number(getQuery(event).limit) || 50

  if (!type) {
    throw createError({
      statusCode: 400,
      message: 'Error type is required',
    })
  }

  return {
    errors: getErrorsByType(type, limit),
  }
})

/**
 * GET /api/admin/monitoring/performance
 * Get performance summary
 */
export const getPerformance = defineEventHandler(async () => {
  return getPerformanceSummary()
})

/**
 * GET /api/admin/monitoring/performance/health
 * Get performance health analysis
 */
export const getPerformanceHealth = defineEventHandler(async () => {
  return analyzePerformanceHealth()
})

/**
 * GET /api/admin/monitoring/performance/endpoints
 * Get all endpoints statistics
 */
export const getEndpointsStats = defineEventHandler(async () => {
  return getAllEndpointStats()
})

/**
 * GET /api/admin/monitoring/performance/endpoint/:endpoint
 * Get specific endpoint statistics
 */
export const getEndpointStatsRoute = defineEventHandler(async (event: H3Event) => {
  const endpoint = getRouterParam(event, 'endpoint')
  const method = getQuery(event).method as string | undefined

  if (!endpoint) {
    throw createError({
      statusCode: 400,
      message: 'Endpoint is required',
    })
  }

  return getEndpointStats(endpoint, method)
})

/**
 * GET /api/admin/monitoring/activities
 * Get activities
 */
export const getActivitiesRoute = defineEventHandler(async (event: H3Event) => {
  const limit = Number(getQuery(event).limit) || 50
  const offset = Number(getQuery(event).offset) || 0
  const userId = getQuery(event).userId as string | undefined
  const action = getQuery(event).action as string | undefined

  return await getActivities({
    limit,
    offset,
    userId,
    action,
  })
})

/**
 * GET /api/admin/monitoring/activities/stats
 * Get activity statistics
 */
export const getActivitiesStats = defineEventHandler(async (event: H3Event) => {
  const startDate = new Date(getQuery(event).startDate as string)
  const endDate = new Date(getQuery(event).endDate as string)

  if (isNaN(startDate.getTime()) || isNaN(endDate.getTime())) {
    throw createError({
      statusCode: 400,
      message: 'Valid startDate and endDate are required',
    })
  }

  return await getActivityStats(startDate, endDate)
})

/**
 * GET /api/admin/monitoring/dashboard
 * Get complete dashboard data
 */
export const getDashboard = defineEventHandler(async () => {
  const allEndpointStats = getAllEndpointStats()
  const endpoints = Array.isArray(allEndpointStats) ? allEndpointStats.slice(0, 10) : allEndpointStats

  return {
    errors: {
      stats: getErrorStats(),
      recent: getRecentErrors(10),
    },
    performance: {
      summary: getPerformanceSummary(),
      health: analyzePerformanceHealth(),
      endpoints,
    },
  }
})
