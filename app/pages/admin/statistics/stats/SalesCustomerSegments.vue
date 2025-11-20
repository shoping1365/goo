<template>
  <div class="bg-white rounded-xl shadow p-6">
    <div class="font-bold text-gray-700 mb-4">بخش‌بندی مشتریان</div>
    <div class="space-y-3">
      <div class="flex justify-between items-center">
        <span class="text-gray-600">مشتریان جدید</span>
        <div class="flex items-center">
          <div class="w-20 h-2 bg-gray-200 rounded-full ml-2">
            <div class="w-3/5 h-full bg-blue-500 rounded-full"></div>
          </div>
          <span class="text-sm font-bold">60%</span>
        </div>
      </div>
      <div class="flex justify-between items-center">
        <span class="text-gray-600">مشتریان بازگشتی</span>
        <div class="flex items-center">
          <div class="w-20 h-2 bg-gray-200 rounded-full ml-2">
            <div class="w-2/5 h-full bg-green-500 rounded-full"></div>
          </div>
          <span class="text-sm font-bold">35%</span>
        </div>
      </div>
      <div class="flex justify-between items-center">
        <span class="text-gray-600">مشتریان VIP</span>
        <div class="flex items-center">
          <div class="w-20 h-2 bg-gray-200 rounded-full ml-2">
            <div class="w-1/10 h-full bg-purple-500 rounded-full"></div>
          </div>
          <span class="text-sm font-bold">5%</span>
        </div>
      </div>
      <div class="mt-4 p-3 bg-purple-50 rounded-lg">
        <div class="text-sm font-semibold text-purple-700">درآمد مشتریان VIP</div>
        <div class="text-lg font-bold text-purple-600">45% از کل درآمد</div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ segments: unknown[] }>()
</script> 
