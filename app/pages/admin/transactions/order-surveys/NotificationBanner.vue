<template>
  <div class="bg-gradient-to-r from-blue-100 to-purple-100 border border-blue-200 rounded-xl p-6 flex items-center gap-6 my-4">
    <svg class="w-8 h-8 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M12 20a8 8 0 100-16 8 8 0 000 16z"/>
    </svg>
    <div class="flex-1 text-blue-800 font-medium">
      با تکمیل نظرسنجی، شانس شرکت در قرعه‌کشی ماهانه را خواهید داشت!
    </div>
  </div>
</template>

<script setup lang="ts">
// Notification banner component with Vue 3 composition API
</script>

<style scoped>
/* Smooth border animation */
.border-blue-200 {
  animation: borderPulse 3s ease-in-out infinite;
}

@keyframes borderPulse {
  0%, 100% {
    border-color: rgb(191 219 254);
  }
  50% {
    border-color: rgb(147 197 253);
  }
}
</style> 
