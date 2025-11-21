<template>
  <div class="product-video-wrapper">
    <div v-if="showPlayer && normalizedUrl" class="video-player">
      <!-- YouTube Video -->
      <div v-if="isYouTubeVideo" class="youtube-video">
        <iframe
          :src="youtubeEmbedUrl"
          :title="title || 'ویدیو محصول'"
          class="video-element"
          frameborder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowfullscreen
        ></iframe>
      </div>
      
      <!-- Vimeo Video -->
      <div v-else-if="isVimeoVideo" class="vimeo-video">
        <iframe
          :src="vimeoEmbedUrl"
          :title="title || 'ویدیو محصول'"
          class="video-element"
          frameborder="0"
          allow="autoplay; fullscreen; picture-in-picture"
          allowfullscreen
        ></iframe>
      </div>
      
      <!-- Aparat Video -->
      <div v-else-if="isAparatVideo" class="aparat-video">
        <iframe
          :src="aparatEmbedUrl"
          :title="title || 'ویدیو محصول'"
          class="video-element"
          frameborder="0"
          allowfullscreen
        ></iframe>
      </div>
      
      <!-- Direct Video File -->
      <div v-else class="direct-video">
        <video 
          v-if="normalizedUrl"
          ref="videoPlayer"
          :controls="props.showControls"
          :autoplay="autoplay"
          :muted="autoplay"
          :loop="false"
          :poster="props.posterImage"
          preload="metadata"
          playsinline
          class="video-element"
          @loadedmetadata="onVideoLoaded"
          @error="onVideoError"
          @play="onVideoPlay"
          @pause="onVideoPause"
          @ended="onVideoEnded"
        >
          <source :src="normalizedUrl" type="video/mp4">
          <source :src="normalizedUrl" type="video/webm">
          <source :src="normalizedUrl" type="video/ogg">
          <p class="video-fallback">مرورگر شما از پخش ویدیو پشتیبانی نمی‌کند.</p>
        </video>
        
        <!-- Error Display -->
        <div v-if="videoError" class="mt-4 p-3 bg-red-50 rounded-lg text-red-800">
          <div class="font-semibold">❌ خطا در بارگذاری ویدیو:</div>
          <div>URL: {{ props.videoUrl }}</div>
          <div>لطفاً فایل را بررسی کنید.</div>
        </div>
        
        <!-- Custom Controls for Direct Videos -->
        <div v-if="!props.showControls" class="custom-controls">
          <button 
            class="play-button"
            :class="{ 'playing': isPlaying }"
            @click="togglePlay"
          >
            <svg v-if="!isPlaying" class="w-8 h-8" fill="currentColor" viewBox="0 0 24 24">
              <path d="M8 5v14l11-7z"/>
            </svg>
            <svg v-else class="w-8 h-8" fill="currentColor" viewBox="0 0 24 24">
              <path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z"/>
            </svg>
          </button>
        </div>
      </div>
      
      <!-- Video Info -->
      <div v-if="title || description" class="video-info mt-4">
        <h3 v-if="title" class="text-lg font-semibold text-gray-900 mb-2">{{ title }}</h3>
        <p v-if="description" class="text-gray-600 text-sm">{{ description }}</p>
      </div>
      
      <!-- کنترل‌های اضافی -->
      <div class="video-controls mt-4">
        <button 
          class="btn-secondary" 
          title="نمایش تمام صفحه"
          @click="toggleFullscreen"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5l-5-5m5 5v-4m0 4h-4"></path>
          </svg>
          تمام صفحه
        </button>
        
        <button 
          v-if="!isExternalVideo"
          class="btn-secondary" 
          title="دانلود ویدیو"
          @click="downloadVideo"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          دانلود
        </button>
      </div>
    </div>
    <!-- Placeholder preview -->
    <div v-else class="no-video" style="cursor:pointer" @click="activatePlayer">
      <img :src="props.posterImage || '/default-product.svg'" alt="Video Preview" class="w-full h-full object-cover rounded-lg">
      <p class="text-center text-gray-600 mt-2">برای پخش کلیک کنید</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'

interface Props {
  videoUrl?: string
  posterImage?: string
  title?: string
  description?: string
  autoplay?: boolean
  showControls?: boolean
  showInGallery?: boolean
  lazyLoad?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  autoplay: false,
  showControls: true,
  showInGallery: true,
  lazyLoad: false
})

const showPlayer = ref(!props.lazyLoad)

function activatePlayer(){
  if (showPlayer.value) return
  showPlayer.value = true
}

// Type definitions for external libraries
type PlyrInstance = {
  destroy: () => void
  on: (event: string, callback: (data?: unknown) => void) => void
}

type HlsInstance = {
  destroy: () => void
  loadSource: (url: string) => void
  attachMedia: (element: HTMLVideoElement) => void
}

const videoPlayer = ref<HTMLVideoElement | null>(null)
const player = ref<PlyrInstance | null>(null)
const hlsInstance = ref<HlsInstance | null>(null)
const videoLoaded = ref(false)
const videoError = ref(false)
const isPlaying = ref(false)
const _currentTime = ref(0)
const duration = ref(0)

// Helper function to safely check URL domain
function isValidVideoDomain(url: string, allowedDomains: string[]): boolean {
  if (!url) return false
  try {
    // If it's a relative URL, it's safe
    if (!url.startsWith('http://') && !url.startsWith('https://')) {
      return true
    }
    const urlObj = new URL(url)
    const hostname = urlObj.hostname.toLowerCase()
    return allowedDomains.some(domain => hostname === domain || hostname.endsWith('.' + domain))
  } catch {
    return false
  }
}

// Computed properties for video type detection with proper URL validation
const isYouTubeVideo = computed(() => {
  if (!normalizedUrl.value) return false
  return isValidVideoDomain(normalizedUrl.value, ['youtube.com', 'youtu.be', 'www.youtube.com'])
})

const isVimeoVideo = computed(() => {
  if (!normalizedUrl.value) return false
  return isValidVideoDomain(normalizedUrl.value, ['vimeo.com', 'player.vimeo.com'])
})

const isAparatVideo = computed(() => {
  if (!normalizedUrl.value) return false
  return isValidVideoDomain(normalizedUrl.value, ['aparat.com', 'www.aparat.com'])
})

const isExternalVideo = computed(() => {
  return isYouTubeVideo.value || isVimeoVideo.value || isAparatVideo.value
})

// URL normalization (fix backslashes and missing leading slash)
const normalizedUrl = computed(() => {
  const raw = (props.videoUrl || '').toString().trim()
  if (!raw) return ''
  let p = raw.replace(/\\\\/g, '/').replace(/\/public\//i, '/')
  // یکسان‌سازی حروف برای uploads
  p = p.replace(/^\/uploads\b/i, '/uploads')
  if (!/^https?:\/\//i.test(p) && !p.startsWith('/')) {
    p = '/' + p
  }
  return p
})

// YouTube embed URL
const youtubeEmbedUrl = computed(() => {
  if (!isYouTubeVideo.value) return ''
  const videoId = extractYouTubeId(normalizedUrl.value)
  return `https://www.youtube.com/embed/${videoId}?autoplay=${props.autoplay ? 1 : 0}&rel=0&modestbranding=1`
})

// Vimeo embed URL
const vimeoEmbedUrl = computed(() => {
  if (!isVimeoVideo.value) return ''
  const videoId = extractVimeoId(normalizedUrl.value)
  return `https://player.vimeo.com/video/${videoId}?autoplay=${props.autoplay ? 1 : 0}&title=0&byline=0&portrait=0`
})

// Aparat embed URL
const aparatEmbedUrl = computed(() => {
  if (!isAparatVideo.value) return ''
  const videoId = extractAparatId(normalizedUrl.value)
  return `https://www.aparat.com/video/video/embed/videohash/${videoId}/vt/frame`
})

// Helper functions for extracting video IDs
function extractYouTubeId(url: string): string | null {
  const regex = /(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})/
  const match = url.match(regex)
  return match ? match[1] : null
}

function extractVimeoId(url: string): string | null {
  const regex = /vimeo\.com\/(?:.*#|.*\/videos\/)?([0-9]+)/
  const match = url.match(regex)
  return match ? match[1] : null
}

function extractAparatId(url: string): string | null {
  const regex = /aparat\.com\/v\/([^\/\?]+)/
  const match = url.match(regex)
  return match ? match[1] : null
}

// Video event handlers
const onVideoLoaded = (): void => {
  videoLoaded.value = true
  videoError.value = false
  if (videoPlayer.value) {
    duration.value = videoPlayer.value.duration
  }
}

const onVideoError = (error: Event): void => {
  console.error('Video error:', error)
  videoError.value = true
  videoLoaded.value = false
}

const onVideoPlay = () => {
  isPlaying.value = true
}

const onVideoPause = () => {
  isPlaying.value = false
}

const onVideoEnded = () => {
  isPlaying.value = false
}

// Custom controls
const togglePlay = () => {
  if (videoPlayer.value) {
    if (isPlaying.value) {
      videoPlayer.value.pause()
    } else {
      videoPlayer.value.play()
    }
  }
}

const toggleFullscreen = () => {
  if (videoPlayer.value) {
    if (document.fullscreenElement) {
      document.exitFullscreen()
    } else {
      videoPlayer.value.requestFullscreen()
    }
  }
}

const downloadVideo = () => {
  if (normalizedUrl.value) {
    const link = document.createElement('a')
    link.href = normalizedUrl.value
    link.download = 'video.mp4'
    link.click()
  }
}

// Initialize Plyr player (client-side only)
onMounted(async () => {
  if (import.meta.client && !isExternalVideo.value && normalizedUrl.value) {
    try {
      const { default: Plyr } = await import('plyr') as { default: new (element: HTMLVideoElement, options?: unknown) => PlyrInstance }
      const { default: Hls } = await import('hls.js') as { default: { isSupported: () => boolean, new (): HlsInstance } }
      
      // @ts-ignore: CSS import
      await import('plyr/dist/plyr.css')
      
      if (videoPlayer.value) {
        // Initialize HLS if needed
        if (normalizedUrl.value.includes('.m3u8')) {
          if (Hls.isSupported()) {
            hlsInstance.value = new Hls()
            hlsInstance.value.loadSource(normalizedUrl.value)
            hlsInstance.value.attachMedia(videoPlayer.value)
          } else if (videoPlayer.value.canPlayType('application/vnd.apple.mpegurl')) {
            videoPlayer.value.src = normalizedUrl.value
          }
        }
        
        // Initialize Plyr
        player.value = new Plyr(videoPlayer.value, {
          controls: props.showControls ? ['play', 'progress', 'current-time', 'duration', 'mute', 'volume', 'fullscreen'] : [],
          autoplay: props.autoplay,
          muted: props.autoplay,
          hideControls: !props.showControls
        })
        
        // Plyr event listeners
        player.value.on('ready', () => {
          videoLoaded.value = true
        })
        
        player.value.on('error', (error) => {
          console.error('Plyr error:', error)
          videoError.value = true
        })
      }
    } catch (error) {
      console.error('Failed to load Plyr:', error)
      videoError.value = true
    }
  }
})

// Cleanup
onUnmounted(() => {
  if (player.value) {
    player.value.destroy()
  }
  if (hlsInstance.value) {
    hlsInstance.value.destroy()
  }
})

// Watch for URL changes
watch(() => props.videoUrl, () => {
  videoError.value = false
  videoLoaded.value = false
  if (player.value) {
    player.value.destroy()
    player.value = null
  }
  if (hlsInstance.value) {
    hlsInstance.value.destroy()
    hlsInstance.value = null
  }
})
</script>

<style scoped>
.product-video-wrapper {
  width: 100%;
  overflow: hidden;
  border-radius: 0.5rem;
}

.video-player {
  position: relative;
}

.video-element {
  width: 100%;
  height: auto;
  max-height: 400px;
  background-color: #111827;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  aspect-ratio: 16/9;
  object-fit: contain;
}

.youtube-video,
.vimeo-video,
.aparat-video {
  position: relative;
}

.youtube-video iframe,
.vimeo-video iframe,
.aparat-video iframe {
  width: 100%;
  height: auto;
  max-height: 400px;
  background-color: #111827;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  aspect-ratio: 16/9;
  object-fit: contain;
}

.direct-video {
  position: relative;
}

.custom-controls {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0, 0, 0, 0.3);
}

.play-button {
  background-color: rgba(255, 255, 255, 0.8);
  border-radius: 9999px;
  padding: 0.75rem;
  transition: all 0.2s;
  color: #1f2937;
}

.play-button:hover {
  background-color: rgba(255, 255, 255, 1);
  color: #000000;
}

.play-button.playing {
  background-color: rgba(255, 255, 255, 0.6);
}

.video-controls {
  display: flex;
  justify-content: center;
  gap: 1rem;
}

.btn-secondary {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background-color: #f3f4f6;
  color: #374151;
  border-radius: 0.5rem;
  transition: background-color 0.2s;
}

.btn-secondary:hover {
  background-color: #e5e7eb;
}

.no-video {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 0;
  background-color: #f9fafb;
  border-radius: 0.5rem;
}

.video-fallback {
  text-align: center;
  color: #6b7280;
  padding: 1rem;
}

.video-info {
  background-color: white;
  border-radius: 0.5rem;
  padding: 1rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
}

/* Responsive styles */
@media (max-width: 768px) {
  .video-element,
  .youtube-video iframe,
  .vimeo-video iframe,
  .aparat-video iframe {
    max-height: 12rem;
  }
  
  .video-controls {
    flex-direction: column;
    align-items: center;
  }
  
  .btn-secondary {
    width: 100%;
    justify-content: center;
  }
}
</style> 