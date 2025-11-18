<template>
  <div class="footer-widget footer-widget--logo">
    <NuxtLink :to="viewModel.link" class="footer-widget__logo-link">
      <img
        v-if="viewModel.logoSrc"
        :src="viewModel.logoSrc"
        :alt="viewModel.altText"
        class="footer-widget__logo"
        loading="lazy"
      />
      <span v-else class="footer-widget__fallback">{{ viewModel.fallbackText }}</span>
    </NuxtLink>
    <p v-if="viewModel.tagline" class="footer-widget__tagline">{{ viewModel.tagline }}</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  src?: string
  alt?: string
  fallbackText?: string
  tagline?: string
  link?: string
}>(), {
  src: '',
  alt: 'Logo',
  fallbackText: 'Your Brand',
  tagline: '',
  link: '/'
})

const viewModel = computed(() => ({
  link: props.link || '/',
  logoSrc: props.src?.trim() || '',
  altText: props.alt?.trim() || 'Logo',
  fallbackText: props.fallbackText?.trim() || 'Your Brand',
  tagline: props.tagline?.trim() || ''
}))
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.5rem;
}

.footer-widget__logo-link {
  display: inline-flex;
  align-items: center;
  text-decoration: none;
}

.footer-widget__logo {
  max-width: 160px;
  height: auto;
}

.footer-widget__fallback {
  font-weight: 600;
  font-size: 1.125rem;
  color: var(--footer-widget-logo-color, #111827);
}

.footer-widget__tagline {
  margin: 0;
  font-size: 0.875rem;
  color: var(--footer-widget-tagline-color, #6b7280);
}
</style>
