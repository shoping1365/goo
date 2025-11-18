<template>
  <div class="bg-white rounded-2xl shadow-xl border border-gray-100 p-8 mb-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-yellow-400 to-pink-500 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">پاداش و امتیاز وفاداری</h2>
        <p class="text-gray-600 mt-1">با تکمیل نظرسنجی، امتیاز یا کد تخفیف دریافت کنید</p>
      </div>
    </div>
    <div class="space-y-6">
      <div class="bg-gradient-to-r from-yellow-50 to-pink-50 rounded-xl p-6 border border-yellow-100 text-center">
        <div class="text-3xl font-bold text-yellow-600 mb-2">🎁</div>
        <div class="text-lg text-gray-800 mb-2">شما با تکمیل این نظرسنجی <span class="font-bold text-pink-600">۵۰ امتیاز وفاداری</span> دریافت می‌کنید!</div>
        <div class="text-sm text-gray-500">امتیازها قابل استفاده برای خریدهای بعدی یا دریافت کد تخفیف هستند.</div>
      </div>
      <div class="bg-gradient-to-r from-pink-50 to-yellow-50 rounded-xl p-6 border border-pink-100 text-center">
        <div class="text-lg text-gray-800 mb-2">کد تخفیف ویژه شما:</div>
        <div class="text-2xl font-bold text-pink-600 tracking-widest select-all">SURVEY2024</div>
        <div class="text-sm text-gray-500 mt-2">این کد را در خرید بعدی وارد کنید و ۱۰٪ تخفیف بگیرید.</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Simple reward component with Vue 3 composition API
</script>

<style scoped>
.select-all {
  user-select: all;
  -webkit-user-select: all;
  -moz-user-select: all;
  -ms-user-select: all;
}
</style> 
