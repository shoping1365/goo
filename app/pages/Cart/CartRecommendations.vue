<template>
  <div class="cart-recommendations bg-white rounded-2xl shadow p-6 mt-8">
    <div class="font-bold text-base text-[#1a2341] mb-4">خریداران این محصولات، محصولات زیر را هم خریده‌اند</div>
    <div class="flex gap-6 overflow-x-auto hide-scrollbar pb-2">
      <div v-for="item in recommendations" :key="item.id" class="min-w-[180px] bg-gradient-to-br from-[#fff] to-[#f7fafd] rounded-xl shadow hover:scale-105 transition p-3 flex flex-col items-center relative">
        <span class="absolute top-2 right-2 text-xs text-[#e60023] font-bold">فروش ویژه</span>
        <img :src="item.image" class="w-20 h-20 rounded-lg object-contain mb-2 border" alt="product" />
        <div class="font-bold text-sm text-[#1a2341] text-center">{{ item.name }}</div>
        <div class="text-xs text-gray-500 mt-1">{{ formatPrice(item.price) }}</div>
        <button class="mt-2 px-3 py-1 bg-[#008eb2] text-white rounded-lg text-xs font-bold hover:bg-[#007a99] transition">افزودن به سبد</button>
      </div>
    </div>
  </div>
</template>

<script setup>
const recommendations = []
function formatPrice(val) { return val.toLocaleString('fa-IR') + ' تومان' }
</script> 
