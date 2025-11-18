<template>
  <div class="min-h-screen" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-purple-700">آمار مشتریان</h1>
          <p class="text-gray-500 mt-1">تحلیل رفتار، درآمد و وفاداری مشتریان</p>
        </div>
        <div class="flex gap-2">
          <button class="px-4 py-2 bg-purple-500 text-white rounded-lg shadow hover:bg-purple-600 transition">
            گزارش جامع
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
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-purple-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">کل مشتریان</p>
              <p class="text-2xl font-bold text-purple-600">{{ customerStats.totalCustomers.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-purple-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-green-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">مشتریان فعال</p>
              <p class="text-2xl font-bold text-green-600">{{ customerStats.activeCustomers.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-green-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-blue-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">مشتریان جدید</p>
              <p class="text-2xl font-bold text-blue-600">{{ customerStats.newCustomers.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-blue-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8 9a3 3 0 100-6 3 3 0 000 6zM8 11a6 6 0 016 6H2a6 6 0 016-6zM16 7a1 1 0 10-2 0v1h-1a1 1 0 100 2h1v1a1 1 0 102 0v-1h1a1 1 0 100-2h-1V7z"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-yellow-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">نرخ وفاداری</p>
              <p class="text-2xl font-bold text-yellow-600">{{ customerStats.loyaltyRate }}%</p>
            </div>
            <div class="bg-yellow-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-red-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">ارزش طول عمر</p>
              <p class="text-2xl font-bold text-red-600">{{ customerStats.avgLifetimeValue.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-red-100 p-2 rounded-full">
              <svg class="w-5 h-5 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0 01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582zM11 12.849v-1.698c.22.071.412.164.567.267.364.243.433.468.433.582 0 .114-.07.34-.433.582a2.305 2.305 0 01-.567.267z"/>
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535 4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324 2.246.48.32 1.054.545 1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51 1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0v-.092a4.535 4.535 0 001.676-.662C13.398 13.766 14 12.991 14 12c0-.99-.602-1.765-1.324-2.246A4.535 4.535 0 0011 9.092V7.151c.391.127.68.317.843.504a1 1 0 101.511-1.31c-.563-.649-1.413-1.076-2.354-1.253V5z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-white rounded-lg shadow-sm p-6 border-r-4 border-indigo-500">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600">مشتریان امروز</p>
              <p class="text-2xl font-bold text-indigo-600">{{ customerStats.todayRegistrations.toLocaleString('fa-IR') }}</p>
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

    <!-- Customer Segmentation -->
    <div class="px-6 space-y-6">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Customer Demographics -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">دموگرافی مشتریان</h3>
          <div class="space-y-4">
            <div>
              <div class="flex justify-between items-center mb-2">
                <span class="text-sm text-gray-600">مردان</span>
                <span class="text-sm font-medium">۵۸٪</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-blue-500 h-2 rounded-full" style="width: 58%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between items-center mb-2">
                <span class="text-sm text-gray-600">زنان</span>
                <span class="text-sm font-medium">۴۲٪</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-pink-500 h-2 rounded-full" style="width: 42%"></div>
              </div>
            </div>
          </div>
          
          <div class="mt-6">
            <h4 class="font-medium mb-3">گروه‌های سنی</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
                <span class="text-sm">۱۸-۲۵ سال</span>
                <span class="font-medium">۲۳٪</span>
              </div>
              <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
                <span class="text-sm">۲۶-۳۵ سال</span>
                <span class="font-medium">۳۴٪</span>
              </div>
              <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
                <span class="text-sm">۳۶-۴۵ سال</span>
                <span class="font-medium">۲۸٪</span>
              </div>
              <div class="flex items-center justify-between p-2 bg-gray-50 rounded">
                <span class="text-sm">۴۶+ سال</span>
                <span class="font-medium">۱۵٪</span>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Customer Value Segments -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">بخش‌بندی مشتریان</h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-6 bg-gradient-to-r from-yellow-50 to-yellow-100 rounded-lg border border-yellow-200">
              <div>
                <span class="font-medium text-yellow-800">مشتریان VIP</span>
                <p class="text-sm text-yellow-600">بالای ۱۰ میلیون تومان خرید</p>
              </div>
              <span class="text-xl font-bold text-yellow-800">۱۲۸</span>
            </div>
            <div class="flex items-center justify-between p-6 bg-gradient-to-r from-green-50 to-green-100 rounded-lg border border-green-200">
              <div>
                <span class="font-medium text-green-800">مشتریان وفادار</span>
                <p class="text-sm text-green-600">۳-۱۰ میلیون تومان خرید</p>
              </div>
              <span class="text-xl font-bold text-green-800">۴۵۶</span>
            </div>
            <div class="flex items-center justify-between p-6 bg-gradient-to-r from-blue-50 to-blue-100 rounded-lg border border-blue-200">
              <div>
                <span class="font-medium text-blue-800">مشتریان عادی</span>
                <p class="text-sm text-blue-600">۱-۳ میلیون تومان خرید</p>
              </div>
              <span class="text-xl font-bold text-blue-800">۱,۲۸۵</span>
            </div>
            <div class="flex items-center justify-between p-6 bg-gradient-to-r from-gray-50 to-gray-100 rounded-lg border border-gray-200">
              <div>
                <span class="font-medium text-gray-800">مشتریان جدید</span>
                <p class="text-sm text-gray-600">کمتر از ۱ میلیون تومان</p>
              </div>
              <span class="text-xl font-bold text-gray-800">۶۷۸</span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Top Customers Table -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">برترین مشتریان</h3>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-gray-50">
              <tr class="text-right">
                <th class="px-4 py-2">رتبه</th>
                <th class="px-4 py-2">نام مشتری</th>
                <th class="px-4 py-2">کل خرید</th>
                <th class="px-4 py-2">تعداد سفارش</th>
                <th class="px-4 py-2">آخرین خرید</th>
                <th class="px-4 py-2">وضعیت</th>
                <th class="px-4 py-2">عملیات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(customer, index) in topCustomers" :key="customer.id" class="border-t">
                <td class="px-4 py-3">
                  <span class="flex items-center">
                    <svg v-if="index === 0" class="w-4 h-4 text-yellow-500 ml-1" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
                    </svg>
                    {{ index + 1 }}
                  </span>
                </td>
                <td class="px-4 py-3 font-medium">{{ customer.name }}</td>
                <td class="px-4 py-3">{{ customer.totalPurchase.toLocaleString('fa-IR') }} تومان</td>
                <td class="px-4 py-3">{{ customer.orderCount }}</td>
                <td class="px-4 py-3">{{ customer.lastPurchase }}</td>
                <td class="px-4 py-3">
                  <span :class="getCustomerStatusClass(customer.status)" class="px-2 py-1 rounded-full text-xs">
                    {{ customer.status }}
                  </span>
                </td>
                <td class="px-4 py-3">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">مشاهده</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      
      <!-- Customer Acquisition Chart -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">روند جذب مشتریان جدید</h3>
        <div class="h-64 flex items-center justify-center bg-gray-50 rounded">
          <p class="text-gray-500">نمودار روند جذب مشتریان در اینجا نمایش داده می‌شود</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useAuth: () => { user: any; hasPermission: (perm: string) => boolean }
declare const useHead: (head: { title?: string; meta?: any[] }) => void
</script>

<script setup lang="ts">

// Set admin layout
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Sample data
const customerStats = {
  totalCustomers: 8547,
  activeCustomers: 6234,
  newCustomers: 245,
  loyaltyRate: 72.5,
  avgLifetimeValue: 3450000,
  todayRegistrations: 18
}

const topCustomers = [
  { id: 1, name: 'احمد احمدی', totalPurchase: 24500000, orderCount: 48, lastPurchase: '۱۴۰۲/۰۸/۱۰', status: 'VIP' },
  { id: 2, name: 'مریم مرادی', totalPurchase: 18750000, orderCount: 35, lastPurchase: '۱۴۰۲/۰۸/۱۲', status: 'VIP' },
  { id: 3, name: 'علی علوی', totalPurchase: 15200000, orderCount: 28, lastPurchase: '۱۴۰۲/۰۸/۰۸', status: 'وفادار' },
  { id: 4, name: 'زهرا زارعی', totalPurchase: 12800000, orderCount: 22, lastPurchase: '۱۴۰۲/۰۸/۱۴', status: 'وفادار' },
  { id: 5, name: 'حسین حسینی', totalPurchase: 9450000, orderCount: 18, lastPurchase: '۱۴۰۲/۰۸/۰۶', status: 'عادی' }
]

const getCustomerStatusClass = (status: string) => {
  switch (status) {
    case 'VIP':
      return 'bg-yellow-100 text-yellow-800'
    case 'وفادار':
      return 'bg-green-100 text-green-800'
    case 'عادی':
      return 'bg-blue-100 text-blue-800'
    case 'جدید':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Meta and SEO
useHead({
  title: 'آمار مشتریان - پنل مدیریت',
  meta: [
    { name: 'description', content: 'تحلیل و آمار کامل مشتریان، دموگرافی و وفاداری' }
  ]
})
</script>

<style scoped>
/* Additional styling if needed */
</style> 

