<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">💸 مدیریت برداشت</h2>
      <p class="text-gray-600">مدیریت درخواست‌های برداشت و تأیید آنها</p>
    </div>

    <!-- آمار برداشت -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- درخواست‌های جدید -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">درخواست‌های جدید</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">جدید</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ withdrawalStats.newRequests }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-blue-500">{{ formatCurrency(withdrawalStats.newAmount) }}</span>
          <span class="mx-2">مبلغ کل</span>
        </div>
      </div>

      <!-- در انتظار تأیید -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">در انتظار تأیید</h3>
          <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">انتظار</span>
        </div>
        <div class="text-3xl font-bold text-yellow-600 mb-2">{{ withdrawalStats.pendingApproval }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-yellow-500">{{ formatCurrency(withdrawalStats.pendingAmount) }}</span>
          <span class="mx-2">مبلغ کل</span>
        </div>
      </div>

      <!-- تأیید شده -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تأیید شده</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">تأیید</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ withdrawalStats.approved }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ formatCurrency(withdrawalStats.approvedAmount) }}</span>
          <span class="mx-2">مبلغ کل</span>
        </div>
      </div>

      <!-- رد شده -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">رد شده</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">رد</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ withdrawalStats.rejected }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ formatCurrency(withdrawalStats.rejectedAmount) }}</span>
          <span class="mx-2">مبلغ کل</span>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="new">جدید</option>
            <option value="pending">در انتظار تأیید</option>
            <option value="approved">تأیید شده</option>
            <option value="rejected">رد شده</option>
          </select>
        </div>

        <!-- بازه مبلغ -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه مبلغ</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="low">کمتر از 1 میلیون</option>
            <option value="medium">1 تا 10 میلیون</option>
            <option value="high">بیش از 10 میلیون</option>
          </select>
        </div>

        <!-- بازه زمانی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="today">امروز</option>
            <option value="week">هفته گذشته</option>
            <option value="month">ماه گذشته</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>

        <!-- جستجو -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input type="text" placeholder="شماره درخواست یا نام کاربر" 
                 class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>

      <div class="flex justify-between items-center">
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          اعمال فیلتر
        </button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
          پاک کردن فیلترها
        </button>
      </div>
    </div>

    <!-- جدول درخواست‌های برداشت -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">درخواست‌های برداشت</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button @click="bulkUpdate('success')" class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">تأیید همه</button>
          <button @click="bulkUpdate('failed')" class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">رد همه</button>
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            خروجی اکسل
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </th>
              <th class="px-4 py-3 text-gray-700 font-medium">شماره درخواست</th>
              <th class="px-4 py-3 text-gray-700 font-medium">کاربر</th>
              <th class="px-4 py-3 text-gray-700 font-medium">مبلغ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">حساب بانکی</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ درخواست</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="withdrawal in withdrawals" :key="withdrawal.id" class="hover:bg-gray-50">
              <td class="px-4 py-3">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </td>
              <td class="px-4 py-3 text-gray-900 font-medium">{{ withdrawal.requestId }}</td>
              <td class="px-4 py-3">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center text-xs font-medium text-gray-600">
                    {{ withdrawal.userInitials }}
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ withdrawal.userName }}</div>
                    <div class="text-xs text-gray-500">{{ withdrawal.userEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 font-medium text-gray-900">{{ formatCurrency(withdrawal.amount) }}</td>
              <td class="px-4 py-3 text-gray-600">
                <div>
                  <div class="text-sm">{{ withdrawal.bankName }}</div>
                  <div class="text-xs text-gray-500">{{ withdrawal.accountNumber }}</div>
                </div>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ withdrawal.requestDate }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(withdrawal.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ withdrawal.status }}
                </span>
              </td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button v-if="withdrawal.status === 'new' || withdrawal.status === 'در انتظار تأیید'" @click="updateStatus(withdrawal.id, 'success')" class="text-green-600 hover:text-green-800 text-sm">تأیید</button>
                  <button v-if="withdrawal.status === 'new' || withdrawal.status === 'در انتظار تأیید'" @click="updateStatus(withdrawal.id, 'failed')" class="text-red-600 hover:text-red-800 text-sm">رد</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="flex items-center justify-between mt-6">
        <div class="text-sm text-gray-700">
          نمایش {{ paginationInfo.start }} تا {{ paginationInfo.end }} از {{ paginationInfo.total }} درخواست
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

    <!-- تنظیمات برداشت -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- محدودیت‌های برداشت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">محدودیت‌های برداشت</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ برداشت</label>
            <input type="number" :value="withdrawalSettings.minAmount" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ برداشت</label>
            <input type="number" :value="withdrawalSettings.maxAmount" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کارمزد برداشت (%)</label>
            <input type="number" step="0.1" :value="withdrawalSettings.feePercentage" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان پردازش (ساعت)</label>
            <input type="number" :value="withdrawalSettings.processingTime" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره تنظیمات
          </button>
        </div>
      </div>

      <!-- آمار سریع -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">آمار سریع</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">میانگین مبلغ برداشت:</span>
            <span class="text-sm font-medium text-gray-900">{{ formatCurrency(withdrawalStats.averageAmount) }}</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">سریع‌ترین پردازش:</span>
            <span class="text-sm font-medium text-gray-900">{{ withdrawalStats.fastestProcessing }} دقیقه</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">کندترین پردازش:</span>
            <span class="text-sm font-medium text-gray-900">{{ withdrawalStats.slowestProcessing }} ساعت</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">نرخ تأیید:</span>
            <span class="text-sm font-medium text-green-600">{{ withdrawalStats.approvalRate }}%</span>
          </div>
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm text-gray-700">نرخ رد:</span>
            <span class="text-sm font-medium text-red-600">{{ withdrawalStats.rejectionRate }}%</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, watchEffect } from 'vue';
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
// آمار برداشت (محاسبه از API)
const withdrawalStats = reactive({
  newRequests: 0,
  newAmount: 0,
  pendingApproval: 0,
  pendingAmount: 0,
  approved: 0,
  approvedAmount: 0,
  rejected: 0,
  rejectedAmount: 0,
  averageAmount: 0,
  fastestProcessing: 0,
  slowestProcessing: 0,
  approvalRate: 0,
  rejectionRate: 0
})

// تنظیمات برداشت
const withdrawalSettings = {
  minAmount: 100000,
  maxAmount: 50000000,
  feePercentage: 1.5,
  processingTime: 24
}

// درخواست‌های برداشت از API تراکنش‌ها (method=withdraw)
const withdrawals = ref<any[]>([])
const page = ref(1)
const pageSize = ref(20)
const statusFilter = ref('')
const { data: txs, refresh } = await useFetch('/api/admin/wallet/transactions', {
  method: 'GET',
  query: { page, pageSize, type: 'debit', method: 'withdraw', status: statusFilter },
  credentials: 'include',
  key: () => `admin-wallet-withdraw-${page.value}-${pageSize.value}-${statusFilter.value}`,
  defaultCache: true,
})
const { data: summary } = await useFetch('/api/admin/wallet/summary', { method:'GET', credentials:'include', key:'admin-wallet-summary', defaultCache:true })
watchEffect(()=>{
  const res: any = txs.value
  if (res && Array.isArray(res.items)) {
    withdrawals.value = res.items.map((r:any)=>({
      id: r.id,
      requestId: r.id,
      userName: r.username || r.user_id,
      userEmail: r.email || '-',
      userInitials: (r.username || '-').slice(0,2),
      amount: Number(r.amount||0),
      bankName: '-',
      accountNumber: '-',
      requestDate: r.created_at,
      status: r.status === 'success' ? 'تأیید شده' : (r.status==='pending' ? 'در انتظار تأیید' : (r.status==='failed' ? 'رد شده' : 'جدید'))
    }))
    // آمار ساده از لیست
    const list = withdrawals.value
    withdrawalStats.newRequests = list.filter((x:any)=>x.status==='جدید').length
    withdrawalStats.pendingApproval = list.filter((x:any)=>x.status==='در انتظار تأیید').length
    withdrawalStats.approved = list.filter((x:any)=>x.status==='تأیید شده').length
    withdrawalStats.rejected = list.filter((x:any)=>x.status==='رد شده').length
    withdrawalStats.newAmount = list.filter((x:any)=>x.status==='جدید').reduce((s:number,x:any)=>s+x.amount,0)
    withdrawalStats.pendingAmount = list.filter((x:any)=>x.status==='در انتظار تأیید').reduce((s:number,x:any)=>s+x.amount,0)
    withdrawalStats.approvedAmount = list.filter((x:any)=>x.status==='تأیید شده').reduce((s:number,x:any)=>s+x.amount,0)
    withdrawalStats.rejectedAmount = list.filter((x:any)=>x.status==='رد شده').reduce((s:number,x:any)=>s+x.amount,0)
    const total = Math.max(1, list.length)
    withdrawalStats.approvalRate = Math.round((withdrawalStats.approved/total)*1000)/10
    withdrawalStats.rejectionRate = Math.round((withdrawalStats.rejected/total)*1000)/10
  }
})

async function updateStatus(id: number, status: 'success'|'failed'|'pending') {
  await $fetch(`/api/admin/wallet/transactions/${id}/status`, { method: 'POST', body: { status }, credentials: 'include' })
  await refresh()
}

async function bulkUpdate(status: 'success'|'failed') {
  const ids = withdrawals.value
    .filter((w:any)=> w.status==='new' || w.status==='در انتظار تأیید')
    .map((w:any)=> w.id)
  for (const id of ids) {
    try { await $fetch(`/api/admin/wallet/transactions/${id}/status`, { method: 'POST', body: { status }, credentials: 'include' }) } catch {}
  }
  await refresh()
}

// اطلاعات صفحه‌بندی
const paginationInfo = reactive({ start:1, end:0, total:0 })
watchEffect(()=>{
  const res:any = txs.value
  if (!res) return
  paginationInfo.total = Number(res.total||0)
  paginationInfo.start = (page.value-1)*pageSize.value + 1
  paginationInfo.end = Math.min(page.value*pageSize.value, paginationInfo.total)
})

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'جدید':
      return 'bg-blue-100 text-blue-800'
    case 'در انتظار تأیید':
      return 'bg-yellow-100 text-yellow-800'
    case 'تأیید شده':
      return 'bg-green-100 text-green-800'
    case 'رد شده':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. آمار کلی درخواست‌های برداشت (جدید، در انتظار، تأیید شده، رد شده)
  2. فیلترهای پیشرفته برای جستجو
  3. جدول درخواست‌های برداشت با امکان تأیید/رد دسته‌ای
  4. تنظیمات محدودیت‌های برداشت
  5. آمار سریع پردازش
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
