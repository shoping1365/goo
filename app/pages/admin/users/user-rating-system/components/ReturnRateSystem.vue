<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">سیستم نرخ بازگشت کالا</h2>
          <p class="mt-1 text-sm text-gray-500">پیگیری مرجوعات و امتیازدهی بر اساس الگوی بازگشت کاربران</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button @click="analyzeReturnPatterns" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
            تحلیل الگوها
          </button>
          <button @click="exportReturnReport" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors">
            خروجی گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- Return Rate Settings -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات سیستم نرخ بازگشت</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه نرخ بازگشت (%)</label>
          <input v-model="returnSettings.returnRateThreshold" type="number" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز منفی برای بازگشت</label>
          <input v-model="returnSettings.returnPenaltyScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز مثبت برای عدم بازگشت</label>
          <input v-model="returnSettings.noReturnBonusScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل تعداد خرید برای محاسبه</label>
          <input v-model="returnSettings.minPurchaseCount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دوره محاسبه (روز)</label>
          <input v-model="returnSettings.calculationPeriod" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای بازگشت معتبر</label>
          <input v-model="returnSettings.validReturnScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      
      <div class="mt-6 space-y-4">
        <label class="flex items-center">
          <input v-model="returnSettings.trackReturnPatterns" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">پیگیری الگوهای بازگشت</span>
        </label>
        <label class="flex items-center">
          <input v-model="returnSettings.autoCalculateScores" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">محاسبه خودکار امتیازات</span>
        </label>
        <label class="flex items-center">
          <input v-model="returnSettings.notifyHighReturnRate" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">اعلان برای نرخ بازگشت بالا</span>
        </label>
      </div>
      
      <div class="mt-4 flex space-x-3 space-x-reverse">
        <button @click="saveReturnSettings" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
          ذخیره تنظیمات
        </button>
        <button @click="testReturnSystem" class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700 transition-colors">
          تست سیستم
        </button>
      </div>
    </div>

    <!-- Return Rate Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران مطمئن</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.reliableUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران مشکوک</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.suspiciousUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">میانگین نرخ بازگشت</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.avgReturnRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">بازگشت‌های معتبر</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.validReturns }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Return Patterns Analysis -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تحلیل الگوهای بازگشت</h3>
      
      <!-- Search and Filter -->
      <div class="flex flex-col sm:flex-row gapx-4 py-4 mb-6">
        <div class="flex-1">
          <input v-model="searchQuery" type="text" placeholder="جستجو در کاربران..." class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <select v-model="filterReturnRate" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه نرخ‌ها</option>
          <option value="low">کم (0-5%)</option>
          <option value="medium">متوسط (5-15%)</option>
          <option value="high">زیاد (15%+)</option>
        </select>
        <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه وضعیت‌ها</option>
          <option value="reliable">مطمئن</option>
          <option value="suspicious">مشکوک</option>
          <option value="at-risk">در معرض خطر</option>
        </select>
      </div>

      <!-- Return Patterns Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد خرید</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد بازگشت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ بازگشت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین بازگشت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="pattern in paginatedPatterns" :key="pattern.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="pattern.userAvatar" :alt="pattern.userName" class="w-8 h-8 rounded-full">
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ pattern.userName }}</div>
                    <div class="text-sm text-gray-500">{{ pattern.userEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ pattern.totalPurchases }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ pattern.totalReturns }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getReturnRateClass(pattern.returnRate)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ pattern.returnRate }}%
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(pattern.lastReturn) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(pattern.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getStatusText(pattern.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ pattern.score }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button @click="viewReturnHistory(pattern)" class="text-blue-600 hover:text-blue-900 ml-3">تاریخچه</button>
                <button @click="analyzeUserReturns(pattern)" class="text-green-600 hover:text-green-900 ml-3">تحلیل</button>
                <button @click="flagUser(pattern)" class="text-red-600 hover:text-red-900">علامت‌گذاری</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="flex items-center justify-between mt-6">
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-700">نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} الگو</span>
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button @click="previousPage" :disabled="currentPage === 1" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50">
            قبلی
          </button>
          <button v-for="page in visiblePages" :key="page" @click="goToPage(page)" :class="page === currentPage ? 'bg-blue-600 text-white' : 'bg-white text-gray-700'" class="px-3 py-1 border border-gray-300 rounded-md text-sm">
            {{ page }}
          </button>
          <button @click="nextPage" :disabled="currentPage === totalPages" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50">
            بعدی
          </button>
        </div>
      </div>
    </div>

    <!-- Return Rate Chart -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">نمودار نرخ بازگشت</h3>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="currentColor" viewBox="0 0 20 20">
            <path d="M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z"></path>
          </svg>
          <p class="text-gray-500">نمودار نرخ بازگشت در اینجا نمایش داده می‌شود</p>
        </div>
      </div>
    </div>

    <!-- High Return Rate Users Alert -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">کاربران با نرخ بازگشت بالا</h3>
      <div class="space-y-4">
        <div v-for="user in highReturnUsers" :key="user.id" class="flex items-center justify-between px-4 py-4 border border-red-200 rounded-lg bg-red-50">
          <div class="flex items-center">
            <img :src="user.avatar" :alt="user.name" class="w-10 h-10 rounded-full">
            <div class="mr-4">
              <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
              <div class="text-sm text-gray-500">{{ user.email }}</div>
              <div class="text-sm text-red-600">نرخ بازگشت: {{ user.returnRate }}%</div>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button @click="investigateUser(user)" class="bg-red-600 text-white px-3 py-1 rounded-md text-sm hover:bg-red-700">
              بررسی
            </button>
            <button @click="viewUserDetails(user)" class="bg-gray-600 text-white px-3 py-1 rounded-md text-sm hover:bg-gray-700">
              مشاهده جزئیات
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

// Props and Emits
defineProps<{
  users?: any[]
}>()

defineEmits<{
  saveSettings: [settings: any]
  analyzePatterns: [users: any[]]
  exportReport: [data: any]
}>()

// Reactive data
const searchQuery = ref('')
const filterReturnRate = ref('')
const filterStatus = ref('')
const currentPage = ref(1)
const itemsPerPage = 10

// Return Settings
const returnSettings = ref({
  returnRateThreshold: 15,
  returnPenaltyScore: 20,
  noReturnBonusScore: 10,
  minPurchaseCount: 5,
  calculationPeriod: 365,
  validReturnScore: 5,
  trackReturnPatterns: true,
  autoCalculateScores: true,
  notifyHighReturnRate: true
})

// Sample return patterns data
const returnPatterns = ref([
  {
    id: 1,
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/avatars/ali.jpg',
    totalPurchases: 25,
    totalReturns: 2,
    returnRate: 8,
    lastReturn: '2024-01-10T10:30:00Z',
    status: 'reliable',
    score: 85
  },
  {
    id: 2,
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/avatars/fateme.jpg',
    totalPurchases: 15,
    totalReturns: 5,
    returnRate: 33.3,
    lastReturn: '2024-01-15T15:45:00Z',
    status: 'suspicious',
    score: 45
  },
  {
    id: 3,
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    userAvatar: '/avatars/mohammad.jpg',
    totalPurchases: 8,
    totalReturns: 6,
    returnRate: 75,
    lastReturn: '2024-01-12T09:20:00Z',
    status: 'at-risk',
    score: 15
  }
])

// Sample high return users
const highReturnUsers = ref([
  {
    id: 1,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    avatar: '/avatars/mohammad.jpg',
    returnRate: 75
  },
  {
    id: 2,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    avatar: '/avatars/zahra.jpg',
    returnRate: 60
  }
])

// Statistics
const stats = ref({
  reliableUsers: 750,
  suspiciousUsers: 120,
  avgReturnRate: 12.5,
  validReturns: 85
})

// Pagination
const pagination = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage + 1
  const end = Math.min(currentPage.value * itemsPerPage, filteredPatterns.value.length)
  return {
    start,
    end,
    total: filteredPatterns.value.length
  }
})

const totalPages = computed(() => Math.ceil(filteredPatterns.value.length / itemsPerPage))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

// Filtered patterns
const filteredPatterns = computed(() => {
  let filtered = returnPatterns.value

  if (searchQuery.value) {
    filtered = filtered.filter(pattern => 
      pattern.userName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      pattern.userEmail.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterReturnRate.value) {
    filtered = filtered.filter(pattern => {
      switch (filterReturnRate.value) {
        case 'low':
          return pattern.returnRate <= 5
        case 'medium':
          return pattern.returnRate > 5 && pattern.returnRate <= 15
        case 'high':
          return pattern.returnRate > 15
        default:
          return true
      }
    })
  }

  if (filterStatus.value) {
    filtered = filtered.filter(pattern => pattern.status === filterStatus.value)
  }

  return filtered
})

const paginatedPatterns = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredPatterns.value.slice(start, end)
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getReturnRateClass = (rate: number) => {
  if (rate <= 5) {
    return 'bg-green-100 text-green-800'
  } else if (rate <= 15) {
    return 'bg-yellow-100 text-yellow-800'
  } else {
    return 'bg-red-100 text-red-800'
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'reliable':
      return 'bg-green-100 text-green-800'
    case 'suspicious':
      return 'bg-yellow-100 text-yellow-800'
    case 'at-risk':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'reliable':
      return 'مطمئن'
    case 'suspicious':
      return 'مشکوک'
    case 'at-risk':
      return 'در معرض خطر'
    default:
      return 'نامشخص'
  }
}

const saveReturnSettings = () => {
  console.log('ذخیره تنظیمات نرخ بازگشت:', returnSettings.value)
  // API call to save return settings
}

const testReturnSystem = () => {
  console.log('تست سیستم نرخ بازگشت')
  // Test return system functionality
}

const analyzeReturnPatterns = () => {
  console.log('تحلیل الگوهای بازگشت')
  // Analyze return patterns for all users
}

const exportReturnReport = () => {
  console.log('خروجی گزارش نرخ بازگشت')
  // Export return report
}

const viewReturnHistory = (pattern: any) => {
  console.log('مشاهده تاریخچه بازگشت:', pattern)
  // Open return history modal
}

const analyzeUserReturns = (pattern: any) => {
  console.log('تحلیل بازگشت کاربر:', pattern)
  // Analyze specific user returns
}

const flagUser = (pattern: any) => {
  console.log('علامت‌گذاری کاربر:', pattern)
  // Flag user for investigation
}

const investigateUser = (user: any) => {
  console.log('بررسی کاربر:', user)
  // Investigate high return rate user
}

const viewUserDetails = (user: any) => {
  console.log('مشاهده جزئیات کاربر:', user)
  // View user details
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

const goToPage = (page: number) => {
  currentPage.value = page
}

// Lifecycle
onMounted(() => {
  console.log('سیستم نرخ بازگشت کالا بارگذاری شد')
})
</script> 
