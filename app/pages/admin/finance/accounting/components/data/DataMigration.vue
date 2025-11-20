<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">انتقال داده‌ها</h4>
        <p class="text-sm text-gray-600 mt-1">انتقال و تبدیل داده‌ها بین سیستم‌های مختلف</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="startMigration"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4" />
          </svg>
          شروع انتقال
        </button>
      </div>
    </div>

    <!-- آمار انتقال -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار انتقال</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کل انتقال‌ها</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ migrationStats.total }}</div>
          <div class="text-xs text-gray-500 mt-1">عملیات</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">انتقال‌های موفق</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ migrationStats.successful }}</div>
          <div class="text-xs text-gray-500 mt-1">عملیات</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">در حال اجرا</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ migrationStats.running }}</div>
          <div class="text-xs text-gray-500 mt-1">عملیات</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">نرخ موفقیت</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ migrationStats.successRate }}%</div>
          <div class="text-xs text-gray-500 mt-1">موفق</div>
        </div>
      </div>
    </div>

    <!-- پروژه‌های انتقال -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">پروژه‌های انتقال</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="project in migrationProjects"
            :key="project.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getProjectStatusColor(project.status)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ project.name }}</h6>
                <p class="text-xs text-gray-500">{{ project.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ project.progress }}%</div>
                <div class="text-xs text-gray-500">پیشرفت</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ project.records }}</div>
                <div class="text-xs text-gray-500">رکورد</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800"
                  @click="viewProject(project)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
                <button
                  class="text-green-600 hover:text-green-800"
                  @click="editProject(project)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  class="text-red-600 hover:text-red-800"
                  @click="deleteProject(project)"
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

    <!-- تاریخچه انتقال -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تاریخچه انتقال</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="history in migrationHistory"
            :key="history.id"
            class="flex items-start space-x-4 space-x-reverse p-6 border border-gray-200 rounded-lg"
          >
            <div
              class="w-2 h-2 rounded-full mt-2"
              :class="getHistoryColor(history.status)"
            ></div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <h6 class="text-sm font-medium text-gray-900">{{ history.projectName }}</h6>
                <span class="text-xs text-gray-500">{{ history.timestamp }}</span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ history.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>رکوردهای انتقال یافته: {{ history.migratedRecords }}</span>
                <span>مدت: {{ history.duration }} دقیقه</span>
                <span>کاربر: {{ history.user }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات انتقال -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات انتقال</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع انتقال</label>
          <select
            v-model="migrationSettings.type"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="full">کامل</option>
            <option value="incremental">افزایشی</option>
            <option value="selective">انتخابی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرمت خروجی</label>
          <select
            v-model="migrationSettings.outputFormat"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="json">JSON</option>
            <option value="xml">XML</option>
            <option value="csv">CSV</option>
            <option value="sql">SQL</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اندازه بسته</label>
          <input
            v-model="migrationSettings.batchSize"
            type="number"
            min="100"
            max="10000"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">رکورد در هر بسته</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تایم‌اوت</label>
          <input
            v-model="migrationSettings.timeout"
            type="number"
            min="30"
            max="3600"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">ثانیه</p>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="migrationSettings.validateData"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعتبارسنجی داده‌ها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="migrationSettings.createBackup"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">ایجاد پشتیبان قبل از انتقال</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="migrationSettings.resumeOnError"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">ادامه در صورت خطا</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="migrationSettings.notifyOnComplete"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعلان پس از تکمیل</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار انتقال
const migrationStats = ref({
  total: 8,
  successful: 7,
  running: 1,
  successRate: 87.5
})

// پروژه‌های انتقال
const migrationProjects = ref([
  {
    id: 1,
    name: 'انتقال تراکنش‌های ۱۴۰۱',
    description: 'انتقال تراکنش‌های سال ۱۴۰۱ به سیستم جدید',
    status: 'running',
    progress: 65,
    records: 15420
  },
  {
    id: 2,
    name: 'انتقال مشتریان',
    description: 'انتقال اطلاعات مشتریان از سیستم قدیمی',
    status: 'completed',
    progress: 100,
    records: 3240
  },
  {
    id: 3,
    name: 'انتقال محصولات',
    description: 'انتقال کاتالوگ محصولات',
    status: 'pending',
    progress: 0,
    records: 1250
  }
])

// تاریخچه انتقال
const migrationHistory = ref([
  {
    id: 1,
    projectName: 'انتقال مشتریان',
    description: 'انتقال موفق اطلاعات مشتریان از سیستم قدیمی',
    migratedRecords: 3240,
    duration: 45,
    status: 'success',
    user: 'مدیر سیستم',
    timestamp: '۲ روز پیش'
  },
  {
    id: 2,
    projectName: 'انتقال تراکنش‌های ۱۴۰۰',
    description: 'انتقال تراکنش‌های سال ۱۴۰۰',
    migratedRecords: 12850,
    duration: 120,
    status: 'success',
    user: 'کاربر حسابداری',
    timestamp: '۱ هفته پیش'
  }
])

// تنظیمات انتقال
const migrationSettings = ref({
  type: 'full',
  outputFormat: 'json',
  batchSize: 1000,
  timeout: 300,
  validateData: true,
  createBackup: true,
  resumeOnError: false,
  notifyOnComplete: true
})

// رنگ‌های وضعیت
const getProjectStatusColor = (status: string) => {
  const colors = {
    running: 'bg-blue-500',
    completed: 'bg-green-500',
    pending: 'bg-yellow-500',
    error: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
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
const startMigration = () => {
  console.log('شروع انتقال')
}

const viewProject = (project: any) => {
  console.log('مشاهده پروژه:', project)
}

const editProject = (project: any) => {
  console.log('ویرایش پروژه:', project)
}

const deleteProject = (project: any) => {
  console.log('حذف پروژه:', project)
}
</script>

<!--
  کامپوننت انتقال داده‌ها
  شامل:
  1. آمار انتقال
  2. پروژه‌های انتقال
  3. تاریخچه انتقال
  4. تنظیمات انتقال
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
