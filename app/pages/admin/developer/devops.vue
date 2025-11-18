<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">DevOps Dashboard</h1>
        <p class="text-gray-600">ابزارهای توسعه و مدیریت سیستم</p>
      </div>

      <!-- Stats Cards -->
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
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">Disk Usage</p>
              <p class="text-2xl font-semibold text-gray-900">{{ systemStats.disk }}%</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">Uptime</p>
              <p class="text-2xl font-semibold text-gray-900">{{ systemStats.uptime }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- System Monitoring -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">System Monitoring</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <div class="flex justify-between items-center mb-2">
                  <span class="text-sm font-medium text-gray-700">CPU Load</span>
                  <span class="text-sm text-gray-500">{{ systemStats.cpu }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2">
                  <div class="bg-blue-600 h-2 rounded-full" :style="{ width: systemStats.cpu + '%' }"></div>
                </div>
              </div>
              
              <div>
                <div class="flex justify-between items-center mb-2">
                  <span class="text-sm font-medium text-gray-700">Memory Usage</span>
                  <span class="text-sm text-gray-500">{{ systemStats.memory }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2">
                  <div class="bg-green-600 h-2 rounded-full" :style="{ width: systemStats.memory + '%' }"></div>
                </div>
              </div>
              
              <div>
                <div class="flex justify-between items-center mb-2">
                  <span class="text-sm font-medium text-gray-700">Disk Space</span>
                  <span class="text-sm text-gray-500">{{ systemStats.disk }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2">
                  <div class="bg-yellow-600 h-2 rounded-full" :style="{ width: systemStats.disk + '%' }"></div>
                </div>
              </div>
            </div>
            
            <button @click="refreshStats" class="mt-4 w-full bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
              بروزرسانی آمار
            </button>
          </div>
        </div>

        <!-- Log Viewer -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">System Logs</h2>
          </div>
          <div class="p-6">
            <div class="mb-4">
              <select v-model="selectedLogLevel" class="w-full p-2 border border-gray-300 rounded-lg">
                <option value="all">همه لاگ‌ها</option>
                <option value="error">خطاها</option>
                <option value="warning">هشدارها</option>
                <option value="info">اطلاعات</option>
              </select>
            </div>
            <div class="bg-gray-900 text-green-400 p-6 rounded-lg h-64 overflow-y-auto font-mono text-sm">
              <div v-for="log in filteredLogs" :key="log.id" class="mb-1">
                <span class="text-gray-500">[{{ log.timestamp }}]</span>
                <span :class="getLogLevelClass(log.level)">[{{ log.level.toUpperCase() }}]</span>
                <span class="text-white">{{ log.message }}</span>
              </div>
            </div>
            <button @click="clearLogs" class="mt-4 w-full bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors">
              پاک کردن لاگ‌ها
            </button>
          </div>
        </div>

        <!-- Database Management -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">Database Management</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
                <div>
                  <h3 class="font-medium text-gray-900">Database Status</h3>
                  <p class="text-sm text-gray-600">{{ dbStatus.connected ? 'متصل' : 'قطع' }}</p>
                </div>
                <div :class="dbStatus.connected ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-3 py-1 rounded-full text-sm font-medium">
                  {{ dbStatus.connected ? 'فعال' : 'غیرفعال' }}
                </div>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <button @click="backupDatabase" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                  پشتیبان‌گیری
                </button>
                <button @click="optimizeDatabase" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
                  بهینه‌سازی
                </button>
              </div>
              
              <div class="mt-4 p-6 bg-yellow-50 border border-yellow-200 rounded-lg">
                <h4 class="font-medium text-yellow-800 mb-2">آمار دیتابیس</h4>
                <div class="text-sm text-yellow-700">
                  <p>تعداد جداول: {{ dbStats.tables }}</p>
                  <p>حجم کل: {{ dbStats.size }}</p>
                  <p>تعداد رکوردها: {{ dbStats.records }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Cache Management -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">Cache Management</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
                <div>
                  <h3 class="font-medium text-gray-900">Cache Status</h3>
                  <p class="text-sm text-gray-600">{{ cacheStatus.enabled ? 'فعال' : 'غیرفعال' }}</p>
                </div>
                <div :class="cacheStatus.enabled ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-3 py-1 rounded-full text-sm font-medium">
                  {{ cacheStatus.enabled ? 'فعال' : 'غیرفعال' }}
                </div>
              </div>
              
              <div class="grid grid-cols-2 gap-6">
                <button @click="clearCache" class="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors">
                  پاک کردن کش
                </button>
                <button @click="warmCache" class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors">
                  گرم کردن کش
                </button>
              </div>
              
              <div class="mt-4 p-6 bg-blue-50 border border-blue-200 rounded-lg">
                <h4 class="font-medium text-blue-800 mb-2">آمار کش</h4>
                <div class="text-sm text-blue-700">
                  <p>Hit Rate: {{ cacheStats.hitRate }}%</p>
                  <p>Miss Rate: {{ cacheStats.missRate }}%</p>
                  <p>حجم کش: {{ cacheStats.size }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Security Scanner -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">Security Scanner</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
                <div>
                  <h3 class="font-medium text-gray-900">Last Scan</h3>
                  <p class="text-sm text-gray-600">{{ securityScan.lastScan }}</p>
                </div>
                <div :class="securityScan.status === 'safe' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-3 py-1 rounded-full text-sm font-medium">
                  {{ securityScan.status === 'safe' ? 'امن' : 'نقض امنیت' }}
                </div>
              </div>
              
              <button @click="runSecurityScan" class="w-full bg-orange-600 text-white px-4 py-2 rounded-lg hover:bg-orange-700 transition-colors">
                اجرای اسکن امنیتی
              </button>
              
              <div class="mt-4 p-6 bg-red-50 border border-red-200 rounded-lg">
                <h4 class="font-medium text-red-800 mb-2">نقض‌های امنیتی</h4>
                <div class="text-sm text-red-700">
                  <p v-for="vulnerability in securityScan.vulnerabilities" :key="vulnerability.id" class="mb-1">
                    • {{ vulnerability.description }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Performance Monitor -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">Performance Monitor</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div class="grid grid-cols-2 gap-6">
                <div class="text-center p-6 bg-gray-50 rounded-lg">
                  <div class="text-2xl font-bold text-blue-600">{{ performanceStats.responseTime }}ms</div>
                  <div class="text-sm text-gray-600">Response Time</div>
                </div>
                <div class="text-center p-6 bg-gray-50 rounded-lg">
                  <div class="text-2xl font-bold text-green-600">{{ performanceStats.requestsPerSecond }}</div>
                  <div class="text-sm text-gray-600">RPS</div>
                </div>
              </div>
              
              <div class="p-6 bg-gray-50 rounded-lg">
                <h4 class="font-medium text-gray-900 mb-2">Slow Queries</h4>
                <div class="text-sm text-gray-600">
                  <p v-for="query in performanceStats.slowQueries" :key="query.id" class="mb-1">
                    {{ query.sql }} ({{ query.duration }}ms)
                  </p>
                </div>
              </div>
              
              <button @click="optimizePerformance" class="w-full bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
                بهینه‌سازی عملکرد
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">Quick Actions</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
            <button @click="restartServices" class="p-6 bg-blue-100 text-blue-800 rounded-lg hover:bg-blue-200 transition-colors">
              <svg class="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              <span class="block text-sm font-medium">Restart Services</span>
            </button>
            
            <button @click="checkUpdates" class="p-6 bg-green-100 text-green-800 rounded-lg hover:bg-green-200 transition-colors">
              <svg class="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"></path>
              </svg>
              <span class="block text-sm font-medium">Check Updates</span>
            </button>
            
            <button @click="exportLogs" class="p-6 bg-yellow-100 text-yellow-800 rounded-lg hover:bg-yellow-200 transition-colors">
              <svg class="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <span class="block text-sm font-medium">Export Logs</span>
            </button>
            
            <button @click="systemHealth" class="p-6 bg-purple-100 text-purple-800 rounded-lg hover:bg-purple-200 transition-colors">
              <svg class="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <span class="block text-sm font-medium">System Health</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// تعریف متغیرهای reactive
const systemStats = ref({
  cpu: 45,
  memory: 62,
  disk: 78,
  uptime: '15d 8h 32m'
})

const selectedLogLevel = ref('all')
const logs = ref([
  { id: 1, timestamp: '2024-01-15 10:30:15', level: 'info', message: 'System started successfully' },
  { id: 2, timestamp: '2024-01-15 10:31:22', level: 'warning', message: 'High memory usage detected' },
  { id: 3, timestamp: '2024-01-15 10:32:45', level: 'error', message: 'Database connection timeout' },
  { id: 4, timestamp: '2024-01-15 10:33:12', level: 'info', message: 'Cache cleared successfully' },
  { id: 5, timestamp: '2024-01-15 10:34:08', level: 'error', message: 'API endpoint /api/users failed' }
])

const dbStatus = ref({
  connected: true
})

const dbStats = ref({
  tables: 24,
  size: '2.4 GB',
  records: '1,234,567'
})

const cacheStatus = ref({
  enabled: true
})

const cacheStats = ref({
  hitRate: 85,
  missRate: 15,
  size: '512 MB'
})

const securityScan = ref({
  lastScan: '2024-01-15 09:00:00',
  status: 'safe',
  vulnerabilities: [
    { id: 1, description: 'SQL Injection vulnerability in user input' },
    { id: 2, description: 'XSS vulnerability in comment system' }
  ]
})

const performanceStats = ref({
  responseTime: 245,
  requestsPerSecond: 1250,
  slowQueries: [
    { id: 1, sql: 'SELECT * FROM users WHERE...', duration: 1200 },
    { id: 2, sql: 'JOIN products ON...', duration: 890 }
  ]
})

// محاسبه لاگ‌های فیلتر شده
const filteredLogs = computed(() => {
  if (selectedLogLevel.value === 'all') {
    return logs.value
  }
  return logs.value.filter(log => log.level === selectedLogLevel.value)
})

// توابع
const refreshStats = () => {
  // شبیه‌سازی بروزرسانی آمار
  systemStats.value.cpu = Math.floor(Math.random() * 100)
  systemStats.value.memory = Math.floor(Math.random() * 100)
  systemStats.value.disk = Math.floor(Math.random() * 100)
}

const clearLogs = () => {
  logs.value = []
}

const backupDatabase = () => {
  // شبیه‌سازی پشتیبان‌گیری
  console.log('Database backup started...')
}

const optimizeDatabase = () => {
  // شبیه‌سازی بهینه‌سازی
  console.log('Database optimization started...')
}

const clearCache = () => {
  // شبیه‌سازی پاک کردن کش
  console.log('Cache cleared...')
}

const warmCache = () => {
  // شبیه‌سازی گرم کردن کش
  console.log('Cache warming started...')
}

const runSecurityScan = () => {
  // شبیه‌سازی اسکن امنیتی
  console.log('Security scan started...')
}

const optimizePerformance = () => {
  // شبیه‌سازی بهینه‌سازی عملکرد
  console.log('Performance optimization started...')
}

const restartServices = () => {
  // شبیه‌سازی راه‌اندازی مجدد سرویس‌ها
  console.log('Services restart initiated...')
}

const checkUpdates = () => {
  // شبیه‌سازی بررسی بروزرسانی‌ها
  console.log('Checking for updates...')
}

const exportLogs = () => {
  // شبیه‌سازی خروجی لاگ‌ها
  console.log('Exporting logs...')
}

const systemHealth = () => {
  // شبیه‌سازی بررسی سلامت سیستم
  console.log('System health check started...')
}

const getLogLevelClass = (level) => {
  switch (level) {
    case 'error':
      return 'text-red-400'
    case 'warning':
      return 'text-yellow-400'
    case 'info':
      return 'text-blue-400'
    default:
      return 'text-gray-400'
  }
}

// بروزرسانی خودکار آمار هر 30 ثانیه
onMounted(() => {
  setInterval(() => {
    refreshStats()
  }, 30000)
})

// تعریف متا اطلاعات صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()
</script> 
