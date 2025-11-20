<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">گزارش تراکنش‌های همگام‌سازی شده</h4>
        <p class="text-sm text-gray-600 mt-1">گزارش جامع تراکنش‌های همگام‌سازی شده با نرم‌افزارهای حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="showFilters = !showFilters"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z" />
          </svg>
          فیلترها
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="exportTransactionsReport"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          صادر کردن گزارش
        </button>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div v-if="showFilters" class="bg-gray-50 rounded-lg p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلترهای پیشرفته</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع تراکنش</label>
          <select
            v-model="filters.transactionType"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          >
            <option value="">همه انواع</option>
            <option value="sale">فروش</option>
            <option value="purchase">خرید</option>
            <option value="payment">پرداخت</option>
            <option value="refund">بازپرداخت</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نرم‌افزار</label>
          <select
            v-model="filters.software"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          >
            <option value="">همه نرم‌افزارها</option>
            <option value="helo">هلو</option>
            <option value="parsian">پارسیان</option>
            <option value="sepidar">سپیدار</option>
            <option value="ghias">قیاس</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ از</label>
          <input
            v-model="filters.dateFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ تا</label>
          <input
            v-model="filters.dateTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          />
        </div>
      </div>
    </div>

    <!-- آمار تراکنش‌های همگام شده -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ stats.totalSynced }}</div>
            <div class="text-sm text-blue-700">کل همگام شده</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ stats.todaySynced }}</div>
            <div class="text-sm text-green-700">امروز همگام شده</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ stats.avgSyncTime }}</div>
            <div class="text-sm text-yellow-700">میانگین زمان همگام</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ stats.syncRate }}%</div>
            <div class="text-sm text-purple-700">نرخ همگام‌سازی</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول تراکنش‌های همگام شده -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تراکنش‌های همگام‌سازی شده</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">شماره تراکنش</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">مبلغ</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">زمان همگام</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="transaction in filteredTransactions" :key="transaction.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div>
                  <div class="font-medium text-gray-900">{{ transaction.transactionNumber }}</div>
                  <div class="text-xs text-gray-500">{{ transaction.reference }}</div>
                </div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getTransactionTypeClass(transaction.type)"
                >
                  {{ getTransactionTypeLabel(transaction.type) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="text-gray-900 font-medium">{{ formatCurrency(transaction.amount) }}</div>
                <div class="text-xs text-gray-500">{{ transaction.currency }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <img :src="transaction.softwareLogo" :alt="transaction.softwareName" class="w-6 h-6 rounded" />
                  <span class="text-sm text-gray-900">{{ transaction.softwareName }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ transaction.syncTime }}</div>
                <div class="text-xs text-gray-500">{{ transaction.syncDuration }}ms</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <div
                    class="w-2 h-2 rounded-full"
                    :class="getSyncStatusColor(transaction.syncStatus)"
                  ></div>
                  <span
                    class="px-2 py-1 rounded-full text-xs font-medium"
                    :class="getSyncStatusClass(transaction.syncStatus)"
                  >
                    {{ getSyncStatusLabel(transaction.syncStatus) }}
                  </span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="مشاهده جزئیات"
                    @click="viewTransaction(transaction)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-green-600 hover:text-green-800"
                    title="همگام‌سازی مجدد"
                    @click="resyncTransaction(transaction)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودارهای تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">توزیع تراکنش‌ها بر اساس نوع</h5>
        <div class="space-y-4">
          <div v-for="type in transactionTypeDistribution" :key="type.type" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :class="getTransactionTypeColor(type.type)"></div>
              <span class="text-sm text-gray-700">{{ getTransactionTypeLabel(type.type) }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="getTransactionTypeColor(type.type)"
                  :style="{ width: type.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ type.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">نرخ همگام‌سازی بر اساس نرم‌افزار</h5>
        <div class="space-y-4">
          <div v-for="software in softwareSyncRates" :key="software.name" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <img :src="software.logo" :alt="software.name" class="w-6 h-6 rounded" />
              <span class="text-sm text-gray-700">{{ software.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="getSyncRateColor(software.rate)"
                  :style="{ width: software.rate + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ software.rate }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- صفحه‌بندی -->
    <div class="flex items-center justify-between">
      <div class="text-sm text-gray-700">
        نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} تراکنش
      </div>
      <div class="flex gap-2">
        <button
          :disabled="pagination.currentPage === 1"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
          @click="previousPage"
        >
          قبلی
        </button>
        <span class="px-3 py-1 text-sm text-gray-700">
          صفحه {{ pagination.currentPage }} از {{ pagination.totalPages }}
        </span>
        <button
          :disabled="pagination.currentPage === pagination.totalPages"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
          @click="nextPage"
        >
          بعدی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// وضعیت فیلترها
const showFilters = ref(false)

// فیلترها
const filters = ref({
  transactionType: '',
  software: '',
  dateFrom: '',
  dateTo: ''
})

// آمار
const stats = ref({
  totalSynced: 15420,
  todaySynced: 1250,
  avgSyncTime: 245,
  syncRate: 96.3
})

// صفحه‌بندی
const pagination = ref({
  currentPage: 1,
  totalPages: 50,
  total: 15420,
  start: 1,
  end: 20,
  perPage: 20
})

// تراکنش‌ها
const transactions = ref([
  {
    id: 1,
    transactionNumber: 'TXN-001',
    reference: 'فاکتور #12345',
    type: 'sale',
    amount: 1250000,
    currency: 'تومان',
    softwareName: 'هلو',
    softwareLogo: '/images/helo-logo.png',
    syncTime: 'امروز ۱۴:۳۰',
    syncDuration: 245,
    syncStatus: 'success'
  },
  {
    id: 2,
    transactionNumber: 'TXN-002',
    reference: 'پرداخت #67890',
    type: 'payment',
    amount: 850000,
    currency: 'تومان',
    softwareName: 'پارسیان',
    softwareLogo: '/images/parsian-logo.png',
    syncTime: 'امروز ۱۴:۲۵',
    syncDuration: 189,
    syncStatus: 'success'
  },
  {
    id: 3,
    transactionNumber: 'TXN-003',
    reference: 'خرید #11111',
    type: 'purchase',
    amount: 2100000,
    currency: 'تومان',
    softwareName: 'سپیدار',
    softwareLogo: '/images/sepidar-logo.png',
    syncTime: 'امروز ۱۳:۴۵',
    syncDuration: 1250,
    syncStatus: 'warning'
  }
])

// تراکنش‌های فیلتر شده
const filteredTransactions = computed(() => {
  return transactions.value.filter(transaction => {
    if (filters.value.transactionType && transaction.type !== filters.value.transactionType) return false
    if (filters.value.software && transaction.softwareName.toLowerCase() !== filters.value.software) return false
    // TODO: فیلتر تاریخ
    return true
  })
})

// توزیع نوع تراکنش
const transactionTypeDistribution = ref([
  { type: 'sale', percentage: 45 },
  { type: 'payment', percentage: 30 },
  { type: 'purchase', percentage: 20 },
  { type: 'refund', percentage: 5 }
])

// نرخ همگام‌سازی نرم‌افزارها
const softwareSyncRates = ref([
  { name: 'هلو', logo: '/images/helo-logo.png', rate: 98.5 },
  { name: 'پارسیان', logo: '/images/parsian-logo.png', rate: 99.2 },
  { name: 'سپیدار', logo: '/images/sepidar-logo.png', rate: 85.3 },
  { name: 'قیاس', logo: '/images/ghias-logo.png', rate: 92.1 }
])

// فرمت ارز
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// کلاس نوع تراکنش
const getTransactionTypeClass = (type: string) => {
  const classes = {
    sale: 'bg-green-100 text-green-700',
    purchase: 'bg-red-100 text-red-700',
    payment: 'bg-blue-100 text-blue-700',
    refund: 'bg-yellow-100 text-yellow-700'
  }
  return classes[type] || 'bg-gray-100 text-gray-700'
}

// رنگ نوع تراکنش
const getTransactionTypeColor = (type: string) => {
  const colors = {
    sale: 'bg-green-500',
    purchase: 'bg-red-500',
    payment: 'bg-blue-500',
    refund: 'bg-yellow-500'
  }
  return colors[type] || 'bg-gray-500'
}

// برچسب نوع تراکنش
const getTransactionTypeLabel = (type: string) => {
  const labels = {
    sale: 'فروش',
    purchase: 'خرید',
    payment: 'پرداخت',
    refund: 'بازپرداخت'
  }
  return labels[type] || type
}

// رنگ وضعیت همگام‌سازی
const getSyncStatusColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// کلاس وضعیت همگام‌سازی
const getSyncStatusClass = (status: string) => {
  const classes = {
    success: 'bg-green-100 text-green-700',
    warning: 'bg-yellow-100 text-yellow-700',
    error: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت همگام‌سازی
const getSyncStatusLabel = (status: string) => {
  const labels = {
    success: 'موفق',
    warning: 'هشدار',
    error: 'خطا'
  }
  return labels[status] || status
}

// رنگ نرخ همگام‌سازی
const getSyncRateColor = (rate: number) => {
  if (rate >= 95) return 'bg-green-500'
  if (rate >= 80) return 'bg-yellow-500'
  return 'bg-red-500'
}

// اعمال فیلترها
const applyFilters = () => {
  // TODO: اعمال فیلترها
  console.log('فیلترهای تراکنش اعمال شد:', filters.value)
}

// صادر کردن گزارش تراکنش‌ها
const exportTransactionsReport = () => {
  // TODO: صادر کردن گزارش
  console.log('صادر کردن گزارش تراکنش‌های همگام شده')
}

// مشاهده تراکنش
const viewTransaction = (transaction: any) => {
  // TODO: مشاهده جزئیات تراکنش
  console.log('مشاهده تراکنش:', transaction)
}

// همگام‌سازی مجدد تراکنش
const resyncTransaction = (transaction: any) => {
  // TODO: همگام‌سازی مجدد
  console.log('همگام‌سازی مجدد تراکنش:', transaction)
}

// صفحه قبلی
const previousPage = () => {
  if (pagination.value.currentPage > 1) {
    pagination.value.currentPage--
    // TODO: بارگذاری داده‌های صفحه جدید
  }
}

// صفحه بعدی
const nextPage = () => {
  if (pagination.value.currentPage < pagination.value.totalPages) {
    pagination.value.currentPage++
    // TODO: بارگذاری داده‌های صفحه جدید
  }
}
</script>

<!--
  کامپوننت گزارش تراکنش‌های همگام‌سازی شده
  شامل:
  1. آمار تراکنش‌های همگام شده
  2. فیلترهای پیشرفته
  3. جدول تراکنش‌ها
  4. نمودارهای تحلیلی
  5. صفحه‌بندی
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
