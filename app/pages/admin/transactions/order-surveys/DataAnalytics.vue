<template>
  <div class="data-analytics bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-teal-400 to-teal-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">تحلیل داده‌ها</h2>
        <p class="text-gray-600 mt-1">تحلیل پیشرفته و بصری‌سازی داده‌های نظرسنجی</p>
      </div>
    </div>

    <!-- Analytics Overview -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-gradient-to-r from-teal-50 to-teal-100 rounded-xl p-6 border border-teal-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-teal-600">کل نظرسنجی‌ها</p>
            <p class="text-3xl font-bold text-teal-900">{{ analyticsOverview.totalSurveys }}</p>
          </div>
          <div class="p-3 bg-teal-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-teal-700">
          {{ analyticsOverview.surveyGrowth }}% رشد نسبت به ماه گذشته
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">میانگین امتیاز</p>
            <p class="text-3xl font-bold text-blue-900">{{ analyticsOverview.averageRating }}/5</p>
          </div>
          <div class="p-3 bg-blue-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-blue-700">
          {{ analyticsOverview.ratingTrend }}% بهبود نسبت به ماه گذشته
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">نرخ پاسخ</p>
            <p class="text-3xl font-bold text-purple-900">{{ analyticsOverview.responseRate }}%</p>
          </div>
          <div class="p-3 bg-purple-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-purple-700">
          {{ analyticsOverview.responseGrowth }}% افزایش نسبت به ماه گذشته
        </div>
      </div>

      <div class="bg-gradient-to-r from-orange-50 to-orange-100 rounded-xl p-6 border border-orange-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-600">رضایت مشتری</p>
            <p class="text-3xl font-bold text-orange-900">{{ analyticsOverview.customerSatisfaction }}%</p>
          </div>
          <div class="p-3 bg-orange-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-orange-700">
          {{ analyticsOverview.satisfactionTrend }}% بهبود نسبت به ماه گذشته
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
      <!-- Rating Trends Chart -->
      <div class="bg-gray-50 rounded-xl p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">روند امتیازات</h3>
        <div class="h-64 bg-white rounded-lg p-6 border">
          <div class="flex items-center justify-between mb-4">
            <div class="flex space-x-4 space-x-reverse">
              <button 
                v-for="period in chartPeriods" 
                :key="period.value"
                :class="selectedPeriod === period.value ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-700'"
                class="px-3 py-1 rounded-lg text-sm font-medium transition-colors"
                @click="selectedPeriod = period.value"
              >
                {{ period.label }}
              </button>
            </div>
          </div>
          <div class="h-48 flex items-end justify-between space-x-2 space-x-reverse">
            <div v-for="(data, index) in ratingTrendData" :key="index" class="flex-1 flex flex-col items-center">
              <div class="text-xs text-gray-500 mb-1">{{ data.date }}</div>
              <div 
                class="w-full bg-gradient-to-t from-blue-500 to-blue-300 rounded-t"
                :style="{ height: `${(data.rating / 5) * 100}%` }"
              ></div>
              <div class="text-xs font-medium mt-1">{{ data.rating }}/5</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Response Rate Chart -->
      <div class="bg-gray-50 rounded-xl p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">نرخ پاسخ</h3>
        <div class="h-64 bg-white rounded-lg p-6 border">
          <div class="flex items-center justify-center h-full">
            <div class="relative">
              <svg class="w-32 h-32 transform -rotate-90">
                <circle
                  cx="64"
                  cy="64"
                  r="56"
                  stroke="currentColor"
                  stroke-width="8"
                  fill="transparent"
                  class="text-gray-200"
                />
                <circle
                  cx="64"
                  cy="64"
                  r="56"
                  stroke="currentColor"
                  stroke-width="8"
                  fill="transparent"
                  :stroke-dasharray="`${(responseRateData.percentage / 100) * 352} 352`"
                  class="text-green-500 transition-all duration-1000"
                />
              </svg>
              <div class="absolute inset-0 flex items-center justify-center">
                <div class="text-center">
                  <div class="text-2xl font-bold text-gray-900">{{ responseRateData.percentage }}%</div>
                  <div class="text-sm text-gray-600">نرخ پاسخ</div>
                </div>
              </div>
            </div>
          </div>
          <div class="mt-4 text-center">
            <div class="text-sm text-gray-600">
              {{ responseRateData.totalSent }} ارسال شده، {{ responseRateData.totalResponses }} پاسخ
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Detailed Analytics -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-8">
      <!-- Category Performance -->
      <div class="bg-gray-50 rounded-xl p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">عملکرد دسته‌بندی‌ها</h3>
        <div class="space-y-4">
          <div v-for="category in categoryPerformance" :key="category.name" class="bg-white rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <span class="font-medium text-gray-900">{{ category.name }}</span>
              <span class="text-sm font-medium text-gray-600">{{ category.rating }}/5</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full transition-all duration-500"
                :style="{ width: `${(category.rating / 5) * 100}%` }"
              ></div>
            </div>
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>{{ category.surveys }} نظرسنجی</span>
              <span>{{ category.responseRate }}% پاسخ</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Sentiment Analysis -->
      <div class="bg-gray-50 rounded-xl p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تحلیل احساسات</h3>
        <div class="space-y-4">
          <div v-for="sentiment in sentimentAnalysis" :key="sentiment.type" class="bg-white rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <div class="flex items-center">
                <div :class="getSentimentColor(sentiment.type)" class="w-3 h-3 rounded-full ml-2"></div>
                <span class="font-medium text-gray-900">{{ getSentimentText(sentiment.type) }}</span>
              </div>
              <span class="text-sm font-medium text-gray-600">{{ sentiment.percentage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                :class="getSentimentGradient(sentiment.type)"
                class="h-2 rounded-full transition-all duration-500"
                :style="{ width: `${sentiment.percentage}%` }"
              ></div>
            </div>
            <div class="text-xs text-gray-500 mt-1">
              {{ sentiment.count }} نظر
            </div>
          </div>
        </div>
      </div>

      <!-- Top Issues -->
      <div class="bg-gray-50 rounded-xl p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">مشکلات برتر</h3>
        <div class="space-y-4">
          <div v-for="issue in topIssues" :key="issue.name" class="bg-white rounded-lg p-6">
            <div class="flex items-center justify-between mb-2">
              <span class="font-medium text-gray-900">{{ issue.name }}</span>
              <span class="text-sm font-medium text-red-600">{{ issue.count }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="bg-gradient-to-r from-red-500 to-orange-500 h-2 rounded-full transition-all duration-500"
                :style="{ width: `${(issue.count / maxIssueCount) * 100}%` }"
              ></div>
            </div>
            <div class="text-xs text-gray-500 mt-1">
              {{ issue.percentage }}% از کل مشکلات
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Advanced Analytics -->
    <div class="bg-gray-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تحلیل پیشرفته</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- Correlation Analysis -->
        <div class="bg-white rounded-lg p-6 border">
          <h4 class="font-medium text-gray-900 mb-3">تحلیل همبستگی</h4>
          <div class="space-y-3">
            <div v-for="correlation in correlationAnalysis" :key="correlation.factor" class="flex items-center justify-between">
              <span class="text-sm text-gray-700">{{ correlation.factor }}</span>
              <div class="flex items-center">
                <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                  <div 
                    :class="getCorrelationColor(correlation.strength)"
                    class="h-2 rounded-full transition-all duration-500"
                    :style="{ width: `${Math.abs(correlation.strength) * 100}%` }"
                  ></div>
                </div>
                <span class="text-xs font-medium" :class="getCorrelationTextColor(correlation.strength)">
                  {{ correlation.strength > 0 ? '+' : '' }}{{ (correlation.strength * 100).toFixed(1) }}%
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Predictive Analytics -->
        <div class="bg-white rounded-lg p-6 border">
          <h4 class="font-medium text-gray-900 mb-3">پیش‌بینی آینده</h4>
          <div class="space-y-3">
            <div v-for="prediction in predictiveAnalytics" :key="prediction.metric" class="flex items-center justify-between">
              <span class="text-sm text-gray-700">{{ prediction.metric }}</span>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ prediction.currentValue }}</div>
                <div class="text-xs" :class="prediction.trend > 0 ? 'text-green-600' : 'text-red-600'">
                  {{ prediction.trend > 0 ? '+' : '' }}{{ prediction.trend }}% در ماه آینده
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Export and Reports -->
    <div class="bg-gray-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش‌ها و خروجی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <h4 class="font-medium text-gray-900 mb-3">گزارش‌های آماده</h4>
          <div class="space-y-2">
            <button 
              v-for="report in availableReports" 
              :key="report.id"
              class="w-full flex items-center justify-between p-3 bg-white rounded-lg border hover:bg-gray-50 transition-colors"
              @click="generateReport(report.id)"
            >
              <div class="flex items-center">
                <svg class="w-5 h-5 text-gray-500 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
                <span class="text-sm font-medium text-gray-900">{{ report.name }}</span>
              </div>
              <span class="text-xs text-gray-500">{{ report.format }}</span>
            </button>
          </div>
        </div>

        <div>
          <h4 class="font-medium text-gray-900 mb-3">خروجی سفارشی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">بازه زمانی</label>
              <select v-model="customExport.period" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500">
                <option value="7">۷ روز گذشته</option>
                <option value="30">۳۰ روز گذشته</option>
                <option value="90">۹۰ روز گذشته</option>
                <option value="365">یک سال گذشته</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">فرمت خروجی</label>
              <select v-model="customExport.format" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500">
                <option value="excel">Excel</option>
                <option value="csv">CSV</option>
                <option value="pdf">PDF</option>
                <option value="json">JSON</option>
              </select>
            </div>
            <button class="w-full px-4 py-2 bg-teal-600 text-white rounded-lg hover:bg-teal-700 transition-colors" @click="exportCustomData">
              ایجاد خروجی سفارشی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Action Buttons -->
    <div class="flex justify-between">
      <button class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="refreshAnalytics">
        بروزرسانی داده‌ها
      </button>
      
      <div class="flex space-x-3 space-x-reverse">
        <button class="px-6 py-3 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors" @click="scheduleReport">
          زمان‌بندی گزارش
        </button>
        <button class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors" @click="shareAnalytics">
          اشتراک‌گذاری
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

interface AnalyticsOverview {
  totalSurveys: number
  surveyGrowth: number
  averageRating: number
  ratingTrend: number
  responseRate: number
  responseGrowth: number
  customerSatisfaction: number
  satisfactionTrend: number
}

interface RatingTrendData {
  date: string
  rating: number
}

interface ResponseRateData {
  percentage: number
  totalSent: number
  totalResponses: number
}

interface CategoryPerformance {
  name: string
  rating: number
  surveys: number
  responseRate: number
}

interface SentimentAnalysis {
  type: 'positive' | 'neutral' | 'negative'
  percentage: number
  count: number
}

interface TopIssue {
  name: string
  count: number
  percentage: number
}

interface CorrelationAnalysis {
  factor: string
  strength: number
}

interface PredictiveAnalytics {
  metric: string
  currentValue: string
  trend: number
}

interface AvailableReport {
  id: number
  name: string
  format: string
}

interface CustomExport {
  period: string
  format: string
}

const analyticsOverview = ref<AnalyticsOverview>({
  totalSurveys: 1247,
  surveyGrowth: 15.3,
  averageRating: 4.2,
  ratingTrend: 8.7,
  responseRate: 78.9,
  responseGrowth: 12.4,
  customerSatisfaction: 85.6,
  satisfactionTrend: 5.2
})

const selectedPeriod = ref('30')
const chartPeriods = [
  { label: '۷ روز', value: '7' },
  { label: '۳۰ روز', value: '30' },
  { label: '۹۰ روز', value: '90' },
  { label: 'یک سال', value: '365' }
]

const ratingTrendData = ref<RatingTrendData[]>([
  { date: '۱', rating: 4.1 },
  { date: '۲', rating: 4.3 },
  { date: '۳', rating: 4.0 },
  { date: '۴', rating: 4.5 },
  { date: '۵', rating: 4.2 },
  { date: '۶', rating: 4.4 },
  { date: '۷', rating: 4.3 }
])

const responseRateData = ref<ResponseRateData>({
  percentage: 78.9,
  totalSent: 1247,
  totalResponses: 984
})

const categoryPerformance = ref<CategoryPerformance[]>([
  { name: 'کیفیت محصول', rating: 4.3, surveys: 456, responseRate: 82.1 },
  { name: 'سرعت تحویل', rating: 3.8, surveys: 423, responseRate: 75.6 },
  { name: 'خدمات پس از فروش', rating: 4.1, surveys: 389, responseRate: 79.2 },
  { name: 'قیمت‌گذاری', rating: 3.9, surveys: 367, responseRate: 76.8 },
  { name: 'بسته‌بندی', rating: 4.5, surveys: 412, responseRate: 84.3 }
])

const sentimentAnalysis = ref<SentimentAnalysis[]>([
  { type: 'positive', percentage: 65.3, count: 642 },
  { type: 'neutral', percentage: 23.7, count: 233 },
  { type: 'negative', percentage: 11.0, count: 108 }
])

const topIssues = ref<TopIssue[]>([
  { name: 'تأخیر در تحویل', count: 89, percentage: 23.4 },
  { name: 'مشکل در بسته‌بندی', count: 67, percentage: 17.6 },
  { name: 'کیفیت پایین', count: 54, percentage: 14.2 },
  { name: 'قیمت بالا', count: 43, percentage: 11.3 },
  { name: 'خدمات ضعیف', count: 38, percentage: 10.0 }
])

const correlationAnalysis = ref<CorrelationAnalysis[]>([
  { factor: 'سرعت تحویل', strength: 0.78 },
  { factor: 'کیفیت محصول', strength: 0.65 },
  { factor: 'قیمت', strength: -0.42 },
  { factor: 'خدمات پس از فروش', strength: 0.56 },
  { factor: 'بسته‌بندی', strength: 0.34 }
])

const predictiveAnalytics = ref<PredictiveAnalytics[]>([
  { metric: 'نرخ پاسخ', currentValue: '78.9%', trend: 5.2 },
  { metric: 'میانگین امتیاز', currentValue: '4.2/5', trend: 3.1 },
  { metric: 'رضایت مشتری', currentValue: '85.6%', trend: 7.8 },
  { metric: 'تعداد نظرسنجی', currentValue: '1247', trend: 12.4 }
])

const availableReports = ref<AvailableReport[]>([
  { id: 1, name: 'گزارش ماهانه کلی', format: 'PDF' },
  { id: 2, name: 'تحلیل روند امتیازات', format: 'Excel' },
  { id: 3, name: 'گزارش مشکلات برتر', format: 'CSV' },
  { id: 4, name: 'تحلیل احساسات', format: 'PDF' }
])

const customExport = ref<CustomExport>({
  period: '30',
  format: 'excel'
})

const maxIssueCount = computed(() => {
  return Math.max(...topIssues.value.map(issue => issue.count))
})

// Methods
const getSentimentColor = (type: string) => {
  switch (type) {
    case 'positive':
      return 'bg-green-500'
    case 'neutral':
      return 'bg-yellow-500'
    case 'negative':
      return 'bg-red-500'
    default:
      return 'bg-gray-500'
  }
}

const getSentimentText = (type: string) => {
  switch (type) {
    case 'positive':
      return 'مثبت'
    case 'neutral':
      return 'خنثی'
    case 'negative':
      return 'منفی'
    default:
      return 'نامشخص'
  }
}

const getSentimentGradient = (type: string) => {
  switch (type) {
    case 'positive':
      return 'bg-gradient-to-r from-green-500 to-green-400'
    case 'neutral':
      return 'bg-gradient-to-r from-yellow-500 to-yellow-400'
    case 'negative':
      return 'bg-gradient-to-r from-red-500 to-red-400'
    default:
      return 'bg-gradient-to-r from-gray-500 to-gray-400'
  }
}

const getCorrelationColor = (strength: number) => {
  if (strength > 0.5) return 'bg-green-500'
  if (strength > 0.2) return 'bg-blue-500'
  if (strength > -0.2) return 'bg-yellow-500'
  return 'bg-red-500'
}

const getCorrelationTextColor = (strength: number) => {
  if (strength > 0.5) return 'text-green-600'
  if (strength > 0.2) return 'text-blue-600'
  if (strength > -0.2) return 'text-yellow-600'
  return 'text-red-600'
}

const generateReport = (_reportId: number) => {
  // console.log('Generating report:', reportId)
}

const exportCustomData = () => {
  // console.log('Exporting custom data:', customExport.value)
}

const refreshAnalytics = () => {
  // console.log('Refreshing analytics data...')
}

const scheduleReport = () => {
  // console.log('Scheduling report...')
}

const shareAnalytics = () => {
  // console.log('Sharing analytics...')
}

// Expose methods for parent component
defineExpose({
  analyticsOverview,
  ratingTrendData,
  responseRateData,
  categoryPerformance,
  sentimentAnalysis,
  topIssues,
  correlationAnalysis,
  predictiveAnalytics
})
</script>

<style scoped>
.data-analytics {
  transition: all 0.3s ease;
}

/* Chart animations */
.bg-gradient-to-t {
  transition: height 0.5s ease;
}

/* Progress bar animations */
.h-2 {
  transition: width 0.5s ease;
}

/* Button hover effects */
button:hover {
  transform: translateY(-1px);
  transition: transform 0.2s ease;
}

/* Chart period buttons */
.space-x-reverse > * + * {
  margin-right: 0.5rem;
  margin-left: 0;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .grid {
    grid-template-columns: 1fr;
  }
  
  .h-64 {
    height: 12rem;
  }
  
  .h-48 {
    height: 8rem;
  }
}
</style> 
