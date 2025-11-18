<template>
  <div class="footer-about-widget" :style="widgetStyle">
    <div class="footer-about-widget__inner" :style="innerStyle">
      <p v-if="description" class="footer-widget-text text-sm leading-6">{{ description }}</p>
      <p v-else class="footer-widget-placeholder text-xs">متن درباره ما در تنظیمات سایت خالی است.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  title?: string
  description?: string
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: undefined,
  description: '',
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  display: 'flex',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

const innerStyle = computed(() => {
  const style: Record<string, string> = {
    textAlign: getTextAlign(props.align),
    maxWidth: '100%',
    width: '100%'
  }

  switch (props.align) {
    case 'left':
      style.margin = '0'
      break
    case 'right':
      style.margin = '0 0 0 auto'
      break
    default:
      style.margin = '0 auto'
      break
  }

  return style
})

function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'
    case 'right':
      return 'flex-end'
    default:
      return 'center'
  }
}

function getTextAlign(align: string): 'left' | 'center' | 'right' {
  switch (align) {
    case 'left':
      return 'left'
    case 'right':
      return 'right'
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-about-widget {
  transition: all 0.2s ease;
}

.footer-about-widget__inner {
  width: 100%;
}

.footer-widget-text {
  color: #1f2937;
}

.footer-widget-placeholder {
  color: #9ca3af;
  font-style: italic;
}
</style>
