<template>
  <div class="bg-white rounded-xl shadow-lg p-6" dir="rtl">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-xl font-bold text-gray-800">گزارشات تفصیلی</h3>
        <p class="text-gray-600 text-sm">تحلیل جامع عملکرد سیستم نظرسنجی</p>
      </div>
      <div class="flex space-x-3 space-x-reverse">
        <button class="px-4 py-2 bg-green-600 text-white rounded-lg text-sm hover:bg-green-700 flex items-center space-x-2 space-x-reverse" @click="exportReport">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          <span>خروجی PDF</span>
        </button>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700" @click="refreshReports">
          بروزرسانی
        </button>
      </div>
    </div>

    <!-- Report Tabs -->
    <div class="mb-6">
      <nav class="flex space-x-8 space-x-reverse border-b border-gray-200">
        <button 
          v-for="tab in tabs" 
          :key="tab.id"
          :class="[
            'py-2 px-1 border-b-2 font-medium text-sm transition-colors',
            activeTab === tab.id
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
          ]"
          @click="activeTab = tab.id"
        >
          {{ tab.name }}
        </button>
      </nav>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      <span class="mr-3 text-gray-600">در حال بارگذاری گزارشات...</span>
    </div>

    <!-- Tab Content -->
    <div v-else>
      <!-- Time-based Reports -->
      <div v-if="activeTab === 'time'" class="space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">گزارش روزانه</h4>
            <DailyReportChart :data="dailyData" />
          </div>
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">گزارش هفتگی</h4>
            <TrendChart :data="weeklyData" />
          </div>
        </div>
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="text-lg font-semibold mb-4 text-gray-800">گزارش ماهانه</h4>
          <TrendChart :data="monthlyData" />
        </div>
      </div>

      <!-- Gender-based Reports -->
      <div v-else-if="activeTab === 'gender'" class="space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">توزیع بر اساس جنسیت</h4>
            <ResponsePieChart :data="genderData" />
          </div>
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">مقایسه پاسخ‌دهی</h4>
            <GenderResponseChart :data="genderResponseData" />
          </div>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h4 class="text-lg font-semibold mb-4 text-gray-800">جدول تفصیلی جنسیت</h4>
          <GenderDetailTable :data="genderDetailData" />
        </div>
      </div>

      <!-- Product-based Reports -->
      <div v-else-if="activeTab === 'product'" class="space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">عملکرد محصولات</h4>
            <TrendChart :data="productData" />
          </div>
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">امتیازات محصولات</h4>
            <ResponsePieChart :data="productRatingData" />
          </div>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h4 class="text-lg font-semibold mb-4 text-gray-800">رتبه‌بندی محصولات</h4>
          <ProductRankingTable :data="productRankingData" />
        </div>
      </div>

      <!-- Error Reports -->
      <div v-else-if="activeTab === 'error'" class="space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">نوع خطاها</h4>
            <ErrorTypeChart :data="errorTypeData" />
          </div>
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">روند خطاها</h4>
            <ErrorTrendChart :data="errorTrendData" />
          </div>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h4 class="text-lg font-semibold mb-4 text-gray-800">لیست خطاهای ارسال</h4>
          <ErrorDetailTable :data="errorDetailData" />
        </div>
      </div>

      <!-- Regional Reports -->
      <div v-else-if="activeTab === 'regional'" class="space-y-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">توزیع جغرافیایی</h4>
            <ResponsePieChart :data="regionalData" />
          </div>
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-lg font-semibold mb-4 text-gray-800">عملکرد استان‌ها</h4>
            <TrendChart :data="provinceData" />
          </div>
        </div>
        <div class="bg-white rounded-lg border border-gray-200 p-6">
          <h4 class="text-lg font-semibold mb-4 text-gray-800">آمار تفصیلی مناطق</h4>
          <RegionalDetailTable :data="regionalDetailData" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import DailyReportChart from './DailyReportChart.vue'
import ErrorDetailTable from './ErrorDetailTable.vue'
import ErrorTrendChart from './ErrorTrendChart.vue'
import ErrorTypeChart from './ErrorTypeChart.vue'
import GenderDetailTable from './GenderDetailTable.vue'
import GenderResponseChart from './GenderResponseChart.vue'
import ProductRankingTable from './ProductRankingTable.vue'
import ResponsePieChart from './ResponsePieChart.vue'
import TrendChart from './TrendChart.vue'

// Reactive data
const activeTab = ref('time')
const loading = ref(false)

const tabs = [
  { id: 'time', name: 'گزارش زمانی' },
  { id: 'gender', name: 'گزارش جنسیت' },
  { id: 'product', name: 'گزارش محصولات' },
  { id: 'error', name: 'گزارش خطاها' },
  { id: 'regional', name: 'گزارش منطقه‌ای' }
]

// Time-based data
const dailyData = ref({
  labels: ['شنبه', 'یکشنبه', 'دوشنبه', 'سه‌شنبه', 'چهارشنبه', 'پنج‌شنبه', 'جمعه'],
  datasets: [
    {
      label: 'SMS ارسال شده',
      data: [45, 52, 48, 60, 55, 42, 38],
      borderColor: '#6366f1',
      backgroundColor: 'rgba(99,102,241,0.1)'
    },
    {
      label: 'پاسخ‌های دریافتی',
      data: [32, 38, 35, 45, 40, 30, 25],
      borderColor: '#10b981',
      backgroundColor: 'rgba(16,185,129,0.1)'
    }
  ]
})

const weeklyData = ref({
  labels: ['هفته 1', 'هفته 2', 'هفته 3', 'هفته 4'],
  datasets: [
    {
      label: 'کل سفارشات',
      data: [280, 320, 350, 380],
      borderColor: '#8b5cf6',
      backgroundColor: 'rgba(139,92,246,0.1)'
    }
  ]
})

const monthlyData = ref({
  labels: ['فروردین', 'اردیبهشت', 'خرداد', 'تیر', 'مرداد', 'شهریور'],
  datasets: [
    {
      label: 'SMS ارسال شده',
      data: [1200, 1350, 1500, 1650, 1800, 1950],
      borderColor: '#f59e0b',
      backgroundColor: 'rgba(245,158,11,0.1)'
    }
  ]
})

// Gender-based data
const genderData = ref({
  labels: ['مرد', 'زن'],
  datasets: [{
    data: [65, 35],
    backgroundColor: ['#3b82f6', '#ec4899']
  }]
})

const genderResponseData = ref({
  labels: ['مرد', 'زن'],
  datasets: [
    {
      label: 'نرخ پاسخ‌دهی',
      data: [78, 82],
      backgroundColor: ['#10b981', '#f59e0b']
    }
  ]
})

const genderDetailData = ref([
  { gender: 'مرد', totalOrders: 1250, sentSMS: 1250, responses: 975, responseRate: '78%', averageRating: 4.2 },
  { gender: 'زن', totalOrders: 850, sentSMS: 850, responses: 697, responseRate: '82%', averageRating: 4.5 }
])

// Product-based data
const productData = ref({
  labels: ['لپ‌تاپ', 'گوشی', 'تبلت', 'ساعت', 'هدفون'],
  datasets: [
    {
      label: 'تعداد نظرسنجی',
      data: [320, 280, 180, 120, 90],
      borderColor: '#3b82f6',
      backgroundColor: 'rgba(59,130,246,0.2)'
    }
  ]
})

const productRatingData = ref({
  labels: ['عالی', 'خوب', 'متوسط', 'ضعیف'],
  datasets: [{
    data: [45, 35, 15, 5],
    backgroundColor: ['#10b981', '#3b82f6', '#f59e0b', '#ef4444']
  }]
})

const productRankingData = ref([
  { rank: 1, product: 'لپ‌تاپ اپل مک‌بوک پرو', averageRating: 4.8, totalOrders: 156, responseRate: '95%' },
  { rank: 2, product: 'گوشی سامسونگ گلکسی S24', averageRating: 4.6, totalOrders: 142, responseRate: '92%' },
  { rank: 3, product: 'تبلت اپل iPad Pro', averageRating: 4.5, totalOrders: 98, responseRate: '89%' },
  { rank: 4, product: 'ساعت اپل Apple Watch', averageRating: 4.3, totalOrders: 87, responseRate: '85%' },
  { rank: 5, product: 'هدفون Sony WH-1000XM4', averageRating: 4.2, totalOrders: 73, responseRate: '82%' }
])

// Error-based data
const errorTypeData = ref({
  labels: ['شماره نامعتبر', 'خطای شبکه', 'محدودیت ارسال', 'خطای سرور', 'سایر'],
  datasets: [{
    data: [35, 25, 20, 15, 5],
    backgroundColor: ['#ef4444', '#f59e0b', '#3b82f6', '#8b5cf6', '#6b7280']
  }]
})

const errorTrendData = ref({
  labels: ['هفته 1', 'هفته 2', 'هفته 3', 'هفته 4'],
  datasets: [
    {
      label: 'خطاهای ارسال',
      data: [12, 8, 15, 6],
      borderColor: '#ef4444',
      backgroundColor: 'rgba(239,68,68,0.1)'
    }
  ]
})

const errorDetailData = ref([
  { orderNumber: 1, phoneNumber: '09123456789', customerName: 'کاربر ناشناس', errorType: 'شماره نامعتبر', errorDate: '1403/01/15', status: 'حل نشده', retryCount: 0 },
  { orderNumber: 2, phoneNumber: '09187654321', customerName: 'کاربر ناشناس', errorType: 'خطای شبکه', errorDate: '1403/01/14', status: 'حل شده', retryCount: 1 },
  { orderNumber: 3, phoneNumber: '09111111111', customerName: 'کاربر ناشناس', errorType: 'محدودیت ارسال', errorDate: '1403/01/13', status: 'در حال بررسی', retryCount: 2 }
])

// Regional data
const regionalData = ref({
  labels: ['تهران', 'اصفهان', 'مشهد', 'شیراز', 'تبریز', 'سایر'],
  datasets: [{
    data: [40, 20, 15, 10, 8, 7],
    backgroundColor: ['#3b82f6', '#10b981', '#f59e0b', '#ef4444', '#8b5cf6', '#6b7280']
  }]
})

const provinceData = ref({
  labels: ['تهران', 'اصفهان', 'مشهد', 'شیراز', 'تبریز'],
  datasets: [
    {
      label: 'نرخ پاسخ‌دهی',
      data: [85, 78, 82, 75, 80],
      borderColor: '#3b82f6',
      backgroundColor: 'rgba(59,130,246,0.2)'
    }
  ]
})

const regionalDetailData = ref([
  { province: 'تهران', total: 850, responded: 723, rate: '85%', avgRating: 4.3 },
  { province: 'اصفهان', total: 420, responded: 328, rate: '78%', avgRating: 4.1 },
  { province: 'مشهد', total: 380, responded: 312, rate: '82%', avgRating: 4.4 },
  { province: 'شیراز', total: 280, responded: 210, rate: '75%', avgRating: 4.0 },
  { province: 'تبریز', total: 220, responded: 176, rate: '80%', avgRating: 4.2 }
])

// Methods
const exportReport = () => {
  // console.log('Exporting report...')
  // Implementation for PDF export
}

const refreshReports = async () => {
  loading.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 2000))
    // console.log('Reports refreshed')
  } catch {
    // console.error('Error refreshing reports:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  refreshReports()
})
</script>

<style scoped>
/* Additional styles if needed */
</style> 
