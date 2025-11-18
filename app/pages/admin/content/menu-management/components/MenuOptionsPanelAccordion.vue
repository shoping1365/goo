<template>
  <div class="border-b">
    <button type="button" class="flex items-center justify-between w-full px-4 py-2 text-xs font-bold text-gray-700 focus:outline-none hover:bg-gray-50" @click="toggle">
      <span>{{ title }}</span>
      <span>{{ isOpen ? '▲' : '▼' }}</span>
    </button>
    <div v-if="isOpen" class="px-4 pb-3">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps({
  title: String,
  open: Boolean
})

const isOpen = ref(props.open)

watch(() => props.open, (val) => {
  isOpen.value = val
})

function toggle() {
  isOpen.value = !isOpen.value
}
</script> 
