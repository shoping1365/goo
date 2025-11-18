<template>
  <div class="bg-white rounded-xl shadow-lg px-4 py-4 mb-8" dir="rtl">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-xl font-bold text-gray-800">داشبورد آماری</h3>
        <p class="text-gray-600 text-sm">آمار کلی سیستم نظرسنجی</p>
      </div>
      <div class="flex space-x-3 space-x-reverse">
        <select v-model="selectedPeriod" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
          <option value="today">امروز</option>
          <option value="week">هفته جاری</option>
          <option value="month">ماه جاری</option>
          <option value="year">سال جاری</option>
        </select>
        <button @click="refreshStats" class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700">
          بروزرسانی
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gapx-4 py-4 mb-8">
      <div class="stat-card bg-gradient-to-br from-blue-500 to-blue-600">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-white">{{ stats.totalOrders }}</div>
            <div class="text-blue-100 text-sm">کل سفارشات</div>
          </div>
          <div class="p-3 bg-white/20 rounded-full">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-blue-100 text-xs">
          <span :class="stats.ordersGrowth >= 0 ? 'text-green-300' : 'text-red-300'">
            {{ stats.ordersGrowth >= 0 ? '+' : '' }}{{ stats.ordersGrowth }}%
          </span>
          نسبت به دوره قبل
        </div>
      </div>

      <div class="stat-card bg-gradient-to-br from-indigo-500 to-indigo-600">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-white">{{ stats.sentSMS }}</div>
            <div class="text-indigo-100 text-sm">SMS ارسال شده</div>
          </div>
          <div class="p-3 bg-white/20 rounded-full">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-indigo-100 text-xs">
          <span :class="stats.smsGrowth >= 0 ? 'text-green-300' : 'text-red-300'">
            {{ stats.smsGrowth >= 0 ? '+' : '' }}{{ stats.smsGrowth }}%
          </span>
          نسبت به دوره قبل
        </div>
      </div>

      <div class="stat-card bg-gradient-to-br from-green-500 to-green-600">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-white">{{ stats.successfulSMS }}</div>
            <div class="text-green-100 text-sm">SMS موفق</div>
          </div>
          <div class="p-3 bg-white/20 rounded-full">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-green-100 text-xs">
          نرخ موفقیت: {{ stats.successRate }}%
        </div>
      </div>

      <div class="stat-card bg-gradient-to-br from-red-500 to-red-600">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-white">{{ stats.failedSMS }}</div>
            <div class="text-red-100 text-sm">SMS ناموفق</div>
          </div>
          <div class="p-3 bg-white/20 rounded-full">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-red-100 text-xs">
          نرخ خطا: {{ stats.errorRate }}%
        </div>
      </div>

      <div class="stat-card bg-gradient-to-br from-yellow-500 to-yellow-600">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-white">{{ stats.responseRate }}%</div>
            <div class="text-yellow-100 text-sm">نرخ پاسخ‌دهی</div>
          </div>
          <div class="p-3 bg-white/20 rounded-full">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-yellow-100 text-xs">
          {{ stats.totalResponses }} پاسخ دریافت شده
        </div>
      </div>

      <div class="stat-card bg-gradient-to-br from-orange-500 to-orange-600">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-white">{{ stats.averageRating }}/5</div>
            <div class="text-orange-100 text-sm">میانگین امتیاز</div>
          </div>
          <div class="p-3 bg-white/20 rounded-full">
            <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-2 text-orange-100 text-xs">
          {{ stats.totalRatings }} امتیاز ثبت شده
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4">
      <!-- Trend Chart -->
      <div class="bg-gray-50 rounded-lg px-4 py-4">
        <h4 class="text-lg font-semibold mb-4 text-gray-800">روند ارسال SMS</h4>
        <div class="h-64">
          <TrendChart :data="trendData" />
        </div>
      </div>

      <!-- Response Distribution -->
      <div class="bg-gray-50 rounded-lg px-4 py-4">
        <h4 class="text-lg font-semibold mb-4 text-gray-800">توزیع پاسخ‌ها</h4>
        <div class="h-64">
          <ResponsePieChart :data="responseData" />
        </div>
      </div>
    </div>

    <!-- Quick Stats -->
    <div class="mt-6 grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
      <div class="bg-blue-50 rounded-lg px-4 py-4">
        <div class="text-sm text-blue-600 font-medium">بیشترین پاسخ‌دهی</div>
        <div class="text-lg font-bold text-blue-800">{{ stats.bestResponseTime }}</div>
        <div class="text-xs text-blue-600">ساعت {{ stats.bestHour }}</div>
      </div>
      
      <div class="bg-green-50 rounded-lg px-4 py-4">
        <div class="text-sm text-green-600 font-medium">بهترین محصول</div>
        <div class="text-lg font-bold text-green-800">{{ stats.bestProduct }}</div>
        <div class="text-xs text-green-600">امتیاز {{ stats.bestProductRating }}/5</div>
      </div>
      
      <div class="bg-purple-50 rounded-lg px-4 py-4">
        <div class="text-sm text-purple-600 font-medium">منطقه فعال</div>
        <div class="text-lg font-bold text-purple-800">{{ stats.activeRegion }}</div>
        <div class="text-xs text-purple-600">{{ stats.regionOrders }} سفارش</div>
      </div>
      
      <div class="bg-orange-50 rounded-lg px-4 py-4">
        <div class="text-sm text-orange-600 font-medium">نرخ بهبود</div>
        <div class="text-lg font-bold text-orange-800">{{ stats.improvementRate }}%</div>
        <div class="text-xs text-orange-600">نسبت به ماه قبل</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, watch } from 'vue'
import TrendChart from './TrendChart.vue'
import ResponsePieChart from './ResponsePieChart.vue'

// Reactive data
const selectedPeriod = ref('month')
const loading = ref(false)

const stats = reactive({
  totalOrders: 0,
  ordersGrowth: 0,
  sentSMS: 0,
  smsGrowth: 0,
  successfulSMS: 0,
  successRate: 0,
  failedSMS: 0,
  errorRate: 0,
  responseRate: 0,
  totalResponses: 0,
  averageRating: 0,
  totalRatings: 0,
  bestResponseTime: '',
  bestHour: '',
  bestProduct: '',
  bestProductRating: 0,
  activeRegion: '',
  regionOrders: 0,
  improvementRate: 0
})

const trendData = ref({
  labels: ['فروردین', 'اردیبهشت', 'خرداد', 'تیر', 'مرداد', 'شهریور'],
  datasets: [
    {
      label: 'SMS ارسال شده',
      data: [120, 150, 180, 200, 220, 250],
      borderColor: '#6366f1',
      backgroundColor: 'rgba(99,102,241,0.1)',
      tension: 0.4
    },
    {
      label: 'پاسخ‌های دریافتی',
      data: [80, 100, 120, 140, 160, 180],
      borderColor: '#10b981',
      backgroundColor: 'rgba(16,185,129,0.1)',
      tension: 0.4
    }
  ]
})

const responseData = ref({
  labels: ['پاسخ داده', 'پاسخ نداده', 'پاسخ ناقص'],
  datasets: [{
    data: [65, 25, 10],
    backgroundColor: ['#10b981', '#6b7280', '#f59e0b']
  }]
})

// Methods
const loadStats = async () => {
  loading.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Mock data based on selected period
    const mockData = {
      today: {
        totalOrders: 45,
        ordersGrowth: 12,
        sentSMS: 40,
        smsGrowth: 8,
        successfulSMS: 38,
        successRate: 95,
        failedSMS: 2,
        errorRate: 5,
        responseRate: 75,
        totalResponses: 30,
        averageRating: 4.2,
        totalRatings: 28,
        bestResponseTime: '2.5 دقیقه',
        bestHour: '14:00-16:00',
        bestProduct: 'لپ‌تاپ',
        bestProductRating: 4.5,
        activeRegion: 'تهران',
        regionOrders: 25,
        improvementRate: 15
      },
      week: {
        totalOrders: 320,
        ordersGrowth: 8,
        sentSMS: 300,
        smsGrowth: 12,
        successfulSMS: 285,
        successRate: 95,
        failedSMS: 15,
        errorRate: 5,
        responseRate: 70,
        totalResponses: 210,
        averageRating: 4.1,
        totalRatings: 195,
        bestResponseTime: '3.2 دقیقه',
        bestHour: '15:00-17:00',
        bestProduct: 'موبایل',
        bestProductRating: 4.3,
        activeRegion: 'اصفهان',
        regionOrders: 180,
        improvementRate: 12
      },
      month: {
        totalOrders: 1250,
        ordersGrowth: 15,
        sentSMS: 1200,
        smsGrowth: 18,
        successfulSMS: 1140,
        successRate: 95,
        failedSMS: 60,
        errorRate: 5,
        responseRate: 68,
        totalResponses: 816,
        averageRating: 4.0,
        totalRatings: 750,
        bestResponseTime: '4.1 دقیقه',
        bestHour: '16:00-18:00',
        bestProduct: 'کتاب',
        bestProductRating: 4.2,
        activeRegion: 'شیراز',
        regionOrders: 680,
        improvementRate: 10
      },
      year: {
        totalOrders: 15000,
        ordersGrowth: 25,
        sentSMS: 14500,
        smsGrowth: 30,
        successfulSMS: 13775,
        successRate: 95,
        failedSMS: 725,
        errorRate: 5,
        responseRate: 65,
        totalResponses: 9425,
        averageRating: 3.9,
        totalRatings: 8700,
        bestResponseTime: '5.2 دقیقه',
        bestHour: '17:00-19:00',
        bestProduct: 'پوشاک',
        bestProductRating: 4.0,
        activeRegion: 'مشهد',
        regionOrders: 8200,
        improvementRate: 8
      }
    }
    
    Object.assign(stats, mockData[selectedPeriod.value as keyof typeof mockData])
    
  } catch (error) {
    console.error('Error loading stats:', error)
  } finally {
    loading.value = false
  }
}

const refreshStats = () => {
  loadStats()
}

// Watchers
watch(selectedPeriod, () => {
  loadStats()
})

// Lifecycle
onMounted(() => {
  loadStats()
})
</script>

<style scoped>
.stat-card {
  border-radius: 0.5rem;
  padding: 1rem;
  color: white;
  transition: all 0.3s;
}

.stat-card:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  transform: scale(1.05);
}
</style>
