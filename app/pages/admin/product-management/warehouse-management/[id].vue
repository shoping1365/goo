<template>
  <div class="min-h-screen bg-gray-50">
    <WarehouseDetail />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import WarehouseDetail from './components/WarehouseDetail.vue';

// تعریف layout صفحه
definePageMeta({
  layout: 'admin-main'
})
</script>
