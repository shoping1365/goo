<template>
  <div class="media-library">
    <div class="mb-4">
      <h3 class="text-lg font-semibold text-gray-800 mb-2">کتابخانه رسانه</h3>
      <div class="flex gap-2">
        <button 
          class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors" 
          @click="uploadFile"
        >
          آپلود فایل
        </button>
        <input 
          ref="fileInput" 
          type="file" 
          multiple 
          accept="image/*,video/*" 
          class="hidden" 
          @change="handleFileSelect"
        />
      </div>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-6">
      <div 
        v-for="item in mediaItems" 
        :key="item.id" 
        class="relative group cursor-pointer"
        @click="selectItem(item)"
      >
        <div class="aspect-square bg-gray-100 rounded-lg overflow-hidden">
          <img 
            v-if="item.type === 'image'" 
            :src="item.url" 
            :alt="item.name"
            class="w-full h-full object-cover"
          />
          <div 
            v-else 
            class="w-full h-full flex items-center justify-center bg-gray-200"
          >
            <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
        
        <div class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-50 transition-all duration-200 flex items-center justify-center">
          <button 
            class="text-white bg-red-500 rounded-full p-1 hover:bg-red-600 transition-colors opacity-0 group-hover:opacity-100"
            @click.stop="deleteItem(item?.id || 0)"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
            </svg>
          </button>
        </div>
        
        <p class="text-xs text-gray-600 mt-1 truncate">{{ item.name }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(['select', 'delete'])

const fileInput = ref(null)
const mediaItems = ref([])

const uploadFile = () => {
  fileInput.value?.click()
}

const handleFileSelect = (event) => {
  const files = Array.from(event.target.files)
  files.forEach(file => {
    const item = {
      id: Date.now() + Math.random(),
      name: file.name,
      type: file.type.startsWith('image/') ? 'image' : 'video',
      url: URL.createObjectURL(file),
      file: file
    }
    mediaItems.value.push(item)
  })
}

const selectItem = (item) => {
  emit('select', item)
}

const deleteItem = (id) => {
  const index = mediaItems.value.findIndex(item => item.id === id)
  if (index > -1) {
    mediaItems.value.splice(index, 1)
    emit('delete', id)
  }
}
</script> 
