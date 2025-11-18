<template>
  <div class="bg-gradient-to-br from-green-50 via-white to-blue-50 rounded-2xl shadow-xl border border-gray-100 p-12 text-center my-8">
    <div class="flex flex-col items-center justify-center">
      <div class="text-5xl mb-4">🙏</div>
      <h2 class="text-3xl font-bold text-green-700 mb-2">از شما متشکریم!</h2>
      <p class="text-lg text-gray-700 mb-4">نظرات شما برای ما بسیار ارزشمند است و به بهبود خدمات ما کمک می‌کند.</p>
      <div class="text-md text-gray-500">در صورت نیاز به پیگیری، تیم پشتیبانی با شما تماس خواهد گرفت.</div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Thank you message component with Vue 3 composition API
</script>

<style scoped>
/* Smooth animation for thank you message */
.bg-gradient-to-br {
  animation: fadeInUp 0.8s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style> 
