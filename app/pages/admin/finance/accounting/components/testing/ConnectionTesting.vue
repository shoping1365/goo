<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">تست اتصالات</h4>
        <p class="text-sm text-gray-600 mt-1">تست و بررسی وضعیت اتصالات حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="runConnectionTests"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
          اجرای تست
        </button>
      </div>
    </div>

    <!-- وضعیت کلی تست‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">وضعیت کلی تست‌ها</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">تست‌های موفق</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ testStats.success }}</div>
          <div class="text-xs text-gray-500 mt-1">از {{ testStats.total }} تست</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">تست‌های ناموفق</h6>
            <div class="w-3 h-3 bg-red-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-red-600">{{ testStats.failed }}</div>
          <div class="text-xs text-gray-500 mt-1">نیاز به بررسی</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">در حال اجرا</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ testStats.running }}</div>
          <div class="text-xs text-gray-500 mt-1">تست‌های فعال</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">زمان متوسط</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ testStats.avgTime }}s</div>
          <div class="text-xs text-gray-500 mt-1">میانگین اجرا</div>
        </div>
      </div>
    </div>

    <!-- لیست اتصالات برای تست -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">اتصالات قابل تست</h5>
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
                <div class="text-sm font-medium text-gray-900">{{ connection.lastTest }}</div>
                <div class="text-xs text-gray-500">آخرین تست</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ connection.responseTime }}ms</div>
                <div class="text-xs text-gray-500">زمان پاسخ</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800"
                  @click="testConnection(connection)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                </button>
                <button
                  class="text-gray-600 hover:text-gray-800"
                  @click="viewConnectionDetails(connection)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نتایج تست‌های اخیر -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">نتایج تست‌های اخیر</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="result in recentResults"
            :key="result.id"
            class="flex items-start space-x-4 space-x-reverse p-6 border border-gray-200 rounded-lg"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getResultColor(result.status)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ result.connectionName }}</h6>
                <span class="text-xs text-gray-500">{{ result.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ result.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>زمان پاسخ: {{ result.responseTime }}ms</span>
                <span>وضعیت: {{ getStatusText(result.status) }}</span>
                <span v-if="result.error" class="text-red-600">خطا: {{ result.error }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات تست -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات تست</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تایم‌اوت تست</label>
          <input
            v-model="testSettings.timeout"
            type="number"
            min="5"
            max="60"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">ثانیه</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد تلاش مجدد</label>
          <input
            v-model="testSettings.retryCount"
            type="number"
            min="0"
            max="5"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">بار</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بین تست‌ها</label>
          <select
            v-model="testSettings.interval"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="5">۵ ثانیه</option>
            <option value="10">۱۰ ثانیه</option>
            <option value="30">۳۰ ثانیه</option>
            <option value="60">۱ دقیقه</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سطح جزئیات</label>
          <select
            v-model="testSettings.logLevel"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="basic">پایه</option>
            <option value="detailed">جزئی</option>
            <option value="verbose">کامل</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="testSettings.autoTest"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">تست خودکار</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="testSettings.notifyOnFailure"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعلان در صورت شکست</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="testSettings.saveResults"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">ذخیره نتایج</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار تست‌ها
const testStats = ref({
  success: 8,
  failed: 2,
  running: 1,
  total: 11,
  avgTime: 2.3
})

// اتصالات قابل تست
const connections = ref([
  {
    id: 1,
    name: 'زرین‌پال',
    description: 'درگاه پرداخت اصلی',
    status: 'success',
    lastTest: '۲ دقیقه پیش',
    responseTime: 180
  },
  {
    id: 2,
    name: 'نکست‌پی',
    description: 'درگاه پرداخت ثانویه',
    status: 'success',
    lastTest: '۵ دقیقه پیش',
    responseTime: 220
  },
  {
    id: 3,
    name: 'ملت',
    description: 'درگاه بانکی',
    status: 'warning',
    lastTest: '۱۵ دقیقه پیش',
    responseTime: 450
  },
  {
    id: 4,
    name: 'پارسیان',
    description: 'درگاه بانکی',
    status: 'error',
    lastTest: '۱ ساعت پیش',
    responseTime: 0
  }
])

// نتایج تست‌های اخیر
const recentResults = ref([
  {
    id: 1,
    connectionName: 'زرین‌پال',
    status: 'success',
    description: 'اتصال با موفقیت برقرار شد',
    responseTime: 180,
    timestamp: '۲ دقیقه پیش'
  },
  {
    id: 2,
    connectionName: 'نکست‌پی',
    status: 'success',
    description: 'تست احراز هویت موفق',
    responseTime: 220,
    timestamp: '۵ دقیقه پیش'
  },
  {
    id: 3,
    connectionName: 'ملت',
    status: 'warning',
    description: 'زمان پاسخ بالا',
    responseTime: 450,
    timestamp: '۱۵ دقیقه پیش'
  },
  {
    id: 4,
    connectionName: 'پارسیان',
    status: 'error',
    description: 'اتصال قطع شده',
    responseTime: 0,
    error: 'Timeout',
    timestamp: '۱ ساعت پیش'
  }
])

// تنظیمات تست
const testSettings = ref({
  timeout: 30,
  retryCount: 2,
  interval: 30,
  logLevel: 'detailed',
  autoTest: true,
  notifyOnFailure: true,
  saveResults: true
})

// رنگ‌های وضعیت
const getStatusColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getResultColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getStatusText = (status: string) => {
  const texts = {
    success: 'موفق',
    warning: 'هشدار',
    error: 'خطا'
  }
  return texts[status] || status
}

// عملیات
const runConnectionTests = () => {
  console.log('اجرای تست‌های اتصال')
}

const testConnection = (connection: any) => {
  console.log('تست اتصال:', connection)
}

const viewConnectionDetails = (connection: any) => {
  console.log('مشاهده جزئیات اتصال:', connection)
}
</script>

<!--
  کامپوننت تست اتصالات
  شامل:
  1. وضعیت کلی تست‌ها
  2. لیست اتصالات قابل تست
  3. نتایج تست‌های اخیر
  4. تنظیمات تست
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
