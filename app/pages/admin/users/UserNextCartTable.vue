<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-violet-700 mb-2 flex justify-between items-center">
      <span>سبد خرید بعدی</span>
      <button class="bg-violet-100 hover:bg-violet-200 text-violet-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-next-cart')">
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
        <tr v-for="item in nextCartItems" :key="item.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ item.name }}</td>
          <td class="p-2 text-right">{{ item.price }} تومان</td>
          <td class="p-2 text-right">{{ item.quantity }}</td>
          <td class="p-2 text-right">{{ item.total }} تومان</td>
          <td class="p-2 text-right">
            <button class="text-violet-600 hover:text-violet-700" @click="$emit('move-to-cart', item)">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
              </svg>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="mt-4 p-3 bg-violet-50 rounded-lg">
      <div class="text-violet-800 font-bold">جمع کل: <span class="text-2xl">{{ total }} تومان</span></div>
    </div>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ user: any }>();
// Mock data for next cart
const total = '۲٬۰۰۰٬۰۰۰';
const nextCartItems = [
  { id: 1, name: 'هندزفری بلوتوث', price: '۱٬۰۰۰٬۰۰۰', quantity: 1, total: '۱٬۰۰۰٬۰۰۰' },
  { id: 2, name: 'پاوربانک', price: '۵۰۰٬۰۰۰', quantity: 2, total: '۱٬۰۰۰٬۰۰۰' },
];
</script> 
