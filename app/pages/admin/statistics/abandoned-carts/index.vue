<template>
  <div class="min-h-screen" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-amber-700">سبدهای خرید رها شده</h1>
          <p class="text-gray-500 mt-1">تحلیل و بازیابی سبدهای خرید رها شده</p>
        </div>
        <div class="flex gap-2">
          <button class="px-4 py-2 bg-amber-500 text-white rounded-lg shadow hover:bg-amber-600 transition">
            ارسال ایمیل بازیابی
          </button>
          <button class="px-4 py-2 bg-green-500 text-white rounded-lg shadow hover:bg-green-600 transition">
            خروجی اکسل
          </button>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="px-6 py-4 bg-white border-b">
      <div class="flex flex-wrap gap-6">
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-amber-500">
          <option>بازه زمانی</option>
          <option>امروز</option>
          <option>هفته گذشته</option>
          <option>ماه گذشته</option>
        </select>
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-amber-500">
          <option>نوع مشتری</option>
          <option>جدید</option>
          <option>بازگشتی</option>
          <option>VIP</option>
        </select>
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-amber-500">
          <option>وضعیت</option>
          <option>ایمیل نشده</option>
          <option>ایمیل شده</option>
          <option>بازیابی شده</option>
        </select>
        <button class="px-4 py-2 bg-amber-500 text-white rounded-lg text-sm hover:bg-amber-600 transition">اعمال فیلتر</button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg text-sm hover:bg-gray-200 transition">پاک کردن</button>
      </div>
    </div>
    
    <!-- Summary Cards -->
    <div class="px-6 py-6">
      <div class="grid grid-cols-2 md:grid-cols-4 xl:grid-cols-6 gap-6">
        <div class="bg-gradient-to-br from-amber-50 to-amber-100 rounded-lg shadow-sm p-6 border border-amber-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-amber-600">کل سبدهای رها شده</p>
              <p class="text-xl font-bold text-amber-700">{{ stats.totalCarts.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-amber-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-amber-700" fill="currentColor" viewBox="0 0 20 20">
                <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-red-50 to-red-100 rounded-lg shadow-sm p-6 border border-red-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-red-600">نرخ رها کردن</p>
              <p class="text-xl font-bold text-red-700">{{ stats.abandonmentRate }}%</p>
            </div>
            <div class="bg-red-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-red-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-lg shadow-sm p-6 border border-purple-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-purple-600">ارزش از دست رفته</p>
              <p class="text-lg font-bold text-purple-700">{{ stats.lostValue.toLocaleString('fa-IR') }}M</p>
            </div>
            <div class="bg-purple-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-purple-700" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0 01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582z"/>
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535 4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324 2.246.48.32 1.054.545 1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51 1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-lg shadow-sm p-6 border border-green-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-green-600">بازیابی شده</p>
              <p class="text-xl font-bold text-green-700">{{ stats.recovered.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-green-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-green-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg shadow-sm p-6 border border-blue-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-blue-600">نرخ بازیابی</p>
              <p class="text-xl font-bold text-blue-700">{{ stats.recoveryRate }}%</p>
            </div>
            <div class="bg-blue-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-blue-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-yellow-50 to-yellow-100 rounded-lg shadow-sm p-6 border border-yellow-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-yellow-600">امروز</p>
              <p class="text-xl font-bold text-yellow-700">{{ stats.todayCount.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-yellow-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-yellow-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div class="px-6 space-y-6">
      <!-- Abandoned Carts Table -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold">سبدهای خرید رها شده</h3>
          <span class="bg-amber-100 text-amber-800 px-2 py-1 rounded-full text-xs">{{ abandonedCarts.length }} مورد</span>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-gray-50">
              <tr class="text-right">
                <th class="px-4 py-2">مشتری</th>
                <th class="px-4 py-2">نوع</th>
                <th class="px-4 py-2">ارزش سبد</th>
                <th class="px-4 py-2">تعداد آیتم</th>
                <th class="px-4 py-2">زمان رها کردن</th>
                <th class="px-4 py-2">وضعیت</th>
                <th class="px-4 py-2">عملیات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="cart in abandonedCarts" :key="cart.id" class="border-t">
                <td class="px-4 py-3">
                  <div>
                    <p class="font-medium">{{ cart.customer }}</p>
                    <p class="text-xs text-gray-600">{{ cart.email }}</p>
                  </div>
                </td>
                <td class="px-4 py-3">
                  <span :class="getCustomerTypeClass(cart.customerType)" class="px-2 py-1 rounded-full text-xs">
                    {{ cart.customerType }}
                  </span>
                </td>
                <td class="px-4 py-3 font-bold">{{ cart.value.toLocaleString('fa-IR') }} تومان</td>
                <td class="px-4 py-3">{{ cart.itemCount }}</td>
                <td class="px-4 py-3">{{ cart.abandonedTime }}</td>
                <td class="px-4 py-3">
                  <span :class="getCartStatusClass(cart.status)" class="px-2 py-1 rounded-full text-xs">
                    {{ cart.status }}
                  </span>
                </td>
                <td class="px-4 py-3">
                  <div class="flex gap-2">
                    <button class="text-amber-600 hover:text-amber-800 text-xs">ایمیل</button>
                    <button class="text-blue-600 hover:text-blue-800 text-xs">مشاهده</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Recovery Analysis -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Time Analysis -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">تحلیل زمانی رها کردن</h3>
          <div class="space-y-3">
            <div v-for="timeSlot in timeAnalysis" :key="timeSlot.period" class="p-3 bg-gray-50 rounded-lg">
              <div class="flex justify-between items-center mb-2">
                <span class="font-medium">{{ timeSlot.period }}</span>
                <span class="text-amber-600 font-bold">{{ timeSlot.count }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-amber-500 h-2 rounded-full" :style="`width: ${timeSlot.percentage}%`"></div>
              </div>
              <div class="text-xs text-gray-600 mt-1">{{ timeSlot.percentage }}% از کل</div>
            </div>
          </div>
        </div>

        <!-- Customer Segments -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">بخش‌بندی مشتریان</h3>
          <div class="space-y-3">
            <div v-for="segment in customerSegments" :key="segment.type" class="p-3 bg-gradient-to-r from-amber-50 to-amber-100 rounded-lg border border-amber-200">
              <div class="flex justify-between items-center mb-2">
                <span class="font-medium text-amber-800">{{ segment.type }}</span>
                <span class="text-amber-700 font-bold">{{ segment.count }}</span>
              </div>
              <div class="grid grid-cols-3 gap-6 text-sm">
                <div>
                  <p class="text-gray-600">میانگین ارزش</p>
                  <p class="font-bold">{{ segment.avgValue }}K</p>
                </div>
                <div>
                  <p class="text-gray-600">نرخ بازیابی</p>
                  <p class="font-bold">{{ segment.recoveryRate }}%</p>
                </div>
                <div>
                  <p class="text-gray-600">درآمد بازیابی</p>
                  <p class="font-bold">{{ segment.recoveredRevenue }}M</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
// declare const useAuth: () => { user: any; hasPermission: (perm: string) => boolean }
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
const stats = {
  totalCarts: 1847,
  abandonmentRate: 68.5,
  lostValue: 285,
  recovered: 245,
  recoveryRate: 13.3,
  todayCount: 89
}

const abandonedCarts = [
  { id: 1, customer: 'احمد احمدی', email: 'ahmad@example.com', customerType: 'جدید', value: 2450000, itemCount: 3, abandonedTime: '۲ ساعت پیش', status: 'منتظر ایمیل' },
  { id: 2, customer: 'مریم مرادی', email: 'maryam@example.com', customerType: 'بازگشتی', value: 1750000, itemCount: 2, abandonedTime: '۴ ساعت پیش', status: 'ایمیل شده' },
  { id: 3, customer: 'علی علوی', email: 'ali@example.com', customerType: 'VIP', value: 5200000, itemCount: 4, abandonedTime: '۶ ساعت پیش', status: 'پیگیری شده' },
  { id: 4, customer: 'زهرا زارعی', email: 'zahra@example.com', customerType: 'معمولی', value: 890000, itemCount: 1, abandonedTime: '۸ ساعت پیش', status: 'منتظر ایمیل' },
  { id: 5, customer: 'حسین حسینی', email: 'hossein@example.com', customerType: 'جدید', value: 3400000, itemCount: 5, abandonedTime: '۱۰ ساعت پیش', status: 'بازیابی شده' }
]

const timeAnalysis = [
  { period: 'کمتر از ۱ ساعت', count: 245, percentage: 35 },
  { period: '۱ تا ۶ ساعت', count: 189, percentage: 27 },
  { period: '۶ تا ۲۴ ساعت', count: 156, percentage: 22 },
  { period: '۱ تا ۳ روز', count: 89, percentage: 13 },
  { period: 'بیش از ۳ روز', count: 23, percentage: 3 }
]

const customerSegments = [
  { type: 'مشتریان جدید', count: 456, avgValue: 1250, recoveryRate: 8.5, recoveredRevenue: 45 },
  { type: 'مشتریان بازگشتی', count: 678, avgValue: 1850, recoveryRate: 15.2, recoveredRevenue: 89 },
  { type: 'مشتریان VIP', count: 123, avgValue: 4200, recoveryRate: 28.6, recoveredRevenue: 156 },
  { type: 'مشتریان معمولی', count: 590, avgValue: 980, recoveryRate: 11.4, recoveredRevenue: 34 }
]

const getCustomerTypeClass = (type: string) => {
  switch (type) {
    case 'VIP':
      return 'bg-purple-100 text-purple-800'
    case 'بازگشتی':
      return 'bg-blue-100 text-blue-800'
    case 'جدید':
      return 'bg-green-100 text-green-800'
    case 'معمولی':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getCartStatusClass = (status: string) => {
  switch (status) {
    case 'بازیابی شده':
      return 'bg-green-100 text-green-800'
    case 'ایمیل شده':
      return 'bg-blue-100 text-blue-800'
    case 'پیگیری شده':
      return 'bg-yellow-100 text-yellow-800'
    case 'منتظر ایمیل':
      return 'bg-amber-100 text-amber-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Meta and SEO
useHead({
  title: 'سبدهای خرید رها شده - پنل مدیریت',
  meta: [
    { name: 'description', content: 'تحلیل و بازیابی سبدهای خرید رها شده' }
  ]
})
</script> 

