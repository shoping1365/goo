<template>
  <div class="rating-stats bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
    <div class="flex items-center justify-between mb-8">
      <div class="flex items-center space-x-3 space-x-reverse">
        <div class="p-3 bg-gradient-to-r from-yellow-400 to-yellow-600 rounded-xl shadow-lg">
          <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.75l-6.172 3.247 1.179-6.873L2 9.753l6.908-1.004L12 2.25l3.092 6.499L22 9.753l-5.007 4.371 1.179 6.873z"></path>
          </svg>
        </div>
        <div>
          <h2 class="text-2xl font-bold text-gray-900">آمار امتیازات</h2>
          <p class="text-gray-600 mt-1">تحلیل دقیق امتیازات و روندها</p>
        </div>
      </div>
      
      <div class="flex space-x-3 space-x-reverse">
        <select v-model="selectedPeriod" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500">
          <option value="7">۷ روز گذشته</option>
          <option value="30">۳۰ روز گذشته</option>
          <option value="90">۹۰ روز گذشته</option>
          <option value="365">یک سال گذشته</option>
        </select>
      </div>
    </div>

    <!-- Overall Rating Summary -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-gradient-to-r from-yellow-50 to-yellow-100 rounded-xl p-6 border border-yellow-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-yellow-600">میانگین کلی</p>
            <p class="text-3xl font-bold text-yellow-900">{{ stats.averageRating.toFixed(1) }}</p>
          </div>
          <div class="p-3 bg-yellow-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.75l-6.172 3.247 1.179-6.873L2 9.753l6.908-1.004L12 2.25l3.092 6.499L22 9.753l-5.007 4.371 1.179 6.873z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 flex items-center">
          <div class="flex space-x-1 space-x-reverse">
            <span v-for="i in 5" :key="i" class="text-yellow-400">
              {{ i <= Math.round(stats.averageRating) ? '★' : '☆' }}
            </span>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-600">رضایت بالا</p>
            <p class="text-3xl font-bold text-green-900">{{ stats.highSatisfaction }}%</p>
          </div>
          <div class="p-3 bg-green-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-green-700">
          امتیاز ۴ و ۵ ستاره
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">تعداد کل</p>
            <p class="text-3xl font-bold text-blue-900">{{ stats.totalRatings }}</p>
          </div>
          <div class="p-3 bg-blue-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-blue-700">
          امتیاز ثبت شده
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">روند</p>
            <p class="text-3xl font-bold text-purple-900">{{ stats.trend >= 0 ? '+' : '' }}{{ stats.trend }}%</p>
          </div>
          <div class="p-3 bg-purple-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-purple-700">
          نسبت به دوره قبل
        </div>
      </div>
    </div>

    <!-- Rating Distribution -->
    <div class="bg-gray-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع امتیازات</h3>
      <div class="space-y-4">
        <div v-for="rating in [5, 4, 3, 2, 1]" :key="rating" class="flex items-center space-x-4 space-x-reverse">
          <div class="flex items-center space-x-2 space-x-reverse w-20">
            <span class="text-sm font-medium">{{ rating }}</span>
            <span class="text-yellow-400">★</span>
          </div>
          <div class="flex-1 bg-gray-200 rounded-full h-4">
            <div 
              class="bg-yellow-400 h-4 rounded-full transition-all duration-500"
              :style="{ width: `${getRatingPercentage(rating)}%` }"
            ></div>
          </div>
          <div class="flex items-center space-x-4 space-x-reverse w-32">
            <span class="text-sm text-gray-600">{{ getRatingCount(rating) }}</span>
            <span class="text-sm text-gray-500">{{ getRatingPercentage(rating) }}%</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Detailed Category Ratings -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
      <div class="bg-white rounded-xl border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">امتیازات جزئی</h3>
        <div class="space-y-4">
          <div v-for="(rating, category) in stats.categoryRatings" :key="category" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 rounded-lg" :class="getCategoryColor(category)">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <span class="text-sm font-medium text-gray-700">{{ getCategoryName(category) }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="flex space-x-1 space-x-reverse">
                <span v-for="i in 5" :key="i" class="text-yellow-400 text-sm">
                  {{ i <= rating ? '★' : '☆' }}
                </span>
              </div>
              <span class="text-sm font-bold text-gray-900">{{ rating.toFixed(1) }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">روند امتیازات</h3>
        <div class="h-48 flex items-end space-x-2 space-x-reverse">
          <div v-for="(data, index) in stats.ratingTrend" :key="index" class="flex-1 flex flex-col items-center">
            <div 
              class="w-full bg-gradient-to-t from-yellow-500 to-yellow-300 rounded-t transition-all duration-300 hover:from-yellow-600 hover:to-yellow-400"
              :style="{ height: `${(data.value / 5) * 100}%` }"
            ></div>
            <span class="text-xs text-gray-600 mt-2">{{ data.label }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Rating Insights -->
    <div class="bg-gradient-to-r from-indigo-50 to-purple-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">نکات کلیدی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="insight in stats.insights" :key="insight.id" class="bg-white rounded-lg p-6 border border-indigo-200">
          <div class="flex items-start space-x-3 space-x-reverse">
            <div class="p-2 bg-indigo-500 rounded-lg">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">{{ insight.title }}</h4>
              <p class="text-sm text-gray-600 mt-1">{{ insight.description }}</p>
              <div class="mt-2">
                <span class="text-xs text-indigo-600 font-medium">{{ insight.change }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Top and Bottom Performers -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div class="bg-white rounded-xl border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">بهترین عملکردها</h3>
        <div class="space-y-3">
          <div v-for="performer in stats.topPerformers" :key="performer.id" class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-green-500 rounded-lg">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
              </div>
              <div>
                <div class="text-sm font-medium text-gray-900">{{ performer.name }}</div>
                <div class="text-xs text-gray-500">{{ performer.category }}</div>
              </div>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="flex space-x-1 space-x-reverse">
                <span v-for="i in 5" :key="i" class="text-yellow-400 text-sm">
                  {{ i <= performer.rating ? '★' : '☆' }}
                </span>
              </div>
              <span class="text-sm font-bold text-green-600">{{ performer.rating.toFixed(1) }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">نیازمند بهبود</h3>
        <div class="space-y-3">
          <div v-for="performer in stats.bottomPerformers" :key="performer.id" class="flex items-center justify-between p-3 bg-red-50 rounded-lg">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-red-500 rounded-lg">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                </svg>
              </div>
              <div>
                <div class="text-sm font-medium text-gray-900">{{ performer.name }}</div>
                <div class="text-xs text-gray-500">{{ performer.category }}</div>
              </div>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="flex space-x-1 space-x-reverse">
                <span v-for="i in 5" :key="i" class="text-yellow-400 text-sm">
                  {{ i <= performer.rating ? '★' : '☆' }}
                </span>
              </div>
              <span class="text-sm font-bold text-red-600">{{ performer.rating.toFixed(1) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

interface RatingStats {
  averageRating: number
  highSatisfaction: number
  totalRatings: number
  trend: number
  ratingDistribution: Record<number, number>
  categoryRatings: Record<string, number>
  ratingTrend: Array<{ label: string; value: number }>
  insights: Array<{
    id: number
    title: string
    description: string
    change: string
  }>
  topPerformers: Array<{
    id: number
    name: string
    category: string
    rating: number
  }>
  bottomPerformers: Array<{
    id: number
    name: string
    category: string
    rating: number
  }>
}

const selectedPeriod = ref('30')

// Sample data
const stats = ref<RatingStats>({
  averageRating: 4.2,
  highSatisfaction: 78.3,
  totalRatings: 1247,
  trend: 8.1,
  ratingDistribution: {
    5: 456,
    4: 389,
    3: 234,
    2: 98,
    1: 70
  },
  categoryRatings: {
    productQuality: 4.3,
    pricing: 3.9,
    deliverySpeed: 4.1,
    packaging: 4.4,
    afterSales: 3.8,
    userInterface: 4.2
  },
  ratingTrend: [
    { label: 'هفته ۱', value: 4.1 },
    { label: 'هفته ۲', value: 4.2 },
    { label: 'هفته ۳', value: 4.0 },
    { label: 'هفته ۴', value: 4.3 }
  ],
  insights: [
    {
      id: 1,
      title: 'بهبود کیفیت محصولات',
      description: 'امتیاز کیفیت محصولات ۰.۳ واحد افزایش یافته است.',
      change: '+۰.۳ واحد'
    },
    {
      id: 2,
      title: 'کاهش رضایت از قیمت',
      description: 'امتیاز قیمت‌گذاری ۰.۲ واحد کاهش یافته است.',
      change: '-۰.۲ واحد'
    },
    {
      id: 3,
      title: 'افزایش رضایت کلی',
      description: 'میانگین امتیاز کلی ۰.۱ واحد افزایش یافته است.',
      change: '+۰.۱ واحد'
    },
    {
      id: 4,
      title: 'بهبود خدمات پس از فروش',
      description: 'امتیاز خدمات پس از فروش ۰.۴ واحد افزایش یافته است.',
      change: '+۰.۴ واحد'
    }
  ],
  topPerformers: [
    {
      id: 1,
      name: 'لپ تاپ اپل مک‌بوک پرو',
      category: 'الکترونیک',
      rating: 4.8
    },
    {
      id: 2,
      name: 'کتاب مدیریت زمان',
      category: 'کتاب',
      rating: 4.7
    },
    {
      id: 3,
      name: 'کیف چرمی برند',
      category: 'پوشاک',
      rating: 4.6
    }
  ],
  bottomPerformers: [
    {
      id: 1,
      name: 'ساعت هوشمند X',
      category: 'الکترونیک',
      rating: 2.8
    },
    {
      id: 2,
      name: 'کفش ورزشی Y',
      category: 'پوشاک',
      rating: 3.1
    },
    {
      id: 3,
      name: 'لوازم آشپزخانه Z',
      category: 'خانه',
      rating: 3.2
    }
  ]
})

// Methods
const getRatingPercentage = (rating: number) => {
  const total = Object.values(stats.value.ratingDistribution).reduce((sum, count) => sum + count, 0)
  return Math.round((stats.value.ratingDistribution[rating] / total) * 100)
}

const getRatingCount = (rating: number) => {
  return stats.value.ratingDistribution[rating] || 0
}

const getCategoryName = (key: string) => {
  const names = {
    productQuality: 'کیفیت محصول',
    pricing: 'قیمت‌گذاری',
    deliverySpeed: 'سرعت تحویل',
    packaging: 'بسته‌بندی',
    afterSales: 'پشتیبانی',
    userInterface: 'رابط کاربری'
  }
  return names[key] || key
}

const getCategoryColor = (category: string) => {
  const colors = {
    productQuality: 'bg-green-500',
    pricing: 'bg-blue-500',
    deliverySpeed: 'bg-purple-500',
    packaging: 'bg-orange-500',
    afterSales: 'bg-red-500',
    userInterface: 'bg-teal-500'
  }
  return colors[category] || 'bg-gray-500'
}

// Watch for period changes
watch(selectedPeriod, () => {
  // Update stats based on selected period
  // console.log('Period changed to:', selectedPeriod.value)
})
</script>

<style scoped>
.rating-stats {
  transition: all 0.3s ease;
}

/* Chart animations */
.bg-yellow-400,
.bg-gradient-to-t {
  transition: all 0.5s ease-out;
}

/* Hover effects */
.hover\:from-yellow-600:hover {
  background: linear-gradient(to top, #d97706, #f59e0b);
}

.hover\:to-yellow-400:hover {
  background: linear-gradient(to top, #f59e0b, #eab308);
}

/* Smooth transitions */
.transition-all {
  transition: all 0.3s ease-in-out;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .grid {
    grid-template-columns: 1fr;
  }
}
</style> 
