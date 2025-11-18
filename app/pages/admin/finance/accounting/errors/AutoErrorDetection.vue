<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">تشخیص خودکار خطاها</h4>
        <p class="text-sm text-gray-600 mt-1">سیستم تشخیص هوشمند خطاهای اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="startMonitoring"
          :disabled="isMonitoring"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          <svg v-if="isMonitoring" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
          {{ isMonitoring ? 'در حال نظارت...' : 'شروع نظارت' }}
        </button>
        <button
          @click="configureDetection"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          تنظیمات تشخیص
        </button>
      </div>
    </div>

    <!-- آمار تشخیص -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ detectionStats.totalErrors }}</div>
            <div class="text-sm text-blue-700">کل خطاهای تشخیص داده شده</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ detectionStats.accuracyRate }}</div>
            <div class="text-sm text-green-700">دقت تشخیص</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ detectionStats.responseTime }}</div>
            <div class="text-sm text-yellow-700">زمان پاسخ (ثانیه)</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-purple-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-purple-600">{{ detectionStats.activeRules }}</div>
            <div class="text-sm text-purple-700">قوانین فعال</div>
          </div>
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات تشخیص -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- قوانین تشخیص -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">قوانین تشخیص خطا</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.connectionTimeout"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تشخیص timeout اتصال</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تشخیص خطاهای timeout در اتصال به نرم‌افزارهای حسابداری</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.authenticationError"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تشخیص خطاهای احراز هویت</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تشخیص خطاهای مربوط به کلیدهای API و احراز هویت</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.dataSyncError"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تشخیص خطاهای همگام‌سازی</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تشخیص خطاهای مربوط به همگام‌سازی داده‌ها</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.formatError"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تشخیص خطاهای فرمت داده</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تشخیص خطاهای مربوط به فرمت و ساختار داده‌ها</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.performanceError"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تشخیص مشکلات عملکرد</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تشخیص کندی و مشکلات عملکرد اتصال</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.securityError"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تشخیص خطاهای امنیتی</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تشخیص خطاهای مربوط به امنیت و رمزگذاری</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات پیشرفته -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات پیشرفته تشخیص</h5>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بررسی (ثانیه)</label>
            <input
              v-model.number="detectionSettings.checkInterval"
              type="number"
              min="5"
              max="300"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500 mt-1">فاصله زمانی بین بررسی‌های تشخیص خطا</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آستانه خطا</label>
            <input
              v-model.number="detectionSettings.errorThreshold"
              type="number"
              min="1"
              max="10"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500 mt-1">تعداد خطاهای مجاز قبل از فعال شدن هشدار</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح حساسیت</label>
            <select
              v-model="detectionSettings.sensitivityLevel"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="low">کم</option>
              <option value="medium">متوسط</option>
              <option value="high">زیاد</option>
            </select>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.machineLearning"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">استفاده از یادگیری ماشین</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">استفاده از الگوریتم‌های ML برای تشخیص بهتر خطاها</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.patternRecognition"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تشخیص الگوهای خطا</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تشخیص الگوهای تکرار شونده در خطاها</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="detectionSettings.predictiveAnalysis"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تحلیل پیش‌بینی‌کننده</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">پیش‌بینی خطاهای احتمالی قبل از وقوع</p>
          </div>
        </div>
      </div>
    </div>

    <!-- خطاهای تشخیص داده شده -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">خطاهای تشخیص داده شده</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع خطا</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">توضیحات</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">اولویت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ تشخیص</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="error in detectedErrors" :key="error.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <div
                    class="w-3 h-3 rounded-full"
                    :class="getErrorTypeColor(error.type)"
                  ></div>
                  <span class="font-medium text-gray-900">{{ error.type }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ error.description }}</div>
                <div class="text-xs text-gray-500">{{ error.details }}</div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getPriorityClass(error.priority)"
                >
                  {{ getPriorityLabel(error.priority) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ error.detectedDate }}</div>
                <div class="text-xs text-gray-500">{{ error.detectedTime }}</div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getStatusClass(error.status)"
                >
                  {{ getStatusLabel(error.status) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    @click="resolveError(error)"
                    class="p-1 text-green-600 hover:text-green-800"
                    title="حل خطا"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                  </button>
                  <button
                    @click="ignoreError(error)"
                    class="p-1 text-yellow-600 hover:text-yellow-800"
                    title="نادیده گرفتن"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                  <button
                    @click="viewErrorDetails(error)"
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="جزئیات"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت نظارت
const isMonitoring = ref(false)

// آمار تشخیص
const detectionStats = ref({
  totalErrors: 23,
  accuracyRate: 94.5,
  responseTime: 2.3,
  activeRules: 12
})

// تنظیمات تشخیص
const detectionSettings = ref({
  connectionTimeout: true,
  authenticationError: true,
  dataSyncError: true,
  formatError: true,
  performanceError: true,
  securityError: true,
  checkInterval: 30,
  errorThreshold: 3,
  sensitivityLevel: 'medium',
  machineLearning: true,
  patternRecognition: true,
  predictiveAnalysis: false
})

// خطاهای تشخیص داده شده
const detectedErrors = ref([
  {
    id: 1,
    type: 'خطای اتصال',
    description: 'Timeout در اتصال به نرم‌افزار هلو',
    details: 'اتصال به سرور هلو پس از ۳۰ ثانیه timeout شد',
    priority: 'high',
    detectedDate: 'امروز',
    detectedTime: '۱۴:۳۰',
    status: 'active'
  },
  {
    id: 2,
    type: 'خطای احراز هویت',
    description: 'کلید API نامعتبر',
    details: 'کلید API برای نرم‌افزار پارسیان منقضی شده است',
    priority: 'medium',
    detectedDate: 'دیروز',
    detectedTime: '۱۸:۴۵',
    status: 'resolved'
  },
  {
    id: 3,
    type: 'خطای همگام‌سازی',
    description: 'خطا در همگام‌سازی تراکنش‌ها',
    details: 'خطای فرمت در داده‌های ارسالی به نرم‌افزار سپیدار',
    priority: 'low',
    detectedDate: 'هفته گذشته',
    detectedTime: '۱۰:۱۵',
    status: 'ignored'
  }
])

// رنگ نوع خطا
const getErrorTypeColor = (type: string) => {
  const colors = {
    'خطای اتصال': 'bg-red-500',
    'خطای احراز هویت': 'bg-yellow-500',
    'خطای همگام‌سازی': 'bg-blue-500',
    'خطای فرمت': 'bg-purple-500',
    'خطای عملکرد': 'bg-orange-500',
    'خطای امنیتی': 'bg-red-600'
  }
  return colors[type] || 'bg-gray-500'
}

// کلاس اولویت
const getPriorityClass = (priority: string) => {
  const classes = {
    high: 'bg-red-100 text-red-700',
    medium: 'bg-yellow-100 text-yellow-700',
    low: 'bg-green-100 text-green-700'
  }
  return classes[priority] || 'bg-gray-100 text-gray-700'
}

// برچسب اولویت
const getPriorityLabel = (priority: string) => {
  const labels = {
    high: 'زیاد',
    medium: 'متوسط',
    low: 'کم'
  }
  return labels[priority] || priority
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-red-100 text-red-700',
    resolved: 'bg-green-100 text-green-700',
    ignored: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    resolved: 'حل شده',
    ignored: 'نادیده گرفته شده'
  }
  return labels[status] || status
}

// شروع نظارت
const startMonitoring = () => {
  isMonitoring.value = !isMonitoring.value
  // TODO: شروع/توقف نظارت
  console.log('نظارت:', isMonitoring.value ? 'شروع شد' : 'متوقف شد')
}

// تنظیمات تشخیص
const configureDetection = () => {
  // TODO: باز کردن مودال تنظیمات
  console.log('باز کردن تنظیمات تشخیص')
}

// حل خطا
const resolveError = (error: any) => {
  // TODO: حل خطا
  console.log('حل خطا:', error)
}

// نادیده گرفتن خطا
const ignoreError = (error: any) => {
  // TODO: نادیده گرفتن خطا
  console.log('نادیده گرفتن خطا:', error)
}

// مشاهده جزئیات خطا
const viewErrorDetails = (error: any) => {
  // TODO: نمایش جزئیات خطا
  console.log('جزئیات خطا:', error)
}
</script>

<!--
  کامپوننت تشخیص خودکار خطاها
  شامل:
  1. آمار تشخیص
  2. تنظیمات قوانین تشخیص
  3. تنظیمات پیشرفته
  4. جدول خطاهای تشخیص داده شده
  5. عملیات مدیریت خطاها
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
