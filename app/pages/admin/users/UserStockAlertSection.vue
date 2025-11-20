<template>
  <div class="bg-white rounded-2xl shadow-lg p-6 mt-2">
    <div class="font-bold text-blue-700 mb-2 flex items-center gap-2">
      <svg class="w-5 h-5 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V4a2 2 0 10-4 0v1.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/></svg>
      اطلاع از موجودی
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
        <tr v-for="item in alerts" :key="item.id" class="border-b hover:bg-gray-50">
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
import type { User } from '~/types/user';

defineProps<{ user?: User }>();
const alerts = [
  { id: 1, name: 'تلویزیون ال‌جی', date: '1402/03/08', status: 'ارسال شد' },
  { id: 2, name: 'دوربین عکاسی', date: '1402/03/07', status: 'در انتظار' },
];
</script> 
