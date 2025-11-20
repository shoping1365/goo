<template>
  <div class="relative">
    <!-- Image Preview with Eye Icon -->
    <div class="relative group">
      <img 
        :src="imageUrl" 
        :alt="alt || 'Image preview'" 
        class="w-full h-auto rounded-lg object-cover"
      />
      <div 
        class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-30 transition-all duration-200 flex items-center justify-center cursor-pointer"
        @click="openViewer"
      >
        <div class="opacity-0 group-hover:opacity-100 transition-opacity duration-200">
          <!-- Eye Icon using Tailwind -->
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
        </div>
      </div>
    </div>

    <!-- Modal Viewer -->
    <Transition
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="transform scale-95 opacity-0"
      enter-to-class="transform scale-100 opacity-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="transform scale-100 opacity-100"
      leave-to-class="transform scale-95 opacity-0"
    >
      <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-75">
        <div class="relative max-w-4xl w-full mx-4">
          <!-- Close button -->
          <button 
            class="absolute -top-10 right-0 text-white hover:text-gray-300 transition-colors duration-200" 
            @click="closeViewer"
          >
            <!-- Close Icon using Tailwind -->
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>

          <!-- Image/Video Container -->
          <div class="bg-white rounded-lg overflow-hidden">
            <div class="relative">
              <img 
                v-if="!isVideo" 
                :src="imageUrl" 
                :alt="alt || 'Full size image'" 
                class="w-full h-auto cursor-zoom-in"
                :class="{ 'scale-150': isZoomed }"
                @click="toggleZoom"
              />
              <video 
                v-else 
                :src="imageUrl" 
                controls 
                class="w-full h-auto"
              ></video>
              <!-- Zoom Controls -->
              <div v-if="!isVideo" class="absolute bottom-4 right-4 flex gap-2">
                <button 
                  class="p-2 bg-white bg-opacity-75 rounded-full hover:bg-opacity-100 transition-all duration-200"
                  @click="zoomIn"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                  </svg>
                </button>
                <button 
                  class="p-2 bg-white bg-opacity-75 rounded-full hover:bg-opacity-100 transition-all duration-200"
                  @click="zoomOut"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
                  </svg>
                </button>
              </div>
            </div>

            <!-- Details Section -->
            <div class="p-6 bg-white">
              <h3 class="text-lg font-semibold mb-2">{{ title || 'Media Details' }}</h3>
              <div class="text-sm text-gray-600">
                <p v-if="description">{{ description }}</p>
                <p v-if="dimensions">ابعاد: {{ dimensions }}</p>
                <p v-if="fileSize">حجم فایل: {{ fileSize }}</p>
                <p v-if="uploadDate">تاریخ آپلود: {{ uploadDate }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Props {
  imageUrl: string
  alt?: string
  title?: string
  description?: string
  dimensions?: string
  fileSize?: string
  uploadDate?: string
  isVideo?: boolean
}

const props = defineProps<Props>()

const isOpen = ref(false)
const isZoomed = ref(false)
const zoomLevel = ref(1)

const openViewer = () => {
  isOpen.value = true
  document.body.style.overflow = 'hidden'
}

const closeViewer = () => {
  isOpen.value = false
  isZoomed.value = false
  zoomLevel.value = 1
  document.body.style.overflow = 'auto'
}

const toggleZoom = () => {
  isZoomed.value = !isZoomed.value
  zoomLevel.value = isZoomed.value ? 1.5 : 1
}

const zoomIn = () => {
  if (zoomLevel.value < 3) {
    zoomLevel.value += 0.5
    isZoomed.value = true
  }
}

const zoomOut = () => {
  if (zoomLevel.value > 1) {
    zoomLevel.value -= 0.5
    isZoomed.value = zoomLevel.value > 1
  }
}
</script>

<style scoped>
.scale-150 {
  transform: scale(1.5);
  transform-origin: center;
  transition: transform 0.3s ease-in-out;
}
</style> 
