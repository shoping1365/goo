/**
 * Activity Logger Utility
 * Track user activities and audit logs
 */

// Note: In production, connect to your database
// For now, using in-memory storage

interface ActivityLog {
  user_id: string
  action: string
  resource_type: string
  resource_id?: string
  details?: Record<string, unknown>
  ip_address: string
  user_agent: string
  status: 'success' | 'failure'
  timestamp: Date
}

// In-memory storage for activity logs
const activityLogs: ActivityLog[] = []

/**
 * Log user activity
 */
export const logActivity = async (activity: ActivityLog): Promise<void> => {
  try {
    activityLogs.push(activity)
    
    // Keep only last 10000 logs in memory
    if (activityLogs.length > 10000) {
      activityLogs.shift()
    }

    // TODO: In production, save to database
  } catch (error) {
    console.error('Failed to log activity:', error)
  }
}

/**
 * Get user activities
 */
export const getUserActivities = async (
  userId: string,
  limit: number = 50,
  offset: number = 0
): Promise<ActivityLog[]> => {
  try {
    const userLogs = activityLogs.filter(log => log.user_id === userId)
    return userLogs.reverse().slice(offset, offset + limit)
  } catch (error) {
    console.error('Failed to get user activities:', error)
    return []
  }
}

/**
 * Get all activities with filtering
 */
export const getActivities = async (filters?: {
  userId?: string
  action?: string
  resourceType?: string
  status?: 'success' | 'failure'
  startDate?: Date
  endDate?: Date
  limit?: number
  offset?: number
}): Promise<ActivityLog[]> => {
  try {
    let filtered = [...activityLogs]

    if (filters?.userId) filtered = filtered.filter(l => l.user_id === filters.userId)
    if (filters?.action) filtered = filtered.filter(l => l.action === filters.action)
    if (filters?.resourceType) filtered = filtered.filter(l => l.resource_type === filters.resourceType)
    if (filters?.status) filtered = filtered.filter(l => l.status === filters.status)

    if (filters?.startDate || filters?.endDate) {
      filtered = filtered.filter(l => {
        if (filters?.startDate && l.timestamp < filters.startDate) return false
        if (filters?.endDate && l.timestamp > filters.endDate) return false
        return true
      })
    }

    const offset = filters?.offset || 0
    const limit = filters?.limit || 100

    return filtered.reverse().slice(offset, offset + limit)
  } catch (error) {
    console.error('Failed to get activities:', error)
    return []
  }
}

/**
 * Get activity statistics
 */
export const getActivityStats = async (
  startDate: Date,
  endDate: Date
): Promise<{
  totalActivities: number
  successCount: number
  failureCount: number
  actionBreakdown: Record<string, number>
}> => {
  try {
    const logs = activityLogs.filter(l => l.timestamp >= startDate && l.timestamp <= endDate)

    const actionBreakdown: Record<string, number> = {}

    logs.forEach((log: ActivityLog) => {
      actionBreakdown[log.action] = (actionBreakdown[log.action] || 0) + 1
    })

    return {
      totalActivities: logs.length,
      successCount: logs.filter((l: ActivityLog) => l.status === 'success').length,
      failureCount: logs.filter((l: ActivityLog) => l.status === 'failure').length,
      actionBreakdown,
    }
  } catch (error) {
    console.error('Failed to get activity stats:', error)
    return {
      totalActivities: 0,
      successCount: 0,
      failureCount: 0,
      actionBreakdown: {},
    }
  }
}

/**
 * Log common actions
 */
export const logUserAction = async (
  userId: string,
  action: string,
  resourceType: string,
  resourceId: string | undefined,
  status: 'success' | 'failure',
  details?: Record<string, unknown>,
  ipAddress?: string,
  userAgent?: string
): Promise<void> => {
  await logActivity({
    user_id: userId,
    action,
    resource_type: resourceType,
    resource_id: resourceId,
    details,
    ip_address: ipAddress || 'unknown',
    user_agent: userAgent || 'unknown',
    status,
    timestamp: new Date(),
  })
}

// Common action logging helpers

export const logLogin = async (
  userId: string,
  status: 'success' | 'failure',
  ipAddress?: string,
  userAgent?: string
): Promise<void> => {
  await logUserAction(userId, 'login', 'auth', undefined, status, {}, ipAddress, userAgent)
}

export const logLogout = async (
  userId: string,
  ipAddress?: string,
  userAgent?: string
): Promise<void> => {
  await logUserAction(userId, 'logout', 'auth', undefined, 'success', {}, ipAddress, userAgent)
}

export const logCreateUser = async (
  actorId: string,
  userId: string,
  status: 'success' | 'failure'
): Promise<void> => {
  await logUserAction(actorId, 'create', 'user', userId, status)
}

export const logUpdateUser = async (
  actorId: string,
  userId: string,
  changes: Record<string, unknown>,
  status: 'success' | 'failure'
): Promise<void> => {
  await logUserAction(actorId, 'update', 'user', userId, status, { changes })
}

export const logDeleteUser = async (
  actorId: string,
  userId: string,
  status: 'success' | 'failure'
): Promise<void> => {
  await logUserAction(actorId, 'delete', 'user', userId, status)
}

export const logRoleChange = async (
  actorId: string,
  userId: string,
  oldRole: string,
  newRole: string,
  status: 'success' | 'failure'
): Promise<void> => {
  await logUserAction(actorId, 'role_change', 'user', userId, status, {
    oldRole,
    newRole,
  })
}

export const logPermissionChange = async (
  actorId: string,
  resourceType: string,
  resourceId: string,
  permissions: string[],
  status: 'success' | 'failure'
): Promise<void> => {
  await logUserAction(actorId, 'permission_change', resourceType, resourceId, status, {
    permissions,
  })
}
