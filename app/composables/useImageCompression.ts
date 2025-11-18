import { reactive, ref, watch } from 'vue'

/*
کامپوزابل مدیریت فشرده سازی تصاویر (image compression)
این ماژول مسئول نگهداری state صفحه، فراخوانی API و فراهم کردن توابع کمکی برای UI است.
تمام کامپوننت‌های صفحه image-compression تنها از این کامپوزابل استفاده خواهند کرد.
*/

// ---------- Types ----------
export interface ImageFile {
  id: number
  name: string
  url: string
  thumbnail?: string
  size: number
  dimensions?: string
  extension?: string
  type?: string
  compressionStatus?: 'processing' | 'completed' | 'error'
  compressionResult?: {
    reduction: number
    newSize: number
  }
  preview?: string
}

export interface CompressionResult {
  id: number
  name: string
  thumbnail: string
  dimensions: string
  originalSize: number
  compressedSize: number
  reduction: number
  status: 'success' | 'error' | 'processing'
  compressedUrl?: string
  outputFormat?: string
  errorMessage?: string
  originalUrl?: string
}

// ---------- Global reactive state ----------
const images = ref<ImageFile[]>([])
const selectedImages = ref<number[]>([])
const isScanning = ref(false)
const scanPerformed = ref(false)

// Compression settings (برگرفته از منطق صفحه فعلی)
const compressionSettings = ref({
  quality: 'medium',
  customQuality: 75,
  format: 'original',
  progressive: true,
  removeMetadata: true,
  optimizeColors: true,
  createBackup: true,
  keepOriginalDate: true,
  autoOrient: true,
  enabled: true
})
const unsavedChanges = ref(false)
const lastLoadedSettings = ref<any>({})

// Compression queue
const compressionQueue = reactive({
  active: false,
  total: 0,
  inProgress: 0,
  completed: 0,
  errors: 0,
  progress: 0
})
const isCompressing = ref(false)
const compressionResults = ref<CompressionResult[]>([])

// Toast
const toast = reactive<{ message: string; type: 'success' | 'error' }>({ message: '', type: 'success' })
function showToast(message: string, type: 'success' | 'error' = 'success') {
  toast.message = message
  toast.type = type
  setTimeout(() => (toast.message = ''), 3000)
}

// Compare modal
const showCompareModal = ref(false)
const compareModalData = ref<CompressionResult | null>(null)
function openCompareModal(data: CompressionResult) {
  compareModalData.value = data
  showCompareModal.value = true
}
function closeCompareModal() {
  showCompareModal.value = false
  compareModalData.value = null
}

// -------- Compression trigger --------
function compressSelected() {
  if (selectedImages.value.length === 0) { showToast('هیچ تصویری انتخاب نشده است', 'error'); return }
  isCompressing.value = true
  showToast('عملیات فشرده‌سازی آغاز شد')
  // در نسخه فعلی شبیه‌سازی می‌کنیم؛ بعداً به API متصل می‌شود
  setTimeout(() => {
    isCompressing.value = false
    compressionQueue.total = selectedImages.value.length
    compressionQueue.completed = selectedImages.value.length
    compressionQueue.progress = 100
    showToast('فشرده‌سازی تکمیل شد')
  }, 1000)
}

// ---------- API helpers ----------
async function loadCompressionSettings() {
  try {
    const res = await $fetch('/api/admin/settings')
    if (Array.isArray(res)) {
      const map: Record<string, any> = {}
      res.forEach((s: any) => {
        map[s.key || s.Key] = s.value ?? s.Value
      })
      // mapping values
      const toBool = (v: any) => (typeof v === 'boolean' ? v : String(v).toLowerCase() === 'true' || v === '1')
      if (map['compression.quality']) compressionSettings.value.quality = map['compression.quality']
      if (map['compression.custom_quality']) compressionSettings.value.customQuality = Number(map['compression.custom_quality']) || 75
      if (map['compression.format']) compressionSettings.value.format = map['compression.format']
      compressionSettings.value.progressive = toBool(map['compression.progressive'])
      compressionSettings.value.removeMetadata = toBool(map['compression.remove_metadata'])
      compressionSettings.value.optimizeColors = toBool(map['compression.optimize_colors'])
      compressionSettings.value.createBackup = toBool(map['compression.create_backup'])
      compressionSettings.value.keepOriginalDate = toBool(map['compression.keep_original_date'])
      compressionSettings.value.autoOrient = toBool(map['compression.auto_orient'])
      compressionSettings.value.enabled = toBool(map['compression.enabled'])
      lastLoadedSettings.value = { ...compressionSettings.value }
    }
    unsavedChanges.value = false
  } catch (e) {
    // silent
  }
}

async function saveCompressionSettings() {
  const payload = [
    { key: 'compression.quality', value: compressionSettings.value.quality, category: 'compression', type: 'string' },
    { key: 'compression.custom_quality', value: String(compressionSettings.value.customQuality), category: 'compression', type: 'number' },
    { key: 'compression.format', value: compressionSettings.value.format, category: 'compression', type: 'string' },
    { key: 'compression.progressive', value: String(compressionSettings.value.progressive), category: 'compression', type: 'boolean' },
    { key: 'compression.remove_metadata', value: String(compressionSettings.value.removeMetadata), category: 'compression', type: 'boolean' },
    { key: 'compression.optimize_colors', value: String(compressionSettings.value.optimizeColors), category: 'compression', type: 'boolean' },
    { key: 'compression.create_backup', value: String(compressionSettings.value.createBackup), category: 'compression', type: 'boolean' },
    { key: 'compression.keep_original_date', value: String(compressionSettings.value.keepOriginalDate), category: 'compression', type: 'boolean' },
    { key: 'compression.auto_orient', value: String(compressionSettings.value.autoOrient), category: 'compression', type: 'boolean' },
    { key: 'compression.enabled', value: String(compressionSettings.value.enabled), category: 'compression', type: 'boolean' },
  ]
  try {
    await $fetch('/api/admin/settings', { method: 'PUT', body: payload })
    showToast('تنظیمات ذخیره شد')
    unsavedChanges.value = false
    lastLoadedSettings.value = { ...compressionSettings.value }
  } catch (e) {
    showToast('خطا در ذخیره تنظیمات', 'error')
  }
}

async function scanImages() {
  if (isScanning.value) return
  scanPerformed.value = true
  try {
    isScanning.value = true
    showToast('در حال اسکن تصاویر...')
    const res = await $fetch('/api/media/list')
    images.value = (res as any[]).map((it: any) => ({
      id: it.id,
      name: it.name,
      url: it.url,
      size: it.size,
      dimensions: it.dimensions
    }))
    showToast('اسکن به اتمام رسید')
  } catch (e) {
    showToast('خطا در اسکن تصاویر', 'error')
  } finally {
    isScanning.value = false
  }
}

// -------------- Watchers --------------
watch(compressionSettings, () => {
  unsavedChanges.value = true
}, { deep: true })

// ---------- Export API ----------
export function useImageCompression() {
  function handleImageSelect(id: number, event?: MouseEvent) {
    if (event?.ctrlKey || event?.metaKey) {
      if (selectedImages.value.includes(id)) {
        selectedImages.value = selectedImages.value.filter(i => i !== id)
      } else {
        selectedImages.value = [...selectedImages.value, id]
      }
    } else {
      selectedImages.value = [id]
    }
  }
  function selectAllImages() {
    selectedImages.value = images.value.map(i => i.id)
  }
  function clearSelection() { selectedImages.value = [] }

  return {
    // state
    images,
    selectedImages,
    isScanning,
    scanPerformed,
    compressionSettings,
    unsavedChanges,
    compressionQueue,
    isCompressing,
    compressionResults,
    toast,
    showCompareModal,
    compareModalData,
    // methods
    loadCompressionSettings,
    saveCompressionSettings,
    scanImages,
    showToast,
    openCompareModal,
    closeCompareModal,
    handleImageSelect,
    selectAllImages,
    clearSelection,
    compressSelected,
  }
} 