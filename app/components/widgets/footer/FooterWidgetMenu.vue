<template>
  <nav class="footer-widget footer-widget--menu" aria-label="Footer menu">
    <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
    <ul class="footer-widget__list">
      <li v-for="(item, index) in viewModel.items" :key="index" class="footer-widget__item">
        <NuxtLink :to="item.href" class="footer-widget__link">{{ item.label }}</NuxtLink>
      </li>
      <li v-if="!viewModel.items.length" class="footer-widget__item footer-widget__item--empty">
        {{ viewModel.emptyText }}
      </li>
    </ul>
  </nav>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type MenuItem = {
  label: string
  href: string
}

const props = withDefaults(defineProps<{
  title?: string
  items?: MenuItem[]
  emptyText?: string
}>(), {
  title: '',
  emptyText: 'No links defined.'
})

const normalizeItems = (items?: MenuItem[]) => {
  if (!Array.isArray(items)) {
    return []
  }
  return items
    .filter((item) => item && typeof item.label === 'string' && typeof item.href === 'string')
    .map((item) => ({
      label: item.label.trim() || 'Untitled',
      href: item.href || '#'
    }))
}

const viewModel = computed(() => ({
  title: props.title?.trim() || '',
  items: normalizeItems(props.items),
  emptyText: props.emptyText || 'No links defined.'
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

.footer-widget__list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.footer-widget__link {
  text-decoration: none;
  color: var(--footer-widget-link-color, #2563eb);
  font-size: 0.9rem;
}

.footer-widget__link:hover {
  text-decoration: underline;
}

.footer-widget__item--empty {
  font-style: italic;
  color: var(--footer-widget-muted-color, #9ca3af);
}
</style>
