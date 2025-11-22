<template>
  <div class="bg-blue-100 p-6 rounded-lg">
    <h2 class="text-lg font-bold text-blue-800">تست کامپوننت</h2>
    <p class="text-blue-600">این یک کامپوننت تست است</p>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})
// Test component
</script> 
