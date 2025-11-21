<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">تنظیمات برنامه ارجاع</h2>
      <button 
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
        @click="saveSettings"
      >
        ذخیره تنظیمات
      </button>
    </div>

    <div class="space-y-8">
      <!-- تنظیمات عمومی -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات عمومی</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فعال بودن برنامه</label>
            <div class="flex items-center">
              <input 
                v-model="settings.isActive" 
                type="checkbox" 
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              >
              <span class="mr-2 text-sm text-gray-700">برنامه ارجاع فعال باشد</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ خرید</label>
            <input 
              v-model="settings.minPurchaseAmount" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 100000"
            >
            <p class="text-xs text-gray-500 mt-1">حداقل مبلغ خرید برای اعتبار ارجاع (تومان)</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر پاداش ماهانه</label>
            <input 
              v-model="settings.maxMonthlyReward" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 500000"
            >
            <p class="text-xs text-gray-500 mt-1">حداکثر پاداش قابل دریافت در ماه (تومان)</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار کد</label>
            <input 
              v-model="settings.codeExpiryDays" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 30"
            >
            <p class="text-xs text-gray-500 mt-1">مدت اعتبار کد ارجاع (روز)</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات پاداش‌دهی -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات پاداش‌دهی</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">درصد پاداش پیش‌فرض</label>
            <input 
              v-model="settings.defaultRewardPercentage" 
              type="number" 
              min="0" 
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 10"
            >
            <p class="text-xs text-gray-500 mt-1">درصد پاداش پیش‌فرض برای ارجاعات موفق</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پاداش ثابت پیش‌فرض</label>
            <input 
              v-model="settings.defaultFixedReward" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 25000"
            >
            <p class="text-xs text-gray-500 mt-1">پاداش ثابت پیش‌فرض (تومان)</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل فاصله زمانی</label>
            <input 
              v-model="settings.minTimeBetweenReferrals" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 24"
            >
            <p class="text-xs text-gray-500 mt-1">حداقل فاصله زمانی بین ارجاعات (ساعت)</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر ارجاع روزانه</label>
            <input 
              v-model="settings.maxDailyReferrals" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 5"
            >
            <p class="text-xs text-gray-500 mt-1">حداکثر تعداد ارجاع مجاز در روز</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات اعلان‌ها -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اعلان‌ها</h3>
        
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان ارجاع موفق</label>
              <p class="text-xs text-gray-500">ارسال اعلان به کاربر هنگام موفقیت ارجاع</p>
            </div>
            <input 
              v-model="settings.notifications.successfulReferral" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            >
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان پاداش جدید</label>
              <p class="text-xs text-gray-500">اعلان به کاربر هنگام دریافت پاداش جدید</p>
            </div>
            <input 
              v-model="settings.notifications.newReward" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            >
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان انقضای کد</label>
              <p class="text-xs text-gray-500">اعلان قبل از انقضای کد ارجاع</p>
            </div>
            <input 
              v-model="settings.notifications.codeExpiry" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            >
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان ایمیل</label>
              <p class="text-xs text-gray-500">ارسال اعلان‌ها از طریق ایمیل</p>
            </div>
            <input 
              v-model="settings.notifications.email" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            >
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان پیامک</label>
              <p class="text-xs text-gray-500">ارسال اعلان‌ها از طریق پیامک</p>
            </div>
            <input 
              v-model="settings.notifications.sms" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            >
          </div>
        </div>
      </div>

      <!-- تنظیمات امنیتی -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات امنیتی</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تایید IP</label>
            <div class="flex items-center">
              <input 
                v-model="settings.security.ipVerification" 
                type="checkbox" 
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              >
              <span class="mr-2 text-sm text-gray-700">تایید IP برای جلوگیری از تقلب</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تایید شماره موبایل</label>
            <div class="flex items-center">
              <input 
                v-model="settings.security.phoneVerification" 
                type="checkbox" 
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              >
              <span class="mr-2 text-sm text-gray-700">تایید شماره موبایل برای ارجاع</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش روزانه</label>
            <input 
              v-model="settings.security.maxDailyAttempts" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 10"
            >
            <p class="text-xs text-gray-500 mt-1">حداکثر تعداد تلاش برای ارجاع در روز</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت مسدودیت</label>
            <input 
              v-model="settings.security.blockDuration" 
              type="number" 
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="مثال: 24"
            >
            <p class="text-xs text-gray-500 mt-1">مدت مسدودیت پس از تجاوز از حد مجاز (ساعت)</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تنظیمات برنامه ارجاع
const settings = ref({
  // تنظیمات عمومی
  isActive: true,
  minPurchaseAmount: 100000,
  maxMonthlyReward: 500000,
  codeExpiryDays: 30,
  
  // تنظیمات پاداش‌دهی
  defaultRewardPercentage: 10,
  defaultFixedReward: 25000,
  minTimeBetweenReferrals: 24,
  maxDailyReferrals: 5,
  
  // تنظیمات اعلان‌ها
  notifications: {
    successfulReferral: true,
    newReward: true,
    codeExpiry: true,
    email: true,
    sms: false
  },
  
  // تنظیمات امنیتی
  security: {
    ipVerification: true,
    phoneVerification: true,
    maxDailyAttempts: 10,
    blockDuration: 24
  }
})

// ذخیره تنظیمات
function saveSettings() {
  // TODO: فراخوانی API برای ذخیره تنظیمات
  alert('تنظیمات با موفقیت ذخیره شد!')
}
</script> 
