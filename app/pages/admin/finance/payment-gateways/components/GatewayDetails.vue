<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div :class="`w-12 h-12 ${gateway.color} rounded-lg flex items-center justify-center ml-4`">
            <span class="text-white font-bold text-lg">{{ gateway.icon }}</span>
          </div>
          <div>
            <h2 class="text-lg font-semibold text-gray-900">{{ gateway.name }}</h2>
            <p class="text-gray-600">{{ gateway.englishName }} - {{ gateway.country }}</p>
          </div>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <span :class="`px-3 py-1 rounded-full text-sm font-medium ${getStatusBadgeClass(gateway.status)}`">
            {{ getStatusName(gateway.status) }}
          </span>
          <button @click="refreshData" :disabled="loading" class="p-2 text-gray-400 hover:text-gray-600 transition-colors">
            <span v-if="loading" class="i-heroicons-arrow-path animate-spin"></span>
            <span v-else class="i-heroicons-arrow-path"></span>
          </button>
        </div>
      </div>
    </div>

    <!-- تب‌های جزئیات -->
    <div class="border-b border-gray-200">
      <div class="flex">
        <button v-for="tab in detailTabs" :key="tab.value" @click="activeTab = tab.value"
          :class="['px-6 py-3 font-medium text-sm focus:outline-none', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <div class="p-6">
      <!-- اطلاعات کلی -->
      <div v-if="activeTab === 'overview'" class="space-y-6">
        <!-- آمار کلی -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="bg-blue-50 rounded-lg p-6">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
                <span class="i-heroicons-arrow-trending-up text-blue-600"></span>
              </div>
              <div>
                <p class="text-sm text-gray-600">تراکنش‌های امروز</p>
                <p class="text-xl font-bold text-gray-900">{{ formatNumber(gateway.todayTransactions) }}</p>
              </div>
            </div>
          </div>
          <div class="bg-green-50 rounded-lg p-6">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center ml-3">
                <span class="i-heroicons-check-circle text-green-600"></span>
              </div>
              <div>
                <p class="text-sm text-gray-600">نرخ موفقیت</p>
                <p class="text-xl font-bold text-gray-900">{{ gatewayStats.successRate }}%</p>
              </div>
            </div>
          </div>
          <div class="bg-yellow-50 rounded-lg p-6">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-yellow-100 rounded-lg flex items-center justify-center ml-3">
                <span class="i-heroicons-currency-dollar text-yellow-600"></span>
              </div>
              <div>
                <p class="text-sm text-gray-600">درآمد امروز</p>
                <p class="text-xl font-bold text-gray-900">{{ formatPrice(gatewayStats.todayRevenue) }}</p>
              </div>
            </div>
          </div>
          <div class="bg-purple-50 rounded-lg p-6">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
                <span class="i-heroicons-clock text-purple-600"></span>
              </div>
              <div>
                <p class="text-sm text-gray-600">متوسط زمان</p>
                <p class="text-xl font-bold text-gray-900">{{ gatewayStats.avgResponseTime }}s</p>
              </div>
            </div>
          </div>
        </div>

        <!-- اطلاعات درگاه -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات درگاه</h3>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="text-gray-600">نوع درگاه:</span>
                <span class="font-medium">{{ getTypeName(gateway.type) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">کارمزد:</span>
                <span class="font-medium">{{ gateway.fee }}%</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">حداقل مبلغ:</span>
                <span class="font-medium">{{ formatPrice(gateway.minAmount) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">حداکثر مبلغ:</span>
                <span class="font-medium">{{ formatPrice(gateway.maxAmount) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">نسخه API:</span>
                <span class="font-medium">{{ gateway.apiVersion }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">تاریخ ایجاد:</span>
                <span class="font-medium">{{ formatDate(gateway.createdAt) }}</span>
              </div>
            </div>
          </div>

          <div class="bg-gray-50 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت سیستم</h3>
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-gray-600">اتصال API:</span>
                <span :class="`px-2 py-1 rounded-full text-xs font-medium ${gatewayStats.apiStatus === 'online' ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'}`">
                  {{ gatewayStats.apiStatus === 'online' ? 'آنلاین' : 'آفلاین' }}
                </span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-gray-600">حالت تست:</span>
                <span :class="`px-2 py-1 rounded-full text-xs font-medium ${gateway.testMode ? 'bg-yellow-100 text-yellow-700' : 'bg-gray-100 text-gray-700'}`">
                  {{ gateway.testMode ? 'فعال' : 'غیرفعال' }}
                </span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-gray-600">آخرین بروزرسانی:</span>
                <span class="font-medium">{{ formatDate(gateway.updatedAt) }}</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-gray-600">زمان انقضا:</span>
                <span class="font-medium">{{ gateway.transactionExpiry || 30 }} دقیقه</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تراکنش‌ها -->
      <div v-if="activeTab === 'transactions'" class="space-y-6">
        <!-- فیلترهای تراکنش -->
        <div class="flex flex-wrap gap-6">
          <select v-model="transactionFilters.status" @change="loadTransactions" class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه وضعیت‌ها</option>
            <option value="success">موفق</option>
            <option value="failed">ناموفق</option>
            <option value="pending">در انتظار</option>
            <option value="cancelled">لغو شده</option>
          </select>
          <input v-model="transactionFilters.search" @input="loadTransactions" type="text" placeholder="جستجو در شماره تراکنش..." class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          <button @click="exportTransactions" class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors">
            <span class="i-heroicons-arrow-down-tray ml-2"></span>
            خروجی
          </button>
        </div>

        <!-- جدول تراکنش‌ها -->
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره تراکنش</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="transaction in transactions" :key="transaction.id">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ transaction.id }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatPrice(transaction.amount) }}</td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="`px-2 py-1 rounded-full text-xs font-medium ${getTransactionStatusClass(transaction.status)}`">
                    {{ getTransactionStatusName(transaction.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(transaction.createdAt) }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button @click="viewTransaction(transaction)" class="text-blue-600 hover:text-blue-900">مشاهده</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- صفحه‌بندی -->
        <div class="flex items-center justify-between">
          <div class="text-sm text-gray-700">
            نمایش {{ pagination.from }} تا {{ pagination.to }} از {{ pagination.total }} نتیجه
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button @click="previousPage" :disabled="pagination.currentPage === 1" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50">
              قبلی
            </button>
            <button @click="nextPage" :disabled="pagination.currentPage === pagination.lastPage" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50">
              بعدی
            </button>
          </div>
        </div>
      </div>

      <!-- لاگ‌ها -->
      <div v-if="activeTab === 'logs'" class="space-y-6">
        <!-- فیلترهای لاگ -->
        <div class="flex flex-wrap gap-6">
          <select v-model="logFilters.level" @change="loadLogs" class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه سطوح</option>
            <option value="error">خطا</option>
            <option value="warning">هشدار</option>
            <option value="info">اطلاعات</option>
            <option value="debug">دیباگ</option>
          </select>
          <input v-model="logFilters.search" @input="loadLogs" type="text" placeholder="جستجو در لاگ‌ها..." class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          <button @click="clearLogs" class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors">
            <span class="i-heroicons-trash ml-2"></span>
            پاک کردن لاگ‌ها
          </button>
        </div>

        <!-- لیست لاگ‌ها -->
        <div class="space-y-2 max-h-96 overflow-y-auto">
          <div v-for="log in logs" :key="log.id" :class="`p-6 rounded-lg border ${getLogLevelClass(log.level)}`">
            <div class="flex items-start justify-between">
              <div class="flex-1">
                <div class="flex items-center space-x-2 space-x-reverse mb-2">
                  <span :class="`px-2 py-1 rounded-full text-xs font-medium ${getLogLevelBadgeClass(log.level)}`">
                    {{ getLogLevelName(log.level) }}
                  </span>
                  <span class="text-sm text-gray-500">{{ formatDate(log.createdAt) }}</span>
                </div>
                <p class="text-sm text-gray-900">{{ log.message }}</p>
                <p v-if="log.details" class="text-xs text-gray-600 mt-1">{{ log.details }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تست -->
      <div v-if="activeTab === 'test'" class="space-y-6">
        <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
          <div class="flex">
            <div class="flex-shrink-0">
              <span class="i-heroicons-exclamation-triangle text-yellow-400"></span>
            </div>
            <div class="mr-3">
              <h3 class="text-sm font-medium text-yellow-800">حالت تست</h3>
              <div class="mt-2 text-sm text-yellow-700">
                <p>در حالت تست، تراکنش‌ها واقعی نیستند و برای تست عملکرد درگاه استفاده می‌شوند.</p>
              </div>
            </div>
          </div>
        </div>

        <!-- فرم تست -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ تست (تومان)</label>
            <input v-model.number="testAmount" type="number" min="1000" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شماره کارت تست</label>
            <input v-model="testCard" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
        </div>

        <div class="flex space-x-3 space-x-reverse">
          <button @click="testPayment" :disabled="testing" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50">
            <span v-if="testing" class="i-heroicons-arrow-path animate-spin ml-2"></span>
            <span v-else class="i-heroicons-beaker ml-2"></span>
            {{ testing ? 'در حال تست...' : 'تست پرداخت' }}
          </button>
          <button @click="testVerification" :disabled="testing" class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors disabled:opacity-50">
            <span v-if="testing" class="i-heroicons-arrow-path animate-spin ml-2"></span>
            <span v-else class="i-heroicons-check ml-2"></span>
            {{ testing ? 'در حال تست...' : 'تست تایید' }}
          </button>
        </div>

        <!-- نتایج تست -->
        <div v-if="testResults.length > 0" class="space-y-4">
          <h3 class="text-lg font-semibold text-gray-900">نتایج تست</h3>
          <div v-for="result in testResults" :key="result.id" :class="`p-6 rounded-lg border ${result.success ? 'border-green-200 bg-green-50' : 'border-red-200 bg-red-50'}`">
            <div class="flex items-start justify-between">
              <div>
                <h4 class="font-medium text-gray-900">{{ result.title }}</h4>
                <p class="text-sm text-gray-600 mt-1">{{ result.message }}</p>
                <p v-if="result.details" class="text-xs text-gray-500 mt-2">{{ result.details }}</p>
              </div>
              <span :class="`px-2 py-1 rounded-full text-xs font-medium ${result.success ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'}`">
                {{ result.success ? 'موفق' : 'ناموفق' }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// تعریف interface ها
interface PaymentGateway {
  id: number
  name: string
  englishName: string
  type: 'iranian' | 'international' | 'digital'
  status: 'active' | 'inactive' | 'maintenance'
  fee: number
  minAmount: number
  maxAmount: number
  todayTransactions: number
  color: string
  icon: string
  country: string
  description: string
  apiVersion: string
  createdAt: string
  updatedAt: string
  testMode?: boolean
  transactionExpiry?: number
}

interface GatewayStats {
  successRate: number
  todayRevenue: number
  avgResponseTime: number
  apiStatus: 'online' | 'offline'
}

interface Transaction {
  id: string
  amount: number
  status: 'success' | 'failed' | 'pending' | 'cancelled'
  createdAt: string
}

interface Log {
  id: number
  level: 'error' | 'warning' | 'info' | 'debug'
  message: string
  details?: string
  createdAt: string
}

interface TestResult {
  id: number
  title: string
  message: string
  details?: string
  success: boolean
}

interface Pagination {
  currentPage: number
  lastPage: number
  total: number
  from: number
  to: number
}

// Props
const props = defineProps<{
  gateway: PaymentGateway
}>()

// متغیرهای reactive
const activeTab = ref('overview')
const loading = ref(false)
const testing = ref(false)

// تب‌های جزئیات
const detailTabs = [
  { value: 'overview', label: 'نمای کلی' },
  { value: 'transactions', label: 'تراکنش‌ها' },
  { value: 'logs', label: 'لاگ‌ها' },
  { value: 'test', label: 'تست' }
]

// آمار درگاه
const gatewayStats = ref<GatewayStats>({
  successRate: 98.5,
  todayRevenue: 450000000,
  avgResponseTime: 2.3,
  apiStatus: 'online'
})

// فیلترهای تراکنش
const transactionFilters = ref({
  status: '',
  search: ''
})

// تراکنش‌ها
const transactions = ref<Transaction[]>([
  { id: 'TXN001', amount: 500000, status: 'success', createdAt: '2024-01-15T10:30:00Z' },
  { id: 'TXN002', amount: 750000, status: 'pending', createdAt: '2024-01-15T10:25:00Z' },
  { id: 'TXN003', amount: 300000, status: 'failed', createdAt: '2024-01-15T10:20:00Z' }
])

// صفحه‌بندی
const pagination = ref<Pagination>({
  currentPage: 1,
  lastPage: 5,
  total: 125,
  from: 1,
  to: 25
})

// فیلترهای لاگ
const logFilters = ref({
  level: '',
  search: ''
})

// لاگ‌ها
const logs = ref<Log[]>([
  { id: 1, level: 'info', message: 'درگاه با موفقیت متصل شد', createdAt: '2024-01-15T10:30:00Z' },
  { id: 2, level: 'warning', message: 'زمان پاسخ API بیش از حد انتظار', createdAt: '2024-01-15T10:25:00Z' },
  { id: 3, level: 'error', message: 'خطا در تایید تراکنش', details: 'کد خطا: 1001', createdAt: '2024-01-15T10:20:00Z' }
])

// تست
const testAmount = ref(10000)
const testCard = ref('6037991234567890')
const testResults = ref<TestResult[]>([])

// توابع کمکی
const getStatusName = (status: string): string => {
  switch (status) {
    case 'active': return 'فعال'
    case 'inactive': return 'غیرفعال'
    case 'maintenance': return 'نگهداری'
    default: return status
  }
}

const getStatusBadgeClass = (status: string): string => {
  switch (status) {
    case 'active': return 'bg-green-100 text-green-700'
    case 'inactive': return 'bg-red-100 text-red-700'
    case 'maintenance': return 'bg-yellow-100 text-yellow-700'
    default: return 'bg-gray-100 text-gray-700'
  }
}

const getTypeName = (type: string): string => {
  switch (type) {
    case 'iranian': return 'ایرانی'
    case 'international': return 'خارجی'
    case 'digital': return 'دیجیتال'
    default: return type
  }
}

const formatNumber = (num: number): string => {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(0) + 'K'
  }
  return new Intl.NumberFormat('fa-IR').format(num)
}

const formatPrice = (price: number): string => {
  if (price >= 1000000000) {
    return (price / 1000000000).toFixed(1) + ' میلیارد'
  } else if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + ' میلیون'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + ' هزار'
  }
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatDate = (date: string): string => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const getTransactionStatusName = (status: string): string => {
  switch (status) {
    case 'success': return 'موفق'
    case 'failed': return 'ناموفق'
    case 'pending': return 'در انتظار'
    case 'cancelled': return 'لغو شده'
    default: return status
  }
}

const getTransactionStatusClass = (status: string): string => {
  switch (status) {
    case 'success': return 'bg-green-100 text-green-700'
    case 'failed': return 'bg-red-100 text-red-700'
    case 'pending': return 'bg-yellow-100 text-yellow-700'
    case 'cancelled': return 'bg-gray-100 text-gray-700'
    default: return 'bg-gray-100 text-gray-700'
  }
}

const getLogLevelName = (level: string): string => {
  switch (level) {
    case 'error': return 'خطا'
    case 'warning': return 'هشدار'
    case 'info': return 'اطلاعات'
    case 'debug': return 'دیباگ'
    default: return level
  }
}

const getLogLevelBadgeClass = (level: string): string => {
  switch (level) {
    case 'error': return 'bg-red-100 text-red-700'
    case 'warning': return 'bg-yellow-100 text-yellow-700'
    case 'info': return 'bg-blue-100 text-blue-700'
    case 'debug': return 'bg-gray-100 text-gray-700'
    default: return 'bg-gray-100 text-gray-700'
  }
}

const getLogLevelClass = (level: string): string => {
  switch (level) {
    case 'error': return 'border-red-200 bg-red-50'
    case 'warning': return 'border-yellow-200 bg-yellow-50'
    case 'info': return 'border-blue-200 bg-blue-50'
    case 'debug': return 'border-gray-200 bg-gray-50'
    default: return 'border-gray-200 bg-gray-50'
  }
}

// توابع عملیات
const refreshData = async () => {
  loading.value = true
  try {
    // TODO: فراخوانی API برای به‌روزرسانی داده‌ها
    await new Promise(resolve => setTimeout(resolve, 1000))
    console.log('داده‌ها به‌روزرسانی شد')
  } catch (error) {
    console.error('خطا در به‌روزرسانی داده‌ها:', error)
  } finally {
    loading.value = false
  }
}

const loadTransactions = async () => {
  try {
    // TODO: فراخوانی API برای بارگذاری تراکنش‌ها
    console.log('بارگذاری تراکنش‌ها...')
  } catch (error) {
    console.error('خطا در بارگذاری تراکنش‌ها:', error)
  }
}

const loadLogs = async () => {
  try {
    // TODO: فراخوانی API برای بارگذاری لاگ‌ها
    console.log('بارگذاری لاگ‌ها...')
  } catch (error) {
    console.error('خطا در بارگذاری لاگ‌ها:', error)
  }
}

const exportTransactions = () => {
  try {
    // TODO: خروجی تراکنش‌ها
    console.log('خروجی تراکنش‌ها...')
  } catch (error) {
    console.error('خطا در خروجی تراکنش‌ها:', error)
  }
}

const clearLogs = async () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید تمام لاگ‌ها را پاک کنید؟')) {
    try {
      // TODO: پاک کردن لاگ‌ها
      console.log('پاک کردن لاگ‌ها...')
    } catch (error) {
      console.error('خطا در پاک کردن لاگ‌ها:', error)
    }
  }
}

const viewTransaction = (transaction: Transaction) => {
  // TODO: نمایش جزئیات تراکنش
  console.log('مشاهده تراکنش:', transaction)
}

const previousPage = () => {
  if (pagination.value.currentPage > 1) {
    pagination.value.currentPage--
    loadTransactions()
  }
}

const nextPage = () => {
  if (pagination.value.currentPage < pagination.value.lastPage) {
    pagination.value.currentPage++
    loadTransactions()
  }
}

const testPayment = async () => {
  testing.value = true
  try {
    // TODO: تست پرداخت
    await new Promise(resolve => setTimeout(resolve, 2000))
    testResults.value.unshift({
      id: Date.now(),
      title: 'تست پرداخت',
      message: 'درخواست پرداخت با موفقیت ارسال شد',
      success: true
    })
  } catch (error) {
    testResults.value.unshift({
      id: Date.now(),
      title: 'تست پرداخت',
      message: 'خطا در ارسال درخواست پرداخت',
      details: error.message,
      success: false
    })
  } finally {
    testing.value = false
  }
}

const testVerification = async () => {
  testing.value = true
  try {
    // TODO: تست تایید
    await new Promise(resolve => setTimeout(resolve, 2000))
    testResults.value.unshift({
      id: Date.now(),
      title: 'تست تایید',
      message: 'تایید تراکنش با موفقیت انجام شد',
      success: true
    })
  } catch (error) {
    testResults.value.unshift({
      id: Date.now(),
      title: 'تست تایید',
      message: 'خطا در تایید تراکنش',
      details: error.message,
      success: false
    })
  } finally {
    testing.value = false
  }
}

// بارگذاری داده‌ها
onMounted(() => {
  loadTransactions()
  loadLogs()
})
</script>

<!--
  کامپوننت جزئیات و مدیریت هر درگاه پرداخت
  شامل:
  1. نمای کلی با آمار و اطلاعات درگاه
  2. مدیریت تراکنش‌ها با فیلتر و صفحه‌بندی
  3. مشاهده لاگ‌های سیستم
  4. تست عملکرد درگاه
  5. عملیات مختلف مدیریتی
  6. طراحی ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
--> 
