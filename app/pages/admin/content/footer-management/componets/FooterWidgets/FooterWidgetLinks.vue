<template>
  <div class="footer-links-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- عنوان بخش -->
      <h4 v-if="title" class="text-white font-semibold mb-4 text-lg">{{ title }}</h4>
      
      <!-- لینک‌های مفید -->
      <div class="space-y-2">
        <NuxtLink
          v-for="link in links"
          :key="link.title"
          :to="link.url"
          class="block text-white text-sm hover:text-blue-200 transition-colors duration-200 py-1"
        >
          {{ link.title }}
        </NuxtLink>
      </div>
      
      <!-- پیام خالی -->
      <div v-if="!links || links.length === 0" class="text-white text-opacity-60 text-sm">
        هیچ لینکی تعریف نشده است
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Link {
  title: string
  url: string
}

interface Props {
  title?: string
  links?: Link[]
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: 'لینک‌های مفید',
  links: () => [
    { title: 'درباره ما', url: '/about' },
    { title: 'تماس با ما', url: '/contact' },
    { title: 'قوانین و مقررات', url: '/terms' },
    { title: 'حریم خصوصی', url: '/privacy' }
  ],
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// استایل کامپوننت بر اساس چینش
const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  display: 'flex',
  alignItems: 'center',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'  // در RTL: چپ = flex-start
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'  // در RTL: راست = flex-end
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-links-widget {
  transition: all 0.2s ease;
}

.footer-links-widget a:hover {
  transform: translateX(-2px);
}
</style>
