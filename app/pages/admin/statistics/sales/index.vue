<template>
  <div class="sales-dashboard min-h-screen" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-green-700">آمار فروش</h1>
          <p class="text-gray-500 mt-1">گزارش کامل فروش، درآمد، محصولات پرفروش و تحلیل پیشرفته</p>
        </div>
        <div class="flex gap-2">
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
    <div class="px-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- Total Revenue -->
        <div class="bg-gradient-to-br from-green-500 to-green-600 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
          <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
          <div class="relative">
            <div class="flex items-center justify-between mb-2">
              <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0 01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582zM11 12.849v-1.698c.22.071.412.164.567.267.364.243.433.468.433.582 0 .114-.07.34-.433.582a2.305 2.305 0 01-.567.267z"/>
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535 4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324 2.246.48.32 1.054.545 1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51 1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0v-.092a4.535 4.535 0 001.676-.662C13.398 13.766 14 12.991 14 12c0-.99-.602-1.765-1.324-2.246A4.535 4.535 0 0011 9.092V7.151c.391.127.68.317.843.504a1 1 0 101.511-1.31c-.563-.649-1.413-1.076-2.354-1.253V5z" clip-rule="evenodd"/>
                </svg>
              </div>
              <div class="text-sm opacity-90">کل درآمد</div>
            </div>
            <div class="text-3xl font-bold mb-1">{{ stats.totalRevenue.toLocaleString('fa-IR') }}</div>
            <div class="flex items-center gap-1 text-sm opacity-90">
              <svg class="w-4 h-4 text-green-200" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M5.293 7.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L6.707 7.707a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
              </svg>
              ۱۲.۵% نسبت به ماه قبل
            </div>
          </div>
        </div>

        <!-- Today's Revenue -->
        <div class="bg-gradient-to-br from-blue-500 to-blue-600 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
          <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
          <div class="relative">
            <div class="flex items-center justify-between mb-2">
              <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd"/>
                </svg>
              </div>
              <div class="text-sm opacity-90">فروش امروز</div>
            </div>
            <div class="text-3xl font-bold mb-1">{{ stats.todayRevenue.toLocaleString('fa-IR') }}</div>
            <div class="text-sm opacity-90">تومان از ساعت 00:00</div>
          </div>
        </div>

        <!-- Total Orders -->
        <div class="bg-gradient-to-br from-purple-500 to-purple-600 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
          <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
          <div class="relative">
            <div class="flex items-center justify-between mb-2">
              <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z" clip-rule="evenodd"/>
                </svg>
              </div>
              <div class="text-sm opacity-90">تعداد سفارشات</div>
            </div>
            <div class="text-3xl font-bold mb-1">{{ stats.totalOrders.toLocaleString('fa-IR') }}</div>
            <div class="text-sm opacity-90">در این ماه</div>
          </div>
        </div>

        <!-- Average Order Value -->
        <div class="bg-gradient-to-br from-orange-500 to-red-500 text-white rounded-xl shadow-lg p-6 relative overflow-hidden">
          <div class="absolute top-0 right-0 w-24 h-24 bg-white bg-opacity-10 rounded-full -translate-y-6 translate-x-6"></div>
          <div class="relative">
            <div class="flex items-center justify-between mb-2">
              <div class="w-10 h-10 bg-white bg-opacity-20 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M4 4a2 2 0 00-2 2v4a2 2 0 002 2V6h10a2 2 0 00-2-2H4zm2 6a2 2 0 012-2h8a2 2 0 012 2v4a2 2 0 01-2 2H8a2 2 0 01-2-2v-4zm6 4a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd"/>
                </svg>
              </div>
              <div class="text-sm opacity-90">میانگین ارزش سفارش</div>
            </div>
            <div class="text-3xl font-bold mb-1">{{ stats.avgOrderValue.toLocaleString('fa-IR') }}</div>
            <div class="text-sm opacity-90">نرخ تبدیل: {{ stats.conversionRate }}%</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="px-6 mb-6">
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">فیلترها و تنظیمات</h3>
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
            <input type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm" placeholder="۱۴۰۲/۰۸/۰۱">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
            <input type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm" placeholder="۱۴۰۲/۰۸/۳۰">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm">
              <option>همه دسته‌ها</option>
              <option>الکترونیک</option>
              <option>پوشاک</option>
              <option>کتاب</option>
            </select>
          </div>
          <div class="flex items-end">
            <button class="w-full px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition">
              اعمال فیلتر
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Analytics Grid -->
    <div class="px-6 grid grid-cols-1 xl:grid-cols-12 gap-6 mb-6">
      <!-- Revenue Chart -->
      <div class="xl:col-span-8">
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">روند درآمد</h3>
          <div class="h-80 flex items-center justify-center bg-gray-50 rounded">
            <p class="text-gray-500">نمودار درآمد (Chart.js یا کتابخانه نمودار)</p>
          </div>
        </div>
      </div>
      
      <!-- Top Products -->
      <div class="xl:col-span-4">
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">محصولات پرفروش</h3>
          <div class="space-y-4">
            <div v-for="product in topProducts" :key="product.name" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div>
                <p class="font-medium">{{ product.name }}</p>
                <p class="text-sm text-gray-600">{{ product.sales }} فروش</p>
              </div>
              <div class="text-left">
                <p class="font-semibold text-green-600">{{ product.revenue.toLocaleString('fa-IR') }}</p>
                <p class="text-xs text-gray-500">تومان</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Secondary Analytics -->
    <div class="px-6 grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-6 mb-6">
      <!-- Revenue Trends -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">روندهای درآمد</h3>
        <div class="space-y-4">
          <div v-for="trend in revenueTrends" :key="trend.period" class="flex items-center justify-between p-3 border border-gray-200 rounded-lg">
            <span class="font-medium">{{ trend.period }}</span>
            <div class="text-left">
              <p class="font-semibold">{{ trend.value.toLocaleString('fa-IR') }}</p>
              <p class="text-sm" :class="trend.growth > 0 ? 'text-green-600' : 'text-red-600'">
                {{ trend.growth > 0 ? '+' : '' }}{{ trend.growth }}%
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Customer Segments -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">بخش‌بندی مشتریان</h3>
        <div class="space-y-4">
          <div v-for="segment in customerSegments" :key="segment.segment" class="p-3 border border-gray-200 rounded-lg">
            <div class="flex items-center justify-between mb-2">
              <span class="font-medium">{{ segment.segment }}</span>
              <span class="text-sm text-gray-600">{{ segment.count }} نفر</span>
            </div>
            <div class="text-sm">
              <span class="font-semibold text-green-600">{{ segment.revenue.toLocaleString('fa-IR') }} تومان</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Advanced Analytics -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">تحلیل‌های پیشرفته</h3>
        <div class="space-y-4">
          <div class="p-3 bg-blue-50 rounded-lg">
            <div class="flex items-center justify-between">
              <span class="text-sm text-blue-700">تحلیل کوهورت</span>
              <span class="font-semibold text-blue-700">{{ analytics.cohortAnalysis }}%</span>
            </div>
          </div>
          <div class="p-3 bg-red-50 rounded-lg">
            <div class="flex items-center justify-between">
              <span class="text-sm text-red-700">نرخ ترک</span>
              <span class="font-semibold text-red-700">{{ analytics.churnRate }}%</span>
            </div>
          </div>
          <div class="p-3 bg-green-50 rounded-lg">
            <div class="flex items-center justify-between">
              <span class="text-sm text-green-700">CLV میانگین</span>
              <span class="font-semibold text-green-700">{{ analytics.clv.toLocaleString('fa-IR') }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Sales Table -->
    <div class="px-6 mb-6">
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">آخرین فروش‌ها</h3>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-gray-50">
              <tr class="text-right">
                <th class="px-4 py-2">محصول</th>
                <th class="px-4 py-2">مشتری</th>
                <th class="px-4 py-2">مبلغ</th>
                <th class="px-4 py-2">تاریخ</th>
                <th class="px-4 py-2">عملیات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="sale in salesData" :key="sale.id" class="border-t">
                <td class="px-4 py-3 font-medium">{{ sale.product }}</td>
                <td class="px-4 py-3">{{ sale.customer }}</td>
                <td class="px-4 py-3">
                  <span class="font-semibold text-green-600">{{ sale.amount.toLocaleString('fa-IR') }} تومان</span>
                </td>
                <td class="px-4 py-3">{{ sale.date }}</td>
                <td class="px-4 py-3">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">مشاهده</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Export Options -->
    <div class="px-6 mb-6">
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">گزارش‌گیری و خروجی</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <button class="flex items-center justify-center gap-2 px-4 py-3 bg-green-100 text-green-700 rounded-lg hover:bg-green-200 transition">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clip-rule="evenodd"/>
            </svg>
            خروجی Excel
          </button>
          <button class="flex items-center justify-center gap-2 px-4 py-3 bg-red-100 text-red-700 rounded-lg hover:bg-red-200 transition">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 11-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L7.707 6.707a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
            </svg>
            خروجی PDF
          </button>
          <button class="flex items-center justify-center gap-2 px-4 py-3 bg-blue-100 text-blue-700 rounded-lg hover:bg-blue-200 transition">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"/>
            </svg>
            گزارش تفصیلی
          </button>
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

// Sample data for components
const stats = {
  totalRevenue: 12500000,
  todayRevenue: 245000,
  weekRevenue: 1850000,
  monthRevenue: 7200000,
  totalOrders: 1245,
  avgOrderValue: 185000,
  conversionRate: 3.2
}

const topProducts = [
  { name: 'گوشی آیفون ۱۴', sales: 125, revenue: 18750000 },
  { name: 'لپ‌تاپ ایسوس', sales: 89, revenue: 12650000 },
  { name: 'تبلت سامسونگ', sales: 156, revenue: 9360000 }
]

const revenueTrends = [
  { period: 'این ماه', value: 7200000, growth: 12.5 },
  { period: 'ماه گذشته', value: 6400000, growth: -2.1 },
  { period: 'سه ماه گذشته', value: 19800000, growth: 8.7 }
]

const customerSegments = [
  { segment: 'مشتریان VIP', count: 128, revenue: 8500000 },
  { segment: 'مشتریان عادی', count: 1850, revenue: 15200000 },
  { segment: 'مشتریان جدید', count: 245, revenue: 2800000 }
]

const salesData = [
  { id: 1, product: 'گوشی موبایل', amount: 1250000, date: '۱۴۰۲/۰۸/۱۵', customer: 'احمد احمدی' },
  { id: 2, product: 'لپ‌تاپ گیمینگ', amount: 8500000, date: '۱۴۰۲/۰۸/۱۴', customer: 'مریم مرادی' },
  { id: 3, product: 'تبلت', amount: 3200000, date: '۱۴۰۲/۰۸/۱۳', customer: 'علی علوی' }
]

const analytics = {
  cohortAnalysis: 85.2,
  churnRate: 2.8,
  clv: 1850000,
  rfmSegmentation: {
    champions: 15,
    loyalCustomers: 35,
    potentialLoyalists: 25,
    newCustomers: 25
  }
}

// Meta and SEO
useHead({
  title: 'آمار فروش - پنل مدیریت',
  meta: [
    { name: 'description', content: 'مشاهده و تحلیل آمار فروش، درآمد و عملکرد محصولات' }
  ]
})
</script> 

