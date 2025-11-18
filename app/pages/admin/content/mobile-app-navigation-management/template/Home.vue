<template>
  <NuxtLink to="/" class="mobile-nav-item">
    <svg class="w-6 h-6 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
    </svg>
    <span class="text-xs font-medium">خانه</span>
  </NuxtLink>
</template>

<script setup>
// آیتم ناوبری موبایل - خانه
// این کامپوننت برای نمایش آیتم خانه در ناوبری پایین موبایل استفاده می‌شود
</script>

<style scoped>
.mobile-nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  transition: color 0.2s;
  color: rgb(107 114 128);
}

.mobile-nav-item:hover {
  color: rgb(55 65 81);
}

.mobile-nav-item.router-link-active {
  color: rgb(55 65 81);
  background-color: rgb(254 242 242);
}
</style>
