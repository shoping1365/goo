<template>
  <div class="space-y-6">
    <!-- Advanced Filters -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-gray-900">فیلترهای پیشرفته</h3>
        <button class="text-sm text-gray-500 hover:text-gray-700" @click="clearFilters">
          پاک کردن فیلترها
        </button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- Date Range -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select v-model="filters.dateRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="today">امروز</option>
            <option value="yesterday">دیروز</option>
            <option value="week">هفته اخیر</option>
            <option value="month">ماه اخیر</option>
            <option value="quarter">سه ماهه</option>
            <option value="year">سال</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>

        <!-- Product Category -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی کالا</label>
          <select v-model="filters.category" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه دسته‌ها</option>
            <option value="mobile">موبایل</option>
            <option value="laptop">لپ‌تاپ</option>
            <option value="tablet">تبلت</option>
            <option value="accessories">لوازم جانبی</option>
          </select>
        </div>

        <!-- Brand -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">برند</label>
          <select v-model="filters.brand" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه برندها</option>
            <option value="samsung">سامسونگ</option>
            <option value="apple">اپل</option>
            <option value="xiaomi">شیائومی</option>
            <option value="huawei">هوآوی</option>
          </select>
        </div>

        <!-- Sales Channel -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کانال فروش</label>
          <select v-model="filters.channel" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه کانال‌ها</option>
            <option value="website">وب‌سایت</option>
            <option value="mobile_app">اپ موبایل</option>
            <option value="instagram">اینستاگرام</option>
            <option value="telegram">تلگرام</option>
          </select>
        </div>
      </div>

      <!-- Custom Date Range -->
      <div v-if="filters.dateRange === 'custom'" class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input v-model="filters.startDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input v-model="filters.endDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
        </div>
      </div>

      <!-- Apply Filters Button -->
      <div class="mt-4 flex justify-end">
        <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="applyFilters">
          اعمال فیلترها
        </button>
      </div>
    </div>

    <!-- Sales Performance Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Sales vs Profit Trend -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">روند فروش و سود</h3>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button 
              v-for="metric in ['revenue', 'profit', 'margin']" 
              :key="metric"
              :class="[
                selectedMetric === metric 
                  ? 'bg-blue-100 text-blue-700' 
                  : 'text-gray-500 hover:text-gray-700',
                'px-3 py-1 rounded-md text-sm font-medium transition-colors'
              ]"
              @click="selectedMetric = metric"
            >
              {{ getMetricText(metric) }}
            </button>
          </div>
        </div>
        <div class="h-80">
          <!-- Chart placeholder -->
          <div class="w-full h-full bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path>
              </svg>
              <p class="text-gray-500">نمودار روند {{ getMetricText(selectedMetric) }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Sales Funnel -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">قیف فروش</h3>
        <div class="space-y-4">
          <div v-for="(stage, index) in salesFunnel" :key="stage.name" class="relative">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-medium text-gray-700">{{ stage.name }}</span>
              <span class="text-sm text-gray-500">{{ stage.count.toLocaleString() }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-3">
              <div 
                class="bg-gradient-to-r from-blue-500 to-blue-600 h-3 rounded-full transition-all duration-500"
                :style="{ width: `${stage.percentage}%` }"
              ></div>
            </div>
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>{{ stage.percentage }}%</span>
              <span v-if="index < salesFunnel.length - 1">
                ریزش: {{ ((salesFunnel[index].count - salesFunnel[index + 1].count) / salesFunnel[index].count * 100).toFixed(1) }}%
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Market Basket Analysis -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تحلیل سبد خرید</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-500">کالاهایی که معمولاً با هم خریداری می‌شوند</span>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="bundle in marketBaskets" :key="bundle.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-3">
            <h4 class="font-medium text-gray-900">{{ bundle.name }}</h4>
            <span class="text-sm bg-green-100 text-green-800 px-2 py-1 rounded-full">
              {{ bundle.confidence }}% اطمینان
            </span>
          </div>
          
          <div class="space-y-2 mb-4">
            <div v-for="product in bundle.products" :key="product.id" class="flex items-center">
              <div class="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center text-xs font-medium text-gray-600 ml-2">
                {{ product.id }}
              </div>
              <div class="flex-1">
                <p class="text-sm text-gray-900">{{ product.name }}</p>
                <p class="text-xs text-gray-500">{{ formatPrice(product.price) }}</p>
              </div>
            </div>
          </div>
          
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-500">{{ bundle.frequency }} بار خرید</span>
            <button class="text-blue-600 hover:text-blue-800 font-medium">
              ایجاد باندل
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Sales by Category and Brand -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Sales by Category -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">فروش بر اساس دسته‌بندی</h3>
        <div class="space-y-4">
          <div v-for="category in salesByCategory" :key="category.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full ml-3" :style="{ backgroundColor: category.color }"></div>
              <span class="text-sm font-medium text-gray-900">{{ category.name }}</span>
            </div>
            <div class="text-left">
              <p class="text-sm font-medium text-gray-900">{{ formatPrice(category.revenue) }}</p>
              <p class="text-xs text-gray-500">{{ category.percentage }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Sales by Brand -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">فروش بر اساس برند</h3>
        <div class="space-y-4">
          <div v-for="brand in salesByBrand" :key="brand.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full ml-3" :style="{ backgroundColor: brand.color }"></div>
              <span class="text-sm font-medium text-gray-900">{{ brand.name }}</span>
            </div>
            <div class="text-left">
              <p class="text-sm font-medium text-gray-900">{{ formatPrice(brand.revenue) }}</p>
              <p class="text-xs text-gray-500">{{ brand.percentage }}%</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Anomaly Detection -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تشخیص ناهنجاری</h3>
        <div class="flex items-center">
          <div class="w-2 h-2 bg-red-500 rounded-full animate-pulse ml-2"></div>
          <span class="text-sm text-red-600 font-medium">{{ anomalies.length }} هشدار</span>
        </div>
      </div>
      
      <div class="space-y-4">
        <div v-for="anomaly in anomalies" :key="anomaly.id" class="border border-red-200 bg-red-50 rounded-lg p-6">
          <div class="flex items-start">
            <div class="w-6 h-6 bg-red-500 rounded-full flex items-center justify-center ml-3 mt-0.5">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="flex-1">
              <h4 class="text-sm font-medium text-red-800">{{ anomaly.title }}</h4>
              <p class="text-sm text-red-700 mt-1">{{ anomaly.description }}</p>
              <div class="flex items-center mt-2">
                <span class="text-xs text-red-600">{{ formatTime(anomaly.timestamp) }}</span>
                <button class="text-xs text-red-600 hover:text-red-800 mr-4">مشاهده جزئیات</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

// Reactive data
const filters = ref({
  dateRange: 'month',
  category: '',
  brand: '',
  channel: '',
  startDate: '',
  endDate: ''
})

const selectedMetric = ref('revenue')

// Sample data
const salesFunnel = ref([
  { name: 'بازدید از صفحه محصول', count: 15420, percentage: 100 },
  { name: 'افزودن به سبد خرید', count: 4626, percentage: 30 },
  { name: 'رفتن به صفحه پرداخت', count: 2776, percentage: 18 },
  { name: 'پرداخت موفق', count: 1943, percentage: 12.6 }
])

const marketBaskets = ref([
  {
    id: 1,
    name: 'باندل موبایل',
    confidence: 85,
    frequency: 234,
    products: [
      { id: 1, name: 'گوشی سامسونگ گلکسی S24', price: 15000000 },
      { id: 2, name: 'قاب محافظ شیشه‌ای', price: 500000 },
      { id: 3, name: 'شارژر بی‌سیم', price: 800000 }
    ]
  },
  {
    id: 2,
    name: 'باندل لپ‌تاپ',
    confidence: 72,
    frequency: 156,
    products: [
      { id: 4, name: 'لپ‌تاپ اپل مک‌بوک پرو', price: 45000000 },
      { id: 5, name: 'ماوس بی‌سیم', price: 1200000 },
      { id: 6, name: 'کیف لپ‌تاپ', price: 800000 }
    ]
  },
  {
    id: 3,
    name: 'باندل صوتی',
    confidence: 68,
    frequency: 89,
    products: [
      { id: 7, name: 'هدفون سونی WH-1000XM5', price: 4500000 },
      { id: 8, name: 'کیف محافظ', price: 300000 }
    ]
  }
])

const salesByCategory = ref([
  { name: 'موبایل', revenue: 4500000000, percentage: 35, color: '#3B82F6' },
  { name: 'لپ‌تاپ', revenue: 3200000000, percentage: 25, color: '#10B981' },
  { name: 'لوازم جانبی', revenue: 2600000000, percentage: 20, color: '#F59E0B' },
  { name: 'تبلت', revenue: 1950000000, percentage: 15, color: '#EF4444' },
  { name: 'سایر', revenue: 650000000, percentage: 5, color: '#8B5CF6' }
])

const salesByBrand = ref([
  { name: 'سامسونگ', revenue: 3800000000, percentage: 30, color: '#3B82F6' },
  { name: 'اپل', revenue: 3500000000, percentage: 27, color: '#10B981' },
  { name: 'شیائومی', revenue: 2200000000, percentage: 17, color: '#F59E0B' },
  { name: 'هوآوی', revenue: 1800000000, percentage: 14, color: '#EF4444' },
  { name: 'سایر', revenue: 1800000000, percentage: 12, color: '#8B5CF6' }
])

const anomalies = ref([
  {
    id: 1,
    title: 'افت نرخ تبدیل',
    description: 'نرخ تبدیل امروز ۳۰٪ کمتر از میانگین هفته گذشته است',
    timestamp: new Date(Date.now() - 3600000)
  },
  {
    id: 2,
    title: 'افزایش غیرعادی بازگشت کالا',
    description: 'میزان بازگشت کالا در دسته موبایل ۲۵٪ افزایش یافته است',
    timestamp: new Date(Date.now() - 7200000)
  }
])

// Methods
const clearFilters = () => {
  filters.value = {
    dateRange: 'month',
    category: '',
    brand: '',
    channel: '',
    startDate: '',
    endDate: ''
  }
}

const applyFilters = () => {

  // Implementation for applying filters
}

const getMetricText = (metric: string): string => {
  const metrics = {
    revenue: 'فروش',
    profit: 'سود',
    margin: 'حاشیه سود'
  }
  return metrics[metric] || metric
}

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatTime = (date: Date): string => {
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const hours = Math.floor(diff / 3600000)
  
  if (hours < 1) return 'کمتر از یک ساعت پیش'
  if (hours < 24) return `${hours} ساعت پیش`
  return date.toLocaleDateString('fa-IR')
}
</script> 

