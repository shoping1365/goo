<template>
  <div class="cart-badge relative" style="font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;">
    <NuxtLink to="/Cart" class="flex items-center gap-2 text-gray-700 hover:text-blue-600 transition-colors">
      <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m-2.4 0L3 3m4 10v6a1 1 0 001 1h8a1 1 0 001-1v-6m-9 0V9a1 1 0 011-1h6a1 1 0 011 1v4.01"></path>
      </svg>
      <span class="text-sm font-medium">سبد خرید</span>
      
      <!-- Badge تعداد -->
      <div v-if="cartCount > 0" class="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center font-bold">
        {{ cartCount > 99 ? '99+' : cartCount }}
      </div>
    </NuxtLink>
  </div>
</template>

<script setup>
// استفاده از composable سبد خرید
const { cartCount, fetchCart } = useCart()

// دریافت سبد خرید در ابتدای بارگذاری
onMounted(() => {
  fetchCart()
})
</script> 