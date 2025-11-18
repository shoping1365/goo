<template>
  <div class="cart-recommendations bg-white/80 rounded-2xl shadow p-5">
    <h4 class="font-bold text-gray-700 mb-4">پیشنهاد برای شما</h4>
    <div class="flex gap-6 overflow-x-auto hide-scrollbar pb-2">
      <div v-for="item in recommendations" :key="item.id" class="min-w-[180px] bg-gradient-to-br from-blue-100 to-purple-100 rounded-xl shadow hover:scale-105 transition p-3 flex flex-col items-center">
        <img :src="item.image" class="w-20 h-20 rounded-lg object-cover mb-2 border" alt="product" />
        <div class="font-bold text-sm text-gray-800 text-center">{{ item.name }}</div>
        <div class="text-xs text-gray-500 mt-1">{{ formatPrice(item.price) }}</div>
        <button class="mt-2 px-3 py-1 bg-blue-500 text-white rounded-lg text-xs font-bold hover:bg-blue-600 transition">افزودن به سبد</button>
      </div>
    </div>
  </div>
</template>

<script setup>
const recommendations = [
  { id: 1, name: 'پاوربانک شیائومی', image: '/public/statics/images/sample3.jpg', price: 850000 },
  { id: 2, name: 'کیف لپ‌تاپ', image: '/public/statics/images/sample4.jpg', price: 420000 },
  { id: 3, name: 'ساعت هوشمند', image: '/public/statics/images/sample5.jpg', price: 3200000 },
  { id: 4, name: 'کابل شارژ سریع', image: '/public/statics/images/sample6.jpg', price: 180000 },
]
function formatPrice(val) { return val.toLocaleString('fa-IR') + ' تومان' }
</script> 
