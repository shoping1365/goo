<template>
  <div class="footer-widget footer-widget--notification">
    <h4 class="footer-widget__title">{{ viewModel.title }}</h4>
    <ul class="footer-widget__list">
      <li v-for="(item, index) in viewModel.items" :key="index" class="footer-widget__item">
        <span class="footer-widget__item-text">{{ item.text }}</span>
        <time v-if="item.date" class="footer-widget__item-date">{{ item.date }}</time>
      </li>
      <li v-if="!viewModel.items.length" class="footer-widget__item footer-widget__item--empty">
        {{ viewModel.emptyText }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type NotificationItem = {
  text: string
  date?: string
}

const props = withDefaults(defineProps<{
  title?: string
  items?: NotificationItem[]
  emptyText?: string
}>(), {
  title: 'Latest updates',
  emptyText: 'No notifications yet.'
})

const viewModel = computed(() => ({
  title: props.title || 'Latest updates',
  items: Array.isArray(props.items) ? props.items : [],
  emptyText: props.emptyText || 'No notifications yet.'
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

.footer-widget__item {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  font-size: 0.85rem;
  color: var(--footer-widget-muted-color, #4b5563);
}

.footer-widget__item-text {
  line-height: 1.4;
}

.footer-widget__item-date {
  font-size: 0.75rem;
  color: var(--footer-widget-date-color, #9ca3af);
}

.footer-widget__item--empty {
  font-style: italic;
}
</style>
