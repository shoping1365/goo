<template>
  <!-- انتخاب زبان -->
  <div :style="widgetStyle" class="language-widget-container">
    <select class="border rounded px-2 py-1 text-sm">
      <option value="fa">فارسی</option>
      <option value="en">English</option>
    </select>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
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
