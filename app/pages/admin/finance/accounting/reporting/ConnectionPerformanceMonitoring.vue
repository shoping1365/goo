<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">نظارت بر عملکرد اتصال</h4>
        <p class="text-sm text-gray-600 mt-1">نظارت و تحلیل عملکرد اتصال‌های حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="refreshMetrics"
          :disabled="isRefreshing"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg v-if="isRefreshing" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          {{ isRefreshing ? 'در حال بروزرسانی...' : 'بروزرسانی' }}
        </button>
        <button
          @click="exportPerformanceReport"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          صادر کردن گزارش
        </button>
      </div>
    </div>

    <!-- متریک‌های کلیدی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ metrics.uptime }}%</div>
            <div class="text-sm text-blue-700">زمان فعالیت</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ metrics.responseTime }}ms</div>
            <div class="text-sm text-green-700">زمان پاسخ</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ metrics.throughput }}</div>
            <div class="text-sm text-yellow-700">تراکنش/دقیقه</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ metrics.successRate }}%</div>
            <div class="text-sm text-purple-700">نرخ موفقیت</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای عملکرد -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">زمان پاسخ در ۲۴ ساعت گذشته</h5>
        <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
          <div
            v-for="(point, index) in responseTimeChart"
            :key="index"
            class="flex-1 bg-blue-500 rounded-t"
            :style="{ height: (point.value / 1000) * 100 + '%' }"
            :title="`${point.time}: ${point.value}ms`"
          ></div>
        </div>
        <div class="mt-4 text-xs text-gray-500 text-center">
          زمان (ساعت)
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">نرخ موفقیت در ۷ روز گذشته</h5>
        <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
          <div
            v-for="(point, index) in successRateChart"
            :key="index"
            class="flex-1 bg-green-500 rounded-t"
            :style="{ height: point.value + '%' }"
            :title="`${point.date}: ${point.value}%`"
          ></div>
        </div>
        <div class="mt-4 text-xs text-gray-500 text-center">
          روز
        </div>
      </div>
    </div>

    <!-- جدول عملکرد نرم‌افزارها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">عملکرد نرم‌افزارهای حسابداری</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">زمان پاسخ</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرخ موفقیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تراکنش/دقیقه</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">آخرین همگام</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="software in softwarePerformance" :key="software.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <img :src="software.logo" :alt="software.name" class="w-8 h-8 rounded" />
                  <div>
                    <div class="font-medium text-gray-900">{{ software.name }}</div>
                    <div class="text-xs text-gray-500">{{ software.version }}</div>
                  </div>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div
                    class="w-2 h-2 rounded-full"
                    :class="getStatusColor(software.status)"
                  ></div>
                  <span
                    class="px-2 py-1 rounded-full text-xs font-medium"
                    :class="getStatusClass(software.status)"
                  >
                    {{ getStatusLabel(software.status) }}
                  </span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ software.responseTime }}ms</div>
                <div class="text-xs text-gray-500">{{ software.responseStatus }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div class="flex-1 bg-gray-200 rounded-full h-2">
                    <div
                      class="h-2 rounded-full"
                      :class="getSuccessRateColor(software.successRate)"
                      :style="{ width: software.successRate + '%' }"
                    ></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ software.successRate }}%</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ software.throughput }}</div>
                <div class="text-xs text-gray-500">تراکنش/دقیقه</div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ software.lastSync }}</div>
                <div class="text-xs text-gray-500">{{ software.lastSyncTime }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    @click="testPerformance(software)"
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="تست عملکرد"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                    </svg>
                  </button>
                  <button
                    @click="viewPerformanceDetails(software)"
                    class="p-1 text-green-600 hover:text-green-800"
                    title="مشاهده جزئیات"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button
                    @click="optimizePerformance(software)"
                    class="p-1 text-yellow-600 hover:text-yellow-800"
                    title="بهینه‌سازی"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- هشدارهای عملکرد -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">هشدارهای عملکرد</h5>
      </div>
      <div class="p-6">
        <div v-if="performanceAlerts.length === 0" class="text-center text-gray-500 py-8">
          <svg class="w-12 h-12 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p>هیچ هشداری وجود ندارد</p>
        </div>
        <div v-else class="space-y-4">
          <div
            v-for="alert in performanceAlerts"
            :key="alert.id"
            class="flex items-start space-x-4 space-x-reverse p-6 rounded-lg"
            :class="getAlertClass(alert.level)"
          >
            <div class="flex-shrink-0">
              <svg class="w-5 h-5" :class="getAlertIconClass(alert.level)" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
              </svg>
            </div>
            <div class="flex-1">
              <h6 class="text-sm font-medium" :class="getAlertTextClass(alert.level)">
                {{ alert.title }}
              </h6>
              <p class="text-sm mt-1" :class="getAlertTextClass(alert.level)">
                {{ alert.message }}
              </p>
              <div class="mt-2 text-xs" :class="getAlertTextClass(alert.level)">
                {{ alert.time }}
              </div>
            </div>
            <div class="flex-shrink-0">
              <button
                @click="dismissAlert(alert)"
                class="text-gray-400 hover:text-gray-600"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت بروزرسانی
const isRefreshing = ref(false)

// متریک‌های کلیدی
const metrics = ref({
  uptime: 99.8,
  responseTime: 245,
  throughput: 1250,
  successRate: 96.3
})

// نمودار زمان پاسخ
const responseTimeChart = ref([
  { time: '00:00', value: 180 },
  { time: '04:00', value: 220 },
  { time: '08:00', value: 280 },
  { time: '12:00', value: 320 },
  { time: '16:00', value: 290 },
  { time: '20:00', value: 250 },
  { time: '24:00', value: 200 }
])

// نمودار نرخ موفقیت
const successRateChart = ref([
  { date: 'شنبه', value: 98 },
  { date: 'یکشنبه', value: 97 },
  { date: 'دوشنبه', value: 99 },
  { date: 'سه‌شنبه', value: 96 },
  { date: 'چهارشنبه', value: 98 },
  { date: 'پنج‌شنبه', value: 95 },
  { date: 'جمعه', value: 97 }
])

// عملکرد نرم‌افزارها
const softwarePerformance = ref([
  {
    id: 1,
    name: 'هلو',
    version: 'نسخه ۹.۲',
    logo: '/images/helo-logo.png',
    status: 'active',
    responseTime: 180,
    responseStatus: 'عالی',
    successRate: 98.5,
    throughput: 1500,
    lastSync: 'امروز ۱۴:۳۰',
    lastSyncTime: '۲ دقیقه پیش'
  },
  {
    id: 2,
    name: 'پارسیان',
    version: 'نسخه ۸.۱',
    logo: '/images/parsian-logo.png',
    status: 'active',
    responseTime: 220,
    responseStatus: 'خوب',
    successRate: 99.2,
    throughput: 1200,
    lastSync: 'امروز ۱۴:۲۵',
    lastSyncTime: '۷ دقیقه پیش'
  },
  {
    id: 3,
    name: 'سپیدار',
    version: 'نسخه ۷.۵',
    logo: '/images/sepidar-logo.png',
    status: 'warning',
    responseTime: 1250,
    responseStatus: 'کند',
    successRate: 85.3,
    throughput: 800,
    lastSync: 'امروز ۱۳:۴۵',
    lastSyncTime: '۴۷ دقیقه پیش'
  }
])

// هشدارهای عملکرد
const performanceAlerts = ref([
  {
    id: 1,
    level: 'warning',
    title: 'زمان پاسخ کند',
    message: 'زمان پاسخ نرم‌افزار سپیدار بیش از حد معمول است',
    time: 'امروز ۱۳:۴۵'
  },
  {
    id: 2,
    level: 'error',
    title: 'اتصال ناموفق',
    message: 'اتصال به نرم‌افزار قیاس ناموفق بود',
    time: 'امروز ۱۲:۳۰'
  }
])

// رنگ وضعیت
const getStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-700',
    warning: 'bg-yellow-100 text-yellow-700',
    error: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    warning: 'هشدار',
    error: 'خطا'
  }
  return labels[status] || status
}

// رنگ نرخ موفقیت
const getSuccessRateColor = (rate: number) => {
  if (rate >= 95) return 'bg-green-500'
  if (rate >= 80) return 'bg-yellow-500'
  return 'bg-red-500'
}

// کلاس هشدار
const getAlertClass = (level: string) => {
  const classes = {
    warning: 'bg-yellow-50 border border-yellow-200',
    error: 'bg-red-50 border border-red-200',
    info: 'bg-blue-50 border border-blue-200'
  }
  return classes[level] || 'bg-gray-50 border border-gray-200'
}

// کلاس آیکون هشدار
const getAlertIconClass = (level: string) => {
  const classes = {
    warning: 'text-yellow-400',
    error: 'text-red-400',
    info: 'text-blue-400'
  }
  return classes[level] || 'text-gray-400'
}

// کلاس متن هشدار
const getAlertTextClass = (level: string) => {
  const classes = {
    warning: 'text-yellow-800',
    error: 'text-red-800',
    info: 'text-blue-800'
  }
  return classes[level] || 'text-gray-800'
}

// بروزرسانی متریک‌ها
const refreshMetrics = async () => {
  try {
    isRefreshing.value = true
    // TODO: بروزرسانی متریک‌ها از سرور
    console.log('بروزرسانی متریک‌های عملکرد')
    
    // شبیه‌سازی بروزرسانی
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    console.log('متریک‌های عملکرد بروزرسانی شد')
  } catch (error) {
    console.error('خطا در بروزرسانی متریک‌ها:', error)
  } finally {
    isRefreshing.value = false
  }
}

// صادر کردن گزارش عملکرد
const exportPerformanceReport = () => {
  // TODO: صادر کردن گزارش عملکرد
  console.log('صادر کردن گزارش عملکرد')
}

// تست عملکرد
const testPerformance = (software: any) => {
  // TODO: تست عملکرد نرم‌افزار
  console.log('تست عملکرد:', software)
}

// مشاهده جزئیات عملکرد
const viewPerformanceDetails = (software: any) => {
  // TODO: مشاهده جزئیات عملکرد
  console.log('مشاهده جزئیات عملکرد:', software)
}

// بهینه‌سازی عملکرد
const optimizePerformance = (software: any) => {
  // TODO: بهینه‌سازی عملکرد
  console.log('بهینه‌سازی عملکرد:', software)
}

// حذف هشدار
const dismissAlert = (alert: any) => {
  // TODO: حذف هشدار
  console.log('حذف هشدار:', alert)
  const index = performanceAlerts.value.findIndex(a => a.id === alert.id)
  if (index > -1) {
    performanceAlerts.value.splice(index, 1)
  }
}
</script>

<!--
  کامپوننت نظارت بر عملکرد اتصال
  شامل:
  1. متریک‌های کلیدی عملکرد
  2. نمودارهای تحلیلی
  3. جدول عملکرد نرم‌افزارها
  4. هشدارهای عملکرد
  5. عملیات مدیریت عملکرد
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
