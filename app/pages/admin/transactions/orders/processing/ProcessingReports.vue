<template>
  <div>
    <!-- آمار کلی گزارشات -->
    <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">آمار کلی گزارشات در حال انجام</h3>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
          <!-- کل سفارشات در حال انجام -->
          <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-blue-700 mb-1">کل سفارشات در حال انجام</p>
                <p class="text-2xl font-bold text-blue-900">{{ reportStats.totalProcessing.toLocaleString('fa-IR') }}</p>
                <p class="text-xs text-blue-600 mt-1">در حال پردازش</p>
              </div>
              <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- مبلغ کل در حال انجام -->
          <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-green-700 mb-1">مبلغ کل در حال انجام</p>
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

          <!-- نرخ تکمیل -->
          <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-purple-700 mb-1">نرخ تکمیل</p>
                <p class="text-2xl font-bold text-purple-900">{{ reportStats.completionRate }}%</p>
                <p class="text-xs text-purple-600 mt-1">سفارشات تکمیل شده</p>
              </div>
              <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
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
                <p class="text-xs text-orange-600 mt-1">از شروع تا تکمیل</p>
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
      <div class="bg-gradient-to-r from-indigo-50 to-purple-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-indigo-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته گزارشات</h3>
          </div>
          <button class="text-sm text-indigo-600 hover:text-indigo-800 transition-colors font-medium hover:bg-indigo-50 px-3 py-1 rounded-lg" @click="showReportFilters = !showReportFilters">
            {{ showReportFilters ? 'مخفی کردن' : 'نمایش' }}
          </button>
        </div>
      </div>
      
      <div v-if="showReportFilters" class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
          <!-- فیلتر دوره زمانی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
            <select v-model="reportFilters.timePeriod" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
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
            <select v-model="reportFilters.amountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه مبالغ</option>
              <option value="0-1000000">کمتر از 1 میلیون تومان</option>
              <option value="1000000-5000000">1 تا 5 میلیون تومان</option>
              <option value="5000000-10000000">5 تا 10 میلیون تومان</option>
              <option value="10000000+">بیش از 10 میلیون تومان</option>
            </select>
          </div>

          <!-- فیلتر نرخ تکمیل -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ تکمیل</label>
            <select v-model="reportFilters.completionRate" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه نرخ‌ها</option>
              <option value="0-25">کمتر از 25%</option>
              <option value="25-50">25 تا 50%</option>
              <option value="50-75">50 تا 75%</option>
              <option value="75-100">75 تا 100%</option>
            </select>
          </div>

          <!-- فیلتر نوع گزارش -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
            <select v-model="reportFilters.reportType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه گزارشات</option>
              <option value="processing">گزارش پردازش</option>
              <option value="performance">گزارش عملکرد</option>
              <option value="efficiency">گزارش کارایی</option>
              <option value="bottleneck">گزارش گلوگاه</option>
            </select>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="mt-6 flex justify-between items-center">
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors" @click="clearReportFilters">
              پاک کردن فیلترها
            </button>
            <button class="px-4 py-2 bg-blue-100 text-blue-700 rounded-md hover:bg-blue-200 transition-colors" @click="resetReportFilters">
              بازنشانی فیلترها
            </button>
          </div>
          
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition-colors flex items-center" @click="exportFilteredReports">
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی Excel
            </button>
            <button class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition-colors flex items-center" @click="exportFilteredReportsPDF">
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
          <h3 class="text-lg font-semibold text-gray-900">گزارشات تفصیلی در حال انجام</h3>
          <div class="relative">
            <input
              v-model="reportSearchTerm"
              type="text"
              placeholder="جستجو در گزارشات..."
              class="w-48 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
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
        <!-- Loading State -->
        <div v-if="isLoading" class="flex justify-center items-center py-8">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
          <span class="mr-2 text-gray-600">در حال بارگذاری گزارشات...</span>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="text-center py-8">
          <div class="text-red-600 mb-4">{{ error }}</div>
          <button class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600" @click="fetchReportData">
            تلاش مجدد
          </button>
        </div>

        <!-- Reports Table -->
        <table v-else class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دوره</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد سفارشات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ کل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تکمیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">متوسط زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-if="paginatedReports.length === 0">
              <td colspan="7" class="px-6 py-8 text-center text-gray-500">
                هیچ گزارش تفصیلی یافت نشد
              </td>
            </tr>
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
                <span :class="getCompletionRateClass(report.completionRate)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ report.completionRate.toFixed(1) }}%
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.avgProcessingTime.toFixed(1) }} ساعت
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getChangeClass(report.change)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ report.change > 0 ? '+' : '' }}{{ report.change.toFixed(1) }}%
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2">
                  <button class="text-blue-600 hover:text-blue-900" @click="viewReportDetails(report)">جزئیات</button>
                  <button class="text-green-600 hover:text-green-900" @click="exportReport(report)">خروجی</button>
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
  totalProcessing: 0,
  totalAmount: 0,
  completionRate: 0,
  avgProcessingTime: 0
})

// داده‌های گزارشات تفصیلی واقعی
const detailedReports = ref([])

const isLoading = ref(false)
const error = ref(null)

// متغیرهای فیلتر گزارشات
const showReportFilters = ref(false)
const reportFilters = ref({
  timePeriod: '',
  amountRange: '',
  completionRate: '',
  reportType: ''
})

// متغیرهای جستجو و صفحه‌بندی گزارشات
const reportSearchTerm = ref('')
const reportCurrentPage = ref(1)
const reportItemsPerPage = ref(10)

// تابع دریافت داده‌های گزارشات از API
const fetchReportData = async () => {
  isLoading.value = true
  error.value = null
  
  try {
    const data = await $fetch('/api/admin/orders/processing/stats')
    
    // آمار اصلی
    reportStats.value = {
      totalProcessing: data.totalProcessing || 0,
      totalAmount: data.totalAmount || 0,
      completionRate: data.completionRate || 0,
      avgProcessingTime: data.avgProcessingTime || 0
    }

    // گزارشات تفصیلی
    if (data.detailedReports && data.detailedReports.length > 0) {
      detailedReports.value = data.detailedReports.map((report, index) => ({
        id: index + 1,
        period: report.period,
        orderCount: report.orderCount || 0,
        totalAmount: report.totalAmount || 0,
        completionRate: report.completionRate || 0,
        avgProcessingTime: report.avgOrderValue || 0,
        change: report.change || 0
      }))
    } else {
      detailedReports.value = []
    }

  } catch (err) {
    error.value = 'خطا در دریافت داده‌های گزارشات'
    console.error('Error fetching report data:', err)
    detailedReports.value = []
  } finally {
    isLoading.value = false
  }
}

// دریافت داده‌ها هنگام بارگذاری کامپوننت
onMounted(() => {
  fetchReportData()
})

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

  // فیلتر بر اساس نرخ تکمیل
  if (reportFilters.value.completionRate) {
    const [min, max] = reportFilters.value.completionRate.split('-').map(Number)
    filtered = filtered.filter(report => {
      if (max) {
        return report.completionRate >= min && report.completionRate <= max
      } else {
        return report.completionRate >= min
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
    completionRate: '',
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
  // اینجا می‌توانید منطق خروجی Excel را پیاده‌سازی کنید
}

const exportFilteredReportsPDF = () => {
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
const viewReportDetails = (_report) => {
  // View report details logic
}

const exportReport = (_report) => {
  // Export report logic
}

// متدهای کمکی
const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const getCompletionRateClass = (rate) => {
  if (rate >= 90) return 'bg-green-100 text-green-800'
  if (rate >= 80) return 'bg-blue-100 text-blue-800'
  if (rate >= 70) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

const getChangeClass = (change) => {
  if (change > 0) return 'bg-green-100 text-green-800'
  if (change < 0) return 'bg-red-100 text-red-800'
  return 'bg-gray-100 text-gray-800'
}
</script> 
