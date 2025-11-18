<template>
  <div class="brands-slider">
    <div class="container mx-auto px-4 py-6">
      <div class="bg-gray-200 rounded-lg flex items-center justify-center h-64">
        <p class="text-gray-500">کامپوننت اسلایدر برندها - در حال توسعه</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Widget } from '~/types/widget'

interface Props {
  widget: Widget
}

const props = defineProps<Props>()
</script> 