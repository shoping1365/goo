<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">مدیریت کیفیت داده</h2>
        <p class="mt-1 text-sm text-gray-600">نظارت و بهبود کیفیت داده‌های پیشنهادات</p>
      </div>
      <button class="bg-teal-600 text-white px-4 py-2 rounded-lg hover:bg-teal-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        اسکن کیفیت
      </button>
    </div>

    <!-- آمار کیفیت داده -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">کیفیت کلی</p>
            <p class="text-2xl font-bold">94.2%</p>
            <p class="text-xs opacity-75">عالی</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">داده‌های کامل</p>
            <p class="text-2xl font-bold">98.7%</p>
            <p class="text-xs opacity-75">از کل رکوردها</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">داده‌های تکراری</p>
            <p class="text-2xl font-bold">1.3%</p>
            <p class="text-xs opacity-75">نیاز به پاکسازی</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">داده‌های نامعتبر</p>
            <p class="text-2xl font-bold">0.8%</p>
            <p class="text-xs opacity-75">نیاز به بررسی</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">آخرین به‌روزرسانی</p>
            <p class="text-2xl font-bold">2h</p>
            <p class="text-xs opacity-75">قبل</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات کیفیت داده -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- قوانین اعتبارسنجی -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">قوانین اعتبارسنجی</h3>
          </div>
          <button class="bg-blue-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-blue-700 transition-colors">
            افزودن قانون
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="rule in validationRules" :key="rule.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <h4 class="font-medium text-gray-900">{{ rule.name }}</h4>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" v-model="rule.enabled" class="sr-only peer">
                <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            <p class="text-sm text-gray-600 mb-2">{{ rule.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>نوع: {{ rule.type }}</span>
              <span>اولویت: {{ rule.priority }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- پاکسازی داده -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پاکسازی داده</h3>
          </div>
          <button class="bg-green-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-green-700 transition-colors">
            شروع پاکسازی
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حذف رکوردهای تکراری</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="dataCleaning.removeDuplicates" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">فعال‌سازی حذف خودکار</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حذف رکوردهای قدیمی</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="dataCleaning.removeOldRecords" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">حذف رکوردهای قدیمی‌تر از</span>
                <input type="number" v-model="dataCleaning.oldRecordDays" min="30" max="365" class="w-20 px-2 py-1 border border-gray-300 rounded text-sm">
                <span class="text-sm text-gray-700 mr-2">روز</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی داده</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="dataCleaning.compression" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">فشرده‌سازی خودکار</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس پاکسازی</label>
            <select v-model="dataCleaning.frequency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="daily">روزانه</option>
              <option value="weekly">هفتگی</option>
              <option value="monthly">ماهانه</option>
              <option value="manual">دستی</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- نظارت بر کیفیت -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- گزارش‌های کیفیت -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">گزارش‌های کیفیت</h3>
          </div>
          <button class="bg-purple-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-purple-700 transition-colors">
            ایجاد گزارش
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="report in qualityReports" :key="report.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <h4 class="font-medium text-gray-900">{{ report.name }}</h4>
              <span class="text-xs px-2 py-1 rounded-full" :class="report.statusClass">{{ report.status }}</span>
            </div>
            <p class="text-sm text-gray-600 mb-2">{{ report.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>آخرین اجرا: {{ report.lastRun }}</span>
              <span>نتیجه: {{ report.result }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- هشدارهای کیفیت -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">هشدارهای کیفیت</h3>
          </div>
          <button class="bg-red-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-red-700 transition-colors">
            تنظیم هشدار
          </button>
        </div>
        
        <div class="space-y-4">
          <div v-for="alert in qualityAlerts" :key="alert.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <h4 class="font-medium text-gray-900">{{ alert.name }}</h4>
              <span class="text-xs px-2 py-1 rounded-full" :class="alert.severityClass">{{ alert.severity }}</span>
            </div>
            <p class="text-sm text-gray-600 mb-2">{{ alert.description }}</p>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>آستانه: {{ alert.threshold }}</span>
              <span>وضعیت: {{ alert.status }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار کیفیت داده -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">آمار کیفیت داده</h3>
      </div>
      <div class="p-6">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع داده</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد کل</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">داده‌های کامل</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">داده‌های ناقص</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کیفیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="dataType in dataQualityStats" :key="dataType.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="dataType.bgColor">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="dataType.icon"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ dataType.name }}</div>
                      <div class="text-sm text-gray-500">{{ dataType.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ dataType.totalCount }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ dataType.completeCount }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ dataType.incompleteCount }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                      <div class="bg-green-600 h-2 rounded-full" :style="{ width: dataType.quality + '%' }"></div>
                    </div>
                    <span class="text-sm text-gray-900">{{ dataType.quality }}%</span>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button class="text-indigo-600 hover:text-indigo-900 mr-3">جزئیات</button>
                  <button class="text-green-600 hover:text-green-900">بهبود</button>
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
// قوانین اعتبارسنجی
const validationRules = ref([
  {
    id: 1,
    name: 'اعتبارسنجی ایمیل',
    description: 'بررسی فرمت صحیح آدرس‌های ایمیل',
    type: 'Format',
    priority: 'High',
    enabled: true
  },
  {
    id: 2,
    name: 'اعتبارسنجی شماره تلفن',
    description: 'بررسی فرمت صحیح شماره‌های تلفن',
    type: 'Format',
    priority: 'High',
    enabled: true
  },
  {
    id: 3,
    name: 'بررسی محدوده قیمت',
    description: 'اعتبارسنجی قیمت‌های محصولات',
    type: 'Range',
    priority: 'Medium',
    enabled: true
  },
  {
    id: 4,
    name: 'بررسی دسته‌بندی',
    description: 'اعتبارسنجی دسته‌بندی‌های محصولات',
    type: 'Reference',
    priority: 'Medium',
    enabled: false
  }
])

// تنظیمات پاکسازی داده
const dataCleaning = ref({
  removeDuplicates: true,
  removeOldRecords: true,
  oldRecordDays: 90,
  compression: true,
  frequency: 'weekly'
})

// گزارش‌های کیفیت
const qualityReports = ref([
  {
    id: 1,
    name: 'گزارش کیفیت روزانه',
    description: 'گزارش کیفیت داده‌های ورودی روزانه',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastRun: '2 ساعت پیش',
    result: 'موفق'
  },
  {
    id: 2,
    name: 'گزارش پاکسازی هفتگی',
    description: 'گزارش پاکسازی داده‌های تکراری',
    status: 'در انتظار',
    statusClass: 'bg-yellow-100 text-yellow-800',
    lastRun: '6 روز پیش',
    result: 'موفق'
  },
  {
    id: 3,
    name: 'گزارش اعتبارسنجی',
    description: 'گزارش اعتبارسنجی داده‌های محصولات',
    status: 'فعال',
    statusClass: 'bg-green-100 text-green-800',
    lastRun: '1 ساعت پیش',
    result: 'موفق'
  }
])

// هشدارهای کیفیت
const qualityAlerts = ref([
  {
    id: 1,
    name: 'کیفیت پایین داده',
    description: 'هشدار در صورت کاهش کیفیت داده‌ها',
    severity: 'متوسط',
    severityClass: 'bg-yellow-100 text-yellow-800',
    threshold: '< 90%',
    status: 'فعال'
  },
  {
    id: 2,
    name: 'داده‌های تکراری زیاد',
    description: 'هشدار در صورت افزایش داده‌های تکراری',
    severity: 'کم',
    severityClass: 'bg-blue-100 text-blue-800',
    threshold: '> 5%',
    status: 'فعال'
  },
  {
    id: 3,
    name: 'داده‌های نامعتبر',
    description: 'هشدار در صورت وجود داده‌های نامعتبر',
    severity: 'زیاد',
    severityClass: 'bg-red-100 text-red-800',
    threshold: '> 2%',
    status: 'غیرفعال'
  }
])

// آمار کیفیت داده
const dataQualityStats = ref([
  {
    id: 1,
    name: 'کاربران',
    description: 'داده‌های کاربران سیستم',
    totalCount: '12,456',
    completeCount: '12,234',
    incompleteCount: '222',
    quality: 98.2,
    bgColor: 'bg-blue-500',
    icon: 'M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z'
  },
  {
    id: 2,
    name: 'محصولات',
    description: 'داده‌های محصولات',
    totalCount: '8,234',
    completeCount: '8,156',
    incompleteCount: '78',
    quality: 99.0,
    bgColor: 'bg-green-500',
    icon: 'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4'
  },
  {
    id: 3,
    name: 'سفارشات',
    description: 'داده‌های سفارشات',
    totalCount: '45,678',
    completeCount: '45,123',
    incompleteCount: '555',
    quality: 98.8,
    bgColor: 'bg-purple-500',
    icon: 'M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z'
  },
  {
    id: 4,
    name: 'بازخوردها',
    description: 'داده‌های بازخورد کاربران',
    totalCount: '23,456',
    completeCount: '22,890',
    incompleteCount: '566',
    quality: 97.6,
    bgColor: 'bg-yellow-500',
    icon: 'M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z'
  }
])
</script> 
