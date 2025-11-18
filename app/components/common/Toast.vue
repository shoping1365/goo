<template>
  <div v-if="isVisible" class="fixed top-6 right-4 z-50">
    <div :class="[
      'px-4 py-3 rounded-lg shadow-lg max-w-sm',
      type === 'success' ? 'bg-green-500 text-white' : '',
      type === 'error' ? 'bg-red-500 text-white' : '',
      type === 'warning' ? 'bg-yellow-500 text-white' : '',
      type === 'info' ? 'bg-blue-500 text-white' : ''
    ]">
      <div class="flex items-center justify-between">
        <span class="text-sm font-medium">{{ message }}</span>
        <button @click="close" class="ml-3 text-white hover:text-gray-200">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  message: {
    type: String,
    required: true
  },
  type: {
    type: String,
    default: 'info',
    validator: (value) => ['success', 'error', 'warning', 'info'].includes(value)
  },
  duration: {
    type: Number,
    default: 3000
  }
})

const emit = defineEmits(['close'])

const isVisible = ref(true)

const close = () => {
  isVisible.value = false
  emit('close')
}

onMounted(() => {
  if (props.duration > 0) {
    setTimeout(() => {
      close()
    }, props.duration)
  }
})
</script> 
