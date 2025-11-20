<template>
  <div class="min-h-screen" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-blue-700">آمار سفارش‌ها</h1>
          <p class="text-gray-500 mt-1">گزارش کامل سفارش‌ها، وضعیت‌ها و تحلیل‌های مرتبط</p>
        </div>
        <div class="flex gap-2">
          <button class="px-4 py-2 bg-blue-500 text-white rounded-lg shadow hover:bg-blue-600 transition">
            گزارش ماهانه
          </button>
          <button class="px-4 py-2 bg-green-500 text-white rounded-lg shadow hover:bg-green-600 transition">
            خروجی اکسل
          </button>
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg shadow hover:bg-gray-200 transition">
            بروزرسانی
          </button>
        </div>
      </div>
    </div>
    
    <!-- Summary Cards -->
    <div class="px-6 py-6">
      <div class="grid grid-cols-2 md:grid-cols-4 xl:grid-cols-6 gap-6">
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-blue-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">کل سفارش‌ها</p>
              <p class="text-2xl font-bold text-blue-600">{{ orderStats.totalOrders.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-blue-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-green-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">سفارش‌های تکمیل شده</p>
              <p class="text-2xl font-bold text-green-600">{{ orderStats.completedOrders.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-green-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-yellow-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">در انتظار پردازش</p>
              <p class="text-2xl font-bold text-yellow-600">{{ orderStats.pendingOrders.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-yellow-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-red-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">لغو شده</p>
              <p class="text-2xl font-bold text-red-600">{{ orderStats.cancelledOrders.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-red-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-purple-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">میانگین ارزش سفارش</p>
              <p class="text-2xl font-bold text-purple-600">{{ orderStats.avgOrderValue.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-purple-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0 01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582zM11 12.849v-1.698c.22.071.412.164.567.267.364.243.433.468.433.582 0 .114-.07.34-.433.582a2.305 2.305 0 01-.567.267z"/>
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535 4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324 2.246.48.32 1.054.545 1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51 1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0v-.092a4.535 4.535 0 001.676-.662C13.398 13.766 14 12.991 14 12c0-.99-.602-1.765-1.324-2.246A4.535 4.535 0 0011 9.092V7.151c.391.127.68.317.843.504a1 1 0 101.511-1.31c-.563-.649-1.413-1.076-2.354-1.253V5z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-indigo-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">سفارش‌های امروز</p>
              <p class="text-2xl font-bold text-indigo-600">{{ orderStats.todayOrders.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-indigo-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-indigo-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="px-6 mb-6">
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">فیلترها</h3>
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
            <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option>۷ روز گذشته</option>
              <option>۳۰ روز گذشته</option>
              <option>۳ ماه گذشته</option>
              <option>سال جاری</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option>همه وضعیت‌ها</option>
              <option>تکمیل شده</option>
              <option>در انتظار</option>
              <option>لغو شده</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ</label>
            <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option>همه مبالغ</option>
              <option>کمتر از ۱ میلیون</option>
              <option>۱ تا ۵ میلیون</option>
              <option>بیشتر از ۵ میلیون</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
            <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
              <option>همه مشتریان</option>
              <option>مشتریان جدید</option>
              <option>مشتریان قدیمی</option>
              <option>مشتریان VIP</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- Charts and Tables -->
    <div class="px-6 space-y-6">
      <!-- Orders Chart -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">روند سفارش‌ها</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <p class="text-gray-500">نمودار روند سفارش‌ها در اینجا نمایش داده می‌شود</p>
        </div>
      </div>

      <!-- Order Status Distribution -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">توزیع وضعیت سفارش‌ها</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-green-50 rounded">
              <span class="text-green-700">تکمیل شده</span>
              <span class="font-semibold text-green-700">۷۲٪</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-yellow-50 rounded">
              <span class="text-yellow-700">در انتظار</span>
              <span class="font-semibold text-yellow-700">۱۸٪</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-red-50 rounded">
              <span class="text-red-700">لغو شده</span>
              <span class="font-semibold text-red-700">۱۰٪</span>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">روش‌های پرداخت</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 bg-blue-50 rounded">
              <span class="text-blue-700">کارت بانکی</span>
              <span class="font-semibold text-blue-700">۶۵٪</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-purple-50 rounded">
              <span class="text-purple-700">کیف پول</span>
              <span class="font-semibold text-purple-700">۲۰٪</span>
            </div>
            <div class="flex items-center justify-between p-3 bg-gray-50 rounded">
              <span class="text-gray-700">پرداخت در محل</span>
              <span class="font-semibold text-gray-700">۱۵٪</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Orders Table -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">آخرین سفارش‌ها</h3>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-gray-50">
              <tr class="text-right">
                <th class="px-4 py-2">شماره سفارش</th>
                <th class="px-4 py-2">مشتری</th>
                <th class="px-4 py-2">مبلغ</th>
                <th class="px-4 py-2">وضعیت</th>
                <th class="px-4 py-2">تاریخ</th>
                <th class="px-4 py-2">عملیات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="order in recentOrders" :key="order.id" class="border-t">
                <td class="px-4 py-3 font-medium">#{{ order.id }}</td>
                <td class="px-4 py-3">{{ order.customer }}</td>
                <td class="px-4 py-3">{{ order.amount.toLocaleString('fa-IR') }} تومان</td>
                <td class="px-4 py-3">
                  <span :class="getStatusClass(order.status)" class="px-2 py-1 rounded-full text-xs">
                    {{ order.status }}
                  </span>
                </td>
                <td class="px-4 py-3">{{ order.date }}</td>
                <td class="px-4 py-3">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">مشاهده</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
// declare const useAuth: () => { user: unknown; hasPermission: (perm: string) => boolean }
declare const useHead: (head: { title?: string; meta?: Record<string, unknown>[] }) => void
</script>

<script setup lang="ts">

// Set admin layout
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Sample data
const orderStats = {
  totalOrders: 2547,
  completedOrders: 1834,
  pendingOrders: 458,
  cancelledOrders: 255,
  avgOrderValue: 1850000,
  todayOrders: 42
}

const recentOrders = [
  { id: 1001, customer: 'احمد احمدی', amount: 2450000, status: 'تکمیل شده', date: '۱۴۰۲/۰۸/۱۵' },
  { id: 1002, customer: 'مریم مرادی', amount: 1750000, status: 'در انتظار', date: '۱۴۰۲/۰۸/۱۵' },
  { id: 1003, customer: 'علی علوی', amount: 3200000, status: 'تکمیل شده', date: '۱۴۰۲/۰۸/۱۴' },
  { id: 1004, customer: 'زهرا زارعی', amount: 890000, status: 'لغو شده', date: '۱۴۰۲/۰۸/۱۴' },
  { id: 1005, customer: 'حسین حسینی', amount: 4500000, status: 'تکمیل شده', date: '۱۴۰۲/۰۸/۱۳' }
]

const getStatusClass = (status: string) => {
  switch (status) {
    case 'تکمیل شده':
      return 'bg-green-100 text-green-800'
    case 'در انتظار':
      return 'bg-yellow-100 text-yellow-800'
    case 'لغو شده':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Meta and SEO
useHead({
  title: 'آمار سفارش‌ها - پنل مدیریت',
  meta: [
    { name: 'description', content: 'مشاهده و تحلیل آمار سفارش‌ها، وضعیت‌ها و روندهای فروش' }
  ]
})
</script>

<style scoped>
/* Additional styling if needed */
</style> 

