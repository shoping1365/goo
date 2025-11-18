<template>
  <!-- سبد خرید با نشانگر تعداد -->
  <div :style="widgetStyle" class="cart-widget-container header-cart-widget">
    <NuxtLink to="/cart" class="relative flex items-center text-xl w-full h-full" :class="getLinkAlignment()">
      🛒
      <span v-if="count && count > 0" class="absolute -top-1 -right-2 text-xs bg-red-500 text-white rounded-full px-1">{{ count }}</span>
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  count?: number
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  count: 0,
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

// تابع تعیین کلاس چینش برای لینک
function getLinkAlignment(): string {
  switch (props.align) {
    case 'left':
      return 'justify-start'
    case 'center':
      return 'justify-center'
    case 'right':
      return 'justify-end'
    default:
      return 'justify-center'
  }
}
</script>

<style scoped>
.header-cart-widget {
  margin: 0;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.header-cart-widget a {
  margin: 0;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
