<template>
  <div class="space-y-6 w-full">
    <!-- هدر بخش -->
    <div class="p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">💰 مدیریت موجودی و تراز</h2>
      <p class="text-gray-600">نمایش و مدیریت موجودی‌های کیف پول کاربران</p>
    </div>

    <!-- کارت‌های آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- موجودی کل -->
      <div class="border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">موجودی کل</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ formatCurrency(balanceStats.totalBalance) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ balanceStats.balanceGrowth }}%</span>
          <span class="mx-2">از ماه قبل</span>
        </div>
      </div>

      <!-- موجودی قابل برداشت -->
      <div class="border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">قابل برداشت</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">آماده</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ formatCurrency(balanceStats.withdrawableBalance) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-blue-500">{{ balanceStats.withdrawablePercentage }}%</span>
          <span class="mx-2">از کل موجودی</span>
        </div>
      </div>

      <!-- موجودی مسدود شده -->
      <div class="border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">مسدود شده</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">مسدود</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ formatCurrency(balanceStats.blockedBalance) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ balanceStats.blockedCount }}</span>
          <span class="mx-2">حساب مسدود</span>
        </div>
      </div>

      <!-- موجودی در انتظار -->
      <div class="border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">در انتظار تأیید</h3>
          <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">انتظار</span>
        </div>
        <div class="text-3xl font-bold text-yellow-600 mb-2">{{ formatCurrency(balanceStats.pendingBalance) }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-yellow-500">{{ balanceStats.pendingCount }}</span>
          <span class="mx-2">تراکنش در انتظار</span>
        </div>
      </div>
    </div>

    <!-- نمودار روند موجودی -->
    <div class="border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">روند موجودی (30 روز گذشته)</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-sm bg-blue-100 text-blue-800 rounded-lg hover:bg-blue-200">روزانه</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">هفتگی</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">ماهانه</button>
        </div>
      </div>
      
      <!-- نمودار -->
      <div class="h-64 w-full overflow-x-auto">
        <div class="flex items-end space-x-1 h-full min-w-max px-4">
          <div v-for="(day, index) in balanceTrend" :key="index" class="flex flex-col items-center flex-shrink-0 min-w-[6px]">
            <div
class="w-full bg-gray-200 rounded-t relative"
                 :style="{ height: getChartHeight(day.balance) + 'px' }">
              <div
class="w-full bg-gradient-to-t from-blue-500 to-indigo-500 rounded-t transition-all duration-300 absolute bottom-0"
                   :style="{ height: getChartHeight(day.balance) + 'px' }"></div>
            </div>
            <span class="text-xs text-gray-500 mt-1 text-center hidden md:block">{{ day.date }}</span>
            <span class="text-xs text-gray-400 mt-1 text-center hidden xl:block">{{ formatCurrency(day.balance) }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول تغییرات موجودی -->
    <div class="border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">تاریخچه تغییرات موجودی</h3>
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          مشاهده همه
        </button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right min-w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium whitespace-nowrap">تاریخ</th>
              <th class="px-4 py-3 text-gray-700 font-medium whitespace-nowrap">نوع تغییر</th>
              <th class="px-4 py-3 text-gray-700 font-medium whitespace-nowrap">مبلغ</th>
              <th class="px-4 py-3 text-gray-700 font-medium whitespace-nowrap">موجودی قبل</th>
              <th class="px-4 py-3 text-gray-700 font-medium whitespace-nowrap">موجودی بعد</th>
              <th class="px-4 py-3 text-gray-700 font-medium whitespace-nowrap">توضیحات</th>
              <th class="px-4 py-3 text-gray-700 font-medium whitespace-nowrap">وضعیت</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="change in balanceChanges" :key="change.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-gray-900 whitespace-nowrap">{{ change.date }}</td>
              <td class="px-4 py-3">
                <span :class="getChangeTypeClass(change.type)" class="px-2 py-1 text-xs rounded-full whitespace-nowrap">
                  {{ change.type }}
                </span>
              </td>
              <td class="px-4 py-3 font-medium whitespace-nowrap" :class="change.amount > 0 ? 'text-green-600' : 'text-red-600'">
                {{ change.amount > 0 ? '+' : '' }}{{ formatCurrency(change.amount) }}
              </td>
              <td class="px-4 py-3 text-gray-600 whitespace-nowrap">{{ formatCurrency(change.balanceBefore) }}</td>
              <td class="px-4 py-3 text-gray-900 font-medium whitespace-nowrap">{{ formatCurrency(change.balanceAfter) }}</td>
              <td class="px-4 py-3 text-gray-600 max-w-xs truncate">{{ change.description }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(change.status)" class="px-2 py-1 text-xs rounded-full whitespace-nowrap">
                  {{ change.status }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- ابزارهای مدیریت -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- شارژ دستی -->
      <div class="border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">شارژ دستی کیف پول</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">انتخاب کاربر</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>انتخاب کنید...</option>
              <option>کاربر 1</option>
              <option>کاربر 2</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ شارژ</label>
            <input
type="number" placeholder="مبلغ را وارد کنید" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea
placeholder="توضیحات شارژ" rows="3"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
          </div>
          <button class="w-full px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
            شارژ کیف پول
          </button>
        </div>
      </div>

      <!-- کسر دستی -->
      <div class="border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">کسر دستی از کیف پول</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">انتخاب کاربر</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>انتخاب کنید...</option>
              <option>کاربر 1</option>
              <option>کاربر 2</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ کسر</label>
            <input
type="number" placeholder="مبلغ را وارد کنید" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دلیل کسر</label>
            <textarea
placeholder="دلیل کسر از کیف پول" rows="3"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
          </div>
          <button class="w-full px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
            کسر از کیف پول
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, watchEffect } from 'vue';
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
// آمار از API
const balanceStats = reactive({
  totalBalance: 0,
  balanceGrowth: 0,
  withdrawableBalance: 0,
  withdrawablePercentage: 0,
  blockedBalance: 0,
  blockedCount: 0,
  pendingBalance: 0,
  pendingCount: 0
})

// روند موجودی (placeholder)
const balanceTrend = ref<Array<{ date: string; balance: number }>>([])

// تغییرات موجودی (placeholder)
interface BalanceChange {
  id?: number | string
  [key: string]: unknown
}
const balanceChanges = ref<BalanceChange[]>([])

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع محاسبه ارتفاع نمودار
const getChartHeight = (balance: number) => {
  const maxBalance = Math.max(...balanceTrend.value.map(item => item.balance))
  const minBalance = Math.min(...balanceTrend.value.map(item => item.balance))
  const range = maxBalance - minBalance
  const height = 200

  if (range === 0) return height

  const percentage = ((balance - minBalance) / range) * 100
  return (percentage / 100) * height
}

// تابع کلاس نوع تغییر
const getChangeTypeClass = (type: string) => {
  switch (type) {
    case 'شارژ':
      return 'bg-green-100 text-green-800'
    case 'برداشت':
      return 'bg-red-100 text-red-800'
    case 'کسر':
      return 'bg-orange-100 text-orange-800'
    case 'انتقال':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'موفق':
      return 'bg-green-100 text-green-800'
    case 'در انتظار':
      return 'bg-yellow-100 text-yellow-800'
    case 'ناموفق':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// واکشی summary و trend از API ادمین
const { data: summary } = await useFetch('/api/admin/wallet/summary', {
  method: 'GET',
  credentials: 'include',
  key: 'admin-wallet-summary',
  defaultCache: true,
})
const { data: trend } = await useFetch('/api/admin/wallet/trend', {
  method: 'GET',
  query: { days: 30 },
  credentials: 'include',
  key: 'admin-wallet-trend-30',
  defaultCache: true,
})

interface Summary {
  totalBalance?: number | string
  blockedCount?: number | string
  pendingCount?: number | string
  [key: string]: unknown
}
interface Trend {
  items?: Array<{
    [key: string]: unknown
  }>
  [key: string]: unknown
}
watchEffect(() => {
  const s = summary.value as Summary | null
  if (s) {
    balanceStats.totalBalance = Number(s.totalBalance || 0)
    balanceStats.blockedCount = Number(s.blockedCount || 0)
    balanceStats.pendingCount = Number(s.pendingCount || 0)
  }
  const t = trend.value as Trend | null
  if (t && Array.isArray(t.items)) {
    interface TrendItem {
      day?: string
      net?: number | string
      [key: string]: unknown
    }
    balanceTrend.value = (t.items as TrendItem[]).map((x) => ({ date: x.day || '', balance: Number(x.net || 0) }))
  }
})
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. نمایش آمار کلی موجودی (کل، قابل برداشت، مسدود، در انتظار)
  2. نمودار روند موجودی در 30 روز گذشته
  3. جدول تاریخچه تغییرات موجودی
  4. ابزارهای شارژ و کسر دستی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
