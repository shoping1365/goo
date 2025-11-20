<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">بازگشت به نسخه قبلی</h4>
        <p class="text-sm text-gray-600 mt-1">انتخاب و بازگشت به نسخه‌های قبلی اتصال</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="createBackupBeforeRollback"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          ایجاد نسخه پشتیبان
        </button>
      </div>
    </div>

    <!-- هشدار مهم -->
    <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg class="h-5 w-5 text-yellow-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
          </svg>
        </div>
        <div class="mr-3">
          <h3 class="text-sm font-medium text-yellow-800">هشدار مهم</h3>
          <div class="mt-2 text-sm text-yellow-700">
            <p>بازگشت به نسخه قبلی ممکن است باعث از دست رفتن تغییرات اخیر شود. حتماً قبل از بازگشت، نسخه پشتیبان ایجاد کنید.</p>
          </div>
        </div>
      </div>
    </div>

    <!-- انتخاب نسخه -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">انتخاب نسخه برای بازگشت</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نسخه فعلی</label>
          <div class="bg-gray-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-sm font-medium text-gray-900">{{ currentVersion.version }}</div>
                <div class="text-sm text-gray-500">{{ currentVersion.connectionName }}</div>
                <div class="text-xs text-gray-400 mt-1">{{ currentVersion.createdAt }}</div>
              </div>
              <span class="px-2 py-1 text-xs font-medium rounded-full bg-green-100 text-green-700">
                فعلی
              </span>
            </div>
            <div class="mt-3 text-sm text-gray-600">
              {{ currentVersion.description }}
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نسخه هدف</label>
          <select
            v-model="selectedTargetVersion"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب نسخه...</option>
            <option
              v-for="version in availableVersions"
              :key="version.id"
              :value="version.id"
            >
              {{ version.version }} - {{ version.createdAt }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <!-- پیش‌نمایش تغییرات -->
    <div v-if="selectedTargetVersion" class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">پیش‌نمایش تغییرات</h5>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-red-50 rounded-lg p-6">
            <h6 class="text-sm font-medium text-red-800 mb-2">حذف خواهد شد</h6>
            <ul class="text-sm text-red-700 space-y-1">
              <li v-for="(item, index) in changesToRemove" :key="index" class="flex items-center">
                <svg class="w-4 h-4 ml-2 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
                {{ item }}
              </li>
            </ul>
          </div>

          <div class="bg-yellow-50 rounded-lg p-6">
            <h6 class="text-sm font-medium text-yellow-800 mb-2">تغییر خواهد یافت</h6>
            <ul class="text-sm text-yellow-700 space-y-1">
              <li v-for="(item, index) in changesToModify" :key="index" class="flex items-center">
                <svg class="w-4 h-4 ml-2 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
                {{ item }}
              </li>
            </ul>
          </div>

          <div class="bg-green-50 rounded-lg p-6">
            <h6 class="text-sm font-medium text-green-800 mb-2">اضافه خواهد شد</h6>
            <ul class="text-sm text-green-700 space-y-1">
              <li v-for="(item, index) in changesToAdd" :key="index" class="flex items-center">
                <svg class="w-4 h-4 ml-2 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
                {{ item }}
              </li>
            </ul>
          </div>
        </div>

        <!-- جزئیات تغییرات -->
        <div class="border-t border-gray-200 pt-4">
          <h6 class="text-sm font-medium text-gray-900 mb-3">جزئیات تغییرات</h6>
          <div class="bg-gray-50 rounded-lg p-6">
            <div class="space-y-3">
              <div v-for="(change, index) in detailedChanges" :key="index" class="flex items-start space-x-3 space-x-reverse">
                <div class="flex-shrink-0 w-2 h-2 rounded-full mt-2" :class="getChangeTypeClass(change.type)"></div>
                <div class="flex-1">
                  <div class="text-sm font-medium text-gray-900">{{ change.title }}</div>
                  <div class="text-sm text-gray-600">{{ change.description }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات بازگشت -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات بازگشت</h5>
      <div class="space-y-4">
        <div class="flex items-center">
          <input
            id="createBackup"
            v-model="createBackup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label for="createBackup" class="mr-3 text-sm text-gray-700">
            ایجاد نسخه پشتیبان قبل از بازگشت
          </label>
        </div>

        <div class="flex items-center">
          <input
            id="notifyUsers"
            v-model="notifyUsers"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label for="notifyUsers" class="mr-3 text-sm text-gray-700">
            اطلاع‌رسانی به کاربران
          </label>
        </div>

        <div class="flex items-center">
          <input
            id="validateBeforeRollback"
            v-model="validateBeforeRollback"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label for="validateBeforeRollback" class="mr-3 text-sm text-gray-700">
            اعتبارسنجی قبل از بازگشت
          </label>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دلیل بازگشت</label>
          <textarea
            v-model="rollbackReason"
            rows="3"
            placeholder="دلیل بازگشت به نسخه قبلی را وارد کنید..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          ></textarea>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex items-center justify-end space-x-3 space-x-reverse">
      <button
        class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        @click="cancelRollback"
      >
        انصراف
      </button>
      <button
        :disabled="!selectedTargetVersion"
        class="px-4 py-2 border border-transparent rounded-md text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        @click="previewRollback"
      >
        پیش‌نمایش بازگشت
      </button>
      <button
        :disabled="!selectedTargetVersion || !rollbackReason"
        class="px-4 py-2 border border-transparent rounded-md text-sm font-medium text-white bg-red-600 hover:bg-red-700 disabled:bg-gray-400 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
        @click="confirmRollback"
      >
        تایید بازگشت
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// انتخاب نسخه
const selectedTargetVersion = ref('')
const rollbackReason = ref('')

// تنظیمات
const createBackup = ref(true)
const notifyUsers = ref(true)
const validateBeforeRollback = ref(true)

// نسخه فعلی
const currentVersion = ref({
  version: 'v2.1.0',
  connectionName: 'اتصال هلو',
  createdAt: '۱۴۰۲/۱۰/۱۵ - ۱۴:۳۰',
  description: 'نسخه پایدار با بهبودهای عملکرد'
})

// نسخه‌های موجود
const availableVersions = ref([
  {
    id: 1,
    version: 'v2.0.5',
    createdAt: '۱۴۰۲/۱۰/۱۰ - ۰۹:۱۵',
    description: 'نسخه خودکار قبل از تغییرات'
  },
  {
    id: 2,
    version: 'v2.0.0',
    createdAt: '۱۴۰۲/۰۹/۲۸ - ۱۶:۴۵',
    description: 'نسخه اصلی با قابلیت‌های جدید'
  },
  {
    id: 3,
    version: 'v1.9.8',
    createdAt: '۱۴۰۲/۰۹/۱۵ - ۱۱:۲۰',
    description: 'آخرین نسخه از سری ۱.x'
  }
])

// تغییرات
const changesToRemove = ref([
  'تنظیمات جدید امنیتی',
  'بهبودهای عملکرد',
  'قابلیت‌های اضافی'
])

const changesToModify = ref([
  'پیکربندی اتصال',
  'تنظیمات همگام‌سازی',
  'محدودیت‌های دسترسی'
])

const changesToAdd = ref([
  'تنظیمات قدیمی',
  'قابلیت‌های پایدار',
  'پیکربندی ساده'
])

const detailedChanges = ref([
  {
    type: 'remove',
    title: 'حذف تنظیمات امنیتی جدید',
    description: 'تنظیمات امنیتی اضافه شده در نسخه ۲.۱.۰ حذف خواهد شد'
  },
  {
    type: 'modify',
    title: 'تغییر پیکربندی اتصال',
    description: 'پیکربندی اتصال به حالت قبلی بازگردانده خواهد شد'
  },
  {
    type: 'add',
    title: 'بازگردانی تنظیمات قدیمی',
    description: 'تنظیمات پایدار نسخه قبلی بازگردانده خواهد شد'
  }
])

// کلاس نوع تغییر
const getChangeTypeClass = (type: string) => {
  const classes = {
    remove: 'bg-red-500',
    modify: 'bg-yellow-500',
    add: 'bg-green-500'
  }
  return classes[type] || 'bg-gray-500'
}

// عملیات
const createBackupBeforeRollback = () => {
  console.log('ایجاد نسخه پشتیبان قبل از بازگشت')
}

const previewRollback = () => {
  console.log('پیش‌نمایش بازگشت به نسخه:', selectedTargetVersion.value)
}

const confirmRollback = () => {
  console.log('تایید بازگشت به نسخه:', selectedTargetVersion.value)
  console.log('دلیل:', rollbackReason.value)
}

const cancelRollback = () => {
  selectedTargetVersion.value = ''
  rollbackReason.value = ''
}
</script>

<!--
  کامپوننت بازگشت به نسخه قبلی
  شامل:
  1. انتخاب نسخه هدف
  2. پیش‌نمایش تغییرات
  3. تنظیمات بازگشت
  4. تایید عملیات
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
