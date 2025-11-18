<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">سیستم کیفیت نظرات</h2>
          <p class="mt-1 text-sm text-gray-500">تحلیل و امتیازدهی خودکار بر اساس کیفیت نظرات کاربران</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button @click="analyzeAllReviews" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
            تحلیل همه نظرات
          </button>
          <button @click="exportQualityReport" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors">
            خروجی گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- AI Settings -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات هوش مصنوعی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل طول نظر (کاراکتر)</label>
          <input v-model="aiSettings.minReviewLength" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل کلمات</label>
          <input v-model="aiSettings.minWordCount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز پایه برای نظر متنی</label>
          <input v-model="aiSettings.baseTextScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز اضافی برای عکس</label>
          <input v-model="aiSettings.photoBonus" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز اضافی برای ویدیو</label>
          <input v-model="aiSettings.videoBonus" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای اولین نظر</label>
          <input v-model="aiSettings.firstReviewBonus" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      
      <div class="mt-6">
        <label class="flex items-center">
          <input v-model="aiSettings.autoAnalysis" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">تحلیل خودکار نظرات جدید</span>
        </label>
      </div>
      
      <div class="mt-4 flex space-x-3 space-x-reverse">
        <button @click="saveAISettings" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
          ذخیره تنظیمات
        </button>
        <button @click="testAI" class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700 transition-colors">
          تست هوش مصنوعی
        </button>
      </div>
    </div>

    <!-- Quality Analysis -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تحلیل کیفیت نظرات</h3>
      
      <!-- Search and Filter -->
      <div class="flex flex-col sm:flex-row gapx-4 py-4 mb-6">
        <div class="flex-1">
          <input v-model="searchQuery" type="text" placeholder="جستجو در نظرات..." class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <select v-model="filterQuality" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه کیفیت‌ها</option>
          <option value="excellent">عالی</option>
          <option value="good">خوب</option>
          <option value="average">متوسط</option>
          <option value="poor">ضعیف</option>
        </select>
        <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه وضعیت‌ها</option>
          <option value="analyzed">تحلیل شده</option>
          <option value="pending">در انتظار</option>
          <option value="flagged">مشکوک</option>
        </select>
      </div>

      <!-- Reviews Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصول</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نظر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کیفیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="review in paginatedReviews" :key="review.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="review.userAvatar" :alt="review.userName" class="w-8 h-8 rounded-full">
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ review.userName }}</div>
                    <div class="text-sm text-gray-500">{{ review.userEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ review.productName }}</div>
                <div class="text-sm text-gray-500">{{ review.productCategory }}</div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm text-gray-900 max-w-xs truncate">{{ review.content }}</div>
                <div class="flex items-center mt-1 space-x-2 space-x-reverse">
                  <span v-if="review.hasPhoto" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    عکس
                  </span>
                  <span v-if="review.hasVideo" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-purple-100 text-purple-800">
                    ویدیو
                  </span>
                  <span v-if="review.isFirstReview" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800">
                    اولین نظر
                  </span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getQualityClass(review.quality)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getQualityText(review.quality) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ review.score }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(review.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getStatusText(review.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button @click="viewReviewDetails(review)" class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button @click="reanalyzeReview(review)" class="text-green-600 hover:text-green-900 ml-3">تحلیل مجدد</button>
                <button @click="flagReview(review)" class="text-red-600 hover:text-red-900">علامت‌گذاری</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="flex items-center justify-between mt-6">
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-700">نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} نظر</span>
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

    <!-- Quality Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">نظرات عالی</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.excellentReviews }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">نظرات خوب</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.goodReviews }}</p>
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
            <p class="text-sm font-medium text-gray-500">نظرات متوسط</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.averageReviews }}</p>
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
            <p class="text-sm font-medium text-gray-500">نظرات ضعیف</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.poorReviews }}</p>
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
  reviews?: any[]
}>()

defineEmits<{
  saveSettings: [settings: any]
  analyzeReviews: [reviews: any[]]
  exportReport: [data: any]
}>()

// Reactive data
const searchQuery = ref('')
const filterQuality = ref('')
const filterStatus = ref('')
const currentPage = ref(1)
const itemsPerPage = 10

// AI Settings
const aiSettings = ref({
  minReviewLength: 50,
  minWordCount: 10,
  baseTextScore: 20,
  photoBonus: 30,
  videoBonus: 50,
  firstReviewBonus: 15,
  autoAnalysis: true
})

// Sample reviews data
const reviews = ref([
  {
    id: 1,
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/avatars/ali.jpg',
    productName: 'گوشی هوشمند سامسونگ',
    productCategory: 'الکترونیک',
    content: 'این گوشی واقعاً عالی است! کیفیت دوربین فوق‌العاده و باتری طولانی مدت دارد. حتماً پیشنهاد می‌کنم.',
    hasPhoto: true,
    hasVideo: false,
    isFirstReview: true,
    quality: 'excellent',
    score: 85,
    status: 'analyzed',
    createdAt: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/avatars/fateme.jpg',
    productName: 'لپ‌تاپ اپل',
    productCategory: 'کامپیوتر',
    content: 'خوب بود',
    hasPhoto: false,
    hasVideo: false,
    isFirstReview: false,
    quality: 'poor',
    score: 5,
    status: 'flagged',
    createdAt: '2024-01-14T15:45:00Z'
  },
  {
    id: 3,
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    userAvatar: '/avatars/mohammad.jpg',
    productName: 'کتاب برنامه‌نویسی',
    productCategory: 'کتاب',
    content: 'کتاب مفیدی است و مثال‌های خوبی دارد. برای مبتدیان مناسب است.',
    hasPhoto: false,
    hasVideo: false,
    isFirstReview: false,
    quality: 'good',
    score: 45,
    status: 'analyzed',
    createdAt: '2024-01-13T09:20:00Z'
  }
])

// Statistics
const stats = ref({
  excellentReviews: 450,
  goodReviews: 380,
  averageReviews: 250,
  poorReviews: 120
})

// Pagination
const pagination = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage + 1
  const end = Math.min(currentPage.value * itemsPerPage, filteredReviews.value.length)
  return {
    start,
    end,
    total: filteredReviews.value.length
  }
})

const totalPages = computed(() => Math.ceil(filteredReviews.value.length / itemsPerPage))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

// Filtered reviews
const filteredReviews = computed(() => {
  let filtered = reviews.value

  if (searchQuery.value) {
    filtered = filtered.filter(review => 
      review.content.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      review.userName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      review.productName.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterQuality.value) {
    filtered = filtered.filter(review => review.quality === filterQuality.value)
  }

  if (filterStatus.value) {
    filtered = filtered.filter(review => review.status === filterStatus.value)
  }

  return filtered
})

const paginatedReviews = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredReviews.value.slice(start, end)
})

// Methods
const getQualityClass = (quality: string) => {
  switch (quality) {
    case 'excellent':
      return 'bg-green-100 text-green-800'
    case 'good':
      return 'bg-blue-100 text-blue-800'
    case 'average':
      return 'bg-yellow-100 text-yellow-800'
    case 'poor':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getQualityText = (quality: string) => {
  switch (quality) {
    case 'excellent':
      return 'عالی'
    case 'good':
      return 'خوب'
    case 'average':
      return 'متوسط'
    case 'poor':
      return 'ضعیف'
    default:
      return 'نامشخص'
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'analyzed':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'flagged':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'analyzed':
      return 'تحلیل شده'
    case 'pending':
      return 'در انتظار'
    case 'flagged':
      return 'مشکوک'
    default:
      return 'نامشخص'
  }
}

const saveAISettings = () => {
  console.log('ذخیره تنظیمات هوش مصنوعی:', aiSettings.value)
  // API call to save AI settings
}

const testAI = () => {
  console.log('تست هوش مصنوعی')
  // Test AI analysis with sample review
}

const analyzeAllReviews = () => {
  console.log('تحلیل همه نظرات')
  // API call to analyze all reviews
}

const exportQualityReport = () => {
  console.log('خروجی گزارش کیفیت')
  // Export quality analysis report
}

const viewReviewDetails = (review: any) => {
  console.log('مشاهده جزئیات نظر:', review)
  // Open review details modal
}

const reanalyzeReview = (review: any) => {
  console.log('تحلیل مجدد نظر:', review)
  // Re-analyze specific review
}

const flagReview = (review: any) => {
  console.log('علامت‌گذاری نظر:', review)
  review.status = 'flagged'
  // API call to flag review
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
  console.log('سیستم کیفیت نظرات بارگذاری شد')
})
</script> 
