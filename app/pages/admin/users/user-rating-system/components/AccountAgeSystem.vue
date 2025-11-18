<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">سیستم قدمت حساب کاربری</h2>
          <p class="mt-1 text-sm text-gray-500">محاسبه امتیاز بر اساس سن حساب و وفاداری طولانی‌مدت کاربران</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button @click="analyzeAccountAges" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
            تحلیل قدمت
          </button>
          <button @click="exportAgeReport" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors">
            خروجی گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- Account Age Settings -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات سیستم قدمت حساب</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز پایه برای حساب جدید</label>
          <input v-model="ageSettings.baseScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز سالانه</label>
          <input v-model="ageSettings.yearlyScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای 5 سال</label>
          <input v-model="ageSettings.fiveYearBonus" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای 10 سال</label>
          <input v-model="ageSettings.tenYearBonus" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر امتیاز قدمت</label>
          <input v-model="ageSettings.maxAgeScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">ضریب فعالیت</label>
          <input v-model="ageSettings.activityMultiplier" type="number" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      
      <div class="mt-6 space-y-4">
        <label class="flex items-center">
          <input v-model="ageSettings.autoCalculateScores" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">محاسبه خودکار امتیازات</span>
        </label>
        <label class="flex items-center">
          <input v-model="ageSettings.celebrateMilestones" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">جشن مراحل مهم</span>
        </label>
        <label class="flex items-center">
          <input v-model="ageSettings.notifyAnniversaries" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">اعلان سالگرد</span>
        </label>
      </div>
      
      <div class="mt-4 flex space-x-3 space-x-reverse">
        <button @click="saveAgeSettings" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
          ذخیره تنظیمات
        </button>
        <button @click="testAgeSystem" class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700 transition-colors">
          تست سیستم
        </button>
      </div>
    </div>

    <!-- Account Age Statistics -->
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
            <p class="text-sm font-medium text-gray-500">کاربران قدیمی</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.oldUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">میانگین قدمت</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.avgAccountAge }} سال</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران وفادار</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.loyalUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران جدید</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.newUsers }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Account Age Analysis -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تحلیل قدمت حساب‌ها</h3>
      
      <!-- Search and Filter -->
      <div class="flex flex-col sm:flex-row gapx-4 py-4 mb-6">
        <div class="flex-1">
          <input v-model="searchQuery" type="text" placeholder="جستجو در کاربران..." class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <select v-model="filterAge" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه سنین</option>
          <option value="new">جدید (0-1 سال)</option>
          <option value="young">جوان (1-3 سال)</option>
          <option value="mature">بالغ (3-7 سال)</option>
          <option value="old">قدیمی (7+ سال)</option>
        </select>
        <select v-model="filterLoyalty" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه سطوح</option>
          <option value="bronze">برنزی</option>
          <option value="silver">نقره‌ای</option>
          <option value="gold">طلایی</option>
          <option value="platinum">پلاتینیوم</option>
          <option value="diamond">الماس</option>
        </select>
      </div>

      <!-- Account Age Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ عضویت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">قدمت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح وفاداری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین فعالیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="account in paginatedAccounts" :key="account.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="account.userAvatar" :alt="account.userName" class="w-8 h-8 rounded-full">
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ account.userName }}</div>
                    <div class="text-sm text-gray-500">{{ account.userEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(account.joinDate) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getAgeClass(account.age)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ account.age }} سال
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getLoyaltyClass(account.loyaltyLevel)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getLoyaltyText(account.loyaltyLevel) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(account.lastActivity) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(account.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getStatusText(account.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ account.score }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button @click="viewAccountHistory(account)" class="text-blue-600 hover:text-blue-900 ml-3">تاریخچه</button>
                <button @click="celebrateMilestone(account)" class="text-green-600 hover:text-green-900 ml-3">جشن</button>
                <button @click="sendAnniversary(account)" class="text-purple-600 hover:text-purple-900">سالگرد</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="flex items-center justify-between mt-6">
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-700">نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} حساب</span>
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

    <!-- Account Age Chart -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">نمودار توزیع قدمت حساب‌ها</h3>
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="currentColor" viewBox="0 0 20 20">
            <path d="M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z"></path>
          </svg>
          <p class="text-gray-500">نمودار توزیع قدمت حساب‌ها در اینجا نمایش داده می‌شود</p>
        </div>
      </div>
    </div>

    <!-- Milestone Celebrations -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">جشن مراحل مهم</h3>
      <div class="space-y-4">
        <div v-for="milestone in milestones" :key="milestone.id" class="flex items-center justify-between px-4 py-4 border border-yellow-200 rounded-lg bg-yellow-50">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-6 h-6 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
              </svg>
            </div>
            <div class="mr-4">
              <div class="text-sm font-medium text-gray-900">{{ milestone.userName }}</div>
              <div class="text-sm text-yellow-600">{{ milestone.milestone }}</div>
              <div class="text-sm text-gray-500">{{ formatDate(milestone.date) }}</div>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button @click="sendCongratulations(milestone)" class="bg-yellow-600 text-white px-3 py-1 rounded-md text-sm hover:bg-yellow-700">
              تبریک
            </button>
            <button @click="viewUserDetails(milestone)" class="bg-gray-600 text-white px-3 py-1 rounded-md text-sm hover:bg-gray-700">
              مشاهده
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
  analyzeAges: [users: any[]]
  exportReport: [data: any]
}>()

// Reactive data
const searchQuery = ref('')
const filterAge = ref('')
const filterLoyalty = ref('')
const currentPage = ref(1)
const itemsPerPage = 10

// Age Settings
const ageSettings = ref({
  baseScore: 10,
  yearlyScore: 5,
  fiveYearBonus: 50,
  tenYearBonus: 100,
  maxAgeScore: 500,
  activityMultiplier: 1.2,
  autoCalculateScores: true,
  celebrateMilestones: true,
  notifyAnniversaries: true
})

// Sample account age data
const accountAges = ref([
  {
    id: 1,
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/avatars/ali.jpg',
    joinDate: '2018-03-15T10:30:00Z',
    age: 5.8,
    loyaltyLevel: 'gold',
    lastActivity: '2024-01-15T10:30:00Z',
    status: 'active',
    score: 125
  },
  {
    id: 2,
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/avatars/fateme.jpg',
    joinDate: '2020-07-20T15:45:00Z',
    age: 3.5,
    loyaltyLevel: 'silver',
    lastActivity: '2024-01-14T15:45:00Z',
    status: 'active',
    score: 75
  },
  {
    id: 3,
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    userAvatar: '/avatars/mohammad.jpg',
    joinDate: '2023-01-10T09:20:00Z',
    age: 1.0,
    loyaltyLevel: 'bronze',
    lastActivity: '2024-01-13T09:20:00Z',
    status: 'active',
    score: 25
  }
])

// Sample milestones
const milestones = ref([
  {
    id: 1,
    userName: 'علی احمدی',
    milestone: '5 سال عضویت',
    date: '2024-03-15T10:30:00Z'
  },
  {
    id: 2,
    userName: 'فاطمه محمدی',
    milestone: '3 سال عضویت',
    date: '2024-07-20T15:45:00Z'
  }
])

// Statistics
const stats = ref({
  oldUsers: 250,
  avgAccountAge: 3.2,
  loyalUsers: 180,
  newUsers: 420
})

// Pagination
const pagination = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage + 1
  const end = Math.min(currentPage.value * itemsPerPage, filteredAccounts.value.length)
  return {
    start,
    end,
    total: filteredAccounts.value.length
  }
})

const totalPages = computed(() => Math.ceil(filteredAccounts.value.length / itemsPerPage))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

// Filtered accounts
const filteredAccounts = computed(() => {
  let filtered = accountAges.value

  if (searchQuery.value) {
    filtered = filtered.filter(account => 
      account.userName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      account.userEmail.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterAge.value) {
    filtered = filtered.filter(account => {
      switch (filterAge.value) {
        case 'new':
          return account.age <= 1
        case 'young':
          return account.age > 1 && account.age <= 3
        case 'mature':
          return account.age > 3 && account.age <= 7
        case 'old':
          return account.age > 7
        default:
          return true
      }
    })
  }

  if (filterLoyalty.value) {
    filtered = filtered.filter(account => account.loyaltyLevel === filterLoyalty.value)
  }

  return filtered
})

const paginatedAccounts = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredAccounts.value.slice(start, end)
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getAgeClass = (age: number) => {
  if (age <= 1) {
    return 'bg-blue-100 text-blue-800'
  } else if (age <= 3) {
    return 'bg-green-100 text-green-800'
  } else if (age <= 7) {
    return 'bg-yellow-100 text-yellow-800'
  } else {
    return 'bg-purple-100 text-purple-800'
  }
}

const getLoyaltyClass = (level: string) => {
  switch (level) {
    case 'bronze':
      return 'bg-orange-100 text-orange-800'
    case 'silver':
      return 'bg-gray-100 text-gray-800'
    case 'gold':
      return 'bg-yellow-100 text-yellow-800'
    case 'platinum':
      return 'bg-blue-100 text-blue-800'
    case 'diamond':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getLoyaltyText = (level: string) => {
  switch (level) {
    case 'bronze':
      return 'برنزی'
    case 'silver':
      return 'نقره‌ای'
    case 'gold':
      return 'طلایی'
    case 'platinum':
      return 'پلاتینیوم'
    case 'diamond':
      return 'الماس'
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
    case 'suspended':
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
    case 'suspended':
      return 'معلق'
    default:
      return 'نامشخص'
  }
}

const saveAgeSettings = () => {
  console.log('ذخیره تنظیمات قدمت حساب:', ageSettings.value)
  // API call to save age settings
}

const testAgeSystem = () => {
  console.log('تست سیستم قدمت حساب')
  // Test age system functionality
}

const analyzeAccountAges = () => {
  console.log('تحلیل قدمت حساب‌ها')
  // Analyze account ages for all users
}

const exportAgeReport = () => {
  console.log('خروجی گزارش قدمت حساب')
  // Export age report
}

const viewAccountHistory = (account: any) => {
  console.log('مشاهده تاریخچه حساب:', account)
  // Open account history modal
}

const celebrateMilestone = (account: any) => {
  console.log('جشن مرحله مهم:', account)
  // Celebrate user milestone
}

const sendAnniversary = (account: any) => {
  console.log('ارسال سالگرد:', account)
  // Send anniversary notification
}

const sendCongratulations = (milestone: any) => {
  console.log('ارسال تبریک:', milestone)
  // Send congratulations message
}

const viewUserDetails = (milestone: any) => {
  console.log('مشاهده جزئیات کاربر:', milestone)
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
  console.log('سیستم قدمت حساب کاربری بارگذاری شد')
})
</script> 
