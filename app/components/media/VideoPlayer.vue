<template>
  <div class="video-player">
    <div class="relative">
      <video
        ref="videoRef"
        :src="src"
        :poster="poster"
        class="w-full rounded-lg"
        controls
        preload="metadata"
        @loadedmetadata="onLoadedMetadata"
        @timeupdate="onTimeUpdate"
        @ended="onEnded"
      >
        مرورگر شما از پخش ویدیو پشتیبانی نمی‌کند.
      </video>
      
      <div v-if="showControls" class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/50 to-transparent p-6">
        <div class="flex items-center justify-between text-white">
          <div class="flex items-center space-x-4 space-x-reverse">
            <button @click="togglePlay" class="p-2 hover:bg-white/20 rounded">
              <svg v-if="!isPlaying" class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8 5v14l11-7z"/>
              </svg>
              <svg v-else class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z"/>
              </svg>
            </button>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm">{{ formatTime(currentTime) }}</span>
              <span class="text-sm">/</span>
              <span class="text-sm">{{ formatTime(duration) }}</span>
            </div>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <button @click="toggleMute" class="p-2 hover:bg-white/20 rounded">
              <svg v-if="!isMuted" class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path d="M16.5 12c0-1.77-1.02-3.29-2.5-4.03v2.21l2.45 2.45c.03-.2.05-.41.05-.63zm2.5 0c0 .94-.2 1.82-.54 2.64l1.51 1.51C20.63 14.91 21 13.5 21 12c0-4.28-2.99-7.86-7-8.77v2.06c2.89.86 5 3.54 5 6.71zM4.27 3L3 4.27 7.73 9H3v6h4l5 5v-6.73l4.25 4.25c-.67.52-1.42.93-2.25 1.18v2.06c1.38-.31 2.63-.95 3.69-1.81L19.73 21 21 19.73l-9-9L4.27 3zM12 4L9.91 6.09 12 8.18V4z"/>
              </svg>
              <svg v-else class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path d="M3 9v6h4l5 5V4L7 9H3zm13.5 3c0-1.77-1.02-3.29-2.5-4.03v8.05c1.48-.73 2.5-2.25 2.5-4.02zM14 3.23v2.06c2.89.86 5 3.54 5 6.71s-2.11 5.85-5 6.71v2.06c4.01-.91 7-4.49 7-8.77s-2.99-7.86-7-8.77z"/>
              </svg>
            </button>
            
            <input
              v-model="volume"
              type="range"
              min="0"
              max="1"
              step="0.1"
              class="w-20"
              @input="onVolumeChange"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  src: {
    type: String,
    required: true
  },
  poster: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['play', 'pause', 'ended', 'timeupdate'])

const videoRef = ref(null)
const isPlaying = ref(false)
const isMuted = ref(false)
const currentTime = ref(0)
const duration = ref(0)
const volume = ref(1)
const showControls = ref(false)

const togglePlay = () => {
  if (isPlaying.value) {
    videoRef.value?.pause()
  } else {
    videoRef.value?.play()
  }
}

const toggleMute = () => {
  if (videoRef.value) {
    videoRef.value.muted = !videoRef.value.muted
    isMuted.value = videoRef.value.muted
  }
}

const onVolumeChange = () => {
  if (videoRef.value) {
    videoRef.value.volume = volume.value
  }
}

const onLoadedMetadata = () => {
  if (videoRef.value) {
    duration.value = videoRef.value.duration
  }
}

const onTimeUpdate = () => {
  if (videoRef.value) {
    currentTime.value = videoRef.value.currentTime
    emit('timeupdate', currentTime.value)
  }
}

const onEnded = () => {
  isPlaying.value = false
  emit('ended')
}

const formatTime = (time) => {
  const minutes = Math.floor(time / 60)
  const seconds = Math.floor(time % 60)
  return `${minutes}:${seconds.toString().padStart(2, '0')}`
}

onMounted(() => {
  if (videoRef.value) {
    videoRef.value.addEventListener('play', () => {
      isPlaying.value = true
      emit('play')
    })
    
    videoRef.value.addEventListener('pause', () => {
      isPlaying.value = false
      emit('pause')
    })
  }
})
</script> 