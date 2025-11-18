<template>
  <div>
    <div v-if="items.length === 0" class="bg-white rounded-2xl shadow p-8 text-center text-gray-400 text-lg">سبد خرید خالی است.</div>
    <div v-else class="flex flex-col">
      <div v-for="item in items" :key="item.id" class="cart-item flex items-center gap-6 p-6 bg-white border border-gray-200 mb-2">
        <!-- تصویر محصول (سمت راست) -->
        <div class="flex-shrink-0">
          <NuxtLink :to="`/product/sku-${item.sku || item.product_id || item.id}`" class="block">
            <img :src="item.image" class="product-img w-20 h-20 rounded-xl object-contain border border-gray-200 bg-[#f8fafc] hover:border-blue-300 transition-colors" alt="product" />
          </NuxtLink>
        </div>
        
        <!-- جزئیات محصول (وسط) -->
        <div class="flex-1 flex flex-col gap-1 text-right">
          <NuxtLink :to="`/product/sku-${item.sku || item.product_id || item.id}`" class="font-bold text-base text-[#1a2341] leading-snug hover:text-blue-600 transition-colors cursor-pointer">
            {{ item.name }}
          </NuxtLink>
          <div v-if="item.features && item.features.length" class="flex flex-wrap gap-1 text-[11px] text-gray-500 mt-1">
            <span v-for="f in item.features" :key="f" class="after:content-['•'] last:after:content-[''] after:px-1">{{ f }}</span>
          </div>
          <div class="text-[11px] text-gray-400 mt-0.5">کد: {{ item.sku }}</div>
        </div>
        
        <!-- قیمت و تعداد (سمت چپ) -->
        <div class="flex flex-col items-end gap-2">
          <div class="font-bold text-lg text-[#1a2341]">{{ formatPrice((item.price || 0) * (item.qty || 1)) }}</div>
          
          <div v-if="!isNext" class="flex items-center gap-2 mt-1">
            <button @click="$emit('change-qty', {item, val: item.qty+1})" class="qty-btn w-7 h-7 flex items-center justify-center rounded-full border border-gray-200 bg-white hover:bg-gray-50 text-gray-700">+</button>
            <span class="w-5 text-center select-none">{{ item.qty }}</span>
            <button v-if="item.qty > 1" @click="$emit('change-qty', {item, val: item.qty-1})" class="qty-btn w-7 h-7 flex items-center justify-center rounded-full border border-gray-200 bg-white hover:bg-gray-50 text-gray-700">-</button>
            <button v-else @click="$emit('remove', item.id)" class="qty-btn w-7 h-7 flex items-center justify-center rounded-full border border-red-200 bg-white hover:bg-red-50">
              <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5-4h4a2 2 0 012 2v2H7V5a2 2 0 012-2zm0 0V3m0 0h4"/></svg>
            </button>
          </div>
          
          <div class="mt-2 text-right flex gap-2 justify-end">
            <button v-if="!isNext" @click="$emit('move-to-next', item)" class="text-blue-500 text-xs hover:underline whitespace-nowrap">انتقال به خرید بعدی</button>
            <button v-else @click="$emit('move-to-cart', item)" class="text-red-500 text-xs hover:underline whitespace-nowrap">انتقال به سبد خرید</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({ items: { type: Array, default: () => [] }, isNext: { type: Boolean, default: false } })
function formatPrice(val) { return val.toLocaleString('fa-IR') + ' تومان' }
</script>

<style scoped>
.cart-item {
  background: #fff;
  transition: background-color .18s;
  align-items: center;
  min-height: 80px;
  font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;
}
.cart-item:hover {
  background-color: #fafafa;
}
.product-img {
  border-radius: 0.5rem;
  border: 1px solid #e2e8f0;
  background: #f8fafc;
}
.qty-btn {
  background: #fff;
  border: 1px solid #e2e8f0;
  color: #334155;
  font-weight: bold;
  border-radius: 50%;
  width: 1.75rem;
  height: 1.75rem;
  transition: background .18s, color .18s;
}
.qty-btn:hover:not(:disabled) {
  background: #f8fafc;
}
.remove-btn {
  color: #e11d48;
  background: #fef2f2;
  border-radius: 0.75rem;
  padding: 0.25rem 0.75rem;
  font-weight: bold;
  transition: background .18s, color .18s;
  margin-right: 0.5rem;
}
.remove-btn:hover {
  background: #e11d48;
  color: #fff;
}
.move-btn {
  color: #0284c7;
  background: #e0f2fe;
  border-radius: 0.75rem;
  padding: 0.25rem 0.75rem;
  font-weight: bold;
  transition: background .18s, color .18s;
}
.move-btn:hover {
  background: #0284c7;
  color: #fff;
}
</style> 
