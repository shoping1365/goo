<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-amber-700 mb-2 flex justify-between items-center">
      <span>هشدارهای موجودی</span>
      <button class="bg-amber-100 hover:bg-amber-200 text-amber-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-stock-alerts')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">قیمت</th>
          <th class="p-2 text-right font-bold">تاریخ ثبت</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
          <th class="p-2 text-right font-bold">عملیات</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="alert in stockAlerts" :key="alert.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ alert.name }}</td>
          <td class="p-2 text-right">{{ alert.price }} تومان</td>
          <td class="p-2 text-right">{{ alert.registeredDate }}</td>
          <td class="p-2 text-right">
            <span :class="alert.isActive ? 'text-green-600' : 'text-gray-600'">{{ alert.isActive ? 'فعال' : 'غیرفعال' }}</span>
          </td>
          <td class="p-2 text-right">
            <button class="text-amber-600 hover:text-amber-700" @click="$emit('toggle-alert', alert)">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/>
              </svg>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
defineEmits(['view-all-stock-alerts', 'toggle-alert']);
// Mock data for stock alerts
const stockAlerts = [
  { id: 1, name: 'لپ‌تاپ لنوو', price: '۲۵٬۰۰۰٬۰۰۰', registeredDate: '1402/03/10', isActive: true },
  { id: 2, name: 'گوشی شیائومی', price: '۱۲٬۰۰۰٬۰۰۰', registeredDate: '1402/03/09', isActive: false },
  { id: 3, name: 'تبلت سامسونگ', price: '۱۰٬۰۰۰٬۰۰۰', registeredDate: '1402/03/08', isActive: true },
];
</script> 
