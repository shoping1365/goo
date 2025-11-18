<template>
  <div v-if="src" class="w-full flex justify-end">
    <img :src="src" :alt="alt" class="rounded-lg shadow-md object-cover max-h-80 border border-gray-100" />
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  src: String,
  alt: String
})
</script> 