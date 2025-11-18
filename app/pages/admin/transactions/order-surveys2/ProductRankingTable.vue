<template>
  <table class="min-w-full divide-y divide-gray-200">
    <thead class="bg-gray-50">
    <tr>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">رتبه</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">محصول</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">کل سفارشات</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">میانگین امتیاز</th>
      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">نرخ پاسخ‌دهی</th>
    </tr>
    </thead>
    <tbody class="bg-white divide-y divide-gray-200">
    <tr v-for="row in data" :key="row.rank">
      <td class="px-6 py-4 whitespace-nowrap">{{ row.rank }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.product }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.totalOrders }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.averageRating }}</td>
      <td class="px-6 py-4 whitespace-nowrap">{{ row.responseRate }}</td>
    </tr>
    </tbody>
  </table>
</template>

<script setup lang="ts">
defineProps<{ data: any[] }>()
</script>
