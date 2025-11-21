<template>
  <div class="min-h-screen bg-gray-50" dir="rtl">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-6">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-4">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">هوش رقابتی CI</h1>
            <p class="text-sm text-gray-600 mt-1">تحلیل جامع بازار، رقبا و روندهای صنعت</p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
              @click="refreshData"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              به‌روزرسانی
            </button>
            <button 
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
              @click="showSettings = true"
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
      <!-- Overview Dashboard Section -->
      <OverviewDashboard />
      
      <!-- Quick Access Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <NuxtLink to="/admin/analytics/ci/competitor-profiles" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-blue-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-blue-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">پروفایل رقبا</h3>
            <p class="text-sm text-gray-600">تحلیل کامل رقبا و استراتژی‌های آن‌ها</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/ci/digital-footprint" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-green-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-green-500 to-green-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-green-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">ردپای دیجیتال</h3>
            <p class="text-sm text-gray-600">پایش فعالیت‌های آنلاین رقبا</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/ci/market-trends" class="group">
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
            <h3 class="text-lg font-semibold text-gray-900 mb-2">روندهای بازار</h3>
            <p class="text-sm text-gray-600">تحلیل روندها و تغییرات صنعت</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/ci/product-analysis" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-orange-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-orange-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-orange-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">تحلیل محصولات</h3>
            <p class="text-sm text-gray-600">مقایسه محصولات و خدمات رقبا</p>
          </div>
        </NuxtLink>

        <NuxtLink to="/admin/analytics/ci/marketing-campaigns" class="group">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-red-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-red-500 to-red-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-red-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">کمپین‌های بازاریابی</h3>
            <p class="text-sm text-gray-600">پایش استراتژی‌های بازاریابی رقبا</p>
          </div>
        </NuxtLink>

        <div class="group cursor-pointer" @click="showAllSections = true">
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all duration-200 group-hover:border-gray-300">
            <div class="flex items-center justify-between mb-4">
              <div class="w-12 h-12 bg-gradient-to-br from-gray-500 to-gray-600 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                </svg>
              </div>
              <svg class="w-5 h-5 text-gray-400 group-hover:text-gray-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">همه بخش‌ها</h3>
            <p class="text-sm text-gray-600">دسترسی به تمام ابزارهای تحلیل رقابتی</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Settings Modal -->
    <div v-if="showSettings" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-md mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تنظیمات CI</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showSettings = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        <p class="text-sm text-gray-600 mb-4">تنظیمات سیستم هوش رقابتی در حال توسعه است.</p>
        <div class="flex justify-end">
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="showSettings = false">
            بستن
          </button>
        </div>
      </div>
    </div>

    <!-- All Sections Modal -->
    <div v-if="showAllSections" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-4xl mx-4 max-h-[80vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">همه بخش‌های هوش رقابتی</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showAllSections = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div v-for="section in allSections" :key="section.id" class="p-6 border border-gray-200 rounded-lg hover:bg-gray-50">
            <h4 class="font-medium text-gray-900 mb-2">{{ section.title }}</h4>
            <p class="text-sm text-gray-600 mb-3">{{ section.description }}</p>
            <div class="flex items-center text-xs text-gray-500">
              <span class="w-2 h-2 rounded-full mr-2" :class="section.status === 'ready' ? 'bg-green-500' : 'bg-yellow-500'"></span>
              {{ section.status === 'ready' ? 'آماده' : 'در حال توسعه' }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import OverviewDashboard from './components/OverviewDashboard.vue'
import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

const showSettings = ref(false)
const showAllSections = ref(false)

// Sample data for all sections
const allSections = ref([
  { id: 1, title: 'نمای کلی', description: 'خلاصه‌ای از وضعیت بازار، رقبا و روندها', status: 'ready' },
  { id: 2, title: 'پروفایل رقبا', description: 'تحلیل کامل رقبا و استراتژی‌های آن‌ها', status: 'ready' },
  { id: 3, title: 'ردپای دیجیتال', description: 'پایش فعالیت‌های آنلاین رقبا', status: 'ready' },
  { id: 4, title: 'تحلیل محصولات', description: 'مقایسه محصولات و خدمات رقبا', status: 'ready' },
  { id: 5, title: 'کمپین‌های بازاریابی', description: 'پایش استراتژی‌های بازاریابی رقبا', status: 'ready' },
  { id: 6, title: 'روندهای بازار', description: 'تحلیل روندها و تغییرات صنعت', status: 'ready' },
  { id: 7, title: 'استراتژی فروش', description: 'ردیابی استراتژی فروش رقبا', status: 'ready' },
  { id: 8, title: 'تحلیل مالی', description: 'تحلیل داده‌های مالی رقبا', status: 'ready' },
  { id: 9, title: 'هشدارها و اعلان‌ها', description: 'اعلان تغییرات مهم و هشدارها', status: 'ready' },
  { id: 10, title: 'ابزارهای تحلیلی', description: 'امکانات تحلیلی و مقایسه‌ای', status: 'ready' },
  { id: 11, title: 'یکپارچگی داده', description: 'اتصال به منابع داده خارجی', status: 'ready' },
  { id: 12, title: 'مدیریت و اشتراک', description: 'سیستم مدیریت و اشتراک‌گذاری اطلاعات', status: 'ready' },
  { id: 13, title: 'بینش‌های هوش مصنوعی', description: 'قابلیت یادگیری و هوش مصنوعی', status: 'ready' },
  { id: 14, title: 'آرشیو و تاریخچه', description: 'ذخیره گزارش‌ها و داده‌های تاریخی', status: 'ready' }
])

const refreshData = () => {
  // TODO: Implement data refresh logic
  // Refreshing CI data
}
</script>

<style scoped>
/* Custom styles */
</style>



