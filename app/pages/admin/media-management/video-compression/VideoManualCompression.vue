<template>
  <div class="order-3 bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4 w-full max-w-full min-w-0 overflow-x-hidden">
    <div class="flex items-center cursor-pointer px-6 py-4 relative">
      <!-- Toggle icon -->
      <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 text-2xl select-none cursor-pointer" @click="toggleSection('videoSelection')">{{ sections.videoSelection ? '−' : '+' }}</span>
      <!-- Scan + dynamic action buttons container -->
      <div class="absolute left-12 top-1/2 -translate-y-1/2 flex items-center gap-2 select-none">
        <!-- Scan -->
        <button class="bg-blue-100 text-blue-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-blue-200 transition-colors" @click.stop="scanVideos">اسکن</button>
        <!-- Select All (shown when at least one selected & not all) -->
        <button v-if="selectedVideos.length > 0 && selectedVideos.length < videos.length" class="bg-indigo-100 text-indigo-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-indigo-200 transition-colors" @click.stop="selectAllVideos">انتخاب همه</button>
        <!-- Compress -->
        <button v-if="selectedVideos.length > 0" :disabled="isCompressing" class="bg-pink-100 text-pink-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-pink-200 transition-colors disabled:opacity-50" @click.stop="startCompression">فشرده‌سازی</button>
      </div>
      <!-- Title pill -->
      <span class="text-base font-bold bg-blue-100 text-blue-700 px-4 py-1 rounded-lg ml-auto">فشرده‌سازی دستی</span>
    </div>
    <div class="border-b border-gray-200 mx-6 w-full max-w-full min-w-0"></div>
    <div v-show="sections.videoSelection">
      <div class="px-6 py-4 flex items-center justify-between w-full max-w-full min-w-0"></div>
      <!-- Videos Grid (shows ~3 rows; scroll for more) -->
      <div class="p-6 max-h-[540px] overflow-y-auto" dir="rtl">
        <div v-if="videos.length === 0" class="text-center py-8">
          <svg class="w-24 h-24 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
          </svg>
          <h3 class="text-lg font-medium text-gray-900 mb-2">ویدیویی برای فشرده‌سازی یافت نشد</h3>
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          <div 
            v-for="video in videos" 
            :key="video.id"
            :class="[
              'relative group cursor-pointer rounded-xl border-2 transition-all duration-200 overflow-hidden shadow-md hover:shadow-lg',
              selectedVideos.includes(video.id) 
                ? 'border-blue-500 bg-blue-50' 
                : 'border-gray-200 hover:border-gray-300'
            ]"
            @click="toggleVideoSelection(video.id)"
          >
            <!-- Checkbox -->
            <div class="absolute top-2 right-2 z-10">
              <input
                type="checkbox"
                :checked="selectedVideos.includes(video.id)"
                class="w-4 h-4 text-blue-600 bg-white border-gray-300 rounded focus:ring-blue-500"
                @click.stop="toggleVideoSelection(video.id)"
              >
            </div>
            <!-- Video Preview -->
            <div class="aspect-video bg-gray-100 flex items-center justify-center relative">
              <img
                v-if="video.thumbnail"
                :src="video.thumbnail"
                :alt="video.name"
                class="w-full h-full object-cover"
              >
              <div v-else class="flex items-center justify-center">
                <svg class="w-16 h-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                </svg>
              </div>
              <!-- Duration Badge -->
              <div class="absolute bottom-2 left-2 bg-black bg-opacity-75 text-white text-xs px-2 py-1 rounded">
                {{ video.duration }}
              </div>
              <!-- Play Button -->
              <div class="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                <div class="bg-black bg-opacity-50 rounded-full p-3">
                  <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M8 5v14l11-7z"/>
                  </svg>
                </div>
              </div>
            </div>
            <!-- Video Info -->
            <div class="p-6 bg-white">
              <p class="text-sm font-medium text-gray-900 truncate" :title="video.name">
                {{ video.name }}
              </p>
              <div class="flex items-center justify-between mt-2">
                <p class="text-xs text-gray-500">{{ formatFileSize(video.size) }}</p>
                <p class="text-xs text-gray-500">{{ video.resolution }}</p>
              </div>
              <div class="flex items-center justify-between mt-1">
                <p class="text-xs text-gray-500">{{ video.format.toUpperCase() }}</p>
                <p class="text-xs text-gray-500">{{ video.bitrate }} kbps</p>
              </div>
            </div>
            <!-- Compression Status -->
            <div v-if="video.compressionStatus" class="absolute inset-0 bg-black bg-opacity-75 flex items-center justify-center">
              <div class="text-center text-white">
                <div v-if="video.compressionStatus === 'processing'" class="space-y-2">
                  <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-white mx-auto"></div>
                  <p class="text-sm">در حال فشرده‌سازی...</p>
                  <div v-if="video.progress !== undefined" class="w-full bg-gray-700 rounded-full h-2">
                    <div class="bg-blue-500 h-2 rounded-full" :style="{ width: video.progress + '%' }"></div>
                  </div>
                  <p class="text-xs">{{ video.progress }}%</p>
                </div>
                <div v-else-if="video.compressionStatus === 'completed'" class="space-y-2">
                  <svg class="w-8 h-8 text-green-400 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                  </svg>
                  <p class="text-sm">تکمیل شد</p>
                  <div v-if="video.compressionResult" class="text-xs">
                    <p>کاهش: {{ video.compressionResult.reduction }}%</p>
                    <p>حجم جدید: {{ formatFileSize(video.compressionResult.newSize) }}</p>
                  </div>
                </div>
                <div v-else-if="video.compressionStatus === 'error'" class="space-y-2">
                  <svg class="w-8 h-8 text-red-400 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                  <p class="text-sm">خطا</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

interface VideoFile {
  id: number
  name: string
  url: string
  thumbnail: string
  size: number
  duration: string
  resolution: string
  format: string
  bitrate: number
  compressionStatus?: 'processing' | 'completed' | 'error'
  progress?: number
  compressionResult?: {
    reduction: number
    newSize: number
  }
}

interface CompressionQueueItem {
  id: number
  name: string
  thumbnail: string
  size: number
  duration: string
  originalFormat: string
  status: 'pending' | 'processing' | 'completed' | 'error' | 'paused'
  progress: number
  eta: string
  error?: string
  originalSize: number
  compressedSize: number
  reduction: number
  processingTime: string
  jobId?: string
  cancel?: () => Promise<void>
  pause?: () => Promise<void>
  resume?: () => Promise<void>
}

// Reactive data
const selectedVideos = ref<number[]>([])
const isCompressing = ref(false)
const compressionQueue = ref<CompressionQueueItem[]>([])
const _allPaused = ref(false)

// Sample videos data
const videos = ref<VideoFile[]>([
  {
    id: 1,
    name: 'sample-video-1.mp4',
    url: '/uploads/videos/sample1.mp4',
    thumbnail: '/uploads/thumbnails/sample1.jpg',
    size: 52428800, // 50MB
    duration: '2:30',
    resolution: '1920x1080',
    format: 'mp4',
    bitrate: 5000
  },
  {
    id: 2,
    name: 'sample-video-2.mov',
    url: '/uploads/videos/sample2.mov',
    thumbnail: '/uploads/thumbnails/sample2.jpg',
    size: 104857600, // 100MB
    duration: '5:15',
    resolution: '3840x2160',
    format: 'mov',
    bitrate: 8000
  }
])

// Collapsible sections state
const defaultSections = {
  videoSelection: true
} as const

const sections = reactive<Record<keyof typeof defaultSections, boolean>>({ ...defaultSections })

function toggleSection(section: keyof typeof sections) {
  sections[section] = !sections[section]
}

const toggleVideoSelection = (videoId: number) => {
  const index = selectedVideos.value.indexOf(videoId)
  if (index > -1) {
    selectedVideos.value.splice(index, 1)
  } else {
    selectedVideos.value.push(videoId)
  }
}

const selectAllVideos = () => {
  selectedVideos.value = videos.value.map(video => video.id)
}

const _clearSelection = () => {
  selectedVideos.value = []
}

const startCompression = async () => {
  if (selectedVideos.value.length === 0) return
  
  isCompressing.value = true
  
  // Add selected videos to compression queue
  selectedVideos.value.forEach(videoId => {
    const video = videos.value.find(v => v.id === videoId)
    if (video) {
      video.compressionStatus = 'processing'
      video.progress = 0
      
      const queueItem: CompressionQueueItem = {
        id: video.id,
        name: video.name,
        thumbnail: video.thumbnail,
        size: video.size,
        duration: video.duration,
        originalFormat: video.format,
        status: 'processing',
        progress: 0,
        eta: 'محاسبه...',
        originalSize: video.size,
        compressedSize: 0,
        reduction: 0,
        processingTime: '0:00'
      }
      
      compressionQueue.value.push(queueItem)
      processVideo(queueItem)
    }
  })
  
  selectedVideos.value = []
  isCompressing.value = false
}

const processVideo = async (queueItem: CompressionQueueItem) => {
  try {
    // شبیه‌سازی فرآیند فشرده‌سازی
    const totalSteps = 100
    for (let i = 0; i <= totalSteps; i++) {
      await new Promise(resolve => setTimeout(resolve, 100))
      queueItem.progress = i
      
      // به‌روزرسانی ویدیو در لیست
      const video = videos.value.find(v => v.id === queueItem.id)
      if (video) {
        video.progress = i
      }
    }
    
    // شبیه‌سازی نتیجه
    queueItem.status = 'completed'
    queueItem.compressedSize = Math.floor(queueItem.originalSize * 0.6)
    queueItem.reduction = 40
    
    const video = videos.value.find(v => v.id === queueItem.id)
    if (video) {
      video.compressionStatus = 'completed'
      video.compressionResult = {
        reduction: 40,
        newSize: queueItem.compressedSize
      }
    }
    
  } catch (e: unknown) {
    interface ErrorResponse {
      message?: string
    }
    const err = e as ErrorResponse
    queueItem.status = 'error'
    queueItem.error = err?.message || 'خطای ناشناخته'
    
    const video = videos.value.find(v => v.id === queueItem.id)
    if (video) {
      video.compressionStatus = 'error'
    }
  }
}

const scanVideos = () => {
  // شبیه‌سازی اسکن ویدیوها
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>

<!--
  کامپوننت فشرده‌سازی دستی ویدیو
  - نمایش لیست ویدیوها برای انتخاب
  - انتخاب چندتایی ویدیوها
  - شروع فرآیند فشرده‌سازی
  - نمایش وضعیت پیشرفت فشرده‌سازی
  - مدیریت صف فشرده‌سازی
  - توضیحات کامل به فارسی در کد
--> 
