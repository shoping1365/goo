<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">گزارش خطاهای همگام‌سازی</h4>
        <p class="text-sm text-gray-600 mt-1">گزارش و مدیریت خطاهای رخ داده در فرآیند همگام‌سازی</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="retryAllErrors"
          :disabled="isRetrying"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-yellow-600 hover:bg-yellow-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-yellow-500"
        >
          <svg v-if="isRetrying" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          {{ isRetrying ? 'در حال تلاش مجدد...' : 'تلاش مجدد همه خطاها' }}
        </button>
        <button
          @click="exportErrorsReport"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          صادر کردن گزارش
        </button>
      </div>
    </div>

    <!-- آمار خطاها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ stats.totalErrors }}</div>
            <div class="text-sm text-red-700">کل خطاها</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ stats.todayErrors }}</div>
            <div class="text-sm text-yellow-700">خطاهای امروز</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ stats.resolvedErrors }}</div>
            <div class="text-sm text-green-700">خطاهای حل شده</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ stats.errorRate }}%</div>
            <div class="text-sm text-blue-700">نرخ خطا</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای خطا -->
    <div class="bg-gray-50 rounded-lg p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلترهای خطا</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع خطا</label>
          <select
            v-model="filters.errorType"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه انواع</option>
            <option value="connection">خطای اتصال</option>
            <option value="authentication">خطای احراز هویت</option>
            <option value="data">خطای داده</option>
            <option value="timeout">خطای زمان انتظار</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نرم‌افزار</label>
          <select
            v-model="filters.software"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه نرم‌افزارها</option>
            <option value="helo">هلو</option>
            <option value="parsian">پارسیان</option>
            <option value="sepidar">سپیدار</option>
            <option value="ghias">قیاس</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="filters.status"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="open">باز</option>
            <option value="resolved">حل شده</option>
            <option value="ignored">نادیده گرفته شده</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
          <select
            v-model="filters.priority"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه اولویت‌ها</option>
            <option value="high">بالا</option>
            <option value="medium">متوسط</option>
            <option value="low">پایین</option>
          </select>
        </div>
      </div>
    </div>

    <!-- جدول خطاها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">خطاهای همگام‌سازی</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">شماره خطا</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع خطا</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">پیام خطا</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">اولویت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="error in filteredErrors" :key="error.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="font-medium text-gray-900">{{ error.errorNumber }}</div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getErrorTypeClass(error.type)"
                >
                  {{ getErrorTypeLabel(error.type) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ error.message }}</div>
                <div class="text-xs text-gray-500">{{ error.details }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <img :src="error.softwareLogo" :alt="error.softwareName" class="w-6 h-6 rounded" />
                  <span class="text-sm text-gray-900">{{ error.softwareName }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getPriorityClass(error.priority)"
                >
                  {{ getPriorityLabel(error.priority) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getStatusClass(error.status)"
                >
                  {{ getStatusLabel(error.status) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ error.date }}</div>
                <div class="text-xs text-gray-500">{{ error.time }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    @click="retryError(error)"
                    class="p-1 text-green-600 hover:text-green-800"
                    title="تلاش مجدد"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                    </svg>
                  </button>
                  <button
                    @click="viewErrorDetails(error)"
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="مشاهده جزئیات"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button
                    @click="ignoreError(error)"
                    class="p-1 text-yellow-600 hover:text-yellow-800"
                    title="نادیده گرفتن"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودارهای تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">توزیع خطاها بر اساس نوع</h5>
        <div class="space-y-4">
          <div v-for="type in errorTypeDistribution" :key="type.type" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :class="getErrorTypeColor(type.type)"></div>
              <span class="text-sm text-gray-700">{{ getErrorTypeLabel(type.type) }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="getErrorTypeColor(type.type)"
                  :style="{ width: type.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ type.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">خطاها بر اساس نرم‌افزار</h5>
        <div class="space-y-4">
          <div v-for="software in softwareErrorDistribution" :key="software.name" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <img :src="software.logo" :alt="software.name" class="w-6 h-6 rounded" />
              <span class="text-sm text-gray-700">{{ software.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full bg-red-500"
                  :style="{ width: software.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ software.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// وضعیت تلاش مجدد
const isRetrying = ref(false)

// فیلترها
const filters = ref({
  errorType: '',
  software: '',
  status: '',
  priority: ''
})

// آمار
const stats = ref({
  totalErrors: 150,
  todayErrors: 12,
  resolvedErrors: 138,
  errorRate: 1.2
})

// خطاها
const errors = ref([
  {
    id: 1,
    errorNumber: 'ERR-001',
    type: 'connection',
    message: 'خطا در اتصال به سرور',
    details: 'Timeout occurred while connecting to server',
    softwareName: 'هلو',
    softwareLogo: '/images/helo-logo.png',
    priority: 'high',
    status: 'open',
    date: 'امروز',
    time: '۱۴:۳۰'
  },
  {
    id: 2,
    errorNumber: 'ERR-002',
    type: 'authentication',
    message: 'خطا در احراز هویت',
    details: 'Invalid credentials provided',
    softwareName: 'پارسیان',
    softwareLogo: '/images/parsian-logo.png',
    priority: 'medium',
    status: 'resolved',
    date: 'امروز',
    time: '۱۳:۴۵'
  },
  {
    id: 3,
    errorNumber: 'ERR-003',
    type: 'data',
    message: 'خطا در پردازش داده',
    details: 'Data format is invalid',
    softwareName: 'سپیدار',
    softwareLogo: '/images/sepidar-logo.png',
    priority: 'low',
    status: 'open',
    date: 'دیروز',
    time: '۱۸:۲۰'
  }
])

// خطاهای فیلتر شده
const filteredErrors = computed(() => {
  return errors.value.filter(error => {
    if (filters.value.errorType && error.type !== filters.value.errorType) return false
    if (filters.value.software && error.softwareName.toLowerCase() !== filters.value.software) return false
    if (filters.value.status && error.status !== filters.value.status) return false
    if (filters.value.priority && error.priority !== filters.value.priority) return false
    return true
  })
})

// توزیع نوع خطا
const errorTypeDistribution = ref([
  { type: 'connection', percentage: 40 },
  { type: 'authentication', percentage: 25 },
  { type: 'data', percentage: 20 },
  { type: 'timeout', percentage: 15 }
])

// توزیع خطا بر اساس نرم‌افزار
const softwareErrorDistribution = ref([
  { name: 'هلو', logo: '/images/helo-logo.png', percentage: 35 },
  { name: 'پارسیان', logo: '/images/parsian-logo.png', percentage: 25 },
  { name: 'سپیدار', logo: '/images/sepidar-logo.png', percentage: 30 },
  { name: 'قیاس', logo: '/images/ghias-logo.png', percentage: 10 }
])

// کلاس نوع خطا
const getErrorTypeClass = (type: string) => {
  const classes = {
    connection: 'bg-red-100 text-red-700',
    authentication: 'bg-yellow-100 text-yellow-700',
    data: 'bg-blue-100 text-blue-700',
    timeout: 'bg-orange-100 text-orange-700'
  }
  return classes[type] || 'bg-gray-100 text-gray-700'
}

// رنگ نوع خطا
const getErrorTypeColor = (type: string) => {
  const colors = {
    connection: 'bg-red-500',
    authentication: 'bg-yellow-500',
    data: 'bg-blue-500',
    timeout: 'bg-orange-500'
  }
  return colors[type] || 'bg-gray-500'
}

// برچسب نوع خطا
const getErrorTypeLabel = (type: string) => {
  const labels = {
    connection: 'خطای اتصال',
    authentication: 'خطای احراز هویت',
    data: 'خطای داده',
    timeout: 'خطای زمان انتظار'
  }
  return labels[type] || type
}

// کلاس اولویت
const getPriorityClass = (priority: string) => {
  const classes = {
    high: 'bg-red-100 text-red-700',
    medium: 'bg-yellow-100 text-yellow-700',
    low: 'bg-green-100 text-green-700'
  }
  return classes[priority] || 'bg-gray-100 text-gray-700'
}

// برچسب اولویت
const getPriorityLabel = (priority: string) => {
  const labels = {
    high: 'بالا',
    medium: 'متوسط',
    low: 'پایین'
  }
  return labels[priority] || priority
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    open: 'bg-red-100 text-red-700',
    resolved: 'bg-green-100 text-green-700',
    ignored: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    open: 'باز',
    resolved: 'حل شده',
    ignored: 'نادیده گرفته شده'
  }
  return labels[status] || status
}

// اعمال فیلترها
const applyFilters = () => {
  // TODO: اعمال فیلترها
  console.log('فیلترهای خطا اعمال شد:', filters.value)
}

// تلاش مجدد همه خطاها
const retryAllErrors = async () => {
  try {
    isRetrying.value = true
    // TODO: تلاش مجدد همه خطاها
    console.log('تلاش مجدد همه خطاها شروع شد')
    
    // شبیه‌سازی تلاش مجدد
    await new Promise(resolve => setTimeout(resolve, 3000))
    
    console.log('تلاش مجدد همه خطاها تکمیل شد')
  } catch (error) {
    console.error('خطا در تلاش مجدد:', error)
  } finally {
    isRetrying.value = false
  }
}

// صادر کردن گزارش خطاها
const exportErrorsReport = () => {
  // TODO: صادر کردن گزارش خطاها
  console.log('صادر کردن گزارش خطاها')
}

// تلاش مجدد خطا
const retryError = (error: any) => {
  // TODO: تلاش مجدد خطای خاص
  console.log('تلاش مجدد خطا:', error)
}

// مشاهده جزئیات خطا
const viewErrorDetails = (error: any) => {
  // TODO: مشاهده جزئیات خطا
  console.log('مشاهده جزئیات خطا:', error)
}

// نادیده گرفتن خطا
const ignoreError = (error: any) => {
  // TODO: نادیده گرفتن خطا
  console.log('نادیده گرفتن خطا:', error)
}
</script>

<!--
  کامپوننت گزارش خطاهای همگام‌سازی
  شامل:
  1. آمار خطاها
  2. فیلترهای پیشرفته
  3. جدول خطاها
  4. نمودارهای تحلیلی
  5. عملیات مدیریت خطا
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
