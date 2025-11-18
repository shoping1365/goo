<template>
  <NuxtLink to="/categories" class="mobile-nav-item">
    <svg class="w-6 h-6 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path>
    </svg>
    <span class="text-xs font-medium">دسته‌ها</span>
  </NuxtLink>
</template>

<script setup>
// آیتم ناوبری موبایل - دسته‌ها
// این کامپوننت برای نمایش آیتم دسته‌ها در ناوبری پایین موبایل استفاده می‌شود
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
  background-color: rgb(167, 155, 202);
}
</style>
