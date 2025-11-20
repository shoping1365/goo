<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">گزارش خطاها</h4>
        <p class="text-sm text-gray-600 mt-1">گزارش‌های جامع و تحلیلی خطاهای اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          :disabled="isGenerating"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="generateReport"
        >
          <svg v-if="isGenerating" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
          </svg>
          {{ isGenerating ? 'در حال تولید...' : 'تولید گزارش' }}
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="exportReport"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          خروجی
        </button>
      </div>
    </div>

    <!-- فیلترهای گزارش -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلترهای گزارش</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select
            v-model="reportFilters.timeRange"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="today">امروز</option>
            <option value="yesterday">دیروز</option>
            <option value="week">هفته گذشته</option>
            <option value="month">ماه گذشته</option>
            <option value="quarter">سه ماه گذشته</option>
            <option value="year">سال گذشته</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع خطا</label>
          <select
            v-model="reportFilters.errorType"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه انواع</option>
            <option value="connection">خطاهای اتصال</option>
            <option value="authentication">خطاهای احراز هویت</option>
            <option value="sync">خطاهای همگام‌سازی</option>
            <option value="data">خطاهای داده</option>
            <option value="performance">خطاهای عملکرد</option>
            <option value="security">خطاهای امنیتی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
          <select
            v-model="reportFilters.priority"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه اولویت‌ها</option>
            <option value="high">زیاد</option>
            <option value="medium">متوسط</option>
            <option value="low">کم</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="reportFilters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="resolved">حل شده</option>
            <option value="ignored">نادیده گرفته شده</option>
          </select>
        </div>
      </div>

      <div v-if="reportFilters.timeRange === 'custom'" class="mt-4 grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input
            v-model="reportFilters.dateFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input
            v-model="reportFilters.dateTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ reportStats.totalErrors }}</div>
            <div class="text-sm text-red-700">کل خطاها</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ reportStats.resolvedErrors }}</div>
            <div class="text-sm text-green-700">خطاهای حل شده</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ reportStats.avgResolutionTime }}</div>
            <div class="text-sm text-yellow-700">میانگین زمان حل</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ reportStats.errorRate }}</div>
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

    <!-- نمودارها و تحلیل‌ها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار توزیع خطاها -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">توزیع خطاها بر اساس نوع</h5>
        <div class="space-y-3">
          <div v-for="errorType in errorTypeDistribution" :key="errorType.type" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div
                class="w-4 h-4 rounded-full"
                :style="{ backgroundColor: errorType.color }"
              ></div>
              <span class="text-sm text-gray-700">{{ errorType.type }}</span>
            </div>
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-32 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :style="{ width: errorType.percentage + '%', backgroundColor: errorType.color }"
                ></div>
              </div>
              <span class="text-sm font-medium text-gray-900">{{ errorType.count }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار روند خطاها -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">روند خطاها در زمان</h5>
        <div class="space-y-3">
          <div v-for="trend in errorTrends" :key="trend.period" class="flex items-center justify-between">
            <span class="text-sm text-gray-700">{{ trend.period }}</span>
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-32 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="trend.trend === 'up' ? 'bg-red-500' : 'bg-green-500'"
                  :style="{ width: Math.abs(trend.change) + '%' }"
                ></div>
              </div>
              <span
                class="text-sm font-medium"
                :class="trend.trend === 'up' ? 'text-red-600' : 'text-green-600'"
              >
                {{ trend.trend === 'up' ? '+' : '-' }}{{ trend.change }}%
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول گزارش خطاها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h5 class="text-md font-medium text-gray-900">گزارش تفصیلی خطاها</h5>
          <div class="flex items-center space-x-3 space-x-reverse">
            <span class="text-sm text-gray-500">{{ filteredErrors.length }} خطا</span>
            <button
              class="text-sm text-blue-600 hover:text-blue-800"
              @click="toggleSort"
            >
              {{ sortDirection === 'asc' ? 'صعودی' : 'نزولی' }}
            </button>
          </div>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">شناسه</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع خطا</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">توضیحات</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">اولویت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ وقوع</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">زمان حل</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="error in paginatedErrors" :key="error.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <span class="text-sm text-gray-900">#{{ error.id }}</span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <div
                    class="w-3 h-3 rounded-full"
                    :class="getErrorTypeColor(error.type)"
                  ></div>
                  <span class="font-medium text-gray-900">{{ error.type }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ error.description }}</div>
                <div class="text-xs text-gray-500">{{ error.details }}</div>
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
                <div class="text-sm text-gray-900">{{ error.occurredDate }}</div>
                <div class="text-xs text-gray-500">{{ error.occurredTime }}</div>
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
                <span class="text-sm text-gray-900">{{ error.resolutionTime || '-' }}</span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="جزئیات"
                    @click="viewErrorDetails(error)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-green-600 hover:text-green-800"
                    title="خروجی"
                    @click="exportError(error)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="px-6 py-4 border-t border-gray-200">
        <div class="flex items-center justify-between">
          <div class="text-sm text-gray-500">
            نمایش {{ (currentPage - 1) * pageSize + 1 }} تا {{ Math.min(currentPage * pageSize, filteredErrors.length) }} از {{ filteredErrors.length }} خطا
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              :disabled="currentPage === 1"
              class="px-3 py-1 text-sm border border-gray-300 rounded-md disabled:bg-gray-100 disabled:text-gray-400"
              @click="previousPage"
            >
              قبلی
            </button>
            <span class="px-3 py-1 text-sm text-gray-700">{{ currentPage }} از {{ totalPages }}</span>
            <button
              :disabled="currentPage === totalPages"
              class="px-3 py-1 text-sm border border-gray-300 rounded-md disabled:bg-gray-100 disabled:text-gray-400"
              @click="nextPage"
            >
              بعدی
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// وضعیت تولید گزارش
const isGenerating = ref(false)

// فیلترهای گزارش
const reportFilters = ref({
  timeRange: 'month',
  errorType: 'all',
  priority: 'all',
  status: 'all',
  dateFrom: '',
  dateTo: ''
})

// آمار گزارش
const reportStats = ref({
  totalErrors: 234,
  resolvedErrors: 198,
  avgResolutionTime: '۲.۱ ساعت',
  errorRate: 3.2
})

// توزیع نوع خطاها
const errorTypeDistribution = ref([
  { type: 'خطاهای اتصال', count: 45, percentage: 19.2, color: '#EF4444' },
  { type: 'خطاهای احراز هویت', count: 32, percentage: 13.7, color: '#F59E0B' },
  { type: 'خطاهای همگام‌سازی', count: 67, percentage: 28.6, color: '#3B82F6' },
  { type: 'خطاهای داده', count: 28, percentage: 12.0, color: '#8B5CF6' },
  { type: 'خطاهای عملکرد', count: 42, percentage: 17.9, color: '#F97316' },
  { type: 'خطاهای امنیتی', count: 20, percentage: 8.6, color: '#DC2626' }
])

// روند خطاها
const errorTrends = ref([
  { period: 'هفته گذشته', change: 15, trend: 'up' },
  { period: 'دو هفته پیش', change: 8, trend: 'down' },
  { period: 'سه هفته پیش', change: 22, trend: 'up' },
  { period: 'ماه گذشته', change: 12, trend: 'down' }
])

// خطاهای فیلتر شده
const filteredErrors = ref([
  {
    id: 1,
    type: 'خطای اتصال',
    description: 'Timeout در اتصال به نرم‌افزار هلو',
    details: 'اتصال به سرور هلو پس از ۳۰ ثانیه timeout شد',
    priority: 'high',
    occurredDate: 'امروز',
    occurredTime: '۱۴:۳۰',
    status: 'resolved',
    resolutionTime: '۱.۲ ساعت'
  },
  {
    id: 2,
    type: 'خطای احراز هویت',
    description: 'کلید API نامعتبر',
    details: 'کلید API برای نرم‌افزار پارسیان منقضی شده است',
    priority: 'medium',
    occurredDate: 'دیروز',
    occurredTime: '۱۸:۴۵',
    status: 'resolved',
    resolutionTime: '۳.۵ ساعت'
  },
  {
    id: 3,
    type: 'خطای همگام‌سازی',
    description: 'خطا در همگام‌سازی تراکنش‌ها',
    details: 'خطای فرمت در داده‌های ارسالی به نرم‌افزار سپیدار',
    priority: 'low',
    occurredDate: 'هفته گذشته',
    occurredTime: '۱۰:۱۵',
    status: 'ignored',
    resolutionTime: null
  }
])

// صفحه‌بندی
const currentPage = ref(1)
const pageSize = ref(10)
const sortDirection = ref('desc')

// محاسبه صفحات
const totalPages = computed(() => Math.ceil(filteredErrors.value.length / pageSize.value))

// خطاهای صفحه فعلی
const paginatedErrors = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredErrors.value.slice(start, end)
})

// رنگ نوع خطا
const getErrorTypeColor = (type: string) => {
  const colors = {
    'خطای اتصال': 'bg-red-500',
    'خطای احراز هویت': 'bg-yellow-500',
    'خطای همگام‌سازی': 'bg-blue-500',
    'خطای داده': 'bg-purple-500',
    'خطای عملکرد': 'bg-orange-500',
    'خطای امنیتی': 'bg-red-600'
  }
  return colors[type] || 'bg-gray-500'
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
    high: 'زیاد',
    medium: 'متوسط',
    low: 'کم'
  }
  return labels[priority] || priority
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-red-100 text-red-700',
    resolved: 'bg-green-100 text-green-700',
    ignored: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    resolved: 'حل شده',
    ignored: 'نادیده گرفته شده'
  }
  return labels[status] || status
}

// تولید گزارش
const generateReport = async () => {
  try {
    isGenerating.value = true
    // TODO: تولید گزارش
    console.log('تولید گزارش با فیلترها:', reportFilters.value)
    
    // شبیه‌سازی تولید گزارش
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    console.log('گزارش تولید شد')
  } catch (error) {
    console.error('خطا در تولید گزارش:', error)
  } finally {
    isGenerating.value = false
  }
}

// خروجی گزارش
const exportReport = () => {
  // TODO: خروجی گزارش
  console.log('خروجی گزارش')
}

// تغییر ترتیب
const toggleSort = () => {
  sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc'
  // TODO: اعمال ترتیب
}

// صفحه قبلی
const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

// صفحه بعدی
const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

// مشاهده جزئیات خطا
const viewErrorDetails = (error: any) => {
  // TODO: نمایش جزئیات خطا
  console.log('جزئیات خطا:', error)
}

// خروجی خطا
const exportError = (error: any) => {
  // TODO: خروجی خطا
  console.log('خروجی خطا:', error)
}
</script>

<!--
  کامپوننت گزارش خطاها
  شامل:
  1. فیلترهای پیشرفته گزارش
  2. آمار کلی خطاها
  3. نمودار توزیع خطاها
  4. نمودار روند خطاها
  5. جدول تفصیلی خطاها
  6. صفحه‌بندی
  7. قابلیت خروجی
  8. طراحی مدرن و کاملاً ریسپانسیو
--> 
