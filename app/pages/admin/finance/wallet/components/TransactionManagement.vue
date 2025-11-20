<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">💳 مدیریت تراکنش‌ها</h2>
      <p class="text-gray-600">مدیریت کامل تراکنش‌های ورودی و خروجی کیف پول</p>
    </div>

    <!-- آمار کلی تراکنش‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کل تراکنش‌ها -->
      <div class="border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کل تراکنش‌ها</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">امروز</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ transactionStats.totalTransactions }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ transactionStats.dailyGrowth }}%</span>
          <span class="mx-2">از دیروز</span>
        </div>
      </div>

      <!-- تراکنش‌های موفق -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">موفق</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">عالی</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ transactionStats.successfulTransactions }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ transactionStats.successRate }}%</span>
          <span class="mx-2">نرخ موفقیت</span>
        </div>
      </div>

      <!-- تراکنش‌های ناموفق -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">ناموفق</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">هشدار</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ transactionStats.failedTransactions }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ transactionStats.failureRate }}%</span>
          <span class="mx-2">نرخ شکست</span>
        </div>
      </div>

      <!-- تراکنش‌های در انتظار -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">در انتظار</h3>
          <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">انتظار</span>
        </div>
        <div class="text-3xl font-bold text-yellow-600 mb-2">{{ transactionStats.pendingTransactions }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-yellow-500">{{ transactionStats.pendingAmount }}</span>
          <span class="mx-2">مبلغ کل</span>
        </div>
      </div>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <!-- نوع تراکنش -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع تراکنش</label>
          <select v-model="typeFilter" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="charge">شارژ</option>
            <option value="withdraw">برداشت</option>
            <option value="transfer">انتقال</option>
            <option value="fee">کارمزد</option>
          </select>
        </div>

        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select v-model="statusFilter" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="success">موفق</option>
            <option value="failed">ناموفق</option>
            <option value="pending">در انتظار</option>
          </select>
        </div>

        <!-- بازه زمانی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <input v-model="fromDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>

        <!-- جستجو -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
v-model="queryText" type="text" placeholder="شماره تراکنش یا نام کاربر" 
                 class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>

      <div class="flex justify-between items-center">
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500" @click="applyFilters">
          اعمال فیلتر
        </button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500" @click="clearFilters">
          پاک کردن فیلترها
        </button>
      </div>
    </div>

    <!-- جدول تراکنش‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">لیست تراکنش‌ها</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500" @click="exportCsv">خروجی CSV</button>
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            گزارش
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">شماره تراکنش</th>
              <th class="px-4 py-3 text-gray-700 font-medium">کاربر</th>
              <th class="px-4 py-3 text-gray-700 font-medium">نوع</th>
              <th class="px-4 py-3 text-gray-700 font-medium">مبلغ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">توضیحات</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="transaction in transactions" :key="transaction.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-gray-900 font-medium">{{ transaction.transactionId }}</td>
              <td class="px-4 py-3">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center text-xs font-medium text-gray-600">
                    {{ transaction.userInitials }}
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ transaction.userName }}</div>
                    <div class="text-xs text-gray-500">{{ transaction.userEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3">
                <span :class="getTransactionTypeClass(transaction.type)" class="px-2 py-1 text-xs rounded-full">
                  {{ transaction.type }}
                </span>
              </td>
              <td class="px-4 py-3 font-medium" :class="transaction.amount > 0 ? 'text-green-600' : 'text-red-600'">
                {{ transaction.amount > 0 ? '+' : '' }}{{ formatCurrency(transaction.amount) }}
              </td>
              <td class="px-4 py-3 text-gray-600">{{ transaction.date }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(transaction.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ transaction.status }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600 max-w-xs truncate">{{ transaction.description }}</td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button v-if="transaction.status === 'pending'" class="text-green-600 hover:text-green-800 text-sm" @click="updateStatus(transaction.id, 'success')">تأیید</button>
                  <button v-if="transaction.status === 'pending'" class="text-red-600 hover:text-red-800 text-sm" @click="updateStatus(transaction.id, 'failed')">رد</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="flex items-center justify-between mt-6">
        <div class="text-sm text-gray-700">
          نمایش {{ paginationInfo.start }} تا {{ paginationInfo.end }} از {{ paginationInfo.total }} تراکنش
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">قبلی</button>
          <button class="px-3 py-1 text-sm bg-blue-600 text-white rounded-lg">1</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">2</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">3</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">بعدی</button>
        </div>
      </div>
    </div>

    <!-- نمودار تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار تراکنش‌ها بر اساس نوع -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تراکنش‌ها بر اساس نوع</h3>
        <div class="space-y-4">
          <div v-for="type in transactionTypes" :key="type.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: type.color }"></div>
              <span class="text-sm text-gray-700">{{ type.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ type.count }}</span>
              <span class="text-sm text-gray-500">({{ type.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار تراکنش‌ها بر اساس وضعیت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تراکنش‌ها بر اساس وضعیت</h3>
        <div class="space-y-4">
          <div v-for="status in transactionStatuses" :key="status.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: status.color }"></div>
              <span class="text-sm text-gray-700">{{ status.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ status.count }}</span>
              <span class="text-sm text-gray-500">({{ status.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref, watchEffect } from 'vue';
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
const transactionStats = reactive({
  totalTransactions: 0,
  dailyGrowth: 0,
  successfulTransactions: 0,
  successRate: 0,
  failedTransactions: 0,
  failureRate: 0,
  pendingTransactions: 0,
  pendingAmount: '0 تومان'
})

const transactions = ref<any[]>([])
const paginationInfo = reactive({ start: 1, end: 0, total: 0 })
const page = ref(1)
const pageSize = ref(20)
const typeFilter = ref('')
const statusFilter = ref('')
const fromDate = ref('')
const queryText = ref('')

const { data, refresh } = await useFetch('/api/admin/wallet/transactions', {
  method: 'GET',
  query: computed(() => ({
    page: page.value,
    pageSize: pageSize.value,
    type: mapType(typeFilter.value),
    status: statusFilter.value,
    from: fromDate.value || undefined,
    q: queryText.value || undefined,
  })),
  credentials: 'include',
  key: () => `admin-wallet-transactions-${page.value}-${pageSize.value}-${typeFilter.value}-${statusFilter.value}-${fromDate.value}-${queryText.value}`,
  defaultCache: true,
  swr: true,
})

watchEffect(() => {
  const res: any = data.value
  if (!res) return
  const items = Array.isArray(res.items) ? res.items : []
  transactions.value = items.map((r: any) => ({
    id: r.id,
    transactionId: r.id,
    userName: r.username || r.user_id,
    userEmail: r.email || '-',
    userInitials: (r.username || '-').slice(0,2),
    type: r.type === 'credit' ? 'شارژ' : 'کسر',
    amount: r.type === 'credit' ? Number(r.amount) : -Number(r.amount),
    date: r.created_at,
    status: r.status === 'success' ? 'موفق' : (r.status === 'pending' ? 'در انتظار' : 'ناموفق'),
    description: r.gateway || r.method || '-'
  }))
  paginationInfo.total = Number(res.total || 0)
  paginationInfo.start = (page.value-1) * pageSize.value + 1
  paginationInfo.end = Math.min(page.value * pageSize.value, paginationInfo.total)

  // آمار ساده
  transactionStats.totalTransactions = paginationInfo.total
  const succ = items.filter((it: any) => it.status === 'success').length
  const fail = items.filter((it: any) => it.status === 'failed').length
  const pend = items.filter((it: any) => it.status === 'pending').length
  const totalShown = Math.max(1, items.length)
  transactionStats.successfulTransactions = succ
  transactionStats.failedTransactions = fail
  transactionStats.pendingTransactions = pend
  transactionStats.successRate = Math.round((succ / totalShown) * 1000) / 10
  transactionStats.failureRate = Math.round((fail / totalShown) * 1000) / 10
})

function mapType(v: string) {
  if (v === 'charge') return 'credit'
  if (v === 'withdraw') return 'debit'
  return ''
}

function applyFilters() {
  page.value = 1
  refresh()
}
function clearFilters() {
  typeFilter.value = ''
  statusFilter.value = ''
  fromDate.value = ''
  queryText.value = ''
  page.value = 1
  refresh()
}

async function updateStatus(id: number, status: 'success' | 'failed') {
  await $fetch(`/api/admin/wallet/transactions/${id}/status`, {
    method: 'POST',
    body: { status },
    credentials: 'include',
  })
  await refresh()
}

function exportCsv() {
  const rows = [
    ['id','user','type','amount','date','status','desc'],
    ...transactions.value.map((t:any)=>[
      t.id,
      t.userName,
      t.type,
      t.amount,
      t.date,
      t.status,
      t.description,
    ])
  ]
  const csv = rows.map(r=>r.map(x=>`"${String(x).replace(/"/g,'""')}"`).join(',')).join('\n')
  const blob = new Blob([csv], { type: 'text/csv;charset=utf-8;' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `transactions_${new Date().toISOString().slice(0,10)}.csv`
  a.click()
  URL.revokeObjectURL(url)
}

const transactionTypes = computed(() => {
  const items: any[] = transactions.value
  const count = (name: string) => items.filter((t) => (name === 'شارژ' ? t.amount > 0 : t.amount < 0)).length
  const total = Math.max(1, items.length)
  return [
    { name: 'شارژ', count: count('شارژ'), percentage: Math.round((count('شارژ')/total)*1000)/10, color: '#10B981' },
    { name: 'کسر', count: count('کسر'), percentage: Math.round((count('کسر')/total)*1000)/10, color: '#EF4444' },
  ]
})

const transactionStatuses = computed(() => {
  const items: any[] = transactions.value
  const succ = items.filter((it) => it.status === 'موفق').length
  const fail = items.filter((it) => it.status === 'ناموفق').length
  const pend = items.filter((it) => it.status === 'در انتظار').length
  const total = Math.max(1, items.length)
  return [
    { name: 'موفق', count: succ, percentage: Math.round((succ/total)*1000)/10, color: '#10B981' },
    { name: 'ناموفق', count: fail, percentage: Math.round((fail/total)*1000)/10, color: '#EF4444' },
    { name: 'در انتظار', count: pend, percentage: Math.round((pend/total)*1000)/10, color: '#F59E0B' },
  ]
})

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع کلاس نوع تراکنش
const getTransactionTypeClass = (type: string) => {
  switch (type) {
    case 'شارژ':
      return 'bg-green-100 text-green-800'
    case 'برداشت':
      return 'bg-red-100 text-red-800'
    case 'انتقال':
      return 'bg-blue-100 text-blue-800'
    case 'کارمزد':
      return 'bg-orange-100 text-orange-800'
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
  1. آمار کلی تراکنش‌ها (کل، موفق، ناموفق، در انتظار)
  2. فیلترها و جستجوی پیشرفته
  3. جدول تراکنش‌ها با جزئیات کامل
  4. نمودارهای تحلیلی بر اساس نوع و وضعیت
  5. صفحه‌بندی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
