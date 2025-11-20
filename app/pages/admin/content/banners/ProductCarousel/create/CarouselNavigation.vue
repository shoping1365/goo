<template>
  <div class="carousel-navigation">
    <!-- Previous Button -->
    <button
      v-if="canGoPrev"
      class="nav-button nav-prev"
      :class="navigationStyleClass"
      :style="{ width: `${navigationSize}px`, height: `${navigationSize}px` }"
      aria-label="Previous slide"
      @click="$emit('prev')"
    >
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
      </svg>
    </button>

    <!-- Next Button -->
    <button
      v-if="canGoNext"
      class="nav-button nav-next"
      :class="navigationStyleClass"
      :style="{ width: `${navigationSize}px`, height: `${navigationSize}px` }"
      aria-label="Next slide"
      @click="$emit('next')"
    >
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
      </svg>
    </button>
  </div>
</template>

<script setup lang="ts">
// Props
interface Props {
  canGoPrev: boolean
  canGoNext: boolean
  navigationStyle?: 'default' | 'minimal' | 'bold'
  navigationSize?: number
}

const props = withDefaults(defineProps<Props>(), {
  navigationStyle: 'default',
  navigationSize: 40
})

// Emits
import { computed } from 'vue';

const emit = defineEmits<{
  prev: []
  next: []
}>()

// Computed
const navigationStyleClass = computed(() => {
  const baseClass = 'nav-button'
  switch (props.navigationStyle) {
    case 'minimal':
      return `${baseClass} nav-minimal`
    case 'bold':
      return `${baseClass} nav-bold`
    default:
      return `${baseClass} nav-default`
  }
})
</script>

<style scoped>
.carousel-navigation {
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  transform: translateY(-50%);
  pointer-events: none;
  z-index: 10;
}

.nav-button {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(255, 255, 255, 0.9);
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  pointer-events: auto;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.nav-button:hover {
  background: rgba(255, 255, 255, 1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  transform: translateY(-50%) scale(1.1);
}

.nav-button:active {
  transform: translateY(-50%) scale(0.95);
}

.nav-prev {
  left: 16px;
}

.nav-next {
  right: 16px;
}

/* Navigation Styles */
.nav-default {
  color: #374151;
}

.nav-minimal {
  background: rgba(255, 255, 255, 0.7);
  color: #6b7280;
}

.nav-bold {
  background: #3b82f6;
  color: white;
}

.nav-bold:hover {
  background: #2563eb;
}

/* Responsive */
@media (max-width: 768px) {
  .nav-button {
    width: 36px !important;
    height: 36px !important;
  }
  
  .nav-prev {
    left: 8px;
  }
  
  .nav-next {
    right: 8px;
  }
}
</style>

