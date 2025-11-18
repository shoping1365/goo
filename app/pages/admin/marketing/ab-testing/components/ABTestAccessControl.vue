<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">کنترل‌های مدیریتی</h3>
      <button 
        @click="refreshLogs" 
        class="px-3 py-2 text-sm bg-blue-50 text-blue-700 rounded-lg hover:bg-blue-100 flex items-center"
      >
        <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
        </svg>
        به‌روزرسانی
      </button>
    </div>

    <!-- مدیریت دسترسی -->
    <div class="mb-8">
      <h4 class="text-md font-medium text-gray-900 mb-4">مدیریت دسترسی</h4>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- مجوزهای کاربران -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-700 mb-3">مجوزهای کاربران</h5>
          
          <div class="space-y-3">
            <div v-for="user in users" :key="user.id" class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-gray-900">{{ user.name }}</p>
                <p class="text-xs text-gray-500">{{ user.role }}</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <label class="flex items-center">
                  <input 
                    v-model="user.permissions.createTest" 
                    type="checkbox" 
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="text-xs text-gray-600 mr-1">ایجاد</span>
                </label>
                <label class="flex items-center">
                  <input 
                    v-model="user.permissions.viewResults" 
                    type="checkbox" 
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="text-xs text-gray-600 mr-1">مشاهده نتایج</span>
                </label>
                <label class="flex items-center">
                  <input 
                    v-model="user.permissions.editTest" 
                    type="checkbox" 
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="text-xs text-gray-600 mr-1">ویرایش</span>
                </label>
              </div>
            </div>
          </div>
          
          <button 
            @click="saveUserPermissions" 
            class="mt-4 w-full px-3 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
          >
            ذخیره مجوزها
          </button>
        </div>

        <!-- محدودیت‌های سیستم -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h5 class="text-sm font-medium text-gray-700 mb-3">محدودیت‌های سیستم</h5>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر تست‌های فعال</label>
              <input 
                v-model="systemLimits.maxActiveTests" 
                type="number" 
                min="1" 
                max="100"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر مدت تست (روز)</label>
              <input 
                v-model="systemLimits.maxTestDuration" 
                type="number" 
                min="1" 
                max="365"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">حداقل نمونه‌گیری</label>
              <input 
                v-model="systemLimits.minSampleSize" 
                type="number" 
                min="100" 
                step="100"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
          
          <button 
            @click="saveSystemLimits" 
            class="mt-4 w-full px-3 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700"
          >
            ذخیره محدودیت‌ها
          </button>
        </div>
      </div>
    </div>

    <!-- لاگ‌گیری -->
    <div>
      <h4 class="text-md font-medium text-gray-900 mb-4">لاگ‌گیری سیستم</h4>
      
      <!-- فیلترهای لاگ -->
      <div class="mb-4 flex flex-wrap gap-2">
        <button 
          v-for="filter in logFilters" 
          :key="filter.value"
          @click="setLogFilter(filter.value)"
          :class="[
            'px-3 py-1 text-xs rounded-full',
            currentLogFilter === filter.value 
              ? 'bg-blue-100 text-blue-800' 
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
        >
          {{ filter.label }}
        </button>
      </div>

      <!-- جدول لاگ‌ها -->
      <div class="bg-gray-50 rounded-lg overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-100">
              <tr>
                <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان</th>
                <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
                <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">جزئیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="log in filteredLogs" :key="log.id" class="hover:bg-gray-50">
                <td class="px-4 py-3 text-sm text-gray-900">
                  {{ formatDate(log.timestamp) }}
                </td>
                <td class="px-4 py-3">
                  <span :class="getLogTypeClass(log.type)" class="inline-flex px-2 py-1 text-xs font-medium rounded-full">
                    {{ getLogTypeLabel(log.type) }}
                  </span>
                </td>
                <td class="px-4 py-3 text-sm text-gray-900">{{ log.user }}</td>
                <td class="px-4 py-3 text-sm text-gray-900">{{ log.action }}</td>
                <td class="px-4 py-3 text-sm text-gray-500">
                  <button 
                    @click="showLogDetails(log)"
                    class="text-blue-600 hover:text-blue-800 text-xs"
                  >
                    مشاهده جزئیات
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- صفحه‌بندی -->
      <div class="mt-4 flex items-center justify-between">
        <div class="text-sm text-gray-700">
          نمایش {{ (currentPage - 1) * pageSize + 1 }} تا {{ Math.min(currentPage * pageSize, totalLogs) }} از {{ totalLogs }} لاگ
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button 
            @click="previousPage"
            :disabled="currentPage === 1"
            :class="[
              'px-3 py-1 text-sm rounded-lg',
              currentPage === 1 
                ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                : 'bg-blue-50 text-blue-700 hover:bg-blue-100'
            ]"
          >
            قبلی
          </button>
          <span class="text-sm text-gray-700">{{ currentPage }} از {{ totalPages }}</span>
          <button 
            @click="nextPage"
            :disabled="currentPage === totalPages"
            :class="[
              'px-3 py-1 text-sm rounded-lg',
              currentPage === totalPages 
                ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                : 'bg-blue-50 text-blue-700 hover:bg-blue-100'
            ]"
          >
            بعدی
          </button>
        </div>
      </div>
    </div>

    <!-- Modal جزئیات لاگ -->
    <div v-if="showLogModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-2xl w-full mx-4 max-h-96 overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">جزئیات لاگ</h3>
          <button @click="closeLogModal" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <div v-if="selectedLog" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">زمان:</label>
            <p class="text-sm text-gray-900">{{ formatDate(selectedLog.timestamp) }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">نوع:</label>
            <p class="text-sm text-gray-900">{{ getLogTypeLabel(selectedLog.type) }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">کاربر:</label>
            <p class="text-sm text-gray-900">{{ selectedLog.user }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">عملیات:</label>
            <p class="text-sm text-gray-900">{{ selectedLog.action }}</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">جزئیات:</label>
            <pre class="text-sm text-gray-900 bg-gray-50 p-3 rounded-lg overflow-x-auto">{{ JSON.stringify(selectedLog.details, null, 2) }}</pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
// Props
interface Props {
  modelValue?: any
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: () => ({})
})

// Emits
const emit = defineEmits<{
  'update:modelValue': [value: any]
  'permissions-changed': [permissions: any]
  'limits-changed': [limits: any]
}>()

// State
const currentLogFilter = ref('all')
const currentPage = ref(1)
const pageSize = 10
const showLogModal = ref(false)
const selectedLog = ref(null)

// فیلترهای لاگ
const logFilters = [
  { label: 'همه', value: 'all' },
  { label: 'تغییرات', value: 'change' },
  { label: 'دسترسی‌ها', value: 'access' },
  { label: 'خطاها', value: 'error' }
]

// کاربران و مجوزها
const users = ref([
  {
    id: 1,
    name: 'مدیر سیستم',
    role: 'مدیر',
    permissions: {
      createTest: true,
      viewResults: true,
      editTest: true
    }
  },
  {
    id: 2,
    name: 'کارشناس بازاریابی',
    role: 'بازاریابی',
    permissions: {
      createTest: true,
      viewResults: true,
      editTest: false
    }
  },
  {
    id: 3,
    name: 'توسعه‌دهنده',
    role: 'توسعه‌دهنده',
    permissions: {
      createTest: false,
      viewResults: true,
      editTest: false
    }
  }
])

// محدودیت‌های سیستم
const systemLimits = ref({
  maxActiveTests: 10,
  maxTestDuration: 30,
  minSampleSize: 1000
})

// لاگ‌های سیستم
const logs = ref([
  {
    id: 1,
    timestamp: new Date('2024-01-15T10:30:00'),
    type: 'change',
    user: 'مدیر سیستم',
    action: 'ایجاد تست جدید',
    details: { testId: 123, testName: 'تست صفحه اصلی' }
  },
  {
    id: 2,
    timestamp: new Date('2024-01-15T09:15:00'),
    type: 'access',
    user: 'کارشناس بازاریابی',
    action: 'مشاهده نتایج',
    details: { testId: 122, results: 'viewed' }
  },
  {
    id: 3,
    timestamp: new Date('2024-01-15T08:45:00'),
    type: 'error',
    user: 'سیستم',
    action: 'خطا در محاسبه نتایج',
    details: { error: 'Database connection failed', testId: 121 }
  },
  {
    id: 4,
    timestamp: new Date('2024-01-14T16:20:00'),
    type: 'change',
    user: 'مدیر سیستم',
    action: 'تغییر تنظیمات',
    details: { setting: 'maxActiveTests', oldValue: 5, newValue: 10 }
  }
])

// محاسبات
const totalLogs = computed(() => logs.value.length)
const totalPages = computed(() => Math.ceil(totalLogs.value / pageSize))

const filteredLogs = computed(() => {
  let filtered = logs.value
  
  if (currentLogFilter.value !== 'all') {
    filtered = filtered.filter(log => log.type === currentLogFilter.value)
  }
  
  const start = (currentPage.value - 1) * pageSize
  const end = start + pageSize
  
  return filtered.slice(start, end)
})

// توابع
const setLogFilter = (filter: string) => {
  currentLogFilter.value = filter
  currentPage.value = 1
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const refreshLogs = () => {
  // در حالت واقعی، لاگ‌ها را از API دریافت می‌کنیم
  console.log('به‌روزرسانی لاگ‌ها')
}

const saveUserPermissions = () => {
  emit('permissions-changed', users.value)
  console.log('مجوزهای کاربران ذخیره شد:', users.value)
}

const saveSystemLimits = () => {
  emit('limits-changed', systemLimits.value)
  console.log('محدودیت‌های سیستم ذخیره شد:', systemLimits.value)
}

const showLogDetails = (log: any) => {
  selectedLog.value = log
  showLogModal.value = true
}

const closeLogModal = () => {
  showLogModal.value = false
  selectedLog.value = null
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

const getLogTypeLabel = (type: string) => {
  const labels: any = {
    change: 'تغییر',
    access: 'دسترسی',
    error: 'خطا'
  }
  return labels[type] || type
}

const getLogTypeClass = (type: string) => {
  const classes: any = {
    change: 'bg-blue-100 text-blue-800',
    access: 'bg-green-100 text-green-800',
    error: 'bg-red-100 text-red-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}
</script> 
