<template>
  <div v-if="!!modelValue" class="fixed inset-0 flex items-center justify-center z-50" @click="emitClose">
    <div class="relative">
      <button @click="emitClose" class="absolute right-0 top-0 -translate-y-1/2 translate-x-1/2 bg-rose-200 hover:bg-rose-300 rounded-full p-2 shadow hover:bg-opacity-100 transition-all z-10">
        <svg class="w-7 h-7 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
      <img
        v-if="image && (image.url || image.src)"
        :src="image.url || image.src"
        class="max-w-[70vw] max-h-[70vh] mx-auto rounded shadow block border-8 border-purple-300"
        :alt="image.name || ''"
        @click.stop
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { PropType } from 'vue'

const props = defineProps({
  modelValue: {
    type: [Boolean, Number],
    default: false
  },
  image: {
    type: Object as PropType<Record<string, any> | null>,
    required: false,
    default: null
  }
})
const emit = defineEmits(['update:modelValue', 'close'])

function formatFileSize(bytes:number){
  if(!bytes) return '0B'
  const k=1024, sizes=['B','KB','MB','GB']
  const i=Math.floor(Math.log(bytes)/Math.log(k))
  return (bytes/Math.pow(k,i)).toFixed(1)+' '+sizes[i]
}

function emitClose() {
  emit('update:modelValue', false)
  emit('close')
}
</script> 
