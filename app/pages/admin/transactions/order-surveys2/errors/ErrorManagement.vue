<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">مدیریت خطاهای ارسال</h3>
        <p class="text-gray-600 text-sm">بررسی و حل مشکلات ارسال SMS</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button 
          :disabled="retryingAll"
          class="px-4 py-2 bg-orange-600 hover:bg-orange-700 disabled:bg-gray-400 text-white rounded-lg text-sm transition-colors flex items-center space-x-2 space-x-reverse"
          @click="retryAllFailed"
        >
          <svg v-if="retryingAll" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          <span>{{ retryingAll ? 'در حال ارسال مجدد...' : 'ارسال مجدد همه' }}</span>
        </button>
        
        <button 
          class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700 transition-colors"
          @click="exportErrorReport"
        >
          خروجی گزارش
        </button>
      </div>
    </div>

    <!-- Error Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-red-800">{{ errorStats.totalFailed }}</div>
            <div class="text-sm text-red-600">کل خطاها</div>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-yellow-800">{{ errorStats.retryable }}</div>
            <div class="text-sm text-yellow-600">قابل ارسال مجدد</div>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-green-800">{{ errorStats.resolved }}</div>
            <div class="text-sm text-green-600">حل شده</div>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <div class="text-2xl font-bold text-blue-800">{{ errorStats.successRate }}%</div>
            <div class="text-sm text-blue-600">نرخ موفقیت</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Error Filters -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h4 class="text-md font-semibold text-gray-800 mb-3">فیلتر خطاها</h4>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع خطا</label>
          <select 
            v-model="filters.errorType"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه انواع</option>
            <option value="invalid_number">شماره نامعتبر</option>
            <option value="service_unavailable">سرویس در دسترس نیست</option>
            <option value="quota_exceeded">محدودیت ارسال</option>
            <option value="network_error">خطای شبکه</option>
            <option value="timeout">زمان انتظار</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select 
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="pending">در انتظار</option>
            <option value="retrying">در حال تلاش مجدد</option>
            <option value="resolved">حل شده</option>
            <option value="permanent_failure">خطای دائمی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">تعداد تلاش</label>
          <select 
            v-model="filters.retryCount"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="0">0 بار</option>
            <option value="1">1 بار</option>
            <option value="2">2 بار</option>
            <option value="3">3 بار یا بیشتر</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ</label>
          <input 
            v-model="filters.date"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
        </div>
      </div>
    </div>

    <!-- Failed SMS Table -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                شماره سفارش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مشتری
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نوع خطا
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                پیام خطا
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تعداد تلاش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                آخرین تلاش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr 
              v-for="sms in filteredFailedSMS" 
              :key="sms.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">#{{ sms.orderNumber }}</div>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <div class="h-8 w-8 rounded-full bg-gradient-to-r from-red-400 to-pink-500 flex items-center justify-center">
                      <span class="text-white font-medium text-xs">{{ getInitials(sms.customerName) }}</span>
                    </div>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ sms.customerName }}</div>
                    <div class="text-sm text-gray-500">{{ sms.phoneNumber }}</div>
                  </div>
                </div>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <span 
                  :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    getErrorTypeClass(sms.errorType)
                  ]"
                >
                  {{ getErrorTypeText(sms.errorType) }}
                </span>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900 max-w-xs truncate" :title="sms.errorMessage">
                  {{ sms.errorMessage }}
                </div>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <span class="text-sm text-gray-900">{{ sms.retryCount }}</span>
                  <span class="text-xs text-gray-500 mr-1">/ {{ maxRetries }}</span>
                </div>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDateTime(sms.lastAttempt) }}
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap">
                <span 
                  :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    getStatusClass(sms.status)
                  ]"
                >
                  {{ getStatusText(sms.status) }}
                </span>
              </td>
              
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    v-if="canRetry(sms)"
                    :disabled="sms.retrying"
                    class="text-blue-600 hover:text-blue-900 disabled:text-gray-400 transition-colors"
                    title="ارسال مجدد"
                    @click="resendSMS(sms.id)"
                  >
                    <svg v-if="sms.retrying" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                    </svg>
                  </button>
                  
                  <button 
                    class="text-green-600 hover:text-green-900 transition-colors"
                    title="تحلیل خطا"
                    @click="analyzeError(sms.id)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
                    </svg>
                  </button>
                  
                  <button 
                    class="text-purple-600 hover:text-purple-900 transition-colors"
                    title="مشاهده جزئیات"
                    @click="viewDetails(sms.id)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Pagination -->
    <div class="flex items-center justify-between">
      <div class="text-sm text-gray-700">
        نمایش {{ startIndex + 1 }} تا {{ endIndex }} از {{ totalErrors }} خطا
      </div>
      <div class="flex space-x-2 space-x-reverse">
        <button 
          :disabled="currentPage === 1"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
          @click="previousPage"
        >
          قبلی
        </button>
        <span class="px-3 py-1 text-sm text-gray-700">
          صفحه {{ currentPage }} از {{ totalPages }}
        </span>
        <button 
          :disabled="currentPage >= totalPages"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
          @click="nextPage"
        >
          بعدی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue'

interface FailedSMS {
  id: number
  orderNumber: string
  customerName: string
  phoneNumber: string
  errorType: string
  errorMessage: string
  retryCount: number
  lastAttempt: string
  status: string
  retrying?: boolean
}

const props = defineProps<{
  failedSMS: FailedSMS[]
}>()

const emit = defineEmits<{
  'resend': [smsId: number]
  'retry-all': []
  'analyze-error': [errorId: number]
}>()

// Reactive data
const retryingAll = ref(false)
const currentPage = ref(1)
const itemsPerPage = ref(20)
const maxRetries = 3

const filters = reactive({
  errorType: '',
  status: '',
  retryCount: '',
  date: ''
})

const errorStats = reactive({
  totalFailed: 0,
  retryable: 0,
  resolved: 0,
  successRate: 0
})

// Computed properties
const filteredFailedSMS = computed(() => {
  let filtered = [...props.failedSMS]
  
  if (filters.errorType) {
    filtered = filtered.filter(sms => sms.errorType === filters.errorType)
  }
  
  if (filters.status) {
    filtered = filtered.filter(sms => sms.status === filters.status)
  }
  
  if (filters.retryCount) {
    const count = parseInt(filters.retryCount)
    if (count === 3) {
      filtered = filtered.filter(sms => sms.retryCount >= 3)
    } else {
      filtered = filtered.filter(sms => sms.retryCount === count)
    }
  }
  
  if (filters.date) {
    filtered = filtered.filter(sms => sms.lastAttempt.startsWith(filters.date))
  }
  
  return filtered
})

const totalErrors = computed(() => filteredFailedSMS.value.length)
const totalPages = computed(() => Math.ceil(totalErrors.value / itemsPerPage.value))
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, totalErrors.value))

// Methods
const resendSMS = async (smsId: number) => {
  emit('resend', smsId)
}

const retryAllFailed = async () => {
  retryingAll.value = true
  try {
    emit('retry-all')
    await new Promise(resolve => setTimeout(resolve, 3000))
  } finally {
    retryingAll.value = false
  }
}

const analyzeError = (errorId: number) => {
  emit('analyze-error', errorId)
}

const viewDetails = (_smsId: number) => {
  // console.log(`Viewing details for SMS ${smsId}`)
}

const exportErrorReport = () => {
  // console.log('Exporting error report...')
}

const canRetry = (sms: FailedSMS) => {
  return sms.retryCount < maxRetries && sms.status !== 'permanent_failure'
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

// Utility functions
const getInitials = (name: string) => {
  return name.split(' ').map(n => n[0]).join('').toUpperCase()
}

const formatDateTime = (dateTime: string) => {
  return new Date(dateTime).toLocaleString('fa-IR')
}

const getErrorTypeClass = (errorType: string) => {
  const classes = {
    invalid_number: 'bg-red-100 text-red-800',
    service_unavailable: 'bg-yellow-100 text-yellow-800',
    quota_exceeded: 'bg-orange-100 text-orange-800',
    network_error: 'bg-purple-100 text-purple-800',
    timeout: 'bg-gray-100 text-gray-800'
  }
  return classes[errorType as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getErrorTypeText = (errorType: string) => {
  const texts = {
    invalid_number: 'شماره نامعتبر',
    service_unavailable: 'سرویس در دسترس نیست',
    quota_exceeded: 'محدودیت ارسال',
    network_error: 'خطای شبکه',
    timeout: 'زمان انتظار'
  }
  return texts[errorType as keyof typeof texts] || errorType
}

const getStatusClass = (status: string) => {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800',
    retrying: 'bg-blue-100 text-blue-800',
    resolved: 'bg-green-100 text-green-800',
    permanent_failure: 'bg-red-100 text-red-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts = {
    pending: 'در انتظار',
    retrying: 'در حال تلاش مجدد',
    resolved: 'حل شده',
    permanent_failure: 'خطای دائمی'
  }
  return texts[status as keyof typeof texts] || status
}

// Update stats when data changes
const updateStats = () => {
  errorStats.totalFailed = props.failedSMS.length
  errorStats.retryable = props.failedSMS.filter(sms => canRetry(sms)).length
  errorStats.resolved = props.failedSMS.filter(sms => sms.status === 'resolved').length
  errorStats.successRate = Math.round((errorStats.resolved / errorStats.totalFailed) * 100) || 0
}

// Watch for changes
import { watch } from 'vue'
watch(() => props.failedSMS, updateStats, { immediate: true })
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
