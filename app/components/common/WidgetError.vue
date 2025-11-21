<template>
  <div class="text-red-500 p-6 border border-red-300 rounded">
    Failed to load widget: {{ widgetType }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  widgetType: string
}>()
</script>

