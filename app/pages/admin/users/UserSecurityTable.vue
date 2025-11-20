<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-red-700 mb-2 flex justify-between items-center">
      <span>امنیت کاربر</span>
      <button class="bg-red-100 hover:bg-red-200 text-red-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-security')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">عملیات</th>
          <th class="p-2 text-right font-bold">تاریخ</th>
          <th class="p-2 text-right font-bold">IP</th>
          <th class="p-2 text-right font-bold">دستگاه</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="log in securityLogs" :key="log.date + log.ip" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ log.action }}</td>
          <td class="p-2 text-right">{{ log.date }}</td>
          <td class="p-2 text-right">{{ log.ip }}</td>
          <td class="p-2 text-right">{{ log.device }}</td>
          <td class="p-2 text-right">
            <span :class="log.status === 'موفق' ? 'text-green-600' : 'text-red-600'">{{ log.status }}</span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
import type { User } from '~/types/user';

defineProps<{ user: User }>();
defineEmits(['view-all-security']);
// Mock data for security logs
const securityLogs = [
  { action: 'ورود به سیستم', date: '1402/03/10 09:12', ip: '192.168.1.10', device: 'موبایل', status: 'موفق' },
  { action: 'تغییر رمز عبور', date: '1402/03/09 18:22', ip: '192.168.1.11', device: 'دسکتاپ', status: 'موفق' },
  { action: 'تلاش ورود ناموفق', date: '1402/03/09 17:45', ip: '192.168.1.12', device: 'موبایل', status: 'ناموفق' },
];
</script> 
