<template>
  <div class="file-uploader">
    <!-- File Upload Modal -->
    <div v-if="showModal" class="upload-modal-overlay" @click.self="closeModal">
      <div class="upload-modal">
        <div class="modal-header">
          <h3>📎 ارسال فایل</h3>
          <button class="close-btn" @click="closeModal">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <div 
          class="drop-zone" 
          :class="{ 'drag-over': isDragOver }"
          @drop="handleDrop"
          @dragover.prevent="isDragOver = true"
          @dragleave="isDragOver = false"
          @click="$refs.fileInput.click()"
        >
          <div class="drop-icon">
            <svg v-if="!isDragOver" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-12 h-12">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5" />
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-12 h-12 animate-bounce">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m.75 12l3 3m0 0l3-3m-3 3v-6m-1.5-9H5.25A2.25 2.25 0 003 5.25v13.5A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V5.25a2.25 2.25 0 00-2.25-2.25H9.75z" />
            </svg>
          </div>
          <h4>{{ isDragOver ? 'فایل را رها کنید!' : 'فایل خود را اینجا بکشید' }}</h4>
          <p>یا کلیک کنید تا فایل انتخاب کنید</p>
          <div class="supported-formats">
            <span>پشتیبانی:</span>
            <span class="format">تصاویر</span>
            <span class="format">اسناد</span>
            <span class="format">ویدیو</span>
            <span class="format">صوت</span>
          </div>
        </div>
        
        <input
          ref="fileInput"
          type="file"
          multiple
          hidden
          accept="image/*,video/*,audio/*,.pdf,.doc,.docx,.txt"
          @change="handleFileSelect"
        />
        
        <!-- Selected Files -->
        <div v-if="selectedFiles.length > 0" class="selected-files">
          <h4>فایل‌های انتخاب شده:</h4>
          <div class="files-list">
            <div v-for="(file, index) in selectedFiles" :key="index" class="file-item">
              <div class="file-icon">{{ getFileIcon(file.type) }}</div>
              <div class="file-info">
                <span class="file-name">{{ file.name }}</span>
                <span class="file-size">{{ formatFileSize(file.size) }}</span>
              </div>
              <button 
                v-if="canDeleteChatFile"
                class="remove-file-btn" 
                @click="removeFile(index)"
              >
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>
          
          <div class="upload-actions">
            <button class="clear-btn" @click="clearFiles">پاک کردن همه</button>
            <button class="upload-btn" :disabled="isUploading" @click="uploadFiles">
              <span v-if="!isUploading">📤 ارسال فایل‌ها</span>
              <span v-else>در حال ارسال...</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'


const { hasPermission } = useAuth()

const canDeleteChatFile = computed(() => hasPermission('chat.delete'))

defineProps({
  show: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'upload'])

const showModal = computed(() => props.show)
const isDragOver = ref(false)
const selectedFiles = ref([])
const isUploading = ref(false)

const closeModal = () => {
  selectedFiles.value = []
  emit('close')
}

const handleDrop = (e) => {
  e.preventDefault()
  isDragOver.value = false
  
  const files = Array.from(e.dataTransfer.files)
  addFiles(files)
}

const handleFileSelect = (e) => {
  const files = Array.from(e.target.files)
  addFiles(files)
}

const addFiles = (files) => {
  files.forEach(file => {
    if (file.size <= 10 * 1024 * 1024) { // 10MB limit
      selectedFiles.value.push(file)
    } else {
      alert(`فایل ${file.name} بیش از 10 مگابایت است!`)
    }
  })
}

const removeFile = (index) => {
  selectedFiles.value.splice(index, 1)
}

const clearFiles = () => {
  selectedFiles.value = []
}

const uploadFiles = async () => {
  if (selectedFiles.value.length === 0) return
  
  isUploading.value = true
  
  try {
    // Simulate upload delay
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    const uploadedFiles = selectedFiles.value.map(file => ({
      name: file.name,
      size: file.size,
      type: file.type,
      url: URL.createObjectURL(file), // In real app, this would be server URL
      uploadTime: new Date().toLocaleTimeString('fa-IR')
    }))
    
    emit('upload', uploadedFiles)
    selectedFiles.value = []
    closeModal()
    
  } catch (error) {
    console.error('Upload failed:', error)
    alert('خطا در آپلود فایل!')
  } finally {
    isUploading.value = false
  }
}

const getFileIcon = (type) => {
  if (type.startsWith('image/')) return '🖼️'
  if (type.startsWith('video/')) return '🎥'
  if (type.startsWith('audio/')) return '🎵'
  if (type.includes('pdf')) return '📄'
  if (type.includes('word') || type.includes('doc')) return '📝'
  return '📎'
}

const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>

<style scoped>
.upload-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(4px);
}

.upload-modal {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 16px;
  width: 90%;
  max-width: 500px;
  max-height: 80vh;
  overflow-y: auto;
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.modal-header h3 {
  margin: 0;
  color: #2d3748;
  font-size: 18px;
  font-weight: 600;
}

.close-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #718096;
  padding: 4px;
  border-radius: 50%;
  transition: all 0.3s ease;
}

.close-btn:hover {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
}

.drop-zone {
  margin: 24px;
  padding: 40px 20px;
  border: 2px dashed #cbd5e0;
  border-radius: 12px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.03), rgba(118, 75, 162, 0.03));
}

.drop-zone:hover,
.drop-zone.drag-over {
  border-color: #667eea;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1), rgba(118, 75, 162, 0.1));
  transform: scale(1.02);
}

.drop-icon {
  margin-bottom: 16px;
  color: #667eea;
}

.drop-zone h4 {
  margin: 0 0 8px 0;
  color: #2d3748;
  font-size: 16px;
  font-weight: 600;
}

.drop-zone p {
  margin: 0 0 16px 0;
  color: #718096;
  font-size: 14px;
}

.supported-formats {
  display: flex;
  gap: 8px;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
}

.supported-formats span {
  font-size: 12px;
  color: #a0aec0;
}

.format {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea !important;
  padding: 2px 8px;
  border-radius: 12px;
  font-weight: 500;
}

.selected-files {
  padding: 24px;
  border-top: 1px solid rgba(102, 126, 234, 0.1);
}

.selected-files h4 {
  margin: 0 0 16px 0;
  color: #2d3748;
  font-size: 14px;
  font-weight: 600;
}

.files-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.file-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: rgba(102, 126, 234, 0.05);
  border-radius: 8px;
  border: 1px solid rgba(102, 126, 234, 0.1);
}

.file-icon {
  font-size: 20px;
  flex-shrink: 0;
}

.file-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.file-name {
  font-size: 13px;
  font-weight: 500;
  color: #2d3748;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.file-size {
  font-size: 11px;
  color: #718096;
}

.remove-file-btn {
  background: rgba(255, 107, 107, 0.1);
  border: none;
  border-radius: 50%;
  padding: 6px;
  cursor: pointer;
  color: #ff6b6b;
  transition: all 0.3s ease;
}

.remove-file-btn:hover {
  background: rgba(255, 107, 107, 0.2);
  transform: scale(1.1);
}

.upload-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.clear-btn {
  padding: 8px 16px;
  background: rgba(255, 107, 107, 0.1);
  color: #ff6b6b;
  border: 1px solid rgba(255, 107, 107, 0.2);
  border-radius: 8px;
  cursor: pointer;
  font-size: 13px;
  transition: all 0.3s ease;
}

.clear-btn:hover {
  background: rgba(255, 107, 107, 0.2);
}

.upload-btn {
  padding: 8px 20px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 13px;
  font-weight: 500;
  transition: all 0.3s ease;
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.3);
}

.upload-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.4);
}

.upload-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}
</style> 
