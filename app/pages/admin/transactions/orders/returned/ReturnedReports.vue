<template>
  <div v-if="hasAccess">
    <!-- آمار کلی گزارشات -->
    <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4 mb-6">
      <!-- کل مرجوع -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center ml-4">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
            </svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-600">کل مرجوع</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.totalReturns || 0 }}</p>
            <p class="text-xs text-green-600">+12% از ماه گذشته</p>
          </div>
        </div>
      </div>

      <!-- مبلغ کل -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-red-500 rounded-lg flex items-center justify-center ml-4">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-600">مبلغ کل</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatPrice(stats.totalAmount || 0) }}</p>
            <p class="text-xs text-red-600">+8% از ماه گذشته</p>
          </div>
        </div>
      </div>

      <!-- نرخ مرجوع -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-yellow-500 rounded-lg flex items-center justify-center ml-4">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-600">نرخ مرجوع</p>
            <p class="text-2xl font-bold text-gray-900">{{ stats.returnRate || 0 }}%</p>
            <p class="text-xs text-yellow-600">-2% از ماه گذشته</p>
          </div>
        </div>
      </div>

      <!-- متوسط مبلغ -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center ml-4">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
            </svg>
          </div>
          <div>
            <p class="text-sm font-medium text-gray-600">متوسط مبلغ</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatPrice(stats.averageAmount || 0) }}</p>
            <p class="text-xs text-blue-600">+5% از ماه گذشته</p>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند مرجوع -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-orange-50 to-red-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-orange-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">روند مرجوع در 30 روز گذشته</h3>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              v-for="period in chartPeriods"
              :key="period.id"
              :class="[
                'px-3 py-1 text-xs font-medium rounded-md transition-colors',
                selectedChartPeriod === period.id
                  ? 'bg-orange-500 text-white'
                  : 'bg-white text-gray-600 hover:bg-gray-50'
              ]"
              @click="selectedChartPeriod = period.id"
            >
              {{ period.name }}
            </button>
          </div>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="h-64 flex items-end justify-between space-x-1 space-x-reverse">
          <div
            v-for="(day, index) in trendChartData"
            :key="index"
            class="flex-1 flex flex-col items-center"
          >
            <div class="text-xs text-gray-500 mb-2">{{ day.date }}</div>
            <div
              class="w-full bg-gradient-to-t from-orange-500 to-orange-400 rounded-t"
              :style="{ height: `${(day.value / maxTrendValue) * 200}px` }"
            ></div>
            <div class="text-xs font-medium text-gray-700 mt-1">{{ day.value }}</div>
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
          <button class="text-sm text-indigo-600 hover:text-indigo-800 transition-colors font-medium hover:bg-indigo-50 px-3 py-1 rounded-lg" @click="showFilters = !showFilters">
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
          
          <!-- فیلتر نرخ مرجوع -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ مرجوع</label>
            <select v-model="filters.returnRate" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
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
          <!-- فیلتر تعداد مرجوع -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد مرجوع</label>
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
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              @click="clearFilters"
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
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد مرجوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ کل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ مرجوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">متوسط مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(report, index) in paginatedReports" :key="report.period || index" class="hover:bg-gray-50">
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
                {{ report.returnRate || 0 }}%
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

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
// Import کامپوننت Pagination
import { computed, onMounted, ref, watch } from 'vue';
import Pagination from '~/components/admin/common/Pagination.vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

// داده‌های آماری
const stats = ref({
  totalReturns: 156,
  totalAmount: 45200000,
  returnRate: 3.2,
  averageAmount: 289743
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
  { date: '1', value: 8 },
  { date: '2', value: 12 },
  { date: '3', value: 10 },
  { date: '4', value: 15 },
  { date: '5', value: 13 },
  { date: '6', value: 18 },
  { date: '7', value: 11 },
  { date: '8', value: 16 },
  { date: '9', value: 20 },
  { date: '10', value: 17 },
  { date: '11', value: 22 },
  { date: '12', value: 19 },
  { date: '13', value: 25 },
  { date: '14', value: 21 },
  { date: '15', value: 28 },
  { date: '16', value: 24 },
  { date: '17', value: 31 },
  { date: '18', value: 27 },
  { date: '19', value: 34 },
  { date: '20', value: 30 },
  { date: '21', value: 36 },
  { date: '22', value: 33 },
  { date: '23', value: 39 },
  { date: '24', value: 35 },
  { date: '25', value: 42 },
  { date: '26', value: 38 },
  { date: '27', value: 45 },
  { date: '28', value: 41 },
  { date: '29', value: 48 },
  { date: '30', value: 44 }
])

// داده‌های گزارشات تفصیلی
const detailedReports = ref([
  {
    period: 'هفته جاری',
    count: 45,
    totalAmount: 12500000,
    returnRate: 2.8,
    averageAmount: 277778,
    change: 12
  },
  {
    period: 'هفته گذشته',
    count: 40,
    totalAmount: 11200000,
    returnRate: 2.5,
    averageAmount: 280000,
    change: -5
  },
  {
    period: 'ماه جاری',
    count: 156,
    totalAmount: 45200000,
    returnRate: 3.2,
    averageAmount: 289743,
    change: 8
  },
  {
    period: 'ماه گذشته',
    count: 144,
    totalAmount: 41800000,
    returnRate: 3.0,
    averageAmount: 290278,
    change: -2
  },
  {
    period: 'سه ماه گذشته',
    count: 432,
    totalAmount: 125000000,
    returnRate: 3.1,
    averageAmount: 289352,
    change: 5
  },
  {
    period: 'شش ماه گذشته',
    count: 856,
    totalAmount: 248000000,
    returnRate: 3.3,
    averageAmount: 289720,
    change: 15
  },
  {
    period: 'سال گذشته',
    count: 1654,
    totalAmount: 478000000,
    returnRate: 3.4,
    averageAmount: 288997,
    change: 22
  }
])

// متغیرهای فیلتر و جستجو
const showFilters = ref(false)
const filters = ref({
  period: '',
  returnRate: '',
  amountRange: '',
  change: '',
  countRange: '',
  averageAmountRange: ''
})

// تعداد آیتم‌ها در هر صفحه
const itemsPerPage = ref(10)

// صفحه فعلی
const currentPage = ref(1)

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

  // فیلتر نرخ مرجوع
  if (filters.value.returnRate) {
    reports = reports.filter(report => {
      const rate = typeof report.returnRate === 'number' ? report.returnRate : parseFloat(String(report.returnRate))
      if (filters.value.returnRate === 'low') {
        return rate < 2
      } else if (filters.value.returnRate === 'medium') {
        return rate >= 2 && rate <= 5
      } else if (filters.value.returnRate === 'high') {
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

  // فیلتر تعداد مرجوع
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
    returnRate: '',
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
