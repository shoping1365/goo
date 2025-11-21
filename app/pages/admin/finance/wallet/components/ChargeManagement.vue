<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-indigo-50 to-blue-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">💳 مدیریت شارژ</h2>
      <p class="text-gray-600">مدیریت روش‌های شارژ و درگاه‌های پرداخت</p>
    </div>

    <!-- آمار شارژ -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کل شارژها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کل شارژها</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">امروز</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ chargeStats.totalCharges }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ chargeStats.dailyGrowth }}%</span>
          <span class="mx-2">از دیروز</span>
        </div>
      </div>

      <!-- مبلغ کل شارژ -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">مبلغ کل شارژ</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">مثبت</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ formatCurrency(chargeStats.totalAmount) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ chargeStats.amountGrowth }}%</span>
          <span class="mx-2">از ماه قبل</span>
        </div>
      </div>

      <!-- شارژهای موفق -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">شارژهای موفق</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">عالی</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ chargeStats.successfulCharges }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ chargeStats.successRate }}%</span>
          <span class="mx-2">نرخ موفقیت</span>
        </div>
      </div>

      <!-- شارژهای ناموفق -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">شارژهای ناموفق</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">هشدار</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ chargeStats.failedCharges }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ chargeStats.failureRate }}%</span>
          <span class="mx-2">نرخ شکست</span>
        </div>
      </div>
    </div>

    <!-- درگاه‌های پرداخت -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">درگاه‌های پرداخت</h3>
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          افزودن درگاه جدید
        </button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
v-for="gateway in paymentGateways" :key="gateway.id" 
             class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center mr-3">
                <span class="text-lg">{{ gateway.icon }}</span>
              </div>
              <div>
                <h4 class="font-medium text-gray-900">{{ gateway.name }}</h4>
                <p class="text-sm text-gray-500">{{ gateway.description }}</p>
              </div>
            </div>
            <span :class="getGatewayStatusClass(gateway.status)" class="px-2 py-1 text-xs rounded-full">
              {{ gateway.status }}
            </span>
          </div>
          
          <div class="space-y-2 mb-4">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">کارمزد:</span>
              <span class="font-medium">{{ gateway.fee }}%</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">حداقل مبلغ:</span>
              <span class="font-medium">{{ formatCurrency(gateway.minAmount) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">حداکثر مبلغ:</span>
              <span class="font-medium">{{ formatCurrency(gateway.maxAmount) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">تراکنش‌های امروز:</span>
              <span class="font-medium">{{ gateway.todayTransactions }}</span>
            </div>
          </div>
          
          <div class="flex space-x-2 space-x-reverse">
            <button class="flex-1 px-3 py-2 bg-blue-100 text-blue-700 text-sm rounded-lg hover:bg-blue-200">
              تنظیمات
            </button>
            <button
:class="gateway.status === 'فعال' ? 'bg-red-100 text-red-700 hover:bg-red-200' : 'bg-green-100 text-green-700 hover:bg-green-200'"
                    class="flex-1 px-3 py-2 text-sm rounded-lg">
              {{ gateway.status === 'فعال' ? 'غیرفعال' : 'فعال' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- روش‌های شارژ -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- شارژ از طریق کارت بانکی -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">شارژ از طریق کارت بانکی</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">کارمزد تراکنش:</span>
            <span class="text-sm font-medium text-gray-900">{{ chargeSettings.cardFee }}%</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">حداقل مبلغ:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(chargeSettings.cardMinAmount) }}</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">حداکثر مبلغ:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(chargeSettings.cardMaxAmount) }}</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">زمان پردازش:</span>
            <span class="text-sm font-medium text-gray-900">{{ chargeSettings.cardProcessingTime }} دقیقه</span>
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            تنظیم کارمزد
          </button>
        </div>
      </div>

      <!-- شارژ از طریق کیف پول دیجیتال -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">شارژ از طریق کیف پول دیجیتال</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">کارمزد تراکنش:</span>
            <span class="text-sm font-medium text-gray-900">{{ chargeSettings.walletFee }}%</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">حداقل مبلغ:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(chargeSettings.walletMinAmount) }}</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">حداکثر مبلغ:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(chargeSettings.walletMaxAmount) }}</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">زمان پردازش:</span>
            <span class="text-sm font-medium text-gray-900">{{ chargeSettings.walletProcessingTime }} دقیقه</span>
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            تنظیم کارمزد
          </button>
        </div>
      </div>
    </div>

    <!-- نمودار شارژها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار شارژ بر اساس درگاه -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">شارژ بر اساس درگاه پرداخت</h3>
        <div class="space-y-4">
          <div v-for="gateway in gatewayStats" :key="gateway.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: gateway.color }"></div>
              <span class="text-sm text-gray-700">{{ gateway.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ gateway.count }}</span>
              <span class="text-sm text-gray-500">({{ gateway.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار روند شارژ -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">روند شارژ (7 روز گذشته)</h3>
        <div class="flex items-end space-x-2 space-x-reverse h-48 overflow-x-auto">
          <div v-for="(day, index) in chargeTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-16">
            <div
class="w-full bg-gray-200 rounded-t relative"
                 :style="{ height: getChartHeight(day.amount) + 'px' }">
              <div
class="w-full bg-gradient-to-t from-blue-500 to-indigo-500 rounded-t transition-all duration-300 absolute bottom-0"
                   :style="{ height: getChartHeight(day.amount) + 'px' }"></div>
            </div>
            <span class="text-xs text-gray-500 mt-1 text-center">{{ day.date }}</span>
            <span class="text-xs text-gray-400 mt-1 text-center">{{ formatCurrency(day.amount) }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول شارژهای اخیر -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">شارژهای اخیر</h3>
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          مشاهده همه
        </button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">شماره تراکنش</th>
              <th class="px-4 py-3 text-gray-700 font-medium">کاربر</th>
              <th class="px-4 py-3 text-gray-700 font-medium">مبلغ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">درگاه پرداخت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="charge in recentCharges" :key="charge.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-gray-900 font-medium">{{ charge.transactionId }}</td>
              <td class="px-4 py-3">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center text-xs font-medium text-gray-600">
                    {{ charge.userInitials }}
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ charge.userName }}</div>
                    <div class="text-xs text-gray-500">{{ charge.userEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 font-medium text-green-600">{{ formatCurrency(charge.amount) }}</td>
              <td class="px-4 py-3 text-gray-600">{{ charge.gateway }}</td>
              <td class="px-4 py-3 text-gray-600">{{ charge.date }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(charge.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ charge.status }}
                </span>
              </td>
              <td class="px-4 py-3">
                <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, watchEffect } from 'vue';
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
// آمار شارژ (از API محاسبه می‌شود)
const chargeStats = reactive({
  totalCharges: 0,
  dailyGrowth: 0,
  totalAmount: 0,
  amountGrowth: 0,
  successfulCharges: 0,
  successRate: 0,
  failedCharges: 0,
  failureRate: 0
})

// تنظیمات شارژ
const chargeSettings = {
  cardFee: 1.5,
  cardMinAmount: 10000,
  cardMaxAmount: 50000000,
  cardProcessingTime: 5,
  walletFee: 0.5,
  walletMinAmount: 5000,
  walletMaxAmount: 10000000,
  walletProcessingTime: 2
}

// درگاه‌های پرداخت (placeholder)
const paymentGateways = [
  {
    id: 1,
    name: 'ملی',
    description: 'درگاه پرداخت ملی',
    icon: '🏦',
    status: 'فعال',
    fee: 1.5,
    minAmount: 10000,
    maxAmount: 50000000,
    todayTransactions: 1250
  },
  {
    id: 2,
    name: 'ملت',
    description: 'درگاه پرداخت ملت',
    icon: '🏛️',
    status: 'فعال',
    fee: 1.8,
    minAmount: 5000,
    maxAmount: 30000000,
    todayTransactions: 890
  },
  {
    id: 3,
    name: 'پارسیان',
    description: 'درگاه پرداخت پارسیان',
    icon: '🏢',
    status: 'غیرفعال',
    fee: 2.0,
    minAmount: 10000,
    maxAmount: 40000000,
    todayTransactions: 0
  },
  {
    id: 4,
    name: 'سپه',
    description: 'درگاه پرداخت سپه',
    icon: '🏦',
    status: 'فعال',
    fee: 1.2,
    minAmount: 5000,
    maxAmount: 25000000,
    todayTransactions: 650
  },
  {
    id: 5,
    name: 'تجارت',
    description: 'درگاه پرداخت تجارت',
    icon: '🏛️',
    status: 'فعال',
    fee: 1.6,
    minAmount: 10000,
    maxAmount: 35000000,
    todayTransactions: 720
  },
  {
    id: 6,
    name: 'کیف پول دیجیتال',
    description: 'شارژ از کیف پول',
    icon: '💳',
    status: 'فعال',
    fee: 0.5,
    minAmount: 1000,
    maxAmount: 10000000,
    todayTransactions: 1850
  }
]

// آمار درگاه‌ها
const gatewayStats = [
  { name: 'ملی', count: 1250, percentage: 28, color: '#10B981' },
  { name: 'ملت', count: 890, percentage: 20, color: '#3B82F6' },
  { name: 'پارسیان', count: 0, percentage: 0, color: '#6B7280' },
  { name: 'سپه', count: 650, percentage: 15, color: '#F59E0B' },
  { name: 'تجارت', count: 720, percentage: 16, color: '#EF4444' },
  { name: 'کیف پول دیجیتال', count: 1850, percentage: 21, color: '#8B5CF6' }
]

// روند شارژ (از API trend کیف پول)
const chargeTrend = ref<Array<{ date: string; amount: number }>>([])

// شارژهای اخیر (از API تراکنش‌ها)
interface ChargeDisplayItem {
  id?: number | string
  transactionId?: number | string
  userName: string | number
  userEmail: string
  userInitials: string
  amount: number
  gateway: string
  date?: string
  status: string
}

const recentCharges = ref<ChargeDisplayItem[]>([])

const page = ref(1)
const pageSize = ref(20)
const { data: txs } = await useFetch('/api/admin/wallet/transactions', {
  method: 'GET',
  query: { page, pageSize, type: 'credit', method: 'online' },
  credentials: 'include',
  key: () => `admin-wallet-credit-online-${page.value}-${pageSize.value}`,
  defaultCache: true,
})
const { data: trend } = await useFetch('/api/admin/wallet/trend', {
  method: 'GET', query: { days: 7 }, credentials: 'include', key: 'admin-wallet-trend-7', defaultCache: true,
})

interface ChargeItem {
  id?: number | string
  transactionId?: number | string
  [key: string]: unknown
}
interface ChargesResponse {
  items?: ChargeItem[]
  [key: string]: unknown
}
watchEffect(() => {
  const res = txs.value as ChargesResponse | null
  if (res && Array.isArray(res.items)) {
    recentCharges.value = res.items.map((r: ChargeItem) => ({
      id: r.id,
      transactionId: r.id,
      userName: r.username || r.user_id,
      userEmail: r.email || '-',
      userInitials: (r.username || '-').slice(0,2),
      amount: Number(r.amount || 0),
      gateway: r.gateway || r.method || '-',
      date: r.created_at,
      status: r.status === 'success' ? 'موفق' : (r.status === 'pending' ? 'در انتظار' : 'ناموفق')
    }))
    // آمار ساده
    chargeStats.totalCharges = Number(res.total || 0)
    const succ = recentCharges.value.filter((x: ChargeDisplayItem)=>x.status==='موفق').length
    const fail = recentCharges.value.filter((x: ChargeDisplayItem)=>x.status==='ناموفق').length
    const totalShown = Math.max(1, recentCharges.value.length)
    chargeStats.successfulCharges = succ
    chargeStats.failedCharges = fail
    chargeStats.successRate = Math.round((succ/totalShown)*1000)/10
    chargeStats.failureRate = Math.round((fail/totalShown)*1000)/10
  }
  interface TrendItem {
    day?: string
    net?: number
    [key: string]: unknown
  }
  interface TrendResponse {
    items?: TrendItem[]
    [key: string]: unknown
  }
  const t = trend.value as TrendResponse | null
  if (t && Array.isArray(t.items)) {
    chargeTrend.value = t.items.map((x: TrendItem)=>({ date: x.day || '', amount: Number(x.net||0) }))
  }
})

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع محاسبه ارتفاع نمودار
const getChartHeight = (amount: number) => {
  const maxAmount = Math.max(...chargeTrend.value.map(item => item.amount))
  const minAmount = Math.min(...chargeTrend.value.map(item => item.amount))
  const range = maxAmount - minAmount
  const height = 150

  if (range === 0) return height

  const percentage = ((amount - minAmount) / range) * 100
  return (percentage / 100) * height
}

// تابع کلاس وضعیت درگاه
const getGatewayStatusClass = (status: string) => {
  switch (status) {
    case 'فعال':
      return 'bg-green-100 text-green-800'
    case 'غیرفعال':
      return 'bg-red-100 text-red-800'
    case 'در حال تعمیر':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'موفق':
      return 'bg-green-100 text-green-800'
    case 'ناموفق':
      return 'bg-red-100 text-red-800'
    case 'در انتظار':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. آمار کلی شارژها (کل، مبلغ، موفق، ناموفق)
  2. مدیریت درگاه‌های پرداخت
  3. تنظیمات روش‌های شارژ
  4. نمودارهای تحلیلی
  5. جدول شارژهای اخیر
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
