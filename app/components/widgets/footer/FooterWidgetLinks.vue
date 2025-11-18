<template>
  <div class="footer-widget footer-widget--links">
    <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
    <div class="footer-widget__grid" :style="{ gridTemplateColumns: viewModel.template }">
      <div v-for="(column, columnIndex) in viewModel.columns" :key="columnIndex" class="footer-widget__column">
        <ul class="footer-widget__list">
          <li v-for="(item, index) in column" :key="index" class="footer-widget__item">
            <NuxtLink :to="item.href" class="footer-widget__link">{{ item.label }}</NuxtLink>
          </li>
        </ul>
      </div>
    </div>
    <p v-if="!viewModel.columns.length" class="footer-widget__empty">No links available.</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type LinkItem = {
  label: string
  href: string
}

const props = withDefaults(defineProps<{
  title?: string
  items?: LinkItem[]
  columns?: number
}>(), {
  title: 'Quick links',
  columns: 2
})

const normalizeItems = (items?: LinkItem[]) => {
  if (!Array.isArray(items)) {
    return []
  }
  return items
    .filter((item) => item && typeof item.label === 'string' && typeof item.href === 'string')
    .map((item) => ({
      label: item.label.trim() || 'Link',
      href: item.href || '#'
    }))
}

const splitIntoColumns = (items: LinkItem[], columnCount: number) => {
  if (columnCount <= 1) {
    return [items]
  }
  const columns: LinkItem[][] = Array.from({ length: columnCount }, () => [])
  items.forEach((item, index) => {
    const target = index % columnCount
    columns[target].push(item)
  })
  return columns.filter((column) => column.length)
}

const viewModel = computed(() => {
  const columnCount = Math.max(1, props.columns || 1)
  const items = normalizeItems(props.items)
  return {
    title: props.title?.trim() || '',
    columns: splitIntoColumns(items, columnCount),
    template: `repeat(${Math.max(1, columnCount)}, minmax(0, 1fr))`
  }
})
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

.footer-widget__grid {
  display: grid;
  gap: 0.75rem;
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
}

.footer-widget__link:hover {
  text-decoration: underline;
}

.footer-widget__empty {
  margin: 0;
  font-style: italic;
  color: var(--footer-widget-muted-color, #9ca3af);
}
</style>
