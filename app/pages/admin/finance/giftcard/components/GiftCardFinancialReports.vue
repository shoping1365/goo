<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">گزارش‌های مالی گیفت کارت</h2>
          <p class="text-gray-600 mt-1">تحلیل و گزارش‌گیری از وضعیت مالی سیستم گیفت کارت</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            @click="exportReport"
          >
            خروجی PDF
          </button>
          <button 
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="exportExcel"
          >
            خروجی Excel
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترهای گزارش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">دوره زمانی</label>
          <select 
            v-model="reportFilters.period"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            @change="updateReportData"
          >
            <option value="today">امروز</option>
            <option value="week">هفته جاری</option>
            <option value="month">ماه جاری</option>
            <option value="quarter">سه ماهه</option>
            <option value="year">سال جاری</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        
        <div v-if="reportFilters.period === 'custom'">
          <label class="block text-sm font-medium text-gray-700 mb-1">از تاریخ</label>
          <input
            v-model="reportFilters.dateFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            @change="updateReportData"
          />
        </div>
        
        <div v-if="reportFilters.period === 'custom'">
          <label class="block text-sm font-medium text-gray-700 mb-1">تا تاریخ</label>
          <input
            v-model="reportFilters.dateTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            @change="updateReportData"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع گزارش</label>
          <select 
            v-model="reportFilters.type"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            @change="updateReportData"
          >
            <option value="summary">خلاصه کلی</option>
            <option value="sales">فروش</option>
            <option value="usage">استفاده</option>
            <option value="refunds">بازپرداخت‌ها</option>
            <option value="expiry">انقضا</option>
          </select>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل درآمد</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(totalRevenue) }}</p>
            <p class="text-xs text-green-600">+{{ revenueGrowth }}% نسبت به دوره قبل</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">تعداد فروش</p>
            <p class="text-2xl font-bold text-gray-900">{{ totalSales }}</p>
            <p class="text-xs text-blue-600">+{{ salesGrowth }}% نسبت به دوره قبل</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">مبلغ استفاده شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(totalUsed) }}</p>
            <p class="text-xs text-yellow-600">{{ usagePercentage }}% از کل کارت‌ها</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">بازپرداخت‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(totalRefunds) }}</p>
            <p class="text-xs text-red-600">{{ refundPercentage }}% از کل فروش</p>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار فروش ماهانه -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">فروش ماهانه</h3>
        <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
          <div 
            v-for="(month, index) in monthlySales" 
            :key="index"
            class="flex-1 bg-blue-500 rounded-t"
            :style="{ height: `${(month.amount / maxMonthlySales) * 100}%` }"
            :title="`${month.name}: ${formatCurrency(month.amount)}`"
          ></div>
        </div>
        <div class="mt-4 grid grid-cols-6 gap-2 text-xs text-gray-600">
          <div v-for="(month, index) in monthlySales" :key="index" class="text-center">
            {{ month.name }}
          </div>
        </div>
      </div>

      <!-- نمودار دایره‌ای انواع کارت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">توزیع انواع کارت</h3>
        <div class="flex items-center justify-center h-64">
          <div class="relative w-48 h-48">
            <svg class="w-full h-full" viewBox="0 0 100 100">
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#e5e7eb"
                stroke-width="8"
              />
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#3b82f6"
                stroke-width="8"
                stroke-dasharray="251.2"
                :stroke-dashoffset="251.2 - (251.2 * cardTypeDistribution.birthday / 100)"
                stroke-linecap="round"
                transform="rotate(-90 50 50)"
              />
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#10b981"
                stroke-width="8"
                stroke-dasharray="251.2"
                :stroke-dashoffset="251.2 - (251.2 * cardTypeDistribution.wedding / 100)"
                stroke-linecap="round"
                transform="rotate(-90 50 50)"
              />
              <circle
                cx="50"
                cy="50"
                r="40"
                fill="none"
                stroke="#f59e0b"
                stroke-width="8"
                stroke-dasharray="251.2"
                :stroke-dashoffset="251.2 - (251.2 * cardTypeDistribution.newyear / 100)"
                stroke-linecap="round"
                transform="rotate(-90 50 50)"
              />
            </svg>
            <div class="absolute inset-0 flex items-center justify-center">
              <div class="text-center">
                <p class="text-2xl font-bold text-gray-900">{{ totalCards }}</p>
                <p class="text-sm text-gray-600">کل کارت‌ها</p>
              </div>
            </div>
          </div>
        </div>
        <div class="mt-4 space-y-2">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-blue-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-700">تولد</span>
            </div>
            <span class="text-sm font-medium text-gray-900">{{ cardTypeDistribution.birthday }}%</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-green-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-700">عروسی</span>
            </div>
            <span class="text-sm font-medium text-gray-900">{{ cardTypeDistribution.wedding }}%</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-yellow-500 rounded-full ml-2"></div>
              <span class="text-sm text-gray-700">سال نو</span>
            </div>
            <span class="text-sm font-medium text-gray-900">{{ cardTypeDistribution.newyear }}%</span>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول جزئیات -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">جزئیات تراکنش‌ها</h3>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نوع
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تعداد
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مبلغ کل
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مبلغ متوسط
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                درصد رشد
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(detail, index) in transactionDetails" :key="index" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(detail.date) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getTypeBadgeClass(detail.type)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getTypeLabel(detail.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ detail.count }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatCurrency(detail.totalAmount) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatCurrency(detail.averageAmount) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm">
                <span :class="detail.growth >= 0 ? 'text-green-600' : 'text-red-600'">
                  {{ detail.growth >= 0 ? '+' : '' }}{{ detail.growth }}%
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- تحلیل‌های پیشرفته -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- بهترین ساعات فروش -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">بهترین ساعات فروش</h3>
        <div class="space-y-3">
          <div v-for="hour in bestSellingHours" :key="hour.hour" class="flex items-center justify-between">
            <span class="text-sm text-gray-700">{{ hour.hour }}</span>
            <div class="flex items-center">
              <div class="w-24 bg-gray-200 rounded-full h-2 ml-2">
                <div 
                  class="bg-blue-600 h-2 rounded-full" 
                  :style="{ width: `${(hour.sales / maxHourlySales) * 100}%` }"
                ></div>
              </div>
              <span class="text-sm font-medium text-gray-900">{{ hour.sales }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- محبوب‌ترین مبالغ -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">محبوب‌ترین مبالغ</h3>
        <div class="space-y-3">
          <div v-for="amount in popularAmounts" :key="amount.amount" class="flex items-center justify-between">
            <span class="text-sm text-gray-700">{{ formatCurrency(amount.amount) }}</span>
            <div class="flex items-center">
              <div class="w-24 bg-gray-200 rounded-full h-2 ml-2">
                <div 
                  class="bg-green-600 h-2 rounded-full" 
                  :style="{ width: `${(amount.count / maxAmountCount) * 100}%` }"
                ></div>
              </div>
              <span class="text-sm font-medium text-gray-900">{{ amount.count }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نرخ تبدیل -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">نرخ تبدیل</h3>
        <div class="space-y-4">
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span class="text-gray-700">بازدید صفحه</span>
              <span class="font-medium text-gray-900">{{ conversionMetrics.pageViews }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-blue-600 h-2 rounded-full" style="width: 100%"></div>
            </div>
          </div>
          
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span class="text-gray-700">افزودن به سبد</span>
              <span class="font-medium text-gray-900">{{ conversionMetrics.addToCart }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="bg-yellow-600 h-2 rounded-full" 
                :style="{ width: `${(conversionMetrics.addToCart / conversionMetrics.pageViews) * 100}%` }"
              ></div>
            </div>
          </div>
          
          <div>
            <div class="flex justify-between text-sm mb-1">
              <span class="text-gray-700">خرید نهایی</span>
              <span class="font-medium text-gray-900">{{ conversionMetrics.purchases }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="bg-green-600 h-2 rounded-full" 
                :style="{ width: `${(conversionMetrics.purchases / conversionMetrics.pageViews) * 100}%` }"
              ></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'

// Reactive data
const reportFilters = reactive({
  period: 'month',
  dateFrom: '',
  dateTo: '',
  type: 'summary'
})

// داده‌های نمونه
const totalRevenue = ref(15000000)
const revenueGrowth = ref(12.5)
const totalSales = ref(1250)
const salesGrowth = ref(8.3)
const totalUsed = ref(8500000)
const usagePercentage = ref(56.7)
const totalRefunds = ref(750000)
const refundPercentage = ref(5.0)

const monthlySales = ref([
  { name: 'فروردین', amount: 1200000 },
  { name: 'اردیبهشت', amount: 1500000 },
  { name: 'خرداد', amount: 1800000 },
  { name: 'تیر', amount: 1400000 },
  { name: 'مرداد', amount: 1600000 },
  { name: 'شهریور', amount: 2000000 }
])

const cardTypeDistribution = reactive({
  birthday: 45,
  wedding: 35,
  newyear: 20
})

const totalCards = ref(2500)

interface TransactionDetail {
  date: Date
  type: string
  count: number
  totalAmount: number
  averageAmount: number
  growth: number
}

const transactionDetails = ref<TransactionDetail[]>([
  {
    date: new Date(Date.now() - 86400000),
    type: 'purchase',
    count: 45,
    totalAmount: 2250000,
    averageAmount: 50000,
    growth: 12.5
  },
  {
    date: new Date(Date.now() - 172800000),
    type: 'usage',
    count: 38,
    totalAmount: 1900000,
    averageAmount: 50000,
    growth: 8.3
  },
  {
    date: new Date(Date.now() - 259200000),
    type: 'refund',
    count: 3,
    totalAmount: 150000,
    averageAmount: 50000,
    growth: -5.2
  }
])

const bestSellingHours = ref([
  { hour: '09:00-11:00', sales: 45 },
  { hour: '14:00-16:00', sales: 38 },
  { hour: '19:00-21:00', sales: 52 },
  { hour: '21:00-23:00', sales: 28 }
])

const popularAmounts = ref([
  { amount: 50000, count: 125 },
  { amount: 100000, count: 89 },
  { amount: 200000, count: 67 },
  { amount: 500000, count: 34 }
])

const conversionMetrics = reactive({
  pageViews: 12500,
  addToCart: 1875,
  purchases: 1250
})

// Computed properties
const maxMonthlySales = computed(() => Math.max(...monthlySales.value.map(m => m.amount)))
const maxHourlySales = computed(() => Math.max(...bestSellingHours.value.map(h => h.sales)))
const maxAmountCount = computed(() => Math.max(...popularAmounts.value.map(a => a.count)))

// Methods
const updateReportData = () => {
  // به‌روزرسانی داده‌ها بر اساس فیلترها

}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR').format(date)
}

const getTypeBadgeClass = (type: string) => {
  const classes = {
    purchase: 'bg-blue-100 text-blue-800',
    usage: 'bg-green-100 text-green-800',
    refund: 'bg-yellow-100 text-yellow-800',
    expiry: 'bg-red-100 text-red-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTypeLabel = (type: string) => {
  const labels = {
    purchase: 'خرید',
    usage: 'استفاده',
    refund: 'بازپرداخت',
    expiry: 'انقضا'
  }
  return labels[type] || type
}

const exportReport = () => {
  // شبیه‌سازی export PDF

  alert('گزارش PDF آماده شد')
}

const exportExcel = () => {
  // شبیه‌سازی export Excel

  alert('گزارش Excel آماده شد')
}

// Lifecycle
onMounted(() => {

  updateReportData()
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
