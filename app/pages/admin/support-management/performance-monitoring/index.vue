<template>
  <div class="performance-monitoring" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">مانیتورینگ عملکرد چت آنلاین</h1>
          <p class="text-sm text-gray-600 mt-1">نظارت بر عملکرد سیستم چت و پشتیبانی</p>
        </div>
        <div class="flex items-center space-x-4 space-x-reverse">
          <!-- دکمه کنترل مانیتورینگ -->
          <button 
            @click="toggleMonitoring"
            :disabled="isToggling"
            :class="[
              'inline-flex items-center px-4 py-2 border text-sm font-medium rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:opacity-50',
              monitoringEnabled 
                ? 'border-red-300 text-red-700 bg-red-50 hover:bg-red-100 focus:ring-red-500' 
                : 'border-green-300 text-green-700 bg-green-50 hover:bg-green-100 focus:ring-green-500'
            ]"
          >
            <svg v-if="isToggling" class="animate-spin -ml-1 mr-3 h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else class="-ml-1 mr-3 h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path v-if="monitoringEnabled" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            {{ isToggling ? 'در حال تغییر...' : (monitoringEnabled ? 'خاموش کردن' : 'روشن کردن') }}
          </button>
          
          <button 
            @click="refreshData"
            :disabled="isRefreshing"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
          >
            <svg v-if="isRefreshing" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else class="-ml-1 mr-3 h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            {{ isRefreshing ? 'در حال بروزرسانی...' : 'بروزرسانی' }}
          </button>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="text-sm text-gray-500">وضعیت:</span>
            <span :class="getStatusClass(systemStatus.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
              <span class="w-2 h-2 rounded-full mr-1" :class="getStatusDotClass(systemStatus.status)"></span>
              {{ getStatusText(systemStatus.status) }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="p-6">
      <!-- Key Metrics Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <!-- Active Connections -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-blue-100 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
            </div>
            <div class="mr-4 flex-1">
              <p class="text-sm font-medium text-gray-600">اتصالات فعال</p>
              <p class="text-2xl font-semibold text-gray-900">{{ metrics.activeConnections }}</p>
            </div>
            <div class="text-right">
              <span :class="getTrendClass(metrics.connectionsTrend)" class="text-sm font-medium">
                {{ metrics.connectionsTrend > 0 ? '+' : '' }}{{ metrics.connectionsTrend }}%
              </span>
            </div>
          </div>
        </div>

        <!-- Rate Limited Messages -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-red-100 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                </svg>
              </div>
            </div>
            <div class="mr-4 flex-1">
              <p class="text-sm font-medium text-gray-600">پیام‌های محدود شده</p>
              <p class="text-2xl font-semibold text-gray-900">{{ metrics.rateLimitedMessages }}</p>
            </div>
            <div class="text-right">
              <span :class="getTrendClass(metrics.rateLimitTrend)" class="text-sm font-medium">
                {{ metrics.rateLimitTrend > 0 ? '+' : '' }}{{ metrics.rateLimitTrend }}%
              </span>
            </div>
          </div>
        </div>

        <!-- Response Time -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-green-100 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
            <div class="mr-4 flex-1">
              <p class="text-sm font-medium text-gray-600">زمان پاسخ</p>
              <p class="text-2xl font-semibold text-gray-900">{{ metrics.avgResponseTime }}ms</p>
            </div>
            <div class="text-right">
              <span :class="getTrendClass(metrics.responseTimeTrend)" class="text-sm font-medium">
                {{ metrics.responseTimeTrend > 0 ? '+' : '' }}{{ metrics.responseTimeTrend }}ms
              </span>
            </div>
          </div>
        </div>

        <!-- Messages Per Second -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-purple-100 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                </svg>
              </div>
            </div>
            <div class="mr-4 flex-1">
              <p class="text-sm font-medium text-gray-600">پیام در ثانیه</p>
              <p class="text-2xl font-semibold text-gray-900">{{ metrics.messagesPerSecond }}</p>
            </div>
            <div class="text-right">
              <span :class="getTrendClass(metrics.messagesTrend)" class="text-sm font-medium">
                {{ metrics.messagesTrend > 0 ? '+' : '' }}{{ metrics.messagesTrend }}%
              </span>
            </div>
          </div>
        </div>

        <!-- Error Rate -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-red-100 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
            <div class="mr-4 flex-1">
              <p class="text-sm font-medium text-gray-600">نرخ خطا</p>
              <p class="text-2xl font-semibold text-gray-900">{{ metrics.errorRate }}%</p>
            </div>
            <div class="text-right">
              <span :class="getTrendClass(-metrics.errorRateTrend)" class="text-sm font-medium">
                {{ metrics.errorRateTrend > 0 ? '+' : '' }}{{ metrics.errorRateTrend }}%
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Charts Section -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <!-- Response Time Chart -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">زمان پاسخ در 24 ساعت گذشته</h3>
          <div class="h-64">
            <canvas ref="responseTimeChart"></canvas>
          </div>
        </div>

        <!-- Messages Chart -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">تعداد پیام‌ها در 24 ساعت گذشته</h3>
          <div class="h-64">
            <canvas ref="messagesChart"></canvas>
          </div>
        </div>
      </div>

      <!-- System Health & Alerts -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- System Health -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">وضعیت سیستم</h3>
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">CPU</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                  <div class="bg-blue-600 h-2 rounded-full" :style="{ width: systemHealth.cpu + '%' }"></div>
                </div>
                <span class="text-sm font-medium text-gray-900">{{ systemHealth.cpu }}%</span>
              </div>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">RAM</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                  <div class="bg-green-600 h-2 rounded-full" :style="{ width: systemHealth.memory + '%' }"></div>
                </div>
                <span class="text-sm font-medium text-gray-900">{{ systemHealth.memory }}%</span>
              </div>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Disk</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                  <div class="bg-yellow-600 h-2 rounded-full" :style="{ width: systemHealth.disk + '%' }"></div>
                </div>
                <span class="text-sm font-medium text-gray-900">{{ systemHealth.disk }}%</span>
              </div>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">Network</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                  <div class="bg-purple-600 h-2 rounded-full" :style="{ width: systemHealth.network + '%' }"></div>
                </div>
                <span class="text-sm font-medium text-gray-900">{{ systemHealth.network }}%</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Recent Alerts -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">هشدارهای اخیر</h3>
          <div class="space-y-3">
            <div v-for="alert in recentAlerts" :key="alert.id" class="flex items-start">
              <div :class="getAlertIconClass(alert.level)" class="flex-shrink-0 mt-0.5">
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path v-if="alert.level === 'error'" fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                  <path v-else-if="alert.level === 'warning'" fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
                  <path v-else fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" />
                </svg>
              </div>
              <div class="mr-3 flex-1">
                <p class="text-sm text-gray-900">{{ alert.message }}</p>
                <p class="text-xs text-gray-500">{{ formatTime(alert.timestamp) }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Performance Score -->
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">امتیاز عملکرد</h3>
          <div class="text-center">
            <div class="relative inline-flex items-center justify-center">
              <svg class="w-24 h-24 transform -rotate-90">
                <circle
                  cx="48"
                  cy="48"
                  r="36"
                  stroke="currentColor"
                  stroke-width="8"
                  fill="transparent"
                  class="text-gray-200"
                />
                <circle
                  cx="48"
                  cy="48"
                  r="36"
                  stroke="currentColor"
                  stroke-width="8"
                  fill="transparent"
                  :stroke-dasharray="`${2 * Math.PI * 36}`"
                  :stroke-dashoffset="`${2 * Math.PI * 36 * (1 - performanceScore / 100)}`"
                  :class="getPerformanceScoreColor(performanceScore)"
                />
              </svg>
              <div class="absolute">
                <span class="text-2xl font-bold text-gray-900">{{ performanceScore }}</span>
                <span class="text-sm text-gray-500">/100</span>
              </div>
            </div>
            <p class="text-sm text-gray-600 mt-2">{{ getPerformanceScoreText(performanceScore) }}</p>
          </div>
        </div>
      </div>

      <!-- Detailed Metrics Table -->
      <div class="bg-white rounded-lg border border-gray-200 mt-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">متریک‌های تفصیلی</h3>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">متریک</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مقدار فعلی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">هدف</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">روند</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="metric in detailedMetrics" :key="metric.name">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ metric.name }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ metric.currentValue }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ metric.target }}</td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="getMetricStatusClass(metric.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                    {{ getMetricStatusText(metric.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  <span :class="getTrendClass(metric.trend)" class="flex items-center">
                    <svg v-if="metric.trend > 0" class="w-4 h-4 ml-1" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M5.293 7.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L6.707 7.707a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                    </svg>
                    <svg v-else class="w-4 h-4 ml-1" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M14.707 12.293a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 111.414-1.414L9 14.586V3a1 1 0 012 0v11.586l2.293-2.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                    {{ metric.trend > 0 ? '+' : '' }}{{ metric.trend }}{{ metric.unit }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
// Dynamic import Chart.js to reduce initial bundle size
let Chart = null

// Load Chart.js dynamically
const loadChartJS = async () => {
  if (!Chart) {
    const chartModule = await import('chart.js/auto')
    Chart = chartModule.default
  }
  return Chart
}

// Page meta
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Reactive data
const isRefreshing = ref(false)
const responseTimeChart = ref(null)
const messagesChart = ref(null)
let chartInstances = {
  responseTime: null,
  messages: null
}

// System status
const systemStatus = ref({
  status: 'healthy', // healthy, warning, error
  uptime: '99.8%',
  lastCheck: new Date()
})

// Key metrics
const metrics = ref({
  activeConnections: 0,
  connectionsTrend: 0,
  avgResponseTime: 0,
  responseTimeTrend: 0,
  messagesPerSecond: 0,
  messagesTrend: 0,
  errorRate: 0,
  errorRateTrend: 0,
  rateLimitedMessages: 0,
  rateLimitTrend: 0,
  websocketConnections: 0,
  websocketConnectionsTrend: 0
})

// System health
const systemHealth = ref({
  cpu: 0,
  memory: 0,
  disk: 0,
  network: 0
})

// Performance score
const performanceScore = ref(0)

// Recent alerts
const recentAlerts = ref([])

// Detailed metrics
const detailedMetrics = ref([])

// Monitoring control
const monitoringEnabled = ref(true)
const isToggling = ref(false)

// Chart data
const chartData = ref({
  responseTime: {
    labels: [],
    data: []
  },
  messages: {
    labels: [],
    data: []
  }
})

// API functions
const fetchMonitoringStatus = async () => {
  try {
    const response = await $fetch('/api/admin/monitoring/status')
    if (response.success) {
      monitoringEnabled.value = response.enabled
    }
  } catch (error) {
    console.error('خطا در دریافت وضعیت مانیتورینگ:', error)
  }
}

const toggleMonitoring = async () => {
  try {
    isToggling.value = true
    const response = await $fetch('/api/admin/monitoring/toggle', {
      method: 'POST',
      body: {
        enabled: !monitoringEnabled.value
      }
    })
    
    if (response.success) {
      monitoringEnabled.value = response.enabled
      // نمایش پیام موفقیت
      alert(response.message)
    }
  } catch (error) {
    console.error('خطا در تغییر وضعیت مانیتورینگ:', error)
    alert('خطا در تغییر وضعیت مانیتورینگ')
  } finally {
    isToggling.value = false
  }
}

const fetchRealTimeMetrics = async () => {
  try {
    const response = await $fetch('/api/admin/chat/metrics')
    if (response.success) {
      // Update metrics
      metrics.value.activeConnections = response.metrics.active_connections
      metrics.value.avgResponseTime = response.metrics.avg_response_time
      metrics.value.messagesPerSecond = response.metrics.messages_per_second
      metrics.value.errorRate = response.metrics.error_rate
      metrics.value.rateLimitedMessages = response.metrics.rate_limited_messages || 0
      metrics.value.rateLimitTrend = response.metrics.rate_limit_trend || 0
      metrics.value.websocketConnections = response.metrics.websocket_connections || 0
      metrics.value.websocketConnectionsTrend = response.metrics.websocket_connections_trend || 0
      
      // Update system health
      systemHealth.value.cpu = response.metrics.cpu_usage
      systemHealth.value.memory = response.metrics.memory_usage
      systemHealth.value.disk = response.metrics.disk_usage
      systemHealth.value.network = response.metrics.network_usage
      
      // Update performance score
      performanceScore.value = response.metrics.performance_score
      
      // Update system status
      systemStatus.value.status = response.health.status
      systemStatus.value.uptime = response.health.uptime
      systemStatus.value.lastCheck = new Date(response.health.last_check)
      
      // Update alerts
      recentAlerts.value = response.alerts.slice(0, 5)
    }
  } catch (error) {
    console.error('خطا در دریافت متریک‌های real-time:', error)
  }
}

const fetchPerformanceChart = async () => {
  try {
    const response = await $fetch('/api/admin/monitoring/chart?time_range=24h&metric=response_time')
    if (response.success) {
      chartData.value.responseTime.labels = response.labels
      chartData.value.responseTime.data = response.data
    }
    
    const messagesResponse = await $fetch('/api/admin/monitoring/chart?time_range=24h&metric=messages')
    if (messagesResponse.success) {
      chartData.value.messages.labels = messagesResponse.labels
      chartData.value.messages.data = messagesResponse.data
    }
    
    createCharts()
  } catch (error) {
    console.error('خطا در دریافت داده‌های نمودار:', error)
  }
}

const fetchDetailedMetrics = async () => {
  try {
    const response = await $fetch('/api/admin/monitoring/performance?time_range=24h&limit=1')
    if (response.success && response.data.length > 0) {
      const latestMetric = response.data[0]
      
      detailedMetrics.value = [
        {
          name: 'زمان پاسخ API',
          currentValue: `${latestMetric.avg_response_time.toFixed(0)}ms`,
          target: '< 200ms',
          status: latestMetric.avg_response_time < 200 ? 'excellent' : latestMetric.avg_response_time < 300 ? 'good' : 'warning',
          trend: 0,
          unit: 'ms'
        },
        {
          name: 'نرخ موفقیت',
          currentValue: `${(100 - latestMetric.error_rate).toFixed(1)}%`,
          target: '> 99%',
          status: latestMetric.error_rate < 1 ? 'excellent' : latestMetric.error_rate < 2 ? 'good' : 'warning',
          trend: 0,
          unit: '%'
        },
        {
          name: 'اتصالات همزمان',
          currentValue: latestMetric.active_connections.toString(),
          target: '< 200',
          status: latestMetric.active_connections < 150 ? 'excellent' : latestMetric.active_connections < 200 ? 'good' : 'warning',
          trend: 0,
          unit: ''
        },
        {
          name: 'حافظه استفاده شده',
          currentValue: `${latestMetric.memory_usage.toFixed(0)}%`,
          target: '< 80%',
          status: latestMetric.memory_usage < 70 ? 'excellent' : latestMetric.memory_usage < 80 ? 'good' : 'warning',
          trend: 0,
          unit: '%'
        },
        {
          name: 'CPU استفاده شده',
          currentValue: `${latestMetric.cpu_usage.toFixed(0)}%`,
          target: '< 70%',
          status: latestMetric.cpu_usage < 50 ? 'excellent' : latestMetric.cpu_usage < 70 ? 'good' : 'warning',
          trend: 0,
          unit: '%'
        },
        {
          name: 'پیام‌های محدود شده',
          currentValue: latestMetric.rate_limited_messages?.toString() || '0',
          target: '< 10',
          status: (latestMetric.rate_limited_messages || 0) < 5 ? 'excellent' : (latestMetric.rate_limited_messages || 0) < 10 ? 'good' : 'warning',
          trend: 0,
          unit: ''
        },
        {
          name: 'اتصالات WebSocket',
          currentValue: latestMetric.websocket_connections?.toString() || '0',
          target: '< 100',
          status: (latestMetric.websocket_connections || 0) < 50 ? 'excellent' : (latestMetric.websocket_connections || 0) < 100 ? 'good' : 'warning',
          trend: 0,
          unit: ''
        }
      ]
    }
  } catch (error) {
    console.error('خطا در دریافت متریک‌های تفصیلی:', error)
  }
}

// Methods
const getStatusClass = (status) => {
  const classes = {
    healthy: 'bg-green-100 text-green-800',
    warning: 'bg-yellow-100 text-yellow-800',
    error: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusDotClass = (status) => {
  const classes = {
    healthy: 'bg-green-400',
    warning: 'bg-yellow-400',
    error: 'bg-red-400'
  }
  return classes[status] || 'bg-gray-400'
}

const getStatusText = (status) => {
  const texts = {
    healthy: 'سالم',
    warning: 'هشدار',
    error: 'خطا'
  }
  return texts[status] || 'نامشخص'
}

const getTrendClass = (trend) => {
  if (trend > 0) return 'text-green-600'
  if (trend < 0) return 'text-red-600'
  return 'text-gray-600'
}

const getAlertIconClass = (level) => {
  const classes = {
    error: 'text-red-500',
    warning: 'text-yellow-500',
    info: 'text-blue-500'
  }
  return classes[level] || 'text-gray-500'
}

const getPerformanceScoreColor = (score) => {
  if (score >= 90) return 'text-green-500'
  if (score >= 70) return 'text-yellow-500'
  return 'text-red-500'
}

const getPerformanceScoreText = (score) => {
  if (score >= 90) return 'عالی'
  if (score >= 70) return 'خوب'
  return 'نیاز به بهبود'
}

const getMetricStatusClass = (status) => {
  const classes = {
    excellent: 'bg-green-100 text-green-800',
    good: 'bg-blue-100 text-blue-800',
    warning: 'bg-yellow-100 text-yellow-800',
    error: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getMetricStatusText = (status) => {
  const texts = {
    excellent: 'عالی',
    good: 'خوب',
    warning: 'هشدار',
    error: 'خطا'
  }
  return texts[status] || 'نامشخص'
}

const formatTime = (timestamp) => {
  const now = new Date()
  const alertTime = new Date(timestamp)
  const diff = now - alertTime
  const minutes = Math.floor(diff / (1000 * 60))
  const hours = Math.floor(diff / (1000 * 60 * 60))
  
  if (minutes < 60) {
    return `${minutes} دقیقه پیش`
  } else if (hours < 24) {
    return `${hours} ساعت پیش`
  } else {
    return alertTime.toLocaleDateString('fa-IR')
  }
}

const createCharts = async () => {
  // Load Chart.js only when needed
  const ChartInstance = await loadChartJS()
  
  // Response Time Chart
  if (chartInstances.responseTime) {
    chartInstances.responseTime.destroy()
  }
  
  chartInstances.responseTime = new ChartInstance(responseTimeChart.value, {
    type: 'line',
    data: {
      labels: chartData.value.responseTime.labels,
      datasets: [{
        label: 'زمان پاسخ (ms)',
        data: chartData.value.responseTime.data,
        borderColor: '#3B82F6',
        backgroundColor: 'rgba(59, 130, 246, 0.1)',
        fill: true,
        tension: 0.4
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          display: false
        }
      },
      scales: {
        y: {
          beginAtZero: true,
          grid: {
            color: 'rgba(0, 0, 0, 0.1)'
          }
        },
        x: {
          grid: {
            display: false
          }
        }
      }
    }
  })
  
  // Messages Chart
  if (chartInstances.messages) {
    chartInstances.messages.destroy()
  }
  
  chartInstances.messages = new ChartInstance(messagesChart.value, {
    type: 'bar',
    data: {
      labels: chartData.value.messages.labels,
      datasets: [{
        label: 'تعداد پیام‌ها',
        data: chartData.value.messages.data,
        backgroundColor: '#10B981',
        borderColor: '#10B981',
        borderWidth: 1
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          display: false
        }
      },
      scales: {
        y: {
          beginAtZero: true,
          grid: {
            color: 'rgba(0, 0, 0, 0.1)'
          }
        },
        x: {
          grid: {
            display: false
          }
        }
      }
    }
  })
}

const refreshData = async () => {
  isRefreshing.value = true
  
  try {
    await Promise.all([
      fetchRealTimeMetrics(),
      fetchPerformanceChart(),
      fetchDetailedMetrics()
    ])
  } catch (error) {
    console.error('خطا در بروزرسانی داده‌ها:', error)
  } finally {
    isRefreshing.value = false
  }
}

// Auto refresh every 30 seconds
let autoRefreshInterval = null

const startAutoRefresh = () => {
  autoRefreshInterval = setInterval(() => {
    fetchRealTimeMetrics()
  }, 30000)
}

const stopAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
    autoRefreshInterval = null
  }
}

// Lifecycle
onMounted(async () => {
  await Promise.all([
    fetchMonitoringStatus(),
    refreshData()
  ])
  startAutoRefresh()
})

onUnmounted(() => {
  stopAutoRefresh()
  if (chartInstances.responseTime) {
    chartInstances.responseTime.destroy()
  }
  if (chartInstances.messages) {
    chartInstances.messages.destroy()
  }
})
</script>

<style scoped>
.performance-monitoring {
  min-height: 100vh;
  background-color: #f9fafb;
}

/* Custom scrollbar */
.overflow-x-auto::-webkit-scrollbar {
  height: 6px;
}

.overflow-x-auto::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.overflow-x-auto::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.overflow-x-auto::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style> 
