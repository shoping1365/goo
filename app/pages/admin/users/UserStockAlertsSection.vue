<template>
  <div class="bg-white rounded-2xl shadow-lg p-6 mt-2">
    <div class="flex justify-between items-center mb-2">
      <div class="flex items-center gap-2 font-bold text-blue-700">
        <svg class="w-5 h-5 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/></svg>
        اطلاع از موجودی
      </div>
      <button class="bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-stock-alerts')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm mt-2">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">تاریخ درخواست</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in alerts" :key="item.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ item.name }}</td>
          <td class="p-2 text-right">{{ item.date }}</td>
          <td class="p-2 text-right">
            <span :class="item.status === 'ارسال شد' ? 'text-green-600' : 'text-gray-500'">{{ item.status }}</span>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="!alerts.length" class="text-center text-gray-400 py-6">درخواستی ثبت نشده است.</div>
  </div>
</template>
<script setup lang="ts">
defineEmits(['view-all-stock-alerts']);
const alerts = [
  { id: 1, name: 'تلویزیون ال‌جی', date: '1402/03/08', status: 'ارسال شد' },
  { id: 2, name: 'دوربین عکاسی', date: '1402/03/07', status: 'در انتظار' },
];
</script> 
