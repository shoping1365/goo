<template>
  <i :class="iconClass" :style="iconStyle"></i>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  name: string
  size?: string | number
  color?: string
}

const props = withDefaults(defineProps<Props>(), {
  size: '24px',
  color: 'currentColor'
})

const iconClass = computed(() => {
  // اگر name با i- شروع میشه، مستقیم استفاده کن (مثلاً i-heroicons-home)
  if (props.name.startsWith('i-')) {
    return props.name
  }
  
  // اگر نه، فرض کن از مجموعه heroicons هست
  return `i-heroicons-${props.name}`
})

const iconStyle = computed(() => ({
  fontSize: typeof props.size === 'number' ? `${props.size}px` : props.size,
  color: props.color,
  display: 'inline-block',
  width: '1em',
  height: '1em'
}))
</script>

<style scoped>
i {
  vertical-align: middle;
}
</style>
