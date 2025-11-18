<template>
  <component :is="tag" v-bind="linkProps" class="footer-widget footer-widget--image">
    <img :src="viewModel.src" :alt="viewModel.alt" class="footer-widget__image" loading="lazy" />
    <p v-if="viewModel.caption" class="footer-widget__caption">{{ viewModel.caption }}</p>
  </component>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  src?: string
  alt?: string
  caption?: string
  href?: string
}>(), {
  src: '/placeholder.svg',
  alt: 'Footer image',
  caption: '',
  href: ''
})

const viewModel = computed(() => ({
  src: props.src || '/placeholder.svg',
  alt: props.alt || 'Footer image',
  caption: props.caption || '',
  href: props.href?.trim() || ''
}))

const tag = computed(() => (viewModel.value.href ? 'NuxtLink' : 'div'))
const linkProps = computed(() => (viewModel.value.href ? { to: viewModel.value.href } : {}))
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  text-decoration: none;
  color: inherit;
}

.footer-widget__image {
  width: 100%;
  height: auto;
  border-radius: 0.75rem;
  object-fit: cover;
}

.footer-widget__caption {
  margin: 0;
  font-size: 0.8rem;
  color: var(--footer-widget-muted-color, #6b7280);
}
</style>
