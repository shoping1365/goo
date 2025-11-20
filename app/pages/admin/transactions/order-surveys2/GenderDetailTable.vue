<template>
  <table class="min-w-full divide-y divide-gray-200">
    <thead class="bg-gray-50">
    <tr>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">جنسیت</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">کل سفارشات</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">SMS ارسال شده</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">پاسخ‌ها</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">نرخ پاسخ‌دهی</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">میانگین امتیاز</th>
    </tr>
    </thead>
    <tbody class="bg-white divide-y divide-gray-200">
    <tr v-for="row in data" :key="row.gender">
      <td class="px-6 py-4 whitespace-nowrap">{{ row.gender }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.totalOrders }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.sentSMS }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.responses }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.responseRate }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.averageRating }}</td>
    </tr>
    </tbody>
  </table>
</template>

<script setup lang="ts">
interface GenderDetail {
  gender: string;
  totalOrders: number;
  sentSMS: number;
  responses: number;
  responseRate: string;
  averageRating: number;
}
defineProps<{ data: GenderDetail[] }>()
</script>
