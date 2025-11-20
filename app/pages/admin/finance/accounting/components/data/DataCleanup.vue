<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">پاکسازی داده‌ها</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و پاکسازی داده‌های غیرضروری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="runCleanup"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          اجرای پاکسازی
        </button>
      </div>
    </div>

    <!-- آمار پاکسازی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار پاکسازی</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">داده‌های قابل پاکسازی</h6>
            <div class="w-3 h-3 bg-red-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-red-600">{{ cleanupStats.cleanable }}</div>
          <div class="text-xs text-gray-500 mt-1">رکورد</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">فضای آزاد شده</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ cleanupStats.freedSpace }}</div>
          <div class="text-xs text-gray-500 mt-1">مگابایت</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">آخرین پاکسازی</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ cleanupStats.lastCleanup }}</div>
          <div class="text-xs text-gray-500 mt-1">روز پیش</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">وضعیت</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ cleanupStats.status }}</div>
          <div class="text-xs text-gray-500 mt-1">نیاز به پاکسازی</div>
        </div>
      </div>
    </div>

    <!-- انواع داده‌های قابل پاکسازی -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">انواع داده‌های قابل پاکسازی</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="dataType in dataTypes"
            :key="dataType.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getDataTypeColor(dataType.priority)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ dataType.name }}</h6>
                <p class="text-xs text-gray-500">{{ dataType.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ dataType.count }}</div>
                <div class="text-xs text-gray-500">رکورد</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ dataType.size }}</div>
                <div class="text-xs text-gray-500">حجم</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800"
                  @click="previewData(dataType)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
                <button
                  class="text-red-600 hover:text-red-800"
                  @click="cleanupDataType(dataType)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه پاکسازی -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تاریخچه پاکسازی</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="history in cleanupHistory"
            :key="history.id"
            class="flex items-start space-x-4 space-x-reverse p-6 border border-gray-200 rounded-lg"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getHistoryColor(history.status)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ history.type }}</h6>
                <span class="text-xs text-gray-500">{{ history.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ history.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>رکوردهای پاک شده: {{ history.deletedRecords }}</span>
                <span>فضای آزاد شده: {{ history.freedSpace }}</span>
                <span>کاربر: {{ history.user }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پاکسازی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات پاکسازی</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل سن داده‌ها</label>
          <select
            v-model="cleanupSettings.minAge"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="30">۳۰ روز</option>
            <option value="90">۹۰ روز</option>
            <option value="180">۱۸۰ روز</option>
            <option value="365">۱ سال</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد رکوردها</label>
          <input
            v-model="cleanupSettings.maxRecords"
            type="number"
            min="100"
            max="10000"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس پاکسازی</label>
          <select
            v-model="cleanupSettings.frequency"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="weekly">هفتگی</option>
            <option value="monthly">ماهانه</option>
            <option value="quarterly">فصلی</option>
            <option value="manual">دستی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع پاکسازی</label>
          <select
            v-model="cleanupSettings.cleanupType"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="soft">نرم (Soft Delete)</option>
            <option value="hard">سخت (Hard Delete)</option>
            <option value="archive">آرشیو</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="cleanupSettings.autoCleanup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">پاکسازی خودکار</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="cleanupSettings.confirmBeforeDelete"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">تأیید قبل از حذف</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="cleanupSettings.createBackup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">ایجاد پشتیبان قبل از پاکسازی</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="cleanupSettings.notifyOnCleanup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعلان پس از پاکسازی</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار پاکسازی
const cleanupStats = ref({
  cleanable: 1247,
  freedSpace: 45.2,
  lastCleanup: 7,
  status: 'نیاز به پاکسازی'
})

// انواع داده‌های قابل پاکسازی
const dataTypes = ref([
  {
    id: 1,
    name: 'لاگ‌های قدیمی',
    description: 'لاگ‌های سیستم با سن بیش از ۹۰ روز',
    priority: 'high',
    count: 456,
    size: '12.5MB'
  },
  {
    id: 2,
    name: 'داده‌های تکراری',
    description: 'رکوردهای تکراری در پایگاه داده',
    priority: 'medium',
    count: 234,
    size: '8.3MB'
  },
  {
    id: 3,
    name: 'فایل‌های موقت',
    description: 'فایل‌های موقت و کش‌های سیستم',
    priority: 'low',
    count: 567,
    size: '24.4MB'
  }
])

// تاریخچه پاکسازی
const cleanupHistory = ref([
  {
    id: 1,
    type: 'پاکسازی لاگ‌های قدیمی',
    description: 'پاکسازی لاگ‌های سیستم با سن بیش از ۹۰ روز',
    deletedRecords: 456,
    freedSpace: '12.5MB',
    status: 'success',
    user: 'سیستم',
    timestamp: '۱ هفته پیش'
  },
  {
    id: 2,
    type: 'پاکسازی داده‌های تکراری',
    description: 'حذف رکوردهای تکراری از پایگاه داده',
    deletedRecords: 234,
    freedSpace: '8.3MB',
    status: 'success',
    user: 'مدیر سیستم',
    timestamp: '۲ هفته پیش'
  }
])

// تنظیمات پاکسازی
const cleanupSettings = ref({
  minAge: 90,
  maxRecords: 1000,
  frequency: 'weekly',
  cleanupType: 'soft',
  autoCleanup: true,
  confirmBeforeDelete: true,
  createBackup: true,
  notifyOnCleanup: true
})

// رنگ‌های اولویت
const getDataTypeColor = (priority: string) => {
  const colors = {
    high: 'bg-red-500',
    medium: 'bg-yellow-500',
    low: 'bg-blue-500'
  }
  return colors[priority] || 'bg-gray-500'
}

const getHistoryColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// عملیات
const runCleanup = () => {
  console.log('اجرای پاکسازی')
}

const previewData = (dataType: any) => {
  console.log('پیش‌نمایش داده:', dataType)
}

const cleanupDataType = (dataType: any) => {
  console.log('پاکسازی نوع داده:', dataType)
}
</script>

<!--
  کامپوننت پاکسازی داده‌ها
  شامل:
  1. آمار پاکسازی
  2. انواع داده‌های قابل پاکسازی
  3. تاریخچه پاکسازی
  4. تنظیمات پاکسازی
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
