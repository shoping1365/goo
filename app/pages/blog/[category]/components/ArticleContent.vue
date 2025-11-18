<template>
  <div class="prose prose-lg max-w-none text-right" v-html="content"></div>
</template>

<script setup lang="ts">
const props = defineProps({
  content: String
})
</script> 