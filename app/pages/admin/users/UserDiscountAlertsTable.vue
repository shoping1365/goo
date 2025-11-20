<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-rose-700 mb-2 flex justify-between items-center">
      <span>هشدارهای تخفیف</span>
      <button class="bg-rose-100 hover:bg-rose-200 text-rose-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-discount-alerts')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">قیمت اصلی</th>
          <th class="p-2 text-right font-bold">درصد تخفیف</th>
          <th class="p-2 text-right font-bold">تاریخ ثبت</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="alert in discountAlerts" :key="alert.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ alert.name }}</td>
          <td class="p-2 text-right">{{ alert.originalPrice }} تومان</td>
          <td class="p-2 text-right text-red-600">{{ alert.discountPercent }}%</td>
          <td class="p-2 text-right">{{ alert.registeredDate }}</td>
          <td class="p-2 text-right">
            <span :class="alert.isActive ? 'text-green-600' : 'text-gray-600'">{{ alert.isActive ? 'فعال' : 'غیرفعال' }}</span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
import type { User } from '~/types/user';

defineProps<{ user: User }>();
defineEmits(['view-all-discount-alerts']);
// Mock data for discount alerts
const discountAlerts = [
  { id: 1, name: 'لپ‌تاپ لنوو', originalPrice: '۳۰٬۰۰۰٬۰۰۰', discountPercent: 20, registeredDate: '1402/03/10', isActive: true },
  { id: 2, name: 'گوشی سامسونگ', originalPrice: '۲۰٬۰۰۰٬۰۰۰', discountPercent: 15, registeredDate: '1402/03/09', isActive: false },
  { id: 3, name: 'تبلت اپل', originalPrice: '۲۵٬۰۰۰٬۰۰۰', discountPercent: 10, registeredDate: '1402/03/08', isActive: true },
];
</script> 
