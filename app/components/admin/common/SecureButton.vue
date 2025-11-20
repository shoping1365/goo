<template>
  <button 
    v-if="hasPermission(permission)"
    :class="buttonClasses"
    :disabled="disabled"
    :type="type"
    :title="title"
    @click="$emit('click', $event)"
  >
    <slot />
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuthState } from '~/composables/useAuthState'

// استفاده از کامپوزابل احراز هویت
const { hasPermission } = useAuthState()

// تعریف props
interface Props {
  permission: string
  variant?: 'primary' | 'secondary' | 'success' | 'danger' | 'warning' | 'info' | 'light' | 'dark'
  size?: 'sm' | 'md' | 'lg'
  disabled?: boolean
  type?: 'button' | 'submit' | 'reset'
  title?: string
  customClass?: string
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  disabled: false,
  type: 'button',
  title: '',
  customClass: ''
})

// تعریف emits
defineEmits<{
  click: [event: MouseEvent]
}>()

// کلاس‌های دکمه بر اساس variant
const variantClasses = {
  primary: 'bg-blue-600 hover:bg-blue-700 text-white',
  secondary: 'bg-gray-600 hover:bg-gray-700 text-white',
  success: 'bg-green-600 hover:bg-green-700 text-white',
  danger: 'bg-red-600 hover:bg-red-700 text-white',
  warning: 'bg-yellow-600 hover:bg-yellow-700 text-white',
  info: 'bg-cyan-600 hover:bg-cyan-700 text-white',
  light: 'bg-gray-100 hover:bg-gray-200 text-gray-800',
  dark: 'bg-gray-800 hover:bg-gray-900 text-white'
}

// کلاس‌های دکمه بر اساس size
const sizeClasses = {
  sm: 'px-2 py-1 text-xs',
  md: 'px-4 py-2 text-sm',
  lg: 'px-6 py-3 text-base'
}

// ترکیب کلاس‌ها
const buttonClasses = computed(() => {
  const baseClasses = 'inline-flex items-center justify-center font-medium rounded-lg transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed'
  const variantClass = variantClasses[props.variant]
  const sizeClass = sizeClasses[props.size]
  
  return `${baseClasses} ${variantClass} ${sizeClass} ${props.customClass}`
})
</script> 