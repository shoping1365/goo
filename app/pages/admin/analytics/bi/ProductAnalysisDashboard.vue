<template>
  <div class="space-y-6">
    <!-- Inventory Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- Total Products -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">کل کالاها</p>
            <p class="text-2xl font-bold text-gray-900">{{ inventoryStats.totalProducts.toLocaleString() }}</p>
            <div class="flex items-center mt-2">
              <span class="text-green-600 text-sm font-medium">+{{ inventoryStats.newProductsToday }}</span>
              <span class="text-sm text-gray-500 mr-2">امروز</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Low Stock Alert -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">کالاهای کم‌موجود</p>
            <p class="text-2xl font-bold text-red-600">{{ inventoryStats.lowStockCount.toLocaleString() }}</p>
            <div class="flex items-center mt-2">
              <span class="text-red-600 text-sm font-medium">{{ inventoryStats.lowStockPercentage }}%</span>
              <span class="text-sm text-gray-500 mr-2">از کل کالاها</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-red-500 to-red-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Out of Stock -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">کالاهای تمام شده</p>
            <p class="text-2xl font-bold text-orange-600">{{ inventoryStats.outOfStockCount.toLocaleString() }}</p>
            <div class="flex items-center mt-2">
              <span class="text-orange-600 text-sm font-medium">{{ inventoryStats.outOfStockPercentage }}%</span>
              <span class="text-sm text-gray-500 mr-2">از کل کالاها</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Inventory Value -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">ارزش کل موجودی</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatPrice(inventoryStats.totalValue) }}</p>
            <div class="flex items-center mt-2">
              <span class="text-green-600 text-sm font-medium">+{{ inventoryStats.valueGrowth }}%</span>
              <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Fast/Slow Moving Items -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Fast Moving Items -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">کالاهای پرگردش</h3>
          <span class="text-sm bg-green-100 text-green-800 px-2 py-1 rounded-full">
            {{ fastMovingItems.length }} کالا
          </span>
        </div>
        
        <div class="space-y-4">
          <div v-for="item in fastMovingItems" :key="item.id" class="border border-gray-100 rounded-lg p-3 hover:shadow-sm transition-shadow">
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center ml-3">
                  <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                  </svg>
                </div>
                <div>
                  <p class="text-sm font-medium text-gray-900">{{ item.name }}</p>
                  <p class="text-xs text-gray-500">{{ item.category }}</p>
                </div>
              </div>
              
              <div class="text-left">
                <p class="text-sm font-medium text-gray-900">{{ item.salesCount.toLocaleString() }}</p>
                <p class="text-xs text-gray-500">فروش ماهانه</p>
              </div>
              
              <div class="text-left">
                <p class="text-sm font-medium text-green-600">{{ item.turnoverRate }}%</p>
                <p class="text-xs text-gray-500">نرخ گردش</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Slow Moving Items -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">کالاهای کم‌گردش</h3>
          <span class="text-sm bg-red-100 text-red-800 px-2 py-1 rounded-full">
            {{ slowMovingItems.length }} کالا
          </span>
        </div>
        
        <div class="space-y-4">
          <div v-for="item in slowMovingItems" :key="item.id" class="border border-gray-100 rounded-lg p-3 hover:shadow-sm transition-shadow">
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center ml-3">
                  <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                  </svg>
                </div>
                <div>
                  <p class="text-sm font-medium text-gray-900">{{ item.name }}</p>
                  <p class="text-xs text-gray-500">{{ item.category }}</p>
                </div>
              </div>
              
              <div class="text-left">
                <p class="text-sm font-medium text-gray-900">{{ item.salesCount.toLocaleString() }}</p>
                <p class="text-xs text-gray-500">فروش ماهانه</p>
              </div>
              
              <div class="text-left">
                <p class="text-sm font-medium text-red-600">{{ item.turnoverRate }}%</p>
                <p class="text-xs text-gray-500">نرخ گردش</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- High Views Low Sales -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">کالاهای پربازدید ولی کم‌فروش</h3>
        <span class="text-sm bg-yellow-100 text-yellow-800 px-2 py-1 rounded-full">
          {{ highViewsLowSales.length }} کالا
        </span>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="item in highViewsLowSales" :key="item.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-3">
            <h4 class="font-medium text-gray-900">{{ item.name }}</h4>
            <span class="text-xs bg-yellow-100 text-yellow-800 px-2 py-1 rounded-full">
              {{ item.conversionRate }}% تبدیل
            </span>
          </div>
          
          <div class="space-y-2 mb-4">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">بازدید:</span>
              <span class="font-medium">{{ item.views.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">فروش:</span>
              <span class="font-medium">{{ item.sales.toLocaleString() }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">قیمت:</span>
              <span class="font-medium">{{ formatPrice(item.price) }}</span>
            </div>
          </div>
          
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">امتیاز:</span>
              <div class="flex items-center">
                <span class="text-yellow-400">★</span>
                <span class="font-medium mr-1">{{ item.rating }}</span>
              </div>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">نظرات:</span>
              <span class="font-medium">{{ item.reviews.toLocaleString() }}</span>
            </div>
          </div>
          
          <div class="mt-4 pt-3 border-t border-gray-100">
            <div class="flex space-x-2 space-x-reverse">
              <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                بررسی قیمت
              </button>
              <button class="text-green-600 hover:text-green-800 text-sm font-medium">
                بهبود اطلاعات
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Demand Forecasting -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Demand Forecast Chart -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">پیش‌بینی تقاضا</h3>
          <select class="text-sm border border-gray-300 rounded-md px-3 py-1">
            <option value="30">۳۰ روز آینده</option>
            <option value="60">۶۰ روز آینده</option>
            <option value="90">۹۰ روز آینده</option>
          </select>
        </div>
        
        <div class="h-80">
          <!-- Chart placeholder -->
          <div class="w-full h-full bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
              <p class="text-gray-500">نمودار پیش‌بینی تقاضا</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Stock Recommendations -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">توصیه‌های موجودی</h3>
        
        <div class="space-y-4">
          <div v-for="recommendation in stockRecommendations" :key="recommendation.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-3">
              <h4 class="font-medium text-gray-900">{{ recommendation.product }}</h4>
              <span :class="getRecommendationBadgeClass(recommendation.type)" class="text-xs px-2 py-1 rounded-full">
                {{ getRecommendationText(recommendation.type) }}
              </span>
            </div>
            
            <p class="text-sm text-gray-600 mb-3">{{ recommendation.description }}</p>
            
            <div class="grid grid-cols-2 gap-6 mb-3">
              <div class="text-sm">
                <span class="text-gray-500">موجودی فعلی:</span>
                <span class="font-medium mr-2">{{ recommendation.currentStock }}</span>
              </div>
              <div class="text-sm">
                <span class="text-gray-500">پیشنهاد:</span>
                <span class="font-medium mr-2">{{ recommendation.suggestedStock }}</span>
              </div>
            </div>
            
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-500">{{ recommendation.confidence }}% اطمینان</span>
              <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">
                اعمال
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Customer Reviews Analysis -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تحلیل نظرات مشتریان</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-500">پردازش شده با NLP</span>
          <span class="text-sm bg-blue-100 text-blue-800 px-2 py-1 rounded-full">
            {{ reviewsAnalysis.totalReviews.toLocaleString() }} نظر
          </span>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <div v-for="sentiment in reviewsAnalysis.sentiments" :key="sentiment.type" class="text-center">
          <div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-3" :class="getSentimentBgClass(sentiment.type)">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getSentimentIcon(sentiment.type)"></path>
            </svg>
          </div>
          <h4 class="font-medium text-gray-900 mb-1">{{ getSentimentText(sentiment.type) }}</h4>
          <p class="text-2xl font-bold" :class="getSentimentTextClass(sentiment.type)">{{ sentiment.percentage }}%</p>
          <p class="text-sm text-gray-500">{{ sentiment.count.toLocaleString() }} نظر</p>
        </div>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Top Positive Keywords -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">کلمات کلیدی مثبت</h4>
          <div class="flex flex-wrap gap-2">
            <span
v-for="keyword in reviewsAnalysis.positiveKeywords" :key="keyword.word" 
                  class="bg-green-100 text-green-800 px-3 py-1 rounded-full text-sm">
              {{ keyword.word }} ({{ keyword.count }})
            </span>
          </div>
        </div>
        
        <!-- Top Negative Keywords -->
        <div class="border border-gray-200 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">کلمات کلیدی منفی</h4>
          <div class="flex flex-wrap gap-2">
            <span
v-for="keyword in reviewsAnalysis.negativeKeywords" :key="keyword.word" 
                  class="bg-red-100 text-red-800 px-3 py-1 rounded-full text-sm">
              {{ keyword.word }} ({{ keyword.count }})
            </span>
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

// Sample data
const inventoryStats = ref({
  totalProducts: 1000000,
  newProductsToday: 45,
  lowStockCount: 1234,
  lowStockPercentage: 0.12,
  outOfStockCount: 567,
  outOfStockPercentage: 0.06,
  totalValue: 4500000000000,
  valueGrowth: 8.5
})

const fastMovingItems = ref([
  { id: 1, name: 'گوشی سامسونگ گلکسی S24', category: 'موبایل', salesCount: 234, turnoverRate: 85 },
  { id: 2, name: 'لپ‌تاپ اپل مک‌بوک پرو', category: 'لپ‌تاپ', salesCount: 156, turnoverRate: 72 },
  { id: 3, name: 'هدفون سونی WH-1000XM5', category: 'صوتی', salesCount: 189, turnoverRate: 68 }
])

const slowMovingItems = ref([
  { id: 4, name: 'کتاب الکترونیک قدیمی', category: 'کتاب', salesCount: 5, turnoverRate: 12 },
  { id: 5, name: 'لوازم جانبی قدیمی', category: 'لوازم جانبی', salesCount: 8, turnoverRate: 15 },
  { id: 6, name: 'نرم‌افزار قدیمی', category: 'نرم‌افزار', salesCount: 3, turnoverRate: 8 }
])

const highViewsLowSales = ref([
  {
    id: 1,
    name: 'گوشی آیفون 15 پرو مکس',
    views: 15420,
    sales: 234,
    conversionRate: 1.5,
    price: 85000000,
    rating: 4.8,
    reviews: 456
  },
  {
    id: 2,
    name: 'لپ‌تاپ گیمینگ ASUS ROG',
    views: 12340,
    sales: 156,
    conversionRate: 1.3,
    price: 65000000,
    rating: 4.6,
    reviews: 234
  },
  {
    id: 3,
    name: 'دوربین DSLR Canon EOS R5',
    views: 9876,
    sales: 89,
    conversionRate: 0.9,
    price: 95000000,
    rating: 4.9,
    reviews: 123
  }
])

const stockRecommendations = ref([
  {
    id: 1,
    product: 'گوشی سامسونگ گلکسی S24',
    type: 'restock',
    description: 'موجودی کم است، پیشنهاد می‌شود موجودی افزایش یابد',
    currentStock: 45,
    suggestedStock: 150,
    confidence: 92
  },
  {
    id: 2,
    product: 'لپ‌تاپ اپل مک‌بوک پرو',
    type: 'reduce',
    description: 'موجودی زیاد است، پیشنهاد می‌شود موجودی کاهش یابد',
    currentStock: 234,
    suggestedStock: 100,
    confidence: 78
  },
  {
    id: 3,
    product: 'هدفون سونی WH-1000XM5',
    type: 'maintain',
    description: 'موجودی در سطح مناسب است',
    currentStock: 67,
    suggestedStock: 70,
    confidence: 85
  }
])

const reviewsAnalysis = ref({
  totalReviews: 45678,
  sentiments: [
    { type: 'positive', percentage: 68, count: 31061 },
    { type: 'neutral', percentage: 22, count: 10049 },
    { type: 'negative', percentage: 10, count: 4568 }
  ],
  positiveKeywords: [
    { word: 'کیفیت عالی', count: 1234 },
    { word: 'قیمت مناسب', count: 987 },
    { word: 'تحویل سریع', count: 756 },
    { word: 'خدمات خوب', count: 654 }
  ],
  negativeKeywords: [
    { word: 'قیمت بالا', count: 234 },
    { word: 'تحویل کند', count: 189 },
    { word: 'کیفیت ضعیف', count: 156 },
    { word: 'خدمات بد', count: 123 }
  ]
})

// Methods
const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const getRecommendationText = (type: string): string => {
  const types = {
    restock: 'افزایش موجودی',
    reduce: 'کاهش موجودی',
    maintain: 'حفظ موجودی'
  }
  return types[type] || type
}

const getRecommendationBadgeClass = (type: string): string => {
  const classes = {
    restock: 'bg-red-100 text-red-800',
    reduce: 'bg-yellow-100 text-yellow-800',
    maintain: 'bg-green-100 text-green-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getSentimentText = (type: string): string => {
  const sentiments = {
    positive: 'مثبت',
    neutral: 'خنثی',
    negative: 'منفی'
  }
  return sentiments[type] || type
}

const getSentimentBgClass = (type: string): string => {
  const classes = {
    positive: 'bg-green-500',
    neutral: 'bg-gray-500',
    negative: 'bg-red-500'
  }
  return classes[type] || 'bg-gray-500'
}

const getSentimentTextClass = (type: string): string => {
  const classes = {
    positive: 'text-green-600',
    neutral: 'text-gray-600',
    negative: 'text-red-600'
  }
  return classes[type] || 'text-gray-600'
}

const getSentimentIcon = (type: string): string => {
  const icons = {
    positive: 'M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    neutral: 'M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z',
    negative: 'M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[type] || icons.neutral
}
</script> 
