<template>
  <div>
    <div class="mb-6">
      <h2 class="text-xl font-semibold text-gray-900 mb-2">تاریخچه ورود</h2>
      <p class="text-gray-600">مشاهده و تحلیل تلاش‌های ورود کاربران</p>
    </div>

    <!-- Filters -->
    <div class="bg-white px-4 py-4 rounded-lg shadow mb-6">
      <div class="grid grid-cols-1 md:grid-cols-5 gapx-4 py-4">
        <div>
          <input
            v-model="filters.search"
            type="text"
            placeholder="جستجو بر اساس موبایل یا یوزرنیم..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <select
            v-model="filters.success"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه نتایج</option>
            <option value="true">موفق</option>
            <option value="false">ناموفق</option>
          </select>
        </div>
        <div>
          <select
            v-model="filters.method"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه روش‌ها</option>
            <option value="otp">OTP</option>
            <option value="password">پسورد</option>
          </select>
        </div>
        <div>
          <input
            v-model="filters.dateFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <button
            @click="loadHistory"
            class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
          >
            جستجو
          </button>
        </div>
      </div>
    </div>

    <!-- Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4 mb-6">
      <div class="bg-white px-4 py-4 rounded-lg shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">ورودهای موفق</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.successfulLogins }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white px-4 py-4 rounded-lg shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">ورودهای ناموفق</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.failedLogins }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white px-4 py-4 rounded-lg shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">نرخ موفقیت</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.successRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white px-4 py-4 rounded-lg shadow">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">فعالیت مشکوک</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.suspiciousActivity }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- History Table -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کاربر
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                روش ورود
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نتیجه
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                آدرس IP
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مرورگر
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ و زمان
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="attempt in history" :key="attempt.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <div class="h-8 w-8 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-xs font-medium text-gray-700">
                        {{ attempt.mobile ? attempt.mobile.charAt(0) : attempt.username ? attempt.username.charAt(0) : '?' }}
                      </span>
                    </div>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">
                      {{ attempt.mobile || attempt.username || 'نامشخص' }}
                    </div>
                    <div v-if="attempt.user_id" class="text-xs text-gray-500">
                      ID: {{ attempt.user_id }}
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="[
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                    attempt.method === 'otp' ? 'bg-blue-100 text-blue-800' : 'bg-green-100 text-green-800'
                  ]"
                >
                  {{ getMethodLabel(attempt.method) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="[
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                    attempt.success ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                  ]"
                >
                  {{ attempt.success ? 'موفق' : 'ناموفق' }}
                </span>
                <div v-if="!attempt.success && attempt.reason" class="text-xs text-gray-500 mt-1">
                  {{ attempt.reason }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ attempt.ip_address }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <div class="max-w-xs truncate" :title="attempt.user_agent">
                  {{ getBrowserInfo(attempt.user_agent) }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDateTime(attempt.created_at) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button
                  @click="viewAttemptDetails(attempt)"
                  class="text-blue-600 hover:text-blue-900"
                >
                  جزئیات
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <button
            @click="previousPage"
            :disabled="currentPage === 1"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
          >
            قبلی
          </button>
          <button
            @click="nextPage"
            :disabled="currentPage >= totalPages"
            class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
          >
            بعدی
          </button>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              نمایش
              <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span>
              تا
              <span class="font-medium">{{ Math.min(currentPage * pageSize, total) }}</span>
              از
              <span class="font-medium">{{ total }}</span>
              نتیجه
            </p>
          </div>
          <div>
            <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
              <button
                @click="previousPage"
                :disabled="currentPage === 1"
                class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
              >
                قبلی
              </button>
              <button
                v-for="page in visiblePages"
                :key="page"
                @click="goToPage(page)"
                :class="[
                  'relative inline-flex items-center px-4 py-2 border text-sm font-medium',
                  page === currentPage
                    ? 'z-10 bg-blue-50 border-blue-500 text-blue-600'
                    : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                ]"
              >
                {{ page }}
              </button>
              <button
                @click="nextPage"
                :disabled="currentPage >= totalPages"
                class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
              >
                بعدی
              </button>
            </nav>
          </div>
        </div>
      </div>
    </div>

    <!-- Attempt Details Modal -->
    <div v-if="selectedAttempt" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-medium text-gray-900">جزئیات تلاش ورود</h3>
            <button
              @click="selectedAttempt = null"
              class="text-gray-400 hover:text-gray-600"
            >
              ✕
            </button>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4 mb-6">
            <div>
              <label class="block text-sm font-medium text-gray-700">کاربر</label>
              <p class="mt-1 text-sm text-gray-900">
                {{ selectedAttempt.mobile || selectedAttempt.username || 'نامشخص' }}
              </p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">روش ورود</label>
              <p class="mt-1 text-sm text-gray-900">{{ getMethodLabel(selectedAttempt.method) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">نتیجه</label>
              <p class="mt-1 text-sm text-gray-900">
                <span :class="selectedAttempt.success ? 'text-green-600' : 'text-red-600'">
                  {{ selectedAttempt.success ? 'موفق' : 'ناموفق' }}
                </span>
              </p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">آدرس IP</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedAttempt.ip_address }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ و زمان</label>
              <p class="mt-1 text-sm text-gray-900">{{ formatDateTime(selectedAttempt.created_at) }}</p>
            </div>
            <div v-if="!selectedAttempt.success">
              <label class="block text-sm font-medium text-gray-700">دلیل شکست</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedAttempt.reason || 'تنظیم نشده' }}</p>
            </div>
          </div>

          <div v-if="selectedAttempt.user_agent" class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">اطلاعات مرورگر</label>
            <div class="bg-gray-50 p-3 rounded-md">
              <p class="text-sm text-gray-900 break-all">{{ selectedAttempt.user_agent }}</p>
            </div>
          </div>

          <div class="flex justify-end">
            <button
              @click="selectedAttempt = null"
              class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
            >
              بستن
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'

// Reactive data
const history = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(20)
const selectedAttempt = ref(null)

const filters = reactive({
  search: '',
  success: '',
  method: '',
  dateFrom: ''
})

const stats = reactive({
  successfulLogins: 0,
  failedLogins: 0,
  successRate: 0,
  suspiciousActivity: 0
})

// Computed
const totalPages = computed(() => Math.ceil(total.value / pageSize.value))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

// Methods
const getMethodLabel = (method) => {
  const labels = {
    'otp': 'OTP',
    'password': 'پسورد'
  }
  return labels[method] || method
}

const getBrowserInfo = (userAgent) => {
  if (!userAgent) return 'نامشخص'
  
  // Simple browser detection
  if (userAgent.includes('Chrome')) return 'Chrome'
  if (userAgent.includes('Firefox')) return 'Firefox'
  if (userAgent.includes('Safari')) return 'Safari'
  if (userAgent.includes('Edge')) return 'Edge'
  if (userAgent.includes('Opera')) return 'Opera'
  
  return 'مرورگر دیگر'
}

const formatDateTime = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString('fa-IR')
}

const loadHistory = async () => {
  try {
    const params = new URLSearchParams({
      page: currentPage.value,
      limit: pageSize.value,
      ...filters
    })

    const { data } = await $fetch(`/api/admin/login-history?${params}`)
    history.value = data.attempts
    total.value = data.total
    
    // Update stats
    updateStats()
  } catch (error) {
    console.error('خطا در بارگذاری تاریخچه:', error)
  }
}

const updateStats = async () => {
  try {
    const { data } = await $fetch('/api/admin/login-stats')
    Object.assign(stats, data)
  } catch (error) {
    console.error('خطا در بارگذاری آمار:', error)
  }
}

const viewAttemptDetails = (attempt) => {
  selectedAttempt.value = attempt
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
    loadHistory()
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    loadHistory()
  }
}

const goToPage = (page) => {
  currentPage.value = page
  loadHistory()
}

// Lifecycle
onMounted(() => {
  loadHistory()
})
</script> 
