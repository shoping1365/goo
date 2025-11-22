<template>
  <div class="bg-white rounded-xl shadow p-6 h-80 flex flex-col">
    <div class="font-bold text-gray-700 mb-2">نقشه بازدیدها</div>
    <div class="flex-1 flex items-center justify-center">
      <!-- Placeholder for map -->
      <div class="w-full h-64 bg-gradient-to-r from-green-100 to-blue-100 rounded-lg flex items-center justify-center text-green-400 text-2xl">
        [نقشه بازدیدها]
      </div>
    </div>
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
defineProps<{ locations: Record<string, unknown>[] }>()
</script> 
