<template>
  <button type="button" class="footer-widget footer-widget--back-to-top" @click="handleClick">
    {{ viewModel.label }}
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  label?: string
  behavior?: ScrollBehavior
}>(), {
  label: 'Back to top',
  behavior: 'smooth'
})

const viewModel = computed(() => ({
  label: props.label?.trim() || 'Back to top',
  behavior: props.behavior || 'smooth'
}))

const handleClick = () => {
  if (typeof window === 'undefined') {
    return
  }
  window.scrollTo({ top: 0, behavior: viewModel.value.behavior })
}
</script>

<style scoped>
.footer-widget {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.6rem 1.2rem;
  border-radius: 999px;
  border: none;
  background: var(--footer-widget-button-bg, #2563eb);
  color: #fff;
  font-weight: 600;
  cursor: pointer;
}

.footer-widget:hover {
  background: var(--footer-widget-button-bg-hover, #1d4ed8);
}
</style>
