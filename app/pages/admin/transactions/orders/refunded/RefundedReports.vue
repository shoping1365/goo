<template>
  <div>
    <!-- کارت‌های آماری گزارشات -->
    <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4 mb-6">
      <!-- گزارش ماهانه -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">گزارش ماهانه</h3>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">کل مسترد شده:</span>
            <span class="text-sm font-semibold text-gray-900">{{ monthlyStats.totalRefunded || 0 }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">مبلغ کل:</span>
            <span class="text-sm font-semibold text-gray-900">{{ formatPrice(monthlyStats.totalAmount || 0) }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">نرخ رشد:</span>
            <span :class="monthlyStats.growthRate >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-semibold">
              {{ (monthlyStats.growthRate || 0) >= 0 ? '+' : '' }}{{ monthlyStats.growthRate || 0 }}%
            </span>
          </div>
        </div>
      </div>

      <!-- گزارش هفتگی -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">گزارش هفتگی</h3>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">مسترد این هفته:</span>
            <span class="text-sm font-semibold text-gray-900">{{ weeklyStats.thisWeek || 0 }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">مسترد هفته قبل:</span>
            <span class="text-sm font-semibold text-gray-900">{{ weeklyStats.lastWeek || 0 }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">تغییر:</span>
            <span :class="weeklyStats.change >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-semibold">
              {{ (weeklyStats.change || 0) >= 0 ? '+' : '' }}{{ weeklyStats.change || 0 }}%
            </span>
          </div>
        </div>
      </div>

      <!-- گزارش روزانه -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">گزارش روزانه</h3>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">امروز:</span>
            <span class="text-sm font-semibold text-gray-900">{{ dailyStats.today || 0 }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">دیروز:</span>
            <span class="text-sm font-semibold text-gray-900">{{ dailyStats.yesterday || 0 }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">متوسط روزانه:</span>
            <span class="text-sm font-semibold text-gray-900">{{ dailyStats.average || 0 }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای گزارشات -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
      <!-- نمودار روند مسترد -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">روند مسترد در 30 روز گذشته</h3>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              v-for="period in chartPeriods"
              :key="period.id"
              @click="selectedChartPeriod = period.id"
              :class="[
                'px-3 py-1 text-xs font-medium rounded-md transition-colors',
                selectedChartPeriod === period.id
                  ? 'bg-blue-500 text-white'
                  : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
              ]"
            >
              {{ period.name }}
            </button>
          </div>
        </div>
        <div class="h-64 flex items-end justify-between space-x-1 space-x-reverse">
          <div
            v-for="(day, index) in trendChartData"
            :key="index"
            class="flex-1 flex flex-col items-center"
          >
            <div class="text-xs text-gray-500 mb-2">{{ day.date }}</div>
            <div
              class="w-full bg-gradient-to-t from-blue-500 to-blue-400 rounded-t"
              :style="{ height: `${(day.value / maxTrendValue) * 200}px` }"
            ></div>
            <div class="text-xs font-medium text-gray-700 mt-1">{{ day.value }}</div>
          </div>
        </div>
      </div>

      <!-- نمودار توزیع دلایل -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع دلایل مسترد</h3>
        <div class="space-y-4">
          <div
            v-for="reason in refundReasonsChart"
            :key="reason.id"
            class="flex items-center justify-between"
          >
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full ml-3" :style="{ backgroundColor: reason.color }"></div>
              <span class="text-sm text-gray-700">{{ reason.name }}</span>
            </div>
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-32 bg-gray-200 rounded-full h-3">
                <div
                  class="h-3 rounded-full"
                  :style="{ width: `${reason.percentage}%`, backgroundColor: reason.color }"
                ></div>
              </div>
              <span class="text-sm font-medium text-gray-900 w-12 text-left">{{ reason.percentage }}%</span>
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
            <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته</h3>
          </div>
          <button @click="showFilters = !showFilters" class="text-sm text-indigo-600 hover:text-indigo-800 transition-colors font-medium hover:bg-indigo-50 px-3 py-1 rounded-lg">
            {{ showFilters ? 'مخفی کردن' : 'نمایش' }}
          </button>
        </div>
      </div>
      
      <div v-if="showFilters" class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <!-- فیلتر دوره -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره</label>
            <select v-model="filters.period" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه دوره‌ها</option>
              <option value="current-week">هفته جاری</option>
              <option value="last-week">هفته گذشته</option>
              <option value="current-month">ماه جاری</option>
              <option value="last-month">ماه گذشته</option>
            </select>
          </div>
          
          <!-- فیلتر نرخ مسترد -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ مسترد</label>
            <select v-model="filters.refundRate" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه نرخ‌ها</option>
              <option value="low">کمتر از 2%</option>
              <option value="medium">2% تا 5%</option>
              <option value="high">بیش از 5%</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4 mt-4">
          <!-- فیلتر مبلغ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ کل</label>
            <select v-model="filters.amountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه مبالغ</option>
              <option value="0-10000000">کمتر از 10 میلیون تومان</option>
              <option value="10000000-50000000">10 تا 50 میلیون تومان</option>
              <option value="50000000-100000000">50 تا 100 میلیون تومان</option>
              <option value="100000000+">بیش از 100 میلیون تومان</option>
            </select>
          </div>
          
          <!-- فیلتر تغییر -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تغییر</label>
            <select v-model="filters.change" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه تغییرات</option>
              <option value="positive">مثبت</option>
              <option value="negative">منفی</option>
              <option value="neutral">بدون تغییر</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4 mt-4">
          <!-- فیلتر تعداد مسترد -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد مسترد</label>
            <select v-model="filters.countRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه تعدادها</option>
              <option value="0-10">کمتر از 10</option>
              <option value="10-50">10 تا 50</option>
              <option value="50-100">50 تا 100</option>
              <option value="100+">بیش از 100</option>
            </select>
          </div>

          <!-- فیلتر متوسط مبلغ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متوسط مبلغ</label>
            <select v-model="filters.averageAmountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه مبالغ</option>
              <option value="0-200000">کمتر از 200 هزار تومان</option>
              <option value="200000-500000">200 تا 500 هزار تومان</option>
              <option value="500000-1000000">500 هزار تا 1 میلیون تومان</option>
              <option value="1000000+">بیش از 1 میلیون تومان</option>
            </select>
          </div>
        </div>

        <!-- دکمه‌های عملیات فیلتر -->
        <div class="flex items-center justify-between mt-4 pt-4 border-t border-gray-200">
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              @click="clearFilters"
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              پاک کردن فیلترها
            </button>
          </div>
          <div class="text-sm text-gray-500">
            {{ filteredReports.length }} گزارش یافت شد
          </div>
        </div>
      </div>
    </div>

    <!-- جدول گزارشات تفصیلی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="bg-gradient-to-r from-gray-50 to-gray-100 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">گزارشات تفصیلی</h3>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دوره</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد مسترد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ کل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ مسترد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">متوسط مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in paginatedReports" :key="report.period || report.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ report.period || 'نامشخص' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.count || 0 }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatPrice(report.totalAmount || 0) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.refundRate || 0 }}%
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatPrice(report.averageAmount || 0) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="report.change >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                  {{ (report.change || 0) >= 0 ? '+' : '' }}{{ report.change || 0 }}%
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <Pagination
        :current-page="currentPage"
        :total-pages="totalPages"
        :total="filteredReports.length"
        :items-per-page="itemsPerPage"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>
  </div>
</template>

<script setup>
// Import کامپوننت Pagination
import Pagination from '~/components/admin/common/Pagination.vue'

// داده‌های آماری
const monthlyStats = ref({})
const weeklyStats = ref({})
const dailyStats = ref({})

// متغیرهای مدیریت وضعیت بارگذاری
const loading = ref(false)
const error = ref(null)

// تابع دریافت آمار مسترد
const fetchStats = async () => {
  try {
    loading.value = true
    error.value = null

    const { data } = await $fetch('/api/admin/orders/refunded/stats', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    // تبدیل داده‌های API به فرمت مورد نیاز کامپوننت
    monthlyStats.value = {
      totalRefunded: data.totalRefunded || 0,
      totalAmount: (data.todaySales || 0) + (data.weeklySales || 0) + (data.monthlySales || 0),
      growthRate: 0 // TODO: محاسبه نرخ رشد
    }

    weeklyStats.value = {
      thisWeek: data.weeklySales || 0,
      lastWeek: data.monthlySales || 0,
      change: 0 // TODO: محاسبه تغییر نسبت به هفته قبل
    }

    dailyStats.value = {
      today: data.todaySales || 0,
      yesterday: data.weeklySales || 0,
      average: data.averageOrder || 0
    }

  } catch (err) {
    console.error('خطا در دریافت آمار:', err)
    error.value = err.message || 'خطا در دریافت آمار'
  } finally {
    loading.value = false
  }
}

// مقداردهی اولیه و دریافت داده‌ها
onMounted(() => {
  fetchStats()
})

// دوره‌های نمودار
const chartPeriods = ref([
  { id: '7d', name: '7 روز' },
  { id: '30d', name: '30 روز' },
  { id: '90d', name: '90 روز' }
])

const selectedChartPeriod = ref('30d')

// داده‌های نمودار روند
const trendChartData = ref([
  { date: '1', value: 12 },
  { date: '2', value: 18 },
  { date: '3', value: 15 },
  { date: '4', value: 22 },
  { date: '5', value: 19 },
  { date: '6', value: 25 },
  { date: '7', value: 16 },
  { date: '8', value: 21 },
  { date: '9', value: 28 },
  { date: '10', value: 24 },
  { date: '11', value: 31 },
  { date: '12', value: 27 },
  { date: '13', value: 33 },
  { date: '14', value: 29 },
  { date: '15', value: 35 },
  { date: '16', value: 32 },
  { date: '17', value: 38 },
  { date: '18', value: 34 },
  { date: '19', value: 41 },
  { date: '20', value: 37 },
  { date: '21', value: 43 },
  { date: '22', value: 39 },
  { date: '23', value: 45 },
  { date: '24', value: 42 },
  { date: '25', value: 48 },
  { date: '26', value: 44 },
  { date: '27', value: 50 },
  { date: '28', value: 47 },
  { date: '29', value: 53 },
  { date: '30', value: 49 }
])

// دلایل مسترد برای نمودار
const refundReasonsChart = ref([
  { id: 1, name: 'عدم رضایت از محصول', percentage: 45, color: '#EF4444' },
  { id: 2, name: 'تأخیر در ارسال', percentage: 25, color: '#F59E0B' },
  { id: 3, name: 'مشکل در کیفیت', percentage: 20, color: '#10B981' },
  { id: 4, name: 'سایر موارد', percentage: 10, color: '#6B7280' }
])

// گزارشات تفصیلی
const detailedReports = ref([
  {
    period: 'هفته جاری',
    count: 23,
    totalAmount: 6800000,
    refundRate: 3.2,
    averageAmount: 295652,
    change: 27.8
  },
  {
    period: 'هفته گذشته',
    count: 18,
    totalAmount: 5200000,
    refundRate: 2.8,
    averageAmount: 288889,
    change: -5.2
  },
  {
    period: 'ماه جاری',
    count: 156,
    totalAmount: 45200000,
    refundRate: 3.1,
    averageAmount: 289744,
    change: 12.5
  },
  {
    period: 'ماه گذشته',
    count: 138,
    totalAmount: 39800000,
    refundRate: 2.9,
    averageAmount: 288406,
    change: 8.7
  }
])

// متغیرهای فیلتر و جستجو
const showFilters = ref(false)
const filters = ref({
  period: '',
  refundRate: '',
  amountRange: '',
  change: '',
  countRange: '',
  averageAmountRange: ''
})

// داده‌های فیلتر شده
const filteredReports = computed(() => {
  let reports = [...detailedReports.value]

  // فیلتر دوره
  if (filters.value.period) {
    reports = reports.filter(report => {
      if (filters.value.period === 'current-week') {
        return report.period === 'هفته جاری'
      } else if (filters.value.period === 'last-week') {
        return report.period === 'هفته گذشته'
      } else if (filters.value.period === 'current-month') {
        return report.period === 'ماه جاری'
      } else if (filters.value.period === 'last-month') {
        return report.period === 'ماه گذشته'
      }
      return true
    })
  }

  // فیلتر نرخ مسترد
  if (filters.value.refundRate) {
    reports = reports.filter(report => {
      const rate = parseFloat(report.refundRate)
      if (filters.value.refundRate === 'low') {
        return rate < 2
      } else if (filters.value.refundRate === 'medium') {
        return rate >= 2 && rate <= 5
      } else if (filters.value.refundRate === 'high') {
        return rate > 5
      }
      return true
    })
  }

  // فیلتر مبلغ کل
  if (filters.value.amountRange) {
    const [min, max] = filters.value.amountRange.split('-')
    if (max === '+') {
      // بیش از مقدار مشخص
      reports = reports.filter(report => report.totalAmount >= parseInt(min))
    } else {
      // بین دو مقدار
      reports = reports.filter(report => 
        report.totalAmount >= parseInt(min) && report.totalAmount <= parseInt(max)
      )
    }
  }

  // فیلتر تغییر
  if (filters.value.change) {
    reports = reports.filter(report => {
      if (filters.value.change === 'positive') {
        return report.change > 0
      } else if (filters.value.change === 'negative') {
        return report.change < 0
      } else if (filters.value.change === 'neutral') {
        return report.change === 0
      }
      return true
    })
  }

  // فیلتر تعداد مسترد
  if (filters.value.countRange) {
    const [min, max] = filters.value.countRange.split('-')
    if (max === '+') {
      // بیش از مقدار مشخص
      reports = reports.filter(report => report.count >= parseInt(min))
    } else {
      // بین دو مقدار
      reports = reports.filter(report => 
        report.count >= parseInt(min) && report.count <= parseInt(max)
      )
    }
  }

  // فیلتر متوسط مبلغ
  if (filters.value.averageAmountRange) {
    const [min, max] = filters.value.averageAmountRange.split('-')
    if (max === '+') {
      // بیش از مقدار مشخص
      reports = reports.filter(report => report.averageAmount >= parseInt(min))
    } else {
      // بین دو مقدار
      reports = reports.filter(report => 
        report.averageAmount >= parseInt(min) && report.averageAmount <= parseInt(max)
      )
    }
  }

  return reports
})

// تعداد آیتم‌ها در هر صفحه
const itemsPerPage = ref(10)

// صفحه فعلی
const currentPage = ref(1)

// تعداد کل صفحات
const totalPages = computed(() => {
  return Math.ceil(filteredReports.value.length / itemsPerPage.value)
})

// داده‌های صفحه بندی شده
const paginatedReports = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredReports.value.slice(start, end)
})

// متدهای عملیاتی صفحه‌بندی
const handlePageChange = (page) => {
  currentPage.value = page
}

const handleItemsPerPageChange = (items) => {
  itemsPerPage.value = items
  currentPage.value = 1 // بازگشت به صفحه اول پس از تغییر تعداد آیتم‌ها
}

// متدهای فیلتر
const clearFilters = () => {
  filters.value = {
    period: '',
    refundRate: '',
    amountRange: '',
    change: '',
    countRange: '',
    averageAmountRange: ''
  }
  currentPage.value = 1
}

// نظارت بر تغییرات فیلترها و بازگشت به صفحه اول
watch(filters, () => {
  currentPage.value = 1
}, { deep: true })

// محاسبه حداکثر مقدار نمودار روند
const maxTrendValue = computed(() => {
  return Math.max(...trendChartData.value.map(item => item.value))
})

// تابع فرمت قیمت
const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}
</script> 
