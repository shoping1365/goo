<template>
  <!-- باکس سفارشی -->
  <div class="p-2 bg-green-50 rounded text-sm">
    📦 {{ text }}
  </div>
</template>

<script setup lang="ts">
// پراپرتی متن نمایشی
defineProps<{ text?: string }>()
</script>
