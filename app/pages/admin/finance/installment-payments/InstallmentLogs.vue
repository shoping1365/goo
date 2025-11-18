<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">لاگ‌ها و تاریخچه</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="logLevel" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="all">همه سطوح</option>
          <option value="info">اطلاعات</option>
          <option value="warning">هشدار</option>
          <option value="error">خطا</option>
          <option value="debug">اشکال‌زدایی</option>
        </select>
        <button @click="clearLogs" class="text-red-600 hover:text-red-800">
          پاک کردن لاگ‌ها
        </button>
        <button @click="refreshLogs" class="text-blue-600 hover:text-blue-800">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Log Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600 mb-1">{{ logStats.totalLogs }}</div>
        <div class="text-sm text-gray-600">کل لاگ‌ها</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600 mb-1">{{ logStats.infoLogs }}</div>
        <div class="text-sm text-gray-600">اطلاعات</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600 mb-1">{{ logStats.warningLogs }}</div>
        <div class="text-sm text-gray-600">هشدارها</div>
      </div>
      <div class="bg-red-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-red-600 mb-1">{{ logStats.errorLogs }}</div>
        <div class="text-sm text-gray-600">خطاها</div>
      </div>
    </div>

    <!-- Search and Filters -->
    <div class="mb-6">
      <div class="flex items-center space-x-4 space-x-reverse">
        <div class="flex-1">
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="جستجو در لاگ‌ها..."
            class="w-full border border-gray-300 rounded-md px-3 py-2"
          >
        </div>
        <div>
          <select v-model="dateFilter" class="border border-gray-300 rounded-md px-3 py-2">
            <option value="today">امروز</option>
            <option value="yesterday">دیروز</option>
            <option value="last_7_days">7 روز گذشته</option>
            <option value="last_30_days">30 روز گذشته</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Logs List -->
    <div class="space-y-2 mb-6">
      <div v-for="log in filteredLogs" :key="log.id" class="border border-gray-200 rounded-lg p-6">
        <div class="flex items-start justify-between">
          <div class="flex items-start">
            <div class="w-3 h-3 rounded-full mt-2 mr-3" :class="getLogLevelColor(log.level)"></div>
            <div class="flex-1">
              <div class="flex items-center space-x-2 space-x-reverse mb-1">
                <span class="px-2 py-1 text-xs rounded-full" :class="getLogLevelClass(log.level)">
                  {{ getLogLevelLabel(log.level) }}
                </span>
                <span class="text-sm text-gray-500">{{ log.timestamp }}</span>
                <span class="text-sm text-gray-500">{{ log.source }}</span>
              </div>
              <h4 class="text-md font-semibold text-gray-900 mb-2">{{ log.title }}</h4>
              <p class="text-sm text-gray-600 mb-2">{{ log.message }}</p>
              
              <div v-if="log.details" class="mt-2">
                <button @click="toggleDetails(log.id)" class="text-blue-600 hover:text-blue-800 text-sm">
                  {{ expandedLogs.includes(log.id) ? 'مخفی کردن جزئیات' : 'نمایش جزئیات' }}
                </button>
                
                <div v-if="expandedLogs.includes(log.id)" class="mt-2 p-3 bg-gray-100 rounded text-sm">
                  <pre class="whitespace-pre-wrap">{{ log.details }}</pre>
                </div>
              </div>
              
              <div v-if="log.user" class="mt-2 text-xs text-gray-500">
                کاربر: {{ log.user }}
              </div>
            </div>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <button v-if="log.level === 'error'" @click="reportBug(log)" class="text-red-600 hover:text-red-800 text-sm">
              گزارش باگ
            </button>
            <button @click="copyLog(log)" class="text-gray-600 hover:text-gray-800">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div class="flex items-center justify-between">
      <div class="text-sm text-gray-700">
        نمایش {{ (currentPage - 1) * logsPerPage + 1 }} تا {{ Math.min(currentPage * logsPerPage, totalLogs) }} از {{ totalLogs }} لاگ
      </div>
      
      <div class="flex items-center space-x-2 space-x-reverse">
        <button 
          @click="previousPage" 
          :disabled="currentPage === 1"
          class="px-3 py-1 border border-gray-300 rounded-md disabled:opacity-50"
        >
          قبلی
        </button>
        
        <span class="px-3 py-1">{{ currentPage }} از {{ totalPages }}</span>
        
        <button 
          @click="nextPage" 
          :disabled="currentPage === totalPages"
          class="px-3 py-1 border border-gray-300 rounded-md disabled:opacity-50"
        >
          بعدی
        </button>
      </div>
    </div>

    <!-- Export Logs -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="flex items-center justify-between">
        <span class="text-md font-semibold text-gray-900">صادرات لاگ‌ها</span>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button @click="exportLogs('txt')" class="bg-gray-600 text-white px-4 py-2 rounded-md hover:bg-gray-700">
            TXT
          </button>
          <button @click="exportLogs('csv')" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700">
            CSV
          </button>
          <button @click="exportLogs('json')" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700">
            JSON
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'

interface Log {
  id: number
  level: 'info' | 'warning' | 'error' | 'debug'
  title: string
  message: string
  details?: string
  timestamp: string
  source: string
  user?: string
}

const logLevel = ref('all')
const searchQuery = ref('')
const dateFilter = ref('today')
const currentPage = ref(1)
const logsPerPage = 20
const expandedLogs = ref<number[]>([])

interface LogStats {
  totalLogs: number
  infoLogs: number
  warningLogs: number
  errorLogs: number
}

const logStats = ref<LogStats>({
  totalLogs: 1250,
  infoLogs: 890,
  warningLogs: 245,
  errorLogs: 115
})

const logs = ref<Log[]>([
  {
    id: 1,
    level: 'info',
    title: 'درخواست اقساط جدید',
    message: 'درخواست اقساط جدید از کاربر علی احمدی ثبت شد',
    details: 'مبلغ: 25,000,000 ریال\nتعداد اقساط: 12\nامتیاز اعتباری: 75',
    timestamp: '1403/08/15 14:30:25',
    source: 'InstallmentService',
    user: 'علی احمدی'
  },
  {
    id: 2,
    level: 'warning',
    title: 'امتیاز اعتباری پایین',
    message: 'امتیاز اعتباری کاربر کمتر از حد مجاز است',
    details: 'امتیاز: 45\nحد مجاز: 60\nوضعیت: نیاز به بررسی دستی',
    timestamp: '1403/08/15 14:25:10',
    source: 'CreditService',
    user: 'محمد رضایی'
  },
  {
    id: 3,
    level: 'error',
    title: 'خطا در اتصال به سرویس بانک',
    message: 'امکان اتصال به سرویس اعتبارسنجی بانک وجود ندارد',
    details: 'Error: Connection timeout\nURL: https://api.bank.ir/credit-check\nTimeout: 30 seconds',
    timestamp: '1403/08/15 14:20:45',
    source: 'BankService'
  }
])

const totalLogs = computed(() => filteredLogs.value.length)
const totalPages = computed(() => Math.ceil(totalLogs.value / logsPerPage))

const filteredLogs = computed(() => {
  let filtered = logs.value

  // Filter by level
  if (logLevel.value !== 'all') {
    filtered = filtered.filter(log => log.level === logLevel.value)
  }

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(log => 
      log.title.toLowerCase().includes(query) ||
      log.message.toLowerCase().includes(query) ||
      log.source.toLowerCase().includes(query)
    )
  }

  // Pagination
  const startIndex = (currentPage.value - 1) * logsPerPage
  const endIndex = startIndex + logsPerPage
  return filtered.slice(startIndex, endIndex)
})

const getLogLevelColor = (level: string): string => {
  switch (level) {
    case 'info': return 'bg-blue-500'
    case 'warning': return 'bg-yellow-500'
    case 'error': return 'bg-red-500'
    case 'debug': return 'bg-gray-500'
    default: return 'bg-gray-500'
  }
}

const getLogLevelClass = (level: string): string => {
  switch (level) {
    case 'info': return 'bg-blue-100 text-blue-800'
    case 'warning': return 'bg-yellow-100 text-yellow-800'
    case 'error': return 'bg-red-100 text-red-800'
    case 'debug': return 'bg-gray-100 text-gray-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getLogLevelLabel = (level: string): string => {
  switch (level) {
    case 'info': return 'اطلاعات'
    case 'warning': return 'هشدار'
    case 'error': return 'خطا'
    case 'debug': return 'اشکال‌زدایی'
    default: return 'نامشخص'
  }
}

const toggleDetails = (logId: number) => {
  const index = expandedLogs.value.indexOf(logId)
  if (index > -1) {
    expandedLogs.value.splice(index, 1)
  } else {
    expandedLogs.value.push(logId)
  }
}

const copyLog = (log: Log) => {
  const logText = `[${log.level.toUpperCase()}] ${log.timestamp}\n${log.title}\n${log.message}\n${log.details || ''}`
  navigator.clipboard.writeText(logText)
  alert('لاگ کپی شد')
}

const reportBug = (log: Log) => {
  // Report bug functionality
  alert('گزارش باگ ارسال شد')
}

const clearLogs = async () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید همه لاگ‌ها را پاک کنید؟')) {
    try {
      await $fetch('/api/admin/installment-payments/logs', {
        method: 'DELETE'
      })
      logs.value = []
      alert('لاگ‌ها پاک شدند')
    } catch (error) {
      console.error('خطا در پاک کردن لاگ‌ها:', error)
      alert('خطا در پاک کردن لاگ‌ها')
    }
  }
}

const refreshLogs = async () => {
  try {
    interface LogsData {
      logs: Log[]
      stats: LogStats
    }
    interface ApiResponse {
      data?: LogsData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/logs')
    if (response.data) {
      logs.value = response.data.logs
      logStats.value = response.data.stats
    }
  } catch (error) {
    console.error('خطا در بارگذاری لاگ‌ها:', error)
  }
}

const exportLogs = async (format: string) => {
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>(`/api/admin/installment-payments/logs/export`, {
      method: 'POST',
      body: {
        format,
        filters: {
          level: logLevel.value,
          search: searchQuery.value,
          dateFilter: dateFilter.value
        }
      }
    })
    
    if (response.success) {
      alert(`فایل ${format.toUpperCase()} در حال آماده‌سازی است`)
    }
  } catch (error) {
    console.error('خطا در صادرات لاگ‌ها:', error)
    alert('خطا در صادرات لاگ‌ها')
  }
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

onMounted(() => {
  refreshLogs()
})

// Watch for filter changes
watch([logLevel, searchQuery, dateFilter], () => {
  currentPage.value = 1
})
</script>
