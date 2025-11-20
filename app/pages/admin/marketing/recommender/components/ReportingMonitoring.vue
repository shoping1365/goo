<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">گزارش‌گیری و مانیتورینگ</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت گزارش‌ها و نظارت بر عملکرد سیستم</p>
      </div>
      <div class="flex space-x-3 space-x-reverse">
        <select class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          <option>امروز</option>
          <option>هفته گذشته</option>
          <option>ماه گذشته</option>
          <option>3 ماه گذشته</option>
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

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">کل گزارش‌ها</p>
            <p class="text-2xl font-bold">1,247</p>
            <p class="text-xs opacity-75">+23 از هفته گذشته</p>
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
            <p class="text-sm opacity-90">گزارش‌های فعال</p>
            <p class="text-2xl font-bold">89</p>
            <p class="text-xs opacity-75">7.1% از کل گزارش‌ها</p>
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
            <p class="text-sm opacity-90">هشدارهای فعال</p>
            <p class="text-2xl font-bold">12</p>
            <p class="text-xs opacity-75">-3 از دیروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">میانگین زمان تولید</p>
            <p class="text-2xl font-bold">2.3s</p>
            <p class="text-xs opacity-75">-0.5s از هفته گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش‌های موجود -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌های موجود</h3>
        <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">ایجاد گزارش جدید</button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام گزارش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فرکانس</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین اجرا</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in availableReports" :key="report.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="p-2 rounded-lg mr-3" :class="getReportColor(report.type).bg">
                    <svg class="w-5 h-5" :class="getReportColor(report.type).text" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getReportIcon(report.type)"></path>
                    </svg>
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ report.name }}</div>
                    <div class="text-sm text-gray-500">{{ report.description }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="getReportColor(report.type).badge">
                  {{ report.type }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ report.frequency }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ report.lastRun }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="report.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                  {{ report.status === 'active' ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-4">اجرا</button>
                <button class="text-green-600 hover:text-green-900 ml-4">دانلود</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مانیتورینگ سیستم -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- وضعیت سیستم -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">وضعیت سیستم</h3>
        
        <div class="space-y-4">
          <div v-for="service in systemServices" :key="service.id" class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
            <div class="flex items-center">
              <div class="w-3 h-3 rounded-full mr-3" :class="service.status === 'healthy' ? 'bg-green-500' : service.status === 'warning' ? 'bg-yellow-500' : 'bg-red-500'"></div>
              <div>
                <p class="text-sm font-medium text-gray-900">{{ service.name }}</p>
                <p class="text-xs text-gray-500">{{ service.description }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ service.responseTime }}ms</p>
              <p class="text-xs text-gray-500">{{ service.uptime }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- هشدارها -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">هشدارهای فعال</h3>
        
        <div class="space-y-4">
          <div v-for="alert in activeAlerts" :key="alert.id" class="border border-gray-200 rounded-lg p-6" :class="getAlertColor(alert.severity).bg">
            <div class="flex items-start">
              <div class="p-2 rounded-lg mr-3" :class="getAlertColor(alert.severity).icon">
                <svg class="w-5 h-5" :class="getAlertColor(alert.severity).text" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                </svg>
              </div>
              <div class="flex-1">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="text-sm font-medium" :class="getAlertColor(alert.severity).title">{{ alert.title }}</h4>
                  <span class="text-xs" :class="getAlertColor(alert.severity).time">{{ alert.time }}</span>
                </div>
                <p class="text-sm" :class="getAlertColor(alert.severity).description">{{ alert.description }}</p>
                <div class="mt-3 flex space-x-2 space-x-reverse">
                  <button class="text-xs px-2 py-1 rounded border" :class="getAlertColor(alert.severity).button">بررسی</button>
                  <button class="text-xs px-2 py-1 rounded border" :class="getAlertColor(alert.severity).button">نادیده گرفتن</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای عملکرد -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- نمودار درخواست‌ها -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">درخواست‌های سیستم</h3>
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-3 py-1 text-xs bg-blue-100 text-blue-800 rounded-lg">ساعت</button>
            <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">روز</button>
            <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">هفته</button>
          </div>
        </div>
        
        <div class="space-y-4">
          <div v-for="(hour, index) in requestData" :key="index" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-12 text-sm text-gray-600">{{ hour.time }}</div>
              <div class="mr-4">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div class="bg-blue-600 h-2 rounded-full" :style="{ width: (hour.requests / 1000) + '%' }"></div>
                </div>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-medium text-gray-900">{{ hour.requests }}</p>
              <p class="text-xs text-gray-500" :class="hour.trend > 0 ? 'text-green-600' : 'text-red-600'">
                {{ hour.trend > 0 ? '+' : '' }}{{ hour.trend }}%
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار خطاها -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">خطاهای سیستم</h3>
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-3 py-1 text-xs bg-red-100 text-red-800 rounded-lg">خطا</button>
            <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">هشدار</button>
          </div>
        </div>
        
        <div class="space-y-4">
          <div v-for="(error, index) in errorData" :key="index" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-12 text-sm text-gray-600">{{ error.time }}</div>
              <div class="mr-4">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div class="bg-red-600 h-2 rounded-full" :style="{ width: (error.count / 10) + '%' }"></div>
                </div>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-medium text-gray-900">{{ error.count }}</p>
              <p class="text-xs text-gray-500">{{ error.type }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش‌های خودکار -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌های خودکار</h3>
        <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">تنظیم زمان‌بندی</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="schedule in reportSchedules" :key="schedule.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <h4 class="text-sm font-medium text-gray-900">{{ schedule.name }}</h4>
            <label class="relative inline-flex items-center cursor-pointer">
              <input v-model="schedule.enabled" type="checkbox" class="sr-only peer">
              <div class="w-8 h-4 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-3 after:w-3 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="space-y-2">
            <div class="flex justify-between text-xs">
              <span class="text-gray-600">فرکانس:</span>
              <span class="font-medium">{{ schedule.frequency }}</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-600">آخرین اجرا:</span>
              <span class="font-medium">{{ schedule.lastRun }}</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-600">اجرای بعدی:</span>
              <span class="font-medium">{{ schedule.nextRun }}</span>
            </div>
          </div>
          
          <div class="mt-3 flex space-x-2 space-x-reverse">
            <button class="text-xs px-2 py-1 bg-blue-100 text-blue-700 rounded hover:bg-blue-200">ویرایش</button>
            <button class="text-xs px-2 py-1 bg-gray-100 text-gray-700 rounded hover:bg-gray-200">اجرا</button>
          </div>
        </div>
      </div>
    </div>

    <!-- لاگ‌های سیستم -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">لاگ‌های سیستم</h3>
        <div class="flex space-x-2 space-x-reverse">
          <select class="px-3 py-1 text-xs border border-gray-300 rounded-lg">
            <option>همه سطوح</option>
            <option>خطا</option>
            <option>هشدار</option>
            <option>اطلاعات</option>
          </select>
          <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">پاک کردن لاگ</button>
        </div>
      </div>
      
      <div class="space-y-2 max-h-64 overflow-y-auto">
        <div v-for="log in systemLogs" :key="log.id" class="flex items-start p-3 border border-gray-200 rounded-lg">
          <div class="p-1 rounded mr-3" :class="getLogColor(log.level).bg">
            <svg class="w-4 h-4" :class="getLogColor(log.level).text" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getLogIcon(log.level)"></path>
            </svg>
          </div>
          <div class="flex-1">
            <div class="flex items-center justify-between mb-1">
              <span class="text-sm font-medium text-gray-900">{{ log.message }}</span>
              <span class="text-xs text-gray-500">{{ log.time }}</span>
            </div>
            <p class="text-xs text-gray-600">{{ log.details }}</p>
            <div class="mt-1 flex items-center space-x-4 space-x-reverse text-xs text-gray-500">
              <span>{{ log.service }}</span>
              <span>{{ log.user }}</span>
              <span>{{ log.ip }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// گزارش‌های موجود
const availableReports = ref([
  {
    id: 1,
    name: 'گزارش عملکرد روزانه',
    description: 'آمار عملکرد سیستم در 24 ساعت گذشته',
    type: 'performance',
    frequency: 'روزانه',
    lastRun: '2 ساعت پیش',
    status: 'active'
  },
  {
    id: 2,
    name: 'گزارش خطاها',
    description: 'لیست خطاهای رخ داده در سیستم',
    type: 'error',
    frequency: 'ساعتی',
    lastRun: '30 دقیقه پیش',
    status: 'active'
  },
  {
    id: 3,
    name: 'گزارش کاربران',
    description: 'آمار فعالیت کاربران و پیشنهادات',
    type: 'user',
    frequency: 'هفتگی',
    lastRun: '2 روز پیش',
    status: 'active'
  },
  {
    id: 4,
    name: 'گزارش امنیت',
    description: 'لاگ‌های امنیتی و دسترسی‌ها',
    type: 'security',
    frequency: 'روزانه',
    lastRun: '1 روز پیش',
    status: 'inactive'
  },
  {
    id: 5,
    name: 'گزارش مالی',
    description: 'آمار درآمد و فروش',
    type: 'financial',
    frequency: 'هفتگی',
    lastRun: '3 روز پیش',
    status: 'active'
  }
])

// سرویس‌های سیستم
const systemServices = ref([
  {
    id: 1,
    name: 'API Gateway',
    description: 'درگاه اصلی API',
    status: 'healthy',
    responseTime: 45,
    uptime: 99.9
  },
  {
    id: 2,
    name: 'Database',
    description: 'پایگاه داده اصلی',
    status: 'healthy',
    responseTime: 12,
    uptime: 99.8
  },
  {
    id: 3,
    name: 'Cache Server',
    description: 'سرور کش Redis',
    status: 'warning',
    responseTime: 8,
    uptime: 98.5
  },
  {
    id: 4,
    name: 'ML Engine',
    description: 'موتور یادگیری ماشین',
    status: 'healthy',
    responseTime: 156,
    uptime: 99.2
  }
])

// هشدارهای فعال
const activeAlerts = ref([
  {
    id: 1,
    severity: 'high',
    title: 'مصرف CPU بالا',
    description: 'مصرف CPU به 90% رسیده است',
    time: '5 دقیقه پیش'
  },
  {
    id: 2,
    severity: 'medium',
    title: 'خطای اتصال به کش',
    description: 'اتصال به Redis با مشکل مواجه شده',
    time: '15 دقیقه پیش'
  },
  {
    id: 3,
    severity: 'low',
    title: 'فضای دیسک کم',
    description: 'فقط 15% فضای دیسک باقی مانده',
    time: '1 ساعت پیش'
  }
])

// داده‌های درخواست
const requestData = ref([
  { time: '00:00', requests: 1247, trend: 5.2 },
  { time: '04:00', requests: 892, trend: -2.1 },
  { time: '08:00', requests: 2156, trend: 12.3 },
  { time: '12:00', requests: 3456, trend: 8.7 },
  { time: '16:00', requests: 2891, trend: 3.4 },
  { time: '20:00', requests: 1678, trend: -1.2 }
])

// داده‌های خطا
const errorData = ref([
  { time: '00:00', count: 3, type: 'خطای شبکه' },
  { time: '04:00', count: 1, type: 'خطای دیتابیس' },
  { time: '08:00', count: 7, type: 'خطای API' },
  { time: '12:00', count: 5, type: 'خطای کش' },
  { time: '16:00', count: 2, type: 'خطای احراز هویت' },
  { time: '20:00', count: 4, type: 'خطای شبکه' }
])

// زمان‌بندی گزارش‌ها
const reportSchedules = ref([
  {
    id: 1,
    name: 'گزارش روزانه عملکرد',
    enabled: true,
    frequency: 'هر روز ساعت 08:00',
    lastRun: 'امروز 08:00',
    nextRun: 'فردا 08:00'
  },
  {
    id: 2,
    name: 'گزارش هفتگی کاربران',
    enabled: true,
    frequency: 'هر شنبه ساعت 09:00',
    lastRun: 'شنبه گذشته',
    nextRun: 'شنبه آینده'
  },
  {
    id: 3,
    name: 'گزارش ماهانه مالی',
    enabled: false,
    frequency: 'اول هر ماه',
    lastRun: '1 ماه پیش',
    nextRun: '1 ماه آینده'
  },
  {
    id: 4,
    name: 'گزارش خطاها',
    enabled: true,
    frequency: 'هر ساعت',
    lastRun: '1 ساعت پیش',
    nextRun: '1 ساعت آینده'
  },
  {
    id: 5,
    name: 'گزارش امنیت',
    enabled: true,
    frequency: 'هر 6 ساعت',
    lastRun: '6 ساعت پیش',
    nextRun: '6 ساعت آینده'
  },
  {
    id: 6,
    name: 'گزارش پشتیبان‌گیری',
    enabled: false,
    frequency: 'هر روز ساعت 02:00',
    lastRun: 'دیروز',
    nextRun: 'امشب'
  }
])

// لاگ‌های سیستم
const systemLogs = ref([
  {
    id: 1,
    level: 'error',
    message: 'خطای اتصال به پایگاه داده',
    details: 'Connection timeout after 30 seconds',
    time: '2 دقیقه پیش',
    service: 'Database',
    user: 'system',
    ip: '127.0.0.1'
  },
  {
    id: 2,
    level: 'warning',
    message: 'مصرف حافظه بالا',
    details: 'Memory usage reached 85%',
    time: '5 دقیقه پیش',
    service: 'System',
    user: 'monitor',
    ip: '192.168.1.100'
  },
  {
    id: 3,
    level: 'info',
    message: 'گزارش روزانه تولید شد',
    details: 'Daily performance report generated successfully',
    time: '1 ساعت پیش',
    service: 'Reporting',
    user: 'scheduler',
    ip: '10.0.0.1'
  },
  {
    id: 4,
    level: 'error',
    message: 'خطای احراز هویت کاربر',
    details: 'Invalid token provided',
    time: '2 ساعت پیش',
    service: 'Auth',
    user: 'user123',
    ip: '203.0.113.45'
  }
])

// توابع کمکی
const getReportColor = (type: string) => {
  const colors = {
    performance: { bg: 'bg-blue-100', text: 'text-blue-600', badge: 'bg-blue-100 text-blue-800' },
    error: { bg: 'bg-red-100', text: 'text-red-600', badge: 'bg-red-100 text-red-800' },
    user: { bg: 'bg-green-100', text: 'text-green-600', badge: 'bg-green-100 text-green-800' },
    security: { bg: 'bg-yellow-100', text: 'text-yellow-600', badge: 'bg-yellow-100 text-yellow-800' },
    financial: { bg: 'bg-purple-100', text: 'text-purple-600', badge: 'bg-purple-100 text-purple-800' }
  }
  return colors[type] || colors.performance
}

const getReportIcon = (type: string) => {
  const icons = {
    performance: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z',
    error: 'M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    user: 'M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z',
    security: 'M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z',
    financial: 'M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1'
  }
  return icons[type] || icons.performance
}

const getAlertColor = (severity: string) => {
  const colors = {
    high: {
      bg: 'bg-red-50',
      icon: 'bg-red-100',
      text: 'text-red-600',
      title: 'text-red-800',
      description: 'text-red-700',
      time: 'text-red-600',
      button: 'border-red-300 text-red-700 hover:bg-red-100'
    },
    medium: {
      bg: 'bg-yellow-50',
      icon: 'bg-yellow-100',
      text: 'text-yellow-600',
      title: 'text-yellow-800',
      description: 'text-yellow-700',
      time: 'text-yellow-600',
      button: 'border-yellow-300 text-yellow-700 hover:bg-yellow-100'
    },
    low: {
      bg: 'bg-blue-50',
      icon: 'bg-blue-100',
      text: 'text-blue-600',
      title: 'text-blue-800',
      description: 'text-blue-700',
      time: 'text-blue-600',
      button: 'border-blue-300 text-blue-700 hover:bg-blue-100'
    }
  }
  return colors[severity] || colors.low
}

const getLogColor = (level: string) => {
  const colors = {
    error: { bg: 'bg-red-100', text: 'text-red-600' },
    warning: { bg: 'bg-yellow-100', text: 'text-yellow-600' },
    info: { bg: 'bg-blue-100', text: 'text-blue-600' },
    debug: { bg: 'bg-gray-100', text: 'text-gray-600' }
  }
  return colors[level] || colors.info
}

const getLogIcon = (level: string) => {
  const icons = {
    error: 'M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    warning: 'M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z',
    info: 'M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    debug: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[level] || icons.info
}
</script> 
