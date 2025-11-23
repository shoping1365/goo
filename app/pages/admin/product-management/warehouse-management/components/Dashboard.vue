<template>
  <div v-if="hasAccess" class="space-y-6">
    <!-- شاخص‌های کلیدی عملکرد (KPIs) -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- ارزش کل موجودی -->
      <div class="bg-gradient-to-br from-emerald-500 to-teal-600 text-white rounded-2xl p-6 shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-emerald-100 text-sm">ارزش کل موجودی</p>
            <p class="text-3xl font-bold">۱۲,۴۵۰,۰۰۰</p>
            <p class="text-emerald-100 text-xs">تومان</p>
          </div>
          <div class="h-16 w-16 bg-white/20 rounded-full flex items-center justify-center">
            <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
        </div>
        <div class="mt-4 text-emerald-100 text-sm">
          <span class="bg-white/20 px-3 py-1 rounded-full">+۱۲٪ نسبت به دیروز</span>
        </div>
      </div>

      <!-- نرخ گردش موجودی -->
      <div class="bg-gradient-to-br from-blue-500 to-indigo-600 text-white rounded-2xl p-6 shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-blue-100 text-sm">نرخ گردش موجودی</p>
            <p class="text-3xl font-bold">۸.۵</p>
            <p class="text-blue-100 text-xs">دفعه در ماه</p>
          </div>
          <div class="h-16 w-16 bg-white/20 rounded-full flex items-center justify-center">
            <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
          </div>
        </div>
        <div class="mt-4 text-blue-100 text-sm">
          <span class="bg-white/20 px-3 py-1 rounded-full">هدف: ۱۰ دفعه</span>
        </div>
      </div>

      <!-- نرخ تحقق سفارش -->
      <div class="bg-gradient-to-br from-green-500 to-emerald-600 text-white rounded-2xl p-6 shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-green-100 text-sm">نرخ تحقق سفارش</p>
            <p class="text-3xl font-bold">۹۶.۸٪</p>
            <p class="text-green-100 text-xs">به‌موقع ارسال شده</p>
          </div>
          <div class="h-16 w-16 bg-white/20 rounded-full flex items-center justify-center">
            <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
        <div class="mt-4 text-green-100 text-sm">
          <span class="bg-white/20 px-3 py-1 rounded-full">+۲.۱٪ نسبت به هفته گذشته</span>
        </div>
      </div>

      <!-- حجم ورودی و خروجی -->
      <div class="bg-gradient-to-br from-purple-500 to-pink-600 text-white rounded-2xl p-6 shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-purple-100 text-sm">حجم ورودی و خروجی</p>
            <p class="text-3xl font-bold">۲۴۷</p>
            <p class="text-purple-100 text-xs">محموله امروز</p>
          </div>
          <div class="h-16 w-16 bg-white/20 rounded-full flex items-center justify-center">
            <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4" />
            </svg>
          </div>
        </div>
        <div class="mt-4 text-purple-100 text-sm">
          <span class="bg-white/20 px-3 py-1 rounded-full">ورودی: ۸۹ | خروجی: ۱۵۸</span>
        </div>
      </div>

      <!-- کالاهای با موجودی کم -->
      <div class="bg-gradient-to-br from-orange-500 to-red-500 text-white rounded-2xl p-6 shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-orange-100 text-sm">کالاهای با موجودی کم</p>
            <p class="text-3xl font-bold">۲۳</p>
            <p class="text-orange-100 text-xs">کالا نیاز به سفارش</p>
          </div>
          <div class="h-16 w-16 bg-white/20 rounded-full flex items-center justify-center">
            <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.34 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
          </div>
        </div>
        <div class="mt-4 text-orange-100 text-sm">
          <span class="bg-white/20 px-3 py-1 rounded-full">نیاز به اقدام فوری</span>
        </div>
      </div>

      <!-- ظرفیت انبار -->
      <div class="bg-gradient-to-br from-gray-600 to-gray-800 text-white rounded-2xl p-6 shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-100 text-sm">ظرفیت انبار</p>
            <p class="text-3xl font-bold">۷۸٪</p>
            <p class="text-gray-100 text-xs">پر شده</p>
          </div>
          <div class="h-16 w-16 bg-white/20 rounded-full flex items-center justify-center">
            <svg class="h-8 w-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
        </div>
        <div class="mt-4">
          <div class="w-full bg-white/20 rounded-full h-3">
            <div class="bg-white h-3 rounded-full transition-all duration-500" style="width: 78%"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- اقدامات سریع -->
    <div class="bg-white rounded-2xl shadow-lg p-6">
      <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
        <svg class="h-5 w-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
        </svg>
        اقدامات سریع
      </h3>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
        <button class="p-6 bg-gradient-to-r from-blue-500 to-blue-600 text-white rounded-xl hover:from-blue-600 hover:to-blue-700 transition-all duration-200 shadow-lg hover:shadow-xl">
          <div class="text-center">
            <svg class="h-8 w-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            <div class="font-medium text-sm">ایجاد سفارش خرید</div>
          </div>
        </button>
        
        <button class="p-6 bg-gradient-to-r from-green-500 to-green-600 text-white rounded-xl hover:from-green-600 hover:to-green-700 transition-all duration-200 shadow-lg hover:shadow-xl">
          <div class="text-center">
            <svg class="h-8 w-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
            <div class="font-medium text-sm">دریافت کالا</div>
          </div>
        </button>
        
        <button class="p-6 bg-gradient-to-r from-purple-500 to-purple-600 text-white rounded-xl hover:from-purple-600 hover:to-purple-700 transition-all duration-200 shadow-lg hover:shadow-xl">
          <div class="text-center">
            <svg class="h-8 w-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
            </svg>
            <div class="font-medium text-sm">ایجاد سفارش فروش</div>
          </div>
        </button>
        
        <button class="p-6 bg-gradient-to-r from-orange-500 to-orange-600 text-white rounded-xl hover:from-orange-600 hover:to-orange-700 transition-all duration-200 shadow-lg hover:shadow-xl">
          <div class="text-center">
            <svg class="h-8 w-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
            </svg>
            <div class="font-medium text-sm">شروع شمارش انبار</div>
          </div>
        </button>
      </div>
    </div>

    <!-- هشدارها و اعلان‌ها -->
    <div class="bg-white rounded-2xl shadow-lg p-6">
      <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
        <svg class="h-5 w-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.34 16.5c-.77.833.192 2.5 1.732 2.5z" />
        </svg>
        هشدارها و اعلان‌ها
      </h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="p-6 bg-red-50 rounded-lg border-r-4 border-red-500">
          <div class="flex items-center gap-3">
            <svg class="h-6 w-6 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.34 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
            <div>
              <div class="font-medium text-red-800">موجودی کم</div>
              <div class="text-sm text-red-600">۲۳ کالا نیاز به سفارش فوری دارند</div>
            </div>
          </div>
        </div>
        
        <div class="p-6 bg-yellow-50 rounded-lg border-r-4 border-yellow-500">
          <div class="flex items-center gap-3">
            <svg class="h-6 w-6 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <div>
              <div class="font-medium text-yellow-800">سفارش‌های تأخیردار</div>
              <div class="text-sm text-yellow-600">۸ سفارش بیش از ۲۴ ساعت تأخیر دارند</div>
            </div>
          </div>
        </div>
        
        <div class="p-6 bg-orange-50 rounded-lg border-r-4 border-orange-500">
          <div class="flex items-center gap-3">
            <svg class="h-6 w-6 text-orange-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <div>
              <div class="font-medium text-orange-800">محصولات در حال انقضا</div>
              <div class="text-sm text-orange-600">۵ محصول تا ۷ روز دیگر منقضی می‌شوند</div>
            </div>
          </div>
        </div>
        
        <div class="p-6 bg-blue-50 rounded-lg border-r-4 border-blue-500">
          <div class="flex items-center gap-3">
            <svg class="h-6 w-6 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
            </svg>
            <div>
              <div class="font-medium text-blue-800">مشکلات ارسال</div>
              <div class="text-sm text-blue-600">۳ محموله در مسیر با مشکل مواجه شده‌اند</div>
            </div>
          </div>
        </div>
        
        <div class="p-6 bg-purple-50 rounded-lg border-r-4 border-purple-500">
          <div class="flex items-center gap-3">
            <svg class="h-6 w-6 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <div>
              <div class="font-medium text-purple-800">کالاهای راکد</div>
              <div class="text-sm text-purple-600">۱۵ کالا بیش از ۶ ماه فروش نرفته‌اند</div>
            </div>
          </div>
        </div>
        
        <div class="p-6 bg-indigo-50 rounded-lg border-r-4 border-indigo-500">
          <div class="flex items-center gap-3">
            <svg class="h-6 w-6 text-indigo-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            <div>
              <div class="font-medium text-indigo-800">انبارگردانی</div>
              <div class="text-sm text-indigo-600">۲ انبار نیاز به شمارش دوره‌ای دارند</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مصورسازی داده‌ها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار سطح موجودی -->
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
          <svg class="h-5 w-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
          </svg>
          نمودار سطح موجودی
        </h3>
        <div class="bg-gradient-to-br from-green-50 to-emerald-100 rounded-xl p-6 h-48 flex items-center justify-center">
          <div class="text-center text-gray-600">
            <svg class="h-16 w-16 mx-auto mb-2 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            <p class="text-sm">نمودار سطح موجودی در طول زمان</p>
            <p class="text-xs text-gray-500">آخرین بروزرسانی: ۲ ساعت پیش</p>
          </div>
        </div>
      </div>

      <!-- تفکیک موجودی بر اساس دسته‌بندی -->
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
          <svg class="h-5 w-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z" />
          </svg>
          تفکیک موجودی بر اساس دسته‌بندی
        </h3>
        <div class="space-y-3">
          <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
            <span class="text-sm text-gray-700">الکترونیک</span>
            <div class="flex items-center gap-2">
              <div class="w-20 bg-blue-200 rounded-full h-2">
                <div class="bg-blue-600 h-2 rounded-full" style="width: 45%"></div>
              </div>
              <span class="text-sm font-medium text-blue-600">۴۵٪</span>
            </div>
          </div>
          <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
            <span class="text-sm text-gray-700">پوشیدنی</span>
            <div class="flex items-center gap-2">
              <div class="w-20 bg-green-200 rounded-full h-2">
                <div class="bg-green-600 h-2 rounded-full" style="width: 28%"></div>
              </div>
              <span class="text-sm font-medium text-green-600">۲۸٪</span>
            </div>
          </div>
          <div class="flex items-center justify-between p-3 bg-purple-50 rounded-lg">
            <span class="text-sm text-gray-700">صوتی</span>
            <div class="flex items-center gap-2">
              <div class="w-20 bg-purple-200 rounded-full h-2">
                <div class="bg-purple-600 h-2 rounded-full" style="width: 18%"></div>
              </div>
              <span class="text-sm font-medium text-purple-600">۱۸٪</span>
            </div>
          </div>
          <div class="flex items-center justify-between p-3 bg-orange-50 rounded-lg">
            <span class="text-sm text-gray-700">سایر</span>
            <div class="flex items-center gap-2">
              <div class="w-20 bg-orange-200 rounded-full h-2">
                <div class="bg-orange-600 h-2 rounded-full" style="width: 9%"></div>
              </div>
              <span class="text-sm font-medium text-orange-600">۹٪</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نقشه مقصدهای ارسال -->
    <div class="bg-white rounded-2xl shadow-lg p-6">
      <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
        <svg class="h-5 w-5 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m0 0L9 7" />
        </svg>
        نقشه مقصدهای ارسال‌های اخیر
      </h3>
      <div class="bg-gradient-to-br from-indigo-50 to-purple-100 rounded-xl p-6 h-48 flex items-center justify-center">
        <div class="text-center text-gray-600">
          <svg class="h-16 w-16 mx-auto mb-2 text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m0 0L9 7" />
          </svg>
          <p class="text-sm">نقشه تعاملی مقصدهای ارسال</p>
          <p class="text-xs text-gray-500">تهران: ۴۵ | اصفهان: ۲۳ | شیراز: ۱۸ | مشهد: ۱۲</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>;
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})
</script>
