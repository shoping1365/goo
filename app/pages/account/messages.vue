<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl" dir="rtl">
          <!-- عنوان و تب‌های ناوبری -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <div class="flex items-center justify-between mb-4">
              <h1 class="text-sm font-bold text-gray-900">پیام ها</h1>
              <button class="text-blue-600 hover:text-blue-700 text-sm font-medium transition-colors flex items-center space-x-2 space-x-reverse">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                <span>خواندن همه</span>
              </button>
            </div>
            <div class="flex space-x-8 space-x-reverse border-b border-gray-200">
            <button 
              class="pb-4 px-1 border-b-2 border-red-500 text-red-600 font-medium text-sm"
              @click="activeTab = 'all'"
            >
              همه پیام ها
            </button>
            <button 
              class="pb-4 px-1 border-b-2 border-transparent text-gray-500 hover:text-gray-700 font-medium text-sm"
              @click="activeTab = 'activities'"
            >
              فعالیت ها
            </button>
            <button 
              class="pb-4 px-1 border-b-2 border-transparent text-gray-500 hover:text-gray-700 font-medium text-sm"
              @click="activeTab = 'orders'"
            >
              سفارش ها
            </button>
            <button 
              class="pb-4 px-1 border-b-2 border-transparent text-gray-500 hover:text-gray-700 font-medium text-sm"
              @click="activeTab = 'discounts'"
            >
              تخفیف ها
            </button>
            </div>
          </div>

          <!-- محتوای خالی -->
          <div class="bg-white rounded-lg border border-gray-200 p-12 text-center">
            <!-- آیکون پیام -->
            <div class="mb-8">
              <div class="w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center mx-auto">
                <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
                </svg>
              </div>
            </div>

            <!-- متن توضیحات -->
            <div class="mb-8">
              <h3 class="text-lg font-medium text-gray-900 mb-2">پیامی وجود ندارد</h3>
              <p class="text-gray-600 leading-relaxed">
                هنوز هیچ پیامی دریافت نکرده‌اید. پیام‌های مربوط به سفارشات، فعالیت‌ها و تخفیف‌ها اینجا نمایش داده می‌شود.
              </p>
            </div>
          </div>
        </div>

        <!-- سایدبار -->
        <div class="lg:w-80 flex-shrink-0 lg:ml-8">
          <AccountSidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import AccountSidebar from '@/components/account/AccountSidebar.vue'

// تنظیم عنوان صفحه
definePageMeta({
  layout: 'default',
  middleware: ['auth']
})

// متغیر برای تب فعال
const activeTab = ref('all')

// تنظیم عنوان صفحه
useHead({
  title: 'پیام ها - حساب کاربری'
})
</script> 