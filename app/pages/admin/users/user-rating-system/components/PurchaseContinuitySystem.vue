<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">سیستم تداوم خرید</h2>
          <p class="mt-1 text-sm text-gray-500">تحلیل الگوی خرید کاربران و امتیازدهی بر اساس تداوم خرید</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button @click="analyzePurchasePatterns" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
            تحلیل الگوها
          </button>
          <button @click="exportContinuityReport" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors">
            خروجی گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- Continuity Settings -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات سیستم تداوم خرید</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای خرید هفتگی</label>
          <input v-model="continuitySettings.weeklyPurchaseScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای خرید ماهانه</label>
          <input v-model="continuitySettings.monthlyPurchaseScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای خرید فصلی</label>
          <input v-model="continuitySettings.quarterlyPurchaseScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل فاصله خرید (روز)</label>
          <input v-model="continuitySettings.minPurchaseInterval" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر فاصله خرید (روز)</label>
          <input v-model="continuitySettings.maxPurchaseInterval" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای تداوم 3 ماهه</label>
          <input v-model="continuitySettings.threeMonthStreakScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      
      <div class="mt-6 space-y-4">
        <label class="flex items-center">
          <input v-model="continuitySettings.trackPurchasePatterns" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">پیگیری الگوهای خرید</span>
        </label>
        <label class="flex items-center">
          <input v-model="continuitySettings.autoCalculateScores" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">محاسبه خودکار امتیازات</span>
        </label>
        <label class="flex items-center">
          <input v-model="continuitySettings.notifyInactiveUsers" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">اعلان به کاربران غیرفعال</span>
        </label>
      </div>
      
      <div class="mt-4 flex space-x-3 space-x-reverse">
        <button @click="saveContinuitySettings" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
          ذخیره تنظیمات
        </button>
        <button @click="testContinuitySystem" class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700 transition-colors">
          تست سیستم
        </button>
      </div>
    </div>

    <!-- Purchase Continuity Statistics -->
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
            <p class="text-sm font-medium text-gray-500">کاربران فعال</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.activeUsers }}</p>
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
            <p class="text-sm font-medium text-gray-500">میانگین فاصله خرید</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.avgPurchaseInterval }} روز</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">تداوم 3 ماهه</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.threeMonthStreak }}</p>
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
            <p class="text-sm font-medium text-gray-500">کاربران غیرفعال</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.inactiveUsers }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Purchase Patterns Analysis -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تحلیل الگوهای خرید</h3>
      
      <!-- Search and Filter -->
      <div class="flex flex-col sm:flex-row gapx-4 py-4 mb-6">
        <div class="flex-1">
          <input v-model="searchQuery" type="text" placeholder="جستجو در کاربران..." class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <select v-model="filterPattern" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه الگوها</option>
          <option value="weekly">هفتگی</option>
          <option value="monthly">ماهانه</option>
          <option value="quarterly">فصلی</option>
          <option value="irregular">نامنظم</option>
        </select>
        <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه وضعیت‌ها</option>
          <option value="active">فعال</option>
          <option value="inactive">غیرفعال</option>
          <option value="at-risk">در معرض خطر</option>
        </select>
      </div>

      <!-- Purchase Patterns Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین خرید</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">الگوی خرید</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فاصله خرید</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تداوم</th>
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
                {{ formatDate(pattern.lastPurchase) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getPatternClass(pattern.pattern)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getPatternText(pattern.pattern) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ pattern.purchaseInterval }} روز
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex space-x-1 space-x-reverse">
                    <div v-for="i in 5" :key="i" :class="i <= pattern.streakLevel ? 'bg-green-500' : 'bg-gray-300'" class="w-2 h-2 rounded-full"></div>
                  </div>
                  <span class="mr-2 text-sm text-gray-500">{{ pattern.streakLevel }}/5</span>
                </div>
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
                <button @click="viewPurchaseHistory(pattern)" class="text-blue-600 hover:text-blue-900 ml-3">تاریخچه</button>
                <button @click="sendReminder(pattern)" class="text-green-600 hover:text-green-900 ml-3">یادآوری</button>
                <button @click="analyzeUserPattern(pattern)" class="text-purple-600 hover:text-purple-900">تحلیل</button>
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

    <!-- Purchase Continuity Chart -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">نمودار تداوم خرید</h3>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="currentColor" viewBox="0 0 20 20">
            <path d="M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z"></path>
          </svg>
          <p class="text-gray-500">نمودار تداوم خرید در اینجا نمایش داده می‌شود</p>
        </div>
      </div>
    </div>

    <!-- Inactive Users Alert -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">کاربران غیرفعال</h3>
      <div class="space-y-4">
        <div v-for="user in inactiveUsers" :key="user.id" class="flex items-center justify-between px-4 py-4 border border-gray-200 rounded-lg">
          <div class="flex items-center">
            <img :src="user.avatar" :alt="user.name" class="w-10 h-10 rounded-full">
            <div class="mr-4">
              <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
              <div class="text-sm text-gray-500">{{ user.email }}</div>
              <div class="text-sm text-gray-500">آخرین خرید: {{ formatDate(user.lastPurchase) }}</div>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button @click="sendInactiveReminder(user)" class="bg-blue-600 text-white px-3 py-1 rounded-md text-sm hover:bg-blue-700">
              ارسال یادآوری
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
const filterPattern = ref('')
const filterStatus = ref('')
const currentPage = ref(1)
const itemsPerPage = 10

// Continuity Settings
const continuitySettings = ref({
  weeklyPurchaseScore: 10,
  monthlyPurchaseScore: 25,
  quarterlyPurchaseScore: 50,
  minPurchaseInterval: 1,
  maxPurchaseInterval: 90,
  threeMonthStreakScore: 100,
  trackPurchasePatterns: true,
  autoCalculateScores: true,
  notifyInactiveUsers: true
})

// Sample purchase patterns data
const purchasePatterns = ref([
  {
    id: 1,
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/avatars/ali.jpg',
    lastPurchase: '2024-01-15T10:30:00Z',
    pattern: 'weekly',
    purchaseInterval: 7,
    streakLevel: 4,
    status: 'active',
    score: 85
  },
  {
    id: 2,
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/avatars/fateme.jpg',
    lastPurchase: '2024-01-10T15:45:00Z',
    pattern: 'monthly',
    purchaseInterval: 30,
    streakLevel: 3,
    status: 'active',
    score: 65
  },
  {
    id: 3,
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    userAvatar: '/avatars/mohammad.jpg',
    lastPurchase: '2023-12-20T09:20:00Z',
    pattern: 'irregular',
    purchaseInterval: 45,
    streakLevel: 1,
    status: 'at-risk',
    score: 25
  }
])

// Sample inactive users
const inactiveUsers = ref([
  {
    id: 1,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    avatar: '/avatars/zahra.jpg',
    lastPurchase: '2023-11-15T14:10:00Z'
  },
  {
    id: 2,
    name: 'حسین نوری',
    email: 'hossein@example.com',
    avatar: '/avatars/hossein.jpg',
    lastPurchase: '2023-10-20T11:30:00Z'
  }
])

// Statistics
const stats = ref({
  activeUsers: 850,
  avgPurchaseInterval: 15,
  threeMonthStreak: 320,
  inactiveUsers: 150
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
  let filtered = purchasePatterns.value

  if (searchQuery.value) {
    filtered = filtered.filter(pattern => 
      pattern.userName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      pattern.userEmail.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterPattern.value) {
    filtered = filtered.filter(pattern => pattern.pattern === filterPattern.value)
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

const getPatternClass = (pattern: string) => {
  switch (pattern) {
    case 'weekly':
      return 'bg-green-100 text-green-800'
    case 'monthly':
      return 'bg-blue-100 text-blue-800'
    case 'quarterly':
      return 'bg-yellow-100 text-yellow-800'
    case 'irregular':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getPatternText = (pattern: string) => {
  switch (pattern) {
    case 'weekly':
      return 'هفتگی'
    case 'monthly':
      return 'ماهانه'
    case 'quarterly':
      return 'فصلی'
    case 'irregular':
      return 'نامنظم'
    default:
      return 'نامشخص'
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    case 'at-risk':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'at-risk':
      return 'در معرض خطر'
    default:
      return 'نامشخص'
  }
}

const saveContinuitySettings = () => {
  console.log('ذخیره تنظیمات تداوم خرید:', continuitySettings.value)
  // API call to save continuity settings
}

const testContinuitySystem = () => {
  console.log('تست سیستم تداوم خرید')
  // Test continuity system functionality
}

const analyzePurchasePatterns = () => {
  console.log('تحلیل الگوهای خرید')
  // Analyze purchase patterns for all users
}

const exportContinuityReport = () => {
  console.log('خروجی گزارش تداوم خرید')
  // Export continuity report
}

const viewPurchaseHistory = (pattern: any) => {
  console.log('مشاهده تاریخچه خرید:', pattern)
  // Open purchase history modal
}

const sendReminder = (pattern: any) => {
  console.log('ارسال یادآوری:', pattern)
  // Send reminder to user
}

const analyzeUserPattern = (pattern: any) => {
  console.log('تحلیل الگوی کاربر:', pattern)
  // Analyze specific user pattern
}

const sendInactiveReminder = (user: any) => {
  console.log('ارسال یادآوری به کاربر غیرفعال:', user)
  // Send reminder to inactive user
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
  console.log('سیستم تداوم خرید بارگذاری شد')
})
</script> 
