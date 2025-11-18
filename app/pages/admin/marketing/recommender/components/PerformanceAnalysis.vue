<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">تحلیل عملکرد</h2>
        <p class="mt-1 text-sm text-gray-600">بررسی و تحلیل عملکرد سیستم پیشنهادات هوشمند</p>
      </div>
      <div class="flex space-x-3 space-x-reverse">
        <select class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          <option>7 روز گذشته</option>
          <option>30 روز گذشته</option>
          <option>90 روز گذشته</option>
          <option>1 سال گذشته</option>
        </select>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"></path>
          </svg>
          خروجی گزارش
        </button>
      </div>
    </div>

    <!-- آمار کلی عملکرد -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">میانگین دقت</p>
            <p class="text-2xl font-bold">87.3%</p>
            <p class="text-xs opacity-75">+2.1% از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">میانگین زمان پاسخ</p>
            <p class="text-2xl font-bold">45ms</p>
            <p class="text-xs opacity-75">-12ms از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">نرخ موفقیت</p>
            <p class="text-2xl font-bold">94.7%</p>
            <p class="text-xs opacity-75">+1.8% از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">درخواست‌های روزانه</p>
            <p class="text-2xl font-bold">1.2M</p>
            <p class="text-xs opacity-75">+15.3% از هفته گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای عملکرد -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- نمودار دقت در طول زمان -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">دقت در طول زمان</h3>
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-3 py-1 text-xs bg-blue-100 text-blue-800 rounded-lg">روزانه</button>
            <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">هفتگی</button>
            <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">ماهانه</button>
          </div>
        </div>
        
        <div class="space-y-4">
          <div v-for="(day, index) in accuracyData" :key="index" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-16 text-sm text-gray-600">{{ day.date }}</div>
              <div class="mr-4">
                <div class="w-48 bg-gray-200 rounded-full h-3">
                  <div class="bg-blue-600 h-3 rounded-full" :style="{ width: day.accuracy + '%' }"></div>
                </div>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-medium text-gray-900">{{ day.accuracy }}%</p>
              <p class="text-xs text-gray-500" :class="day.trend > 0 ? 'text-green-600' : 'text-red-600'">
                {{ day.trend > 0 ? '+' : '' }}{{ day.trend }}%
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار زمان پاسخ -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">زمان پاسخ</h3>
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-3 py-1 text-xs bg-green-100 text-green-800 rounded-lg">میلی‌ثانیه</button>
            <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">ثانیه</button>
          </div>
        </div>
        
        <div class="space-y-4">
          <div v-for="(time, index) in responseTimeData" :key="index" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-16 text-sm text-gray-600">{{ time.time }}</div>
              <div class="mr-4">
                <div class="w-48 bg-gray-200 rounded-full h-3">
                  <div class="bg-green-600 h-3 rounded-full" :style="{ width: (time.responseTime / 100) + '%' }"></div>
                </div>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-medium text-gray-900">{{ time.responseTime }}ms</p>
              <p class="text-xs text-gray-500" :class="time.trend < 0 ? 'text-green-600' : 'text-red-600'">
                {{ time.trend > 0 ? '+' : '' }}{{ time.trend }}ms
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل الگوریتم‌ها -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">عملکرد الگوریتم‌ها</h3>
        <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">مشاهده جزئیات</button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">الگوریتم</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دقت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان پاسخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ موفقیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درخواست‌ها</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="algorithm in algorithmPerformance" :key="algorithm.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="p-2 rounded-lg mr-3" :class="getAlgorithmColor(algorithm.type).bg">
                    <svg class="w-5 h-5" :class="getAlgorithmColor(algorithm.type).text" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getAlgorithmIcon(algorithm.type)"></path>
                    </svg>
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ algorithm.name }}</div>
                    <div class="text-sm text-gray-500">{{ algorithm.description }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div class="bg-blue-600 h-2 rounded-full" :style="{ width: algorithm.accuracy + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ algorithm.accuracy }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ algorithm.responseTime }}ms</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div class="bg-green-600 h-2 rounded-full" :style="{ width: algorithm.successRate + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ algorithm.successRate }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ algorithm.requests.toLocaleString() }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="algorithm.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                  {{ algorithm.status === 'active' ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- تحلیل خطاها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- انواع خطاها -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">انواع خطاها</h3>
        
        <div class="space-y-4">
          <div v-for="error in errorTypes" :key="error.id" class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
            <div class="flex items-center">
              <div class="p-2 rounded-lg mr-3" :class="getErrorColor(error.severity).bg">
                <svg class="w-5 h-5" :class="getErrorColor(error.severity).text" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                </svg>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-900">{{ error.name }}</p>
                <p class="text-xs text-gray-500">{{ error.description }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ error.count }}</p>
              <p class="text-xs text-gray-500">{{ error.percentage }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- خطاهای اخیر -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">خطاهای اخیر</h3>
        
        <div class="space-y-4 max-h-80 overflow-y-auto">
          <div v-for="error in recentErrors" :key="error.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-start justify-between">
              <div class="flex-1">
                <div class="flex items-center mb-2">
                  <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full mr-2" :class="getErrorColor(error.severity).badge">
                    {{ error.severity }}
                  </span>
                  <span class="text-xs text-gray-500">{{ error.time }}</span>
                </div>
                <p class="text-sm font-medium text-gray-900 mb-1">{{ error.message }}</p>
                <p class="text-xs text-gray-500">{{ error.details }}</p>
              </div>
              <button class="text-blue-600 hover:text-blue-800 text-sm">بررسی</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل منابع -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تحلیل منابع سیستم</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-xs bg-purple-100 text-purple-800 rounded-lg">CPU</button>
          <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">RAM</button>
          <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">Disk</button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div v-for="resource in systemResources" :key="resource.id" class="text-center">
          <div class="relative w-24 h-24 mx-auto mb-4">
            <svg class="w-24 h-24 transform -rotate-90" viewBox="0 0 36 36">
              <path class="text-gray-200" stroke="currentColor" stroke-width="3" fill="none" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
              <path class="text-blue-600" stroke="currentColor" stroke-width="3" fill="none" stroke-dasharray="100, 100" stroke-dashoffset="25" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"/>
            </svg>
            <div class="absolute inset-0 flex items-center justify-center">
              <span class="text-lg font-bold text-gray-900">{{ resource.usage }}%</span>
            </div>
          </div>
          <h4 class="text-sm font-medium text-gray-900 mb-1">{{ resource.name }}</h4>
          <p class="text-xs text-gray-500">{{ resource.current }}/{{ resource.total }}</p>
          <p class="text-xs" :class="resource.trend > 0 ? 'text-red-600' : 'text-green-600'">
            {{ resource.trend > 0 ? '+' : '' }}{{ resource.trend }}% از ساعت گذشته
          </p>
        </div>
      </div>
    </div>

    <!-- هشدارها و توصیه‌ها -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-6">هشدارها و توصیه‌ها</h3>
      
      <div class="space-y-4">
        <div v-for="alert in alerts" :key="alert.id" class="flex items-start p-6 rounded-lg" :class="getAlertColor(alert.type).bg">
          <div class="p-2 rounded-lg mr-3" :class="getAlertColor(alert.type).icon">
            <svg class="w-5 h-5" :class="getAlertColor(alert.type).text" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getAlertIcon(alert.type)"></path>
            </svg>
          </div>
          <div class="flex-1">
            <h4 class="text-sm font-medium" :class="getAlertColor(alert.type).title">{{ alert.title }}</h4>
            <p class="text-sm mt-1" :class="getAlertColor(alert.type).description">{{ alert.description }}</p>
            <div class="mt-2 flex space-x-2 space-x-reverse">
              <button class="text-xs px-2 py-1 rounded border" :class="getAlertColor(alert.type).button">بررسی</button>
              <button class="text-xs px-2 py-1 rounded border" :class="getAlertColor(alert.type).button">نادیده گرفتن</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// داده‌های دقت
const accuracyData = ref([
  { date: 'دوشنبه', accuracy: 85.2, trend: 2.1 },
  { date: 'سه‌شنبه', accuracy: 87.8, trend: 1.5 },
  { date: 'چهارشنبه', accuracy: 89.1, trend: 3.2 },
  { date: 'پنج‌شنبه', accuracy: 86.7, trend: -1.8 },
  { date: 'جمعه', accuracy: 88.9, trend: 2.4 },
  { date: 'شنبه', accuracy: 90.2, trend: 4.1 },
  { date: 'یکشنبه', accuracy: 87.3, trend: -0.5 }
])

// داده‌های زمان پاسخ
const responseTimeData = ref([
  { time: '00:00', responseTime: 42, trend: -3 },
  { time: '04:00', responseTime: 38, trend: -5 },
  { time: '08:00', responseTime: 45, trend: 2 },
  { time: '12:00', responseTime: 52, trend: 8 },
  { time: '16:00', responseTime: 48, trend: 3 },
  { time: '20:00', responseTime: 41, trend: -4 }
])

// عملکرد الگوریتم‌ها
const algorithmPerformance = ref([
  {
    id: 1,
    name: 'Collaborative Filtering',
    description: 'پیشنهاد بر اساس رفتار کاربران مشابه',
    type: 'collaborative',
    accuracy: 82.5,
    responseTime: 45,
    successRate: 94.2,
    requests: 456789,
    status: 'active'
  },
  {
    id: 2,
    name: 'Content-Based Filtering',
    description: 'پیشنهاد بر اساس ویژگی‌های محصول',
    type: 'content',
    accuracy: 78.3,
    responseTime: 32,
    successRate: 91.7,
    requests: 234567,
    status: 'active'
  },
  {
    id: 3,
    name: 'Hybrid Approach',
    description: 'ترکیب الگوریتم‌های مختلف',
    type: 'hybrid',
    accuracy: 87.3,
    responseTime: 52,
    successRate: 96.8,
    requests: 678901,
    status: 'active'
  },
  {
    id: 4,
    name: 'Real-time Learning',
    description: 'یادگیری مداوم از تعاملات',
    type: 'realtime',
    accuracy: 85.1,
    responseTime: 38,
    successRate: 93.5,
    requests: 345678,
    status: 'active'
  }
])

// انواع خطاها
const errorTypes = ref([
  {
    id: 1,
    name: 'خطای شبکه',
    description: 'مشکلات اتصال به سرور',
    severity: 'high',
    count: 156,
    percentage: 45.2
  },
  {
    id: 2,
    name: 'خطای حافظه',
    description: 'مشکلات مربوط به RAM',
    severity: 'medium',
    count: 89,
    percentage: 25.8
  },
  {
    id: 3,
    name: 'خطای پایگاه داده',
    description: 'مشکلات دسترسی به دیتابیس',
    severity: 'high',
    count: 67,
    percentage: 19.4
  },
  {
    id: 4,
    name: 'خطای الگوریتم',
    description: 'مشکلات در پردازش الگوریتم',
    severity: 'low',
    count: 33,
    percentage: 9.6
  }
])

// خطاهای اخیر
const recentErrors = ref([
  {
    id: 1,
    severity: 'high',
    time: '2 دقیقه پیش',
    message: 'خطای اتصال به سرور Redis',
    details: 'Connection timeout after 30 seconds'
  },
  {
    id: 2,
    severity: 'medium',
    time: '5 دقیقه پیش',
    message: 'مصرف حافظه بالا',
    details: 'Memory usage reached 85%'
  },
  {
    id: 3,
    severity: 'low',
    time: '8 دقیقه پیش',
    message: 'خطای کوچک در الگوریتم',
    details: 'Minor calculation error in collaborative filtering'
  },
  {
    id: 4,
    severity: 'high',
    time: '12 دقیقه پیش',
    message: 'خطای پایگاه داده',
    details: 'Database connection pool exhausted'
  }
])

// منابع سیستم
const systemResources = ref([
  {
    id: 1,
    name: 'CPU',
    usage: 75,
    current: '75%',
    total: '100%',
    trend: 5
  },
  {
    id: 2,
    name: 'RAM',
    usage: 68,
    current: '13.6GB',
    total: '20GB',
    trend: -2
  },
  {
    id: 3,
    name: 'Disk',
    usage: 45,
    current: '450GB',
    total: '1TB',
    trend: 1
  }
])

// هشدارها
const alerts = ref([
  {
    id: 1,
    type: 'warning',
    title: 'مصرف CPU بالا',
    description: 'مصرف CPU به 85% رسیده است. توصیه می‌شود منابع بیشتری تخصیص دهید.'
  },
  {
    id: 2,
    type: 'error',
    title: 'خطای مکرر شبکه',
    description: 'در 10 دقیقه گذشته 15 خطای شبکه رخ داده است. بررسی اتصال شبکه ضروری است.'
  },
  {
    id: 3,
    type: 'info',
    title: 'به‌روزرسانی الگوریتم',
    description: 'نسخه جدید الگوریتم Collaborative Filtering آماده نصب است.'
  },
  {
    id: 4,
    type: 'success',
    title: 'بهبود عملکرد',
    description: 'عملکرد سیستم 15% بهبود یافته است. تغییرات اعمال شده موثر بوده‌اند.'
  }
])

// توابع کمکی
const getAlgorithmColor = (type: string) => {
  const colors = {
    collaborative: { bg: 'bg-blue-100', text: 'text-blue-600' },
    content: { bg: 'bg-green-100', text: 'text-green-600' },
    hybrid: { bg: 'bg-purple-100', text: 'text-purple-600' },
    realtime: { bg: 'bg-yellow-100', text: 'text-yellow-600' }
  }
  return colors[type] || colors.collaborative
}

const getAlgorithmIcon = (type: string) => {
  const icons = {
    collaborative: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z',
    content: '',
    hybrid: 'M13 10V3L4 14h7v7l9-11h-7z',
    realtime: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[type] || icons.collaborative
}

const getErrorColor = (severity: string) => {
  const colors = {
    high: { bg: 'bg-red-100', text: 'text-red-600', badge: 'bg-red-100 text-red-800' },
    medium: { bg: 'bg-yellow-100', text: 'text-yellow-600', badge: 'bg-yellow-100 text-yellow-800' },
    low: { bg: 'bg-blue-100', text: 'text-blue-600', badge: 'bg-blue-100 text-blue-800' }
  }
  return colors[severity] || colors.low
}

const getAlertColor = (type: string) => {
  const colors = {
    warning: {
      bg: 'bg-yellow-50',
      icon: 'bg-yellow-100',
      text: 'text-yellow-600',
      title: 'text-yellow-800',
      description: 'text-yellow-700',
      button: 'border-yellow-300 text-yellow-700 hover:bg-yellow-100'
    },
    error: {
      bg: 'bg-red-50',
      icon: 'bg-red-100',
      text: 'text-red-600',
      title: 'text-red-800',
      description: 'text-red-700',
      button: 'border-red-300 text-red-700 hover:bg-red-100'
    },
    info: {
      bg: 'bg-blue-50',
      icon: 'bg-blue-100',
      text: 'text-blue-600',
      title: 'text-blue-800',
      description: 'text-blue-700',
      button: 'border-blue-300 text-blue-700 hover:bg-blue-100'
    },
    success: {
      bg: 'bg-green-50',
      icon: 'bg-green-100',
      text: 'text-green-600',
      title: 'text-green-800',
      description: 'text-green-700',
      button: 'border-green-300 text-green-700 hover:bg-green-100'
    }
  }
  return colors[type] || colors.info
}

const getAlertIcon = (type: string) => {
  const icons = {
    warning: 'M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z',
    error: 'M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    info: 'M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    success: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[type] || icons.info
}
</script> 
