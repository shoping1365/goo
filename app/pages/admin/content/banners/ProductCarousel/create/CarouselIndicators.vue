<template>
  <div class="carousel-indicators">
    <div class="indicators-container" :class="indicatorStyleClass">
      <button
        v-for="index in totalSlides"
        :key="index"
        class="indicator-dot"
        :class="{
          'active': index - 1 === currentIndex,
          'indicator-small': indicatorStyle === 'small',
          'indicator-large': indicatorStyle === 'large'
        }"
        :style="{ backgroundColor: index - 1 === currentIndex ? indicatorColor : undefined }"
        :aria-label="`Go to slide ${index}`"
        @click="$emit('change', index - 1)"
      ></button>
    </div>
  </div>
</template>

<script setup lang="ts">
// Props
interface Props {
  totalSlides: number
  currentIndex: number
  indicatorStyle?: 'default' | 'small' | 'large'
  indicatorColor?: string
}

const props = withDefaults(defineProps<Props>(), {
  indicatorStyle: 'default',
  indicatorColor: '#3b82f6'
})

// Emits
const emit = defineEmits<{
  change: [index: number]
}>()

import { computed } from 'vue'

// Computed
const indicatorStyleClass = computed(() => {
  switch (props.indicatorStyle) {
    case 'small':
      return 'indicators-small'
    case 'large':
      return 'indicators-large'
    default:
      return 'indicators-default'
  }
})
</script>

<style scoped>
.carousel-indicators {
  position: absolute;
  bottom: 16px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 10;
}

.indicators-container {
  display: flex;
  gap: 8px;
  align-items: center;
}

.indicator-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  border: none;
  background-color: rgba(255, 255, 255, 0.5);
  cursor: pointer;
  transition: all 0.3s ease;
  padding: 0;
}

.indicator-dot:hover {
  background-color: rgba(255, 255, 255, 0.8);
  transform: scale(1.2);
}

.indicator-dot.active {
  background-color: v-bind(indicatorColor);
  transform: scale(1.3);
}

/* Indicator Styles */
.indicators-small .indicator-dot {
  width: 6px;
  height: 6px;
}

.indicators-large .indicator-dot {
  width: 12px;
  height: 12px;
}

.indicator-small {
  width: 6px !important;
  height: 6px !important;
}

.indicator-large {
  width: 12px !important;
  height: 12px !important;
}

/* Responsive */
@media (max-width: 768px) {
  .carousel-indicators {
    bottom: 12px;
  }
  
  .indicators-container {
    gap: 6px;
  }
  
  .indicator-dot {
    width: 6px;
    height: 6px;
  }
  
  .indicators-large .indicator-dot {
    width: 10px;
    height: 10px;
  }
}
</style>

