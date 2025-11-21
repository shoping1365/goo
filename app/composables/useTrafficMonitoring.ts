// تعریف interface برای response های API
interface ApiResponse<T = unknown> {
  data?: T
  success?: boolean
  message?: string
}

interface TrafficStats {
  online_users: number
  today_requests: number
  suspicious_requests: number
  blocked_attacks: number
  hourly_traffic: unknown[]
}

interface OnlineUser {
  [key: string]: unknown
}

export const useTrafficMonitoring = () => {
  // به‌روزرسانی سشن کاربر
  const updateUserSession = async (sessionData: {
    session_id: string
    current_page?: string
    ip_address?: string
    user_agent?: string
  }) => {
    try {
      await $fetch('/api/admin/traffic/update-session', {
        method: 'POST',
        body: sessionData
      })
    } catch (error) {
      console.warn('خطا در به‌روزرسانی سشن:', error)
    }
  }

  // ثبت درخواست ترافیک
  const logTrafficRequest = async (requestData: {
    ip_address: string
    user_agent?: string
    request_path: string
    method: string
    status_code: number
    response_time?: number
  }) => {
    try {
      await $fetch('/api/admin/traffic/log-request', {
        method: 'POST',
        body: requestData
      })
    } catch (error) {
      console.warn('خطا در ثبت لاگ ترافیک:', error)
    }
  }

  // دریافت آمار ترافیک
  const getTrafficStats = async () => {
    try {
      const res = await $fetch<ApiResponse<TrafficStats>>('/api/admin/traffic/stats')
      return res.data
    } catch (error) {
      console.warn('خطا در دریافت آمار ترافیک:', error)
      return {
        online_users: 0,
        today_requests: 0,
        suspicious_requests: 0,
        blocked_attacks: 0,
        hourly_traffic: []
      }
    }
  }

  // دریافت کاربران آنلاین
  const getOnlineUsers = async () => {
    try {
      const res = await $fetch<ApiResponse<OnlineUser[]>>('/api/admin/traffic/online-users')
      return res.data || []
    } catch (error) {
      console.warn('خطا در دریافت کاربران آنلاین:', error)
      return []
    }
  }

  // مسدود کردن IP
  const blockIP = async (blockData: {
    ip_address: string
    reason: string
    duration: number
  }) => {
    try {
      const response = await $fetch('/api/admin/traffic/block-ip', {
        method: 'POST',
        body: blockData
      })
      return response
    } catch (error) {
      console.warn('خطا در مسدود کردن IP:', error)
      throw error
    }
  }

  // آزاد کردن IP
  const unblockIP = async (ip: string) => {
    try {
      const response = await $fetch(`/api/admin/traffic/unblock-ip/${ip}`, {
        method: 'DELETE'
      })
      return response
    } catch (error) {
      console.warn('خطا در آزاد کردن IP:', error)
      throw error
    }
  }

  // دریافت لیست IP های مسدود شده
  const getBlockedIPs = async () => {
    try {
      const res = await $fetch<ApiResponse<Array<Record<string, unknown>>>>('/api/admin/traffic/blocked-ips')
      return res.data || []
    } catch (error) {
      console.error('خطا در دریافت لیست IP های مسدود شده:', error)
      return []
    }
  }

  // پاکسازی سشن‌های منقضی شده
  const cleanupExpiredSessions = async () => {
    try {
      const response = await $fetch('/api/admin/traffic/cleanup', {
        method: 'POST'
      })
      return response
    } catch (error) {
      console.error('خطا در پاکسازی سشن‌ها:', error)
      throw error
    }
  }

  return {
    updateUserSession,
    logTrafficRequest,
    getTrafficStats,
    getOnlineUsers,
    blockIP,
    unblockIP,
    getBlockedIPs,
    cleanupExpiredSessions
  }
} 