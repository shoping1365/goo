<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">پشتیبان‌گیری و بازیابی</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت پشتیبان‌ها و بازیابی داده‌ها</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="createBackup"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
          </svg>
          ایجاد پشتیبان
        </button>
      </div>
    </div>

    <!-- آمار پشتیبان‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار پشتیبان‌ها</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کل پشتیبان‌ها</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ backupStats.total }}</div>
          <div class="text-xs text-gray-500 mt-1">فایل</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">آخرین پشتیبان</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ backupStats.lastBackup }}</div>
          <div class="text-xs text-gray-500 mt-1">ساعت پیش</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">حجم کل</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ backupStats.totalSize }}</div>
          <div class="text-xs text-gray-500 mt-1">گیگابایت</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">وضعیت</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ backupStats.status }}</div>
          <div class="text-xs text-gray-500 mt-1">سالم</div>
        </div>
      </div>
    </div>

    <!-- لیست پشتیبان‌ها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">پشتیبان‌های موجود</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="backup in backups"
            :key="backup.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getBackupStatusColor(backup.status)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ backup.name }}</h6>
                <p class="text-xs text-gray-500">{{ backup.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ backup.size }}</div>
                <div class="text-xs text-gray-500">حجم</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ backup.date }}</div>
                <div class="text-xs text-gray-500">تاریخ</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  @click="restoreBackup(backup)"
                  class="text-blue-600 hover:text-blue-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                </button>
                <button
                  @click="downloadBackup(backup)"
                  class="text-green-600 hover:text-green-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                </button>
                <button
                  @click="deleteBackup(backup)"
                  class="text-red-600 hover:text-red-800"
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

    <!-- تاریخچه بازیابی -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تاریخچه بازیابی</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="restore in restoreHistory"
            :key="restore.id"
            class="flex items-start space-x-4 space-x-reverse p-6 border border-gray-200 rounded-lg"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getRestoreColor(restore.status)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ restore.backupName }}</h6>
                <span class="text-xs text-gray-500">{{ restore.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ restore.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>وضعیت: {{ getRestoreStatusText(restore.status) }}</span>
                <span>مدت: {{ restore.duration }} دقیقه</span>
                <span>کاربر: {{ restore.user }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پشتیبان‌گیری -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات پشتیبان‌گیری</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس پشتیبان‌گیری</label>
          <select
            v-model="backupSettings.frequency"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="daily">روزانه</option>
            <option value="weekly">هفتگی</option>
            <option value="monthly">ماهانه</option>
            <option value="manual">دستی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">زمان پشتیبان‌گیری</label>
          <input
            v-model="backupSettings.time"
            type="time"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حفظ پشتیبان‌ها</label>
          <select
            v-model="backupSettings.retention"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="7">۷ روز</option>
            <option value="30">۳۰ روز</option>
            <option value="90">۹۰ روز</option>
            <option value="365">۱ سال</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی</label>
          <select
            v-model="backupSettings.compression"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="none">بدون فشرده‌سازی</option>
            <option value="gzip">GZIP</option>
            <option value="zip">ZIP</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="backupSettings.autoBackup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">پشتیبان‌گیری خودکار</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="backupSettings.encryptBackup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">رمزگذاری پشتیبان‌ها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="backupSettings.notifyOnBackup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعلان پس از پشتیبان‌گیری</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار پشتیبان‌ها
const backupStats = ref({
  total: 24,
  lastBackup: 3,
  totalSize: 2.4,
  status: 'سالم'
})

// پشتیبان‌های موجود
const backups = ref([
  {
    id: 1,
    name: 'پشتیبان کامل - ۱۵ دی ۱۴۰۲',
    description: 'پشتیبان کامل از تمام داده‌های حسابداری',
    status: 'success',
    size: '150MB',
    date: '۱۵ دی ۱۴۰۲'
  },
  {
    id: 2,
    name: 'پشتیبان تفاضلی - ۱۴ دی ۱۴۰۲',
    description: 'پشتیبان تفاضلی از تغییرات اخیر',
    status: 'success',
    size: '45MB',
    date: '۱۴ دی ۱۴۰۲'
  },
  {
    id: 3,
    name: 'پشتیبان کامل - ۱۰ دی ۱۴۰۲',
    description: 'پشتیبان کامل هفتگی',
    status: 'warning',
    size: '148MB',
    date: '۱۰ دی ۱۴۰۲'
  }
])

// تاریخچه بازیابی
const restoreHistory = ref([
  {
    id: 1,
    backupName: 'پشتیبان کامل - ۱۵ دی ۱۴۰۲',
    description: 'بازیابی کامل سیستم پس از بروزرسانی',
    status: 'success',
    duration: 15,
    user: 'مدیر سیستم',
    timestamp: '۲ ساعت پیش'
  },
  {
    id: 2,
    backupName: 'پشتیبان تفاضلی - ۱۴ دی ۱۴۰۲',
    description: 'بازیابی بخشی از داده‌ها',
    status: 'success',
    duration: 8,
    user: 'کاربر حسابداری',
    timestamp: '۱ روز پیش'
  }
])

// تنظیمات پشتیبان‌گیری
const backupSettings = ref({
  frequency: 'daily',
  time: '02:00',
  retention: 30,
  compression: 'gzip',
  autoBackup: true,
  encryptBackup: true,
  notifyOnBackup: true
})

// رنگ‌های وضعیت
const getBackupStatusColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getRestoreColor = (status: string) => {
  const colors = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

const getRestoreStatusText = (status: string) => {
  const texts = {
    success: 'موفق',
    warning: 'هشدار',
    error: 'خطا'
  }
  return texts[status] || status
}

// عملیات
const createBackup = () => {
  console.log('ایجاد پشتیبان')
}

const restoreBackup = (backup: any) => {
  console.log('بازیابی پشتیبان:', backup)
}

const downloadBackup = (backup: any) => {
  console.log('دانلود پشتیبان:', backup)
}

const deleteBackup = (backup: any) => {
  console.log('حذف پشتیبان:', backup)
}
</script>

<!--
  کامپوننت پشتیبان‌گیری و بازیابی
  شامل:
  1. آمار پشتیبان‌ها
  2. لیست پشتیبان‌ها
  3. تاریخچه بازیابی
  4. تنظیمات پشتیبان‌گیری
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
