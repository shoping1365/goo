<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">مدیریت تراکنش‌های دوطرفه</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و نظارت بر تراکنش‌های دوطرفه بین سیستم و نرم‌افزار حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          :disabled="isSyncing"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="syncDualTransactions"
        >
          <svg v-if="isSyncing" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          {{ isSyncing ? 'در حال همگام‌سازی...' : 'همگام‌سازی' }}
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="showAddForm = true"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          افزودن تراکنش جدید
        </button>
      </div>
    </div>

    <!-- آمار تراکنش‌های دوطرفه -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ stats.totalTransactions }}</div>
            <div class="text-sm text-blue-700">کل تراکنش‌ها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ stats.syncedTransactions }}</div>
            <div class="text-sm text-green-700">همگام شده</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ stats.pendingTransactions }}</div>
            <div class="text-sm text-yellow-700">در انتظار</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ stats.conflictTransactions }}</div>
            <div class="text-sm text-red-700">تضاد</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="bg-gray-50 rounded-lg p-6">
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
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="synced">همگام شده</option>
            <option value="pending">در انتظار</option>
            <option value="conflict">تضاد</option>
            <option value="error">خطا</option>
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
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ</label>
          <input
            v-model="filters.date"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="applyFilters"
          />
        </div>
      </div>
    </div>

    <!-- جدول تراکنش‌های دوطرفه -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تراکنش‌های دوطرفه</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">شماره تراکنش</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">مبلغ</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ</th>
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
                <div class="flex items-center gap-2">
                  <div
                    class="w-2 h-2 rounded-full"
                    :class="getStatusColor(transaction.status)"
                  ></div>
                  <span
                    class="px-2 py-1 rounded-full text-xs font-medium"
                    :class="getStatusClass(transaction.status)"
                  >
                    {{ getStatusLabel(transaction.status) }}
                  </span>
                </div>
              </td>
              <td class="py-3 px-4 text-gray-900">{{ transaction.date }}</td>
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
                  <button
                    v-if="transaction.status === 'conflict'"
                    class="p-1 text-yellow-600 hover:text-yellow-800"
                    title="حل تضاد"
                    @click="resolveConflict(transaction)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
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

// وضعیت همگام‌سازی
const isSyncing = ref(false)
const showAddForm = ref(false)

// فیلترها
const filters = ref({
  transactionType: '',
  status: '',
  software: '',
  date: ''
})

// آمار
const stats = ref({
  totalTransactions: 15420,
  syncedTransactions: 14850,
  pendingTransactions: 420,
  conflictTransactions: 150
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
    date: '2024-01-15',
    softwareName: 'هلو',
    softwareLogo: '/images/helo-logo.png',
    status: 'synced'
  },
  {
    id: 2,
    transactionNumber: 'TXN-002',
    reference: 'پرداخت #67890',
    type: 'payment',
    amount: 850000,
    currency: 'تومان',
    date: '2024-01-15',
    softwareName: 'پارسیان',
    softwareLogo: '/images/parsian-logo.png',
    status: 'pending'
  },
  {
    id: 3,
    transactionNumber: 'TXN-003',
    reference: 'خرید #11111',
    type: 'purchase',
    amount: 2100000,
    currency: 'تومان',
    date: '2024-01-14',
    softwareName: 'سپیدار',
    softwareLogo: '/images/sepidar-logo.png',
    status: 'conflict'
  }
])

// تراکنش‌های فیلتر شده
const filteredTransactions = computed(() => {
  return transactions.value.filter(transaction => {
    if (filters.value.transactionType && transaction.type !== filters.value.transactionType) return false
    if (filters.value.status && transaction.status !== filters.value.status) return false
    if (filters.value.software && transaction.softwareName.toLowerCase() !== filters.value.software) return false
    if (filters.value.date && transaction.date !== filters.value.date) return false
    return true
  })
})

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

// رنگ وضعیت
const getStatusColor = (status: string) => {
  const colors = {
    synced: 'bg-green-500',
    pending: 'bg-yellow-500',
    conflict: 'bg-red-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    synced: 'bg-green-100 text-green-700',
    pending: 'bg-yellow-100 text-yellow-700',
    conflict: 'bg-red-100 text-red-700',
    error: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    synced: 'همگام شده',
    pending: 'در انتظار',
    conflict: 'تضاد',
    error: 'خطا'
  }
  return labels[status] || status
}

// اعمال فیلترها
const applyFilters = () => {
  // TODO: اعمال فیلترها
  console.log('فیلترهای تراکنش اعمال شد:', filters.value)
}

// همگام‌سازی تراکنش‌های دوطرفه
const syncDualTransactions = async () => {
  try {
    isSyncing.value = true
    // TODO: همگام‌سازی تراکنش‌ها
    console.log('همگام‌سازی تراکنش‌های دوطرفه شروع شد')
    
    // شبیه‌سازی همگام‌سازی
    await new Promise(resolve => setTimeout(resolve, 3000))
    
    console.log('همگام‌سازی تراکنش‌های دوطرفه تکمیل شد')
  } catch (error) {
    console.error('خطا در همگام‌سازی تراکنش‌های دوطرفه:', error)
  } finally {
    isSyncing.value = false
  }
}

// مشاهده تراکنش
const viewTransaction = (transaction: any) => {
  // TODO: مشاهده جزئیات تراکنش
  console.log('مشاهده تراکنش:', transaction)
}

// همگام‌سازی مجدد تراکنش
const resyncTransaction = async (transaction: any) => {
  try {
    // TODO: همگام‌سازی مجدد تراکنش
    console.log('همگام‌سازی مجدد تراکنش:', transaction)
  } catch (error) {
    console.error('خطا در همگام‌سازی مجدد:', error)
  }
}

// حل تضاد
const resolveConflict = (transaction: any) => {
  // TODO: حل تضاد تراکنش
  console.log('حل تضاد تراکنش:', transaction)
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
  کامپوننت مدیریت تراکنش‌های دوطرفه
  شامل:
  1. آمار و خلاصه وضعیت
  2. فیلترهای پیشرفته
  3. جدول تراکنش‌های دوطرفه
  4. عملیات همگام‌سازی
  5. حل تضادها
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
