<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl">
          <!-- عنوان و تب‌های ناوبری -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <h1 class="text-sm font-bold text-gray-900 text-right mb-4">کارت های هدیه</h1>
            <div class="flex justify-end space-x-6 md:space-x-10 space-x-reverse pr-3 border-b border-gray-200">
            <button
              :class="[
                'pb-4 px-3 font-medium text-sm border-b-2',
                activeTab === 'received' ? 'border-red-500 text-red-600' : 'border-transparent text-gray-500 hover:text-gray-700'
              ]"
              @click="activeTab = 'received'"
            >
              هدیه گرفتم
            </button>
            <button 
              :class="[
                'pb-4 px-3 font-medium text-sm border-b-2',
                activeTab === 'given' ? 'border-red-500 text-red-600' : 'border-transparent text-gray-500 hover:text-gray-700'
              ]"
              @click="activeTab = 'given'"
            >
              هدیه دادم
            </button>
            </div>
          </div>

          <!-- محتوای خالی -->
          <div class="bg-white rounded-lg border border-gray-200 p-12 text-center">
            <!-- تصویر توستر و کارت‌های هدیه -->
            <div class="mb-8">
              <div class="relative inline-block">
                <!-- توستر نارنجی -->
                <div class="w-32 h-24 bg-orange-400 rounded-lg relative mx-auto mb-4">
                  <!-- شکاف‌های توستر -->
                  <div class="absolute top-2 left-4 w-6 h-16 bg-orange-600 rounded"></div>
                  <div class="absolute top-2 right-4 w-6 h-16 bg-orange-600 rounded"></div>
                  <!-- دسته توستر -->
                  <div class="absolute -right-2 top-8 w-4 h-8 bg-orange-600 rounded-r-lg"></div>
                </div>
                
                <!-- کارت‌های هدیه که از توستر بیرون می‌آیند -->
                <div class="absolute -top-8 -left-4">
                  <div class="w-8 h-6 bg-purple-400 rounded transform rotate-12"></div>
                </div>
                <div class="absolute -top-6 -right-2">
                  <div class="w-6 h-4 bg-yellow-400 rounded transform -rotate-12"></div>
                </div>
                <div class="absolute -top-6 left-8">
                  <div class="w-7 h-5 bg-pink-400 rounded transform rotate-6"></div>
                </div>
                
                <!-- ستاره‌های تزئینی -->
                <div class="absolute -top-2 -left-6">
                  <svg class="w-4 h-4 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.363-1.118l-2.8-2.034c-.784-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                </div>
                <div class="absolute -top-1 -right-4">
                  <svg class="w-3 h-3 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.363-1.118l-2.8-2.034c-.784-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                </div>
                <div class="absolute top-2 -left-8">
                  <svg class="w-2 h-2 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.363-1.118l-2.8-2.034c-.784-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                </div>
              </div>
            </div>

            <!-- متن توضیحات -->
            <div class="mb-8">
              <p class="text-gray-600 text-lg leading-relaxed">
                کسی به شما کارت هدیه داده؟ اینجا اضافه کنید تا بعدا دنبالش نگردید
              </p>
            </div>

            <!-- دکمه اضافه کردن کارت هدیه -->
            <button class="bg-pink-500 hover:bg-pink-600 text-white font-medium py-3 px-8 rounded-lg transition-colors">
              اضافه کردن کارت هدیه
            </button>
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
const activeTab = ref('received')

// تنظیم عنوان صفحه
useHead({
  title: 'کارت های هدیه - حساب کاربری'
})
</script>
