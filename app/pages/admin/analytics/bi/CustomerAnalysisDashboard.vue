<template>
  <div class="space-y-6">
    <!-- Customer Overview Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- Total Customers -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">کل مشتریان</p>
            <p class="text-2xl font-bold text-gray-900">{{ customerStats.totalCustomers.toLocaleString() }}</p>
            <div class="flex items-center mt-2">
              <span class="text-green-600 text-sm font-medium">+{{ customerStats.newCustomersToday }}</span>
              <span class="text-sm text-gray-500 mr-2">امروز</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Active Customers -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">مشتریان فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ customerStats.activeCustomers.toLocaleString() }}</p>
            <div class="flex items-center mt-2">
              <span class="text-green-600 text-sm font-medium">{{ customerStats.activePercentage }}%</span>
              <span class="text-sm text-gray-500 mr-2">از کل مشتریان</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Average CLV -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">میانگین CLV</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatPrice(customerStats.averageCLV) }}</p>
            <div class="flex items-center mt-2">
              <span class="text-green-600 text-sm font-medium">+{{ customerStats.clvGrowth }}%</span>
              <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Churn Rate -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">نرخ ریزش</p>
            <p class="text-2xl font-bold text-gray-900">{{ customerStats.churnRate }}%</p>
            <div class="flex items-center mt-2">
              <span :class="customerStats.churnChange <= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                {{ customerStats.churnChange <= 0 ? '-' : '+' }}{{ Math.abs(customerStats.churnChange) }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-red-500 to-red-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- RFM Customer Segmentation -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">بخش‌بندی مشتریان (RFM)</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button 
            v-for="segment in ['all', 'champions', 'loyal', 'at_risk', 'lost']" 
            :key="segment"
            :class="[
              selectedSegment === segment 
                ? 'bg-blue-100 text-blue-700' 
                : 'text-gray-500 hover:text-gray-700',
              'px-3 py-1 rounded-md text-sm font-medium transition-colors'
            ]"
            @click="selectedSegment = segment"
          >
            {{ getSegmentText(segment) }}
          </button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div v-for="segment in rfmSegments" :key="segment.name" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-3">
            <h4 class="font-medium text-gray-900">{{ segment.name }}</h4>
            <span :class="getSegmentBadgeClass(segment.type)" class="text-xs px-2 py-1 rounded-full">
              {{ segment.count.toLocaleString() }}
            </span>
          </div>
          
          <div class="space-y-2 mb-4">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">میانگین خرید:</span>
              <span class="font-medium">{{ segment.avgOrders }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">میانگین ارزش:</span>
              <span class="font-medium">{{ formatPrice(segment.avgValue) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">آخرین خرید:</span>
              <span class="font-medium">{{ segment.lastPurchase }} روز پیش</span>
            </div>
          </div>
          
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-500">{{ segment.percentage }}%</span>
            <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
              مشاهده لیست
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Customer Lifetime Value Analysis -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- CLV Distribution -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">توزیع ارزش طول عمر مشتری</h3>
        <div class="h-80">
          <!-- Chart placeholder -->
          <div class="w-full h-full bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
              <p class="text-gray-500">نمودار توزیع CLV</p>
            </div>
          </div>
        </div>
      </div>

      <!-- New vs Returning Customers -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">مشتریان جدید در مقابل بازگشتی</h3>
        <div class="space-y-6">
          <div v-for="(period, index) in customerAcquisition" :key="index" class="border-b border-gray-100 pb-4 last:border-b-0">
            <div class="flex items-center justify-between mb-3">
              <h4 class="font-medium text-gray-900">{{ period.name }}</h4>
              <span class="text-sm text-gray-500">{{ period.total.toLocaleString() }} کل</span>
            </div>
            
            <div class="grid grid-cols-2 gap-6">
              <div class="bg-blue-50 rounded-lg p-3">
                <div class="flex items-center justify-between mb-2">
                  <span class="text-sm font-medium text-blue-800">جدید</span>
                  <span class="text-sm text-blue-600">{{ period.new.toLocaleString() }}</span>
                </div>
                <div class="w-full bg-blue-200 rounded-full h-2">
                  <div 
                    class="bg-blue-500 h-2 rounded-full"
                    :style="{ width: `${(period.new / period.total) * 100}%` }"
                  ></div>
                </div>
                <p class="text-xs text-blue-600 mt-1">{{ ((period.new / period.total) * 100).toFixed(1) }}%</p>
              </div>
              
              <div class="bg-green-50 rounded-lg p-3">
                <div class="flex items-center justify-between mb-2">
                  <span class="text-sm font-medium text-green-800">بازگشتی</span>
                  <span class="text-sm text-green-600">{{ period.returning.toLocaleString() }}</span>
                </div>
                <div class="w-full bg-green-200 rounded-full h-2">
                  <div 
                    class="bg-green-500 h-2 rounded-full"
                    :style="{ width: `${(period.returning / period.total) * 100}%` }"
                  ></div>
                </div>
                <p class="text-xs text-green-600 mt-1">{{ ((period.returning / period.total) * 100).toFixed(1) }}%</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Customer Behavior Analysis -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-6">تحلیل رفتار مشتریان</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="behavior in customerBehaviors" :key="behavior.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <h4 class="font-medium text-gray-900">{{ behavior.title }}</h4>
            <span :class="getBehaviorBadgeClass(behavior.type)" class="text-xs px-2 py-1 rounded-full">
              {{ behavior.count.toLocaleString() }}
            </span>
          </div>
          
          <p class="text-sm text-gray-600 mb-3">{{ behavior.description }}</p>
          
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">میانگین سن:</span>
              <span class="font-medium">{{ behavior.avgAge }} سال</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">محبوب‌ترین شهر:</span>
              <span class="font-medium">{{ behavior.topCity }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">میانگین ارزش:</span>
              <span class="font-medium">{{ formatPrice(behavior.avgValue) }}</span>
            </div>
          </div>
          
          <div class="mt-4 pt-3 border-t border-gray-100">
            <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
              تحلیل عمیق
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Churn Prediction -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">پیش‌بینی ریزش مشتریان</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-500">احتمال ریزش بالا</span>
          <span class="text-sm bg-red-100 text-red-800 px-2 py-1 rounded-full font-medium">
            {{ churnPrediction.highRisk }} مشتری
          </span>
        </div>
      </div>
      
      <div class="space-y-4">
        <div v-for="customer in churnPrediction.customers" :key="customer.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-gray-300 rounded-full flex items-center justify-center ml-3">
                <span class="text-sm font-medium text-gray-600">{{ customer.name.charAt(0) }}</span>
              </div>
              <div>
                <h4 class="font-medium text-gray-900">{{ customer.name }}</h4>
                <p class="text-sm text-gray-500">{{ customer.email }}</p>
              </div>
            </div>
            
            <div class="text-left">
              <div class="flex items-center space-x-2 space-x-reverse">
                <span class="text-sm font-medium text-red-600">{{ customer.churnProbability }}%</span>
                <span class="text-xs text-gray-500">احتمال ریزش</span>
              </div>
              <p class="text-xs text-gray-500">{{ customer.lastPurchase }} روز پیش</p>
            </div>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                ارسال پیام
              </button>
              <button class="text-green-600 hover:text-green-800 text-sm font-medium">
                پیشنهاد ویژه
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '~/composables/useAuth';

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Reactive data
const selectedSegment = ref('all')

// Sample data
const customerStats = ref({
  totalCustomers: 125847,
  newCustomersToday: 234,
  activeCustomers: 89456,
  activePercentage: 71,
  averageCLV: 8500000,
  clvGrowth: 12.5,
  churnRate: 2.3,
  churnChange: -0.5
})

const rfmSegments = ref([
  {
    name: 'مشتریان وفادار',
    type: 'champions',
    count: 15678,
    percentage: 12.5,
    avgOrders: 8.5,
    avgValue: 12000000,
    lastPurchase: 15
  },
  {
    name: 'مشتریان وفادار',
    type: 'loyal',
    count: 23456,
    percentage: 18.6,
    avgOrders: 4.2,
    avgValue: 6500000,
    lastPurchase: 45
  },
  {
    name: 'در خطر ریزش',
    type: 'at_risk',
    count: 18934,
    percentage: 15.0,
    avgOrders: 2.1,
    avgValue: 3200000,
    lastPurchase: 120
  },
  {
    name: 'مشتریان از دست رفته',
    type: 'lost',
    count: 45678,
    percentage: 36.3,
    avgOrders: 1.8,
    avgValue: 2800000,
    lastPurchase: 365
  }
])

const customerAcquisition = ref([
  {
    name: 'امروز',
    total: 234,
    new: 189,
    returning: 45
  },
  {
    name: 'هفته اخیر',
    total: 1847,
    new: 1456,
    returning: 391
  },
  {
    name: 'ماه اخیر',
    total: 8234,
    new: 6234,
    returning: 2000
  }
])

const customerBehaviors = ref([
  {
    id: 1,
    title: 'خریداران فصلی',
    type: 'seasonal',
    count: 12345,
    description: 'مشتریانی که در فصل‌های خاص خرید می‌کنند',
    avgAge: 28,
    topCity: 'تهران',
    avgValue: 8500000
  },
  {
    id: 2,
    title: 'خریداران شبانه',
    type: 'night',
    count: 8765,
    description: 'مشتریانی که عمدتاً شب‌ها خرید می‌کنند',
    avgAge: 32,
    topCity: 'اصفهان',
    avgValue: 6200000
  },
  {
    id: 3,
    title: 'خریداران آخر هفته',
    type: 'weekend',
    count: 15678,
    description: 'مشتریانی که در آخر هفته خرید می‌کنند',
    avgAge: 35,
    topCity: 'مشهد',
    avgValue: 7800000
  }
])

const churnPrediction = ref({
  highRisk: 234,
  customers: [
    {
      id: 1,
      name: 'علی احمدی',
      email: 'ali@example.com',
      churnProbability: 85,
      lastPurchase: 45
    },
    {
      id: 2,
      name: 'فاطمه محمدی',
      email: 'fateme@example.com',
      churnProbability: 78,
      lastPurchase: 67
    },
    {
      id: 3,
      name: 'محمد رضایی',
      email: 'mohammad@example.com',
      churnProbability: 72,
      lastPurchase: 89
    }
  ]
})

// Methods
const getSegmentText = (segment: string): string => {
  const segments = {
    all: 'همه',
    champions: 'وفادار',
    loyal: 'عادی',
    at_risk: 'در خطر',
    lost: 'از دست رفته'
  }
  return segments[segment] || segment
}

const getSegmentBadgeClass = (type: string): string => {
  const classes = {
    champions: 'bg-green-100 text-green-800',
    loyal: 'bg-blue-100 text-blue-800',
    at_risk: 'bg-yellow-100 text-yellow-800',
    lost: 'bg-red-100 text-red-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getBehaviorBadgeClass = (type: string): string => {
  const classes = {
    seasonal: 'bg-purple-100 text-purple-800',
    night: 'bg-indigo-100 text-indigo-800',
    weekend: 'bg-pink-100 text-pink-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}
</script> 

