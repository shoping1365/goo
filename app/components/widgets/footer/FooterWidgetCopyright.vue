<template>
  <div class="footer-widget footer-widget--copyright">
    <div class="footer-widget__text" v-html="viewModel.html"></div>
    <ul v-if="viewModel.links.length" class="footer-widget__links">
      <li v-for="(link, index) in viewModel.links" :key="index" class="footer-widget__link-item">
        <NuxtLink :to="link.href" class="footer-widget__link">{{ link.label }}</NuxtLink>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type FooterLink = {
  label: string
  href: string
}

const props = withDefaults(defineProps<{
  year?: number
  company?: string
  text?: string
  links?: FooterLink[]
}>(), {
  year: undefined,
  company: undefined,
  text: '',
  links: () => []
})

const normalizeLinks = (links?: FooterLink[]) => {
  if (!Array.isArray(links)) {
    return []
  }
  return links
    .filter((link) => link && typeof link.label === 'string' && typeof link.href === 'string')
    .map((link) => ({
      label: link.label.trim() || 'Link',
      href: link.href || '#'
    }))
}

const viewModel = computed(() => {
  if (!props.text || !props.text.trim().length) {
    return {
      html: '',
      links: []
    }
  }

  const html = props.text.trim()
  return {
    html,
    links: normalizeLinks(props.links)
  }
})
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  font-size: 0.85rem;
  color: var(--footer-widget-muted-color, #6b7280);
}

.footer-widget__links {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  list-style: none;
  padding: 0;
  margin: 0;
}

.footer-widget__link {
  color: var(--footer-widget-link-color, #2563eb);
  text-decoration: none;
}

.footer-widget__link:hover {
  text-decoration: underline;
}
</style>
