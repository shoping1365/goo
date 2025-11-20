<template>
  <div>
    <!-- آمار کلی گزارشات -->
    <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">آمار کلی گزارشات ارسال</h3>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
          <!-- کل ارسال‌ها -->
          <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-green-700 mb-1">کل ارسال‌ها</p>
                <p class="text-2xl font-bold text-green-900">{{ reportStats.totalShipped?.toLocaleString('fa-IR') || '0' }}</p>
                <p class="text-xs text-green-600 mt-1">در 30 روز گذشته</p>
              </div>
              <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-14 0h14"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- نرخ تحویل -->
          <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-blue-700 mb-1">نرخ تحویل</p>
                <p class="text-2xl font-bold text-blue-900">{{ reportStats.deliveryRate }}%</p>
                <p class="text-xs text-blue-600 mt-1">از کل ارسال‌ها</p>
              </div>
              <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- متوسط زمان تحویل -->
          <div class="bg-gradient-to-br from-emerald-50 to-emerald-100 border border-emerald-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-emerald-700 mb-1">متوسط زمان تحویل</p>
                <p class="text-2xl font-bold text-emerald-900">{{ reportStats.avgDeliveryTime }} روز</p>
                <p class="text-xs text-emerald-600 mt-1">از ارسال تا تحویل</p>
              </div>
              <div class="w-12 h-12 bg-emerald-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- رضایت مشتری -->
          <div class="bg-gradient-to-br from-yellow-50 to-yellow-100 border border-yellow-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-yellow-700 mb-1">رضایت مشتری</p>
                <p class="text-2xl font-bold text-yellow-900">{{ reportStats.customerSatisfaction }}%</p>
                <p class="text-xs text-yellow-600 mt-1">از نظرات دریافتی</p>
              </div>
              <div class="w-12 h-12 bg-yellow-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند 30 روزه -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4 mb-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-gray-900">روند ارسال در 30 روز گذشته</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button
            v-for="period in chartPeriods"
            :key="period.id"
            :class="[
              'px-3 py-1 text-xs font-medium rounded-md transition-colors',
              selectedChartPeriod === period.id
                ? 'bg-green-500 text-white'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="selectedChartPeriod = period.id"
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
            class="w-full bg-gradient-to-t from-green-500 to-green-400 rounded-t"
            :style="{ height: `${(day.value / maxTrendValue) * 200}px` }"
          ></div>
          <div class="text-xs font-medium text-gray-700 mt-1">{{ day.value }}</div>
        </div>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته گزارشات</h3>
          </div>
          <button class="text-sm text-green-600 hover:text-green-800 transition-colors font-medium hover:bg-green-50 px-3 py-1 rounded-lg" @click="showFilters = !showFilters">
            {{ showFilters ? 'مخفی کردن' : 'نمایش' }}
          </button>
        </div>
      </div>

      <div v-if="showFilters" class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
          <!-- فیلتر دوره -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
            <select v-model="filters.period" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="">همه دوره‌ها</option>
              <option value="7d">7 روز گذشته</option>
              <option value="30d">30 روز گذشته</option>
              <option value="90d">90 روز گذشته</option>
              <option value="1y">1 سال گذشته</option>
            </select>
          </div>

          <!-- فیلتر نرخ تحویل -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ تحویل</label>
            <select v-model="filters.deliveryRate" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="">همه نرخ‌ها</option>
              <option value="0-50">کمتر از 50%</option>
              <option value="50-80">50% تا 80%</option>
              <option value="80-95">80% تا 95%</option>
              <option value="95-100">بیش از 95%</option>
            </select>
          </div>

          <!-- فیلتر محدوده مبلغ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدوده مبلغ</label>
            <select v-model="filters.amountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="">همه مبالغ</option>
              <option value="0-1000000">کمتر از 1 میلیون تومان</option>
              <option value="1000000-5000000">1 تا 5 میلیون تومان</option>
              <option value="5000000-10000000">5 تا 10 میلیون تومان</option>
              <option value="10000000+">بیش از 10 میلیون تومان</option>
            </select>
          </div>

          <!-- فیلتر تغییر -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع تغییر</label>
            <select v-model="filters.change" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="">همه تغییرات</option>
              <option value="increase">افزایش</option>
              <option value="decrease">کاهش</option>
              <option value="stable">ثابت</option>
            </select>
          </div>

          <!-- فیلتر محدوده تعداد -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدوده تعداد</label>
            <select v-model="filters.countRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="">همه تعدادها</option>
              <option value="0-10">کمتر از 10</option>
              <option value="10-50">10 تا 50</option>
              <option value="50-100">50 تا 100</option>
              <option value="100+">بیش از 100</option>
            </select>
          </div>

          <!-- فیلتر محدوده متوسط مبلغ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">متوسط مبلغ</label>
            <select v-model="filters.averageAmountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="">همه متوسط‌ها</option>
              <option value="0-500000">کمتر از 500 هزار تومان</option>
              <option value="500000-1000000">500 هزار تا 1 میلیون تومان</option>
              <option value="1000000-2000000">1 تا 2 میلیون تومان</option>
              <option value="2000000+">بیش از 2 میلیون تومان</option>
            </select>
          </div>
        </div>

        <div class="mt-4 flex justify-end">
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors" @click="clearFilters">
            پاک کردن فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- جدول گزارشات تفصیلی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">گزارشات تفصیلی ارسال</h3>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <!-- Search -->
            <div class="relative">
              <input 
                v-model="searchTerm"
                type="text" 
                class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-xs shadow-sm" 
                placeholder="جستجو در گزارشات..."
                dir="rtl"
              />
              <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
                <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Table Content -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دوره</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد ارسال</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ کل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تحویل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">متوسط زمان تحویل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in paginatedReports" :key="report.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ report.period || 'نامشخص' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.shippedCount?.toLocaleString('fa-IR') || '0' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatPrice(report.totalAmount || 0) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getDeliveryRateClass(report.deliveryRate)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ report.deliveryRate || 0 }}%
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.avgDeliveryTime || 0 }} روز
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getChangeClass(report.change)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ (report.change || 0) > 0 ? '+' : '' }}{{ report.change || 0 }}%
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900" @click="viewReport(report)">جزئیات</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
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

// داده‌های آماری گزارشات
const reportStats = ref({
  totalShipped: 234,
  deliveryRate: 89.5,
  avgDeliveryTime: 2.8,
  customerSatisfaction: 92.3
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
  { date: '13', value: 34 },
  { date: '14', value: 30 },
  { date: '15', value: 37 },
  { date: '16', value: 33 },
  { date: '17', value: 40 },
  { date: '18', value: 36 },
  { date: '19', value: 43 },
  { date: '20', value: 39 },
  { date: '21', value: 45 },
  { date: '22', value: 42 },
  { date: '23', value: 48 },
  { date: '24', value: 44 },
  { date: '25', value: 51 },
  { date: '26', value: 47 },
  { date: '27', value: 54 },
  { date: '28', value: 50 },
  { date: '29', value: 57 },
  { date: '30', value: 53 }
])

// داده‌های گزارشات تفصیلی
const detailedReports = ref([
  {
    id: 1,
    period: 'هفته گذشته',
    shippedCount: 156,
    totalAmount: 45600000,
    deliveryRate: 92.3,
    avgDeliveryTime: 2.5,
    change: 5.2
  },
  {
    id: 2,
    period: 'ماه گذشته',
    shippedCount: 678,
    totalAmount: 189000000,
    deliveryRate: 88.7,
    avgDeliveryTime: 3.1,
    change: -2.1
  },
  {
    id: 3,
    period: '3 ماه گذشته',
    shippedCount: 1890,
    totalAmount: 523000000,
    deliveryRate: 91.2,
    avgDeliveryTime: 2.8,
    change: 3.4
  },
  {
    id: 4,
    period: '6 ماه گذشته',
    shippedCount: 3456,
    totalAmount: 987000000,
    deliveryRate: 89.8,
    avgDeliveryTime: 3.2,
    change: 1.7
  },
  {
    id: 5,
    period: 'سال گذشته',
    shippedCount: 6789,
    totalAmount: 1890000000,
    deliveryRate: 87.5,
    avgDeliveryTime: 3.5,
    change: -1.2
  },
  {
    id: 6,
    period: 'امروز',
    shippedCount: 23,
    totalAmount: 6700000,
    deliveryRate: 95.7,
    avgDeliveryTime: 2.1,
    change: 8.3
  },
  {
    id: 7,
    period: 'دیروز',
    shippedCount: 45,
    totalAmount: 12300000,
    deliveryRate: 93.2,
    avgDeliveryTime: 2.3,
    change: 4.1
  },
  {
    id: 8,
    period: 'هفته جاری',
    shippedCount: 234,
    totalAmount: 67800000,
    deliveryRate: 90.1,
    avgDeliveryTime: 2.7,
    change: 2.8
  }
])

// متغیرهای فیلتر و جستجو
const showFilters = ref(false)
const searchTerm = ref('')
const filters = ref({
  period: '',
  deliveryRate: '',
  amountRange: '',
  change: '',
  countRange: '',
  averageAmountRange: ''
})

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)

// محاسبه حداکثر مقدار نمودار روند
const maxTrendValue = computed(() => {
  return Math.max(...trendChartData.value.map(item => item.value))
})

// محاسبه گزارشات فیلتر شده
const filteredReports = computed(() => {
  let filtered = detailedReports.value

  // فیلتر بر اساس جستجو
  if (searchTerm.value) {
    const search = searchTerm.value.toLowerCase()
    filtered = filtered.filter(report => 
      report.period.toLowerCase().includes(search)
    )
  }

  // فیلتر بر اساس دوره
  if (filters.value.period) {
    filtered = filtered.filter(_report => {
      // اینجا می‌توانید منطق فیلتر بر اساس دوره را پیاده‌سازی کنید
      return true
    })
  }

  // فیلتر بر اساس نرخ تحویل
  if (filters.value.deliveryRate) {
    const [min, max] = filters.value.deliveryRate.split('-').map(Number)
    filtered = filtered.filter(report => {
      if (max) {
        return report.deliveryRate >= min && report.deliveryRate <= max
      } else {
        return report.deliveryRate >= min
      }
    })
  }

  // فیلتر بر اساس محدوده مبلغ
  if (filters.value.amountRange) {
    const [min, max] = filters.value.amountRange.split('-').map(Number)
    filtered = filtered.filter(report => {
      if (max) {
        return report.totalAmount >= min && report.totalAmount <= max
      } else {
        return report.totalAmount >= min
      }
    })
  }

  // فیلتر بر اساس نوع تغییر
  if (filters.value.change) {
    filtered = filtered.filter(report => {
      switch (filters.value.change) {
        case 'increase':
          return report.change > 0
        case 'decrease':
          return report.change < 0
        case 'stable':
          return report.change === 0
        default:
          return true
      }
    })
  }

  // فیلتر بر اساس محدوده تعداد
  if (filters.value.countRange) {
    const [min, max] = filters.value.countRange.split('-').map(Number)
    filtered = filtered.filter(report => {
      if (max) {
        return report.shippedCount >= min && report.shippedCount <= max
      } else {
        return report.shippedCount >= min
      }
    })
  }

  // فیلتر بر اساس محدوده متوسط مبلغ
  if (filters.value.averageAmountRange) {
    const [min, max] = filters.value.averageAmountRange.split('-').map(Number)
    filtered = filtered.filter(report => {
      const avgAmount = report.totalAmount / report.shippedCount
      if (max) {
        return avgAmount >= min && avgAmount <= max
      } else {
        return avgAmount >= min
      }
    })
  }

  return filtered
})

// محاسبه تعداد صفحات
const totalPages = computed(() => {
  return Math.ceil(filteredReports.value.length / itemsPerPage.value)
})

// محاسبه گزارشات نمایش داده شده در صفحه فعلی
const paginatedReports = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredReports.value.slice(start, end)
})

// Watch برای تغییرات فیلترها
watch(filters, () => {
  currentPage.value = 1 // بازگشت به صفحه اول
}, { deep: true })

// تابع تغییر صفحه
const handlePageChange = (page) => {
  currentPage.value = page
}

// تابع تغییر تعداد آیتم در هر صفحه
const handleItemsPerPageChange = (newItemsPerPage) => {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1
}

// تابع پاک کردن فیلترها
const clearFilters = () => {
  filters.value = {
    period: '',
    deliveryRate: '',
    amountRange: '',
    change: '',
    countRange: '',
    averageAmountRange: ''
  }
  searchTerm.value = ''
}

// تابع مشاهده گزارش
const viewReport = (_report) => {
  // View report logic
}

// تابع فرمت قیمت
const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

// تابع دریافت کلاس نرخ تحویل
const getDeliveryRateClass = (rate) => {
  const rateValue = rate || 0
  if (rateValue >= 90) return 'bg-green-100 text-green-800'
  if (rateValue >= 80) return 'bg-blue-100 text-blue-800'
  if (rateValue >= 70) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

// تابع دریافت کلاس تغییر
const getChangeClass = (change) => {
  const changeValue = change || 0
  if (changeValue > 0) return 'bg-green-100 text-green-800'
  if (changeValue < 0) return 'bg-red-100 text-red-800'
  return 'bg-gray-100 text-gray-800'
}
</script> 
