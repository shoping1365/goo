<template>
  <div v-if="isVisible" class="fixed inset-0 z-50 flex items-center justify-center">
    <div class="absolute inset-0 bg-black bg-opacity-75" @click="close"></div>
    
    <div class="relative bg-white rounded-lg max-w-4xl w-full mx-4 max-h-[90vh] overflow-hidden">
      <div class="flex items-center justify-between p-6 border-b">
        <h3 class="text-lg font-semibold text-gray-800">{{ title }}</h3>
        <button 
          @click="close"
          class="text-gray-400 hover:text-gray-600 transition-colors"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>
      
      <div class="p-6">
        <VideoPlayer 
          :src="videoSrc" 
          :poster="poster"
          @play="onPlay"
          @pause="onPause"
          @ended="onEnded"
        />
      </div>
      
      <div class="p-6 border-t">
        <div class="flex justify-between items-center">
          <div class="text-sm text-gray-600">
            <p>نام فایل: {{ fileName }}</p>
            <p>حجم: {{ fileSize }}</p>
          </div>
          
          <div class="flex space-x-2 space-x-reverse">
            <button 
              @click="download"
              class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
            >
              دانلود
            </button>
            <button 
              @click="close"
              class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 transition-colors"
            >
              بستن
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  isVisible: {
    type: Boolean,
    default: false
  },
  videoSrc: {
    type: String,
    required: true
  },
  poster: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: 'پیش‌نمایش ویدیو'
  },
  fileName: {
    type: String,
    default: ''
  },
  fileSize: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['close', 'play', 'pause', 'ended'])

const close = () => {
  emit('close')
}

const download = () => {
  const link = document.createElement('a')
  link.href = props.videoSrc
  link.download = props.fileName || 'video.mp4'
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

const onPlay = () => {
  emit('play')
}

const onPause = () => {
  emit('pause')
}

const onEnded = () => {
  emit('ended')
}
</script> 