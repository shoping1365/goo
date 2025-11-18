<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">عملکرد سیستم</h1>
      <p class="text-gray-600">نظارت و تست عملکرد سرور و سیستم</p>
    </div>

    <!-- کارت‌های وضعیت سیستم -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">CPU Usage</p>
            <p class="text-2xl font-semibold text-gray-900">{{ systemStats.cpu }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Memory Usage</p>
            <p class="text-2xl font-semibold text-gray-900">{{ systemStats.memory }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Active Processes</p>
            <p class="text-2xl font-semibold text-gray-900">{{ systemStats.processes }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Response Time</p>
            <p class="text-2xl font-semibold text-gray-900">{{ systemStats.responseTime }}ms</p>
          </div>
        </div>
      </div>
    </div>

    <!-- بخش تنظیمات سیستم -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-8">
      <!-- تنظیمات هسته‌ها -->
      <div class="bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تنظیمات هسته‌ها</h2>
          <p class="text-sm text-gray-600 mt-1">مدیریت تعداد هسته‌های CPU</p>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تعداد هسته‌های فعال:</label>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input 
                  v-model.number="coreSettings.instances" 
                  type="number" 
                  min="1" 
                  max="32" 
                  class="flex-1 border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <button 
                  @click="updateCoreSettings" 
                  :disabled="updatingCores"
                  class="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  <span v-if="updatingCores">🔄</span>
                  <span v-else>💾</span>
                </button>
              </div>
            </div>
            
            <div class="text-sm text-gray-600">
              <p>هسته‌های فیزیکی: {{ cpuInfo.totalCores || 0 }}</p>
              <p>هسته‌های منطقی: {{ cpuInfo.logicalCores || 0 }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات حافظه -->
      <div class="bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تنظیمات حافظه</h2>
          <p class="text-sm text-gray-600 mt-1">مدیریت RAM سیستم</p>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر حافظه (MB):</label>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input 
                  v-model.number="memorySettings.maxMemory" 
                  type="number" 
                  min="100" 
                  max="10000" 
                  class="flex-1 border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <button 
                  @click="updateMemorySettings" 
                  :disabled="updatingMemory"
                  class="bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  <span v-if="updatingMemory">🔄</span>
                  <span v-else>💾</span>
                </button>
              </div>
            </div>
            
            <div class="text-sm text-gray-600">
              <p>حافظه کل: {{ memoryInfo.physical?.total || 0 }} MB</p>
              <p>حافظه آزاد: {{ memoryInfo.physical?.free || 0 }} MB</p>
            </div>
          </div>
        </div>
      </div>

      <!-- عملکرد هسته‌ها -->
      <div class="bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">عملکرد هسته‌ها</h2>
          <p class="text-sm text-gray-600 mt-1">نمایش درصد استفاده هر هسته</p>
        </div>
        <div class="p-6">
          <div class="space-y-3">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">CPU کلی:</span>
              <span class="text-lg font-bold" :class="getCPUColor(cpuInfo.perCoreUsage)">
                {{ cpuInfo.perCoreUsage || 0 }}%
              </span>
            </div>
            
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="bg-blue-600 h-2 rounded-full transition-all duration-300" 
                :style="{ width: (cpuInfo.perCoreUsage || 0) + '%' }"
              ></div>
            </div>
            
            <div class="grid grid-cols-2 gap-2">
              <button 
                @click="refreshCPUInfo" 
                :disabled="refreshingCPU"
                class="bg-purple-600 hover:bg-purple-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
              >
                <span v-if="refreshingCPU">🔄</span>
                <span v-else>🔄</span>
              </button>
              
              <button 
                @click="showCPUDetails" 
                :disabled="showingCPUDetails"
                class="bg-indigo-600 hover:bg-indigo-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
              >
                <span v-if="showingCPUDetails">🔍</span>
                <span v-else>🔍</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- بخش تست‌های بارگذاری -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- تست‌های PM2 -->
      <div class="bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تست‌های PM2 Cluster Mode</h2>
          <p class="text-sm text-gray-600 mt-1">بررسی عملکرد چند هسته‌ای و cluster mode</p>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div>
              <button 
                @click="runPM2Test" 
                :disabled="pm2TestRunning"
                class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
              >
                <span v-if="pm2TestRunning">🔄 در حال اجرا...</span>
                <span v-else>🚀 راه‌اندازی و تست PM2</span>
              </button>
            </div>
            
            <div>
              <button 
                @click="runClusterTest" 
                :disabled="clusterTestRunning"
                class="w-full bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
              >
                <span v-if="clusterTestRunning">🔄 در حال تست...</span>
                <span v-else>📊 تست عملکرد چند هسته‌ای</span>
              </button>
            </div>

            <div class="grid grid-cols-2 gap-6">
              <button 
                @click="showPM2Status" 
                class="bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium py-2 px-4 rounded-lg transition-colors"
              >
                📋 وضعیت PM2
              </button>
              
              <button 
                @click="showPM2Logs" 
                class="bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium py-2 px-4 rounded-lg transition-colors"
              >
                📝 نمایش Logs
              </button>
            </div>
          </div>

          <!-- نتایج PM2 -->
          <div v-if="pm2Results" class="mt-6 p-6 bg-gray-50 rounded-lg">
            <h3 class="font-medium text-gray-900 mb-2">نتایج PM2:</h3>
            <pre class="text-sm text-gray-700 whitespace-pre-wrap">{{ pm2Results }}</pre>
          </div>
        </div>
      </div>

      <!-- تست‌های بارگذاری -->
      <div class="bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">تست‌های بارگذاری</h2>
          <p class="text-sm text-gray-600 mt-1">بررسی عملکرد سرور تحت بار</p>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع تست:</label>
              <select v-model="selectedLoadTest" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option value="quick">تست سریع (1000 req/s - 30s)</option>
                <option value="full">تست کامل (1000 req/s - 60s)</option>
                <option value="custom">تست سفارشی</option>
              </select>
            </div>

            <!-- تنظیمات سفارشی -->
            <div v-if="selectedLoadTest === 'custom'" class="space-y-3">
              <div>
                <label class="block text-sm font-medium text-gray-700">درخواست در ثانیه:</label>
                <input v-model.number="customConfig.requestsPerSecond" type="number" min="1" max="10000" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">مدت زمان (ثانیه):</label>
                <input v-model.number="customConfig.duration" type="number" min="1" max="300" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">کاربران همزمان:</label>
                <input v-model.number="customConfig.concurrentUsers" type="number" min="1" max="1000" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
              </div>
            </div>

            <button 
              @click="runLoadTest" 
              :disabled="loadTestRunning"
              class="w-full bg-red-600 hover:bg-red-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
            >
              <span v-if="loadTestRunning">🔄 در حال تست بارگذاری...</span>
              <span v-else>⚡ شروع تست بارگذاری</span>
            </button>
          </div>

          <!-- نتایج تست بارگذاری -->
          <div v-if="loadTestResults" class="mt-6 p-6 bg-gray-50 rounded-lg">
            <h3 class="font-medium text-gray-900 mb-2">نتایج تست بارگذاری:</h3>
            <pre class="text-sm text-gray-700 whitespace-pre-wrap">{{ loadTestResults }}</pre>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار عملکرد -->
    <div class="mt-8 bg-white rounded-lg shadow">
      <div class="p-6 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900">نمودار عملکرد</h2>
        <p class="text-sm text-gray-600 mt-1">نمایش روند عملکرد سیستم در طول زمان</p>
      </div>
      <div class="p-6">
        <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
          <p class="text-gray-500">نمودار عملکرد در حال توسعه...</p>
        </div>
      </div>
    </div>

    <!-- لاگ‌های سیستم -->
    <div class="mt-8 bg-white rounded-lg shadow">
      <div class="p-6 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900">لاگ‌های سیستم</h2>
        <p class="text-sm text-gray-600 mt-1">آخرین رویدادها و خطاهای سیستم</p>
      </div>
      <div class="p-6">
        <div class="h-64 bg-gray-900 rounded-lg p-6 overflow-y-auto">
          <pre class="text-green-400 text-sm">{{ systemLogs }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
// تعریف interface ها
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
  result?: string
  logs?: string
}

interface CPUInfo {
  cores: number
  usage: number
  temperature?: number
  totalCores?: number
  logicalCores?: number
  perCoreUsage?: number
}

interface MemoryInfo {
  total: number
  used: number
  free: number
  usage: number
  physical?: {
    total: number
    free: number
  }
}

interface PM2Response {
  success?: boolean
  message?: string
  result?: string
}

interface LoadTestConfig {
  requestsPerSecond: number
  duration: number
  concurrentUsers: number
}

// Type declaration for Nuxt 4 auto-imported definePageMeta
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void

export default {
  name: 'SystemPerformancePage'
}
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// وضعیت سیستم
const systemStats = ref({
  cpu: 0,
  memory: 0,
  processes: 0,
  responseTime: 0
})

// اطلاعات CPU و حافظه
const cpuInfo = ref<CPUInfo | null>(null)
const memoryInfo = ref<MemoryInfo | null>(null)

// تنظیمات سیستم
const coreSettings = ref({
  instances: 4
})

const memorySettings = ref({
  maxMemory: 1024
})

// وضعیت تست‌ها
const pm2TestRunning = ref(false)
const clusterTestRunning = ref(false)
const loadTestRunning = ref(false)
const updatingCores = ref(false)
const updatingMemory = ref(false)
const refreshingCPU = ref(false)
const showingCPUDetails = ref(false)

// نتایج تست‌ها
const pm2Results = ref('')
const loadTestResults = ref('')

// تنظیمات تست بارگذاری
const selectedLoadTest = ref('quick')
const customConfig = ref({
  requestsPerSecond: 1000,
  duration: 30,
  concurrentUsers: 100
})

// لاگ‌های سیستم
const systemLogs = ref('در حال بارگذاری لاگ‌های سیستم...')

// تابع دریافت آمار سیستم
const fetchSystemStats = async () => {
  try {
    // شبیه‌سازی آمار سیستم
    systemStats.value = {
      cpu: Math.floor(Math.random() * 100),
      memory: Math.floor(Math.random() * 100),
      processes: Math.floor(Math.random() * 20) + 1,
      responseTime: Math.floor(Math.random() * 500) + 50
    }
  } catch (error) {
    console.error('خطا در دریافت آمار سیستم:', error)
  }
}

// تابع دریافت اطلاعات CPU
const fetchCPUInfo = async () => {
  try {
    const response: ApiResponse<CPUInfo> = await $fetch('/api/admin/system/cpu-cores', {
      method: 'GET'
    })
    if (response.success) {
      cpuInfo.value = response.data
    }
  } catch (error) {
    console.error('خطا در دریافت اطلاعات CPU:', error)
  }
}

// تابع دریافت اطلاعات حافظه
const fetchMemoryInfo = async () => {
  try {
    const response: ApiResponse<MemoryInfo> = await $fetch('/api/admin/system/memory-info', {
      method: 'GET'
    })
    if (response.success) {
      memoryInfo.value = response.data
    }
  } catch (error) {
    console.error('خطا در دریافت اطلاعات حافظه:', error)
  }
}

// تابع به‌روزرسانی تنظیمات هسته‌ها
const updateCoreSettings = async (): Promise<void> => {
  updatingCores.value = true
  try {
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/pm2-update', {
      method: 'POST',
      body: {
        instances: coreSettings.value.instances
      }
    })
    if (response.success && response.data) {
      pm2Results.value = (response.data.message || '') + '\n\n' + (response.data.result || '')
    } else if (response.message || response.result) {
      pm2Results.value = (response.message || '') + '\n\n' + (response.result || '')
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    pm2Results.value = `خطا در به‌روزرسانی تنظیمات هسته‌ها: ${errorMessage}`
  } finally {
    updatingCores.value = false
  }
}

// تابع به‌روزرسانی تنظیمات حافظه
const updateMemorySettings = async (): Promise<void> => {
  updatingMemory.value = true
  try {
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/pm2-update', {
      method: 'POST',
      body: {
        maxMemory: memorySettings.value.maxMemory
      }
    })
    if (response.success && response.data) {
      pm2Results.value = (response.data.message || '') + '\n\n' + (response.data.result || '')
    } else if (response.message || response.result) {
      pm2Results.value = (response.message || '') + '\n\n' + (response.result || '')
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    pm2Results.value = `خطا در به‌روزرسانی تنظیمات حافظه: ${errorMessage}`
  } finally {
    updatingMemory.value = false
  }
}

// تابع به‌روزرسانی اطلاعات CPU
const refreshCPUInfo = async () => {
  refreshingCPU.value = true
  try {
    await fetchCPUInfo()
  } catch (error) {
    console.error('خطا در به‌روزرسانی اطلاعات CPU:', error)
  } finally {
    refreshingCPU.value = false
  }
}

// تابع تعیین رنگ CPU بر اساس درصد استفاده
const getCPUColor = (usage: number | undefined): string => {
  const usageValue = usage || 0
  if (usageValue < 50) return 'text-green-600'
  if (usageValue < 80) return 'text-yellow-600'
  return 'text-red-600'
}

// تابع نمایش جزئیات CPU
const showCPUDetails = async (): Promise<void> => {
  showingCPUDetails.value = true
  try {
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/cpu-details', {
      method: 'POST'
    })
    if (response.success && response.data) {
      pm2Results.value = 'جزئیات CPU:\n\n' + (response.data.result || response.result || '')
    } else if (response.result) {
      pm2Results.value = 'جزئیات CPU:\n\n' + response.result
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    pm2Results.value = `خطا در دریافت جزئیات CPU: ${errorMessage}`
  } finally {
    showingCPUDetails.value = false
  }
}

// تابع اجرای تست PM2
const runPM2Test = async (): Promise<void> => {
  pm2TestRunning.value = true
  pm2Results.value = ''
  
  try {
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/pm2-test', {
      method: 'POST'
    })
    pm2Results.value = response.result || response.data?.result || ''
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    pm2Results.value = `خطا در اجرای تست PM2: ${errorMessage}`
  } finally {
    pm2TestRunning.value = false
  }
}

// تابع اجرای تست cluster
const runClusterTest = async (): Promise<void> => {
  clusterTestRunning.value = true
  pm2Results.value = ''
  
  try {
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/cluster-test', {
      method: 'POST'
    })
    pm2Results.value = response.result || response.data?.result || ''
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    pm2Results.value = `خطا در اجرای تست cluster: ${errorMessage}`
  } finally {
    clusterTestRunning.value = false
  }
}

// تابع نمایش وضعیت PM2
const showPM2Status = async (): Promise<void> => {
  try {
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/pm2-status', {
      method: 'GET'
    })
    pm2Results.value = response.result || response.data?.result || ''
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    pm2Results.value = `خطا در دریافت وضعیت PM2: ${errorMessage}`
  }
}

// تابع نمایش لاگ‌های PM2
const showPM2Logs = async (): Promise<void> => {
  try {
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/pm2-logs', {
      method: 'GET'
    })
    pm2Results.value = response.result || response.data?.result || ''
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    pm2Results.value = `خطا در دریافت لاگ‌های PM2: ${errorMessage}`
  }
}

// تابع اجرای تست بارگذاری
const runLoadTest = async (): Promise<void> => {
  loadTestRunning.value = true
  loadTestResults.value = ''
  
  try {
    let config: LoadTestConfig
    
    if (selectedLoadTest.value === 'quick') {
      config = { requestsPerSecond: 1000, duration: 30, concurrentUsers: 100 }
    } else if (selectedLoadTest.value === 'full') {
      config = { requestsPerSecond: 1000, duration: 60, concurrentUsers: 100 }
    } else {
      config = customConfig.value
    }
    
    const response = await $fetch<ApiResponse<PM2Response>>('/api/admin/system/load-test', {
      method: 'POST',
      body: config
    })
    loadTestResults.value = response.result || response.data?.result || ''
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    loadTestResults.value = `خطا در اجرای تست بارگذاری: ${errorMessage}`
  } finally {
    loadTestRunning.value = false
  }
}

// تابع دریافت لاگ‌های سیستم
const fetchSystemLogs = async (): Promise<void> => {
  try {
    const response = await $fetch<ApiResponse<{ logs?: string }>>('/api/admin/system/logs', {
      method: 'GET'
    })
    systemLogs.value = response.logs || response.data?.logs || 'در حال بارگذاری لاگ‌های سیستم...'
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    systemLogs.value = `خطا در دریافت لاگ‌های سیستم: ${errorMessage}`
  }
}

// بارگذاری اولیه
onMounted(() => {
  fetchSystemStats()
  fetchSystemLogs()
  fetchCPUInfo()
  fetchMemoryInfo()
  
  // به‌روزرسانی آمار هر 30 ثانیه
  setInterval(fetchSystemStats, 30000)
  
  // به‌روزرسانی لاگ‌ها هر 10 ثانیه
  setInterval(fetchSystemLogs, 10000)
  
  // به‌روزرسانی اطلاعات CPU و حافظه هر 60 ثانیه
  setInterval(() => {
    fetchCPUInfo()
    fetchMemoryInfo()
  }, 60000)
})
</script> 

