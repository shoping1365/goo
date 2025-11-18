<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">تنظیمات برنامه وفاداری</h2>
      <button 
        @click="saveSettings"
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
      >
        ذخیره تنظیمات
      </button>
    </div>

    <!-- تنظیمات عمومی -->
    <div class="space-y-6">
      <!-- وضعیت برنامه -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت برنامه</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">فعال کردن برنامه وفاداری</label>
              <p class="text-sm text-gray-500">برنامه وفاداری را فعال یا غیرفعال کنید</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.programEnabled" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">نمایش عمومی</label>
              <p class="text-sm text-gray-500">نمایش برنامه وفاداری برای همه کاربران</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.publicDisplay" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>
      </div>

      <!-- تنظیمات امتیازدهی -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات امتیازدهی</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700">امتیاز برای هر تومان خرید</label>
            <input 
              v-model="settings.pointsPerToman" 
              type="number" 
              min="0"
              step="0.01"
              class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700">حداقل مبلغ خرید برای امتیاز</label>
            <input 
              v-model="settings.minPurchaseForPoints" 
              type="number" 
              min="0"
              class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700">حداکثر امتیاز روزانه</label>
            <input 
              v-model="settings.maxDailyPoints" 
              type="number" 
              min="0"
              class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700">حداکثر امتیاز ماهانه</label>
            <input 
              v-model="settings.maxMonthlyPoints" 
              type="number" 
              min="0"
              class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
          </div>
        </div>
      </div>

      <!-- تنظیمات انقضا -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات انقضا</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700">انقضای امتیازات (روز)</label>
            <input 
              v-model="settings.pointsExpiryDays" 
              type="number" 
              min="0"
              class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
            <p class="text-sm text-gray-500 mt-1">0 = بدون انقضا</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700">انقضای پاداش‌ها (روز)</label>
            <input 
              v-model="settings.rewardsExpiryDays" 
              type="number" 
              min="0"
              class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
            <p class="text-sm text-gray-500 mt-1">0 = بدون انقضا</p>
          </div>
        </div>
      </div>

      <!-- تنظیمات امنیتی -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات امنیتی</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">تایید دو مرحله‌ای برای پاداش‌های بزرگ</label>
              <p class="text-sm text-gray-500">تایید اضافی برای پاداش‌های بالای 1000 امتیاز</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.twoFactorForLargeRewards" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">محدودیت IP</label>
              <p class="text-sm text-gray-500">محدود کردن استفاده از یک IP</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.ipRestriction" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">لاگ کامل فعالیت‌ها</label>
              <p class="text-sm text-gray-500">ثبت تمام فعالیت‌های برنامه وفاداری</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.fullActivityLog" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>
      </div>

      <!-- تنظیمات اعلان‌ها -->
      <div class="border border-gray-200 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اعلان‌ها</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان دریافت امتیاز</label>
              <p class="text-sm text-gray-500">ارسال اعلان هنگام دریافت امتیاز</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.notifyPointsEarned" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان ارتقا سطح</label>
              <p class="text-sm text-gray-500">ارسال اعلان هنگام ارتقا سطح</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.notifyLevelUpgrade" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <label class="text-sm font-medium text-gray-700">اعلان انقضای امتیاز</label>
              <p class="text-sm text-gray-500">اعلان قبل از انقضای امتیازات</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="settings.notifyPointsExpiry" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تنظیمات برنامه وفاداری
const settings = ref({
  // وضعیت برنامه
  programEnabled: true,
  publicDisplay: true,
  
  // تنظیمات امتیازدهی
  pointsPerToman: 1,
  minPurchaseForPoints: 10000,
  maxDailyPoints: 1000,
  maxMonthlyPoints: 10000,
  
  // تنظیمات انقضا
  pointsExpiryDays: 365,
  rewardsExpiryDays: 30,
  
  // تنظیمات امنیتی
  twoFactorForLargeRewards: true,
  ipRestriction: false,
  fullActivityLog: true,
  
  // تنظیمات اعلان‌ها
  notifyPointsEarned: true,
  notifyLevelUpgrade: true,
  notifyPointsExpiry: true
})

// ذخیره تنظیمات
function saveSettings() {
  // TODO: فراخوانی API برای ذخیره تنظیمات
  console.log('تنظیمات ذخیره شد:', settings.value)
  alert('تنظیمات با موفقیت ذخیره شد')
}
</script> 
