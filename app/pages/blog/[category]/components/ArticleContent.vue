<template>
  <!-- eslint-disable-next-line vue/no-v-html -->
  <div class="prose prose-lg max-w-none text-right" v-html="sanitizedContent"></div>
</template>

<script setup lang="ts">
import DOMPurify from 'dompurify'
import { computed } from 'vue'

const props = defineProps({
  content: String
})

const sanitizedContent = computed(() => DOMPurify.sanitize(props.content || ''))
</script> 