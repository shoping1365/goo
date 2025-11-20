<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-indigo-50 to-purple-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">🔌 مدیریت API و وب‌هوک‌ها</h2>
      <p class="text-gray-600">مدیریت API ها، کلیدها و وب‌هوک‌های کیف پول</p>
    </div>

    <!-- آمار API -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کل درخواست‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کل درخواست‌ها</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">امروز</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ apiStats.totalRequests }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ apiStats.requestGrowth }}%</span>
          <span class="mx-2">از دیروز</span>
        </div>
      </div>

      <!-- نرخ موفقیت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">نرخ موفقیت</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">عالی</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ apiStats.successRate }}%</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ apiStats.successfulRequests }}</span>
          <span class="mx-2">درخواست موفق</span>
        </div>
      </div>

      <!-- کلیدهای فعال -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کلیدهای فعال</h3>
          <span class="text-xs bg-purple-100 text-purple-800 rounded-full px-3 py-1">فعال</span>
        </div>
        <div class="text-3xl font-bold text-purple-600 mb-2">{{ apiStats.activeKeys }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-purple-500">{{ apiStats.keyUsage }}%</span>
          <span class="mx-2">نرخ استفاده</span>
        </div>
      </div>

      <!-- وب‌هوک‌های فعال -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">وب‌هوک‌های فعال</h3>
          <span class="text-xs bg-orange-100 text-orange-800 rounded-full px-3 py-1">فعال</span>
        </div>
        <div class="text-3xl font-bold text-orange-600 mb-2">{{ apiStats.activeWebhooks }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-orange-500">{{ apiStats.webhookDeliveryRate }}%</span>
          <span class="mx-2">نرخ تحویل</span>
        </div>
      </div>
    </div>

    <!-- مدیریت کلیدهای API -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">کلیدهای API</h3>
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          ایجاد کلید جدید
        </button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">نام کلید</th>
              <th class="px-4 py-3 text-gray-700 font-medium">کلید API</th>
              <th class="px-4 py-3 text-gray-700 font-medium">سطح دسترسی</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ ایجاد</th>
              <th class="px-4 py-3 text-gray-700 font-medium">آخرین استفاده</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="key in apiKeys" :key="key.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-gray-900 font-medium">{{ key.name }}</td>
              <td class="px-4 py-3">
                <div class="flex items-center">
                  <span class="font-mono text-xs text-gray-600">{{ key.keyPrefix }}...</span>
                  <button class="mr-2 text-blue-600 hover:text-blue-800 text-sm">نمایش</button>
                </div>
              </td>
              <td class="px-4 py-3">
                <span :class="getAccessLevelClass(key.accessLevel)" class="px-2 py-1 text-xs rounded-full">
                  {{ key.accessLevel }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ key.createdAt }}</td>
              <td class="px-4 py-3 text-gray-600">{{ key.lastUsed }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(key.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ key.status }}
                </span>
              </td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button
v-if="key.status === 'فعال'" 
                          class="text-red-600 hover:text-red-800 text-sm">غیرفعال</button>
                  <button
v-if="key.status === 'غیرفعال'" 
                          class="text-green-600 hover:text-green-800 text-sm">فعال</button>
                  <button class="text-red-600 hover:text-red-800 text-sm">حذف</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مدیریت وب‌هوک‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">وب‌هوک‌ها</h3>
        <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
          افزودن وب‌هوک جدید
        </button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">نام</th>
              <th class="px-4 py-3 text-gray-700 font-medium">URL</th>
              <th class="px-4 py-3 text-gray-700 font-medium">رویدادها</th>
              <th class="px-4 py-3 text-gray-700 font-medium">آخرین تحویل</th>
              <th class="px-4 py-3 text-gray-700 font-medium">نرخ موفقیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="webhook in webhooks" :key="webhook.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-gray-900 font-medium">{{ webhook.name }}</td>
              <td class="px-4 py-3">
                <div class="max-w-xs truncate text-gray-600">{{ webhook.url }}</div>
              </td>
              <td class="px-4 py-3">
                <div class="flex flex-wrap gap-1">
                  <span
v-for="event in webhook.events" :key="event" 
                        class="px-2 py-1 text-xs bg-blue-100 text-blue-800 rounded-full">
                    {{ event }}
                  </span>
                </div>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ webhook.lastDelivery }}</td>
              <td class="px-4 py-3">
                <span :class="getSuccessRateClass(webhook.successRate)" class="px-2 py-1 text-xs rounded-full">
                  {{ webhook.successRate }}%
                </span>
              </td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(webhook.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ webhook.status }}
                </span>
              </td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button class="text-green-600 hover:text-green-800 text-sm">تست</button>
                  <button
v-if="webhook.status === 'فعال'" 
                          class="text-red-600 hover:text-red-800 text-sm">غیرفعال</button>
                  <button
v-if="webhook.status === 'غیرفعال'" 
                          class="text-green-600 hover:text-green-800 text-sm">فعال</button>
                  <button class="text-red-600 hover:text-red-800 text-sm">حذف</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- لاگ‌های API -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">لاگ‌های API</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
            پاک کردن لاگ‌ها
          </button>
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            خروجی اکسل
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">زمان</th>
              <th class="px-4 py-3 text-gray-700 font-medium">کلید API</th>
              <th class="px-4 py-3 text-gray-700 font-medium">روش</th>
              <th class="px-4 py-3 text-gray-700 font-medium">مسیر</th>
              <th class="px-4 py-3 text-gray-700 font-medium">کد پاسخ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">زمان پاسخ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">IP</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="log in apiLogs" :key="log.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-gray-600">{{ log.timestamp }}</td>
              <td class="px-4 py-3 text-gray-600 font-mono text-xs">{{ log.apiKey }}</td>
              <td class="px-4 py-3">
                <span :class="getMethodClass(log.method)" class="px-2 py-1 text-xs rounded-full">
                  {{ log.method }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600 font-mono text-xs">{{ log.path }}</td>
              <td class="px-4 py-3">
                <span :class="getResponseCodeClass(log.responseCode)" class="px-2 py-1 text-xs rounded-full">
                  {{ log.responseCode }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ log.responseTime }}ms</td>
              <td class="px-4 py-3 text-gray-600 font-mono text-xs">{{ log.ipAddress }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- تنظیمات API -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات عمومی -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات عمومی API</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت درخواست (در دقیقه)</label>
            <input
type="number" :value="apiSettings.rateLimit" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان انقضای توکن (ساعت)</label>
            <input
type="number" :value="apiSettings.tokenExpiry" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اندازه درخواست (MB)</label>
            <input
type="number" :value="apiSettings.maxRequestSize" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div class="flex items-center">
            <input
type="checkbox" :checked="apiSettings.enableCors" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <label class="mr-2 text-sm text-gray-700">فعال کردن CORS</label>
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره تنظیمات
          </button>
        </div>
      </div>

      <!-- تنظیمات وب‌هوک -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات وب‌هوک</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش تحویل</label>
            <input
type="number" :value="webhookSettings.maxRetries" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله تلاش مجدد (ثانیه)</label>
            <input
type="number" :value="webhookSettings.retryInterval" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان تایم‌اوت (ثانیه)</label>
            <input
type="number" :value="webhookSettings.timeout" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div class="flex items-center">
            <input
type="checkbox" :checked="webhookSettings.enableRetries" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <label class="mr-2 text-sm text-gray-700">فعال کردن تلاش مجدد</label>
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- آمار تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار درخواست‌ها بر اساس روش -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">درخواست‌ها بر اساس روش</h3>
        <div class="space-y-4">
          <div v-for="method in methodStats" :key="method.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: method.color }"></div>
              <span class="text-sm text-gray-700">{{ method.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ method.count }}</span>
              <span class="text-sm text-gray-500">({{ method.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار کدهای پاسخ -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">کدهای پاسخ</h3>
        <div class="space-y-4">
          <div v-for="code in responseCodeStats" :key="code.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: code.color }"></div>
              <span class="text-sm text-gray-700">{{ code.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ code.count }}</span>
              <span class="text-sm text-gray-500">({{ code.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// آمار API
const apiStats = {
  totalRequests: 15420,
  requestGrowth: 12.5,
  successRate: 98.2,
  successfulRequests: 15142,
  activeKeys: 25,
  keyUsage: 85.3,
  activeWebhooks: 8,
  webhookDeliveryRate: 96.8
}

// تنظیمات API
const apiSettings = {
  rateLimit: 1000,
  tokenExpiry: 24,
  maxRequestSize: 10,
  enableCors: true
}

// تنظیمات وب‌هوک
const webhookSettings = {
  maxRetries: 3,
  retryInterval: 30,
  timeout: 10,
  enableRetries: true
}

// کلیدهای API
const apiKeys = [
  {
    id: 1,
    name: 'کلید اصلی',
    keyPrefix: 'sk_live_1234567890abcdef',
    accessLevel: 'کامل',
    createdAt: '1402/08/15',
    lastUsed: '2 دقیقه پیش',
    status: 'فعال'
  },
  {
    id: 2,
    name: 'کلید تست',
    keyPrefix: 'sk_test_abcdef1234567890',
    accessLevel: 'محدود',
    createdAt: '1402/10/01',
    lastUsed: '1 ساعت پیش',
    status: 'فعال'
  },
  {
    id: 3,
    name: 'کلید قدیمی',
    keyPrefix: 'sk_old_9876543210fedcba',
    accessLevel: 'کامل',
    createdAt: '1402/06/20',
    lastUsed: '1 هفته پیش',
    status: 'غیرفعال'
  }
]

// وب‌هوک‌ها
const webhooks = [
  {
    id: 1,
    name: 'اعلان تراکنش',
    url: 'https://example.com/webhook/transaction',
    events: ['transaction.created', 'transaction.completed'],
    lastDelivery: '2 دقیقه پیش',
    successRate: 98.5,
    status: 'فعال'
  },
  {
    id: 2,
    name: 'گزارش روزانه',
    url: 'https://reports.example.com/daily',
    events: ['daily.report'],
    lastDelivery: '1 روز پیش',
    successRate: 95.2,
    status: 'فعال'
  },
  {
    id: 3,
    name: 'هشدار امنیتی',
    url: 'https://security.example.com/alerts',
    events: ['security.alert', 'fraud.detected'],
    lastDelivery: 'هرگز',
    successRate: 0,
    status: 'غیرفعال'
  }
]

// لاگ‌های API
const apiLogs = [
  {
    id: 1,
    timestamp: '1402/10/15 - 14:30:25',
    apiKey: 'sk_live_1234...',
    method: 'GET',
    path: '/api/v1/transactions',
    responseCode: 200,
    responseTime: 45,
    ipAddress: '192.168.1.100'
  },
  {
    id: 2,
    timestamp: '1402/10/15 - 14:29:15',
    apiKey: 'sk_test_abcd...',
    method: 'POST',
    path: '/api/v1/charges',
    responseCode: 201,
    responseTime: 120,
    ipAddress: '192.168.1.101'
  },
  {
    id: 3,
    timestamp: '1402/10/15 - 14:28:30',
    apiKey: 'sk_live_1234...',
    method: 'GET',
    path: '/api/v1/users',
    responseCode: 401,
    responseTime: 15,
    ipAddress: '192.168.1.102'
  }
]

// آمار روش‌های درخواست
const methodStats = [
  { name: 'GET', count: 8500, percentage: 55, color: '#10B981' },
  { name: 'POST', count: 5200, percentage: 34, color: '#3B82F6' },
  { name: 'PUT', count: 1200, percentage: 8, color: '#F59E0B' },
  { name: 'DELETE', count: 520, percentage: 3, color: '#EF4444' }
]

// آمار کدهای پاسخ
const responseCodeStats = [
  { name: '200 (موفق)', count: 13500, percentage: 87.5, color: '#10B981' },
  { name: '201 (ایجاد شده)', count: 1200, percentage: 7.8, color: '#3B82F6' },
  { name: '400 (خطای درخواست)', count: 420, percentage: 2.7, color: '#F59E0B' },
  { name: '401 (غیرمجاز)', count: 180, percentage: 1.2, color: '#EF4444' },
  { name: '500 (خطای سرور)', count: 120, percentage: 0.8, color: '#DC2626' }
]

// تابع کلاس سطح دسترسی
const getAccessLevelClass = (level: string) => {
  switch (level) {
    case 'کامل':
      return 'bg-green-100 text-green-800'
    case 'محدود':
      return 'bg-yellow-100 text-yellow-800'
    case 'فقط خواندن':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'فعال':
      return 'bg-green-100 text-green-800'
    case 'غیرفعال':
      return 'bg-red-100 text-red-800'
    case 'در انتظار':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس نرخ موفقیت
const getSuccessRateClass = (rate: number) => {
  if (rate >= 95) return 'bg-green-100 text-green-800'
  if (rate >= 80) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

// تابع کلاس روش
const getMethodClass = (method: string) => {
  switch (method) {
    case 'GET':
      return 'bg-green-100 text-green-800'
    case 'POST':
      return 'bg-blue-100 text-blue-800'
    case 'PUT':
      return 'bg-yellow-100 text-yellow-800'
    case 'DELETE':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس کد پاسخ
const getResponseCodeClass = (code: number) => {
  if (code >= 200 && code < 300) return 'bg-green-100 text-green-800'
  if (code >= 400 && code < 500) return 'bg-yellow-100 text-yellow-800'
  if (code >= 500) return 'bg-red-100 text-red-800'
  return 'bg-gray-100 text-gray-800'
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. آمار API (درخواست‌ها، نرخ موفقیت، کلیدها، وب‌هوک‌ها)
  2. مدیریت کلیدهای API
  3. مدیریت وب‌هوک‌ها
  4. لاگ‌های API
  5. تنظیمات API و وب‌هوک
  6. نمودارهای تحلیلی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
