<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center" @click="closeModal">
    <div class="relative bg-black rounded-lg shadow-2xl max-w-5xl max-h-[90vh] overflow-hidden" @click.stop>
      <!-- Close Button -->
      <button 
        @click="closeModal" 
        class="absolute top-6 right-4 z-20 bg-blue-500 hover:bg-blue-600 rounded-full p-2 transition-all"
      >
        <svg class="w-6 h-6 text-white" fill="currentColor" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>

      <!-- Main Image Container -->
      <div class="relative flex items-center justify-center p-6">
        <!-- Navigation Arrows -->
        <button 
          v-if="hasMultipleImages"
          @click="previousImage"
          class="absolute left-4 z-10 bg-blue-500 hover:bg-blue-600 rounded-full p-3 transition-all"
        >
          <svg class="w-6 h-6 text-white" fill="currentColor" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>
        
        <button 
          v-if="hasMultipleImages"
          @click="nextImage"
          class="absolute right-4 z-10 bg-blue-500 hover:bg-blue-600 rounded-full p-3 transition-all"
        >
          <svg class="w-6 h-6 text-white" fill="currentColor" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>

        <!-- Main Image -->
        <img
          v-if="currentImage"
          :src="currentImage"
          :alt="productName"
          class="max-w-full max-h-[60vh] object-contain rounded-lg"
          @click.stop
        />
      </div>

      <!-- Image Gallery -->
      <div v-if="hasMultipleImages" class="px-4 pb-4 flex flex-col items-center space-y-3">
        <!-- Gallery Title -->
        <h3 class="text-white text-lg font-medium">عکس و ویدئو محصول</h3>
        
        <!-- Thumbnails -->
        <div class="flex justify-center gap-2 overflow-hidden">
          <button
            v-for="(image, index) in images"
            :key="index"
            @click="setCurrentImage(index)"
            :class="[
              'w-12 h-12 rounded-lg border-2 overflow-hidden transition-all',
              currentImageIndex === index 
                ? 'border-blue-500 shadow-lg' 
                : 'border-gray-300 hover:border-gray-400'
            ]"
          >
            <img 
              :src="image" 
              :alt="productName"
              class="w-full h-full object-cover"
            />
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'

interface Props {
  isOpen: boolean
  images: string[]
  currentIndex?: number
  productName?: string
}

const props = withDefaults(defineProps<Props>(), {
  currentIndex: 0,
  productName: 'محصول'
})

const emit = defineEmits(['close'])

const currentImageIndex = ref(props.currentIndex)

// بررسی وجود چندین تصویر
const hasMultipleImages = computed(() => props.images.length > 1)

// تصویر فعلی
const currentImage = computed(() => {
  if (props.images.length === 0) return ''
  return props.images[currentImageIndex.value] || props.images[0]
})

// تغییر تصویر فعلی
function setCurrentImage(index: number) {
  currentImageIndex.value = index
}

// تصویر قبلی
function previousImage() {
  if (currentImageIndex.value > 0) {
    currentImageIndex.value--
  } else {
    currentImageIndex.value = props.images.length - 1
  }
}

// تصویر بعدی
function nextImage() {
  if (currentImageIndex.value < props.images.length - 1) {
    currentImageIndex.value++
  } else {
    currentImageIndex.value = 0
  }
}

// بستن مودال
function closeModal() {
  emit('close')
}

// تماشای تغییرات props
watch(() => props.currentIndex, (newIndex) => {
  currentImageIndex.value = newIndex
})

// کلیدهای کیبورد
function handleKeydown(event: KeyboardEvent) {
  if (!props.isOpen) return
  
  switch (event.key) {
    case 'Escape':
      closeModal()
      break
    case 'ArrowLeft':
      previousImage()
      break
    case 'ArrowRight':
      nextImage()
      break
  }
}

// اضافه کردن event listener برای کیبورد
onMounted(() => {
  if (typeof window !== 'undefined') {
    window.addEventListener('keydown', handleKeydown)
  }
})

onUnmounted(() => {
  if (typeof window !== 'undefined') {
    window.removeEventListener('keydown', handleKeydown)
  }
})
</script>
