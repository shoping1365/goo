<template>
  <button class="mobile-nav-item" @click="toggleDarkMode">
    <svg v-if="isDark" class="w-6 h-6 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"></path>
    </svg>
    <svg v-else class="w-6 h-6 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"></path>
    </svg>
    <span class="text-xs font-medium">حالت تاریک</span>
  </button>
</template>

<script setup>
// آیتم ناوبری موبایل - حالت تاریک
// این کامپوننت برای نمایش آیتم تغییر تم در ناوبری پایین موبایل استفاده می‌شود

const isDark = ref(false)

const toggleDarkMode = () => {
  isDark.value = !isDark.value
  // TODO: اتصال به سیستم تغییر تم
  // useTheme().toggle()
  
  // تغییر کلاس dark در html
  if (isDark.value) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
}

// بررسی حالت اولیه
onMounted(() => {
  isDark.value = document.documentElement.classList.contains('dark')
})
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

.mobile-nav-item:active {
  color: rgb(55 65 81);
  background-color: rgb(254 242 242);
}
</style>
