<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-pink-700 mb-2 flex justify-between items-center">
      <span>لیست علاقه‌مندی‌ها</span>
      <button class="bg-pink-100 hover:bg-pink-200 text-pink-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-wishlist')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">قیمت</th>
          <th class="p-2 text-right font-bold">تاریخ افزودن</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
          <th class="p-2 text-right font-bold">عملیات</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in wishlistItems" :key="item.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ item.name }}</td>
          <td class="p-2 text-right">{{ item.price }} تومان</td>
          <td class="p-2 text-right">{{ item.addedDate }}</td>
          <td class="p-2 text-right">
            <span :class="item.inStock ? 'text-green-600' : 'text-red-600'">{{ item.inStock ? 'موجود' : 'ناموجود' }}</span>
          </td>
          <td class="p-2 text-right">
            <button class="text-pink-600 hover:text-pink-700" @click="$emit('add-to-cart', item)">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ user: any }>();
// Mock data for wishlist
const wishlistItems = [
  { id: 1, name: 'لپ‌تاپ اپل', price: '۵۰٬۰۰۰٬۰۰۰', addedDate: '1402/03/10', inStock: true },
  { id: 2, name: 'گوشی سامسونگ', price: '۱۵٬۰۰۰٬۰۰۰', addedDate: '1402/03/09', inStock: false },
  { id: 3, name: 'تبلت لنوو', price: '۸٬۰۰۰٬۰۰۰', addedDate: '1402/03/08', inStock: true },
];
</script> 
