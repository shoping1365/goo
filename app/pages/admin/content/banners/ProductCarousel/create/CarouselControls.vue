<template>
  <div class="carousel-controls">
    <div class="controls-container">
      <!-- Autoplay Toggle -->
      <button
        @click="$emit('toggle-autoplay')"
        class="control-button autoplay-toggle"
        :class="{ 'active': isPlaying }"
        :title="isPlaying ? 'Pause autoplay' : 'Start autoplay'"
      >
        <svg v-if="isPlaying" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6"/>
        </svg>
        <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
      </button>

      <!-- Autoplay Delay Control -->
      <div class="delay-control">
        <label class="delay-label">Delay:</label>
        <select
          :value="autoplayDelay / 1000"
          @change="handleDelayChange"
          class="delay-select"
        >
          <option value="2">2s</option>
          <option value="3">3s</option>
          <option value="5">5s</option>
          <option value="7">7s</option>
          <option value="10">10s</option>
        </select>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Props
interface Props {
  autoplayEnabled: boolean
  autoplayDelay: number
  isPlaying: boolean
}

const props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  'toggle-autoplay': []
  'change-delay': [delay: number]
}>()

// Methods
const handleDelayChange = (event: Event) => {
  const target = event.target as HTMLSelectElement
  const newDelay = parseInt(target.value) * 1000
  emit('change-delay', newDelay)
}
</script>

<style scoped>
.carousel-controls {
  position: absolute;
  top: 16px;
  right: 16px;
  z-index: 10;
}

.controls-container {
  display: flex;
  align-items: center;
  gap: 12px;
  background: rgba(0, 0, 0, 0.7);
  padding: 8px 12px;
  border-radius: 20px;
  backdrop-filter: blur(10px);
}

.control-button {
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.control-button:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(1.1);
}

.control-button.active {
  background: #3b82f6;
}

.delay-control {
  display: flex;
  align-items: center;
  gap: 6px;
}

.delay-label {
  color: white;
  font-size: 12px;
  font-weight: 500;
}

.delay-select {
  background: rgba(255, 255, 255, 0.9);
  border: none;
  border-radius: 12px;
  padding: 4px 8px;
  font-size: 12px;
  color: #374151;
  cursor: pointer;
  transition: all 0.3s ease;
}

.delay-select:hover {
  background: white;
}

.delay-select:focus {
  outline: none;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}

/* Responsive */
@media (max-width: 768px) {
  .carousel-controls {
    top: 12px;
    right: 12px;
  }
  
  .controls-container {
    padding: 6px 10px;
    gap: 8px;
  }
  
  .control-button {
    width: 28px;
    height: 28px;
  }
  
  .delay-label {
    font-size: 11px;
  }
  
  .delay-select {
    font-size: 11px;
    padding: 3px 6px;
  }
}
</style>

