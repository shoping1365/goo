<template>
  <div v-if="hasDescription" class="footer-widget footer-widget--about" :style="containerStyle">
    <div class="footer-widget__inner" :style="innerStyle">
      <p class="footer-widget__description">{{ resolvedDescription }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  title?: string
  description?: string
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}>(), {
  title: '',
  description: '',
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

const resolvedDescription = computed(() => props.description?.trim() || '')
const hasDescription = computed(() => Boolean(resolvedDescription.value))

const containerStyle = computed(() => ({
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

function getJustifyContent(align: string | undefined): string {
  switch (align) {
    case 'left':
      return 'flex-start'
    case 'right':
      return 'flex-end'
    default:
      return 'center'
  }
}

function getTextAlign(align: string | undefined): 'left' | 'center' | 'right' {
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
.footer-widget {
  display: flex;
  align-items: center;
}

.footer-widget__inner {
  width: 100%;
}

.footer-widget__description {
  margin: 0;
  font-size: 0.95rem;
  line-height: 1.9;
  color: var(--footer-widget-text-color, #374151);
}

</style>
