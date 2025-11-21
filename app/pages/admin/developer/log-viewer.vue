<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مشاهده لاگ‌ها</h1>
        <p class="text-gray-600">ابزارهای مشاهده و تحلیل لاگ‌های سیستم</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">کل لاگ‌ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.totalLogs }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">خطاها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.errors }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">هشدارها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.warnings }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">اطلاعات</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.info }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
        <!-- Filters Sidebar -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">فیلترها</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <!-- Log Level Filter -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سطح لاگ</label>
                <div class="space-y-2">
                  <label class="flex items-center">
                    <input v-model="filters.levels" type="checkbox" value="error" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                    <span class="ml-2 text-sm text-red-600">خطا</span>
                  </label>
                  <label class="flex items-center">
                    <input v-model="filters.levels" type="checkbox" value="warning" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                    <span class="ml-2 text-sm text-yellow-600">هشدار</span>
                  </label>
                  <label class="flex items-center">
                    <input v-model="filters.levels" type="checkbox" value="info" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <span class="ml-2 text-sm text-blue-600">اطلاعات</span>
                  </label>
                  <label class="flex items-center">
                    <input v-model="filters.levels" type="checkbox" value="debug" class="rounded border-gray-300 text-gray-600 focus:ring-gray-500">
                    <span class="ml-2 text-sm text-gray-600">دیباگ</span>
                  </label>
                </div>
              </div>

              <!-- Date Range Filter -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
                <select v-model="filters.dateRange" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="1h">آخرین ساعت</option>
                  <option value="24h">آخرین 24 ساعت</option>
                  <option value="7d">آخرین 7 روز</option>
                  <option value="30d">آخرین 30 روز</option>
                  <option value="custom">سفارشی</option>
                </select>
              </div>

              <!-- Custom Date Range -->
              <div v-if="filters.dateRange === 'custom'" class="space-y-2">
                <input 
                  v-model="filters.startDate"
                  type="datetime-local"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <input 
                  v-model="filters.endDate"
                  type="datetime-local"
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <!-- Search Filter -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
                <input 
                  v-model="filters.search"
                  type="text"
                  placeholder="جستجو در لاگ‌ها..."
                  class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
              </div>

              <!-- Service Filter -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سرویس</label>
                <select v-model="filters.service" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="">همه سرویس‌ها</option>
                  <option value="api">API</option>
                  <option value="database">Database</option>
                  <option value="auth">Authentication</option>
                  <option value="payment">Payment</option>
                  <option value="email">Email</option>
                </select>
              </div>

              <!-- Actions -->
              <div class="space-y-2">
                <button 
                  class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="applyFilters"
                >
                  اعمال فیلتر
                </button>
                <button 
                  class="w-full bg-gray-500 hover:bg-gray-600 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="clearFilters"
                >
                  پاک کردن فیلترها
                </button>
                <button 
                  class="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="exportLogs"
                >
                  خروجی CSV
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Log Viewer -->
        <div class="lg:col-span-3 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">لاگ‌ها</h2>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                  @click="refreshLogs"
                >
                  بارگذاری مجدد
                </button>
                <button 
                  :class="[
                    'font-medium py-2 px-4 rounded-lg transition-colors',
                    autoRefresh ? 'bg-green-600 hover:bg-green-700 text-white' : 'bg-gray-500 hover:bg-gray-600 text-white'
                  ]"
                  @click="autoRefresh = !autoRefresh"
                >
                  {{ autoRefresh ? 'توقف' : 'اتوماتیک' }}
                </button>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-2">
              <div v-for="log in filteredLogs" :key="log.id" class="border rounded-lg p-6 hover:bg-gray-50">
                <div class="flex items-start justify-between">
                  <div class="flex items-start space-x-3 space-x-reverse">
                    <!-- Log Level Badge -->
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      log.level === 'error' ? 'bg-red-100 text-red-800' :
                      log.level === 'warning' ? 'bg-yellow-100 text-yellow-800' :
                      log.level === 'info' ? 'bg-blue-100 text-blue-800' :
                      'bg-gray-100 text-gray-800'
                    ]">
                      {{ log.level.toUpperCase() }}
                    </span>

                    <!-- Log Content -->
                    <div class="flex-1">
                      <p class="text-sm font-medium text-gray-900">{{ log.message }}</p>
                      <div class="mt-1 text-xs text-gray-500 space-y-1">
                        <p>زمان: {{ formatDate(log.timestamp) }}</p>
                        <p>سرویس: {{ log.service }}</p>
                        <p v-if="log.user">کاربر: {{ log.user }}</p>
                        <p v-if="log.ip">IP: {{ log.ip }}</p>
                      </div>
                    </div>
                  </div>

                  <!-- Actions -->
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="viewLogDetails(log)"
                    >
                      جزئیات
                    </button>
                    <button 
                      class="text-gray-600 hover:text-gray-800 text-sm font-medium"
                      @click="copyLog(log)"
                    >
                      کپی
                    </button>
                  </div>
                </div>

                <!-- Stack Trace (if error) -->
                <div v-if="log.level === 'error' && log.stackTrace" class="mt-3">
                  <details class="text-xs">
                    <summary class="cursor-pointer text-red-600 font-medium">نمایش Stack Trace</summary>
                    <pre class="mt-2 bg-red-50 border border-red-200 p-3 rounded text-red-800 overflow-x-auto">{{ log.stackTrace }}</pre>
                  </details>
                </div>
              </div>

              <!-- Empty State -->
              <div v-if="filteredLogs.length === 0" class="text-center py-8">
                <svg class="w-12 h-12 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
                <p class="text-gray-500">لاگی برای نمایش وجود ندارد</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Log Details Modal -->
      <div v-if="logDetailsModal.show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-4xl max-h-4xl overflow-auto">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-semibold text-gray-900">جزئیات لاگ</h3>
              <button class="text-gray-400 hover:text-gray-600" @click="logDetailsModal.show = false">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">پیام</label>
                <p class="text-gray-900">{{ logDetailsModal.log?.message }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سطح</label>
                <span
:class="[
                  'px-2 py-1 rounded text-xs font-medium',
                  logDetailsModal.log?.level === 'error' ? 'bg-red-100 text-red-800' :
                  logDetailsModal.log?.level === 'warning' ? 'bg-yellow-100 text-yellow-800' :
                  logDetailsModal.log?.level === 'info' ? 'bg-blue-100 text-blue-800' :
                  'bg-gray-100 text-gray-800'
                ]">
                  {{ logDetailsModal.log?.level?.toUpperCase() }}
                </span>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">زمان</label>
                <p class="text-gray-900">{{ formatDate(logDetailsModal.log?.timestamp) }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">سرویس</label>
                <p class="text-gray-900">{{ logDetailsModal.log?.service }}</p>
              </div>
              <div v-if="logDetailsModal.log?.stackTrace">
                <label class="block text-sm font-medium text-gray-700 mb-2">Stack Trace</label>
                <pre class="bg-gray-100 p-3 rounded text-sm overflow-x-auto">{{ logDetailsModal.log.stackTrace }}</pre>
              </div>
              <div v-if="logDetailsModal.log?.metadata">
                <label class="block text-sm font-medium text-gray-700 mb-2">متادیتا</label>
                <pre class="bg-gray-100 p-3 rounded text-sm overflow-x-auto">{{ JSON.stringify(logDetailsModal.log.metadata, null, 2) }}</pre>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted, reactive, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const logs = ref([])
const autoRefresh = ref(false)
const refreshInterval = ref(null)

const stats = reactive({
  totalLogs: 0,
  errors: 0,
  warnings: 0,
  info: 0
})

const filters = reactive({
  levels: ['error', 'warning', 'info', 'debug'],
  dateRange: '24h',
  startDate: '',
  endDate: '',
  search: '',
  service: ''
})

const logDetailsModal = reactive({
  show: false,
  log: null
})

// Computed
const filteredLogs = computed(() => {
  let filtered = logs.value

  // Filter by level
  if (filters.levels.length > 0) {
    filtered = filtered.filter(log => filters.levels.includes(log.level))
  }

  // Filter by search
  if (filters.search) {
    filtered = filtered.filter(log => 
      log.message.toLowerCase().includes(filters.search.toLowerCase()) ||
      log.service.toLowerCase().includes(filters.search.toLowerCase())
    )
  }

  // Filter by service
  if (filters.service) {
    filtered = filtered.filter(log => log.service === filters.service)
  }

  // Filter by date range
  if (filters.dateRange !== 'custom') {
    const now = new Date()
    const startDate = new Date()
    
    switch (filters.dateRange) {
      case '1h':
        startDate.setHours(now.getHours() - 1)
        break
      case '24h':
        startDate.setDate(now.getDate() - 1)
        break
      case '7d':
        startDate.setDate(now.getDate() - 7)
        break
      case '30d':
        startDate.setDate(now.getDate() - 30)
        break
    }
    
    filtered = filtered.filter(log => new Date(log.timestamp) >= startDate)
  } else if (filters.startDate && filters.endDate) {
    const start = new Date(filters.startDate)
    const end = new Date(filters.endDate)
    filtered = filtered.filter(log => {
      const logDate = new Date(log.timestamp)
      return logDate >= start && logDate <= end
    })
  }

  return filtered
})

// Methods
async function fetchLogs() {
  try {
    // Simulate API call
    const mockLogs = generateMockLogs()
    logs.value = mockLogs
    updateStats()
  } catch (error) {
    console.error('خطا در دریافت لاگ‌ها:', error)
  }
}

function generateMockLogs() {
  const levels = ['error', 'warning', 'info', 'debug']
  const services = ['api', 'database', 'auth', 'payment', 'email']
  const messages = [
    'User authentication failed',
    'Database connection established',
    'Payment processed successfully',
    'Email sent to user@example.com',
    'API rate limit exceeded',
    'Cache miss for key: user_profile_123',
    'File upload completed',
    'Background job started',
    'Memory usage high: 85%',
    'SSL certificate expires in 30 days'
  ]

  const logs = []
  for (let i = 0; i < 50; i++) {
    const level = levels[Math.floor(Math.random() * levels.length)]
    const service = services[Math.floor(Math.random() * services.length)]
    const message = messages[Math.floor(Math.random() * messages.length)]
    
    logs.push({
      id: Date.now() + i,
      level,
      message,
      service,
      timestamp: new Date(Date.now() - Math.random() * 7 * 24 * 60 * 60 * 1000).toISOString(),
      user: Math.random() > 0.5 ? `user${Math.floor(Math.random() * 1000)}` : null,
      ip: Math.random() > 0.3 ? `192.168.1.${Math.floor(Math.random() * 255)}` : null,
      stackTrace: level === 'error' ? generateStackTrace() : null,
      metadata: Math.random() > 0.7 ? { requestId: `req_${Math.random().toString(36).substr(2, 9)}` } : null
    })
  }

  return logs.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp))
}

function generateStackTrace() {
  return `Error: Something went wrong
    at processRequest (/app/src/handlers/user.js:45:12)
    at async handleUserAction (/app/src/controllers/user.js:23:8)
    at async router.get (/app/src/routes/user.js:15:4)
    at async dispatch (/app/src/middleware/auth.js:12:6)`
}

function updateStats() {
  stats.totalLogs = logs.value.length
  stats.errors = logs.value.filter(log => log.level === 'error').length
  stats.warnings = logs.value.filter(log => log.level === 'warning').length
  stats.info = logs.value.filter(log => log.level === 'info').length
}

function applyFilters() {
  // Filters are applied automatically via computed property

}

function clearFilters() {
  filters.levels = ['error', 'warning', 'info', 'debug']
  filters.dateRange = '24h'
  filters.startDate = ''
  filters.endDate = ''
  filters.search = ''
  filters.service = ''
}

function exportLogs() {
  const csvContent = generateCSV(filteredLogs.value)
  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
  const link = document.createElement('a')
  link.href = URL.createObjectURL(blob)
  link.download = `logs_${new Date().toISOString().split('T')[0]}.csv`
  link.click()
}

function generateCSV(logs) {
  const headers = ['Level', 'Message', 'Service', 'Timestamp', 'User', 'IP']
  const rows = logs.map(log => [
    log.level,
    log.message,
    log.service,
    log.timestamp,
    log.user || '',
    log.ip || ''
  ])
  
  return [headers, ...rows].map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
}

function refreshLogs() {
  fetchLogs()
}

function viewLogDetails(log) {
  logDetailsModal.log = log
  logDetailsModal.show = true
}

function copyLog(log) {
  const logText = `${log.level.toUpperCase()} - ${log.message} (${log.service}) - ${log.timestamp}`
  navigator.clipboard.writeText(logText).then(() => {
    alert('لاگ کپی شد')
  })
}

function formatDate(dateString) {
  return new Date(dateString).toLocaleString('fa-IR')
}

// Auto refresh
watch(autoRefresh, (newValue) => {
  if (newValue) {
    refreshInterval.value = setInterval(fetchLogs, 5000) // Refresh every 5 seconds
  } else {
    if (refreshInterval.value) {
      clearInterval(refreshInterval.value)
      refreshInterval.value = null
    }
  }
})

// Load initial data
onMounted(() => {
  fetchLogs()
})

onUnmounted(() => {
  if (refreshInterval.value) {
    clearInterval(refreshInterval.value)
  }
})
</script> 
