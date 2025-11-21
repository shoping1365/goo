<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">📊 گزارش‌های مالی</h2>
      <p class="text-gray-600">گزارش‌های جامع مالی و تحلیلی کیف پول</p>
    </div>

    <!-- آمار کلی مالی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- درآمد کل -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">درآمد کل</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">مثبت</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ formatCurrency(financialStats.totalRevenue) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ financialStats.revenueGrowth }}%</span>
          <span class="mx-2">از ماه قبل</span>
        </div>
      </div>

      <!-- سود خالص -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">سود خالص</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">سود</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ formatCurrency(financialStats.netProfit) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-blue-500">{{ financialStats.profitMargin }}%</span>
          <span class="mx-2">حاشیه سود</span>
        </div>
      </div>

      <!-- کارمزدها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کارمزدها</h3>
          <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">درآمد</span>
        </div>
        <div class="text-3xl font-bold text-purple-600 mb-2">{{ formatCurrency(financialStats.totalFees) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-purple-500">{{ financialStats.feePercentage }}%</span>
          <span class="mx-2">از کل درآمد</span>
        </div>
      </div>

      <!-- هزینه‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">هزینه‌ها</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">هزینه</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ formatCurrency(financialStats.totalExpenses) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ financialStats.expensePercentage }}%</span>
          <span class="mx-2">از کل درآمد</span>
        </div>
      </div>
    </div>

    <!-- گزارش‌های دوره‌ای -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- گزارش درآمد روزانه -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">درآمد روزانه</h3>
          <select class="px-3 py-1 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>7 روز گذشته</option>
            <option>30 روز گذشته</option>
            <option>90 روز گذشته</option>
          </select>
        </div>
        
        <div class="space-y-3">
          <div v-for="day in dailyRevenue" :key="day.date" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-green-500 rounded-full mr-3"></div>
              <span class="text-sm text-gray-700">{{ day.date }}</span>
            </div>
            <div class="text-right">
              <div class="text-sm font-medium text-gray-900">{{ formatCurrency(day.revenue) }}</div>
              <div class="text-xs text-gray-500">{{ day.transactions }} تراکنش</div>
            </div>
          </div>
        </div>
      </div>

      <!-- گزارش درآمد ماهانه -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">درآمد ماهانه</h3>
          <select class="px-3 py-1 text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>6 ماه گذشته</option>
            <option>12 ماه گذشته</option>
            <option>24 ماه گذشته</option>
          </select>
        </div>
        
        <div class="space-y-3">
          <div v-for="month in monthlyRevenue" :key="month.month" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-blue-500 rounded-full mr-3"></div>
              <span class="text-sm text-gray-700">{{ month.month }}</span>
            </div>
            <div class="text-right">
              <div class="text-sm font-medium text-gray-900">{{ formatCurrency(month.revenue) }}</div>
              <div class="text-xs text-gray-500">{{ month.growth > 0 ? '+' : '' }}{{ month.growth }}%</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار روند درآمد -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">روند درآمد (30 روز گذشته)</h3>
        <div class="h-64 flex items-end space-x-2 space-x-reverse overflow-x-auto">
          <div v-for="(day, index) in revenueTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-12">
            <div
class="w-full bg-gray-200 rounded-t relative"
                 :style="{ height: getChartHeight(day.revenue) + 'px' }">
              <div
class="w-full bg-gradient-to-t from-green-500 to-blue-500 rounded-t transition-all duration-300 absolute bottom-0"
                   :style="{ height: getChartHeight(day.revenue) + 'px' }"></div>
            </div>
            <span class="text-xs text-gray-500 mt-1 text-center">{{ day.date }}</span>
            <span class="text-xs text-gray-400 mt-1 text-center">{{ formatCurrency(day.revenue) }}</span>
          </div>
        </div>
      </div>

      <!-- نمودار توزیع درآمد -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع درآمد بر اساس منبع</h3>
        <div class="space-y-4">
          <div v-for="source in revenueSources" :key="source.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: source.color }"></div>
              <span class="text-sm text-gray-700">{{ source.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ formatCurrency(source.amount) }}</span>
              <span class="text-sm text-gray-500">({{ source.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش‌های تفصیلی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">گزارش‌های تفصیلی</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
            خروجی PDF
          </button>
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            خروجی اکسل
          </button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- گزارش تراکنش‌ها -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-3">گزارش تراکنش‌ها</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کل تراکنش‌ها:</span>
              <span class="font-medium">{{ detailedReports.totalTransactions }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">تراکنش‌های موفق:</span>
              <span class="font-medium text-green-600">{{ detailedReports.successfulTransactions }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">تراکنش‌های ناموفق:</span>
              <span class="font-medium text-red-600">{{ detailedReports.failedTransactions }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">نرخ موفقیت:</span>
              <span class="font-medium">{{ detailedReports.successRate }}%</span>
            </div>
          </div>
        </div>

        <!-- گزارش کاربران -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-3">گزارش کاربران</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کل کاربران:</span>
              <span class="font-medium">{{ detailedReports.totalUsers }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کاربران فعال:</span>
              <span class="font-medium text-green-600">{{ detailedReports.activeUsers }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کاربران جدید:</span>
              <span class="font-medium text-blue-600">{{ detailedReports.newUsers }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">میانگین تراکنش:</span>
              <span class="font-medium">{{ detailedReports.averageTransactions }}</span>
            </div>
          </div>
        </div>

        <!-- گزارش کارمزدها -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-3">گزارش کارمزدها</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کل کارمزدها:</span>
              <span class="font-medium">{{ formatCurrency(detailedReports.totalFees) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کارمزد شارژ:</span>
              <span class="font-medium text-green-600">{{ formatCurrency(detailedReports.chargeFees) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کارمزد برداشت:</span>
              <span class="font-medium text-blue-600">{{ formatCurrency(detailedReports.withdrawalFees) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">میانگین کارمزد:</span>
              <span class="font-medium">{{ detailedReports.averageFee }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌بینی‌های مالی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">پیش‌بینی‌های مالی</h3>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center p-6 bg-blue-50 rounded-lg">
          <div class="text-2xl font-bold text-blue-600">{{ formatCurrency(forecasts.nextMonthRevenue) }}</div>
          <div class="text-sm text-gray-600">درآمد ماه آینده</div>
          <div class="text-xs text-blue-500 mt-1">+{{ forecasts.revenueGrowth }}%</div>
        </div>
        
        <div class="text-center p-6 bg-green-50 rounded-lg">
          <div class="text-2xl font-bold text-green-600">{{ formatCurrency(forecasts.nextMonthProfit) }}</div>
          <div class="text-sm text-gray-600">سود ماه آینده</div>
          <div class="text-xs text-green-500 mt-1">+{{ forecasts.profitGrowth }}%</div>
        </div>
        
        <div class="text-center p-6 bg-purple-50 rounded-lg">
          <div class="text-2xl font-bold text-purple-600">{{ forecasts.nextMonthUsers }}</div>
          <div class="text-sm text-gray-600">کاربران جدید</div>
          <div class="text-xs text-purple-500 mt-1">+{{ forecasts.userGrowth }}%</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watchEffect } from 'vue';
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>

interface SummaryData {
  totalBalance?: number
  todayTxCount?: number
  [key: string]: unknown
}

const summary = ref<SummaryData | null>(null)
const days = ref(30)
const { data: summaryData } = await useFetch('/api/admin/wallet/summary', {
  credentials: 'include',
  key: 'admin-wallet-summary',
  defaultCache: true,
  swr: true,
})
const { data: trendData } = await useFetch('/api/admin/wallet/trend', {
  query: { days },
  credentials: 'include',
  key: () => `admin-wallet-trend-${days.value}`,
  defaultCache: true,
  swr: true,
})

const financialStats = computed(() => ({
  totalRevenue: Number(summary.value?.totalBalance || 0),
  revenueGrowth: 0,
  netProfit: Number(summary.value?.totalBalance || 0),
  profitMargin: 0,
  totalFees: 0,
  feePercentage: 0,
  totalExpenses: 0,
  expensePercentage: 0,
}))

interface TrendItem {
  Day?: string
  Net?: number | string
  [key: string]: unknown
}
const dailyRevenue = computed(() => {
  const trend = trendData.value as { items?: TrendItem[] } | undefined
  const items: TrendItem[] = trend?.items || []
  return items.slice(-7).map((it: TrendItem) => ({ date: it.Day || '', revenue: Number(it.Net || 0), transactions: undefined }))
})

const monthlyRevenue = computed(() => {
  const trend = trendData.value as { items?: TrendItem[] } | undefined
  const items: TrendItem[] = trend?.items || []
  return [
    { month: '۳ ماه اخیر', revenue: sumRange(items, 0, 30), growth: 0 },
    { month: '۶ ماه اخیر', revenue: sumRange(items, 0, 30), growth: 0 },
  ]
})

const revenueTrend = computed(() => {
  const trend = trendData.value as { items?: TrendItem[] } | undefined
  const items: TrendItem[] = trend?.items || []
  return items.map((it: TrendItem) => ({ date: it.Day?.slice(5) || '', revenue: Number(it.Net || 0) }))
})

const revenueSources = ref([
  { name: 'کارمزد شارژ', amount: 0, percentage: 0, color: '#10B981' },
  { name: 'کارمزد برداشت', amount: 0, percentage: 0, color: '#3B82F6' },
  { name: 'کارمزد انتقال', amount: 0, percentage: 0, color: '#F59E0B' },
  { name: 'سایر', amount: 0, percentage: 0, color: '#8B5CF6' }
])

const detailedReports = computed(() => ({
  totalTransactions: 0,
  successfulTransactions: Number(summary.value?.todayTxCount || 0),
  failedTransactions: 0,
  successRate: 0,
  totalUsers: 0,
  activeUsers: 0,
  newUsers: 0,
  averageTransactions: 0,
  totalFees: 0,
  chargeFees: 0,
  withdrawalFees: 0,
  averageFee: 0,
}))

const forecasts = ref({
  nextMonthRevenue: 0,
  revenueGrowth: 0,
  nextMonthProfit: 0,
  profitGrowth: 0,
  nextMonthUsers: 0,
  userGrowth: 0,
})

watchEffect(()=>{
  if (summaryData.value) summary.value = summaryData.value
})

function sumRange(items: TrendItem[], start:number, count:number){
  return items.slice(start, start+count).reduce((acc, it: TrendItem)=>acc + Number(it.Net||0), 0)
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

interface RevenueItem {
  date: string
  revenue: number
}

const getChartHeight = (revenue: number) => {
  const arr = revenueTrend.value
  const maxRevenue = Math.max(...arr.map((item: RevenueItem) => item.revenue), 0)
  const minRevenue = Math.min(...arr.map((item: RevenueItem) => item.revenue), 0)
  const range = maxRevenue - minRevenue
  const height = 200
  if (range === 0) return height
  const percentage = ((revenue - minRevenue) / range) * 100
  return (percentage / 100) * height
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. آمار کلی مالی (درآمد، سود، کارمزد، هزینه)
  2. گزارش‌های دوره‌ای (روزانه و ماهانه)
  3. نمودارهای تحلیلی (روند درآمد و توزیع منابع)
  4. گزارش‌های تفصیلی (تراکنش‌ها، کاربران، کارمزدها)
  5. پیش‌بینی‌های مالی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
