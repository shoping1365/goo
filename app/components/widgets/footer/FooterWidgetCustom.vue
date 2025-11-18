<template>
  <div class="footer-widget footer-widget--custom">
    <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
    <div v-if="viewModel.html" class="footer-widget__content" v-html="viewModel.html"></div>
    <slot v-else />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  title?: string
  html?: string
}>(), {
  title: '',
  html: ''
})

const viewModel = computed(() => ({
  title: props.title?.trim() || '',
  html: props.html || ''
}))
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-widget__title {
  margin: 0;
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.footer-widget__content :deep(*) {
  color: inherit;
  font-size: 0.9rem;
  line-height: 1.6;
}
</style>
