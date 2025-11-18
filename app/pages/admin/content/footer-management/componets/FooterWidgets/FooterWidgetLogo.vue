<template>
  <div class="footer-logo-widget" :style="widgetStyle">
    <NuxtLink 
      :to="logoLink || '/'"
      class="flex items-center space-x-3 space-x-reverse hover:opacity-80 transition-opacity"
    >
      <!-- لوگوی فوتر -->
      <div class="flex-shrink-0">
        <img
          v-if="logoUrl"
          :src="logoUrl"
          :alt="logoAlt || 'لوگوی سایت'"
          class="h-12 w-auto object-contain"
        />
        <div
          v-else
          class="h-12 w-32 bg-white bg-opacity-20 rounded-lg flex items-center justify-center border border-white border-opacity-30"
        >
          <span class="text-white text-lg font-bold">{{ logoAlt || 'LOGO' }}</span>
        </div>
      </div>
      
      <!-- نام سایت -->
      <div v-if="showSiteName" class="text-white">
        <h3 class="text-lg font-bold">{{ siteName || 'نام سایت' }}</h3>
        <p v-if="siteDescription" class="text-sm opacity-80">{{ siteDescription }}</p>
      </div>
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  logoUrl?: string
  logoAlt?: string
  logoLink?: string
  siteName?: string
  siteDescription?: string
  showSiteName?: boolean
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  logoUrl: '',
  logoAlt: 'لوگوی سایت',
  logoLink: '/',
  siteName: 'نام سایت',
  siteDescription: '',
  showSiteName: true,
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
.footer-logo-widget {
  transition: all 0.2s ease;
}

.footer-logo-widget:hover {
  transform: scale(1.02);
}
</style>
