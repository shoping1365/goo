<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">وضعیت سیستم</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <span class="text-sm text-gray-500">آخرین به‌روزرسانی: {{ lastUpdate }}</span>
        <button class="text-blue-600 hover:text-blue-800" @click="refreshStatus">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- System Health Status -->
    <div class="mb-6">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-semibold text-gray-900">سلامت سیستم</h4>
        <div class="flex items-center">
          <div class="w-3 h-3 rounded-full mr-2" :class="getStatusColor(status.overallHealth)"></div>
          <span class="text-sm font-medium" :class="getStatusTextColor(status.overallHealth)">
            {{ getStatusText(status.overallHealth) }}
          </span>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-gray-900">پایگاه داده</span>
            <div class="flex items-center">
              <div class="w-2 h-2 rounded-full mr-1" :class="getStatusColor(status.database)"></div>
              <span class="text-xs" :class="getStatusTextColor(status.database)">
                {{ getStatusText(status.database) }}
              </span>
            </div>
          </div>
          <div class="text-xs text-gray-500">
            زمان پاسخ: {{ status.databaseResponseTime }}ms
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-gray-900">API</span>
            <div class="flex items-center">
              <div class="w-2 h-2 rounded-full mr-1" :class="getStatusColor(status.api)"></div>
              <span class="text-xs" :class="getStatusTextColor(status.api)">
                {{ getStatusText(status.api) }}
              </span>
            </div>
          </div>
          <div class="text-xs text-gray-500">
            زمان پاسخ: {{ status.apiResponseTime }}ms
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-medium text-gray-900">سیستم اعتبارسنجی</span>
            <div class="flex items-center">
              <div class="w-2 h-2 rounded-full mr-1" :class="getStatusColor(status.creditSystem)"></div>
              <span class="text-xs" :class="getStatusTextColor(status.creditSystem)">
                {{ getStatusText(status.creditSystem) }}
              </span>
            </div>
          </div>
          <div class="text-xs text-gray-500">
            در دسترس: {{ status.creditSystemAvailability }}%
          </div>
        </div>
      </div>
    </div>

    <!-- Current Activity -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">فعالیت فعلی</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div class="bg-blue-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-blue-600 mb-1">{{ status.activeRequests }}</div>
          <div class="text-sm text-gray-600">درخواست‌های فعال</div>
        </div>
        <div class="bg-yellow-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-yellow-600 mb-1">{{ status.pendingApprovals }}</div>
          <div class="text-sm text-gray-600">در انتظار تایید</div>
        </div>
        <div class="bg-green-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-green-600 mb-1">{{ status.processingRequests }}</div>
          <div class="text-sm text-gray-600">در حال پردازش</div>
        </div>
        <div class="bg-purple-50 rounded-lg p-6 text-center">
          <div class="text-2xl font-bold text-purple-600 mb-1">{{ status.completedToday }}</div>
          <div class="text-sm text-gray-600">تکمیل شده امروز</div>
        </div>
      </div>
    </div>

    <!-- Performance Metrics -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">معیارهای عملکرد</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-900 mb-3">زمان پاسخ</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین</span>
              <span class="text-sm font-medium text-gray-900">{{ status.avgResponseTime }}ms</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حداکثر</span>
              <span class="text-sm font-medium text-gray-900">{{ status.maxResponseTime }}ms</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حداقل</span>
              <span class="text-sm font-medium text-gray-900">{{ status.minResponseTime }}ms</span>
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-900 mb-3">نرخ موفقیت</h5>
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درخواست‌ها</span>
              <span class="text-sm font-medium text-gray-900">{{ status.requestSuccessRate }}%</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پردازش</span>
              <span class="text-sm font-medium text-gray-900">{{ status.processingSuccessRate }}%</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تایید</span>
              <span class="text-sm font-medium text-gray-900">{{ status.approvalSuccessRate }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Alerts -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">هشدارهای اخیر</h4>
      <div class="space-y-3">
        <div v-for="alert in status.recentAlerts" :key="alert.id" class="flex items-start p-3 rounded-lg" :class="getAlertBackground(alert.level)">
          <div class="w-2 h-2 rounded-full mt-2 mr-3" :class="getAlertColor(alert.level)"></div>
          <div class="flex-1">
            <div class="flex items-center justify-between mb-1">
              <span class="text-sm font-medium text-gray-900">{{ alert.title }}</span>
              <span class="text-xs text-gray-500">{{ alert.time }}</span>
            </div>
            <p class="text-sm text-gray-600">{{ alert.message }}</p>
          </div>
        </div>
        <div v-if="status.recentAlerts.length === 0" class="text-center py-4 text-gray-500">
          هیچ هشداری وجود ندارد
        </div>
      </div>
    </div>

    <!-- System Information -->
    <div>
      <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات سیستم</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-900 mb-3">نسخه و اطلاعات</h5>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">نسخه سیستم:</span>
              <span class="font-medium">{{ status.systemVersion }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">آخرین به‌روزرسانی:</span>
              <span class="font-medium">{{ status.lastSystemUpdate }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">وضعیت به‌روزرسانی:</span>
              <span class="font-medium" :class="status.updateAvailable ? 'text-blue-600' : 'text-green-600'">
                {{ status.updateAvailable ? 'در دسترس' : 'به‌روز' }}
              </span>
            </div>
          </div>
        </div>

        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-900 mb-3">منابع سیستم</h5>
          <div class="space-y-3">
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-600">CPU</span>
                <span class="font-medium">{{ status.cpuUsage }}%</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-blue-500 h-2 rounded-full" :style="{ width: status.cpuUsage + '%' }"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-600">حافظه</span>
                <span class="font-medium">{{ status.memoryUsage }}%</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-500 h-2 rounded-full" :style="{ width: status.memoryUsage + '%' }"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between text-sm mb-1">
                <span class="text-gray-600">دیسک</span>
                <span class="font-medium">{{ status.diskUsage }}%</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-yellow-500 h-2 rounded-full" :style="{ width: status.diskUsage + '%' }"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

interface Alert {
  id: number
  title: string
  message: string
  level: 'info' | 'warning' | 'error'
  time: string
}

interface Status {
  overallHealth: 'healthy' | 'warning' | 'error'
  database: 'healthy' | 'warning' | 'error'
  api: 'healthy' | 'warning' | 'error'
  creditSystem: 'healthy' | 'warning' | 'error'
  databaseResponseTime: number
  apiResponseTime: number
  creditSystemAvailability: number
  activeRequests: number
  pendingApprovals: number
  processingRequests: number
  completedToday: number
  avgResponseTime: number
  maxResponseTime: number
  minResponseTime: number
  requestSuccessRate: number
  processingSuccessRate: number
  approvalSuccessRate: number
  recentAlerts: Alert[]
  systemVersion: string
  lastSystemUpdate: string
  updateAvailable: boolean
  cpuUsage: number
  memoryUsage: number
  diskUsage: number
}

const status = ref<Status>({
  overallHealth: 'healthy',
  database: 'healthy',
  api: 'healthy',
  creditSystem: 'healthy',
  databaseResponseTime: 0,
  apiResponseTime: 0,
  creditSystemAvailability: 0,
  activeRequests: 0,
  pendingApprovals: 0,
  processingRequests: 0,
  completedToday: 0,
  avgResponseTime: 0,
  maxResponseTime: 0,
  minResponseTime: 0,
  requestSuccessRate: 0,
  processingSuccessRate: 0,
  approvalSuccessRate: 0,
  recentAlerts: [],
  systemVersion: '',
  lastSystemUpdate: '',
  updateAvailable: false,
  cpuUsage: 0,
  memoryUsage: 0,
  diskUsage: 0
})

const lastUpdate = ref('')

const getStatusColor = (health: string): string => {
  switch (health) {
    case 'healthy': return 'bg-green-500'
    case 'warning': return 'bg-yellow-500'
    case 'error': return 'bg-red-500'
    default: return 'bg-gray-500'
  }
}

const getStatusTextColor = (health: string): string => {
  switch (health) {
    case 'healthy': return 'text-green-600'
    case 'warning': return 'text-yellow-600'
    case 'error': return 'text-red-600'
    default: return 'text-gray-600'
  }
}

const getStatusText = (health: string): string => {
  switch (health) {
    case 'healthy': return 'سالم'
    case 'warning': return 'هشدار'
    case 'error': return 'خطا'
    default: return 'نامشخص'
  }
}

const getAlertBackground = (level: string): string => {
  switch (level) {
    case 'info': return 'bg-blue-50'
    case 'warning': return 'bg-yellow-50'
    case 'error': return 'bg-red-50'
    default: return 'bg-gray-50'
  }
}

const getAlertColor = (level: string): string => {
  switch (level) {
    case 'info': return 'bg-blue-500'
    case 'warning': return 'bg-yellow-500'
    case 'error': return 'bg-red-500'
    default: return 'bg-gray-500'
  }
}

const formatTime = (): string => {
  const now = new Date()
  return now.toLocaleTimeString('fa-IR', {
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
}

const fetchStatus = async () => {
  try {
    interface StatusData {
      [key: string]: unknown
    }
    interface ApiResponse {
      data?: StatusData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/status')
    
    if (response.data) {
      status.value = response.data as unknown as Status
      lastUpdate.value = formatTime()
    }
  } catch (error) {
    console.error('خطا در دریافت وضعیت:', error)
    // در حالت واقعی، اینجا از createError استفاده می‌کنیم
  }
}

const refreshStatus = () => {
  fetchStatus()
}

onMounted(() => {
  fetchStatus()
  lastUpdate.value = formatTime()
})

// Auto-refresh every 30 seconds
setInterval(() => {
  fetchStatus()
}, 30000)
</script> 
