<template>
  <div v-if="modelValue" class="fixed inset-0 bg-transparent flex items-center justify-center z-50" @click="emitClose">
    <div class="relative bg-blue-50 rounded-lg overflow-hidden w-full max-w-[65vw] flex flex-col md:flex-row-reverse ring-4 ring-purple-300" @click.stop>
      <!-- Image preview with actions below -->
      <div class="flex-1 flex flex-col items-center justify-start bg-gray-100 p-6">
        <div v-if="file && (file.url || file.preview)" ref="containerEl" class="relative" @mousemove="onMouseMove" @mouseleave="onMouseLeave">
          <template v-if="file.type === 'video'">
            <div class="w-full h-[540px] bg-gray-200 rounded-xl flex items-center justify-center overflow-hidden">
              <video
                v-if="file.preview || file.url"
                :src="file.preview || file.url"
                controls
                preload="metadata"
                playsinline
                class="w-full h-full object-contain rounded-xl"
                @error="videoError = true"
              >
                <source :src="file.preview || file.url" :type="videoMimeType" />
                مرورگر شما از پخش ویدیو پشتیبانی نمی‌کند.
              </video>
              <div v-else class="text-red-600 text-center w-full">آدرس ویدیو نامعتبر است یا فایل وجود ندارد.</div>
            </div>
          </template>
          <template v-else>
            <img ref="imgEl" :src="file.preview || file.url" class="max-h-[70vh] object-contain select-none" :alt="file.name" draggable="false" @load="updateImageMetrics" />
            <!-- lens -->
            <div v-if="magnifierEnabled && lensVisible" :style="lensStyle" class="pointer-events-none"></div>
          </template>
        </div>

        <!-- Global action + compression buttons (single row) -->
        <div class="mt-4 flex items-center justify-center gap-2 flex-wrap md:flex-nowrap">
          <button :class="['px-4 py-2 rounded text-sm font-semibold transition', dirty ? 'bg-blue-500 hover:bg-blue-600 text-white' : 'bg-green-400 text-white']" @click="emitSave">ذخیره</button>
          <button v-if="canDeleteMedia" class="px-4 py-2 rounded bg-red-100 hover:bg-red-200 text-red-600 text-sm font-semibold" @click="emitDelete">حذف</button>
          <button class="px-4 py-2 rounded bg-blue-100 hover:bg-blue-200 text-blue-700 text-sm font-semibold" @click="copyLink">کپی لینک</button>
          <button class="px-4 py-2 rounded bg-purple-100 hover:bg-purple-200 text-purple-700 text-sm font-semibold" @click="downloadFile">دانلود</button>
          <!-- magnify button -->
          <button :class="['px-3 py-2 rounded text-sm font-semibold', magnifierEnabled ? 'bg-gray-300 text-gray-800' : 'bg-gray-100 hover:bg-gray-200 text-gray-700']" title="ابزار ذره بین" @click="magnifierEnabled = !magnifierEnabled">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5a6 6 0 016 6c0 1.227-.368 2.366-1 3.318l4.707 4.707a1 1 0 11-1.414 1.414l-4.707-4.707A5.978 5.978 0 0111 17a6 6 0 110-12z" />
            </svg>
          </button>
          <select v-if="magnifierEnabled" v-model="zoom" class="px-2 py-1 border rounded text-sm bg-white ml-2">
            <option v-for="z in [2,3,4,5]" :key="z" :value="z">{{ z }}×</option>
          </select>

          <!-- Compression controls (shown only for images) -->
          <template v-if="file.type === 'image'">
            <!-- toggle compress controls -->
            <button class="px-2 py-1 bg-sky-100 hover:bg-sky-200 text-sky-700 rounded text-sm border ml-10" title="ابزار فشرده‌سازی" @click="showCompressControls = !showCompressControls">⚙</button>
            <template v-if="showCompressControls">
              <!-- زمانی که فایل هنوز فشرده نشده -->
              <template v-if="!isCompressed">
                <button class="px-3 py-2 bg-rose-100 hover:bg-rose-200 text-rose-700 rounded text-sm border" @click="decQuality">-</button>
                <span class="px-1 text-sm font-medium select-none">{{ quality }}%</span>
                <button class="px-3 py-2 bg-emerald-100 hover:bg-emerald-200 text-emerald-700 rounded text-sm border" @click="incQuality">+</button>
                <select v-model="selectedFormat" class="px-2 py-1 border rounded text-sm bg-white">
                  <option value="auto">خودکار</option>
                  <option value="webp">WebP</option>
                  <option value="jpeg">JPEG</option>
                  <option value="png">PNG</option>
                </select>
                <button class="px-4 py-2 bg-yellow-400 hover:bg-yellow-500 text-white rounded text-sm font-semibold" @click="emitCompress">فشرده سازی</button>
                <button class="px-4 py-2 bg-indigo-400 hover:bg-indigo-500 text-white rounded text-sm font-semibold" @click="emitSmartCompress">فشرده‌سازی هوشمند</button>
              </template>
              <!-- زمانی که قبلاً فشرده شده، فقط دکمه بازگردانی -->
              <button v-if="isCompressed" class="px-4 py-2 bg-red-400 hover:bg-red-500 text-white rounded text-sm font-semibold" @click="emitRevert">بازگردانی</button>
            </template>
          </template>
        </div>
        <!-- removed separate compression container -->
      </div>

      <!-- Details panel -->
      <div class="flex-none w-full md:max-w-[240px] lg:max-w-[300px] p-3 md:p-6 overflow-y-auto space-y-3 md:space-y-4">
        <h2 class="text-center text-lg font-bold mb-2">جزئیات پیوست</h2>

        <!-- Static info -->
        <div class="text-xs text-gray-700 space-y-1 mb-10 text-center">
          <p v-if="file.uploadDate || file.created_at || file.createdAt">تاریخ: {{ formatDate(file.uploadDate || file.created_at || file.createdAt) }}</p>
          <p v-if="uploader">بارگذاری شده توسط: {{ uploader }}</p>
          <p>نام فایل: {{ file.name }}</p>
          <p>فرمت فایل: {{ file.extension || file.mime_type || '' }}</p>
          <p v-if="file.size">حجم قبل از فشرده‌سازی: {{ formatFileSize(file.size) }}</p>
          <p v-if="file.compressed_size">حجم بعد از فشرده‌سازی: {{ formatFileSize(file.compressed_size) }}</p>
        </div>

        <!-- Editable fields -->
        <div style="margin-top:48px">
          <template v-if="file.type === 'image'">
            <div class="mt-4 mb-6">
              <label class="block text-sm mb-1">متن جایگزین</label>
              <input v-model="form.alt" type="text" class="w-full max-w-[280px] border rounded px-2 py-1 text-sm" />
            </div>
            <div>
              <label class="block text-sm mb-1">عنوان</label>
              <input v-model="form.title" type="text" class="w-full max-w-[280px] border rounded px-2 py-1 text-sm" />
            </div>
            <div class="mt-6">
              <label class="block text-sm mb-1">توضیحات مختصر (caption)</label>
              <textarea v-model="form.caption" rows="3" class="w-full max-w-[280px] border rounded px-2 py-1 text-sm" />
            </div>
            <div class="mt-6">
              <label class="block text-sm mb-1">توضیح</label>
              <textarea v-model="form.description" rows="3" class="w-full max-w-[280px] border rounded px-2 py-1 text-sm" />
            </div>
          </template>
          <template v-else-if="file.type === 'video'">
            <div class="mt-4 mb-6">
              <label class="block text-sm mb-1">نام فایل ویدیو</label>
              <div class="w-full max-w-[280px] border rounded px-2 py-1 text-sm bg-gray-50">{{ file.name }}</div>
            </div>
            <div class="mt-6">
              <label class="block text-sm mb-1">توضیحات ویدیو</label>
              <textarea v-model="form.description" rows="6" class="w-full max-w-[280px] border rounded px-2 py-1 text-sm bg-white"></textarea>
            </div>
            <div v-if="videoError" class="mt-4 text-red-600 text-center text-base">خطا در بارگذاری یا پخش ویدیو. لطفاً فایل را بررسی کنید.</div>
          </template>
        </div>

        <!-- Action & compression buttons moved under image to reduce height -->

      </div>

      <!-- Close button -->
      <button class="absolute left-4 top-6 bg-white bg-opacity-80 rounded-full p-1 shadow hover:bg-opacity-100 transition-all" @click="emitClose">
        <svg class="w-6 h-6 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  modelValue: Boolean,
  file: {
    type: Object,
    required: false,
    default: null
  },
  canDelete: {
    type: Boolean,
    default: true
  }
})
import { nextTick, reactive, computed, watch, ref } from 'vue'
import type { CSSProperties } from 'vue'

const emit = defineEmits(['update:modelValue', 'close', 'save', 'delete', 'compress', 'revert'])

// Computed برای چک کردن پرمیژن حذف (بدون اتکا به useAuth در این کامپوننت)
const canDeleteMedia = computed(() => props.canDelete)

function emitClose() {
  emit('update:modelValue', false)
  emit('close')
}


// Determine uploader string
const uploader = computed(() => {
  if (!props.file) return ''
  // prefer explicit uploader_name then author
  return props.file.uploader_name || props.file.author || props.file.uploader_username || ''
})

const form = reactive({
  alt: '',
  title: '',
  caption: '',
  description: ''
})

// initialise form when file changes
watch(() => props.file, (newFile: any) => {
  if (!newFile) return
  form.alt = newFile.alt || ''
  form.title = newFile.title || ''
  form.caption = newFile.caption || ''
  form.description = newFile.description || ''
}, { immediate: true })

const dirty = computed(() => {
  if (!props.file) return false
  return form.alt !== (props.file.alt || '') ||
         form.title !== (props.file.title || '') ||
         form.caption !== (props.file.caption || '') ||
         form.description !== (props.file.description || '')
})

function emitSave() {
  if (!props.file || !dirty.value) return
  const payload = {
    id: props.file.id,
    alt: form.alt,
    title: form.title,
    caption: form.caption,
    description: form.description
  }

  // Send PUT request to backend
  fetch(`/api/media/${props.file.id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  }).catch(err => console.error('media update error', err))

  emit('save', payload)
}

function emitDelete() {
  if (!props.file) return
  
  // Emit delete event and let parent handle confirmation
  emit('delete', props.file.id)
}

function copyLink() {
  if (props.file && props.file.url) {
    navigator.clipboard.writeText(props.file.url)
    // optional toast
  }
}

function refreshImage() {
  // force reload image
  if (props.file && (props.file.url || props.file.preview)) {
    const img = new Image()
    img.src = (props.file.url || props.file.preview) + '?t=' + Date.now()
  }
}

function downloadFile() {
  if (!props.file || !props.file.url) return
  const a = document.createElement('a')
  a.href = props.file.url
  a.download = props.file.name || 'download'
  a.target = '_blank'
  a.click()
}

// define quality
const quality = ref(75)

function incQuality(){ if(quality.value < 100) quality.value += 5 }
function decQuality(){ if(quality.value > 5) quality.value -= 5 }

// reactive format selection (auto/webp/jpeg/png)
const selectedFormat = ref<'auto' | 'webp' | 'jpeg' | 'png'>('auto')

// toggle state
const showCompressControls = ref(false)

const magnifierEnabled = ref(false)
const lensVisible = ref(false)
const pointerInsideImage = ref(false)
const lensX = ref(0) // container-relative center X
const lensY = ref(0) // container-relative center Y
const imgRelX = ref(0) // image-relative center X
const imgRelY = ref(0) // image-relative center Y
const imgWidth = ref(0)
const imgHeight = ref(0)
const zoom = ref(2) // magnification factor
function setZoom(val) {
  zoom.value = Number(val)
}
const lensSize = 300
const effectiveLensSize = computed(() => {
  const w = imgWidth.value || 0
  const h = imgHeight.value || 0
  if (!w || !h) return lensSize
  return Math.min(lensSize, w, h)
})
const imgEl = ref<HTMLImageElement|null>(null)
const containerEl = ref<HTMLElement|null>(null)

const lensStyle = computed<CSSProperties>(() => {
  const size = effectiveLensSize.value
  const z = Number(zoom.value)
  // Allow full travel: background-position ranges from +size/2 (when x=0) to -(img*w*z - size/2)
  const maxBgX = size / 2
  const maxBgY = size / 2
  const minBgX = -(imgWidth.value * z - size / 2)
  const minBgY = -(imgHeight.value * z - size / 2)
  const rawBgX = size / 2 - (imgRelX.value * z)
  const rawBgY = size / 2 - (imgRelY.value * z)
  const bgX = Math.max(minBgX, Math.min(maxBgX, rawBgX))
  const bgY = Math.max(minBgY, Math.min(maxBgY, rawBgY))
  return {
    width: size + 'px',
    height: size + 'px',
    position: 'absolute',
    left: (lensX.value - size/2)+'px',
    top: (lensY.value - size/2)+'px',
    border: '2px solid #888',
    borderRadius: '50%',
    backgroundImage: pointerInsideImage.value && props.file ? `url(${props.file.preview || props.file.url})` : 'none',
    backgroundRepeat: 'no-repeat',
    backgroundSize: `${imgWidth.value*z}px ${imgHeight.value*z}px`,
    backgroundPosition: `${bgX}px ${bgY}px`,
    boxShadow: '0 0 4px rgba(0,0,0,0.3)',
    zIndex: 10,
    willChange: 'background-position,left,top'
  } as CSSProperties
})

let rafScheduled = false
let lastClientX = 0
let lastClientY = 0
function onMouseMove(e: MouseEvent) {
  if (!magnifierEnabled.value) return
  lastClientX = e.clientX
  lastClientY = e.clientY
  if (rafScheduled) return
  rafScheduled = true
  requestAnimationFrame(() => {
    rafScheduled = false
    const container = containerEl.value
    const img = imgEl.value
    if (!container || !img) return
    const rect = container.getBoundingClientRect()
    const imgRect = img.getBoundingClientRect()
    const x = lastClientX - rect.left
    const y = lastClientY - rect.top
    const size = effectiveLensSize.value
    // allow lens to move outside image; only background coords are clamped separately
    const offsetX = imgRect.left - rect.left
    const offsetY = imgRect.top - rect.top
    const imgCenterX = Math.max(0, Math.min(imgRect.width, x - offsetX))
    const imgCenterY = Math.max(0, Math.min(imgRect.height, y - offsetY))
    lensX.value = x
    lensY.value = y
    imgRelX.value = imgCenterX
    imgRelY.value = imgCenterY
    pointerInsideImage.value =
      lastClientX >= imgRect.left && lastClientX <= imgRect.right &&
      lastClientY >= imgRect.top && lastClientY <= imgRect.bottom
    lensVisible.value = true
  })
}

function onMouseLeave() {
  lensVisible.value = false
}

watch([imgEl, () => props.file], async () => {
  await nextTick()
  if (imgEl.value) {
    imgWidth.value = imgEl.value.clientWidth
    imgHeight.value = imgEl.value.clientHeight
  }
}, { immediate: true })

function updateImageMetrics() {
  if (imgEl.value) {
    imgWidth.value = imgEl.value.clientWidth
    imgHeight.value = imgEl.value.clientHeight
  }
}

const isCompressed = computed(() => {
  if (!props.file) return false
  if (props.file.compressed === true) return true
  if (props.file.compressed_size && props.file.size) {
    return props.file.compressed_size < props.file.size
  }
  return false
})

function emitCompress(){ if(!props.file) return; emit('compress', { id: props.file.id, quality: quality.value, format: selectedFormat.value === 'auto' ? undefined : selectedFormat.value }) }

function emitSmartCompress(){
  if(!props.file) return;
  emit('compress', { id: props.file.id, smart: true, format: selectedFormat.value === 'auto' ? undefined : selectedFormat.value })
}

function emitRevert(){ if(!props.file) return; emit('revert', props.file.id) }

function formatFileSize(bytes: number): string {
  if (!bytes) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

function formatDate(dateStr: string) {
  if (!dateStr) return ''
  const d = new Date(dateStr)
  return d.toLocaleString('fa-IR')
}

const videoError = ref(false)
const videoMimeType = computed(() => {
  if (!props.file) return 'video/mp4'
  const t = (props.file.mime_type || '').toLowerCase()
  if (t.startsWith('video/')) return t
  const url = props.file.preview || props.file.url || ''
  if (/\.webm(\?|#|$)/i.test(url)) return 'video/webm'
  if (/\.ogv(\?|#|$)/i.test(url)) return 'video/ogg'
  if (/\.m4v(\?|#|$)/i.test(url)) return 'video/mp4'
  if (/\.mov(\?|#|$)/i.test(url)) return 'video/quicktime'
  return 'video/mp4'
})

// reset videoError when file changes
watch(() => props.file, () => { videoError.value = false })
</script> 
