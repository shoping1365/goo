<template>
  <div class="space-y-6">
    <!-- KPI Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <!-- Total Revenue -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">درآمد کل</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatPrice(kpis.totalRevenue) }}</p>
            <div class="flex items-center mt-2">
              <span :class="kpis.revenueChange >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                {{ kpis.revenueChange >= 0 ? '+' : '' }}{{ kpis.revenueChange }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به دیروز</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Total Orders -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">تعداد سفارشات</p>
            <p class="text-2xl font-bold text-gray-900">{{ kpis.totalOrders.toLocaleString() }}</p>
            <div class="flex items-center mt-2">
              <span :class="kpis.ordersChange >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                {{ kpis.ordersChange >= 0 ? '+' : '' }}{{ kpis.ordersChange }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به دیروز</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Conversion Rate -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">نرخ تبدیل</p>
            <p class="text-2xl font-bold text-gray-900">{{ kpis.conversionRate }}%</p>
            <div class="flex items-center mt-2">
              <span :class="kpis.conversionChange >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                {{ kpis.conversionChange >= 0 ? '+' : '' }}{{ kpis.conversionChange }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به دیروز</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Average Order Value -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">میانگین سفارش</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatPrice(kpis.averageOrderValue) }}</p>
            <div class="flex items-center mt-2">
              <span :class="kpis.aovChange >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
                {{ kpis.aovChange >= 0 ? '+' : '' }}{{ kpis.aovChange }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به دیروز</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 13v-1a4 4 0 014-4 4 4 0 014 4v1m0 0h-3m3 0h3m-9 7a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Live Users -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-600 mb-1">کاربران آنلاین</p>
            <p class="text-2xl font-bold text-gray-900">{{ kpis.liveUsers.toLocaleString() }}</p>
            <div class="flex items-center mt-2">
              <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse ml-2"></div>
              <span class="text-sm text-gray-500">فعال</span>
            </div>
          </div>
          <div class="w-12 h-12 bg-gradient-to-br from-red-500 to-red-600 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Charts Row -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Sales Trend Chart -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">روند فروش ۳۰ روز گذشته</h3>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button 
              v-for="period in ['7d', '30d', '90d']" 
              :key="period"
              @click="salesPeriod = period"
              :class="[
                salesPeriod === period 
                  ? 'bg-blue-100 text-blue-700' 
                  : 'text-gray-500 hover:text-gray-700',
                'px-3 py-1 rounded-md text-sm font-medium transition-colors'
              ]"
            >
              {{ period === '7d' ? '۷ روز' : period === '30d' ? '۳۰ روز' : '۹۰ روز' }}
            </button>
          </div>
        </div>
        <div class="h-80">
          <!-- Chart placeholder - will be replaced with actual chart library -->
          <div class="w-full h-full bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path>
              </svg>
              <p class="text-gray-500">نمودار روند فروش</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Top Selling Products -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">پرفروش‌ترین کالاها</h3>
          <select class="text-sm border border-gray-300 rounded-md px-3 py-1">
            <option value="today">امروز</option>
            <option value="week">هفته اخیر</option>
            <option value="month">ماه اخیر</option>
          </select>
        </div>
        <div class="space-y-4">
          <div v-for="(product, index) in topProducts" :key="product.id" class="flex items-center">
            <div class="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center text-sm font-medium text-gray-600 ml-3">
              {{ index + 1 }}
            </div>
            <div class="flex-1">
              <p class="text-sm font-medium text-gray-900">{{ product.name }}</p>
              <p class="text-xs text-gray-500">{{ product.category }}</p>
            </div>
            <div class="text-left">
              <p class="text-sm font-medium text-gray-900">{{ product.sales.toLocaleString() }}</p>
              <p class="text-xs text-gray-500">{{ formatPrice(product.revenue) }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Geographic Sales Map -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">نقشه جغرافیایی فروش</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button class="px-3 py-1 bg-blue-100 text-blue-700 rounded-md text-sm font-medium">
            لحظه‌ای
          </button>
          <button class="px-3 py-1 text-gray-500 hover:text-gray-700 rounded-md text-sm font-medium">
            امروز
          </button>
          <button class="px-3 py-1 text-gray-500 hover:text-gray-700 rounded-md text-sm font-medium">
            هفته
          </button>
        </div>
      </div>
      <div class="h-96 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-20 h-20 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m-6 3l6-3"></path>
          </svg>
          <p class="text-gray-500">نقشه جغرافیایی فروش</p>
          <p class="text-sm text-gray-400 mt-1">نمایش پراکندگی سفارشات در سراسر کشور</p>
        </div>
      </div>
    </div>

    <!-- Real-time Activity Feed -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">فعالیت‌های لحظه‌ای</h3>
        <div class="flex items-center">
          <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse ml-2"></div>
          <span class="text-sm text-gray-500">زنده</span>
        </div>
      </div>
      <div class="space-y-3 max-h-64 overflow-y-auto">
        <div v-for="activity in realTimeActivities" :key="activity.id" class="flex items-center p-3 bg-gray-50 rounded-lg">
          <div class="w-2 h-2 bg-blue-500 rounded-full ml-3"></div>
          <div class="flex-1">
            <p class="text-sm text-gray-900">{{ activity.message }}</p>
            <p class="text-xs text-gray-500">{{ formatTime(activity.timestamp) }}</p>
          </div>
          <div class="text-left">
            <span :class="getActivityBadgeClass(activity.type)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
              {{ getActivityTypeText(activity.type) }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
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
const salesPeriod = ref('30d')

// Sample data
const kpis = ref({
  totalRevenue: 1250000000,
  revenueChange: 12.5,
  totalOrders: 2847,
  ordersChange: 8.3,
  conversionRate: 3.2,
  conversionChange: -0.5,
  averageOrderValue: 439000,
  aovChange: 4.1,
  liveUsers: 1247
})

const topProducts = ref([
  { id: 1, name: 'گوشی سامسونگ گلکسی S24', category: 'موبایل', sales: 156, revenue: 234000000 },
  { id: 2, name: 'لپ‌تاپ اپل مک‌بوک پرو', category: 'لپ‌تاپ', sales: 89, revenue: 445000000 },
  { id: 3, name: 'ساعت هوشمند اپل واچ', category: 'ساعت', sales: 234, revenue: 117000000 },
  { id: 4, name: 'هدفون سونی WH-1000XM5', category: 'صوتی', sales: 178, revenue: 89000000 },
  { id: 5, name: 'تبلت اپل آیپد پرو', category: 'تبلت', sales: 67, revenue: 134000000 }
])

const realTimeActivities = ref([
  { id: 1, message: 'سفارش جدید #ORD-2024-001 ثبت شد', type: 'order', timestamp: new Date() },
  { id: 2, message: 'کاربر جدید ثبت‌نام کرد', type: 'user', timestamp: new Date(Date.now() - 30000) },
  { id: 3, message: 'پرداخت موفق #TXN-123456', type: 'payment', timestamp: new Date(Date.now() - 60000) },
  { id: 4, message: 'محصول به سبد خرید اضافه شد', type: 'cart', timestamp: new Date(Date.now() - 90000) },
  { id: 5, message: 'بازدید از صفحه محصول', type: 'view', timestamp: new Date(Date.now() - 120000) }
])

// Methods
const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatTime = (date: Date): string => {
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const minutes = Math.floor(diff / 60000)
  
  if (minutes < 1) return 'همین الان'
  if (minutes < 60) return `${minutes} دقیقه پیش`
  return date.toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
}

const getActivityTypeText = (type: string): string => {
  const types = {
    order: 'سفارش',
    user: 'کاربر',
    payment: 'پرداخت',
    cart: 'سبد خرید',
    view: 'بازدید'
  }
  return types[type] || type
}

const getActivityBadgeClass = (type: string): string => {
  const classes = {
    order: 'bg-green-100 text-green-800',
    user: 'bg-blue-100 text-blue-800',
    payment: 'bg-purple-100 text-purple-800',
    cart: 'bg-yellow-100 text-yellow-800',
    view: 'bg-gray-100 text-gray-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

// Lifecycle
onMounted(() => {
  // Initialize real-time data updates
  console.log('Overview dashboard mounted')
})
</script> 

