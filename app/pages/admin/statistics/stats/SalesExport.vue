<template>
  <div class="bg-white rounded-xl shadow p-6 flex gap-6 items-center">
    <div class="font-bold text-gray-700">خروجی گرفتن از داده‌های فروش</div>
    <button class="px-4 py-2 bg-green-500 text-white rounded-lg shadow hover:bg-green-600 transition">Excel</button>
    <button class="px-4 py-2 bg-blue-500 text-white rounded-lg shadow hover:bg-blue-600 transition">CSV</button>
    <button class="px-4 py-2 bg-red-500 text-white rounded-lg shadow hover:bg-red-600 transition">PDF</button>
    <button class="px-4 py-2 bg-purple-500 text-white rounded-lg shadow hover:bg-purple-600 transition">گزارش مالی</button>
  </div>
</template>
<script setup lang="ts">
defineProps<{ sales: unknown[] }>()
</script> 
