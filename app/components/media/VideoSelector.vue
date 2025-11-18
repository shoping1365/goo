<template>
  <div class="video-selector">
    <div class="mb-4">
      <h3 class="text-lg font-semibold text-gray-800 mb-2">انتخاب ویدیو</h3>
      <div class="flex gap-2">
        <button 
          @click="openFileDialog" 
          class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
        >
          انتخاب فایل
        </button>
        <input 
          ref="fileInput" 
          type="file" 
          accept="video/*" 
          class="hidden" 
          @change="handleFileSelect"
        />
      </div>
    </div>

    <div v-if="selectedVideo" class="border rounded-lg p-6 bg-gray-50">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3 space-x-reverse">
          <div class="w-16 h-16 bg-gray-200 rounded-lg flex items-center justify-center">
            <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
          </div>
          <div>
            <h4 class="font-medium text-gray-800">{{ selectedVideo.name }}</h4>
            <p class="text-sm text-gray-600">{{ formatFileSize(selectedVideo.size) }}</p>
          </div>
        </div>
        
        <div class="flex space-x-2 space-x-reverse">
          <button 
            @click="previewVideo"
            class="px-3 py-1 bg-green-500 text-white rounded text-sm hover:bg-green-600 transition-colors"
          >
            پیش‌نمایش
          </button>
          <button 
            @click="removeVideo"
            class="px-3 py-1 bg-red-500 text-white rounded text-sm hover:bg-red-600 transition-colors"
          >
            حذف
          </button>
        </div>
      </div>
    </div>

    <div v-if="showPreview" class="mt-4">
      <VideoPreviewModal
        :is-visible="showPreview"
        :video-src="previewUrl"
        :title="selectedVideo?.name || 'پیش‌نمایش ویدیو'"
        :file-name="selectedVideo?.name || ''"
        :file-size="selectedVideo ? formatFileSize(selectedVideo.size) : ''"
        @close="closePreview"
      />
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(['select', 'remove'])

const fileInput = ref(null)
const selectedVideo = ref(null)
const showPreview = ref(false)
const previewUrl = ref('')

const openFileDialog = () => {
  fileInput.value?.click()
}

const handleFileSelect = (event) => {
  const file = event.target.files[0]
  if (file && file.type.startsWith('video/')) {
    selectedVideo.value = file
    previewUrl.value = URL.createObjectURL(file)
    emit('select', file)
  }
}

const removeVideo = () => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
  }
  selectedVideo.value = null
  previewUrl.value = ''
  emit('remove')
}

const previewVideo = () => {
  if (selectedVideo.value) {
    showPreview.value = true
  }
}

const closePreview = () => {
  showPreview.value = false
}

const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

onUnmounted(() => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
  }
})
</script> 