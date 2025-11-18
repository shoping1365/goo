<template>
  <div class="space-y-6">
    <!-- فرم امتیازدهی دستی -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">امتیازدهی دستی</h3>
      </div>
      <div class="px-4 py-4">
        <form @submit.prevent="submitManualScore" class="space-y-6">
          <!-- انتخاب کاربر -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">انتخاب کاربر</label>
            <div class="flex space-x-3 space-x-reverse">
              <input 
                v-model="searchUser" 
                type="text" 
                placeholder="جستجو بر اساس نام یا ایمیل..." 
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                @input="searchUsers"
              >
              <button type="button" @click="showUserSelector = true" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
                انتخاب کاربر
              </button>
            </div>
            
            <!-- نتایج جستجو -->
            <div v-if="searchResults.length > 0 && searchUser" class="mt-2 border border-gray-200 rounded-md max-h-48 overflow-y-auto">
              <div 
                v-for="user in searchResults" 
                :key="user.id" 
                @click="selectUser(user)"
                class="p-3 hover:bg-gray-50 cursor-pointer border-b border-gray-100 last:border-b-0"
              >
                <div class="flex items-center">
                  <img :src="user.avatar || '/default-avatar.png'" :alt="user.name" class="w-8 h-8 rounded-full mr-3">
                  <div>
                    <div class="font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">{{ user.email }}</div>
                    <div class="text-xs text-gray-400">امتیاز فعلی: {{ user.currentScore }}</div>
                  </div>
                </div>
              </div>
            </div>

            <!-- کاربر انتخاب شده -->
            <div v-if="selectedUser" class="mt-4 px-4 py-4 bg-blue-50 rounded-lg">
              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <img :src="selectedUser.avatar || '/default-avatar.png'" :alt="selectedUser.name" class="w-12 h-12 rounded-full mr-4">
                  <div>
                    <div class="font-medium text-gray-900">{{ selectedUser.name }}</div>
                    <div class="text-sm text-gray-500">{{ selectedUser.email }}</div>
                    <div class="text-sm text-gray-600">امتیاز فعلی: {{ selectedUser.currentScore }}</div>
                  </div>
                </div>
                <button type="button" @click="selectedUser = null" class="text-red-600 hover:text-red-800">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>

          <!-- نوع امتیاز -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع امتیاز</label>
            <select v-model="scoreType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="positive">امتیاز مثبت</option>
              <option value="negative">امتیاز منفی</option>
              <option value="bonus">امتیاز ویژه</option>
              <option value="penalty">جریمه</option>
            </select>
          </div>

          <!-- مقدار امتیاز -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مقدار امتیاز</label>
            <input 
              v-model="scoreValue" 
              type="number" 
              :min="scoreType === 'negative' || scoreType === 'penalty' ? -1000 : 0"
              :max="scoreType === 'positive' || scoreType === 'bonus' ? 1000 : 0"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="مقدار امتیاز را وارد کنید"
            >
          </div>

          <!-- دلیل امتیازدهی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دلیل امتیازدهی</label>
            <select v-model="scoreReason" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="">انتخاب دلیل</option>
              <option v-if="scoreType === 'positive' || scoreType === 'bonus'" value="excellent_service">خدمات عالی</option>
              <option v-if="scoreType === 'positive' || scoreType === 'bonus'" value="helpful_review">نظر مفید</option>
              <option v-if="scoreType === 'positive' || scoreType === 'bonus'" value="referral">ارجاع موفق</option>
              <option v-if="scoreType === 'positive' || scoreType === 'bonus'" value="loyalty">وفاداری</option>
              <option v-if="scoreType === 'positive' || scoreType === 'bonus'" value="special_contribution">مشارکت ویژه</option>
              <option v-if="scoreType === 'negative' || scoreType === 'penalty'" value="inappropriate_behavior">رفتار نامناسب</option>
              <option v-if="scoreType === 'negative' || scoreType === 'penalty'" value="spam_reviews">نظرات اسپم</option>
              <option v-if="scoreType === 'negative' || scoreType === 'penalty'" value="rule_violation">تخلف از قوانین</option>
              <option v-if="scoreType === 'negative' || scoreType === 'penalty'" value="fraud">کلاهبرداری</option>
              <option value="other">سایر</option>
            </select>
          </div>

          <!-- توضیحات -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات (اختیاری)</label>
            <textarea 
              v-model="scoreDescription" 
              rows="3" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="توضیحات اضافی درباره دلیل امتیازدهی..."
            ></textarea>
          </div>

          <!-- پیش‌نمایش تغییرات -->
          <div v-if="selectedUser" class="px-4 py-4 bg-gray-50 rounded-lg">
            <h4 class="font-medium text-gray-900 mb-2">پیش‌نمایش تغییرات</h4>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between">
                <span>امتیاز فعلی:</span>
                <span class="font-medium">{{ selectedUser.currentScore }}</span>
              </div>
              <div class="flex justify-between">
                <span>امتیاز جدید:</span>
                <span :class="getScoreChangeClass()" class="font-medium">
                  {{ selectedUser.currentScore + scoreValueNumber }}
                </span>
              </div>
              <div class="flex justify-between">
                <span>تغییر:</span>
                <span :class="getScoreChangeClass()" class="font-medium">
                  {{ scoreValueNumber > 0 ? '+' : '' }}{{ scoreValue || 0 }}
                </span>
              </div>
            </div>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="flex justify-end space-x-3 space-x-reverse">
            <button type="button" @click="resetForm" class="px-6 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
              انصراف
            </button>
            <button 
              type="submit" 
              :disabled="!selectedUser || !scoreValue || !scoreReason"
              class="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              اعمال امتیاز
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- تاریخچه امتیازدهی دستی -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">تاریخچه امتیازدهی دستی</h3>
          <div class="flex space-x-3 space-x-reverse">
            <select v-model="filterPeriod" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="all">همه زمان‌ها</option>
              <option value="today">امروز</option>
              <option value="week">هفته جاری</option>
              <option value="month">ماه جاری</option>
            </select>
            <button @click="exportHistory" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors">
              خروجی اکسل
            </button>
          </div>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مقدار</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدیر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="record in filteredHistory" :key="record.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="record.user.avatar || '/default-avatar.png'" :alt="record.user.name" class="w-8 h-8 rounded-full mr-3">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ record.user.name }}</div>
                    <div class="text-sm text-gray-500">{{ record.user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getScoreTypeClass(record.type)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getScoreTypeText(record.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="record.value > 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                  {{ record.value > 0 ? '+' : '' }}{{ record.value }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ getReasonText(record.reason) }}</div>
                <div v-if="record.description" class="text-sm text-gray-500">{{ record.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ record.admin.name }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(record.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button @click="viewDetails(record)" class="text-blue-600 hover:text-blue-900">جزئیات</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <button @click="previousPage" :disabled="currentPage === 1" class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50">
            قبلی
          </button>
          <button @click="nextPage" :disabled="currentPage === totalPages" class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50">
            بعدی
          </button>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              نمایش <span class="font-medium">{{ paginationStart }}</span> تا <span class="font-medium">{{ paginationEnd }}</span> از <span class="font-medium">{{ totalRecords }}</span> نتیجه
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal جزئیات -->
    <div v-if="showDetailsModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">جزئیات امتیازدهی</h3>
          <div v-if="selectedRecord" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">کاربر:</label>
              <p class="text-sm text-gray-900">{{ selectedRecord.user.name }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">نوع امتیاز:</label>
              <p class="text-sm text-gray-900">{{ getScoreTypeText(selectedRecord.type) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">مقدار:</label>
              <p class="text-sm text-gray-900">{{ selectedRecord.value }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">دلیل:</label>
              <p class="text-sm text-gray-900">{{ getReasonText(selectedRecord.reason) }}</p>
            </div>
            <div v-if="selectedRecord.description">
              <label class="block text-sm font-medium text-gray-700">توضیحات:</label>
              <p class="text-sm text-gray-900">{{ selectedRecord.description }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">مدیر:</label>
              <p class="text-sm text-gray-900">{{ selectedRecord.admin.name }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ:</label>
              <p class="text-sm text-gray-900">{{ formatDate(selectedRecord.createdAt) }}</p>
            </div>
          </div>
          <div class="flex justify-end mt-6">
            <button @click="showDetailsModal = false" class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400">بستن</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';

// Props
const props = defineProps<{
  users: any[]
}>()

// Emits
const emit = defineEmits<{
  submitScore: [data: any]
  exportHistory: [data: any]
}>()

// Reactive data
const searchUser = ref('')
const selectedUser = ref(null)
const showUserSelector = ref(false)
const searchResults = ref([])
const scoreType = ref('positive')
const scoreValue = ref('')
const scoreReason = ref('')
const scoreDescription = ref('')
const filterPeriod = ref('all')
const currentPage = ref(1)
const showDetailsModal = ref(false)
const selectedRecord = ref(null)

// Sample history data
const manualScoreHistory = ref([
  {
    id: 1,
    user: { name: 'علی احمدی', email: 'ali@example.com', avatar: '/avatars/ali.jpg' },
    type: 'positive',
    value: 50,
    reason: 'excellent_service',
    description: 'خدمات عالی به مشتریان',
    admin: { name: 'مدیر سیستم' },
    createdAt: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    user: { name: 'فاطمه محمدی', email: 'fateme@example.com', avatar: '/avatars/fateme.jpg' },
    type: 'negative',
    value: -20,
    reason: 'inappropriate_behavior',
    description: 'رفتار نامناسب در نظرات',
    admin: { name: 'مدیر سیستم' },
    createdAt: '2024-01-14T15:45:00Z'
  }
])

// Computed properties
const scoreValueNumber = computed(() => parseInt(scoreValue.value || '0'))

const filteredHistory = computed(() => {
  let filtered = manualScoreHistory.value

  if (filterPeriod.value !== 'all') {
    const now = new Date()
    const filterDate = new Date()
    
    switch (filterPeriod.value) {
      case 'today':
        filterDate.setHours(0, 0, 0, 0)
        break
      case 'week':
        filterDate.setDate(now.getDate() - 7)
        break
      case 'month':
        filterDate.setMonth(now.getMonth() - 1)
        break
    }
    
    filtered = filtered.filter(record => new Date(record.createdAt) >= filterDate)
  }

  return filtered
})

const totalRecords = computed(() => filteredHistory.value.length)
const totalPages = computed(() => Math.ceil(totalRecords.value / 10))
const paginationStart = computed(() => (currentPage.value - 1) * 10 + 1)
const paginationEnd = computed(() => Math.min(currentPage.value * 10, totalRecords.value))

// Methods
const searchUsers = () => {
  if (searchUser.value.length < 2) {
    searchResults.value = []
    return
  }

  searchResults.value = props.users.filter(user => 
    user.name.toLowerCase().includes(searchUser.value.toLowerCase()) ||
    user.email.toLowerCase().includes(searchUser.value.toLowerCase())
  ).slice(0, 5)
}

const selectUser = (user: any) => {
  selectedUser.value = user
  searchUser.value = user.name
  searchResults.value = []
}

const getScoreChangeClass = () => {
  const value = scoreValueNumber.value
  return value > 0 ? 'text-green-600' : value < 0 ? 'text-red-600' : 'text-gray-600'
}

const getScoreTypeClass = (type: string) => {
  switch (type) {
    case 'positive':
    case 'bonus':
      return 'bg-green-100 text-green-800'
    case 'negative':
    case 'penalty':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getScoreTypeText = (type: string) => {
  switch (type) {
    case 'positive':
      return 'مثبت'
    case 'negative':
      return 'منفی'
    case 'bonus':
      return 'ویژه'
    case 'penalty':
      return 'جریمه'
    default:
      return 'نامشخص'
  }
}

const getReasonText = (reason: string) => {
  const reasons = {
    excellent_service: 'خدمات عالی',
    helpful_review: 'نظر مفید',
    referral: 'ارجاع موفق',
    loyalty: 'وفاداری',
    special_contribution: 'مشارکت ویژه',
    inappropriate_behavior: 'رفتار نامناسب',
    spam_reviews: 'نظرات اسپم',
    rule_violation: 'تخلف از قوانین',
    fraud: 'کلاهبرداری',
    other: 'سایر'
  }
  return reasons[reason] || reason
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const submitManualScore = () => {
  if (!selectedUser.value || !scoreValue.value || !scoreReason.value) {
    return
  }

  const scoreData = {
    userId: selectedUser.value.id,
    type: scoreType.value,
    value: parseInt(scoreValue.value),
    reason: scoreReason.value,
    description: scoreDescription.value,
    adminId: 1 // در حالت واقعی از context مدیر گرفته می‌شود
  }

  emit('submitScore', scoreData)
  resetForm()
}

const resetForm = () => {
  selectedUser.value = null
  searchUser.value = ''
  scoreType.value = 'positive'
  scoreValue.value = ''
  scoreReason.value = ''
  scoreDescription.value = ''
  searchResults.value = []
}

const exportHistory = () => {
  emit('exportHistory', {
    history: filteredHistory.value,
    period: filterPeriod.value
  })
}

const viewDetails = (record: any) => {
  selectedRecord.value = record
  showDetailsModal.value = true
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

// Watchers
watch(searchUser, () => {
  searchUsers()
})
</script> 
