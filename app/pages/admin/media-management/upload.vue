<template>
  <div class="min-h-screen flex flex-col">
    <div class="w-full bg-white shadow-lg border-b border-blue-200 sticky top-0 z-10 px-8 py-6 flex flex-col gap-1">
      <div class="flex flex-row-reverse items-center justify-between w-full">
        <NuxtLink
            to="/admin/media-management/library"
            class="bg-gradient-to-l from-blue-600 to-blue-400 text-white px-5 py-2 rounded-xl shadow hover:from-blue-700 hover:to-blue-500 transition-colors flex items-center gap-2 font-medium text-base"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          بازگشت به کتابخانه
        </NuxtLink>
        <h1 class="text-2xl font-bold text-gray-900">افزودن رسانه جدید</h1>
      </div>
      <p class="text-xs text-gray-600 mt-1 w-full">آپلود فایل‌ها، تصاویر و ویدیوها</p>
    </div>
    <div class="p-8 w-full flex flex-col items-center">
      <div class="w-full max-w-3xl mx-auto">
        <div class="bg-white rounded-2xl shadow-2xl border-2 border-blue-200 p-10 hover:shadow-2xl transition-shadow duration-300 w-full min-h-[350px]">
          <h2 class="text-lg font-semibold text-gray-900 mb-4 text-center">آپلود با کشیدن و رها کردن</h2>
          <div
              :class="[
              'border-2 border-dashed rounded-xl flex flex-col items-center justify-center p-8 text-center transition-all duration-300 shadow-inner',
              isDragOver
                ? 'border-blue-500 bg-blue-50 shadow-lg'
                : 'border-blue-200 hover:border-blue-400 hover:shadow-md',
              'min-h-[300px]'
            ]"
              class="mx-auto w-full max-w-full"
              @drop="handleDrop"
              @dragover.prevent
              @dragenter.prevent
              @dragenter="isDragOver = true"
              @dragleave="isDragOver = false"
          >
            <svg class="w-16 h-16 text-blue-400 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
            </svg>
            <p class="text-lg font-bold text-gray-900 mb-2">فایل‌ها را اینجا بکشید</p>
            <p class="text-xs text-gray-600 mb-3">یا کلیک کنید تا فایل انتخاب کنید</p>
            <input
                ref="fileInput"
                type="file"
                multiple
                class="hidden"
                accept="image/*,video/*,audio/*,.pdf,.doc,.docx,.txt"
                @change="handleFileSelect"
            >
            <button
                class="px-6 py-2 rounded-lg font-bold bg-gradient-to-br from-cyan-200 to-green-200 text-cyan-900 border border-cyan-300 shadow-sm hover:from-cyan-300 hover:to-green-300 hover:text-cyan-950 transition-all duration-200 font-medium text-base mt-2"
                @click="($refs.fileInput as HTMLInputElement)?.click()"
            >
              انتخاب فایل‌ها
            </button>
            <div class="flex items-center justify-center mt-6">
              <label class="flex items-center gap-2 px-4 py-2 rounded-xl bg-blue-50 border border-blue-200 text-blue-700 cursor-pointer shadow-sm hover:bg-blue-100 transition-all">
                <input v-model="autoCompress" type="checkbox" class="accent-blue-500 w-5 h-5" />
                <span class="font-medium">فشرده‌سازی هوشمند</span>
              </label>
            </div>
            <p class="text-xs text-gray-500 mt-4">
              فرمت‌های مجاز: JPG, PNG, GIF, MP4, MP3, PDF, DOC, TXT (حداکثر 1GB)
            </p>
          </div>
        </div>
        <div class="w-full h-auto py-4 bg-gradient-to-l from-green-200 to-blue-200 rounded-full shadow-md flex flex-wrap items-center justify-center gap-8 mt-8">
          <label class="flex items-center gap-2 text-base text-gray-700">
            <input v-model="selectedCategory" type="radio" value="library" class="accent-green-500 w-5 h-5" />
            کتابخانه
          </label>
          <label class="flex items-center gap-2 text-base text-gray-700">
            <input v-model="selectedCategory" type="radio" value="products" class="accent-blue-500 w-5 h-5" />
            محصولات
          </label>
          <label class="flex items-center gap-2 text-base text-gray-700">
            <input v-model="selectedCategory" type="radio" value="product-categories" class="accent-purple-500 w-5 h-5" />
            دسته‌بندی محصولات
          </label>
          <label class="flex items-center gap-2 text-base text-gray-700">
            <input v-model="selectedCategory" type="radio" value="brands" class="accent-pink-500 w-5 h-5" />
            برندها
          </label>
          <label class="flex items-center gap-2 text-base text-gray-700">
            <input v-model="selectedCategory" type="radio" value="banners" class="accent-yellow-500 w-5 h-5" />
            بنرهای تبلیغاتی
          </label>
        </div>
      </div>

      <div v-if="uploadQueue.length > 0" class="bg-white rounded-2xl shadow-2xl border-2 border-blue-200 overflow-hidden mb-8 w-full mt-8">
        <div class="px-8 py-6 border-b border-blue-100 flex items-center justify-between gap-6">
          <h2 class="text-lg font-semibold text-gray-900">فایل‌های در حال آپلود</h2>
          <div v-if="uploadQueue.length" class="flex flex-col gap-1">
            <!-- Current file progress -->
            <div class="relative w-[800px] bg-gray-200 rounded-full h-3 overflow-hidden rtl:ml-4 ltr:mr-4">
              <div :style="{ width: currentFileProgress + '%' }" class="absolute left-0 top-0 bg-green-500 h-full transition-all duration-200"></div>
              <span class="absolute right-2 top-1/2 -translate-y-1/2 text-xs font-semibold text-green-600 rtl:left-2 rtl:right-auto">{{ currentFileProgress }}٪</span>
            </div>
            <!-- Overall progress -->
            <div class="relative w-[800px] bg-gray-200 rounded-full h-3 overflow-hidden rtl:ml-4 ltr:mr-4 mt-2">
              <div :style="{ width: totalProgress + '%' }" class="absolute left-0 top-0 bg-blue-500 h-full transition-all duration-200"></div>
              <span class="absolute right-2 top-1/2 -translate-y-1/2 text-xs font-semibold text-blue-600 rtl:left-2 rtl:right-auto">{{ totalProgress }}٪</span>
            </div>
          </div>
        </div>
        <div class="grid gap-6 p-6 [grid-template-columns:repeat(auto-fit,minmax(220px,1fr))]">
          <div v-for="file in uploadQueue" :key="file.id" class="rounded-xl bg-blue-50 border border-blue-200 shadow flex flex-col gap-2 p-6 relative text-right min-w-0 w-full max-w-sm min-w-[200px] max-w-[320px]">
            <div class="flex items-center justify-between mb-2 flex-row-reverse">
              <div class="font-bold text-base text-gray-800 truncate">{{ file.name }}</div>
              <div class="flex items-center gap-2">
                <!-- Thumbnail + Eye Icon -->
                <div class="relative w-[50px] h-[50px] mr-2 flex items-center justify-center">
                  <img v-if="file.type.startsWith('image') && file.preview" :src="file.thumbnail || file.preview" alt="thumbnail" class="w-[50px] h-[50px] object-cover rounded border" />
                  <video v-else-if="file.type.startsWith('video') && file.preview" :src="file.preview" class="w-[50px] h-[50px] object-cover rounded border" muted></video>
                  <div v-else class="w-[50px] h-[50px] bg-gray-200 flex items-center justify-center rounded border">
                    <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path></svg>
                  </div>
                  <button class="absolute left-0 top-0 bg-white bg-opacity-80 rounded-full p-1 shadow hover:bg-opacity-100 transition-all" style="z-index:2;" @click="openImagePreview(file)">
                    <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                </div>
              <div class="text-xs text-gray-500">KB {{ (file.size / 1024).toLocaleString(undefined, {maximumFractionDigits:2}) }}</div>
              </div>
            </div>
            <div class="grid grid-cols-2 gap-6 mb-2">
              <div>
                <label class="block text-gray-600 mb-1">متن جایگزین (Alt):</label>
                <input v-model="file.alt" type="text" placeholder="توضیح تصویر" class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500" :disabled="file.status !== 'completed'" @input="markDirty(file)" @blur="handleSave(file)">
              </div>
              <div>
                <label class="block text-gray-600 mb-1">عنوان:</label>
                <input v-model="file.title" type="text" placeholder="عنوان فایل" class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500" :disabled="file.status !== 'completed'" @input="markDirty(file)" @blur="handleSave(file)">
              </div>
              <div class="col-span-2">
                <label class="block text-gray-600 mb-1">توضیحات مختصر (caption):</label>
                <input v-model="file.caption" type="text" placeholder="توضیحات مختصر فایل" class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500" :disabled="file.status !== 'completed'" @input="markDirty(file)" @blur="handleSave(file)">
              </div>
              <div class="col-span-2">
                <label class="block text-gray-600 mb-1">توضیحات:</label>
                <textarea v-model="file.description" placeholder="توضیحات فایل" rows="2" class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500" :disabled="file.status !== 'completed'" @input="markDirty(file)" @blur="handleSave(file)"></textarea>
              </div>
            </div>
            <!-- Per file progress bar -->
            <div class="relative w-full bg-gray-200 rounded-full h-3 overflow-hidden">
              <div :style="{ width: file.progress + '%' }" :class="['absolute right-0 top-0 h-full rounded-full transition-all duration-200', file.status === 'error' ? 'bg-red-500' : 'bg-green-500']"></div>
            </div>
            <p class="text-xs text-right text-gray-600">{{ file.progress }}%</p>
            <div class="flex items-center justify-between mt-2">
              <button :class="['px-4 py-2 rounded font-bold transition-all', (saveStatus[file.id]==='saved') ? 'bg-green-400 hover:bg-green-500 text-white' : 'bg-blue-500 hover:bg-blue-600 text-white']" :disabled="file.status !== 'completed'" @click="handleSave(file)">ذخیره اطلاعات</button>
              <div class="flex items-center gap-2">
                <span v-if="file.status === 'completed'" class="text-green-600 font-bold">تکمیل شد</span>
                <span v-else-if="file.status === 'uploading'" class="text-blue-600 font-bold">در حال آپلود...</span>
                <span v-else-if="file.status === 'error'" class="text-red-600 font-bold">خطا: {{ file.error }}</span>
                <button class="text-red-500 hover:text-red-700 flex items-center gap-1" title="حذف" @click="removeFromQueue(file.id)">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path></svg>
                  <span class="text-xs">حذف</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="recentUploads.length > 0" class="bg-white rounded-lg border border-gray-200 overflow-hidden w-full max-w-[1600px] mx-auto">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">آپلودهای اخیر</h2>
        </div>

        <div class="grid gap-6 p-6 [grid-template-columns:repeat(auto-fit,minmax(220px,1fr))]">
          <div v-for="file in recentUploads" :key="file.id" class="group">
            <div class="aspect-square bg-gray-100 rounded-lg overflow-hidden mb-2 relative">
              <img
                  v-if="file.type.startsWith('image')"
                  :src="file.thumbnail"
                  :alt="file.name"
                  class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-200"
              >
              <div v-else class="w-full h-full flex items-center justify-center">
                <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
              </div>

              <div class="absolute inset-0 bg-black bg-opacity-50 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                <button
                    class="bg-white text-gray-900 p-2 rounded-lg hover:bg-gray-100 transition-colors"
                    @click="viewFile(file)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                </button>
              </div>
            </div>
            <p class="text-sm font-medium text-gray-900 truncate" :title="file.name">{{ file.name }}</p>
            <p class="text-xs text-gray-500">{{ formatFileSize(file.size) }}</p>
          </div>
        </div>
      </div>

      <div v-if="itemsToShow < filteredFiles.length" class="flex justify-center my-6">
        <button class="bg-red-500 text-white px-6 py-2 rounded-lg hover:bg-red-600 transition-colors" @click="itemsToShow += 30">
          بارگذاری بیشتر
        </button>
      </div>
    </div>

    <MediaPreviewModal
      v-if="previewModalVisible && previewModalFile"
      :file="previewModalFile"
      @close="closePreviewModal"
      @save="onModalSave"
      @delete="onModalDelete"
    />
    <ImagePreviewModal v-if="imgPreviewVisible" v-model="imgPreviewVisible" :image="imgPreviewFile" />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue'
import MediaPreviewModal from '~/components/media/MediaPreviewModal.vue'

import { useRouter } from 'vue-router'

interface UploadItem {
  id: string
  file: File
  name: string
  size: number
  type: string
  status: 'pending' | 'uploading' | 'completed' | 'error' | 'paused'
  progress: number
  speed: number
  error?: string
  preview?: string
  url?: string
  thumbnail?: string
  alt?: string
  title?: string
  description?: string
  caption?: string
  extension?: string
  uploadDate?: string
  author?: string
  width?: number
  height?: number
  category?: string
  compressed_size?: number | null
  serverId?: number | null
  [key: string]: unknown
}

interface RecentFile {
  id: number
  name: string
  url: string
  thumbnail: string
  type: string
  size: number
  author: string
  uploadDate: string
  extension: string
  width?: number
  height?: number
}

interface SaveStatusMap {
  [id: string]: 'saved-anim' | 'saved' | undefined;
}

const saveStatus: SaveStatusMap = reactive({});

const selectedCategory = ref('library')

const _categoryChecks = reactive({
  library: false,
  products: false
})

declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'افزودن رسانه جدید - پنل مدیریت'
})

// Reactive data
const isDragOver = ref(false)
const uploadQueue = ref<UploadItem[]>([])
const recentUploads = ref<RecentFile[]>([])
const _allPaused = ref(false)
const _modalImage = ref<string|null>(null)
const _showModal = ref(false)
// NEW: state for simple image preview
const imgPreviewVisible = ref(false)
interface PreviewFileData {
  [key: string]: unknown
}
const imgPreviewFile = ref<PreviewFileData | null>(null)
const _selectedFile = ref<UploadItem | RecentFile | null>(null)
const itemsToShow = ref(30)
const autoCompress = ref(true)
const previewModalFile = ref<UploadItem|null>(null)
const previewModalVisible = ref(false)


const _router = useRouter()

const filteredFiles = computed(() => uploadQueue.value)

// NEW: overall progress computed
const totalProgress = computed(() => {
  if (uploadQueue.value.length === 0) return 0
  const total = uploadQueue.value.reduce((sum, f) => sum + (f.progress || 0), 0)
  return Math.round(total / uploadQueue.value.length)
})

// NEW: progress of the file currently uploading (first in queue with status uploading)
const currentFileProgress = computed(() => {
  const uploading = uploadQueue.value.find(f => f.status === 'uploading')
  if (uploading) return uploading.progress || 0
  // اگر فایلی در حال آپلود نیست، اگر همه تکمیل شده‌اند 100، در غیر اینصورت 0
  if (uploadQueue.value.length && uploadQueue.value.every(f => f.status === 'completed')) return 100
  return 0
})

// NEW: helper to actually upload a single file with progress tracking
const uploadFile = (item: UploadItem) => {
  return new Promise<void>((resolve, reject) => {
    const formData = new FormData()
    formData.append('file', item.file)
    formData.append('category', selectedCategory.value)
    formData.append('author', '')
    formData.append('uploaded_by', '')
    formData.append('compress', autoCompress.value ? '1' : '0')

    const xhr = new XMLHttpRequest()
    // اگر کاربر فشرده‌سازی را فعال کرده باشد، علاوه بر پارامتر compress، فلگ smart=true را نیز در QueryString قرار می‌دهیم تا سمت سرور منطق هوشمند اجرا شود.
    const smartQuery = autoCompress.value ? '&smart=true' : ''
    xhr.open('POST', `/api/media/upload?category=${selectedCategory.value}${smartQuery}`)



    xhr.upload.onprogress = (event) => {
      /*
       * در برخی مرورگرها (به‌خصوص هنگام استفاده از proxy یا HTTP/2)‎
       * ‎lengthComputable ممکن است «false» شود و در نتیجه ‎total‎ صفر بماند.
       * ‎برای جلوگیری از گیرکردن در حالت «در حال آپلود…»‎، اگر ‎lengthComputable‎
       * ‎درست نبود، از حجم فایل روی کلاینت برای محاسبه درصد استفاده می‌کنیم.
       */
      const total = event.lengthComputable ? event.total : item.file.size
      if (total) {
        item.progress = Math.min(100, Math.round((event.loaded / total) * 100))
      }

      // در صورتی که به هر شکلی ۱۰۰٪ شد اما رویداد onload فراخوانی نشد،
      // حداکثر بعد از ۵۰۰ میلی‌ثانیه وضعیت را کامل اعلام می‌کنیم.
      if (item.progress >= 100 && item.status === 'uploading') {
          setTimeout(() => {
            if (item.status === 'uploading') {
              item.status = 'completed'
              // علامت‌گذاری برای refresh کردن کتابخانه رسانه
              sessionStorage.setItem('refreshMediaLibrary', 'true')
            }
        }, 500)
      }
    }

    xhr.onreadystatechange = () => {
      if (xhr.readyState === 4) {
        if (item.status === 'uploading') {
          if (xhr.status >= 200 && xhr.status < 300) {
            item.status = 'completed'
            item.progress = 100
          } else {
            item.status = 'error'
            item.error = 'خطا در آپلود'
          }
        }
      }
    }

    xhr.onload = () => {
      try {
        const isJson = xhr.getResponseHeader('content-type')?.includes('application/json')
        const result = isJson ? JSON.parse(xhr.responseText || '{}') : {}
        if (xhr.status >= 200 && xhr.status < 300) {
          item.serverId = result?.files?.[0]?.id || item.serverId
          item.url = result?.files?.[0]?.url || item.url
          item.preview = result?.files?.[0]?.url || item.preview
          item.compressed_size = result?.files?.[0]?.compressed_size || item.compressed_size
          item.status = 'completed'
          item.progress = 100
          // علامت‌گذاری برای refresh کردن کتابخانه رسانه
          sessionStorage.setItem('refreshMediaLibrary', 'true')
          resolve()
        } else {
          throw new Error(result.error || 'خطا در آپلود')
        }
      } catch (e: unknown) {
        item.status = 'error'
        item.error = e instanceof Error ? e.message : 'خطای ناشناخته'
        reject(e)
      }
    }

    xhr.onerror = () => {
      item.status = 'error'
      item.error = 'خطا در اتصال'
      reject(new Error('Network Error'))
    }

    xhr.onloadend = () => {
      if (item.status === 'uploading') {
        if (xhr.status >= 200 && xhr.status < 300) {
          item.status = 'completed'
          item.progress = 100
          // علامت‌گذاری برای refresh کردن کتابخانه رسانه
          sessionStorage.setItem('refreshMediaLibrary', 'true')
        } else {
          item.status = 'error'
          item.error = xhr.statusText || 'خطا در آپلود'
        }
      }
    }

    xhr.send(formData)
  })
}

// File handling
const handleDrop = (event: DragEvent) => {
  event.preventDefault()
  isDragOver.value = false

  const files = Array.from(event.dataTransfer?.files || [])
  processFiles(files)
}

// REPLACE handleFileSelect implementation
const handleFileSelect = (event: Event) => {
  const input = event.target as HTMLInputElement
  if (!input.files?.length) return

  const files = Array.from(input.files)
  for (const file of files) {
    const uploadItem: UploadItem = reactive({
      id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
      file,
      name: file.name,
      size: file.size,
      type: file.type,
      status: 'uploading',
      progress: 0,
      speed: 0,
      error: undefined,
      description: '',
      alt: '',
      title: '',
      preview: '',
      caption: '',
      extension: file.name.split('.').pop()?.toLowerCase() || '',
      uploadDate: new Date().toISOString(),
      author: 'ادمین',
      category: selectedCategory.value,
      compressed_size: null,
      serverId: null,
    })
    uploadQueue.value.push(uploadItem)
    uploadFile(uploadItem)
  }
  input.value = '' // Reset input
}

// REPLACE processFiles implementation
const processFiles = (files: File[]) => {
  files.forEach(file => {
    // Validate file size (1GB limit)
    if (file.size > 1024 * 1024 * 1024) {
      alert(`فایل ${file.name} بیش از 1 گیگابایت است و قابل آپلود نیست.`)
      return
    }

    // Validate file type
    const allowedTypes = [
      'image/jpeg', 'image/png', 'image/gif', 'image/webp',
      'video/mp4', 'video/webm', 'video/avi',
      'audio/mp3', 'audio/wav', 'audio/ogg',
      'application/pdf', 'application/msword', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
      'text/plain'
    ]

    if (!allowedTypes.includes(file.type)) {
      alert(`فایل ${file.name} با فرمت ${file.type} قابل آپلود نیست.`)
      return
    }

    const uploadItem: UploadItem = reactive({
      id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
      file,
      name: file.name,
      size: file.size,
      type: file.type,
      status: 'uploading',
      progress: 0,
      speed: 0,
      error: undefined,
      description: '',
      alt: '',
      title: '',
      preview: '',
      caption: '',
      extension: file.name.split('.').pop()?.toLowerCase() || '',
      uploadDate: new Date().toISOString(),
      author: 'ادمین',
      category: selectedCategory.value,
      compressed_size: null,
      serverId: null,
    })

    uploadQueue.value.push(uploadItem)
    uploadFile(uploadItem)
  })
}

const handleSave = async (file: UploadItem) => {
  if (!file || !file.id) return;
  try {
    const headers: Record<string,string> = { 'Content-Type': 'application/json' }

    const res = await fetch(`/api/media/${file.serverId || file.id}`, {
      method: 'PUT',
      headers,
      body: JSON.stringify({
        title: file.title,
        caption: file.caption,
        alt_text: file.alt,
        description: file.description,
        mime_type: file.type,
        author: file.author || ''
      })
    });
    if (!res.ok) throw new Error('خطا در ذخیره اطلاعات');
    saveStatus[file.id] = 'saved';
  } catch (_e) {
    alert('خطا در ذخیره اطلاعات');
  }
};

const markDirty = (file: UploadItem) => {
  if (saveStatus[file.id] === 'saved') {
    saveStatus[file.id] = undefined
  }
}

const removeFromQueue = async (id: string, skipConfirm: boolean = false) => {
  if (!skipConfirm && !confirm('آیا از حذف این فایل اطمینان دارید؟')) return;
  
  const index = uploadQueue.value.findIndex(item => item.id === id)
  if (index === -1) return

  const item = uploadQueue.value[index]

  // اگر فایل روی سرور ذخیره شده است، ابتدا درخواست حذف بفرستیم
  if (item.serverId) {
    try {
      const res = await fetch(`/api/media/delete?id=${item.serverId}`, {
        method: 'DELETE',
  
      })
      if (!res.ok) {
        throw new Error('حذف روی سرور ناموفق بود')
      }
    } catch (e) {
      alert((e as Error).message || 'خطا در حذف فایل')
      return // اگر حذف سرور ناموفق، آیتم را در UI نگه می‌داریم
    }
  }

    uploadQueue.value.splice(index, 1)
}

const _openPreviewModal = (file: UploadItem) => {
  previewModalFile.value = file
  previewModalVisible.value = true
}

// NEW: open simple preview modal (ImagePreviewModal)
function openImagePreview(file: UploadItem) {
  imgPreviewFile.value = file
  imgPreviewVisible.value = true
}

function closePreviewModal() {
  previewModalVisible.value = false
  previewModalFile.value = null
}

interface ModalSavePayload {
  alt?: string
  title?: string
  caption?: string
  description?: string
  id: string
}

async function onModalSave(payload: ModalSavePayload) {
  // payload = {alt,title,caption,description,id}
  const item = uploadQueue.value.find(f => f.id === payload.id)
  if (item) {
    item.alt = payload.alt
    item.title = payload.title
    item.caption = payload.caption
    item.description = payload.description
    await handleSave(item)
  }
}

async function onModalDelete(id: string | number) {
  // پیدا کردن فایل در upload queue بر اساس id
  const item = uploadQueue.value.find(f => f.id === id.toString() || f.serverId === id)
  if (item) {
    await removeFromQueue(item.id, true) // skipConfirm = true چون confirm در modal انجام شده
    closePreviewModal()
  }
}

// ---------------------------
// Utility / helper functions
// ---------------------------
// نمایش فایل در مودال پیش‌نمایش
const viewFile = (file: UploadItem | RecentFile) => {
  previewModalFile.value = file as UploadItem
  previewModalVisible.value = true
}

// تبدیل بایت به رشته قابل خواندن
const formatFileSize = (bytes: number): string => {
  if (!bytes) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>
