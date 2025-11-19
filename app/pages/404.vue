<template>
  <div dir="rtl" class="h-screen flex flex-col items-center justify-center px-4 text-center font-iransans bg-gradient-to-b from-red-50 via-orange-50 to-yellow-50">
      <!-- کد خطا -->
      <div class="w-full text-center mb-6">
        <h1 class="color-pulse w-full text-center text-[100px] sm:text-[130px] md:text-[160px] lg:text-[180px] font-black text-transparent bg-clip-text bg-gradient-to-r from-red-600 via-orange-600 to-yellow-600 leading-none drop-shadow-2xl">
          404
        </h1>
      </div>

      <!-- متن توضیحات -->
      <div class="w-full text-center mb-8 space-y-3 max-w-xl px-4">
        <h2 class="w-full text-center text-xl md:text-2xl font-bold text-gray-800 drop-shadow-lg">
          صفحه یافت نشد
        </h2>
        <p class="w-full text-center text-sm md:text-lg text-gray-700">
          صفحه‌ای که به دنبال آن هستید وجود ندارد.
        </p>
      </div>

      <!-- دکمه بازگشت در باکس -->
      <div class="bg-white/80 backdrop-blur-sm rounded-3xl p-6 md:p-8 shadow-2xl border-2 border-orange-200 max-w-sm w-full">
        <NuxtLink
          to="/"
          class="group relative inline-flex w-full justify-center px-8 md:px-10 py-4 md:py-5 bg-gradient-to-r from-red-600 via-orange-600 to-yellow-600 text-white rounded-2xl font-semibold text-base md:text-lg shadow-xl hover:shadow-2xl transform hover:scale-105 transition-all duration-300 overflow-hidden"
        >
          <span class="absolute inset-0 bg-gradient-to-r from-orange-600 via-yellow-600 to-red-600 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></span>
          <span class="relative z-10 flex items-center justify-center space-x-2 rtl:space-x-reverse">
            <svg class="w-6 h-6 transform group-hover:-translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
            </svg>
            <span>بازگشت به صفحه اصلی</span>
          </span>
        </NuxtLink>
      </div>
  </div>
</template>

<script setup lang="ts">
// صفحه اختصاصی 404
// @ts-ignore: Nuxt macro
definePageMeta({
  layout: false
})
</script>

<style scoped>
h1 { 
  animation: glow 2s ease-in-out infinite; 
}

.color-pulse {
  animation: color-pulse 3s ease-in-out infinite, glow 2s ease-in-out infinite;
}

@keyframes color-pulse {
  0%, 100% {
    background: linear-gradient(to right, #dc2626, #ea580c, #ca8a04);
    -webkit-background-clip: text;
    background-clip: text;
    opacity: 0.8;
  }
  50% {
    background: linear-gradient(to right, #ef4444, #f97316, #eab308);
    -webkit-background-clip: text;
    background-clip: text;
    opacity: 1;
  }
}

@keyframes glow {
  0%, 100% {
    filter: drop-shadow(0 0 20px rgba(251, 146, 60, 0.5)) drop-shadow(0 0 40px rgba(249, 115, 22, 0.3));
  }
  50% {
    filter: drop-shadow(0 0 30px rgba(251, 146, 60, 0.8)) drop-shadow(0 0 60px rgba(249, 115, 22, 0.5)) drop-shadow(0 0 90px rgba(234, 88, 12, 0.3));
  }
}
</style>
