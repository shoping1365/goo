<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex items-center">
      <div class="flex-shrink-0">
        <div :class="iconBgClass" class="w-8 h-8 rounded-md flex items-center justify-center">
          <slot name="icon"></slot>
        </div>
      </div>
      <div class="mr-4">
        <p class="text-sm font-medium text-gray-500">{{ title }}</p>
        <p class="text-2xl font-semibold text-gray-900">{{ value }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string
  value: string | number
  iconBgClass?: string
}

withDefaults(defineProps<Props>(), {
  iconBgClass: 'bg-blue-100'
})
</script> 