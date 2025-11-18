<template>
  <div v-if="compressionQueue.total" class="bg-white rounded-xl shadow border border-gray-200 p-6 mb-6">
    <div class="flex items-center justify-between mb-2">
      <span class="font-bold text-gray-700">وضعیت صف فشرده‌سازی</span>
      <span class="text-sm text-gray-500">{{ compressionQueue.completed }}/{{ compressionQueue.total }}</span>
    </div>
    <div class="w-full bg-gray-100 rounded-full h-4 overflow-hidden">
      <div class="bg-blue-500 h-4 transition-all" :style="{ width: compressionQueue.progress + '%' }"></div>
    </div>
    <div class="text-xs text-gray-600 mt-2 flex gap-6">
      <span>در حال پردازش: {{ compressionQueue.inProgress }}</span>
      <span>موفق: {{ compressionQueue.completed }}</span>
      <span>خطا: {{ compressionQueue.errors }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageCompression } from '~/composables/useImageCompression';
const { compressionQueue } = useImageCompression()
</script> 
