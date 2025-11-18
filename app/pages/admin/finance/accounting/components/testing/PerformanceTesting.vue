<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">تست عملکرد</h4>
        <p class="text-sm text-gray-600 mt-1">تست سرعت، بار و استرس سیستم</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="runPerformanceTest"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
          اجرای تست
        </button>
      </div>
    </div>

    <!-- آمار عملکرد -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار عملکرد</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">زمان پاسخ متوسط</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ performanceStats.avgResponseTime }}ms</div>
          <div class="text-xs text-gray-500 mt-1">میانگین</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">درخواست در ثانیه</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ performanceStats.requestsPerSecond }}</div>
          <div class="text-xs text-gray-500 mt-1">RPS</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">نرخ خطا</h6>
            <div class="w-3 h-3 bg-red-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-red-600">{{ performanceStats.errorRate }}%</div>
          <div class="text-xs text-gray-500 mt-1">خطاها</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">استفاده از CPU</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ performanceStats.cpuUsage }}%</div>
          <div class="text-xs text-gray-500 mt-1">پردازنده</div>
        </div>
      </div>
    </div>

    <!-- انواع تست -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">انواع تست عملکرد</h5>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center mb-3">
            <svg class="w-8 h-8 text-blue-600 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            <h6 class="text-sm font-medium text-gray-900">تست سرعت</h6>
          </div>
          <p class="text-xs text-gray-600 mb-3">تست سرعت پاسخ سیستم در شرایط عادی</p>
          <div class="space-y-2 text-xs text-gray-500">
            <div>زمان پاسخ: {{ speedTest.responseTime }}ms</div>
            <div>درخواست‌ها: {{ speedTest.requests }}</div>
            <div>وضعیت: {{ getTestStatus(speedTest.status) }}</div>
          </div>
          <button
            @click="runSpeedTest"
            class="mt-3 w-full px-3 py-2 bg-blue-600 text-white text-xs rounded hover:bg-blue-700"
          >
            اجرای تست سرعت
          </button>
        </div>

        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center mb-3">
            <svg class="w-8 h-8 text-green-600 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            <h6 class="text-sm font-medium text-gray-900">تست بار</h6>
          </div>
          <p class="text-xs text-gray-600 mb-3">تست عملکرد سیستم تحت بار بالا</p>
          <div class="space-y-2 text-xs text-gray-500">
            <div>زمان پاسخ: {{ loadTest.responseTime }}ms</div>
            <div>درخواست‌ها: {{ loadTest.requests }}</div>
            <div>وضعیت: {{ getTestStatus(loadTest.status) }}</div>
          </div>
          <button
            @click="runLoadTest"
            class="mt-3 w-full px-3 py-2 bg-green-600 text-white text-xs rounded hover:bg-green-700"
          >
            اجرای تست بار
          </button>
        </div>

        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center mb-3">
            <svg class="w-8 h-8 text-red-600 ml-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
            <h6 class="text-sm font-medium text-gray-900">تست استرس</h6>
          </div>
          <p class="text-xs text-gray-600 mb-3">تست حد نهایی تحمل سیستم</p>
          <div class="space-y-2 text-xs text-gray-500">
            <div>زمان پاسخ: {{ stressTest.responseTime }}ms</div>
            <div>درخواست‌ها: {{ stressTest.requests }}</div>
            <div>وضعیت: {{ getTestStatus(stressTest.status) }}</div>
          </div>
          <button
            @click="runStressTest"
            class="mt-3 w-full px-3 py-2 bg-red-600 text-white text-xs rounded hover:bg-red-700"
          >
            اجرای تست استرس
          </button>
        </div>
      </div>
    </div>

    <!-- نمودار عملکرد -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">نمودار عملکرد</h5>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 mx-auto mb-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
          </svg>
          <p class="text-gray-500">نمودار عملکرد</p>
          <p class="text-sm text-gray-400 mt-1">نمایش روند عملکرد در زمان واقعی</p>
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
                <h6 class="text-sm font-medium text-gray-900">{{ result.testType }}</h6>
                <span class="text-xs text-gray-500">{{ result.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ result.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>زمان پاسخ: {{ result.responseTime }}ms</span>
                <span>درخواست‌ها: {{ result.requests }}</span>
                <span>نرخ خطا: {{ result.errorRate }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات تست -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات تست عملکرد</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مدت زمان تست</label>
          <select
            v-model="testSettings.duration"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="30">۳۰ ثانیه</option>
            <option value="60">۱ دقیقه</option>
            <option value="300">۵ دقیقه</option>
            <option value="600">۱۰ دقیقه</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد کاربران همزمان</label>
          <input
            v-model="testSettings.concurrentUsers"
            type="number"
            min="1"
            max="1000"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تایم‌اوت درخواست</label>
          <input
            v-model="testSettings.requestTimeout"
            type="number"
            min="1"
            max="60"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">ثانیه</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع درخواست</label>
          <select
            v-model="testSettings.requestType"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="mixed">ترکیبی</option>
            <option value="read">فقط خواندن</option>
            <option value="write">فقط نوشتن</option>
            <option value="sync">همگام‌سازی</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="testSettings.autoStop"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">توقف خودکار در صورت خطا</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="testSettings.saveResults"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">ذخیره نتایج</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="testSettings.generateReport"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">تولید گزارش</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار عملکرد
const performanceStats = ref({
  avgResponseTime: 240,
  requestsPerSecond: 45,
  errorRate: 2.3,
  cpuUsage: 65
})

// تست سرعت
const speedTest = ref({
  responseTime: 180,
  requests: 100,
  status: 'success'
})

// تست بار
const loadTest = ref({
  responseTime: 350,
  requests: 500,
  status: 'warning'
})

// تست استرس
const stressTest = ref({
  responseTime: 1200,
  requests: 1000,
  status: 'error'
})

// نتایج تست‌های اخیر
const recentResults = ref([
  {
    id: 1,
    testType: 'تست سرعت',
    description: 'تست سرعت در شرایط عادی',
    responseTime: 180,
    requests: 100,
    errorRate: 0.5,
    status: 'success',
    timestamp: '۲ دقیقه پیش'
  },
  {
    id: 2,
    testType: 'تست بار',
    description: 'تست بار با ۵۰۰ کاربر همزمان',
    responseTime: 350,
    requests: 500,
    errorRate: 2.1,
    status: 'warning',
    timestamp: '۱۵ دقیقه پیش'
  },
  {
    id: 3,
    testType: 'تست استرس',
    description: 'تست استرس با ۱۰۰۰ کاربر همزمان',
    responseTime: 1200,
    requests: 1000,
    errorRate: 8.5,
    status: 'error',
    timestamp: '۱ ساعت پیش'
  }
])

// تنظیمات تست
const testSettings = ref({
  duration: 60,
  concurrentUsers: 100,
  requestTimeout: 30,
  requestType: 'mixed',
  autoStop: true,
  saveResults: true,
  generateReport: true
})

// رنگ‌های وضعیت
const getResultColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getTestStatus = (status: string) => {
  const texts = {
    success: 'موفق',
    warning: 'هشدار',
    error: 'خطا'
  }
  return texts[status] || status
}

// عملیات
const runPerformanceTest = () => {
  console.log('اجرای تست عملکرد')
}

const runSpeedTest = () => {
  console.log('اجرای تست سرعت')
}

const runLoadTest = () => {
  console.log('اجرای تست بار')
}

const runStressTest = () => {
  console.log('اجرای تست استرس')
}
</script>

<!--
  کامپوننت تست عملکرد
  شامل:
  1. آمار عملکرد
  2. انواع تست (سرعت، بار، استرس)
  3. نمودار عملکرد
  4. نتایج تست‌های اخیر
  5. تنظیمات تست
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
