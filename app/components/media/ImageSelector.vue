<template>
  <div class="grid grid-cols-3 md:grid-cols-6 lg:grid-cols-8 xl:grid-cols-10 gap-3">
    <div 
      v-for="image in images" 
      :key="image.id"
      :class="[
        'relative group rounded-xl border-2 transition-all duration-200 overflow-hidden shadow-md hover:shadow-lg',
        selectedImages.includes(image.id) 
          ? 'border-blue-500 bg-blue-50' 
          : 'border-gray-200 hover:border-gray-300'
      ]"
    >
      <!-- Image Preview -->
      <div 
        class="relative bg-gray-100 overflow-hidden flex items-center justify-center cursor-pointer" 
        style="padding-bottom:100%;"
        @click.stop="handlePreview(image, $event)"
      >
        <template v-if="image.type === 'video'">
          <svg class="absolute inset-0 m-auto w-12 h-12 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2" fill="#e0e7ff"/>
            <polygon points="10,8 16,12 10,16" fill="#3b82f6" />
          </svg>
        </template>
        <template v-else>
          <img 
            :src="image.url" 
            loading="lazy" 
            :alt="image.name"
            class="absolute inset-0 w-full h-full object-cover"
            @error="(e) => console.error('Image load error:', image.name, image.url, e)"
          >
        </template>
      </div>
      <!-- Image Info -->
      <div 
        class="p-3 bg-white relative info-section"
        style="cursor:pointer"
        @click.stop="handleSelect(image.id, $event)"
      >
        <p class="text-sm font-medium text-gray-900 truncate" :title="image.name">
          {{ image.name }}
        </p>
        <div class="flex items-center justify-between mt-1">
          <p class="text-xs text-gray-500">{{ formatFileSize(image.compressed_size ?? image.size) }}</p>
          <p class="text-xs text-gray-500">{{ image.dimensions }}</p>
        </div>
      </div>
      <!-- Compression Status -->
      <div v-if="image.compressionStatus" class="absolute inset-0 bg-black bg-opacity-75 flex items-center justify-center pointer-events-none">
        <div class="text-center text-white">
          <svg v-if="image.compressionStatus === 'processing'" class="animate-spin w-8 h-8 mx-auto mb-2" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else-if="image.compressionStatus === 'completed'" class="w-8 h-8 mx-auto mb-2 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
          </svg>
          <svg v-else-if="image.compressionStatus === 'error'" class="w-8 h-8 mx-auto mb-2 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
          <p class="text-sm font-medium">
            <span v-if="image.compressionStatus === 'processing'">در حال پردازش...</span>
            <span v-else-if="image.compressionStatus === 'completed'">تکمیل شد</span>
            <span v-else-if="image.compressionStatus === 'error'">خطا</span>
          </p>
          <p v-if="image.compressionStatus === 'completed' && image.compressionResult" class="text-xs mt-1">
            {{ image.compressionResult.reduction }}% کاهش حجم
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface ImageFile {
  id: number
  name: string
  url: string
  thumbnail: string
  size: number
  [key: string]: unknown
}

function handlePreview(img: ImageFile, e:MouseEvent){
  // اگر ctrl/meta فشرده شده باشد، حالت انتخاب چندگانه است
  if((e.ctrlKey || e.metaKey)) {
    emit('select', img.id, e)
    return
  }
  // در غیر این صورت پیش‌نمایش را باز کن
  emit('preview', img)
}

function handleSelect(imageId: number, e: MouseEvent) {
  // اگر ctrl/meta فشرده شده باشد، حالت انتخاب چندگانه است
  if((e.ctrlKey || e.metaKey)) {
    emit('select', imageId, e)
  } else {
    // در غیر این صورت، فقط این آیتم را انتخاب کن (انتخاب قبلی را پاک کن)
    emit('select', imageId, e)
  }
}
interface ImageFile {
  id: number
  name: string
  url: string // اضافه شد
  thumbnail: string
  size: number
  compressed_size?: number | null
  dimensions?: string
  compressionStatus?: 'processing' | 'completed' | 'error'
  compressionResult?: {
    reduction: number
    newSize: number
  }
  type?: string
  extension?: string // برای شناسایی فایل‌های GIF
}
defineProps<{
  images: ImageFile[],
  selectedImages: number[],
  formatFileSize: (size: number) => string
}>()
const emit = defineEmits(['preview', 'select'])
</script> 
