<template>
  <div class="bg-white rounded-2xl shadow-lg p-6 mt-2">
    <div class="font-bold text-orange-700 mb-2 flex items-center gap-2">
      <svg class="w-5 h-5 text-orange-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 1.343-3 3s1.343 3 3 3 3-1.343 3-3-1.343-3-3-3zm0 13c-4.418 0-8-3.582-8-8s3.582-8 8-8 8 3.582 8 8-3.582 8-8 8z"/></svg>
      اطلاع از تخفیف
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
const props = defineProps<{ user?: any }>();
const alerts = [
  { id: 1, name: 'کنسول بازی', date: '1402/03/05', status: 'ارسال شد' },
  { id: 2, name: 'کفش ورزشی', date: '1402/03/04', status: 'در انتظار' },
];
</script> 
