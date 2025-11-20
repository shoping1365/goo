<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-teal-700 mb-2 flex justify-between items-center">
      <span>سبد خرید</span>
      <button class="bg-teal-100 hover:bg-teal-200 text-teal-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-cart')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">قیمت</th>
          <th class="p-2 text-right font-bold">تعداد</th>
          <th class="p-2 text-right font-bold">قیمت کل</th>
          <th class="p-2 text-right font-bold">عملیات</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in cartItems" :key="item.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ item.name }}</td>
          <td class="p-2 text-right">{{ item.price }} تومان</td>
          <td class="p-2 text-right">{{ item.quantity }}</td>
          <td class="p-2 text-right">{{ item.total }} تومان</td>
          <td class="p-2 text-right">
            <button v-if="canDeleteItems" class="text-red-600 hover:text-red-700" @click="$emit('remove-item', item)">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
              </svg>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="mt-4 p-3 bg-teal-50 rounded-lg">
      <div class="text-teal-800 font-bold">جمع کل: <span class="text-2xl">{{ total }} تومان</span></div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { computed } from 'vue';
import { useAuth } from '~/composables/useAuth';
import type { User } from '~/types/user';

// استفاده از کامپوزابل احراز هویت
const { hasPermission } = useAuth();

// بررسی دسترسی در client-side
const canDeleteItems = computed(() => {
  if (process.server) return true
  return hasPermission && hasPermission('user.delete')
})


defineProps<{ user: User }>();
defineEmits(['view-all-cart', 'remove-item']);
// Mock data for cart
const total = '۳٬۵۰۰٬۰۰۰';
const cartItems = [
  { id: 1, name: 'لپ‌تاپ لنوو', price: '۲٬۰۰۰٬۰۰۰', quantity: 1, total: '۲٬۰۰۰٬۰۰۰' },
  { id: 2, name: 'ماوس گیمینگ', price: '۵۰۰٬۰۰۰', quantity: 2, total: '۱٬۰۰۰٬۰۰۰' },
  { id: 3, name: 'کیبورد مکانیکال', price: '۵۰۰٬۰۰۰', quantity: 1, total: '۵۰۰٬۰۰۰' },
];
</script> 
