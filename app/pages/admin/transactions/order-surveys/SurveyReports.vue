<template>
  <div class="survey-reports bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
    <div class="flex items-center justify-between mb-8">
      <div class="flex items-center space-x-3 space-x-reverse">
        <div class="p-3 bg-gradient-to-r from-purple-400 to-purple-600 rounded-xl shadow-lg">
          <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
        </div>
        <div>
          <h2 class="text-2xl font-bold text-gray-900">گزارش‌های نظرسنجی</h2>
          <p class="text-gray-600 mt-1">تحلیل جامع و گزارش‌های آماری نظرسنجی‌ها</p>
        </div>
      </div>
      
      <div class="flex space-x-3 space-x-reverse">
        <button @click="generateReport" class="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors flex items-center space-x-2 space-x-reverse">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          <span>تولید گزارش</span>
        </button>
        
        <button @click="exportReport" class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors">
          خروجی PDF
        </button>
      </div>
    </div>

    <!-- Report Filters -->
    <div class="bg-gray-50 rounded-xl p-6 mb-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
          <select v-model="reportType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
            <option value="overview">گزارش کلی</option>
            <option value="detailed">گزارش تفصیلی</option>
            <option value="trends">گزارش روند</option>
            <option value="comparison">گزارش مقایسه‌ای</option>
            <option value="custom">گزارش سفارشی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select v-model="timeRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
            <option value="7">۷ روز گذشته</option>
            <option value="30">۳۰ روز گذشته</option>
            <option value="90">۹۰ روز گذشته</option>
            <option value="365">یک سال گذشته</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
          <select v-model="category" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
            <option value="">همه دسته‌ها</option>
            <option value="electronics">الکترونیک</option>
            <option value="clothing">پوشاک</option>
            <option value="books">کتاب</option>
            <option value="home">خانه و آشپزخانه</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرمت خروجی</label>
          <select v-model="exportFormat" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
            <option value="pdf">PDF</option>
            <option value="excel">Excel</option>
            <option value="csv">CSV</option>
            <option value="json">JSON</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Report Content -->
    <div v-if="currentReport" class="space-y-8">
      <!-- Executive Summary -->
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-6 border border-blue-200">
        <h3 class="text-xl font-bold text-gray-900 mb-4">خلاصه اجرایی</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="text-center">
            <div class="text-3xl font-bold text-blue-600">{{ currentReport.totalSurveys }}</div>
            <div class="text-sm text-gray-600">کل نظرسنجی‌ها</div>
          </div>
          <div class="text-center">
            <div class="text-3xl font-bold text-green-600">{{ currentReport.responseRate }}%</div>
            <div class="text-sm text-gray-600">نرخ پاسخ</div>
          </div>
          <div class="text-center">
            <div class="text-3xl font-bold text-purple-600">{{ currentReport.averageRating.toFixed(1) }}</div>
            <div class="text-sm text-gray-600">میانگین امتیاز</div>
          </div>
        </div>
      </div>

      <!-- Rating Analysis -->
      <div class="bg-white rounded-xl border border-gray-200 p-6">
        <h3 class="text-xl font-bold text-gray-900 mb-4">تحلیل امتیازات</h3>
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <!-- Overall Rating Distribution -->
          <div>
            <h4 class="text-lg font-semibold text-gray-800 mb-4">توزیع امتیازات کلی</h4>
            <div class="space-y-3">
              <div v-for="rating in [5, 4, 3, 2, 1]" :key="rating" class="flex items-center space-x-3 space-x-reverse">
                <div class="flex items-center space-x-1 space-x-reverse w-16">
                  <span class="text-sm font-medium">{{ rating }}</span>
                  <span class="text-yellow-400">★</span>
                </div>
                <div class="flex-1 bg-gray-200 rounded-full h-4">
                  <div 
                    class="bg-yellow-400 h-4 rounded-full transition-all duration-500"
                    :style="{ width: `${getRatingPercentage(rating)}%` }"
                  ></div>
                </div>
                <span class="text-sm text-gray-600 w-16">{{ getRatingCount(rating) }}</span>
                <span class="text-sm text-gray-500 w-12">{{ getRatingPercentage(rating) }}%</span>
              </div>
            </div>
          </div>

          <!-- Detailed Ratings -->
          <div>
            <h4 class="text-lg font-semibold text-gray-800 mb-4">امتیازات جزئی</h4>
            <div class="space-y-3">
              <div v-for="(rating, category) in currentReport.detailedRatings" :key="category" class="flex items-center justify-between">
                <span class="text-sm text-gray-700">{{ getCategoryName(category) }}</span>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div class="flex space-x-1 space-x-reverse">
                    <span v-for="i in 5" :key="i" class="text-yellow-400 text-sm">
                      {{ i <= rating ? '★' : '☆' }}
                    </span>
                  </div>
                  <span class="text-sm font-medium text-gray-900">{{ rating.toFixed(1) }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Trends Analysis -->
      <div class="bg-white rounded-xl border border-gray-200 p-6">
        <h3 class="text-xl font-bold text-gray-900 mb-4">تحلیل روند</h3>
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <!-- Response Trend -->
          <div>
            <h4 class="text-lg font-semibold text-gray-800 mb-4">روند پاسخ‌دهی</h4>
            <div class="h-48 flex items-end space-x-2 space-x-reverse">
              <div v-for="(data, index) in currentReport.responseTrend" :key="index" class="flex-1 flex flex-col items-center">
                <div 
                  class="w-full bg-gradient-to-t from-blue-500 to-blue-300 rounded-t transition-all duration-300"
                  :style="{ height: `${(data.value / maxTrendValue) * 100}%` }"
                ></div>
                <span class="text-xs text-gray-600 mt-2">{{ data.label }}</span>
              </div>
            </div>
          </div>

          <!-- Rating Trend -->
          <div>
            <h4 class="text-lg font-semibold text-gray-800 mb-4">روند امتیازات</h4>
            <div class="h-48 flex items-end space-x-2 space-x-reverse">
              <div v-for="(data, index) in currentReport.ratingTrend" :key="index" class="flex-1 flex flex-col items-center">
                <div 
                  class="w-full bg-gradient-to-t from-green-500 to-green-300 rounded-t transition-all duration-300"
                  :style="{ height: `${(data.value / 5) * 100}%` }"
                ></div>
                <span class="text-xs text-gray-600 mt-2">{{ data.label }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Category Analysis -->
      <div class="bg-white rounded-xl border border-gray-200 p-6">
        <h3 class="text-xl font-bold text-gray-900 mb-4">تحلیل بر اساس دسته‌بندی</h3>
        <div class="overflow-x-auto">
          <table class="w-full">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">دسته‌بندی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">تعداد نظرسنجی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">میانگین امتیاز</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">نرخ پاسخ</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">رضایت</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="category in currentReport.categoryAnalysis" :key="category.name" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ category.name }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ category.surveyCount }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="flex space-x-1 space-x-reverse">
                      <span v-for="i in 5" :key="i" class="text-yellow-400 text-sm">
                        {{ i <= category.averageRating ? '★' : '☆' }}
                      </span>
                    </div>
                    <span class="text-sm text-gray-900 mr-2">({{ category.averageRating.toFixed(1) }})</span>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ category.responseRate }}%
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="getSatisfactionClass(category.satisfaction)" class="px-2 py-1 text-xs font-medium rounded-full">
                    {{ category.satisfaction }}%
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Key Insights -->
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-200">
        <h3 class="text-xl font-bold text-gray-900 mb-4">نکات کلیدی</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div v-for="insight in currentReport.keyInsights" :key="insight.id" class="bg-white rounded-lg p-6 border border-green-200">
            <div class="flex items-start space-x-3 space-x-reverse">
              <div class="p-2 bg-green-500 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <div>
                <h4 class="font-medium text-gray-900">{{ insight.title }}</h4>
                <p class="text-sm text-gray-600 mt-1">{{ insight.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Recommendations -->
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-6 border border-blue-200">
        <h3 class="text-xl font-bold text-gray-900 mb-4">توصیه‌ها</h3>
        <div class="space-y-4">
          <div v-for="recommendation in currentReport.recommendations" :key="recommendation.id" class="flex items-start space-x-3 space-x-reverse">
            <div class="p-2 bg-blue-500 rounded-lg">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">{{ recommendation.title }}</h4>
              <p class="text-sm text-gray-600 mt-1">{{ recommendation.description }}</p>
              <div class="mt-2">
                <span class="text-xs text-blue-600 font-medium">اولویت: {{ recommendation.priority }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-purple-600 mx-auto"></div>
      <p class="text-gray-600 mt-4">در حال تولید گزارش...</p>
    </div>

    <!-- No Report State -->
    <div v-if="!currentReport && !loading" class="text-center py-12">
      <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
      </svg>
      <p class="text-gray-600">برای مشاهده گزارش، ابتدا یک گزارش تولید کنید</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface ReportData {
  totalSurveys: number
  responseRate: number
  averageRating: number
  ratingDistribution: Record<number, number>
  detailedRatings: Record<string, number>
  responseTrend: Array<{ label: string; value: number }>
  ratingTrend: Array<{ label: string; value: number }>
  categoryAnalysis: Array<{
    name: string
    surveyCount: number
    averageRating: number
    responseRate: number
    satisfaction: number
  }>
  keyInsights: Array<{
    id: number
    title: string
    description: string
  }>
  recommendations: Array<{
    id: number
    title: string
    description: string
    priority: string
  }>
}

const reportType = ref('overview')
const timeRange = ref('30')
const category = ref('')
const exportFormat = ref('pdf')
const loading = ref(false)
const currentReport = ref<ReportData | null>(null)

// Sample report data
const sampleReport: ReportData = {
  totalSurveys: 1247,
  responseRate: 78.3,
  averageRating: 4.2,
  ratingDistribution: {
    5: 456,
    4: 389,
    3: 234,
    2: 98,
    1: 70
  },
  detailedRatings: {
    productQuality: 4.3,
    pricing: 3.9,
    deliverySpeed: 4.1,
    packaging: 4.4,
    afterSales: 3.8,
    userInterface: 4.2
  },
  responseTrend: [
    { label: 'شنبه', value: 45 },
    { label: 'یکشنبه', value: 52 },
    { label: 'دوشنبه', value: 38 },
    { label: 'سه‌شنبه', value: 61 },
    { label: 'چهارشنبه', value: 48 },
    { label: 'پنج‌شنبه', value: 55 },
    { label: 'جمعه', value: 42 }
  ],
  ratingTrend: [
    { label: 'هفته ۱', value: 4.1 },
    { label: 'هفته ۲', value: 4.2 },
    { label: 'هفته ۳', value: 4.0 },
    { label: 'هفته ۴', value: 4.3 }
  ],
  categoryAnalysis: [
    {
      name: 'الکترونیک',
      surveyCount: 456,
      averageRating: 4.3,
      responseRate: 82.1,
      satisfaction: 85.2
    },
    {
      name: 'پوشاک',
      surveyCount: 234,
      averageRating: 4.0,
      responseRate: 75.6,
      satisfaction: 78.9
    },
    {
      name: 'کتاب',
      surveyCount: 189,
      averageRating: 4.4,
      responseRate: 88.3,
      satisfaction: 91.7
    },
    {
      name: 'خانه و آشپزخانه',
      surveyCount: 368,
      averageRating: 4.1,
      responseRate: 79.2,
      satisfaction: 82.4
    }
  ],
  keyInsights: [
    {
      id: 1,
      title: 'افزایش رضایت مشتریان',
      description: 'رضایت مشتریان در ۳۰ روز گذشته ۸.۱٪ افزایش یافته است.'
    },
    {
      id: 2,
      title: 'بهبود سرعت تحویل',
      description: 'امتیاز سرعت تحویل از ۳.۸ به ۴.۱ بهبود یافته است.'
    },
    {
      id: 3,
      title: 'کاهش شکایات',
      description: 'تعداد شکایات ۱۵٪ کاهش یافته است.'
    },
    {
      id: 4,
      title: 'افزایش نرخ پاسخ',
      description: 'نرخ پاسخ به نظرسنجی‌ها ۵.۲٪ افزایش یافته است.'
    }
  ],
  recommendations: [
    {
      id: 1,
      title: 'بهبود خدمات پس از فروش',
      description: 'با توجه به امتیاز پایین خدمات پس از فروش، پیشنهاد می‌شود تیم پشتیبانی تقویت شود.',
      priority: 'بالا'
    },
    {
      id: 2,
      title: 'بهینه‌سازی قیمت‌گذاری',
      description: 'امتیاز قیمت‌گذاری نیاز به بهبود دارد. بررسی استراتژی قیمت‌گذاری توصیه می‌شود.',
      priority: 'متوسط'
    },
    {
      id: 3,
      title: 'ارتقای رابط کاربری',
      description: 'رابط کاربری سایت نیاز به بهبود دارد تا تجربه کاربری بهتری ارائه شود.',
      priority: 'متوسط'
    }
  ]
}

// Computed properties
const maxTrendValue = computed(() => {
  if (!currentReport.value) return 0
  return Math.max(...currentReport.value.responseTrend.map(item => item.value))
})

// Methods
const getRatingPercentage = (rating: number) => {
  if (!currentReport.value) return 0
  const total = Object.values(currentReport.value.ratingDistribution).reduce((sum, count) => sum + count, 0)
  return Math.round((currentReport.value.ratingDistribution[rating] / total) * 100)
}

const getRatingCount = (rating: number) => {
  if (!currentReport.value) return 0
  return currentReport.value.ratingDistribution[rating] || 0
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

const getSatisfactionClass = (satisfaction: number) => {
  if (satisfaction >= 90) return 'bg-green-100 text-green-800'
  if (satisfaction >= 80) return 'bg-blue-100 text-blue-800'
  if (satisfaction >= 70) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

const generateReport = async () => {
  loading.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 2000))
    currentReport.value = sampleReport
  } catch (error) {
    console.error('Error generating report:', error)
  } finally {
    loading.value = false
  }
}

const exportReport = () => {
  if (!currentReport.value) {
    alert('ابتدا یک گزارش تولید کنید')
    return
  }
  
  // Implement export functionality
  console.log('Exporting report in', exportFormat.value, 'format')
}
</script>

<style scoped>
.survey-reports {
  transition: all 0.3s ease;
}

/* Chart animations */
.bg-yellow-400,
.bg-gradient-to-t {
  transition: all 0.5s ease-out;
}

/* Hover effects */
.hover\:bg-gray-50:hover {
  background-color: #f9fafb;
}

/* Loading animation */
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .grid {
    grid-template-columns: 1fr;
  }
  
  .overflow-x-auto {
    overflow-x: auto;
  }
}
</style> 
