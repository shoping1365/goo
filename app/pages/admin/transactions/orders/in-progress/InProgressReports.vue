<template>
  <div>
    <!-- آمار کلی گزارشات -->
    <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-indigo-50 to-purple-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-indigo-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">آمار کلی گزارشات در صف پردازش</h3>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
          <!-- کل سفارشات در صف -->
          <div class="bg-gradient-to-br from-indigo-50 to-indigo-100 border border-indigo-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-indigo-700 mb-1">کل سفارشات در صف</p>
                <p class="text-2xl font-bold text-indigo-900">{{ reportStats.totalInQueue.toLocaleString('fa-IR') }}</p>
                <p class="text-xs text-indigo-600 mt-1">در حال پردازش</p>
              </div>
              <div class="w-12 h-12 bg-indigo-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  
                </svg>
              </div>
            </div>
          </div>

          <!-- مبلغ کل در صف -->
          <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-green-700 mb-1">مبلغ کل در صف</p>
                <p class="text-2xl font-bold text-green-900">{{ formatPrice(reportStats.totalAmount) }}</p>
                <p class="text-xs text-green-600 mt-1">در حال پردازش</p>
              </div>
              <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- نرخ تبدیل -->
          <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-blue-700 mb-1">نرخ تبدیل</p>
                <p class="text-2xl font-bold text-blue-900">{{ reportStats.conversionRate }}%</p>
                <p class="text-xs text-blue-600 mt-1">از کل سفارشات</p>
              </div>
              <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- متوسط زمان پردازش -->
          <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-orange-700 mb-1">متوسط زمان پردازش</p>
                <p class="text-2xl font-bold text-orange-900">{{ reportStats.avgProcessingTime }} ساعت</p>
                <p class="text-xs text-orange-600 mt-1">از ثبت تا تکمیل</p>
              </div>
              <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-purple-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-purple-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته گزارشات</h3>
          </div>
          <button @click="showReportFilters = !showReportFilters" class="text-sm text-purple-600 hover:text-purple-800 transition-colors font-medium hover:bg-purple-50 px-3 py-1 rounded-lg">
            {{ showReportFilters ? 'مخفی کردن' : 'نمایش' }}
          </button>
        </div>
      </div>
      
      <div v-if="showReportFilters" class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
          <!-- فیلتر دوره زمانی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
            <select v-model="reportFilters.timePeriod" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="">همه دوره‌ها</option>
              <option value="today">امروز</option>
              <option value="week">هفته جاری</option>
              <option value="month">ماه جاری</option>
              <option value="quarter">فصل جاری</option>
              <option value="year">سال جاری</option>
              <option value="custom">انتخاب دستی</option>
            </select>
          </div>

          <!-- فیلتر محدوده مبلغ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدوده مبلغ</label>
            <select v-model="reportFilters.amountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="">همه مبالغ</option>
              <option value="0-1000000">کمتر از 1 میلیون تومان</option>
              <option value="1000000-5000000">1 تا 5 میلیون تومان</option>
              <option value="5000000-10000000">5 تا 10 میلیون تومان</option>
              <option value="10000000+">بیش از 10 میلیون تومان</option>
            </select>
          </div>

          <!-- فیلتر نرخ تبدیل -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ تبدیل</label>
            <select v-model="reportFilters.conversionRate" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="">همه نرخ‌ها</option>
              <option value="0-1">کمتر از 1%</option>
              <option value="1-2">1 تا 2%</option>
              <option value="2-3">2 تا 3%</option>
              <option value="3-4">3 تا 4%</option>
              <option value="4+">بیش از 4%</option>
            </select>
          </div>

          <!-- فیلتر نوع گزارش -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
            <select v-model="reportFilters.reportType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="">همه گزارشات</option>
              <option value="processing">گزارش پردازش</option>
              <option value="payment">گزارش پرداخت</option>
              <option value="fraud">گزارش تقلب</option>
              <option value="performance">گزارش عملکرد</option>
            </select>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="mt-6 flex justify-between items-center">
          <div class="flex space-x-2 space-x-reverse">
            <button @click="clearReportFilters" class="px-4 py-2 bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors">
              پاک کردن فیلترها
            </button>
            <button @click="resetReportFilters" class="px-4 py-2 bg-blue-100 text-blue-700 rounded-md hover:bg-blue-200 transition-colors">
              بازنشانی فیلترها
            </button>
          </div>
          
          <div class="flex space-x-2 space-x-reverse">
            <button @click="exportFilteredReports" class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition-colors flex items-center">
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی Excel
            </button>
            <button @click="exportFilteredReportsPDF" class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition-colors flex items-center">
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
              </svg>
              خروجی PDF
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول گزارشات تفصیلی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="px-4 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">گزارشات تفصیلی در صف پردازش</h3>
          <div class="relative">
            <input
              type="text"
              v-model="reportSearchTerm"
              placeholder="جستجو در گزارشات..."
              class="w-48 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            />
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
          </div>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دوره</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد سفارشات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ کل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تبدیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">متوسط زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in paginatedReports" :key="report.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ report.period }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.orderCount.toLocaleString('fa-IR') }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatPrice(report.totalAmount) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getConversionRateClass(report.conversionRate)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ report.conversionRate }}%
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.avgProcessingTime }} ساعت
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getChangeClass(report.change)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ report.change > 0 ? '+' : '' }}{{ report.change }}%
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2">
                  <button @click="viewReportDetails(report)" class="text-blue-600 hover:text-blue-900">جزئیات</button>
                  <button @click="exportReport(report)" class="text-green-600 hover:text-green-900">خروجی</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- صفحه‌بندی -->
      <Pagination 
        :current-page="reportCurrentPage"
        :total-pages="reportTotalPages"
        :total="filteredReports.length"
        :items-per-page="reportItemsPerPage"
        @page-changed="handleReportPageChange"
        @items-per-page-changed="handleReportItemsPerPageChange"
      />
    </div>
  </div>
</template>

<script setup>
// Import کامپوننت Pagination
import Pagination from '~/components/admin/common/Pagination.vue'

// داده‌های آماری گزارشات
const reportStats = ref({
  totalInQueue: 234,
  totalAmount: 67800000,
  conversionRate: 3.2,
  avgProcessingTime: 3.2
})

// متغیرهای فیلتر گزارشات
const showReportFilters = ref(false)
const reportFilters = ref({
  timePeriod: '',
  amountRange: '',
  conversionRate: '',
  reportType: ''
})

// متغیرهای جستجو و صفحه‌بندی گزارشات
const reportSearchTerm = ref('')
const reportCurrentPage = ref(1)
const reportItemsPerPage = ref(10)

// داده‌های گزارشات تفصیلی
const detailedReports = ref([
  {
    id: 1,
    period: 'امروز',
    orderCount: 45,
    totalAmount: 12500000,
    conversionRate: 3.5,
    avgProcessingTime: 2.8,
    change: 12.5
  },
  {
    id: 2,
    period: 'هفته جاری',
    orderCount: 234,
    totalAmount: 67800000,
    conversionRate: 3.2,
    avgProcessingTime: 3.2,
    change: -2.1
  },
  {
    id: 3,
    period: 'ماه جاری',
    orderCount: 892,
    totalAmount: 245000000,
    conversionRate: 2.8,
    avgProcessingTime: 3.5,
    change: 8.7
  },
  {
    id: 4,
    period: 'فصل جاری',
    orderCount: 2345,
    totalAmount: 678000000,
    conversionRate: 2.9,
    avgProcessingTime: 3.1,
    change: 15.3
  },
  {
    id: 5,
    period: 'سال جاری',
    orderCount: 8923,
    totalAmount: 2450000000,
    conversionRate: 3.1,
    avgProcessingTime: 2.9,
    change: 22.8
  }
])

// محاسبات فیلتر و جستجو گزارشات
const filteredReports = computed(() => {
  let filtered = detailedReports.value

  // فیلتر بر اساس جستجو
  if (reportSearchTerm.value) {
    filtered = filtered.filter(report => 
      report.period.toLowerCase().includes(reportSearchTerm.value.toLowerCase())
    )
  }

  // فیلتر بر اساس دوره زمانی
  if (reportFilters.value.timePeriod) {
    // اینجا می‌توانید منطق فیلتر زمانی را پیاده‌سازی کنید
  }

  // فیلتر بر اساس محدوده مبلغ
  if (reportFilters.value.amountRange) {
    const [min, max] = reportFilters.value.amountRange.split('-').map(Number)
    filtered = filtered.filter(report => {
      if (max) {
        return report.totalAmount >= min && report.totalAmount <= max
      } else {
        return report.totalAmount >= min
      }
    })
  }

  // فیلتر بر اساس نرخ تبدیل
  if (reportFilters.value.conversionRate) {
    const [min, max] = reportFilters.value.conversionRate.split('-').map(Number)
    filtered = filtered.filter(report => {
      if (max) {
        return report.conversionRate >= min && report.conversionRate <= max
      } else {
        return report.conversionRate >= min
      }
    })
  }

  return filtered
})

// محاسبات صفحه‌بندی گزارشات
const reportTotalPages = computed(() => Math.ceil(filteredReports.value.length / reportItemsPerPage.value))

const paginatedReports = computed(() => {
  const start = (reportCurrentPage.value - 1) * reportItemsPerPage.value
  const end = start + reportItemsPerPage.value
  return filteredReports.value.slice(start, end)
})

// متدهای عملیاتی گزارشات
const clearReportFilters = () => {
  reportFilters.value = {
    timePeriod: '',
    amountRange: '',
    conversionRate: '',
    reportType: ''
  }
  reportSearchTerm.value = ''
  reportCurrentPage.value = 1
}

const resetReportFilters = () => {
  clearReportFilters()
  reportCurrentPage.value = 1
}

const exportFilteredReports = () => {
  console.log('خروجی Excel گزارشات فیلتر شده:', reportFilters.value)
  // اینجا می‌توانید منطق خروجی Excel را پیاده‌سازی کنید
}

const exportFilteredReportsPDF = () => {
  console.log('خروجی PDF گزارشات فیلتر شده:', reportFilters.value)
  // اینجا می‌توانید منطق خروجی PDF را پیاده‌سازی کنید
}

// نظارت بر تغییرات فیلترها و بازگشت به صفحه اول
watch([reportFilters, reportSearchTerm], () => {
  reportCurrentPage.value = 1
}, { deep: true })

// متدهای صفحه‌بندی گزارشات
const handleReportPageChange = (page) => {
  reportCurrentPage.value = page
}

const handleReportItemsPerPageChange = (newItemsPerPage) => {
  reportItemsPerPage.value = newItemsPerPage
  reportCurrentPage.value = 1
}

// متدهای عملیات گزارشات
const viewReportDetails = (report) => {
  console.log('مشاهده جزئیات گزارش:', report)
}

const exportReport = (report) => {
  console.log('خروجی گزارش:', report)
}

// متدهای کمکی
const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const getConversionRateClass = (rate) => {
  if (rate >= 4) return 'bg-green-100 text-green-800'
  if (rate >= 3) return 'bg-blue-100 text-blue-800'
  if (rate >= 2) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

const getChangeClass = (change) => {
  if (change > 0) return 'bg-green-100 text-green-800'
  if (change < 0) return 'bg-red-100 text-red-800'
  return 'bg-gray-100 text-gray-800'
}
</script> 
