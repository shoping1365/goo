<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">گزارش وضعیت اتصال</h4>
        <p class="text-sm text-gray-600 mt-1">وضعیت اتصال به نرم‌افزارهای حسابداری و آمار کلی</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          :disabled="isTesting"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="testAllConnections"
        >
          <svg v-if="isTesting" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          {{ isTesting ? 'در حال تست...' : 'تست همه اتصالات' }}
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="exportStatusReport"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          صادر کردن گزارش
        </button>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ stats.totalConnections }}</div>
            <div class="text-sm text-blue-700">کل اتصالات</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.111 16.404a5.5 5.5 0 017.778 0M12 20h.01m-7.08-7.071c3.904-3.905 10.236-3.905 14.141 0M1.394 9.393c5.857-5.857 15.355-5.857 21.213 0" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ stats.activeConnections }}</div>
            <div class="text-sm text-green-700">اتصالات فعال</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ stats.warningConnections }}</div>
            <div class="text-sm text-yellow-700">اتصالات هشدار</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ stats.failedConnections }}</div>
            <div class="text-sm text-red-700">اتصالات ناموفق</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول وضعیت اتصالات -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">وضعیت اتصالات</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">آخرین همگام‌سازی</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">زمان پاسخ</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرخ موفقیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="connection in connections" :key="connection.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <img :src="connection.logo" :alt="connection.name" class="w-8 h-8 rounded" />
                  <div>
                    <div class="font-medium text-gray-900">{{ connection.name }}</div>
                    <div class="text-xs text-gray-500">{{ connection.version }}</div>
                  </div>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div
                    class="w-2 h-2 rounded-full"
                    :class="getStatusColor(connection.status)"
                  ></div>
                  <span
                    class="px-2 py-1 rounded-full text-xs font-medium"
                    :class="getStatusClass(connection.status)"
                  >
                    {{ getStatusLabel(connection.status) }}
                  </span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ connection.lastSync }}</div>
                <div class="text-xs text-gray-500">{{ connection.lastSyncTime }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ connection.responseTime }}ms</div>
                <div class="text-xs text-gray-500">{{ connection.responseStatus }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div class="flex-1 bg-gray-200 rounded-full h-2">
                    <div
                      class="h-2 rounded-full"
                      :class="getSuccessRateColor(connection.successRate)"
                      :style="{ width: connection.successRate + '%' }"
                    ></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ connection.successRate }}%</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="تست اتصال"
                    @click="testConnection(connection)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-green-600 hover:text-green-800"
                    title="مشاهده جزئیات"
                    @click="viewConnectionDetails(connection)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-yellow-600 hover:text-yellow-800"
                    title="اتصال مجدد"
                    @click="reconnectConnection(connection)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودار وضعیت اتصالات -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">توزیع وضعیت اتصالات</h5>
        <div class="space-y-4">
          <div v-for="status in connectionStatusDistribution" :key="status.status" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :class="getStatusColor(status.status)"></div>
              <span class="text-sm text-gray-700">{{ getStatusLabel(status.status) }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="getStatusColor(status.status)"
                  :style="{ width: status.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ status.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">نرخ موفقیت اتصالات</h5>
        <div class="space-y-4">
          <div v-for="software in successRates" :key="software.name" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <img :src="software.logo" :alt="software.name" class="w-6 h-6 rounded" />
              <span class="text-sm text-gray-700">{{ software.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="getSuccessRateColor(software.rate)"
                  :style="{ width: software.rate + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ software.rate }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت تست
const isTesting = ref(false)

// آمار کلی
const stats = ref({
  totalConnections: 8,
  activeConnections: 6,
  warningConnections: 1,
  failedConnections: 1
})

// اتصالات
const connections = ref([
  {
    id: 1,
    name: 'هلو',
    version: 'نسخه ۹.۲',
    logo: '/images/helo-logo.png',
    status: 'active',
    lastSync: 'امروز ۱۴:۳۰',
    lastSyncTime: '۲ دقیقه پیش',
    responseTime: 245,
    responseStatus: 'موفق',
    successRate: 98.5
  },
  {
    id: 2,
    name: 'پارسیان',
    version: 'نسخه ۸.۱',
    logo: '/images/parsian-logo.png',
    status: 'active',
    lastSync: 'امروز ۱۴:۲۵',
    lastSyncTime: '۷ دقیقه پیش',
    responseTime: 189,
    responseStatus: 'موفق',
    successRate: 99.2
  },
  {
    id: 3,
    name: 'سپیدار',
    version: 'نسخه ۷.۵',
    logo: '/images/sepidar-logo.png',
    status: 'warning',
    lastSync: 'امروز ۱۳:۴۵',
    lastSyncTime: '۴۷ دقیقه پیش',
    responseTime: 1250,
    responseStatus: 'کند',
    successRate: 85.3
  },
  {
    id: 4,
    name: 'قیاس',
    version: 'نسخه ۶.۸',
    logo: '/images/ghias-logo.png',
    status: 'failed',
    lastSync: 'دیروز ۱۸:۲۰',
    lastSyncTime: '۲۰ ساعت پیش',
    responseTime: 0,
    responseStatus: 'ناموفق',
    successRate: 0
  }
])

// توزیع وضعیت اتصالات
const connectionStatusDistribution = ref([
  { status: 'active', percentage: 75 },
  { status: 'warning', percentage: 12.5 },
  { status: 'failed', percentage: 12.5 }
])

// نرخ موفقیت نرم‌افزارها
const successRates = ref([
  { name: 'هلو', logo: '/images/helo-logo.png', rate: 98.5 },
  { name: 'پارسیان', logo: '/images/parsian-logo.png', rate: 99.2 },
  { name: 'سپیدار', logo: '/images/sepidar-logo.png', rate: 85.3 },
  { name: 'قیاس', logo: '/images/ghias-logo.png', rate: 0 }
])

// رنگ وضعیت
const getStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    warning: 'bg-yellow-500',
    failed: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-700',
    warning: 'bg-yellow-100 text-yellow-700',
    failed: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    warning: 'هشدار',
    failed: 'ناموفق'
  }
  return labels[status] || status
}

// رنگ نرخ موفقیت
const getSuccessRateColor = (rate: number) => {
  if (rate >= 95) return 'bg-green-500'
  if (rate >= 80) return 'bg-yellow-500'
  return 'bg-red-500'
}

// تست همه اتصالات
const testAllConnections = async () => {
  try {
    isTesting.value = true
    // TODO: تست همه اتصالات

    // شبیه‌سازی تست
    await new Promise(resolve => setTimeout(resolve, 3000))

  } catch (error) {
    console.error('خطا در تست اتصالات:', error)
  } finally {
    isTesting.value = false
  }
}

// صادر کردن گزارش وضعیت
const exportStatusReport = () => {
  // TODO: صادر کردن گزارش وضعیت

}

interface Connection {
  id?: number | string
  [key: string]: unknown
}

// تست اتصال
const testConnection = (_connection: Connection) => {
  // TODO: تست اتصال خاص

}

// مشاهده جزئیات اتصال
const viewConnectionDetails = (_connection: Connection) => {
  // TODO: مشاهده جزئیات اتصال

}

// اتصال مجدد
const reconnectConnection = (_connection: Connection) => {
  // TODO: اتصال مجدد

}
</script>

<!--
  کامپوننت گزارش وضعیت اتصال
  شامل:
  1. آمار کلی اتصالات
  2. جدول وضعیت اتصالات
  3. نمودارهای تحلیلی
  4. عملیات تست و مدیریت
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
