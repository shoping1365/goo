<template>
  <div class="space-y-6">
    <!-- آمار لاگ‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <!-- کل لاگ‌ها -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-slate-50 to-slate-100 border border-slate-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-slate-600">کل لاگ‌ها</dt>
                <dd class="text-base font-bold text-slate-800">{{ formatNumber(logStats.total) }}</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-slate-400 to-slate-500">
                📄
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📈
            <span class="text-green-600">+5.2%</span>
            <span class="text-slate-600 mr-2">نسبت به روز قبل</span>
          </div>
        </div>
      </div>

      <!-- خطاها -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-red-600">خطاها</dt>
                <dd class="text-base font-bold text-red-800">{{ formatNumber(logStats.errors) }}</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-red-400 to-red-500">
                ❌
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📉
            <span class="text-red-600">-12.3%</span>
            <span class="text-red-600 mr-2">نسبت به روز قبل</span>
          </div>
        </div>
      </div>

      <!-- هشدارها -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-amber-50 to-amber-100 border border-amber-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-amber-600">هشدارها</dt>
                <dd class="text-base font-bold text-amber-800">{{ formatNumber(logStats.warnings) }}</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-amber-400 to-amber-500">
                ⚠️
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📈
            <span class="text-green-600">+3.1%</span>
            <span class="text-amber-600 mr-2">نسبت به روز قبل</span>
          </div>
        </div>
      </div>

      <!-- موفق -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-emerald-50 to-emerald-100 border border-emerald-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-emerald-600">موفق</dt>
                <dd class="text-base font-bold text-emerald-800">{{ formatNumber(logStats.success) }}</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-emerald-400 to-emerald-500">
                ✅
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📈
            <span class="text-green-600">+8.7%</span>
            <span class="text-emerald-600 mr-2">نسبت به روز قبل</span>
          </div>
        </div>
      </div>
    </div>





    <!-- جدول لاگ‌ها -->
    <div class="bg-white rounded-xl shadow-sm">
      <div class="p-6 border-b border-gray-200">
        <!-- تایتل -->
        <div class="flex items-center justify-between mb-6">
          <div>
            <h4 class="text-xl font-bold text-gray-900">لاگ‌های سیستم</h4>
            <p class="text-sm text-gray-600 mt-1">مدیریت و مشاهده لاگ‌های سیستم و تراکنش‌های پرداخت</p>
          </div>
          <div class="flex items-center">
            <TemplateButton 
              v-if="hasActiveFilters"
              bg-gradient="bg-gradient-to-r from-red-500 to-pink-600" 
              text-color="text-white"
              border-color="border border-red-500"
              hover-class="hover:from-red-600 hover:to-pink-700"
              focus-class="focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
              size="medium"
              @click="clearAllFilters"
            >
              🗑️
              پاک کردن فیلترها
            </TemplateButton>
            <div class="ml-4"></div>
            <button
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-cyan-400 to-cyan-600 hover:from-cyan-500 hover:to-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
                @click="showFilters = !showFilters"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
              </svg>
              فیلترها
            </button>
            <div class="ml-4"></div>
            <button
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
                @click="exportLogs"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی
            </button>
          </div>
        </div>

        <!-- فیلترهای پیشرفته -->
        <div v-if="showFilters" class="bg-gray-50 p-6 rounded-lg mb-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع لاگ</label>
              <select v-model="selectedLogType" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="all">همه لاگ‌ها</option>
                <option value="error">خطاها</option>
                <option value="warning">هشدارها</option>
                <option value="info">اطلاعات</option>
                <option value="debug">دیباگ</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">درگاه پرداخت</label>
              <select v-model="selectedGateway" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="all">همه درگاه‌ها</option>
                <option v-for="gateway in gateways" :key="gateway.id" :value="gateway.id">
                  {{ gateway.name }}
                </option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
              <input 
                v-model="filters.startDate" 
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
              <input 
                v-model="filters.endDate" 
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
              <input 
                v-model="filters.search" 
                type="text" 
                placeholder="جستجو در لاگ‌ها..."
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
            </div>
          </div>
        </div>


      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درگاه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پیام</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">جزئیات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="log in filteredLogs" :key="log.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDateTime(log.timestamp) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
:class="[
                  'inline-flex px-2 py-1 text-xs font-semibold rounded-full',
                  log.level === 'error' ? 'bg-red-100 text-red-800' : '',
                  log.level === 'warning' ? 'bg-yellow-100 text-yellow-800' : '',
                  log.level === 'info' ? 'bg-blue-100 text-blue-800' : '',
                  log.level === 'debug' ? 'bg-gray-100 text-gray-800' : ''
                ]">
                  {{ getLevelText(log.level) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div v-if="log.gateway" :class="['w-6 h-6 rounded-full flex items-center justify-center text-white text-xs font-medium', log.gatewayColor]">
                    {{ log.gatewayIcon }}
                  </div>
                  <span class="mr-2 text-sm text-gray-900">{{ log.gateway || 'سیستم' }}</span>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900 max-w-xs truncate">
                {{ log.message }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span v-if="log.transactionId" class="text-blue-600">{{ log.transactionId }}</span>
                <span v-else>-</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900" @click="viewLogDetails(log)">
                  جزئیات
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="px-6 py-4 border-t border-gray-200">
        <div class="flex items-center justify-between">
          <div class="text-sm text-gray-700">
            نمایش {{ (currentPage - 1) * filters.limit + 1 }} تا {{ Math.min(currentPage * filters.limit, totalLogs) }} از {{ formatNumber(totalLogs) }} لاگ
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button 
              :disabled="currentPage === 1" 
              class="px-3 py-1 text-sm border border-gray-300 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed"
              @click="currentPage--"
            >
              قبلی
            </button>
            <span class="px-3 py-1 text-sm text-gray-700">{{ currentPage }} از {{ totalPages }}</span>
            <button 
              :disabled="currentPage === totalPages" 
              class="px-3 py-1 text-sm border border-gray-300 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed"
              @click="currentPage++"
            >
              بعدی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال جزئیات لاگ -->
    <div v-if="selectedLog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 max-w-2xl w-full mx-4 max-h-[80vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">جزئیات لاگ</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="selectedLog = null">
            ❌
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">زمان</label>
            <p class="text-sm text-gray-900">{{ formatDateTime(selectedLog.timestamp) }}</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">سطح</label>
            <span
:class="[
              'inline-flex px-2 py-1 text-xs font-semibold rounded-full',
              selectedLog.level === 'error' ? 'bg-red-100 text-red-800' : '',
              selectedLog.level === 'warning' ? 'bg-yellow-100 text-yellow-800' : '',
              selectedLog.level === 'info' ? 'bg-blue-100 text-blue-800' : '',
              selectedLog.level === 'debug' ? 'bg-gray-100 text-gray-800' : ''
            ]">
              {{ getLevelText(selectedLog.level) }}
            </span>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">درگاه</label>
            <p class="text-sm text-gray-900">{{ selectedLog.gateway || 'سیستم' }}</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">پیام</label>
            <p class="text-sm text-gray-900">{{ selectedLog.message }}</p>
          </div>
          
          <div v-if="selectedLog.transactionId">
            <label class="block text-sm font-medium text-gray-700 mb-1">شماره تراکنش</label>
            <p class="text-sm text-blue-600">{{ selectedLog.transactionId }}</p>
          </div>
          
          <div v-if="selectedLog.details">
            <label class="block text-sm font-medium text-gray-700 mb-1">جزئیات اضافی</label>
            <pre class="text-sm text-gray-900 bg-gray-50 p-3 rounded-lg overflow-x-auto">{{ JSON.stringify(selectedLog.details, null, 2) }}</pre>
          </div>
          
          <div v-if="selectedLog.stackTrace">
            <label class="block text-sm font-medium text-gray-700 mb-1">Stack Trace</label>
            <pre class="text-sm text-gray-900 bg-gray-50 p-3 rounded-lg overflow-x-auto max-h-40">{{ selectedLog.stackTrace }}</pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import TemplateButton from '~/components/common/TemplateButton.vue'

// متغیرهای reactive
const selectedLogType = ref('all')
const selectedGateway = ref('all')
const autoRefresh = ref(false)
const currentPage = ref(1)
const selectedLog = ref(null)
const showFilters = ref(false)

// فیلترها
const filters = ref({
  startDate: '',
  endDate: '',
  search: '',
  limit: 100
})

// آمار لاگ‌ها
const logStats = ref({
  total: 15420,
  errors: 125,
  warnings: 89,
  success: 15206
})

// درگاه‌ها
const gateways = ref([
  { id: 1, name: 'زرین‌پال' },
  { id: 2, name: 'نکست‌پی' },
  { id: 3, name: 'آیدی‌پی' },
  { id: 4, name: 'پی‌ایر' },
  { id: 5, name: 'کیف پول دیجیتال' },
  { id: 6, name: 'پی‌پال' },
  { id: 7, name: 'استرایپ' }
])

// لاگ‌های نمونه
const logs = ref([
  {
    id: 1,
    timestamp: '2024-01-15T10:30:00',
    level: 'info',
    gateway: 'زرین‌پال',
    gatewayIcon: 'زر',
    gatewayColor: 'bg-green-500',
    message: 'تراکنش با موفقیت انجام شد',
    transactionId: 'TXN-2024-001',
    details: { amount: 250000, status: 'success' }
  },
  {
    id: 2,
    timestamp: '2024-01-15T10:29:00',
    level: 'error',
    gateway: 'آیدی‌پی',
    gatewayIcon: 'آ',
    gatewayColor: 'bg-purple-500',
    message: 'خطا در اتصال به درگاه پرداخت',
    transactionId: 'TXN-2024-002',
    details: { errorCode: 'CONNECTION_TIMEOUT', retryCount: 3 },
    stackTrace: 'Error: Connection timeout\n    at PaymentGateway.connect()\n    at processPayment()'
  },
  {
    id: 3,
    timestamp: '2024-01-15T10:28:00',
    level: 'warning',
    gateway: 'نکست‌پی',
    gatewayIcon: 'ن',
    gatewayColor: 'bg-blue-500',
    message: 'مبلغ تراکنش کمتر از حداقل مجاز است',
    transactionId: 'TXN-2024-003',
    details: { amount: 5000, minAmount: 10000 }
  },
  {
    id: 4,
    timestamp: '2024-01-15T10:27:00',
    level: 'debug',
    gateway: 'کیف پول دیجیتال',
    gatewayIcon: 'ک',
    gatewayColor: 'bg-violet-500',
    message: 'درخواست پرداخت ارسال شد',
    transactionId: 'TXN-2024-004',
    details: { requestId: 'req_123456', amount: 75000 }
  },
  {
    id: 5,
    timestamp: '2024-01-15T10:26:00',
    level: 'info',
    gateway: 'پی‌پال',
    gatewayIcon: 'P',
    gatewayColor: 'bg-blue-600',
    message: 'تراکنش در انتظار تایید کاربر',
    transactionId: 'TXN-2024-005',
    details: { status: 'pending', redirectUrl: 'https://paypal.com/checkout' }
  },
  {
    id: 6,
    timestamp: '2024-01-15T10:25:00',
    level: 'error',
    gateway: null,
    message: 'خطا در ذخیره‌سازی لاگ تراکنش',
    details: { errorCode: 'DATABASE_ERROR', table: 'transaction_logs' },
    stackTrace: 'Error: Database connection failed\n    at saveLog()\n    at logTransaction()'
  },
  {
    id: 7,
    timestamp: '2024-01-15T10:24:00',
    level: 'warning',
    gateway: 'استرایپ',
    gatewayIcon: 'S',
    gatewayColor: 'bg-purple-600',
    message: 'کارت بانکی منقضی شده است',
    transactionId: 'TXN-2024-006',
    details: { cardLast4: '1234', expiryMonth: 12, expiryYear: 2023 }
  },
  {
    id: 8,
    timestamp: '2024-01-15T10:23:00',
    level: 'info',
    gateway: 'زرین‌پال',
    gatewayIcon: 'زر',
    gatewayColor: 'bg-green-500',
    message: 'بازگشت از درگاه پرداخت',
    transactionId: 'TXN-2024-007',
    details: { authority: 'A000000000000000000000000000000000000', status: 'OK' }
  }
])

// محاسبات
const filteredLogs = computed(() => {
  let filtered = logs.value

  // فیلتر بر اساس نوع لاگ
  if (selectedLogType.value !== 'all') {
    filtered = filtered.filter(log => log.level === selectedLogType.value)
  }

  // فیلتر بر اساس درگاه
  if (selectedGateway.value !== 'all') {
    const gateway = gateways.value.find(g => g.id === parseInt(selectedGateway.value))
    filtered = filtered.filter(log => log.gateway === gateway?.name)
  }

  // فیلتر بر اساس جستجو
  if (filters.value.search) {
    const search = String(filters.value.search).toLowerCase()
    filtered = filtered.filter(log => 
      log.message.toLowerCase().includes(search) ||
      (log.transactionId && String(log.transactionId).toLowerCase().includes(search)) ||
      (log.gateway && String(log.gateway).toLowerCase().includes(search))
    )
  }

  return filtered
})

const totalLogs = computed(() => filteredLogs.value.length)
const totalPages = computed(() => Math.ceil(totalLogs.value / filters.value.limit))

// بررسی وجود فیلترهای فعال
const hasActiveFilters = computed(() => {
  return selectedLogType.value !== 'all' ||
         selectedGateway.value !== 'all' ||
         filters.value.startDate ||
         filters.value.endDate ||
         filters.value.search
})

// توابع کمکی
const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

const formatDateTime = (dateString: string): string => {
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
}

const getLevelText = (level: string): string => {
  const levelMap = {
    error: 'خطا',
    warning: 'هشدار',
    info: 'اطلاعات',
    debug: 'دیباگ'
  }
  return levelMap[level] || level
}

// توابع عملیات
const refreshLogs = () => {

}

const exportLogs = () => {

}

const _clearLogs = () => {
  if (confirm('آیا از پاک کردن همه لاگ‌ها اطمینان دارید؟')) {

  }
}

const clearAllFilters = () => {
  selectedLogType.value = 'all'
  selectedGateway.value = 'all'
  filters.value.startDate = ''
  filters.value.endDate = ''
  filters.value.search = ''
}

interface Log {
  [key: string]: unknown
}

const viewLogDetails = (log: Log) => {
  selectedLog.value = log
}



// بروزرسانی خودکار
let autoRefreshInterval: ReturnType<typeof setInterval> | null = null

onMounted(() => {
  if (autoRefresh.value) {
    autoRefreshInterval = setInterval(refreshLogs, 30000) // هر 30 ثانیه
  }
})

// نظارت بر تغییرات autoRefresh
watch(autoRefresh, (newValue) => {
  // پاکسازی interval قبلی
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
    autoRefreshInterval = null
  }
  
  if (newValue) {
    autoRefreshInterval = setInterval(refreshLogs, 30000)
  }
})
</script> 
