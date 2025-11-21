<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">نمایش وضعیت زنده</h4>
        <p class="text-sm text-gray-600 mt-1">نمایش وضعیت زنده اتصالات و سیستم</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <div class="flex items-center">
          <div class="w-2 h-2 bg-green-500 rounded-full ml-2"></div>
          <span class="text-sm text-gray-600">زنده</span>
        </div>
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="refreshStatus"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          به‌روزرسانی
        </button>
      </div>
    </div>

    <!-- وضعیت کلی سیستم -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">وضعیت کلی سیستم</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div
          v-for="status in systemStatus"
          :key="status.id"
          class="p-6 border border-gray-200 rounded-lg"
        >
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">{{ status.name }}</h6>
            <div
              class="w-3 h-3 rounded-full"
              :class="getStatusColor(status.status)"
            ></div>
          </div>
          <div class="text-2xl font-bold" :class="getStatusTextColor(status.status)">
            {{ status.value }}
          </div>
          <div class="text-xs text-gray-500 mt-1">{{ status.description }}</div>
        </div>
      </div>
    </div>

    <!-- وضعیت اتصالات -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">وضعیت اتصالات</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="connection in connections"
            :key="connection.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getStatusColor(connection.status)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ connection.name }}</h6>
                <p class="text-xs text-gray-500">{{ connection.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ connection.lastSync }}</div>
                <div class="text-xs text-gray-500">آخرین همگام‌سازی</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ connection.transactionCount }}</div>
                <div class="text-xs text-gray-500">تراکنش</div>
              </div>
              <button
                class="text-blue-600 hover:text-blue-800"
                @click="testConnection(connection)"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار وضعیت زنده -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">نمودار وضعیت زنده</h5>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 mx-auto mb-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
          </svg>
          <p class="text-gray-500">نمودار وضعیت زنده</p>
          <p class="text-sm text-gray-400 mt-1">نمایش روند وضعیت در زمان واقعی</p>
        </div>
      </div>
    </div>

    <!-- رویدادهای اخیر -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">رویدادهای اخیر</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="event in recentEvents"
            :key="event.id"
            class="flex items-start space-x-4 space-x-reverse"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getEventColor(event.type)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ event.title }}</h6>
                <span class="text-xs text-gray-500">{{ event.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ event.description }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار عملکرد -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار عملکرد</h5>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <h6 class="text-sm font-medium text-gray-700 mb-2">تراکنش‌های امروز</h6>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span>موفق</span>
              <span class="font-medium text-green-600">{{ performanceStats.today.success }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span>ناموفق</span>
              <span class="font-medium text-red-600">{{ performanceStats.today.failed }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span>در انتظار</span>
              <span class="font-medium text-yellow-600">{{ performanceStats.today.pending }}</span>
            </div>
          </div>
        </div>

        <div>
          <h6 class="text-sm font-medium text-gray-700 mb-2">میانگین زمان پاسخ</h6>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span>زرین‌پال</span>
              <span class="font-medium">{{ performanceStats.responseTime.zarinpal }}ms</span>
            </div>
            <div class="flex justify-between text-sm">
              <span>نکست‌پی</span>
              <span class="font-medium">{{ performanceStats.responseTime.nextpay }}ms</span>
            </div>
            <div class="flex justify-between text-sm">
              <span>کلی</span>
              <span class="font-medium">{{ performanceStats.responseTime.average }}ms</span>
            </div>
          </div>
        </div>

        <div>
          <h6 class="text-sm font-medium text-gray-700 mb-2">نرخ موفقیت</h6>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span>امروز</span>
              <span class="font-medium text-green-600">{{ performanceStats.successRate.today }}%</span>
            </div>
            <div class="flex justify-between text-sm">
              <span>هفته</span>
              <span class="font-medium text-blue-600">{{ performanceStats.successRate.week }}%</span>
            </div>
            <div class="flex justify-between text-sm">
              <span>ماه</span>
              <span class="font-medium text-purple-600">{{ performanceStats.successRate.month }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات نمایش -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات نمایش</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فاصله به‌روزرسانی</label>
          <select
            v-model="displaySettings.refreshInterval"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="5">۵ ثانیه</option>
            <option value="10">۱۰ ثانیه</option>
            <option value="30">۳۰ ثانیه</option>
            <option value="60">۱ دقیقه</option>
            <option value="300">۵ دقیقه</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نمایش اعلان‌ها</label>
          <select
            v-model="displaySettings.showNotifications"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه</option>
            <option value="errors">فقط خطاها</option>
            <option value="warnings">خطاها و هشدارها</option>
            <option value="none">هیچ‌کدام</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نمایش نمودار</label>
          <select
            v-model="displaySettings.showChart"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="always">همیشه</option>
            <option value="onDemand">در صورت درخواست</option>
            <option value="never">هرگز</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نمایش رویدادها</label>
          <input
            v-model="displaySettings.maxEvents"
            type="number"
            min="5"
            max="50"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="displaySettings.autoRefresh"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">به‌روزرسانی خودکار</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="displaySettings.showTimestamps"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">نمایش زمان‌ها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="displaySettings.animateChanges"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">انیمیشن تغییرات</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت کلی سیستم
const systemStatus = ref([
  {
    id: 1,
    name: 'اتصالات فعال',
    value: '۸',
    status: 'success',
    description: 'از ۱۰ اتصال'
  },
  {
    id: 2,
    name: 'تراکنش‌های امروز',
    value: '۱,۲۴۵',
    status: 'success',
    description: '+۱۲% نسبت به دیروز'
  },
  {
    id: 3,
    name: 'خطاهای امروز',
    value: '۲۳',
    status: 'warning',
    description: '-۵% نسبت به دیروز'
  },
  {
    id: 4,
    name: 'زمان پاسخ',
    value: '۲۴۰ms',
    status: 'success',
    description: 'میانگین'
  }
])

// وضعیت اتصالات
const connections = ref([
  {
    id: 1,
    name: 'زرین‌پال',
    description: 'درگاه پرداخت اصلی',
    status: 'success',
    lastSync: '۲ دقیقه پیش',
    transactionCount: 456
  },
  {
    id: 2,
    name: 'نکست‌پی',
    description: 'درگاه پرداخت ثانویه',
    status: 'success',
    lastSync: '۵ دقیقه پیش',
    transactionCount: 234
  },
  {
    id: 3,
    name: 'ملت',
    description: 'درگاه بانکی',
    status: 'warning',
    lastSync: '۱۵ دقیقه پیش',
    transactionCount: 89
  },
  {
    id: 4,
    name: 'پارسیان',
    description: 'درگاه بانکی',
    status: 'error',
    lastSync: '۱ ساعت پیش',
    transactionCount: 0
  }
])

// رویدادهای اخیر
const recentEvents = ref([
  {
    id: 1,
    type: 'success',
    title: 'اتصال زرین‌پال برقرار شد',
    description: 'اتصال به درگاه زرین‌پال با موفقیت برقرار شد',
    timestamp: '۲ دقیقه پیش'
  },
  {
    id: 2,
    type: 'warning',
    title: 'تاخیر در اتصال ملت',
    description: 'زمان پاسخ اتصال ملت بیش از حد معمول است',
    timestamp: '۱۵ دقیقه پیش'
  },
  {
    id: 3,
    type: 'error',
    title: 'قطع اتصال پارسیان',
    description: 'اتصال به درگاه پارسیان قطع شده است',
    timestamp: '۱ ساعت پیش'
  },
  {
    id: 4,
    type: 'info',
    title: 'به‌روزرسانی سیستم',
    description: 'سیستم به‌روزرسانی شد و تمام اتصالات بررسی شدند',
    timestamp: '۲ ساعت پیش'
  }
])

// آمار عملکرد
const performanceStats = ref({
  today: {
    success: 1189,
    failed: 34,
    pending: 22
  },
  responseTime: {
    zarinpal: 180,
    nextpay: 220,
    average: 240
  },
  successRate: {
    today: 96.5,
    week: 95.2,
    month: 94.8
  }
})

// تنظیمات نمایش
const displaySettings = ref({
  refreshInterval: 30,
  showNotifications: 'all',
  showChart: 'always',
  maxEvents: 20,
  autoRefresh: true,
  showTimestamps: true,
  animateChanges: true
})

// رنگ‌های وضعیت
const getStatusColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500',
    info: 'bg-blue-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getStatusTextColor = (status: string) => {
  const colors = {
    success: 'text-green-600',
    warning: 'text-yellow-600',
    error: 'text-red-600',
    info: 'text-blue-600'
  }
  return colors[status] || 'text-gray-600'
}

const getEventColor = (type: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500',
    info: 'bg-blue-500'
  }
  return colors[type] || 'bg-gray-500'
}

// عملیات
const refreshStatus = () => {

}

interface Connection {
  id: number | string
  [key: string]: unknown
}

const testConnection = (_connection: Connection) => {

}
</script>

<!--
  کامپوننت نمایش وضعیت زنده
  شامل:
  1. وضعیت کلی سیستم
  2. وضعیت اتصالات
  3. نمودار وضعیت زنده
  4. رویدادهای اخیر
  5. آمار عملکرد
  6. تنظیمات نمایش
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 
