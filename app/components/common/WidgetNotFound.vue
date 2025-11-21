<template>
  <div class="text-red-500 p-6 border border-red-300 rounded">
    <h3 class="font-bold">Widget نادیده</h3>
    <p>نوع ویجت "{{ widgetType }}" پیدا نشد.</p>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  widgetType: string
}>()
</script>

