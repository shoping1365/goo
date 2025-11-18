<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">لاگ فعالیت‌های اتصال</h4>
        <p class="text-sm text-gray-600 mt-1">ثبت و مشاهده تمام فعالیت‌های مربوط به اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="showFilters = !showFilters"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z" />
          </svg>
          فیلترها
        </button>
        <button
          @click="exportLogs"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          صادر کردن لاگ
        </button>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div v-if="showFilters" class="bg-gray-50 rounded-lg p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلترهای لاگ</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سطح لاگ</label>
          <select
            v-model="filters.level"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه سطوح</option>
            <option value="info">اطلاعات</option>
            <option value="warning">هشدار</option>
            <option value="error">خطا</option>
            <option value="debug">دیباگ</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع فعالیت</label>
          <select
            v-model="filters.activityType"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه انواع</option>
            <option value="connection">اتصال</option>
            <option value="sync">همگام‌سازی</option>
            <option value="auth">احراز هویت</option>
            <option value="data">داده</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نرم‌افزار</label>
          <select
            v-model="filters.software"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه نرم‌افزارها</option>
            <option value="helo">هلو</option>
            <option value="parsian">پارسیان</option>
            <option value="sepidar">سپیدار</option>
            <option value="ghias">قیاس</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ</label>
          <input
            v-model="filters.date"
            type="date"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
      </div>
    </div>

    <!-- آمار لاگ -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ stats.totalLogs }}</div>
            <div class="text-sm text-blue-700">کل لاگ‌ها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ stats.todayLogs }}</div>
            <div class="text-sm text-green-700">لاگ‌های امروز</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ stats.warningLogs }}</div>
            <div class="text-sm text-yellow-700">لاگ‌های هشدار</div>
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
            <div class="text-2xl font-bold text-red-600">{{ stats.errorLogs }}</div>
            <div class="text-sm text-red-700">لاگ‌های خطا</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول لاگ‌ها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">فعالیت‌های اتصال</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">زمان</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">سطح</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع فعالیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">پیام</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">کاربر</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">جزئیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="log in filteredLogs" :key="log.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ log.time }}</div>
                <div class="text-xs text-gray-500">{{ log.date }}</div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getLogLevelClass(log.level)"
                >
                  {{ getLogLevelLabel(log.level) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getActivityTypeClass(log.activityType)"
                >
                  {{ getActivityTypeLabel(log.activityType) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ log.message }}</div>
                <div class="text-xs text-gray-500">{{ log.details }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <img :src="log.softwareLogo" :alt="log.softwareName" class="w-6 h-6 rounded" />
                  <span class="text-sm text-gray-900">{{ log.softwareName }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ log.user }}</div>
                <div class="text-xs text-gray-500">{{ log.userRole }}</div>
              </td>
              <td class="py-3 px-4">
                <button
                  @click="viewLogDetails(log)"
                  class="p-1 text-blue-600 hover:text-blue-800"
                  title="مشاهده جزئیات"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودارهای تحلیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">توزیع لاگ‌ها بر اساس سطح</h5>
        <div class="space-y-4">
          <div v-for="level in logLevelDistribution" :key="level.level" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :class="getLogLevelColor(level.level)"></div>
              <span class="text-sm text-gray-700">{{ getLogLevelLabel(level.level) }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="getLogLevelColor(level.level)"
                  :style="{ width: level.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ level.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">فعالیت‌ها بر اساس نرم‌افزار</h5>
        <div class="space-y-4">
          <div v-for="software in softwareActivityDistribution" :key="software.name" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <img :src="software.logo" :alt="software.name" class="w-6 h-6 rounded" />
              <span class="text-sm text-gray-700">{{ software.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full bg-blue-500"
                  :style="{ width: software.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ software.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- صفحه‌بندی -->
    <div class="flex items-center justify-between">
      <div class="text-sm text-gray-700">
        نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} لاگ
      </div>
      <div class="flex gap-2">
        <button
          @click="previousPage"
          :disabled="pagination.currentPage === 1"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
        >
          قبلی
        </button>
        <span class="px-3 py-1 text-sm text-gray-700">
          صفحه {{ pagination.currentPage }} از {{ pagination.totalPages }}
        </span>
        <button
          @click="nextPage"
          :disabled="pagination.currentPage === pagination.totalPages"
          class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
        >
          بعدی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// وضعیت فیلترها
const showFilters = ref(false)

// فیلترها
const filters = ref({
  level: '',
  activityType: '',
  software: '',
  date: ''
})

// آمار
const stats = ref({
  totalLogs: 25480,
  todayLogs: 1250,
  warningLogs: 180,
  errorLogs: 45
})

// صفحه‌بندی
const pagination = ref({
  currentPage: 1,
  totalPages: 100,
  total: 25480,
  start: 1,
  end: 20,
  perPage: 20
})

// لاگ‌ها
const logs = ref([
  {
    id: 1,
    time: '۱۴:۳۰:۲۵',
    date: 'امروز',
    level: 'info',
    activityType: 'connection',
    message: 'اتصال موفق به سرور هلو',
    details: 'Connection established successfully',
    softwareName: 'هلو',
    softwareLogo: '/images/helo-logo.png',
    user: 'علی احمدی',
    userRole: 'مدیر سیستم'
  },
  {
    id: 2,
    time: '۱۴:۲۵:۱۸',
    date: 'امروز',
    level: 'warning',
    activityType: 'sync',
    message: 'همگام‌سازی کند انجام شد',
    details: 'Sync completed with delays',
    softwareName: 'پارسیان',
    softwareLogo: '/images/parsian-logo.png',
    user: 'سیستم',
    userRole: 'اتوماتیک'
  },
  {
    id: 3,
    time: '۱۴:۲۰:۴۵',
    date: 'امروز',
    level: 'error',
    activityType: 'auth',
    message: 'خطا در احراز هویت',
    details: 'Authentication failed',
    softwareName: 'سپیدار',
    softwareLogo: '/images/sepidar-logo.png',
    user: 'فاطمه محمدی',
    userRole: 'کاربر'
  }
])

// لاگ‌های فیلتر شده
const filteredLogs = computed(() => {
  return logs.value.filter(log => {
    if (filters.value.level && log.level !== filters.value.level) return false
    if (filters.value.activityType && log.activityType !== filters.value.activityType) return false
    if (filters.value.software && log.softwareName.toLowerCase() !== filters.value.software) return false
    // TODO: فیلتر تاریخ
    return true
  })
})

// توزیع سطح لاگ
const logLevelDistribution = ref([
  { level: 'info', percentage: 70 },
  { level: 'warning', percentage: 20 },
  { level: 'error', percentage: 8 },
  { level: 'debug', percentage: 2 }
])

// توزیع فعالیت بر اساس نرم‌افزار
const softwareActivityDistribution = ref([
  { name: 'هلو', logo: '/images/helo-logo.png', percentage: 40 },
  { name: 'پارسیان', logo: '/images/parsian-logo.png', percentage: 30 },
  { name: 'سپیدار', logo: '/images/sepidar-logo.png', percentage: 20 },
  { name: 'قیاس', logo: '/images/ghias-logo.png', percentage: 10 }
])

// کلاس سطح لاگ
const getLogLevelClass = (level: string) => {
  const classes = {
    info: 'bg-blue-100 text-blue-700',
    warning: 'bg-yellow-100 text-yellow-700',
    error: 'bg-red-100 text-red-700',
    debug: 'bg-gray-100 text-gray-700'
  }
  return classes[level] || 'bg-gray-100 text-gray-700'
}

// رنگ سطح لاگ
const getLogLevelColor = (level: string) => {
  const colors = {
    info: 'bg-blue-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500',
    debug: 'bg-gray-500'
  }
  return colors[level] || 'bg-gray-500'
}

// برچسب سطح لاگ
const getLogLevelLabel = (level: string) => {
  const labels = {
    info: 'اطلاعات',
    warning: 'هشدار',
    error: 'خطا',
    debug: 'دیباگ'
  }
  return labels[level] || level
}

// کلاس نوع فعالیت
const getActivityTypeClass = (type: string) => {
  const classes = {
    connection: 'bg-green-100 text-green-700',
    sync: 'bg-blue-100 text-blue-700',
    auth: 'bg-purple-100 text-purple-700',
    data: 'bg-orange-100 text-orange-700'
  }
  return classes[type] || 'bg-gray-100 text-gray-700'
}

// برچسب نوع فعالیت
const getActivityTypeLabel = (type: string) => {
  const labels = {
    connection: 'اتصال',
    sync: 'همگام‌سازی',
    auth: 'احراز هویت',
    data: 'داده'
  }
  return labels[type] || type
}

// اعمال فیلترها
const applyFilters = () => {
  // TODO: اعمال فیلترها
  console.log('فیلترهای لاگ اعمال شد:', filters.value)
}

// صادر کردن لاگ‌ها
const exportLogs = () => {
  // TODO: صادر کردن لاگ‌ها
  console.log('صادر کردن لاگ‌ها')
}

// مشاهده جزئیات لاگ
const viewLogDetails = (log: any) => {
  // TODO: مشاهده جزئیات لاگ
  console.log('مشاهده جزئیات لاگ:', log)
}

// صفحه قبلی
const previousPage = () => {
  if (pagination.value.currentPage > 1) {
    pagination.value.currentPage--
    // TODO: بارگذاری داده‌های صفحه جدید
  }
}

// صفحه بعدی
const nextPage = () => {
  if (pagination.value.currentPage < pagination.value.totalPages) {
    pagination.value.currentPage++
    // TODO: بارگذاری داده‌های صفحه جدید
  }
}
</script>

<!--
  کامپوننت لاگ فعالیت‌های اتصال
  شامل:
  1. آمار لاگ‌ها
  2. فیلترهای پیشرفته
  3. جدول لاگ‌ها
  4. نمودارهای تحلیلی
  5. صفحه‌بندی
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
