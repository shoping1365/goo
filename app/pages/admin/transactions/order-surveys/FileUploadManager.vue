<template>
  <div class="file-upload-manager bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-blue-400 to-blue-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">مدیریت آپلود فایل</h2>
        <p class="text-gray-600 mt-1">آپلود و مدیریت فایل‌های نظرسنجی با فشرده‌سازی خودکار</p>
      </div>
    </div>

    <!-- Upload Area -->
    <div 
      class="upload-area border-2 border-dashed border-gray-300 rounded-xl p-8 text-center mb-6 transition-all duration-300"
      :class="{
        'border-blue-500 bg-blue-50': isDragOver,
        'hover:border-blue-400 hover:bg-gray-50': !isDragOver
      }"
      @drop="handleDrop"
      @dragover="handleDragOver"
      @dragleave="handleDragLeave"
      @click="triggerFileInput"
    >
      <div class="upload-icon mb-4">
        <svg class="w-16 h-16 text-gray-400 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
        </svg>
      </div>
      
      <h3 class="text-lg font-medium text-gray-900 mb-2">فایل‌های خود را اینجا رها کنید</h3>
      <p class="text-gray-600 mb-4">یا برای انتخاب فایل کلیک کنید</p>
      
      <div class="text-sm text-gray-500 space-y-1">
        <p>پشتیبانی از: JPG, PNG, GIF, MP4, MOV (حداکثر ۱۰ مگابایت)</p>
        <p>فشرده‌سازی خودکار برای بهینه‌سازی حجم</p>
      </div>
      
      <input 
        ref="fileInput"
        type="file"
        multiple
        accept="image/*,video/*"
        class="hidden"
        @change="handleFileSelect"
      >
    </div>

    <!-- Upload Progress -->
    <div v-if="uploadingFiles.length > 0" class="mb-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">در حال آپلود</h3>
      <div class="space-y-3">
        <div v-for="file in uploadingFiles" :key="file.id" class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="p-2 bg-blue-500 rounded-lg">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
                </svg>
              </div>
              <div>
                <div class="text-sm font-medium text-gray-900">{{ file.name }}</div>
                <div class="text-xs text-gray-500">{{ formatFileSize(file.size) }}</div>
              </div>
            </div>
            <div class="text-sm text-gray-500">{{ file.progress }}%</div>
          </div>
          
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div 
              class="bg-blue-500 h-2 rounded-full transition-all duration-300"
              :style="{ width: `${file.progress}%` }"
            ></div>
          </div>
          
          <div v-if="file.status === 'compressing'" class="mt-2 text-xs text-blue-600">
            در حال فشرده‌سازی...
          </div>
          
          <div v-if="file.status === 'error'" class="mt-2 text-xs text-red-600">
            خطا: {{ file.error }}
          </div>
        </div>
      </div>
    </div>

    <!-- Uploaded Files -->
    <div v-if="uploadedFiles.length > 0" class="mb-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">فایل‌های آپلود شده</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="file in uploadedFiles" :key="file.id" class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="p-1 bg-green-500 rounded">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
              </div>
              <span class="text-sm font-medium text-gray-900">{{ file.name }}</span>
            </div>
            <button @click="removeFile(file.id)" class="text-red-600 hover:text-red-800">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
          
          <!-- File Preview -->
          <div class="mb-3">
            <img v-if="file.type.startsWith('image/')" :src="file.url" :alt="file.name" class="w-full h-32 object-cover rounded">
            <video v-else-if="file.type.startsWith('video/')" :src="file.url" class="w-full h-32 object-cover rounded" controls></video>
          </div>
          
          <div class="text-xs text-gray-500 space-y-1">
            <div>حجم اصلی: {{ formatFileSize(file.originalSize) }}</div>
            <div>حجم فشرده: {{ formatFileSize(file.compressedSize) }}</div>
            <div>کاهش حجم: {{ file.compressionRatio }}%</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Storage Info -->
    <div class="bg-blue-50 rounded-xl p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات ذخیره‌سازی</h3>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ storageInfo.usedSpace }}</div>
          <div class="text-sm text-gray-600">فضای استفاده شده</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ storageInfo.savedSpace }}</div>
          <div class="text-sm text-gray-600">فضای صرفه‌جویی شده</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-purple-600">{{ storageInfo.totalFiles }}</div>
          <div class="text-sm text-gray-600">تعداد کل فایل‌ها</div>
        </div>
      </div>
      
      <div class="mt-4">
        <div class="flex justify-between text-sm text-gray-600 mb-1">
          <span>فضای ذخیره‌سازی</span>
          <span>{{ storageInfo.usedSpace }} / {{ storageInfo.totalSpace }}</span>
        </div>
        <div class="w-full bg-gray-200 rounded-full h-2">
          <div 
            class="bg-blue-500 h-2 rounded-full transition-all duration-300"
            :style="{ width: `${storageInfo.usagePercentage}%` }"
          ></div>
        </div>
      </div>
    </div>

    <!-- Settings -->
    <div class="bg-gray-50 rounded-xl p-6 mt-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات آپلود</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر حجم فایل (MB)</label>
          <input 
            v-model.number="settings.maxFileSize" 
            type="number" 
            min="1" 
            max="50"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          >
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کیفیت فشرده‌سازی</label>
          <select v-model="settings.compressionQuality" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
            <option value="high">بالا (کیفیت بهتر)</option>
            <option value="medium">متوسط (تعادل)</option>
            <option value="low">پایین (حجم کمتر)</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرمت‌های مجاز</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="settings.allowedFormats.images" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">تصاویر (JPG, PNG, GIF)</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="settings.allowedFormats.videos" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">ویدیوها (MP4, MOV)</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="settings.allowedFormats.documents" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">اسناد (PDF, DOC)</span>
            </label>
          </div>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی خودکار</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="settings.autoCompress" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فشرده‌سازی خودکار فایل‌های بزرگ</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="settings.keepOriginal" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">نگهداری نسخه اصلی</span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface UploadingFile {
  id: string
  name: string
  size: number
  type: string
  progress: number
  status: 'uploading' | 'compressing' | 'completed' | 'error'
  error?: string
}

interface UploadedFile {
  id: string
  name: string
  type: string
  url: string
  originalSize: number
  compressedSize: number
  compressionRatio: number
  uploadedAt: Date
}

interface StorageInfo {
  usedSpace: string
  totalSpace: string
  savedSpace: string
  totalFiles: number
  usagePercentage: number
}

interface UploadSettings {
  maxFileSize: number
  compressionQuality: 'high' | 'medium' | 'low'
  allowedFormats: {
    images: boolean
    videos: boolean
    documents: boolean
  }
  autoCompress: boolean
  keepOriginal: boolean
}

const fileInput = ref<HTMLInputElement | null>(null)
const isDragOver = ref(false)
const uploadingFiles = ref<UploadingFile[]>([])
const uploadedFiles = ref<UploadedFile[]>([])

const settings = ref<UploadSettings>({
  maxFileSize: 10,
  compressionQuality: 'medium',
  allowedFormats: {
    images: true,
    videos: true,
    documents: false
  },
  autoCompress: true,
  keepOriginal: false
})

const storageInfo = ref<StorageInfo>({
  usedSpace: '۲.۵ GB',
  totalSpace: '۱۰ GB',
  savedSpace: '۱.۲ GB',
  totalFiles: 1247,
  usagePercentage: 25
})

// Methods
const triggerFileInput = () => {
  fileInput.value?.click()
}

const handleDragOver = (event: DragEvent) => {
  event.preventDefault()
  isDragOver.value = true
}

const handleDragLeave = (event: DragEvent) => {
  event.preventDefault()
  isDragOver.value = false
}

const handleDrop = (event: DragEvent) => {
  event.preventDefault()
  isDragOver.value = false
  
  const files = event.dataTransfer?.files
  if (files) {
    handleFiles(Array.from(files))
  }
}

const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    handleFiles(Array.from(target.files))
  }
}

const handleFiles = (files: File[]) => {
  files.forEach(file => {
    if (validateFile(file)) {
      uploadFile(file)
    }
  })
}

const validateFile = (file: File): boolean => {
  // Check file size
  if (file.size > settings.value.maxFileSize * 1024 * 1024) {
    alert(`فایل ${file.name} بزرگتر از حد مجاز است (${settings.value.maxFileSize}MB)`)
    return false
  }
  
  // Check file type
  const allowedTypes = []
  if (settings.value.allowedFormats.images) {
    allowedTypes.push('image/')
  }
  if (settings.value.allowedFormats.videos) {
    allowedTypes.push('video/')
  }
  if (settings.value.allowedFormats.documents) {
    allowedTypes.push('application/pdf', 'application/msword')
  }
  
  const isValidType = allowedTypes.some(type => file.type.startsWith(type))
  if (!isValidType) {
    alert(`نوع فایل ${file.name} پشتیبانی نمی‌شود`)
    return false
  }
  
  return true
}

const uploadFile = async (file: File) => {
  const uploadingFile: UploadingFile = {
    id: Date.now().toString(),
    name: file.name,
    size: file.size,
    type: file.type,
    progress: 0,
    status: 'uploading'
  }
  
  uploadingFiles.value.push(uploadingFile)
  
  try {
    // Simulate upload progress
    for (let i = 0; i <= 100; i += 10) {
      uploadingFile.progress = i
      await new Promise(resolve => setTimeout(resolve, 100))
    }
    
    // Simulate compression
    if (settings.value.autoCompress && file.size > 1024 * 1024) {
      uploadingFile.status = 'compressing'
      await new Promise(resolve => setTimeout(resolve, 1000))
    }
    
    uploadingFile.status = 'completed'
    
    // Add to uploaded files
    const uploadedFile: UploadedFile = {
      id: uploadingFile.id,
      name: file.name,
      type: file.type,
      url: URL.createObjectURL(file),
      originalSize: file.size,
      compressedSize: Math.round(file.size * 0.7), // Simulate compression
      compressionRatio: 30,
      uploadedAt: new Date()
    }
    
    uploadedFiles.value.push(uploadedFile)
    
    // Remove from uploading files
    const index = uploadingFiles.value.findIndex(f => f.id === uploadingFile.id)
    if (index > -1) {
      uploadingFiles.value.splice(index, 1)
    }
    
  } catch (error) {
    uploadingFile.status = 'error'
    uploadingFile.error = 'خطا در آپلود فایل'
  }
}

const removeFile = (fileId: string) => {
  const index = uploadedFiles.value.findIndex(f => f.id === fileId)
  if (index > -1) {
    uploadedFiles.value.splice(index, 1)
  }
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '۰ بایت'
  
  const k = 1024
  const sizes = ['بایت', 'کیلوبایت', 'مگابایت', 'گیگابایت']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// Expose methods for parent component
defineExpose({
  uploadedFiles,
  settings,
  uploadFile,
  removeFile
})
</script>

<style scoped>
.file-upload-manager {
  transition: all 0.3s ease;
}

.upload-area {
  cursor: pointer;
  transition: all 0.3s ease;
}

.upload-area:hover {
  transform: translateY(-2px);
}

/* Progress bar animation */
.bg-blue-500 {
  transition: width 0.3s ease-out;
}

/* File preview hover effects */
.uploaded-files img:hover,
.uploaded-files video:hover {
  transform: scale(1.05);
  transition: transform 0.2s ease;
}

/* Drag and drop animations */
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.7;
  }
}

.border-blue-500 {
  animation: pulse 2s infinite;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .grid {
    grid-template-columns: 1fr;
  }
  
  .upload-area {
    padding: 2rem 1rem;
  }
}
</style> 
