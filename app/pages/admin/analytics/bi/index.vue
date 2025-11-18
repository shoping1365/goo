<template>
  <div class="min-h-screen" dir="rtl">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-6">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-4">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">هوش تجاری و داده‌کاوی</h1>
            <p class="text-sm text-gray-600 mt-1">تحلیل جامع داده‌ها و گزارش‌های هوشمند کسب‌وکار</p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              @click="showSettings = true"
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              </svg>
              تنظیمات
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Dashboard Overview Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <!-- Total Revenue -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 mb-1">کل درآمد</p>
              <p class="text-2xl font-bold text-gray-900">{{ formatPrice(overview.totalRevenue) }}</p>
              <div class="flex items-center mt-2">
                <span class="text-green-600 text-sm font-medium">+{{ overview.revenueGrowth }}%</span>
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

        <!-- Total Orders -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 mb-1">کل سفارشات</p>
              <p class="text-2xl font-bold text-gray-900">{{ overview.totalOrders.toLocaleString() }}</p>
              <div class="flex items-center mt-2">
                <span class="text-blue-600 text-sm font-medium">+{{ overview.ordersGrowth }}%</span>
                <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
              </div>
            </div>
            <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
              </svg>
            </div>
          </div>
        </div>

        <!-- Conversion Rate -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 mb-1">نرخ تبدیل</p>
              <p class="text-2xl font-bold text-gray-900">{{ overview.conversionRate }}%</p>
              <div class="flex items-center mt-2">
                <span class="text-purple-600 text-sm font-medium">+{{ overview.conversionGrowth }}%</span>
                <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
              </div>
            </div>
            <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
              </svg>
            </div>
          </div>
        </div>

        <!-- Active Users -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 mb-1">کاربران فعال</p>
              <p class="text-2xl font-bold text-gray-900">{{ overview.activeUsers.toLocaleString() }}</p>
              <div class="flex items-center mt-2">
                <span class="text-orange-600 text-sm font-medium">+{{ overview.usersGrowth }}%</span>
                <span class="text-sm text-gray-500 mr-2">نسبت به ماه قبل</span>
              </div>
            </div>
            <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Access Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <NuxtLink to="/admin/analytics/bi/OverviewDashboard" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-blue-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-blue-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">داشبورد کلی</h3>
            <p class="text-sm text-gray-600">نمای کلی از تمام شاخص‌های کلیدی عملکرد</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/bi/ProductAnalysisDashboard" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-green-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-green-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">تحلیل محصولات</h3>
            <p class="text-sm text-gray-600">تحلیل عملکرد محصولات و موجودی</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/bi/SalesAnalysisDashboard" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-purple-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-purple-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">تحلیل فروش</h3>
            <p class="text-sm text-gray-600">تحلیل روند فروش و پیش‌بینی‌ها</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/bi/CustomerAnalysisDashboard" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-orange-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
              </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-orange-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">تحلیل مشتریان</h3>
            <p class="text-sm text-gray-600">تحلیل رفتار و رضایت مشتریان</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/Reports" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-red-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-red-500 to-red-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-red-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">گزارش‌ها</h3>
            <p class="text-sm text-gray-600">گزارش‌های تفصیلی و قابل دانلود</p>
          </div>
        </NuxtLink>

        <div @click="showSettings = true" class="group cursor-pointer">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-gray-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-gray-500 to-gray-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-gray-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">تنظیمات BI</h3>
            <p class="text-sm text-gray-600">تنظیمات و پیکربندی سیستم هوش تجاری</p>
          </div>
        </div>
      </div>

      <!-- Recent Activity -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">فعالیت‌های اخیر</h3>
          <button class="text-sm text-blue-600 hover:text-blue-700 font-medium">مشاهده همه</button>
        </div>
        <div class="space-y-4">
          <div v-for="activity in recentActivities" :key="activity.id" class="flex items-center p-3 bg-gray-50 rounded-lg">
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

    <!-- Settings Modal -->
    <div v-if="showSettings" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-md mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تنظیمات BI</h3>
          <button @click="showSettings = false" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        <p class="text-sm text-gray-600 mb-4">تنظیمات سیستم هوش تجاری در حال توسعه است.</p>
        <div class="flex justify-end">
          <button @click="showSettings = false" class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors">
            بستن
          </button>
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

const showSettings = ref(false)

// Sample data for overview
const overview = ref({
  totalRevenue: 1250000000,
  revenueGrowth: 12.5,
  totalOrders: 2847,
  ordersGrowth: 8.3,
  conversionRate: 3.2,
  conversionGrowth: -0.5,
  activeUsers: 1247,
  usersGrowth: 15.2
})

// Sample data for recent activities
const recentActivities = ref([
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
</script>

<style scoped>
/* Custom styles */
</style> 

