<template>
  <div ref="containerRef" :class="props.class"></div>
</template>

<script setup lang="ts">
import DOMPurify from 'dompurify';
import { computed, onMounted, ref, watch } from 'vue';

const props = defineProps<{
  content?: string
  class?: string
}>()

const containerRef = ref<HTMLDivElement | null>(null)

// Sanitize HTML content to prevent XSS attacks
const sanitizedContent = computed(() => {
  if (!props.content) return ''
  
  // Only sanitize on client side (DOMPurify requires DOM)
  if (typeof window === 'undefined') {
    return props.content
  }
  
  return DOMPurify.sanitize(props.content)
})

// Update innerHTML using DOM API instead of v-html
const updateContent = () => {
  if (containerRef.value) {
    if (sanitizedContent.value) {
      containerRef.value.innerHTML = sanitizedContent.value
    } else {
      containerRef.value.innerHTML = ''
    }
  }
}

watch(() => sanitizedContent.value, () => {
  updateContent()
})

onMounted(() => {
  updateContent()
})
</script>

