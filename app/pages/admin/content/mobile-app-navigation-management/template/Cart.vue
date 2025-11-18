<template>
  <NuxtLink to="/cart" class="mobile-nav-item">
    <div class="relative">
      <svg class="w-6 h-6 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
      </svg>
      <!-- نمایش تعداد آیتم‌های سبد خرید -->
      <span v-if="cartCount > 0" class="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">
        {{ cartCount }}
      </span>
    </div>
    <span class="text-xs font-medium">سبد خرید</span>
  </NuxtLink>
</template>

<script setup>
// آیتم ناوبری موبایل - سبد خرید
// این کامپوننت برای نمایش آیتم سبد خرید در ناوبری پایین موبایل استفاده می‌شود

// دریافت تعداد آیتم‌های سبد خرید
const cartCount = ref(0)

// TODO: اتصال به store یا composable سبد خرید
onMounted(() => {
  // اینجا باید تعداد آیتم‌های سبد خرید را از store دریافت کنید
  // cartCount.value = useCartStore().items.length
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

.mobile-nav-item.router-link-active {
  color: rgb(55 65 81);
  background-color: rgb(254 242 242);
}
</style>
