<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-xl font-semibold text-gray-900">پایش ترافیک وب‌سایت</h2>
        <p class="text-sm text-gray-600 mt-1">تحلیل ترافیک و تغییرات آن در طول زمان</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <select v-model="selectedPeriod" class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="7d">7 روز گذشته</option>
          <option value="30d">30 روز گذشته</option>
          <option value="90d">90 روز گذشته</option>
          <option value="1y">1 سال گذشته</option>
        </select>
        <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 text-sm" @click="exportData">
          <svg class="w-4 h-4 ml-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی
        </button>
      </div>
    </div>

    <!-- Key Metrics Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">کل بازدیدکنندگان</p>
            <p class="text-2xl font-bold text-blue-900">{{ formatNumber(trafficData.totalVisitors) }}</p>
            <p class="text-sm text-blue-700 mt-1">
              <span :class="trafficData.visitorGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ trafficData.visitorGrowth >= 0 ? '+' : '' }}{{ trafficData.visitorGrowth }}%
              </span>
              نسبت به دوره قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-600">صفحات مشاهده شده</p>
            <p class="text-2xl font-bold text-green-900">{{ formatNumber(trafficData.pageViews) }}</p>
            <p class="text-sm text-green-700 mt-1">
              <span :class="trafficData.pageViewGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ trafficData.pageViewGrowth >= 0 ? '+' : '' }}{{ trafficData.pageViewGrowth }}%
              </span>
              نسبت به دوره قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">زمان حضور</p>
            <p class="text-2xl font-bold text-purple-900">{{ trafficData.avgSessionDuration }} دقیقه</p>
            <p class="text-sm text-purple-700 mt-1">
              <span :class="trafficData.durationGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ trafficData.durationGrowth >= 0 ? '+' : '' }}{{ trafficData.durationGrowth }}%
              </span>
              نسبت به دوره قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-orange-50 to-orange-100 rounded-xl p-6 border border-orange-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-600">نرخ پرش</p>
            <p class="text-2xl font-bold text-orange-900">{{ trafficData.bounceRate }}%</p>
            <p class="text-sm text-orange-700 mt-1">
              <span :class="trafficData.bounceRateChange <= 0 ? 'text-green-600' : 'text-red-600'">
                {{ trafficData.bounceRateChange >= 0 ? '+' : '' }}{{ trafficData.bounceRateChange }}%
              </span>
              نسبت به دوره قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Traffic Chart -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-semibold text-gray-900">روند ترافیک در طول زمان</h3>
        <div class="flex items-center space-x-4 space-x-reverse">
          <button 
            v-for="metric in chartMetrics" 
            :key="metric.id"
            :class="[
              selectedMetric === metric.id
                ? 'bg-blue-100 text-blue-700 border-blue-300'
                : 'bg-gray-100 text-gray-600 border-gray-300',
              'px-3 py-1 rounded-md text-sm border transition-colors duration-200'
            ]"
            @click="selectedMetric = metric.id"
          >
            {{ metric.name }}
          </button>
        </div>
      </div>
      
      <!-- Chart Placeholder -->
      <div class="h-80 bg-gray-50 rounded-lg border-2 border-dashed border-gray-300 flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
          <p class="text-gray-500 text-lg">نمودار ترافیک</p>
          <p class="text-gray-400 text-sm">در اینجا نمودار روند ترافیک نمایش داده می‌شود</p>
        </div>
      </div>
    </div>

    <!-- Traffic Sources -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">منابع ترافیک</h3>
        <div class="space-y-4">
          <div v-for="source in trafficSources" :key="source.name" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: source.color }"></div>
              <span class="text-sm font-medium text-gray-700">{{ source.name }}</span>
            </div>
            <div class="flex items-center space-x-3 space-x-reverse">
              <span class="text-sm text-gray-600">{{ source.percentage }}%</span>
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div class="h-2 rounded-full" :style="{ width: source.percentage + '%', backgroundColor: source.color }"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">صفحات پر بازدید</h3>
        <div class="space-y-4">
          <div v-for="page in topPages" :key="page.url" class="flex items-center justify-between">
            <div class="flex-1 min-w-0">
              <p class="text-sm font-medium text-gray-900 truncate">{{ page.title }}</p>
              <p class="text-xs text-gray-500 truncate">{{ page.url }}</p>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <span class="text-sm text-gray-600">{{ formatNumber(page.views) }} بازدید</span>
              <span class="text-sm text-gray-500">{{ page.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// متغیرهای reactive
const selectedPeriod = ref('30d')
const selectedMetric = ref('visitors')

// داده‌های ترافیک
const trafficData = ref({
  totalVisitors: 125430,
  visitorGrowth: 12.5,
  pageViews: 456789,
  pageViewGrowth: 8.3,
  avgSessionDuration: 4.2,
  durationGrowth: -2.1,
  bounceRate: 34.2,
  bounceRateChange: -5.8
})

// متریک‌های نمودار
const chartMetrics = [
  { id: 'visitors', name: 'بازدیدکنندگان' },
  { id: 'pageViews', name: 'صفحات مشاهده شده' },
  { id: 'sessions', name: 'جلسات' },
  { id: 'bounceRate', name: 'نرخ پرش' }
]

// منابع ترافیک
const trafficSources = ref([
  { name: 'جستجوی ارگانیک', percentage: 45, color: '#3B82F6' },
  { name: 'جستجوی پولی', percentage: 25, color: '#10B981' },
  { name: 'شبکه‌های اجتماعی', percentage: 18, color: '#8B5CF6' },
  { name: 'مستقیم', percentage: 12, color: '#F59E0B' }
])

// صفحات پر بازدید
const topPages = ref([
  { title: 'صفحه اصلی', url: '/', views: 45678, percentage: 15.2 },
  { title: 'محصولات', url: '/products', views: 34256, percentage: 11.4 },
  { title: 'درباره ما', url: '/about', views: 28934, percentage: 9.6 },
  { title: 'تماس با ما', url: '/contact', views: 23456, percentage: 7.8 },
  { title: 'وبلاگ', url: '/blog', views: 19876, percentage: 6.6 }
])

// توابع
const formatNumber = (num) => {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toLocaleString('fa-IR')
}

const exportData = () => {

  // منطق خروجی داده‌ها
}
</script>
