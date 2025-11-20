<template>
  <div class="cart-items-list bg-white rounded-2xl shadow-xl p-6">
    <h2 class="text-xl font-bold mb-6 text-gray-800">سبد خرید شما</h2>
    <div v-if="items.length === 0" class="text-center text-gray-400 py-12">
      سبد خرید شما خالی است.
    </div>
    <div v-else class="flex flex-col gap-6">
      <div v-for="item in items" :key="item.id" class="flex items-center gap-6 border-b pb-5 last:border-b-0 last:pb-0">
        <img :src="item.image" class="w-20 h-20 rounded-xl object-cover border" alt="product" />
        <div class="flex-1">
          <div class="font-bold text-lg text-gray-800">{{ item.name }}</div>
          <div class="text-sm text-gray-500 mt-1">{{ item.variant }}</div>
          <div class="text-xs text-gray-400 mt-1">کد: {{ item.sku }}</div>
        </div>
        <div class="flex items-center gap-2">
          <button class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-100 hover:bg-gray-200 text-lg font-bold" @click="decrease(item)">-</button>
          <span class="w-8 text-center">{{ item.qty }}</span>
          <button class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-100 hover:bg-gray-200 text-lg font-bold" @click="increase(item)">+</button>
        </div>
        <div class="w-24 text-center font-bold text-base text-blue-600">{{ formatPrice(item.price * item.qty) }}</div>
        <button class="ml-2 text-red-500 hover:text-red-700" @click="remove(item)"><svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg></button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

interface CartItem {
  id: number;
  name: string;
  variant: string;
  sku: string;
  image: string;
  price: number;
  qty: number;
}

const items = ref<CartItem[]>([
  { id: 1, name: 'گوشی موبایل سامسونگ', variant: '128GB - مشکی', sku: 'SM-A123', image: '/public/statics/images/sample1.jpg', price: 18500000, qty: 1 },
  { id: 2, name: 'هدفون بی‌سیم سونی', variant: 'WH-1000XM4', sku: 'SONY-1000XM4', image: '/public/statics/images/sample2.jpg', price: 9500000, qty: 2 },
])

function increase(item: CartItem) { item.qty++ }
function decrease(item: CartItem) { if (item.qty > 1) item.qty-- }
function remove(item: CartItem) { items.value = items.value.filter(i => i.id !== item.id) }
function formatPrice(val: number) { return val.toLocaleString('fa-IR') + ' تومان' }
</script> 
