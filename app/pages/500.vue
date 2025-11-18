<template>
  <div dir="rtl" class="h-screen flex flex-col items-center justify-center px-4 text-center font-iransans bg-gradient-to-b from-red-50 via-orange-50 to-yellow-50">
      <!-- کد خطا -->
      <div class="w-full text-center mb-12">
        <h1 class="color-pulse w-full text-center text-[80px] sm:text-[100px] md:text-[120px] lg:text-[140px] font-black text-transparent bg-clip-text bg-gradient-to-r from-red-600 via-orange-600 to-yellow-600 leading-none drop-shadow-2xl">
          500
        </h1>
      </div>

      <!-- متن توضیحات -->
      <div class="w-full text-center mb-8 space-y-3 max-w-xl px-4">
        <p class="w-full text-center text-sm md:text-lg text-gray-700">
          بعضی وقت‌ها بعضی چیزا درست کار نمیکنه، یه کم دیگه بهمون سر بزنید.
        </p>
      </div>
  </div>
</template>

<script setup lang="ts">
// صفحه اختصاصی 500
definePageMeta({
  layout: false
})
</script>

<style scoped>
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
