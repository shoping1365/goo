<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">تحلیل‌های پیشرفته</h2>
        <p class="mt-1 text-sm text-gray-600">تحلیل عمیق رفتار کاربران و عملکرد سیستم</p>
      </div>
      <button class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
        ایجاد گزارش جدید
      </button>
    </div>

    <!-- آمار کلیدی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">CTR</p>
            <p class="text-2xl font-bold">12.4%</p>
            <p class="text-xs opacity-75">+2.1% از ماه قبل</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Conversion</p>
            <p class="text-2xl font-bold">8.7%</p>
            <p class="text-xs opacity-75">+1.3% از ماه قبل</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Revenue</p>
            <p class="text-2xl font-bold">$45.2K</p>
            <p class="text-xs opacity-75">+15.3% از ماه قبل</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Engagement</p>
            <p class="text-2xl font-bold">67.8%</p>
            <p class="text-xs opacity-75">+5.2% از ماه قبل</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">Churn Rate</p>
            <p class="text-2xl font-bold">3.2%</p>
            <p class="text-xs opacity-75">-0.8% از ماه قبل</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل رفتار کاربران -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- الگوهای خرید -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">الگوهای خرید</h3>
          </div>
          <button class="bg-blue-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-blue-700 transition-colors">
            جزئیات
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="pattern in purchasePatterns" :key="pattern.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <h4 class="font-medium text-gray-900">{{ pattern.name }}</h4>
              <span class="text-sm font-medium" :class="pattern.trendClass">{{ pattern.trend }}</span>
            </div>
            <p class="text-sm text-gray-600 mb-2">{{ pattern.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>فرکانس: {{ pattern.frequency }}</span>
              <span>میانگین: {{ pattern.average }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- تحلیل سشن‌ها -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">تحلیل سشن‌ها</h3>
          </div>
          <button class="bg-green-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-green-700 transition-colors">
            جزئیات
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="session in sessionAnalytics" :key="session.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <h4 class="font-medium text-gray-900">{{ session.name }}</h4>
              <span class="text-sm font-medium" :class="session.valueClass">{{ session.value }}</span>
            </div>
            <p class="text-sm text-gray-600 mb-2">{{ session.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>تغییر: {{ session.change }}</span>
              <span>روند: {{ session.trend }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل محصولات -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- محصولات پرفروش -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">محصولات پرفروش</h3>
          </div>
          <button class="bg-purple-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-purple-700 transition-colors">
            مشاهده همه
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="product in topProducts" :key="product.id" class="flex items-center space-x-4 space-x-reverse">
            <div class="w-10 h-10 bg-gray-200 rounded-lg flex items-center justify-center">
              <span class="text-sm font-medium text-gray-600">{{ product.rank }}</span>
            </div>
            <div class="flex-1">
              <h4 class="font-medium text-gray-900">{{ product.name }}</h4>
              <p class="text-sm text-gray-600">{{ product.category }}</p>
            </div>
            <div class="text-right">
              <p class="font-medium text-gray-900">{{ product.sales }}</p>
              <p class="text-xs text-gray-500">{{ product.revenue }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- تحلیل دسته‌بندی‌ها -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">تحلیل دسته‌بندی‌ها</h3>
          </div>
          <button class="bg-yellow-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-yellow-700 transition-colors">
            جزئیات
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="category in categoryAnalytics" :key="category.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <h4 class="font-medium text-gray-900">{{ category.name }}</h4>
              <span class="text-sm font-medium text-gray-900">{{ category.performance }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2 mb-2">
              <div class="bg-yellow-500 h-2 rounded-full" :style="{ width: category.performance + '%' }"></div>
            </div>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>محصولات: {{ category.products }}</span>
              <span>فروش: {{ category.sales }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل پیشرفته -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-indigo-100 rounded-lg">
            <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">تحلیل پیشرفته</h3>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
            تحلیل A/B
          </button>
          <button class="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors">
            پیش‌بینی
          </button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div class="space-y-4">
          <h4 class="font-medium text-gray-900">تحلیل همبستگی</h4>
          <div v-for="correlation in correlations" :key="correlation.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-medium text-gray-900">{{ correlation.factor1 }}</span>
              <span class="text-sm font-medium text-gray-900">{{ correlation.correlation }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-indigo-500 h-2 rounded-full" :style="{ width: Math.abs(correlation.correlation) * 100 + '%' }"></div>
            </div>
            <p class="text-xs text-gray-500 mt-1">{{ correlation.factor2 }}</p>
          </div>
        </div>
        
        <div class="space-y-4">
          <h4 class="font-medium text-gray-900">تحلیل روند</h4>
          <div v-for="trend in trends" :key="trend.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-medium text-gray-900">{{ trend.metric }}</span>
              <span class="text-sm font-medium" :class="trend.trendClass">{{ trend.change }}</span>
            </div>
            <p class="text-xs text-gray-600 mb-2">{{ trend.description }}</p>
            <div class="flex items-center text-xs text-gray-500">
              <svg class="w-4 h-4 mr-1" :class="trend.iconClass" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="trend.icon"></path>
              </svg>
              {{ trend.period }}
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <h4 class="font-medium text-gray-900">پیش‌بینی‌ها</h4>
          <div v-for="prediction in predictions" :key="prediction.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-medium text-gray-900">{{ prediction.metric }}</span>
              <span class="text-sm font-medium text-gray-900">{{ prediction.value }}</span>
            </div>
            <p class="text-xs text-gray-600 mb-2">{{ prediction.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>دقت: {{ prediction.accuracy }}%</span>
              <span>بازه: {{ prediction.range }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش‌های تحلیلی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌های تحلیلی</h3>
      </div>
      <div class="p-6">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">گزارش</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین اجرا</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="report in analyticalReports" :key="report.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="report.bgColor">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="report.icon"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ report.name }}</div>
                      <div class="text-sm text-gray-500">{{ report.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ report.type }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ report.lastRun }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="report.statusClass">
                    {{ report.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button class="text-indigo-600 hover:text-indigo-900 mr-3">مشاهده</button>
                  <button class="text-green-600 hover:text-green-900 mr-3">اجرا</button>
                  <button class="text-red-600 hover:text-red-900">حذف</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// الگوهای خرید
const purchasePatterns = ref([
  {
    id: 1,
    name: 'خرید در آخر هفته',
    description: 'کاربران تمایل بیشتری به خرید در روزهای آخر هفته دارند',
    frequency: 'هفتگی',
    average: '2.3 محصول',
    trend: '+15%',
    trendClass: 'text-green-600'
  },
  {
    id: 2,
    name: 'خرید گروهی',
    description: 'کاربران اغلب چندین محصول مرتبط را با هم خریداری می‌کنند',
    frequency: 'روزانه',
    average: '3.1 محصول',
    trend: '+8%',
    trendClass: 'text-green-600'
  },
  {
    id: 3,
    name: 'خرید فصلی',
    description: 'الگوی خرید بر اساس فصل‌های سال',
    frequency: 'فصلی',
    average: '5.2 محصول',
    trend: '+22%',
    trendClass: 'text-green-600'
  }
])

// تحلیل سشن‌ها
const sessionAnalytics = ref([
  {
    id: 1,
    name: 'مدت زمان سشن',
    description: 'میانگین مدت زمان حضور کاربران در سایت',
    value: '8.5 دقیقه',
    valueClass: 'text-blue-600',
    change: '+12%',
    trend: 'صعودی'
  },
  {
    id: 2,
    name: 'صفحات بازدید شده',
    description: 'تعداد صفحات بازدید شده در هر سشن',
    value: '4.2 صفحه',
    valueClass: 'text-green-600',
    change: '+8%',
    trend: 'صعودی'
  },
  {
    id: 3,
    name: 'نرخ خروج',
    description: 'درصد کاربرانی که بدون خرید سایت را ترک می‌کنند',
    value: '23.4%',
    valueClass: 'text-red-600',
    change: '-5%',
    trend: 'نزولی'
  }
])

// محصولات پرفروش
const topProducts = ref([
  {
    id: 1,
    rank: 1,
    name: 'لپ‌تاپ گیمینگ',
    category: 'الکترونیک',
    sales: '234 فروش',
    revenue: '$45,678'
  },
  {
    id: 2,
    rank: 2,
    name: 'هندزفری بی‌سیم',
    category: 'الکترونیک',
    sales: '189 فروش',
    revenue: '$12,345'
  },
  {
    id: 3,
    rank: 3,
    name: 'کتاب مدیریت',
    category: 'کتاب',
    sales: '156 فروش',
    revenue: '$3,456'
  },
  {
    id: 4,
    rank: 4,
    name: 'کفش ورزشی',
    category: 'پوشاک',
    sales: '134 فروش',
    revenue: '$8,901'
  }
])

// تحلیل دسته‌بندی‌ها
const categoryAnalytics = ref([
  {
    id: 1,
    name: 'الکترونیک',
    performance: 85,
    products: 234,
    sales: '1,234'
  },
  {
    id: 2,
    name: 'پوشاک',
    performance: 72,
    products: 456,
    sales: '2,345'
  },
  {
    id: 3,
    name: 'کتاب',
    performance: 68,
    products: 123,
    sales: '567'
  },
  {
    id: 4,
    name: 'خانه و آشپزخانه',
    performance: 61,
    products: 345,
    sales: '890'
  }
])

// تحلیل همبستگی
const correlations = ref([
  {
    id: 1,
    factor1: 'قیمت',
    factor2: 'فروش',
    correlation: 0.75
  },
  {
    id: 2,
    factor1: 'بازخورد',
    factor2: 'فروش',
    correlation: 0.82
  },
  {
    id: 3,
    factor1: 'تخفیف',
    factor2: 'فروش',
    correlation: 0.68
  }
])

// تحلیل روند
const trends = ref([
  {
    id: 1,
    metric: 'فروش ماهانه',
    change: '+15.3%',
    trendClass: 'text-green-600',
    description: 'افزایش مداوم فروش در 6 ماه گذشته',
    period: '6 ماه',
    icon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6',
    iconClass: 'text-green-500'
  },
  {
    id: 2,
    metric: 'کاربران فعال',
    change: '+8.7%',
    trendClass: 'text-green-600',
    description: 'رشد پایدار کاربران فعال',
    period: '3 ماه',
    icon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6',
    iconClass: 'text-green-500'
  },
  {
    id: 3,
    metric: 'نرخ تبدیل',
    change: '+2.1%',
    trendClass: 'text-green-600',
    description: 'بهبود تدریجی نرخ تبدیل',
    period: '1 ماه',
    icon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6',
    iconClass: 'text-green-500'
  }
])

// پیش‌بینی‌ها
const predictions = ref([
  {
    id: 1,
    metric: 'فروش آینده',
    value: '$52.3K',
    description: 'پیش‌بینی فروش ماه آینده',
    accuracy: 94,
    range: '±5%'
  },
  {
    id: 2,
    metric: 'کاربران جدید',
    value: '1,234',
    description: 'پیش‌بینی کاربران جدید',
    accuracy: 87,
    range: '±8%'
  },
  {
    id: 3,
    metric: 'محصولات پرفروش',
    value: '15 محصول',
    description: 'پیش‌بینی محصولات پرفروش',
    accuracy: 92,
    range: '±3%'
  }
])

// گزارش‌های تحلیلی
const analyticalReports = ref([
  {
    id: 1,
    name: 'گزارش عملکرد ماهانه',
    description: 'تحلیل جامع عملکرد سیستم در ماه گذشته',
    type: 'ماهانه',
    lastRun: '2 روز پیش',
    status: 'تکمیل شده',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-blue-500',
    icon: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'
  },
  {
    id: 2,
    name: 'تحلیل رفتار کاربران',
    description: 'بررسی الگوهای رفتاری کاربران',
    type: 'هفتگی',
    lastRun: '1 روز پیش',
    status: 'در حال اجرا',
    statusClass: 'bg-yellow-100 text-yellow-800',
    bgColor: 'bg-green-500',
    icon: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z'
  },
  {
    id: 3,
    name: 'پیش‌بینی فروش',
    description: 'پیش‌بینی فروش آینده بر اساس داده‌های تاریخی',
    type: 'روزانه',
    lastRun: '6 ساعت پیش',
    status: 'تکمیل شده',
    statusClass: 'bg-green-100 text-green-800',
    bgColor: 'bg-purple-500',
    icon: 'M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z'
  }
])
</script> 
