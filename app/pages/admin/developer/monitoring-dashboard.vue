<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 mb-2">داشبورد مانیتورینگ</h1>
            <p class="text-gray-600">نظارت بر عملکرد سیستم، منابع و سرویس‌ها</p>
          </div>
          <div class="flex items-center gap-6">
            <!-- سوییچ روشن/خاموش مانیتورینگ -->
            <label class="flex items-center cursor-pointer select-none">
              <span class="ml-3 text-sm text-gray-600">مانیتورینگ</span>
              <input v-model="monitoringEnabled" type="checkbox" class="sr-only" @change="toggleMonitoring">
              <div :class="['w-12 h-7 rounded-full p-1 transition', monitoringEnabled ? 'bg-green-500' : 'bg-gray-300']">
                <div :class="['w-5 h-5 bg-white rounded-full transition', monitoringEnabled ? 'translate-x-5' : 'translate-x-0']"></div>
              </div>
            </label>
            <!-- سوییچ روشن/خاموش ثبت لاگ ترافیک -->
            <label class="flex items-center cursor-pointer select-none">
              <span class="ml-3 text-sm text-gray-600">ثبت لاگ ترافیک</span>
              <input v-model="trafficLoggingEnabled" type="checkbox" class="sr-only" @change="toggleTrafficLoggingMD">
              <div :class="['w-12 h-7 rounded-full p-1 transition', trafficLoggingEnabled ? 'bg-green-500' : 'bg-gray-300']">
                <div :class="['w-5 h-5 bg-white rounded-full transition', trafficLoggingEnabled ? 'translate-x-5' : 'translate-x-0']"></div>
              </div>
            </label>
            <button class="px-3 py-2 text-sm bg-white border border-gray-200 rounded-lg hover:bg-gray-50" @click="loadMonitoringStatus">به‌روزرسانی وضعیت</button>
          </div>
        </div>
      </div>

      <!-- System Overview Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">CPU</p>
              <p class="text-2xl font-bold text-gray-900">{{ systemStats.cpu }}%</p>
            </div>
            <div class="relative">
              <svg class="w-12 h-12 transform -rotate-90" viewBox="0 0 36 36">
                <path class="text-gray-200" stroke="currentColor" stroke-width="2" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
                <path class="text-blue-600" stroke="currentColor" stroke-width="2" stroke-dasharray="100, 100" stroke-dashoffset="25" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
              </svg>
              <span class="absolute inset-0 flex items-center justify-center text-xs font-medium text-gray-600">
                {{ systemStats.cpu }}%
              </span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">RAM</p>
              <p class="text-2xl font-bold text-gray-900">{{ systemStats.memory }}%</p>
            </div>
            <div class="relative">
              <svg class="w-12 h-12 transform -rotate-90" viewBox="0 0 36 36">
                <path class="text-gray-200" stroke="currentColor" stroke-width="2" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
                <path class="text-green-600" stroke="currentColor" stroke-width="2" stroke-dasharray="100, 100" stroke-dashoffset="35" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
              </svg>
              <span class="absolute inset-0 flex items-center justify-center text-xs font-medium text-gray-600">
                {{ systemStats.memory }}%
              </span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">دیسک</p>
              <p class="text-2xl font-bold text-gray-900">{{ systemStats.disk }}%</p>
            </div>
            <div class="relative">
              <svg class="w-12 h-12 transform -rotate-90" viewBox="0 0 36 36">
                <path class="text-gray-200" stroke="currentColor" stroke-width="2" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
                <path class="text-yellow-600" stroke="currentColor" stroke-width="2" stroke-dasharray="100, 100" stroke-dashoffset="45" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
              </svg>
              <span class="absolute inset-0 flex items-center justify-center text-xs font-medium text-gray-600">
                {{ systemStats.disk }}%
              </span>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600">شبکه</p>
              <p class="text-2xl font-bold text-gray-900">{{ systemStats.network }} MB/s</p>
            </div>
            <div class="relative">
              <svg class="w-12 h-12 transform -rotate-90" viewBox="0 0 36 36">
                <path class="text-gray-200" stroke="currentColor" stroke-width="2" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
                <path class="text-purple-600" stroke="currentColor" stroke-width="2" stroke-dasharray="100, 100" stroke-dashoffset="15" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
              </svg>
              <span class="absolute inset-0 flex items-center justify-center text-xs font-medium text-gray-600">
                {{ systemStats.network }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Charts Section -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
        <!-- CPU Usage Chart -->
        <div class="bg-white rounded-lg shadow">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">استفاده از CPU</h2>
          </div>
          <div class="p-6">
            <div class="h-64 flex items-end justify-between space-x-1">
              <div
v-for="(value, index) in cpuHistory" :key="index" 
                   class="flex-1 bg-blue-500 rounded-t" 
                   :style="{ height: `${value}%` }">
              </div>
            </div>
            <div class="mt-4 flex justify-between text-sm text-gray-600">
              <span>0%</span>
              <span>50%</span>
              <span>100%</span>
            </div>
          </div>
        </div>

        <!-- Memory Usage Chart -->
        <div class="bg-white rounded-lg shadow">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">استفاده از RAM</h2>
          </div>
          <div class="p-6">
            <div class="h-64 flex items-end justify-between space-x-1">
              <div
v-for="(value, index) in memoryHistory" :key="index" 
                   class="flex-1 bg-green-500 rounded-t" 
                   :style="{ height: `${value}%` }">
              </div>
            </div>
            <div class="mt-4 flex justify-between text-sm text-gray-600">
              <span>0%</span>
              <span>50%</span>
              <span>100%</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Services Status -->
      <div class="bg-white rounded-lg shadow mb-8">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">وضعیت سرویس‌ها</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
v-for="service in services" :key="service.name" 
                 class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
              <div class="flex items-center">
                <div
class="w-3 h-3 rounded-full mr-3" 
                     :class="service.status === 'running' ? 'bg-green-500' : 'bg-red-500'"></div>
                <div>
                  <p class="font-medium text-gray-900">{{ service.name }}</p>
                  <p class="text-sm text-gray-600">{{ service.description }}</p>
                </div>
              </div>
              <div class="text-right">
                <p
class="text-sm font-medium" 
                   :class="service.status === 'running' ? 'text-green-600' : 'text-red-600'">
                  {{ service.status === 'running' ? 'فعال' : 'غیرفعال' }}
                </p>
                <p class="text-xs text-gray-500">{{ service.uptime }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Performance Metrics -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
        <!-- Response Time -->
        <div class="bg-white rounded-lg shadow">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">زمان پاسخ‌دهی</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="metric in responseTimeMetrics" :key="metric.name" class="flex justify-between items-center">
                <span class="text-sm text-gray-600">{{ metric.name }}</span>
                <div class="flex items-center">
                  <div class="w-32 bg-gray-200 rounded-full h-2 mr-3">
                    <div class="bg-blue-600 h-2 rounded-full" :style="{ width: `${metric.percentage}%` }"></div>
                  </div>
                  <span class="text-sm font-medium text-gray-900">{{ metric.value }}ms</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Error Rate -->
        <div class="bg-white rounded-lg shadow">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-900">نرخ خطا</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="error in errorMetrics" :key="error.name" class="flex justify-between items-center">
                <span class="text-sm text-gray-600">{{ error.name }}</span>
                <div class="flex items-center">
                  <div class="w-32 bg-gray-200 rounded-full h-2 mr-3">
                    <div class="bg-red-600 h-2 rounded-full" :style="{ width: `${error.percentage}%` }"></div>
                  </div>
                  <span class="text-sm font-medium text-gray-900">{{ error.value }}%</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Alerts and Notifications -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
          <h2 class="text-lg font-semibold text-gray-900">هشدارها و اعلان‌ها</h2>
          <button class="text-blue-600 hover:text-blue-800 text-sm" @click="refreshAlerts">
            به‌روزرسانی
          </button>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div
v-for="alert in alerts" :key="alert.id" 
                 class="flex items-start p-6 border border-gray-200 rounded-lg"
                 :class="getAlertClass(alert.severity)">
              <div class="flex-shrink-0 mr-3">
                <svg class="w-5 h-5" :class="getAlertIconClass(alert.severity)" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
                </svg>
              </div>
              <div class="flex-1">
                <h3 class="text-sm font-medium" :class="getAlertTitleClass(alert.severity)">
                  {{ alert.title }}
                </h3>
                <p class="text-sm text-gray-600 mt-1">{{ alert.message }}</p>
                <p class="text-xs text-gray-500 mt-2">{{ formatDate(alert.timestamp) }}</p>
              </div>
              <button class="text-gray-400 hover:text-gray-600" @click="dismissAlert(alert.id)">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// وضعیت مانیتورینگ (آن/آف)
const monitoringEnabled = ref(true)
const { $toast } = useNuxtApp()
const loadMonitoringStatus = async () => {
  try {
    const resp = await $fetch('/api/admin/monitoring/status')
    monitoringEnabled.value = !!resp?.enabled
  } catch (e) {
    monitoringEnabled.value = true
  }
}

// --- Traffic logging toggle (مدیریت کلید traffic.logging.enabled) ---
const trafficLoggingEnabled = ref(true)
const loadTrafficLoggingMD = async () => {
  try {
    const setting = await $fetch('/api/admin/settings/traffic.logging.enabled')
    const val = (setting?.value || '').toString().toLowerCase()
    trafficLoggingEnabled.value = val === 'true' || val === '1' || val === 'yes'
  } catch (e) {
    trafficLoggingEnabled.value = true
  }
}
const toggleTrafficLoggingMD = async () => {
  try {
    const newVal = trafficLoggingEnabled.value ? 'true' : 'false'
    await $fetch('/api/admin/settings/traffic.logging.enabled', {
      method: 'PUT',
      body: { key: 'traffic.logging.enabled', value: newVal, category: 'traffic', type: 'boolean' }
    })
    $toast && $toast.success(trafficLoggingEnabled.value ? 'ثبت لاگ ترافیک فعال شد' : 'ثبت لاگ ترافیک غیرفعال شد')
  } catch (e) {
    $toast && $toast.error('خطا در تغییر وضعیت ثبت لاگ ترافیک')
    await loadTrafficLoggingMD()
  }
}
const toggleMonitoring = async () => {
  try {
    const desired = monitoringEnabled.value
    await $fetch('/api/admin/monitoring/toggle', { method: 'POST', body: { enabled: desired } })
    $toast && $toast.success(desired ? 'مانیتورینگ فعال شد' : 'مانیتورینگ غیرفعال شد')
  } catch (e) {
    // در صورت خطا وضعیت را به حالت قبل برگردان
    $toast && $toast.error('خطا در تغییر وضعیت مانیتورینگ')
    await loadMonitoringStatus()
  }
}

// Reactive data
const systemStats = ref({
  cpu: 45,
  memory: 65,
  disk: 55,
  network: 12.5
})

const cpuHistory = ref([45, 52, 38, 67, 43, 58, 49, 61, 44, 53, 47, 55])
const memoryHistory = ref([65, 68, 62, 71, 64, 69, 66, 73, 61, 67, 65, 70])

const services = ref([
  {
    name: 'Web Server',
    description: 'Nginx/Apache',
    status: 'running',
    uptime: '15 روز'
  },
  {
    name: 'Database',
    description: 'PostgreSQL',
    status: 'running',
    uptime: '12 روز'
  },
  {
    name: 'Redis Cache',
    description: 'Redis Server',
    status: 'running',
    uptime: '8 روز'
  },
  {
    name: 'Application',
    description: 'Node.js/Go',
    status: 'running',
    uptime: '3 روز'
  },
  {
    name: 'Email Service',
    description: 'SMTP Server',
    status: 'stopped',
    uptime: '0 روز'
  },
  {
    name: 'Backup Service',
    description: 'Automated Backup',
    status: 'running',
    uptime: '1 روز'
  }
])

const responseTimeMetrics = ref([
  { name: 'صفحه اصلی', value: 120, percentage: 60 },
  { name: 'API', value: 85, percentage: 42 },
  { name: 'دیتابیس', value: 45, percentage: 22 },
  { name: 'فایل‌ها', value: 200, percentage: 100 }
])

const errorMetrics = ref([
  { name: 'خطاهای 404', value: 2.1, percentage: 21 },
  { name: 'خطاهای 500', value: 0.5, percentage: 5 },
  { name: 'خطاهای دیتابیس', value: 0.8, percentage: 8 },
  { name: 'خطاهای شبکه', value: 1.2, percentage: 12 }
])

const alerts = ref([
  {
    id: 1,
    title: 'استفاده بالای CPU',
    message: 'استفاده از CPU به 85% رسیده است. بررسی کنید.',
    severity: 'warning',
    timestamp: new Date('2024-01-15T10:30:00')
  },
  {
    id: 2,
    title: 'سرویس ایمیل متوقف شد',
    message: 'سرویس SMTP متوقف شده است. بررسی کنید.',
    severity: 'error',
    timestamp: new Date('2024-01-15T09:15:00')
  },
  {
    id: 3,
    title: 'فضای دیسک کم',
    message: 'فقط 15% فضای دیسک باقی مانده است.',
    severity: 'warning',
    timestamp: new Date('2024-01-15T08:45:00')
  }
])

// Methods
const refreshAlerts = () => {
  // Simulate refreshing alerts
  console.log('Refreshing alerts...')
}

const dismissAlert = (id) => {
  const index = alerts.value.findIndex(alert => alert.id === id)
  if (index > -1) {
    alerts.value.splice(index, 1)
  }
}

const getAlertClass = (severity) => {
  const classes = {
    error: 'border-red-200 bg-red-50',
    warning: 'border-yellow-200 bg-yellow-50',
    info: 'border-blue-200 bg-blue-50'
  }
  return classes[severity] || 'border-gray-200 bg-gray-50'
}

const getAlertIconClass = (severity) => {
  const classes = {
    error: 'text-red-400',
    warning: 'text-yellow-400',
    info: 'text-blue-400'
  }
  return classes[severity] || 'text-gray-400'
}

const getAlertTitleClass = (severity) => {
  const classes = {
    error: 'text-red-800',
    warning: 'text-yellow-800',
    info: 'text-blue-800'
  }
  return classes[severity] || 'text-gray-800'
}

const formatDate = (date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

// Auto-refresh system stats
onMounted(() => {
  loadMonitoringStatus()
  loadTrafficLoggingMD()
  const interval = setInterval(() => {
    systemStats.value.cpu = Math.floor(Math.random() * 30) + 30
    systemStats.value.memory = Math.floor(Math.random() * 20) + 55
    systemStats.value.disk = Math.floor(Math.random() * 10) + 50
    systemStats.value.network = (Math.random() * 10 + 5).toFixed(1)
  }, 5000)

  onUnmounted(() => {
    clearInterval(interval)
  })
})
</script> 
