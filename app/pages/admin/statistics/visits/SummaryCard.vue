<template>
  <div :class="['rounded-xl p-6 shadow bg-white flex flex-col items-center', 'border-t-4', 'border-' + color + '-500']">
    <div class="mb-2">
      <span :class="['mdi', icon, 'text-2xl', 'text-' + color + '-500']"></span>
    </div>
    <div class="text-lg font-bold text-gray-800">{{ value }}</div>
    <div class="text-xs text-gray-500 mt-1">{{ title }}</div>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ title: string, value: any, icon: string, color: string }>()
</script> 
