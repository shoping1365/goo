<template>
  <div class="footer-widget footer-widget--social">
    <h4 v-if="props.title" class="footer-widget__title">{{ props.title }}</h4>
    <div class="social-icons">
      <a 
        v-for="(item, index) in visibleLinks" 
        :key="index" 
        :href="item.href" 
        class="social-icon-link" 
        target="_blank" 
        rel="noopener noreferrer"
        :title="item.label"
      >
        {{ item.label }}
      </a>
      <p v-if="!visibleLinks.length" class="empty-text">
        {{ props.emptyText }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  title?: string
  links?: Array<any>
  socials?: Array<{
    platform?: string
    label?: string
    title?: string
    href?: string
    url?: string
    enabled?: boolean
  }>
  emptyText?: string
}>(), {
  title: 'شبکه‌های اجتماعی',
  emptyText: 'شبکه‌های اجتماعی تعریف نشده است'
})

// نقشه ایکون‌ها برای هر پلتفرم
const platformIcons: Record<string, string> = {
  instagram: 'i-mdi-instagram',
  telegram: 'i-mdi-telegram',
  whatsapp: 'i-mdi-whatsapp',
  twitter: 'i-mdi-twitter',
  linkedin: 'i-mdi-linkedin',
  facebook: 'i-mdi-facebook',
  youtube: 'i-mdi-youtube',
  aparat: 'i-mdi-play-box',
  rubika: 'i-mdi-message',
  eitaa: 'i-mdi-message-text',
  bale: 'i-mdi-message-reply',
  tiktok: 'i-mdi-music-note',
  pinterest: 'i-mdi-pinterest',
  discord: 'i-mdi-discord',
  github: 'i-mdi-github'
}

const visibleLinks = computed(() => {
  const source = props.socials || props.links || []
  console.log('⚡️ Computing visibleLinks from:', JSON.parse(JSON.stringify(source)))

  if (!Array.isArray(source) || source.length === 0) {
    console.log('⚠️ Source is empty or not an array.')
    return []
  }

  const result = source
    .filter(item => {
      const isEnabled = item && (item.enabled !== false)
      console.log(`  - Filtering ${item.platform}: enabled=${isEnabled}`)
      return isEnabled
    })
    .map(item => {
      const platform = (item.platform || '').toLowerCase()
      const href = (item.href || item.url || '#').trim()
      const label = (item.label || item.title || item.platform || 'Link').trim()
      
      const mappedItem = {
        platform,
        label,
        href: href || '#',
        iconClass: platformIcons[platform] || 'i-mdi-link'
      }
      console.log(`  - Mapping ${item.platform} to:`, mappedItem)
      return mappedItem
    })
  
  console.log('✅ Computed visibleLinks result:', JSON.parse(JSON.stringify(result)))
  return result
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

.social-icons {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  align-items: center;
}

.social-icon-link {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: var(--footer-social-bg, #f3f4f6);
  color: var(--footer-social-color, #374151);
  text-decoration: none;
  transition: all 0.3s ease;
}

.social-icon-link:hover {
  background: var(--footer-social-hover-bg, #e5e7eb);
  color: var(--footer-social-hover-color, #1f2937);
  transform: translateY(-2px);
}

.social-icon {
  font-size: 20px;
}

.empty-text {
  font-style: italic;
  color: var(--footer-widget-muted-color, #9ca3af);
  margin: 0;
  font-size: 0.875rem;
}
</style>
