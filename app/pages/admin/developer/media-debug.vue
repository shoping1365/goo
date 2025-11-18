<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">دیباگ پیشرفته رسانه</h1>
        <p class="text-gray-600">ابزارهای مدیریت و دیباگ فایل‌های رسانه</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">کل فایل‌ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.totalFiles }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">حجم کل</p>
              <p class="text-2xl font-semibold text-gray-900">{{ formatFileSize(stats.totalSize) }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">فایل‌های مشکل‌دار</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.corruptedFiles }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">فایل‌های بهینه‌سازی شده</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.optimizedFiles }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- File Upload & Management -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">آپلود و مدیریت</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <!-- Upload Area -->
              <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
                <svg class="w-12 h-12 mx-auto text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
                </svg>
                <p class="text-gray-600 mb-2">فایل‌ها را اینجا بکشید یا کلیک کنید</p>
                <input 
                  type="file" 
                  multiple 
                  @change="handleFileUpload"
                  class="hidden"
                  ref="fileInput"
                  accept="image/*,video/*,audio/*"
                >
                <button 
                  @click="$refs.fileInput.click()"
                  class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  انتخاب فایل
                </button>
              </div>

              <!-- Upload Progress -->
              <div v-if="uploadProgress.length > 0" class="space-y-2">
                <h3 class="font-medium text-gray-700">پیشرفت آپلود</h3>
                <div v-for="file in uploadProgress" :key="file.name" class="bg-gray-50 rounded p-3">
                  <div class="flex items-center justify-between mb-2">
                    <span class="text-sm font-medium">{{ file.name }}</span>
                    <span class="text-sm text-gray-600">{{ file.progress }}%</span>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div class="bg-blue-600 h-2 rounded-full" :style="{ width: file.progress + '%' }"></div>
                  </div>
                </div>
              </div>

              <!-- Actions -->
              <div class="space-y-2">
                <button 
                  @click="refreshFileList"
                  class="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  بارگذاری مجدد لیست
                </button>
                <button 
                  @click="optimizeAllFiles"
                  :disabled="optimizing"
                  class="w-full bg-purple-600 hover:bg-purple-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  <span v-if="optimizing">در حال بهینه‌سازی...</span>
                  <span v-else>بهینه‌سازی همه فایل‌ها</span>
                </button>
                <button 
                  @click="cleanupOrphanedFiles"
                  class="w-full bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                  پاکسازی فایل‌های یتیم
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- File List -->
        <div class="lg:col-span-2 bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">لیست فایل‌ها</h2>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input 
                  v-model="searchQuery"
                  type="text"
                  placeholder="جستجو در فایل‌ها..."
                  class="border border-gray-300 rounded-lg px-3 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <select 
                  v-model="filterType"
                  class="border border-gray-300 rounded-lg px-3 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="">همه انواع</option>
                  <option value="image">تصاویر</option>
                  <option value="video">ویدیوها</option>
                  <option value="audio">فایل‌های صوتی</option>
                </select>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="file in filteredFiles" :key="file.id" class="border rounded-lg p-6 hover:bg-gray-50">
                <div class="flex items-center justify-between">
                  <div class="flex items-center space-x-4 space-x-reverse">
                    <!-- File Icon -->
                    <div class="w-12 h-12 bg-gray-100 rounded-lg flex items-center justify-center">
                      <svg v-if="file.type.startsWith('image')" class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                      </svg>
                      <svg v-else-if="file.type.startsWith('video')" class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                      </svg>
                      <svg v-else class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                      </svg>
                    </div>

                    <!-- File Info -->
                    <div>
                      <h3 class="font-medium text-gray-900">{{ file.name }}</h3>
                      <p class="text-sm text-gray-600">{{ file.type }} - {{ formatFileSize(file.size) }}</p>
                      <p class="text-xs text-gray-500">{{ file.path }}</p>
                    </div>
                  </div>

                  <!-- Actions -->
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      @click="previewFile(file)"
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                    >
                      پیش‌نمایش
                    </button>
                    <button 
                      @click="optimizeFile(file)"
                      class="text-purple-600 hover:text-purple-800 text-sm font-medium"
                    >
                      بهینه‌سازی
                    </button>
                    <button 
                      @click="deleteFile(file)"
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                    >
                      حذف
                    </button>
                  </div>
                </div>

                <!-- File Status -->
                <div class="mt-3 flex items-center space-x-4 space-x-reverse">
                  <span :class="[
                    'px-2 py-1 rounded text-xs font-medium',
                    file.status === 'optimized' ? 'bg-green-100 text-green-800' : 
                    file.status === 'corrupted' ? 'bg-red-100 text-red-800' : 'bg-yellow-100 text-yellow-800'
                  ]">
                    {{ getStatusText(file.status) }}
                  </span>
                  <span class="text-xs text-gray-500">آپلود شده در {{ formatDate(file.uploadedAt) }}</span>
                </div>
              </div>

              <!-- Empty State -->
              <div v-if="filteredFiles.length === 0" class="text-center py-8">
                <svg class="w-12 h-12 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
                <p class="text-gray-500">فایلی برای نمایش وجود ندارد</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- File Preview Modal -->
      <div v-if="previewModal.show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-4xl max-h-4xl overflow-auto">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-semibold text-gray-900">{{ previewModal.file?.name }}</h3>
              <button @click="previewModal.show = false" class="text-gray-400 hover:text-gray-600">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="p-6">
            <img v-if="previewModal.file?.type.startsWith('image')" :src="previewModal.file?.url" :alt="previewModal.file?.name" class="max-w-full h-auto">
            <video v-else-if="previewModal.file?.type.startsWith('video')" :src="previewModal.file?.url" controls class="max-w-full h-auto"></video>
            <audio v-else-if="previewModal.file?.type.startsWith('audio')" :src="previewModal.file?.url" controls class="w-full"></audio>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Reactive data
const files = ref([])
const uploadProgress = ref([])
const optimizing = ref(false)
const searchQuery = ref('')
const filterType = ref('')

const previewModal = reactive({
  show: false,
  file: null
})

const stats = reactive({
  totalFiles: 0,
  totalSize: 0,
  corruptedFiles: 0,
  optimizedFiles: 0
})

// Computed
const filteredFiles = computed(() => {
  let filtered = files.value

  if (searchQuery.value) {
    filtered = filtered.filter(file => 
      file.name.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterType.value) {
    filtered = filtered.filter(file => file.type.startsWith(filterType.value))
  }

  return filtered
})

// Methods
async function fetchMediaList() {
  try {
    const response = await $fetch('/api/media/list')
    files.value = response.files || []
    updateStats()
  } catch (error) {
    console.error('خطا در دریافت لیست رسانه:', error)
  }
}

function updateStats() {
  stats.totalFiles = files.value.length
  stats.totalSize = files.value.reduce((sum, file) => sum + file.size, 0)
  stats.corruptedFiles = files.value.filter(file => file.status === 'corrupted').length
  stats.optimizedFiles = files.value.filter(file => file.status === 'optimized').length
}

function handleFileUpload(event) {
  const selectedFiles = Array.from(event.target.files)
  
  selectedFiles.forEach(file => {
    const progress = {
      name: file.name,
      progress: 0,
      status: 'uploading'
    }
    
    uploadProgress.value.push(progress)
    
    // Simulate upload progress
    const interval = setInterval(() => {
      progress.progress += Math.random() * 20
      if (progress.progress >= 100) {
        progress.progress = 100
        progress.status = 'completed'
        clearInterval(interval)
        
        // Add to files list
        files.value.unshift({
          id: Date.now(),
          name: file.name,
          type: file.type,
          size: file.size,
          path: `/uploads/${file.name}`,
          url: URL.createObjectURL(file),
          status: 'pending',
          uploadedAt: new Date()
        })
        
        updateStats()
        
        // Remove from progress after delay
        setTimeout(() => {
          const index = uploadProgress.value.indexOf(progress)
          if (index > -1) {
            uploadProgress.value.splice(index, 1)
          }
        }, 2000)
      }
    }, 200)
  })
}

async function optimizeFile(file) {
  try {
    file.status = 'optimizing'
    // Simulate optimization
    await new Promise(resolve => setTimeout(resolve, 2000))
    file.status = 'optimized'
    updateStats()
  } catch (error) {
    file.status = 'error'
    console.error('خطا در بهینه‌سازی فایل:', error)
  }
}

async function optimizeAllFiles() {
  optimizing.value = true
  try {
    const pendingFiles = files.value.filter(file => file.status !== 'optimized')
    for (const file of pendingFiles) {
      await optimizeFile(file)
    }
  } finally {
    optimizing.value = false
  }
}

async function deleteFile(file) {
  if (confirm(`آیا از حذف فایل "${file.name}" اطمینان دارید؟`)) {
    try {
      const index = files.value.indexOf(file)
      if (index > -1) {
        files.value.splice(index, 1)
        updateStats()
      }
    } catch (error) {
      console.error('خطا در حذف فایل:', error)
    }
  }
}

function previewFile(file) {
  previewModal.file = file
  previewModal.show = true
}

async function cleanupOrphanedFiles() {
  if (confirm('آیا از پاکسازی فایل‌های یتیم اطمینان دارید؟')) {
    try {
      // Simulate cleanup
      await new Promise(resolve => setTimeout(resolve, 1000))
      alert('فایل‌های یتیم با موفقیت پاکسازی شدند')
    } catch (error) {
      console.error('خطا در پاکسازی فایل‌ها:', error)
    }
  }
}

function refreshFileList() {
  fetchMediaList()
}

function formatFileSize(bytes) {
  if (bytes === 0) return '0 بایت'
  const k = 1024
  const sizes = ['بایت', 'کیلوبایت', 'مگابایت', 'گیگابایت']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

function formatDate(date) {
  return new Date(date).toLocaleString('fa-IR')
}

function getStatusText(status) {
  const statusMap = {
    'pending': 'در انتظار',
    'optimized': 'بهینه‌سازی شده',
    'corrupted': 'خراب',
    'error': 'خطا'
  }
  return statusMap[status] || status
}

// Load initial data
onMounted(() => {
  fetchMediaList()
})
</script>

<style scoped>
.media-debug {
  padding: 2rem;
}

.debug-controls {
  margin: 2rem 0;
}

.debug-button {
  padding: 0.75rem 1.5rem;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}

.debug-button:hover {
  background-color: #45a049;
}

.media-list {
  display: grid;
  gap: 1rem;
}

.file-item {
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background-color: #f9f9f9;
}

.file-info {
  display: grid;
  gap: 0.5rem;
}

.file-info p {
  margin: 0;
  font-family: monospace;
}
</style> 
