<template>
  <div class="bg-white rounded-2xl shadow-xl border border-gray-100 p-8 mb-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-indigo-400 to-indigo-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">امتیازدهی جزئی‌تر</h2>
        <p class="text-gray-600 mt-1">نظر خود را درباره جنبه‌های مختلف تجربه خرید ثبت کنید</p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- ستون اول -->
      <div class="space-y-6">
        <!-- کیفیت محصولات -->
        <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-100">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-green-500 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-gray-900">کیفیت محصولات</h3>
            </div>
            <span class="text-sm text-green-600 font-medium">{{ getRatingText(ratings.productQuality) }}</span>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse mb-3">
            <button
              v-for="i in 5"
              :key="i"
              type="button"
              class="text-2xl transition-all duration-200 hover:scale-110"
              :class="i <= ratings.productQuality ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
              @click="ratings.productQuality = i"
            >★</button>
          </div>
          
          <p class="text-sm text-gray-600">کیفیت و دوام محصولات خریداری شده</p>
        </div>

        <!-- قیمت‌گذاری -->
        <div class="bg-gradient-to-r from-blue-50 to-cyan-50 rounded-xl p-6 border border-blue-100">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-blue-500 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-gray-900">قیمت‌گذاری</h3>
            </div>
            <span class="text-sm text-blue-600 font-medium">{{ getRatingText(ratings.pricing) }}</span>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse mb-3">
            <button
              v-for="i in 5"
              :key="i"
              type="button"
              class="text-2xl transition-all duration-200 hover:scale-110"
              :class="i <= ratings.pricing ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
              @click="ratings.pricing = i"
            >★</button>
          </div>
          
          <p class="text-sm text-gray-600">مناسب بودن قیمت نسبت به کیفیت محصول</p>
        </div>

        <!-- سرعت تحویل -->
        <div class="bg-gradient-to-r from-purple-50 to-violet-50 rounded-xl p-6 border border-purple-100">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-purple-500 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-gray-900">سرعت تحویل</h3>
            </div>
            <span class="text-sm text-purple-600 font-medium">{{ getRatingText(ratings.deliverySpeed) }}</span>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse mb-3">
            <button
              v-for="i in 5"
              :key="i"
              type="button"
              class="text-2xl transition-all duration-200 hover:scale-110"
              :class="i <= ratings.deliverySpeed ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
              @click="ratings.deliverySpeed = i"
            >★</button>
          </div>
          
          <p class="text-sm text-gray-600">سرعت و به موقع بودن تحویل سفارش</p>
        </div>
      </div>

      <!-- ستون دوم -->
      <div class="space-y-6">
        <!-- کیفیت بسته‌بندی -->
        <div class="bg-gradient-to-r from-orange-50 to-amber-50 rounded-xl p-6 border border-orange-100">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-orange-500 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-gray-900">کیفیت بسته‌بندی</h3>
            </div>
            <span class="text-sm text-orange-600 font-medium">{{ getRatingText(ratings.packaging) }}</span>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse mb-3">
            <button
              v-for="i in 5"
              :key="i"
              type="button"
              class="text-2xl transition-all duration-200 hover:scale-110"
              :class="i <= ratings.packaging ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
              @click="ratings.packaging = i"
            >★</button>
          </div>
          
          <p class="text-sm text-gray-600">کیفیت و امنیت بسته‌بندی محصولات</p>
        </div>

        <!-- خدمات پس از فروش -->
        <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-xl p-6 border border-red-100">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-red-500 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192L5.636 18.364M12 2.25a9.75 9.75 0 109.75 9.75A9.75 9.75 0 0012 2.25z"></path>
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-gray-900">خدمات پس از فروش</h3>
            </div>
            <span class="text-sm text-red-600 font-medium">{{ getRatingText(ratings.afterSales) }}</span>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse mb-3">
            <button
              v-for="i in 5"
              :key="i"
              type="button"
              class="text-2xl transition-all duration-200 hover:scale-110"
              :class="i <= ratings.afterSales ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
              @click="ratings.afterSales = i"
            >★</button>
          </div>
          
          <p class="text-sm text-gray-600">کیفیت خدمات پشتیبانی و گارانتی</p>
        </div>

        <!-- رابط کاربری سایت -->
        <div class="bg-gradient-to-r from-teal-50 to-cyan-50 rounded-xl p-6 border border-teal-100">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-teal-500 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
              </div>
              <h3 class="text-lg font-semibold text-gray-900">رابط کاربری سایت</h3>
            </div>
            <span class="text-sm text-teal-600 font-medium">{{ getRatingText(ratings.userInterface) }}</span>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse mb-3">
            <button
              v-for="i in 5"
              :key="i"
              type="button"
              class="text-2xl transition-all duration-200 hover:scale-110"
              :class="i <= ratings.userInterface ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
              @click="ratings.userInterface = i"
            >★</button>
          </div>
          
          <p class="text-sm text-gray-600">سهولت استفاده و زیبایی ظاهری سایت</p>
        </div>
      </div>
    </div>

    <!-- خلاصه امتیازات -->
    <div class="mt-8 bg-gradient-to-r from-gray-50 to-blue-50 rounded-xl p-6 border border-gray-200">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">خلاصه امتیازات</h3>
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-6">
        <div v-for="(rating, key) in ratings" :key="key" class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ rating }}</div>
          <div class="text-sm text-gray-600">{{ getCategoryName(key) }}</div>
        </div>
      </div>
      <div class="mt-4 pt-4 border-t border-gray-200">
        <div class="flex items-center justify-between">
          <span class="text-lg font-semibold text-gray-900">میانگین کلی:</span>
          <span class="text-2xl font-bold text-indigo-600">{{ averageRating.toFixed(1) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface DetailedRatings {
  productQuality: number
  pricing: number
  deliverySpeed: number
  packaging: number
  afterSales: number
  userInterface: number
}

const ratings = ref<DetailedRatings>({
  productQuality: 0,
  pricing: 0,
  deliverySpeed: 0,
  packaging: 0,
  afterSales: 0,
  userInterface: 0
})

const getRatingText = (rating: number): string => {
  switch (rating) {
    case 1: return 'خیلی بد'
    case 2: return 'بد'
    case 3: return 'متوسط'
    case 4: return 'خوب'
    case 5: return 'عالی'
    default: return 'امتیاز دهید'
  }
}

const getCategoryName = (key: string): string => {
  const names = {
    productQuality: 'کیفیت',
    pricing: 'قیمت',
    deliverySpeed: 'تحویل',
    packaging: 'بسته‌بندی',
    afterSales: 'پشتیبانی',
    userInterface: 'رابط کاربری'
  }
  return names[key] || key
}

const averageRating = computed(() => {
  const values = Object.values(ratings.value)
  const sum = values.reduce((acc, val) => acc + val, 0)
  return sum / values.length
})
</script>

<style scoped>
button:focus {
  outline: 2px solid #6366f1;
  outline-offset: 2px;
}

/* Custom hover effects */
.hover\:scale-110:hover {
  transform: scale(1.1);
}

/* Smooth transitions */
.transition-all {
  transition: all 0.2s ease-in-out;
}
</style> 
